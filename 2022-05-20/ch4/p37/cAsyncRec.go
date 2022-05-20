package p37

import (
	"sync"
)

// var PostReceiveError = postReceiveError

// ----- legacy code -----
type CAsyncSsRec struct {
	mLock          sync.Mutex
	SslInitialized bool
	Count          int
}

func (c *CAsyncSsRec) Init() bool {
	if c.SslInitialized {
		return true
	}
	c.mLock.Lock()
	defer c.mLock.Unlock()
	c.Count++

	PostReceiveError(SocketCallback, SSLFailure)
	c.SslInitialized = true

	// A lot of code…

	return true
}

const (
	SocketCallback = 1
	SSLFailure     = 2
)

// ----- legacy code -----

func PostReceiveError(eType, errorCode uint) {
	panic("postReceiveError")
}
