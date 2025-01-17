package utils

import (
	"errors"
	"sync"
)

type Lock[K comparable, V any] struct {
	running map[K]*Task[V]
	lock    sync.Mutex
}

type Task[V any] struct {
	ready     sync.WaitGroup
	listeners []chan Result[V]
}

type Result[V any] struct {
	ok  V
	err error
}

func NewLock[K comparable, V any]() Lock[K, V] {
	return Lock[K, V]{
		running: make(map[K]*Task[V]),
	}
}

func (l *Lock[K, V]) Start(key K) (func() (V, error), func(val V, err error) (V, error)) {
	l.lock.Lock()
	defer l.lock.Unlock()

	task, exists := l.running[key]
	if exists {
		ret := make(chan Result[V])
		task.listeners = append(task.listeners, ret)
		return func() (V, error) {
			res := <-ret
			return res.ok, res.err
		}, nil
	}

	l.running[key] = &Task[V]{
		listeners: make([]chan Result[V], 0),
	}

	return nil, func(val V, err error) (V, error) {
		l.lock.Lock()
		defer l.lock.Unlock()

		task, ok := l.running[key]
		if !ok {
			return val, errors.New("invalid run lock state. aborting")
		}

		for _, listener := range task.listeners {
			listener <- Result[V]{ok: val, err: err}
			close(listener)
		}
		delete(l.running, key)
		return val, err
	}
}

func (l *Lock[K, V]) WaitFor(key K) (V, error) {
	l.lock.Lock()
	task, exists := l.running[key]

	if !exists {
		l.lock.Unlock()
		var val V
		return val, nil
	}

	ret := make(chan Result[V])
	task.listeners = append(task.listeners, ret)

	l.lock.Unlock()
	res := <-ret
	return res.ok, res.err
}
