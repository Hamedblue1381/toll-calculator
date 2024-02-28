package main

import (
	"fmt"

	"github.com/hamedblue1381/tolling/types"
)

type MemoryStore struct {
	data map[int]float64
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]float64),
	}
}

func (m *MemoryStore) Insert(d types.Distance) error {
	m.data[d.OBUID] += d.Value
	return nil
}

func (m *MemoryStore) Get(obudID int) (float64, error) {
	distance, ok := m.data[obudID]
	if !ok {
		return 0.0, fmt.Errorf("invalid obuID %d", obudID)
	}
	return distance, nil
}
