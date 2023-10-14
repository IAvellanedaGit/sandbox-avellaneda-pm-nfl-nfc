package nfc

import (
	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
	"github.com/IAvellanedaGit/sandbox-avellaneda-pm-nfl-nfc/service/model"
	flatbuffers "github.com/google/flatbuffers/go"
)

type Manager struct {
	b      *bridge.Bridge
	bus    *bridge.MessageBus
	logger bridge.Logger
	svc    *Service
}

const NFC_VERSION = "0.0.1"

func NewManager(b *bridge.Bridge) *Manager {
	m := &Manager{
		b:      b,
		bus:    b.MessageBusInstance(),
		logger: b,
	}

	m.bus.Subscribe("nfl:nfc:initialize", m.initialize)

	return m
}

func (mgr *Manager) initialize(m *bridge.Message) ([]byte, error) {

	request := model.GetRootAsInitializeRequest(m.Data, 0)

	mgr.logger.Debug("Initializing NFL nfc Manager [%s] LogLevel [%d]",
		NFC_VERSION,
		request.LogLevel())

	mgr.logger = bridge.NewLogger("nfc", bridge.LogLevel(request.LogLevel()), mgr.logger)

	mgr.svc = NewService(mgr.b, mgr.logger)

	// subscribe to other messages
	mgr.bus.Subscribe("nfl:nfc:list-teams", mgr.listTeams)
	mgr.bus.Subscribe("nfl:nfc:validate-team", mgr.validateTeam)
	mgr.bus.Subscribe("nfl:nfc:predict-score", mgr.predictScore)

	builder := flatbuffers.NewBuilder(0)

	vOffset := builder.CreateString(NFC_VERSION)

	// create response message
	model.StringResponseStart(builder)
	model.StringResponseAddOutput(builder, vOffset)
	builder.Finish(model.StringResponseEnd(builder))

	return builder.FinishedBytes(), nil
}

func (mgr *Manager) listTeams(m *bridge.Message) ([]byte, error) {

	teams := mgr.svc.Teams()

	builder := flatbuffers.NewBuilder(0)

	teamOffsets := []flatbuffers.UOffsetT{}
	for _, t := range teams {
		teamOffsets = append(teamOffsets, builder.CreateString(t))
	}

	model.ListTeamsResponseStartTeamsVector(builder, len(teamOffsets))
	for _, tOffset := range teamOffsets {
		builder.PrependUOffsetT(tOffset)
	}
	teamsOffset := builder.EndVector(len(teamOffsets))

	model.ListTeamsResponseStart(builder)
	model.ListTeamsResponseAddTeams(builder, teamsOffset)
	builder.Finish(model.ListTeamsResponseEnd(builder))

	return builder.FinishedBytes(), nil
}

func (mgr *Manager) validateTeam(m *bridge.Message) ([]byte, error) {

	request := model.GetRootAsTeamRequest(m.Data, 0)
	ok := mgr.svc.ValidateTeam(string(request.Team()))

	builder := flatbuffers.NewBuilder(0)

	model.BoolResponseStart(builder)
	model.BoolResponseAddOutput(builder, ok)
	builder.Finish(model.BoolResponseEnd(builder))

	return builder.FinishedBytes(), nil
}

func (mgr *Manager) predictScore(m *bridge.Message) ([]byte, error) {
	request := model.GetRootAsTeamRequest(m.Data, 0)
	score := mgr.svc.PredictScore(string(request.Team()))

	builder := flatbuffers.NewBuilder(0)

	model.IntResponseStart(builder)
	model.IntResponseAddOutput(builder, int32(score))
	builder.Finish(model.IntResponseEnd(builder))

	return builder.FinishedBytes(), nil
}
