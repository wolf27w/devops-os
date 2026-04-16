package common

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileStorage struct {
	basePath string
}

func NewFileStorage(basePath string) *FileStorage {
	return &FileStorage{basePath: basePath}
}

func (fs *FileStorage) ReadJSON(path string, v interface{}) error {
	fullPath := filepath.Join(fs.basePath, path)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", fullPath, err)
	}

	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from %s: %w", fullPath, err)
	}

	return nil
}

func (fs *FileStorage) WriteJSON(path string, v interface{}) error {
	fullPath := filepath.Join(fs.basePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", fullPath, err)
	}

	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON for %s: %w", fullPath, err)
	}

	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", fullPath, err)
	}

	return nil
}

func (fs *FileStorage) ListFiles(path string) ([]string, error) {
	fullPath := filepath.Join(fs.basePath, path)

	files, err := os.ReadDir(fullPath)
	if err != nil {
		return nil, fmt.Errorf("failed to list files in %s: %w", fullPath, err)
	}

	var result []string
	for _, file := range files {
		result = append(result, file.Name())
	}

	return result, nil
}

func (fs *FileStorage) AppendLog(path, message string) error {
	fullPath := filepath.Join(fs.basePath, path)

	if err := os.MkdirAll(filepath.Dir(fullPath), 0755); err != nil {
		return fmt.Errorf("failed to create directory for %s: %w", fullPath, err)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)

	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file %s: %w", fullPath, err)
	}
	defer file.Close()

	if _, err := file.WriteString(logEntry); err != nil {
		return fmt.Errorf("failed to write to log file %s: %w", fullPath, err)
	}

	return nil
}

func (fs *FileStorage) PathExists(path string) bool {
	fullPath := filepath.Join(fs.basePath, path)
	_, err := os.Stat(fullPath)
	return err == nil
}

func (fs *FileStorage) EnsureDir(path string) error {
	fullPath := filepath.Join(fs.basePath, path)
	return os.MkdirAll(fullPath, 0755)
}
