package cacheddns

import (
	"net"
	"sync"
	"time"

	"github.com/juju/errors"
)

// CachedDomain stores info about the cached domain in memory.
type CachedDomain struct {
	lock     *sync.RWMutex
	domain   string
	address  string
	ttl      time.Duration
	deadline time.Time
}

// New creates a new cached domain in memory. This will NOT call the DNS server yet.
func New(domain string, ttl time.Duration) *CachedDomain {
	return &CachedDomain{
		lock:   new(sync.RWMutex),
		domain: domain,
		ttl:    ttl,
	}
}

// Get returns the currently cached address and a true flag. If it has not been
// resolved yet or the TTL has passed it will return an empty string with the
// flag set to false. It's recommended to use Resolve() directly instead.
func (c *CachedDomain) Get() (string, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	if c.deadline.After(time.Now()) {
		return c.address, true
	}

	return "", false
}

// Update the cached address if needed and return it. Any connection error will
// be returned. It's recommended to use Resolve() directly instead.
func (c *CachedDomain) Update() (string, error) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.deadline.After(time.Now()) {
		return c.address, nil
	}

	ips, err := net.LookupIP(c.domain)
	if err != nil {
		return "", errors.Trace(err)
	}

	if len(ips) == 0 {
		return "", errors.Errorf("cannot resolve domain: %s", c.domain)
	}

	c.address = ips[0].String()
	c.deadline = time.Now().Add(c.ttl)

	return c.address, nil
}

// Resolve returns the cached address or checks with the DNS server the setting
// if the TTL has expired. It returns the first address found.
func (c *CachedDomain) Resolve() (string, error) {
	address, ok := c.Get()
	if !ok {
		var err error
		address, err = c.Update()
		if err != nil {
			return "", errors.Trace(err)
		}
	}

	return address, nil
}
