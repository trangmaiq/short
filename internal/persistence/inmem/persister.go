package inmem

import (
	"github.com/trangmaiq/short/internal/handler/url"
	"sync"
)

type Persister struct {
	urls []url.URL

	l *sync.Mutex
}

func NewPersister() *Persister {
	return &Persister{
		urls: make([]url.URL, 0),
	}
}

func (p *Persister) CreateURL(u *url.URL) error {
	p.l.Lock()
	defer p.l.Unlock()

	p.urls = append(p.urls, *u)
	return nil
}
