package user

import (
	"log"
	"sync"
)

type ServiceProxy struct {
	realService Service
	cache       map[int]string
	mu          sync.RWMutex
}

func NewServiceProxy(realService Service) *ServiceProxy {
	return &ServiceProxy{
		realService: realService,
		cache:       make(map[int]string),
	}
}

func (p *ServiceProxy) GetUser(id int) (string, error) {
	p.mu.RLock()
	user, found := p.cache[id]
	p.mu.RUnlock()

	if found {
		log.Printf("Cache hit for the user %d\n", id)
		return user, nil
	}

	log.Printf("Cache miss, calling real service for user %d\n", id)
	user, err := p.realService.GetUser(id)
	if err != nil {
		return "", err
	}

	p.mu.Lock()
	p.cache[id] = user
	p.mu.Unlock()
	return user, nil
}
