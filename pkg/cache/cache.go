package cache

import (
	"log/slog"
	"sync"

	"github.com/amirmohammadkariimi/interview-task/internal/pkg/models"
)

// Cache struct with map and sync.RWMutex for concurrency
type Cache struct {
	mu    sync.RWMutex
	cache map[string]models.DNSQuery
}

// NewCache initializes a new cache
func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]models.DNSQuery),
	}
}

// Set adds a DNSQuery to the cache
func (c *Cache) Set(domain string, query models.DNSQuery) {
	c.mu.Lock()         // Lock for writing
	defer c.mu.Unlock() // Unlock after writing
	c.cache[domain] = query
	slog.Info("Added to cache", "domain", domain)
}

// Get retrieves a DNSQuery from the cache by Domain
func (c *Cache) Get(domain string) (models.DNSQuery, bool) {
	c.mu.RLock()         // Lock for reading
	defer c.mu.RUnlock() // Unlock after reading
	query, found := c.cache[domain]
	return query, found
}

// Delete removes a DNSQuery from the cache by Domain
func (c *Cache) Delete(domain string) {
	c.mu.Lock()         // Lock for writing
	defer c.mu.Unlock() // Unlock after writing
	delete(c.cache, domain)
	slog.Info("Deleted from cache", "domain", domain)
}
