package sync

import isync "sync"

type Mutex struct {
	once isync.Once
	ch   chan struct{}
}

type Locker interface {
	Lock()
	Unlock()
}

func (m *Mutex) init() {
	m.once.Do(func() {
		m.ch = make(chan struct{}, 1)
		m.ch <- struct{}{}
	})
}

func (m *Mutex) Lock() {
	m.init()
	<-m.ch
}

func (m *Mutex) TryLock() bool {
	m.init()
	select {
	case <-m.ch:
		return true
	default:
		return false
	}
}

func (m *Mutex) Unlock() {
	m.init()
	select {
	case m.ch <- struct{}{}:
	default:
		panic("sync: unlock of unlocked mutex")
	}
}
