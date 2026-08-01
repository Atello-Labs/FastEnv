package store

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Store struct {
	BaseDir string
}

func NewStore() (*Store, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("could not locate user home directory: %w", err)
	}

	baseDir := filepath.Join(home, ".fast-env", "store")
	if err := os.MkdirAll(baseDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create global store directory: %w", err)
	}

	return &Store{BaseDir: baseDir}, nil

}

func (s *Store) GetEnvPath(hash string) string {
	return filepath.Join(s.BaseDir, hash)
}

func (s *Store) Exist(hash string) (bool, error) {
	envPath := s.GetEnvPath(hash)
	info, err := os.Stat(envPath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return false, nil
		}
		return false, fmt.Errorf("error inspecting store path %s: %w", envPath, err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("store target %s exist but is not a directory", envPath)
	}
	return true, nil
}
