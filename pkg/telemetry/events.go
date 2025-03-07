package telemetry

import (
	"context"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/livekit/livekit-server/pkg/telemetry/prometheus"
	"github.com/livekit/protocol/livekit"
	"github.com/livekit/protocol/logger"
	"github.com/livekit/protocol/utils"
	"github.com/livekit/protocol/webhook"
)

func (t *telemetryService) NotifyEvent(ctx context.Context, event *livekit.WebhookEvent) {
	if t.notifier == nil {
		return
	}

	event.CreatedAt = time.Now().Unix()
	event.Id = utils.NewGuid("EV_")

	t.webhookPool.Submit(func() {
		if err := t.notifier.Notify(ctx, event); err != nil {
			logger.Warnw("failed to notify webhook", err, "event", event.Event)
		}
	})
}

func (t *telemetryService) RoomStarted(ctx context.Context, room *livekit.Room) {
	t.enqueue(func() {
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event: webhook.EventRoomStarted,
			Room:  room,
		})

		t.SendEvent(ctx, &livekit.AnalyticsEvent{
			Type:      livekit.AnalyticsEventType_ROOM_CREATED,
			Timestamp: &timestamppb.Timestamp{Seconds: room.CreationTime},
			Room:      room,
		})
	})
}

func (t *telemetryService) RoomEnded(ctx context.Context, room *livekit.Room) {
	t.enqueue(func() {
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event: webhook.EventRoomFinished,
			Room:  room,
		})

		t.SendEvent(ctx, &livekit.AnalyticsEvent{
			Type:      livekit.AnalyticsEventType_ROOM_ENDED,
			Timestamp: timestamppb.Now(),
			RoomId:    room.Sid,
			Room:      room,
		})
	})
}

func (t *telemetryService) ParticipantJoined(
	ctx context.Context,
	room *livekit.Room,
	participant *livekit.ParticipantInfo,
	clientInfo *livekit.ClientInfo,
	clientMeta *livekit.AnalyticsClientMeta,
	shouldSendEvent bool,
) {
	t.enqueue(func() {
		prometheus.IncrementParticipantJoin(1, true)
		prometheus.AddParticipant()

		t.createWorker(
			ctx,
			livekit.RoomID(room.Sid),
			livekit.RoomName(room.Name),
			livekit.ParticipantID(participant.Sid),
			livekit.ParticipantIdentity(participant.Identity),
		)

		if shouldSendEvent {
			ev := newParticipantEvent(livekit.AnalyticsEventType_PARTICIPANT_JOINED, room, participant)
			ev.ClientInfo = clientInfo
			ev.ClientMeta = clientMeta
			t.SendEvent(ctx, ev)
		}
	})
}

