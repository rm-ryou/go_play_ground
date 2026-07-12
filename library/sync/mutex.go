package sync

import (
	"runtime"
	"sync/atomic"
)

const mutexLocked int32 = 1 << iota

type Mutex struct {
	state int32
}

type Locker interface {
	Lock()
	Unlock()
}

func (m *Mutex) Lock() {
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	m.lockSlow()
}

func (m *Mutex) lockSlow() {
	for !atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		runtime.Gosched()
	}
}

func (m *Mutex) TryLock() bool {
	return atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked)
}

func (m *Mutex) Unlock() {
	if atomic.AddInt32(&m.state, -mutexLocked) != 0 {
		panic("sync: unlock of unlocked mutex")
	}
}
