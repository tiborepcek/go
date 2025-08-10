# Learning Go

This project is a collection of practical Go code snippets demonstrating common tasks such as file manipulation, network information retrieval, and system inspection.

## Prerequisites

- Go (version 1.23 or newer)

## Getting Started

Follow these steps to get the project running on your local machine.

1.  **Clone the repository:**
    ```sh
    git clone https://github.com/tiborepcek/go.git
    ```

2.  **Navigate to the project directory:**
    ```sh
    cd go
    ```

3.  **Install dependencies:**
    The Go toolchain will automatically handle dependencies when you build or run the project. The required modules are defined in `go.mod` and their integrity is verified using `go.sum`.

## Usage

To run the main application and see all the snippets in action, execute the following command from the root of the project:

```sh
go run .
```

The program will execute and clean up after itself, removing any files or directories it created during its run.

## Demonstrated Snippets

The application demonstrates the following functionalities:

1.  **Get Hostname:** Retrieves and prints the system's hostname.
2.  **Get IP Addresses:** Lists all active non-loopback IPv4 addresses.
3.  **Get CPU Info:** Shows the number of physical CPU cores using the `gopsutil` library.
4.  **Zip a File:** Creates a text file and compresses it into a `.zip` archive.
5.  **Unzip a File:** Decompresses the previously created archive into a new directory.
6.  **Zip a Directory:** Creates a directory with sub-folders and files, then compresses the entire directory structure into a `.zip` archive.