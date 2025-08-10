package snippets

import "os"

// GetHostname retrieves the host name of the current system.
func GetHostname() (string, error) {
	return os.Hostname()
}
