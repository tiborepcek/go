package snippets

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// ZipFile compresses a single source file into a destination zip archive.
// The file within the archive will have the same name as the source file.
func ZipFile(sourcePath, destPath string) error {
	// 1. Create the destination zip file.
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file %s: %w", destPath, err)
	}
	defer destFile.Close()

	// 2. Create a new zip writer.
	zipWriter := zip.NewWriter(destFile)
	defer zipWriter.Close()

	// 3. Open the source file for reading.
	sourceFile, err := os.Open(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to open source file %s: %w", sourcePath, err)
	}
	defer sourceFile.Close()

	// 4. Create a new file entry in the zip archive. We use filepath.Base
	// to ensure the file in the zip doesn't have the full path.
	writer, err := zipWriter.Create(filepath.Base(sourceFile.Name()))
	if err != nil {
		return fmt.Errorf("failed to create entry in zip file: %w", err)
	}

	// 5. Copy the source file's content into the zip entry.
	_, err = io.Copy(writer, sourceFile)
	return err
}

// ZipDirectory compresses a source directory into a destination zip archive.
func ZipDirectory(sourceDir, destPath string) error {
	// 1. Create the destination zip file.
	destFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create zip file %s: %w", destPath, err)
	}
	defer destFile.Close()

	// 2. Create a new zip writer.
	zipWriter := zip.NewWriter(destFile)
	defer zipWriter.Close()

	// 3. Walk the directory tree.
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// We only care about files. Directories are created implicitly by the file paths.
		if info.IsDir() {
			return nil
		}

		// 4. Get the relative path of the file. This will be the path inside the zip.
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		// 5. Create a new file entry in the zip archive.
		writer, err := zipWriter.Create(relPath)
		if err != nil {
			return fmt.Errorf("failed to create entry for %s in zip file: %w", relPath, err)
		}

		// 6. Open the source file.
		fileToZip, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open source file %s: %w", path, err)
		}
		defer fileToZip.Close()

		// 7. Copy the source file's content into the zip entry.
		_, err = io.Copy(writer, fileToZip)
		return err
	})
}
