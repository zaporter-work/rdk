package wx250s

import (
	"sync"

	"go.viam.com/robotcore/api"
)

type Provider struct {
	moveLock *sync.Mutex
}

func (p *Provider) Ready(r api.Robot) error {
	return nil
}

func getProviderOrCreate(r api.MutableRobot) *Provider {
	p := r.ProviderByName("wx250s")
	if p == nil {
		p = &Provider{&sync.Mutex{}}
		r.AddProvider(p, api.ComponentConfig{})
	}
	return p.(*Provider)
}
