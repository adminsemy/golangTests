package testAsyncMap

import (
	"errors"
	"sync"
)

type CasheMap interface {
	Read(string) (string, error)
	Write(string, string) error
	Delete(string) error
}

type asyncMap struct {
	mux      *sync.Mutex
	asyncMap map[string]string
}

func NewAsyncMap() CasheMap {
	return &asyncMap{
		&sync.Mutex{},
		map[string]string{},
	}
}

func (a *asyncMap) Read(key string) (string, error) {
	a.mux.Lock()
	defer a.mux.Unlock()
	res, ok := a.asyncMap[key]
	if ok {
		return res, nil
	}
	return "", errors.New("data not found")
}

func (a *asyncMap) Write(key, data string) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	a.asyncMap[key] = data
	return nil
}

func (a *asyncMap) Delete(key string) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	delete(a.asyncMap, key)
	return nil
}
