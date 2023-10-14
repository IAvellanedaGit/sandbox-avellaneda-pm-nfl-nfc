package nfc

import (
	"math/rand"
	"strings"

	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
)

type Service struct {
	b      *bridge.Bridge
	logger bridge.Logger
}

func NewService(b *bridge.Bridge, logger bridge.Logger) *Service {
	svc := &Service{
		b:      b,
		logger: logger,
	}

	b.RegisterService("nfl-nfc", svc)

	return svc
}

func (s *Service) Teams() []string {
	return []string{
		"seahawks",
		"rams",
		"cardinals",
		"49ers",
		"lions",
		"packers",
		"bears",
		"vikings",
		"cowboys",
		"giants",
		"eagles",
		"commanders",
		"saints",
		"buccaneers",
		"falcons",
		"panthers",
	}
}

func (s *Service) ValidateTeam(team string) bool {

	t := strings.ToLower(team)
	for _, n := range s.Teams() {
		if t == n {
			return true
		}
	}

	return false
}

func (s *Service) PredictScore(team string) int {

	score := rand.Intn(45)
	return score
}
