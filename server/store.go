package main

import (
	"sync"

	"github.com/ihcsim/bluelens"
)

var (
	// singleton instance of the store
	instance core.Store

	// to use a different store, assign this varible to the store's constructor function.
	// store is a function of type func() core.Store
	store = inmemoryStore

	once sync.Once
)

func inmemoryStore() core.Store {
	once.Do(func() {
		instance = core.NewInMemoryStore()
	})

	return instance
}
