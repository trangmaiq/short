package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/trangmaiq/short/internal/handler/url"
	"time"
)

type Persister struct {
	s *gocql.Session
}

func NewPersister() (*Persister, error) {
	cluster := gocql.NewCluster("localhost")
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.ConnectTimeout = 6 * time.Second
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "secret",
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &Persister{s: session}, nil
}

func (p *Persister) CreateURL(u *url.URL) error {
	err := p.s.Query(
		fmt.Sprintf(
			"INSERT INTO short.urls (hash, created_at, expired_at, original_url, user_id) VALUES ('%s', '%s', '%s', '%s', '%s')",
			u.Hash,
			u.CreatedAt.Format("2006-01-02"),
			u.ExpiredAt.Format("2006-01-02"),
			u.OriginalURL,
			u.UserID,
		)).Exec()
	if err != nil {
		return fmt.Errorf("create url failed: %w", err)
	}

	return nil
}
