package engine

import (
	"sync"

	"github.com/google/uuid"
)

type Clients struct {
	mu   sync.Mutex
	repo map[uuid.UUID]*Client
}

func (c *Clients) Add(client *Client) {
	c.mu.Lock()
	key := client.Id
	c.repo[key] = client
	c.mu.Unlock()
}

func (c *Clients) Load(id uuid.UUID) *Client {
	c.mu.Lock()
	defer c.mu.Unlock()
	client, ok := c.repo[id]
	if !ok {
		return nil
	}
	return client
}

func (c *Clients) Remove(id uuid.UUID) {
	c.mu.Lock()
	delete(c.repo, id)
	c.mu.Unlock()
}
