package builder

import (
	"fmt"
	"os"
	"os/exec"
)

type Builder struct {
	pythonPath string
}

func NewBuilder() (*Builder, error) {
	path, err := exec.LookPath("python3")
	if err != nil {
		path, err = exec.LookPath("python")
		if err != nil {
			return nil, fmt.Errorf("neither python3 nor python was found on system PATH : %w", err)
		}

	}
	return &Builder{pythonPath: path}, nil
}

func (b *Builder) BuildVenv(targetPath string) error {
	cmd := exec.Command(b.pythonPath, "-m", "venv", targetPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("python venv creation failed at %s: %w", targetPath, err)
	}
	return nil
}
