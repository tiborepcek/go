package snippets

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Unzip will decompress a zip archive to a specified destination directory.
func Unzip(source, destination string) error {
	// 1. Open the zip archive for reading.
	reader, err := zip.OpenReader(source)
	if err != nil {
		return fmt.Errorf("failed to open zip archive %s: %w", source, err)
	}
	defer reader.Close()

	// 2. Iterate through each file/directory in the archive.
	for _, file := range reader.File {
		filePath := filepath.Join(destination, file.Name)

		// 3. Check for ZipSlip vulnerability. This prevents files from being written
		// outside the destination directory.
		if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", filePath)
		}

		// 4. Handle directories.
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", filePath, err)
			}
			continue
		}

		// 5. Handle files.
		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			return fmt.Errorf("failed to create parent directory for %s: %w", filePath, err)
		}

		destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return fmt.Errorf("failed to create destination file %s: %w", filePath, err)
		}
		defer destFile.Close()

		zippedFile, err := file.Open()
		if err != nil {
			return fmt.Errorf("failed to open file within zip %s: %w", file.Name, err)
		}
		defer zippedFile.Close()

		if _, err = io.Copy(destFile, zippedFile); err != nil {
			return fmt.Errorf("failed to copy content to %s: %w", filePath, err)
		}
	}
	return nil
}
