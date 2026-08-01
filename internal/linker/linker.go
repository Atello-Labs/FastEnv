package linker

import (
	"fmt"
	"os"
)

type Linker struct {
	LinkName string // Default is usually ".venv"
}

func NewLinker(linkName string) *Linker {
	if linkName == "" {
		linkName = ".venv"
	}
	return &Linker{LinkName: linkName}
}

func (l *Linker) LinkSymlink(targetPath string) error {

	_, err := os.Lstat(l.LinkName)
	if err == nil {
		if err := os.Remove(l.LinkName); err != nil {
			return fmt.Errorf("failed to remove existing linkg %s: %w", l.LinkName, err)
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("failed to inspect local linkg path %s: %w", l.LinkName, err)
		}

		if err := os.Symlink(targetPath, l.LinkName); err != nil {
			return fmt.Errorf("failed to create symlink from %s to %s: %w", l.LinkName, targetPath, err)
		}
	}
	return nil
}
