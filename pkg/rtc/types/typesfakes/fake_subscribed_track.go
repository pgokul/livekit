// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/livekit-server/pkg/sfu"
	"github.com/livekit/protocol/livekit"
)

type FakeSubscribedTrack struct {
	DownTrackStub        func() *sfu.DownTrack
	downTrackMutex       sync.RWMutex
	downTrackArgsForCall []struct {
	}
	downTrackReturns struct {
		result1 *sfu.DownTrack
	}
	downTrackReturnsOnCall map[int]struct {
		result1 *sfu.DownTrack
	}
	IDStub        func() livekit.TrackID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 livekit.TrackID
	}
	iDReturnsOnCall map[int]struct {
		result1 livekit.TrackID
	}
	IsMutedStub        func() bool
	isMutedMutex       sync.RWMutex
	isMutedArgsForCall []struct {
	}
	isMutedReturns struct {
		result1 bool
	}
	isMutedReturnsOnCall map[int]struct {
		result1 bool
	}
	MediaTrackStub        func() types.MediaTrack
	mediaTrackMutex       sync.RWMutex
	mediaTrackArgsForCall []struct {
	}
	mediaTrackReturns struct {
		result1 types.MediaTrack
	}
	mediaTrackReturnsOnCall map[int]struct {
		result1 types.MediaTrack
	}
	OnBindStub        func(func())
	onBindMutex       sync.RWMutex
	onBindArgsForCall []struct {
		arg1 func()
	}
	PublisherIDStub        func() livekit.ParticipantID
	publisherIDMutex       sync.RWMutex
	publisherIDArgsForCall []struct {
	}
	publisherIDReturns struct {
		result1 livekit.ParticipantID
	}
	publisherIDReturnsOnCall map[int]struct {
		result1 livekit.ParticipantID
	}
	PublisherIdentityStub        func() livekit.ParticipantIdentity
	publisherIdentityMutex       sync.RWMutex
	publisherIdentityArgsForCall []struct {
	}
	publisherIdentityReturns struct {
		result1 livekit.ParticipantIdentity
	}
	publisherIdentityReturnsOnCall map[int]struct {
		result1 livekit.ParticipantIdentity
	}
	PublisherVersionStub        func() uint32
	publisherVersionMutex       sync.RWMutex
	publisherVersionArgsForCall []struct {
	}
	publisherVersionReturns struct {
		result1 uint32
	}
	publisherVersionReturnsOnCall map[int]struct {
		result1 uint32
	}
	SetPublisherMutedStub        func(bool)
	setPublisherMutedMutex       sync.RWMutex
	setPublisherMutedArgsForCall []struct {
		arg1 bool
	}
	SubscriberStub        func() types.LocalParticipant
	subscriberMutex       sync.RWMutex
	subscriberArgsForCall []struct {
	}
	subscriberReturns struct {
		result1 types.LocalParticipant
	}
	subscriberReturnsOnCall map[int]struct {
		result1 types.LocalParticipant
	}
	SubscriberIDStub        func() livekit.ParticipantID
	subscriberIDMutex       sync.RWMutex
	subscriberIDArgsForCall []struct {
	}
	subscriberIDReturns struct {
		result1 livekit.ParticipantID
	}
	subscriberIDReturnsOnCall map[int]struct {
		result1 livekit.ParticipantID
	}
	SubscriberIdentityStub        func() livekit.ParticipantIdentity
	subscriberIdentityMutex       sync.RWMutex
	subscriberIdentityArgsForCall []struct {
	}
	subscriberIdentityReturns struct {
		result1 livekit.ParticipantIdentity
	}
	subscriberIdentityReturnsOnCall map[int]struct {
		result1 livekit.ParticipantIdentity
	}
	UpdateSubscriberSettingsStub        func(*livekit.UpdateTrackSettings)
	updateSubscriberSettingsMutex       sync.RWMutex
	updateSubscriberSettingsArgsForCall []struct {
		arg1 *livekit.UpdateTrackSettings
	}
	UpdateVideoLayerStub        func()
	updateVideoLayerMutex       sync.RWMutex
	updateVideoLayerArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSubscribedTrack) DownTrack() *sfu.DownTrack {
	fake.downTrackMutex.Lock()
	ret, specificReturn := fake.downTrackReturnsOnCall[len(fake.downTrackArgsForCall)]
	fake.downTrackArgsForCall = append(fake.downTrackArgsForCall, struct {
	}{})
	stub := fake.DownTrackStub
	fakeReturns := fake.downTrackReturns
	fake.recordInvocation("DownTrack", []interface{}{})
	fake.downTrackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) DownTrackCallCount() int {
	fake.downTrackMutex.RLock()
	defer fake.downTrackMutex.RUnlock()
	return len(fake.downTrackArgsForCall)
}

