package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tiborepcek/go/snippets"

	"github.com/shirou/gopsutil/v3/cpu"
)

func main() {
	showHostname()
	showIPAddresses()
	showCPUInfo()

	dummyFileName, zipFileName := demonstrateFileZipping()
	defer os.Remove(dummyFileName)
	defer os.Remove(zipFileName)

	unzipDir := demonstrateFileUnzipping(zipFileName)
	defer os.RemoveAll(unzipDir)

	sourceDir, dirZipFileName := demonstrateDirectoryZipping()
	defer os.RemoveAll(sourceDir)
	defer os.Remove(dirZipFileName)
}

// showHostname demonstrates getting and printing the system's hostname.
func showHostname() {
	fmt.Println("--- 1. Get Hostname ---")
	hostname, err := snippets.GetHostname()
	if err != nil {
		log.Fatalf("Error getting hostname: %v", err)
	}
	fmt.Printf("Hostname: %s\n\n", hostname)
}

// showIPAddresses demonstrates getting and printing all active IPv4 addresses.
func showIPAddresses() {
	fmt.Println("--- 2. Get all active IP Addresses ---")
	ipv4s, err := snippets.GetIPv4s()
	if err != nil {
		log.Fatalf("Error getting IPv4 addresses: %v", err)
	}

	fmt.Println("Active IPv4 Addresses:")
	for _, ip := range ipv4s {
		fmt.Printf("- %s\n", ip)
	}
	fmt.Println()
}

// showCPUInfo demonstrates getting and printing CPU core count.
func showCPUInfo() {
	fmt.Println("--- 3. Get CPU Info ---")
	// Get core count (physical cores)
	coreCount, err := cpu.Counts(false)
	if err != nil {
		log.Fatalf("Error getting CPU core count: %v", err)
	}
	fmt.Printf("Physical CPU Cores: %d\n\n", coreCount)
}

// demonstrateFileZipping creates a dummy file and zips it.
// It returns the names of the created files for cleanup in main.
func demonstrateFileZipping() (dummyFileName, zipFileName string) {
	fmt.Println("--- 4. Zip a file ---")
	dummyFileName = "file_to_zip.txt"
	zipFileName = "archive.zip"
	content := []byte("Hello, this is the content of the file to be zipped.")
	if err := os.WriteFile(dummyFileName, content, 0666); err != nil {
		log.Fatalf("Failed to create dummy file: %v", err)
	}

	fmt.Printf("Zipping '%s' into '%s'...\n", dummyFileName, zipFileName)
	if err := snippets.ZipFile(dummyFileName, zipFileName); err != nil {
		log.Fatalf("Failed to zip file: %v", err)
	}
	fmt.Println("File zipped successfully!")
	fmt.Println()
	return
}

// demonstrateFileUnzipping unzips an archive into a directory.
// It returns the name of the created directory for cleanup in main.
func demonstrateFileUnzipping(zipFileName string) (unzipDir string) {
	fmt.Println("--- 5. Unzip a file ---")
	unzipDir = "unzipped_contents"
	// Clean up the unzip directory in case it exists from a previous run
	if err := os.RemoveAll(unzipDir); err != nil {
		log.Fatalf("Failed to clean up old unzip directory: %v", err)
	}

	fmt.Printf("Unzipping '%s' into directory '%s'...\n", zipFileName, unzipDir)
	if err := snippets.Unzip(zipFileName, unzipDir); err != nil {
		log.Fatalf("Failed to unzip file: %v", err)
	}
	fmt.Println("File unzipped successfully!")
	fmt.Println()
	return
}

// demonstrateDirectoryZipping creates a dummy directory with files and zips it.
// It returns the names of the created artifacts for cleanup in main.
func demonstrateDirectoryZipping() (sourceDir, dirZipFileName string) {
	fmt.Println("--- 6. Zip a directory ---")
	sourceDir = "dir_to_zip"
	dirZipFileName = "dir_archive.zip"

	// Create a dummy directory and some files in it
	fmt.Printf("Creating dummy directory '%s' for zipping...\n", sourceDir)
	if err := os.MkdirAll(sourceDir, 0755); err != nil {
		log.Fatalf("Failed to create source directory: %v", err)
	}

	// Create a file in the root of the directory
	file1Content := []byte("This is file one.")
	if err := os.WriteFile(filepath.Join(sourceDir, "file1.txt"), file1Content, 0666); err != nil {
		log.Fatalf("Failed to create file1.txt: %v", err)
	}

	// Create a subdirectory and a file inside it
	subDir := filepath.Join(sourceDir, "subdir")
	if err := os.MkdirAll(subDir, 0755); err != nil {
		log.Fatalf("Failed to create subdirectory: %v", err)
	}
	file2Content := []byte("This is file two, in a subdirectory.")
	if err := os.WriteFile(filepath.Join(subDir, "file2.txt"), file2Content, 0666); err != nil {
		log.Fatalf("Failed to create file2.txt: %v", err)
	}

	fmt.Printf("Zipping directory '%s' into '%s'...\n", sourceDir, dirZipFileName)
	if err := snippets.ZipDirectory(sourceDir, dirZipFileName); err != nil {
		log.Fatalf("Failed to zip directory: %v", err)
	}
	fmt.Println("Directory zipped successfully!")
	fmt.Println()
	return
}
