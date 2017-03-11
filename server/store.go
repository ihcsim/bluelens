package main

import (
	"github.com/ihcsim/bluelens"
)

var (
	// singleton instance of the store
	instance core.Store

	// to use a different store, assign this varible to the store's constructor function.
	// store is a function of type func() core.Store
	store = inmemoryStore
)

func inmemoryStore() core.Store {
	if instance == nil {
		instance = core.NewInMemoryStore()
	}

	return instance
}
