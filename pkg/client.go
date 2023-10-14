package nfc

import (
	"fmt"

	bridge "github.com/IAvellanedaGit/nuevo-platform-mobile-bridge"
	infc "github.com/IAvellanedaGit/sandbox-avellaneda-pm-nfl-nfc/service"
)

type Client struct {
	svc *infc.Service
}

func NewClient(b *bridge.Bridge) (*Client, error) {
	svc, ok := b.ServiceInstance("nfl-nfc").(*infc.Service)
	if !ok || svc == nil {
		return nil, fmt.Errorf("could not locate nfc service")
	}

	return &Client{
		svc: svc,
	}, nil
}

func (c *Client) Teams() []string {
	return c.svc.Teams()
}

func (c *Client) ValidateTeam(team string) bool {
	return c.svc.ValidateTeam(team)
}

func (c *Client) PredictScore(team string) int {
	return c.svc.PredictScore(team)
}