func (fake *FakeSubscribedTrack) DownTrackCalls(stub func() *sfu.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = stub
}

func (fake *FakeSubscribedTrack) DownTrackReturns(result1 *sfu.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = nil
	fake.downTrackReturns = struct {
		result1 *sfu.DownTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) DownTrackReturnsOnCall(i int, result1 *sfu.DownTrack) {
	fake.downTrackMutex.Lock()
	defer fake.downTrackMutex.Unlock()
	fake.DownTrackStub = nil
	if fake.downTrackReturnsOnCall == nil {
		fake.downTrackReturnsOnCall = make(map[int]struct {
			result1 *sfu.DownTrack
		})
	}
	fake.downTrackReturnsOnCall[i] = struct {
		result1 *sfu.DownTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) ID() livekit.TrackID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeSubscribedTrack) IDCalls(stub func() livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeSubscribedTrack) IDReturns(result1 livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 livekit.TrackID
	}{result1}
}

func (fake *FakeSubscribedTrack) IDReturnsOnCall(i int, result1 livekit.TrackID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 livekit.TrackID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 livekit.TrackID
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMuted() bool {
	fake.isMutedMutex.Lock()
	ret, specificReturn := fake.isMutedReturnsOnCall[len(fake.isMutedArgsForCall)]
	fake.isMutedArgsForCall = append(fake.isMutedArgsForCall, struct {
	}{})
	stub := fake.IsMutedStub
	fakeReturns := fake.isMutedReturns
	fake.recordInvocation("IsMuted", []interface{}{})
	fake.isMutedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) IsMutedCallCount() int {
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	return len(fake.isMutedArgsForCall)
}

func (fake *FakeSubscribedTrack) IsMutedCalls(stub func() bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = stub
}

func (fake *FakeSubscribedTrack) IsMutedReturns(result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	fake.isMutedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) IsMutedReturnsOnCall(i int, result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	if fake.isMutedReturnsOnCall == nil {
		fake.isMutedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMutedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSubscribedTrack) MediaTrack() types.MediaTrack {
	fake.mediaTrackMutex.Lock()
	ret, specificReturn := fake.mediaTrackReturnsOnCall[len(fake.mediaTrackArgsForCall)]
	fake.mediaTrackArgsForCall = append(fake.mediaTrackArgsForCall, struct {
	}{})
	stub := fake.MediaTrackStub
	fakeReturns := fake.mediaTrackReturns
	fake.recordInvocation("MediaTrack", []interface{}{})
	fake.mediaTrackMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) MediaTrackCallCount() int {
	fake.mediaTrackMutex.RLock()
	defer fake.mediaTrackMutex.RUnlock()
	return len(fake.mediaTrackArgsForCall)
}

func (fake *FakeSubscribedTrack) MediaTrackCalls(stub func() types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = stub
}

func (fake *FakeSubscribedTrack) MediaTrackReturns(result1 types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = nil
	fake.mediaTrackReturns = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) MediaTrackReturnsOnCall(i int, result1 types.MediaTrack) {
	fake.mediaTrackMutex.Lock()
	defer fake.mediaTrackMutex.Unlock()
	fake.MediaTrackStub = nil
	if fake.mediaTrackReturnsOnCall == nil {
		fake.mediaTrackReturnsOnCall = make(map[int]struct {
			result1 types.MediaTrack
		})
	}
	fake.mediaTrackReturnsOnCall[i] = struct {
		result1 types.MediaTrack
	}{result1}
}

func (fake *FakeSubscribedTrack) OnBind(arg1 func()) {
	fake.onBindMutex.Lock()
	fake.onBindArgsForCall = append(fake.onBindArgsForCall, struct {
		arg1 func()
	}{arg1})
	stub := fake.OnBindStub
	fake.recordInvocation("OnBind", []interface{}{arg1})
	fake.onBindMutex.Unlock()
	if stub != nil {
		fake.OnBindStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) OnBindCallCount() int {
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	return len(fake.onBindArgsForCall)
}

func (fake *FakeSubscribedTrack) OnBindCalls(stub func(func())) {
	fake.onBindMutex.Lock()
	defer fake.onBindMutex.Unlock()
	fake.OnBindStub = stub
}

func (fake *FakeSubscribedTrack) OnBindArgsForCall(i int) func() {
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	argsForCall := fake.onBindArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) PublisherID() livekit.ParticipantID {
	fake.publisherIDMutex.Lock()
	ret, specificReturn := fake.publisherIDReturnsOnCall[len(fake.publisherIDArgsForCall)]
	fake.publisherIDArgsForCall = append(fake.publisherIDArgsForCall, struct {
	}{})
	stub := fake.PublisherIDStub
	fakeReturns := fake.publisherIDReturns
	fake.recordInvocation("PublisherID", []interface{}{})
	fake.publisherIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherIDCallCount() int {
	fake.publisherIDMutex.RLock()
	defer fake.publisherIDMutex.RUnlock()
	return len(fake.publisherIDArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherIDCalls(stub func() livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = stub
}

func (fake *FakeSubscribedTrack) PublisherIDReturns(result1 livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = nil
	fake.publisherIDReturns = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIDReturnsOnCall(i int, result1 livekit.ParticipantID) {
	fake.publisherIDMutex.Lock()
	defer fake.publisherIDMutex.Unlock()
	fake.PublisherIDStub = nil
	if fake.publisherIDReturnsOnCall == nil {
		fake.publisherIDReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantID
		})
	}
	fake.publisherIDReturnsOnCall[i] = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIdentity() livekit.ParticipantIdentity {
	fake.publisherIdentityMutex.Lock()
	ret, specificReturn := fake.publisherIdentityReturnsOnCall[len(fake.publisherIdentityArgsForCall)]
	fake.publisherIdentityArgsForCall = append(fake.publisherIdentityArgsForCall, struct {
	}{})
	stub := fake.PublisherIdentityStub
	fakeReturns := fake.publisherIdentityReturns
	fake.recordInvocation("PublisherIdentity", []interface{}{})
	fake.publisherIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherIdentityCallCount() int {
	fake.publisherIdentityMutex.RLock()
	defer fake.publisherIdentityMutex.RUnlock()
	return len(fake.publisherIdentityArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherIdentityCalls(stub func() livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = stub
}

func (fake *FakeSubscribedTrack) PublisherIdentityReturns(result1 livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = nil
	fake.publisherIdentityReturns = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherIdentityReturnsOnCall(i int, result1 livekit.ParticipantIdentity) {
	fake.publisherIdentityMutex.Lock()
	defer fake.publisherIdentityMutex.Unlock()
	fake.PublisherIdentityStub = nil
	if fake.publisherIdentityReturnsOnCall == nil {
		fake.publisherIdentityReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantIdentity
		})
	}
	fake.publisherIdentityReturnsOnCall[i] = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherVersion() uint32 {
	fake.publisherVersionMutex.Lock()
	ret, specificReturn := fake.publisherVersionReturnsOnCall[len(fake.publisherVersionArgsForCall)]
	fake.publisherVersionArgsForCall = append(fake.publisherVersionArgsForCall, struct {
	}{})
	stub := fake.PublisherVersionStub
	fakeReturns := fake.publisherVersionReturns
	fake.recordInvocation("PublisherVersion", []interface{}{})
	fake.publisherVersionMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) PublisherVersionCallCount() int {
	fake.publisherVersionMutex.RLock()
	defer fake.publisherVersionMutex.RUnlock()
	return len(fake.publisherVersionArgsForCall)
}

func (fake *FakeSubscribedTrack) PublisherVersionCalls(stub func() uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = stub
}

func (fake *FakeSubscribedTrack) PublisherVersionReturns(result1 uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = nil
	fake.publisherVersionReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeSubscribedTrack) PublisherVersionReturnsOnCall(i int, result1 uint32) {
	fake.publisherVersionMutex.Lock()
	defer fake.publisherVersionMutex.Unlock()
	fake.PublisherVersionStub = nil
	if fake.publisherVersionReturnsOnCall == nil {
		fake.publisherVersionReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.publisherVersionReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeSubscribedTrack) SetPublisherMuted(arg1 bool) {
	fake.setPublisherMutedMutex.Lock()
	fake.setPublisherMutedArgsForCall = append(fake.setPublisherMutedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetPublisherMutedStub
	fake.recordInvocation("SetPublisherMuted", []interface{}{arg1})
	fake.setPublisherMutedMutex.Unlock()
	if stub != nil {
		fake.SetPublisherMutedStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) SetPublisherMutedCallCount() int {
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	return len(fake.setPublisherMutedArgsForCall)
}

func (fake *FakeSubscribedTrack) SetPublisherMutedCalls(stub func(bool)) {
	fake.setPublisherMutedMutex.Lock()
	defer fake.setPublisherMutedMutex.Unlock()
	fake.SetPublisherMutedStub = stub
}

func (fake *FakeSubscribedTrack) SetPublisherMutedArgsForCall(i int) bool {
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	argsForCall := fake.setPublisherMutedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) Subscriber() types.LocalParticipant {
	fake.subscriberMutex.Lock()
	ret, specificReturn := fake.subscriberReturnsOnCall[len(fake.subscriberArgsForCall)]
	fake.subscriberArgsForCall = append(fake.subscriberArgsForCall, struct {
	}{})
	stub := fake.SubscriberStub
	fakeReturns := fake.subscriberReturns
	fake.recordInvocation("Subscriber", []interface{}{})
	fake.subscriberMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberCallCount() int {
	fake.subscriberMutex.RLock()
	defer fake.subscriberMutex.RUnlock()
	return len(fake.subscriberArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberCalls(stub func() types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberReturns(result1 types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = nil
	fake.subscriberReturns = struct {
		result1 types.LocalParticipant
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberReturnsOnCall(i int, result1 types.LocalParticipant) {
	fake.subscriberMutex.Lock()
	defer fake.subscriberMutex.Unlock()
	fake.SubscriberStub = nil
	if fake.subscriberReturnsOnCall == nil {
		fake.subscriberReturnsOnCall = make(map[int]struct {
			result1 types.LocalParticipant
		})
	}
	fake.subscriberReturnsOnCall[i] = struct {
		result1 types.LocalParticipant
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberID() livekit.ParticipantID {
	fake.subscriberIDMutex.Lock()
	ret, specificReturn := fake.subscriberIDReturnsOnCall[len(fake.subscriberIDArgsForCall)]
	fake.subscriberIDArgsForCall = append(fake.subscriberIDArgsForCall, struct {
	}{})
	stub := fake.SubscriberIDStub
	fakeReturns := fake.subscriberIDReturns
	fake.recordInvocation("SubscriberID", []interface{}{})
	fake.subscriberIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberIDCallCount() int {
	fake.subscriberIDMutex.RLock()
	defer fake.subscriberIDMutex.RUnlock()
	return len(fake.subscriberIDArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberIDCalls(stub func() livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberIDReturns(result1 livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = nil
	fake.subscriberIDReturns = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIDReturnsOnCall(i int, result1 livekit.ParticipantID) {
	fake.subscriberIDMutex.Lock()
	defer fake.subscriberIDMutex.Unlock()
	fake.SubscriberIDStub = nil
	if fake.subscriberIDReturnsOnCall == nil {
		fake.subscriberIDReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantID
		})
	}
	fake.subscriberIDReturnsOnCall[i] = struct {
		result1 livekit.ParticipantID
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIdentity() livekit.ParticipantIdentity {
	fake.subscriberIdentityMutex.Lock()
	ret, specificReturn := fake.subscriberIdentityReturnsOnCall[len(fake.subscriberIdentityArgsForCall)]
	fake.subscriberIdentityArgsForCall = append(fake.subscriberIdentityArgsForCall, struct {
	}{})
	stub := fake.SubscriberIdentityStub
	fakeReturns := fake.subscriberIdentityReturns
	fake.recordInvocation("SubscriberIdentity", []interface{}{})
	fake.subscriberIdentityMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeSubscribedTrack) SubscriberIdentityCallCount() int {
	fake.subscriberIdentityMutex.RLock()
	defer fake.subscriberIdentityMutex.RUnlock()
	return len(fake.subscriberIdentityArgsForCall)
}

func (fake *FakeSubscribedTrack) SubscriberIdentityCalls(stub func() livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = stub
}

func (fake *FakeSubscribedTrack) SubscriberIdentityReturns(result1 livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = nil
	fake.subscriberIdentityReturns = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) SubscriberIdentityReturnsOnCall(i int, result1 livekit.ParticipantIdentity) {
	fake.subscriberIdentityMutex.Lock()
	defer fake.subscriberIdentityMutex.Unlock()
	fake.SubscriberIdentityStub = nil
	if fake.subscriberIdentityReturnsOnCall == nil {
		fake.subscriberIdentityReturnsOnCall = make(map[int]struct {
			result1 livekit.ParticipantIdentity
		})
	}
	fake.subscriberIdentityReturnsOnCall[i] = struct {
		result1 livekit.ParticipantIdentity
	}{result1}
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettings(arg1 *livekit.UpdateTrackSettings) {
	fake.updateSubscriberSettingsMutex.Lock()
	fake.updateSubscriberSettingsArgsForCall = append(fake.updateSubscriberSettingsArgsForCall, struct {
		arg1 *livekit.UpdateTrackSettings
	}{arg1})
	stub := fake.UpdateSubscriberSettingsStub
	fake.recordInvocation("UpdateSubscriberSettings", []interface{}{arg1})
	fake.updateSubscriberSettingsMutex.Unlock()
	if stub != nil {
		fake.UpdateSubscriberSettingsStub(arg1)
	}
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsCallCount() int {
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	return len(fake.updateSubscriberSettingsArgsForCall)
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsCalls(stub func(*livekit.UpdateTrackSettings)) {
	fake.updateSubscriberSettingsMutex.Lock()
	defer fake.updateSubscriberSettingsMutex.Unlock()
	fake.UpdateSubscriberSettingsStub = stub
}

func (fake *FakeSubscribedTrack) UpdateSubscriberSettingsArgsForCall(i int) *livekit.UpdateTrackSettings {
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	argsForCall := fake.updateSubscriberSettingsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeSubscribedTrack) UpdateVideoLayer() {
	fake.updateVideoLayerMutex.Lock()
	fake.updateVideoLayerArgsForCall = append(fake.updateVideoLayerArgsForCall, struct {
	}{})
	stub := fake.UpdateVideoLayerStub
	fake.recordInvocation("UpdateVideoLayer", []interface{}{})
	fake.updateVideoLayerMutex.Unlock()
	if stub != nil {
		fake.UpdateVideoLayerStub()
	}
}

func (fake *FakeSubscribedTrack) UpdateVideoLayerCallCount() int {
	fake.updateVideoLayerMutex.RLock()
	defer fake.updateVideoLayerMutex.RUnlock()
	return len(fake.updateVideoLayerArgsForCall)
}

func (fake *FakeSubscribedTrack) UpdateVideoLayerCalls(stub func()) {
	fake.updateVideoLayerMutex.Lock()
	defer fake.updateVideoLayerMutex.Unlock()
	fake.UpdateVideoLayerStub = stub
}

func (fake *FakeSubscribedTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.downTrackMutex.RLock()
	defer fake.downTrackMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	fake.mediaTrackMutex.RLock()
	defer fake.mediaTrackMutex.RUnlock()
	fake.onBindMutex.RLock()
	defer fake.onBindMutex.RUnlock()
	fake.publisherIDMutex.RLock()
	defer fake.publisherIDMutex.RUnlock()
	fake.publisherIdentityMutex.RLock()
	defer fake.publisherIdentityMutex.RUnlock()
	fake.publisherVersionMutex.RLock()
	defer fake.publisherVersionMutex.RUnlock()
	fake.setPublisherMutedMutex.RLock()
	defer fake.setPublisherMutedMutex.RUnlock()
	fake.subscriberMutex.RLock()
	defer fake.subscriberMutex.RUnlock()
	fake.subscriberIDMutex.RLock()
	defer fake.subscriberIDMutex.RUnlock()
	fake.subscriberIdentityMutex.RLock()
	defer fake.subscriberIdentityMutex.RUnlock()
	fake.updateSubscriberSettingsMutex.RLock()
	defer fake.updateSubscriberSettingsMutex.RUnlock()
	fake.updateVideoLayerMutex.RLock()
	defer fake.updateVideoLayerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSubscribedTrack) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.SubscribedTrack = new(FakeSubscribedTrack)
