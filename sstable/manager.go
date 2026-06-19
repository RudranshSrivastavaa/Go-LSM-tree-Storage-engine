package sstable

import "path/filepath"

type Manager struct {
	nextID uint64
	files  []string
}

func LoadManager() (*Manager, error) {

	files, err := filepath.Glob(
		"data/sst_*.db",
	)

	if err != nil {
		return nil, err
	}

	manager := &Manager{}

	for _, f := range files {
		manager.Add(f)
	}

	return manager, nil
}

func (m *Manager) Add(
	path string,
) {
	m.files = append(
		m.files,
		path,
	)
}

func (m *Manager) Files() []string {
	return m.files
}