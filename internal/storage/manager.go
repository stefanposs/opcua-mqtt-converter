package storage

import (
	"os"
	"path/filepath"
)

type Manager struct {
	path string
}

func NewManager(path string) *Manager {
	return &Manager{path: path}
}

func (m *Manager) Save(data map[string]interface{}) error {
	file, err := os.Create(filepath.Join(m.path, "data.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	// Implement saving logic
	return nil
}

func (m *Manager) Load() (map[string]interface{}, error) {
	file, err := os.Open(filepath.Join(m.path, "data.json"))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Implement loading logic
	return nil, nil
}
