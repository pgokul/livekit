package service

import (
	"context"
	"time"

	"github.com/livekit/livekit-server/pkg/config"
	"github.com/livekit/livekit-server/pkg/service/rpc"
	"github.com/livekit/livekit-server/pkg/telemetry"
	"github.com/livekit/protocol/ingress"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
	"github.com/livekit/psrpc"
)

var (
	initialTimeout = time.Second * 3
	retryTimeout   = time.Minute * 1
)

type IngressService struct {
	conf        *config.IngressConfig
	nodeID      livekit.NodeID
	bus         psrpc.MessageBus
	psrpcClient rpc.IngressClient
	rpcClient   ingress.RPCClient
	store       IngressStore
	roomService livekit.RoomService
	telemetry   telemetry.TelemetryService
}

func NewIngressService(
	conf *config.IngressConfig,
	nodeID livekit.NodeID,
	bus psrpc.MessageBus,
	psrpcClient rpc.IngressClient,
	rpcClient ingress.RPCClient,
	store IngressStore,
	rs livekit.RoomService,
	ts telemetry.TelemetryService,
) *IngressService {

	return &IngressService{
		conf:        conf,
		nodeID:      nodeID,
		bus:         bus,
		psrpcClient: psrpcClient,
		rpcClient:   rpcClient,
		store:       store,
		roomService: rs,
		telemetry:   ts,
	}
}

func (s *IngressService) CreateIngress(ctx context.Context, req *livekit.CreateIngressRequest) (*livekit.IngressInfo, error) {
	fields := []interface{}{
		"inputType", req.InputType,
		"name", req.Name,
	}
	if req.RoomName != "" {
		fields = append(fields, "room", req.RoomName, "identity", req.ParticipantIdentity)
	}
	defer func() {
		AppendLogFields(ctx, fields...)
	}()

	ig, err := s.CreateIngressWithUrlPrefix(ctx, s.conf.RTMPBaseURL, req)
	if err != nil {
		return nil, err
	}
	fields = append(fields, "ingressID", ig.IngressId)

	return ig, nil
}

func (s *IngressService) CreateIngressWithUrlPrefix(ctx context.Context, urlPrefix string, req *livekit.CreateIngressRequest) (*livekit.IngressInfo, error) {
	err := EnsureIngressAdminPermission(ctx)
	if err != nil {
		return nil, twirpAuthError(err)
	}
	if s.store == nil {
		return nil, ErrIngressNotConnected
	}

	sk := utils.NewGuid("")

	info := &livekit.IngressInfo{
		IngressId:           utils.NewGuid(utils.IngressPrefix),
		Name:                req.Name,
		StreamKey:           sk,
		Url:                 urlPrefix,
		InputType:           req.InputType,
		Audio:               req.Audio,
		Video:               req.Video,
		RoomName:            req.RoomName,
		ParticipantIdentity: req.ParticipantIdentity,
		ParticipantName:     req.ParticipantName,
		Reusable:            req.InputType == livekit.IngressInput_RTMP_INPUT,
		State:               &livekit.IngressState{},
	}

	if err := s.store.StoreIngress(ctx, info); err != nil {
		logger.Errorw("could not write ingress info", err)
		return nil, err
	}

	return info, nil
}

func (s *IngressService) sendRPCWithRetry(ctx context.Context, req *livekit.IngressRequest) (*livekit.IngressState, error) {
	type result struct {
		state *livekit.IngressState
		err   error
	}

	resChan := make(chan result, 1)

	go func() {
		cctx, _ := context.WithTimeout(context.Background(), retryTimeout)

		for {
			select {
			case <-cctx.Done():
				resChan <- result{nil, ingress.ErrNoResponse}
				return
			default:
			}

			s, err := s.rpcClient.SendRequest(cctx, req)
			if err != ingress.ErrNoResponse {
				resChan <- result{s, err}
				return
			}
		}
	}()

	select {
	case res := <-resChan:
		return res.state, res.err
	case <-time.After(initialTimeout):
		return nil, ingress.ErrNoResponse
	}
}

