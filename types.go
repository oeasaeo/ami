package ami

import (
	"sync"
	"time"

	"git.aqq.me/go/retrier"
	"github.com/go-redis/redis"
)

// Qu client for ami
type Qu struct {
	rDB       *redis.ClusterClient
	wgCons    *sync.WaitGroup
	wgProd    *sync.WaitGroup
	opt       Options
	cCons     chan Message
	needClose bool
	retr      *retrier.Retrier
	cProd     chan string
}

// Options - options for Qu client for ami
type Options struct {
	Name              string
	Consumer          string
	ShardsCount       int8
	PrefetchCount     int64
	Block             time.Duration
	PendingBufferSize int64
}

// Message from Qu
type Message struct {
	Body   string
	ID     string
	Stream string
	Group  string
}