func (t *telemetryService) ParticipantActive(
	ctx context.Context,
	room *livekit.Room,
	participant *livekit.ParticipantInfo,
	clientMeta *livekit.AnalyticsClientMeta,
) {
	t.enqueue(func() {
		// consider participant joined only when they became active
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event:       webhook.EventParticipantJoined,
			Room:        room,
			Participant: participant,
		})

		worker, ok := t.getWorker(livekit.ParticipantID(participant.Sid))
		if !ok {
			// in case of session migration, we may not have seen a Join event take place.
			// we'd need to create the worker here before being able to process events
			worker = t.createWorker(
				ctx,
				livekit.RoomID(room.Sid),
				livekit.RoomName(room.Name),
				livekit.ParticipantID(participant.Sid),
				livekit.ParticipantIdentity(participant.Identity),
			)

			// need to also account for participant count
			prometheus.AddParticipant()
		}
		worker.SetConnected()

		ev := newParticipantEvent(livekit.AnalyticsEventType_PARTICIPANT_ACTIVE, room, participant)
		ev.ClientMeta = clientMeta
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) ParticipantResumed(
	ctx context.Context,
	room *livekit.Room,
	participant *livekit.ParticipantInfo,
	nodeID livekit.NodeID,
) {
	t.enqueue(func() {
		ev := newParticipantEvent(livekit.AnalyticsEventType_PARTICIPANT_RESUMED, room, participant)
		ev.ClientMeta = &livekit.AnalyticsClientMeta{
			Node: string(nodeID),
		}
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) ParticipantLeft(ctx context.Context,
	room *livekit.Room,
	participant *livekit.ParticipantInfo,
	shouldSendEvent bool,
) {
	t.enqueue(func() {
		isConnected := false
		hasWorker := false
		if worker, ok := t.getWorker(livekit.ParticipantID(participant.Sid)); ok {
			hasWorker = true
			isConnected = worker.IsConnected()
			worker.Close()
		}

		if hasWorker {
			// signifies we had incremented participant count
			prometheus.SubParticipant()
		}

		if isConnected && shouldSendEvent {
			t.NotifyEvent(ctx, &livekit.WebhookEvent{
				Event:       webhook.EventParticipantLeft,
				Room:        room,
				Participant: participant,
			})

			t.SendEvent(ctx, newParticipantEvent(livekit.AnalyticsEventType_PARTICIPANT_LEFT, room, participant))
		}
	})
}

func (t *telemetryService) TrackPublishRequested(
	ctx context.Context,
	participantID livekit.ParticipantID,
	identity livekit.ParticipantIdentity,
	track *livekit.TrackInfo,
) {
	t.enqueue(func() {
		prometheus.AddPublishAttempt(track.Type.String())
		room := t.getRoomDetails(participantID)
		ev := newTrackEvent(livekit.AnalyticsEventType_TRACK_PUBLISH_REQUESTED, room, participantID, track)
		if ev.Participant != nil {
			ev.Participant.Identity = string(identity)
		}
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) TrackPublished(
	ctx context.Context,
	participantID livekit.ParticipantID,
	identity livekit.ParticipantIdentity,
	track *livekit.TrackInfo,
) {
	t.enqueue(func() {
		prometheus.AddPublishedTrack(track.Type.String())
		prometheus.AddPublishSuccess(track.Type.String())

		room := t.getRoomDetails(participantID)
		participant := &livekit.ParticipantInfo{
			Sid:      string(participantID),
			Identity: string(identity),
		}
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event:       webhook.EventTrackPublished,
			Room:        room,
			Participant: participant,
			Track:       track,
		})

		ev := newTrackEvent(livekit.AnalyticsEventType_TRACK_PUBLISHED, room, participantID, track)
		ev.Participant = participant
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) TrackPublishedUpdate(ctx context.Context, participantID livekit.ParticipantID, track *livekit.TrackInfo) {
	t.enqueue(func() {
		room := t.getRoomDetails(participantID)
		t.SendEvent(ctx, newTrackEvent(livekit.AnalyticsEventType_TRACK_PUBLISHED_UPDATE, room, participantID, track))
	})
}

func (t *telemetryService) TrackMaxSubscribedVideoQuality(
	ctx context.Context,
	participantID livekit.ParticipantID,
	track *livekit.TrackInfo,
	mime string,
	maxQuality livekit.VideoQuality,
) {
	t.enqueue(func() {
		room := t.getRoomDetails(participantID)
		ev := newTrackEvent(livekit.AnalyticsEventType_TRACK_MAX_SUBSCRIBED_VIDEO_QUALITY, room, participantID, track)
		ev.MaxSubscribedVideoQuality = maxQuality
		ev.Mime = mime
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) TrackSubscribeRequested(
	ctx context.Context,
	participantID livekit.ParticipantID,
	track *livekit.TrackInfo,
	publisher *livekit.ParticipantInfo,
) {
	t.enqueue(func() {
		room := t.getRoomDetails(participantID)
		ev := newTrackEvent(livekit.AnalyticsEventType_TRACK_SUBSCRIBE_REQUESTED, room, participantID, track)
		ev.Publisher = publisher
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) TrackSubscribed(
	ctx context.Context,
	participantID livekit.ParticipantID,
	track *livekit.TrackInfo,
	publisher *livekit.ParticipantInfo,
) {
	t.enqueue(func() {
		prometheus.AddSubscribedTrack(track.Type.String())

		room := t.getRoomDetails(participantID)
		ev := newTrackEvent(livekit.AnalyticsEventType_TRACK_SUBSCRIBED, room, participantID, track)
		ev.Publisher = publisher
		t.SendEvent(ctx, ev)
	})
}

func (t *telemetryService) TrackUnsubscribed(ctx context.Context, participantID livekit.ParticipantID, track *livekit.TrackInfo) {
	t.enqueue(func() {
		prometheus.SubSubscribedTrack(track.Type.String())

		room := t.getRoomDetails(participantID)
		t.SendEvent(ctx, newTrackEvent(livekit.AnalyticsEventType_TRACK_UNSUBSCRIBED, room, participantID, track))
	})
}

func (t *telemetryService) TrackUnpublished(
	ctx context.Context,
	participantID livekit.ParticipantID,
	identity livekit.ParticipantIdentity,
	track *livekit.TrackInfo,
	shouldSendEvent bool,
) {
	t.enqueue(func() {
		prometheus.SubPublishedTrack(track.Type.String())
		if !shouldSendEvent {
			return
		}

		room := t.getRoomDetails(participantID)
		participant := &livekit.ParticipantInfo{
			Sid:      string(participantID),
			Identity: string(identity),
		}
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event:       webhook.EventTrackUnpublished,
			Room:        room,
			Participant: participant,
			Track:       track,
		})

		t.SendEvent(ctx, newTrackEvent(livekit.AnalyticsEventType_TRACK_UNPUBLISHED, room, participantID, track))
	})
}

func (t *telemetryService) TrackMuted(
	ctx context.Context,
	participantID livekit.ParticipantID,
	track *livekit.TrackInfo,
) {
	t.enqueue(func() {
		room := t.getRoomDetails(participantID)
		t.SendEvent(ctx, newTrackEvent(livekit.AnalyticsEventType_TRACK_MUTED, room, participantID, track))
	})
}

func (t *telemetryService) TrackUnmuted(
	ctx context.Context,
	participantID livekit.ParticipantID,
	track *livekit.TrackInfo,
) {
	t.enqueue(func() {
		room := t.getRoomDetails(participantID)
		t.SendEvent(ctx, newTrackEvent(livekit.AnalyticsEventType_TRACK_UNMUTED, room, participantID, track))
	})
}

func (t *telemetryService) EgressStarted(ctx context.Context, info *livekit.EgressInfo) {
	t.enqueue(func() {
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event:      webhook.EventEgressStarted,
			EgressInfo: info,
		})

		t.SendEvent(ctx, newEgressEvent(livekit.AnalyticsEventType_EGRESS_STARTED, info))
	})
}

func (t *telemetryService) EgressEnded(ctx context.Context, info *livekit.EgressInfo) {
	t.enqueue(func() {
		t.NotifyEvent(ctx, &livekit.WebhookEvent{
			Event:      webhook.EventEgressEnded,
			EgressInfo: info,
		})

		t.SendEvent(ctx, newEgressEvent(livekit.AnalyticsEventType_EGRESS_ENDED, info))
	})
}

// returns a livekit.Room with only name and sid filled out
// returns nil if room is not found
func (t *telemetryService) getRoomDetails(participantID livekit.ParticipantID) *livekit.Room {
	if worker, ok := t.getWorker(participantID); ok {
		return &livekit.Room{
			Sid:  string(worker.roomID),
			Name: string(worker.roomName),
		}
	}

	return nil
}

func newRoomEvent(event livekit.AnalyticsEventType, room *livekit.Room) *livekit.AnalyticsEvent {
	ev := &livekit.AnalyticsEvent{
		Type:      event,
		Timestamp: timestamppb.Now(),
	}
	if room != nil {
		ev.Room = room
		ev.RoomId = room.Sid
	}
	return ev
}

func newParticipantEvent(event livekit.AnalyticsEventType, room *livekit.Room, participant *livekit.ParticipantInfo) *livekit.AnalyticsEvent {
	ev := newRoomEvent(event, room)
	if participant != nil {
		ev.ParticipantId = participant.Sid
		ev.Participant = participant
	}
	return ev
}

func newTrackEvent(event livekit.AnalyticsEventType, room *livekit.Room, participantID livekit.ParticipantID, track *livekit.TrackInfo) *livekit.AnalyticsEvent {
	ev := newParticipantEvent(event, room, &livekit.ParticipantInfo{
		Sid: string(participantID),
	})
	if track != nil {
		ev.TrackId = track.Sid
		ev.Track = track
	}
	return ev
}

func newEgressEvent(event livekit.AnalyticsEventType, egress *livekit.EgressInfo) *livekit.AnalyticsEvent {
	return &livekit.AnalyticsEvent{
		Type:      event,
		Timestamp: timestamppb.Now(),
		EgressId:  egress.EgressId,
		RoomId:    egress.RoomId,
		Egress:    egress,
	}
}