func (s *IngressService) UpdateIngress(ctx context.Context, req *livekit.UpdateIngressRequest) (*livekit.IngressInfo, error) {
	fields := []interface{}{
		"ingress", req.IngressId,
		"name", req.Name,
	}
	if req.RoomName != "" {
		fields = append(fields, "room", req.RoomName, "identity", req.ParticipantIdentity)
	}
	AppendLogFields(ctx, fields...)
	err := EnsureIngressAdminPermission(ctx)
	if err != nil {
		return nil, twirpAuthError(err)
	}

	if s.psrpcClient == nil && s.rpcClient == nil {
		return nil, ErrIngressNotConnected
	}

	info, err := s.store.LoadIngress(ctx, req.IngressId)
	if err != nil {
		logger.Errorw("could not load ingress info", err)
		return nil, err
	}

	switch info.State.Status {
	case livekit.IngressState_ENDPOINT_ERROR:
		info.State.Status = livekit.IngressState_ENDPOINT_INACTIVE
		err = s.store.UpdateIngressState(ctx, req.IngressId, info.State)
		if err != nil {
			logger.Warnw("could not store ingress state", err)
		}
		fallthrough

	case livekit.IngressState_ENDPOINT_INACTIVE:
		if req.Name != "" {
			info.Name = req.Name
		}
		if req.RoomName != "" {
			info.RoomName = req.RoomName
		}
		if req.ParticipantIdentity != "" {
			info.ParticipantIdentity = req.ParticipantIdentity
		}
		if req.ParticipantName != "" {
			info.ParticipantName = req.ParticipantName
		}
		if req.Audio != nil {
			info.Audio = req.Audio
		}
		if req.Video != nil {
			info.Video = req.Video
		}

	case livekit.IngressState_ENDPOINT_BUFFERING,
		livekit.IngressState_ENDPOINT_PUBLISHING:
		// Do not update store the returned state as the ingress service will do it
		race := rpc.NewRace[livekit.IngressState](ctx)
		if s.rpcClient != nil {
			race.Go(func(ctx context.Context) (*livekit.IngressState, error) {
				return s.sendRPCWithRetry(ctx, &livekit.IngressRequest{
					IngressId: req.IngressId,
					Request:   &livekit.IngressRequest_Update{Update: req},
				})
			})
		}
		if s.psrpcClient != nil {
			race.Go(func(ctx context.Context) (*livekit.IngressState, error) {
				return s.psrpcClient.UpdateIngress(ctx, req.IngressId, req)
			})
		}
		if _, _, err := race.Wait(); err != nil {
			logger.Warnw("could not update active ingress", err)
		}
	}

	err = s.store.UpdateIngress(ctx, info)
	if err != nil {
		logger.Errorw("could not update ingress info", err)
		return nil, err
	}

	return info, nil
}

func (s *IngressService) ListIngress(ctx context.Context, req *livekit.ListIngressRequest) (*livekit.ListIngressResponse, error) {
	AppendLogFields(ctx, "room", req.RoomName)
	err := EnsureIngressAdminPermission(ctx)
	if err != nil {
		return nil, twirpAuthError(err)
	}
	if s.store == nil {
		return nil, ErrIngressNotConnected
	}

	infos, err := s.store.ListIngress(ctx, livekit.RoomName(req.RoomName))
	if err != nil {
		logger.Errorw("could not list ingress info", err)
		return nil, err
	}

	return &livekit.ListIngressResponse{Items: infos}, nil
}

func (s *IngressService) DeleteIngress(ctx context.Context, req *livekit.DeleteIngressRequest) (*livekit.IngressInfo, error) {
	AppendLogFields(ctx, "ingressID", req.IngressId)
	if err := EnsureIngressAdminPermission(ctx); err != nil {
		return nil, twirpAuthError(err)
	}

	if s.psrpcClient == nil && s.rpcClient == nil {
		return nil, ErrIngressNotConnected
	}

	info, err := s.store.LoadIngress(ctx, req.IngressId)
	if err != nil {
		return nil, err
	}

	switch info.State.Status {
	case livekit.IngressState_ENDPOINT_BUFFERING,
		livekit.IngressState_ENDPOINT_PUBLISHING:
		race := rpc.NewRace[livekit.IngressState](ctx)
		if s.rpcClient != nil {
			race.Go(func(ctx context.Context) (*livekit.IngressState, error) {
				return s.sendRPCWithRetry(ctx, &livekit.IngressRequest{
					IngressId: req.IngressId,
					Request:   &livekit.IngressRequest_Delete{Delete: req},
				})
			})
		}
		if s.psrpcClient != nil {
			race.Go(func(ctx context.Context) (*livekit.IngressState, error) {
				return s.psrpcClient.DeleteIngress(ctx, req.IngressId, req)
			})
		}
		if _, _, err := race.Wait(); err != nil {
			logger.Warnw("could not stop active ingress", err)
		}
	}

	err = s.store.DeleteIngress(ctx, info)
	if err != nil {
		logger.Errorw("could not delete ingress info", err)
		return nil, err
	}

	info.State.Status = livekit.IngressState_ENDPOINT_INACTIVE
	return info, nil
}
