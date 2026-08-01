package hasher

import (
	"fmt"
	"github.com/cespare/xxhash/v2"
	"io"
	"os"
)

func HashFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file:%w", err)
	}
	defer file.Close()

	hasher := xxhash.New()
	if _, err := io.Copy(hasher, file); err != nil {
		return "", fmt.Errorf("failed to hash file contents: %w", err)
	}

	return fmt.Sprintf("%016x", hasher.Sum64()), nil

}
