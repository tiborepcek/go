package snippets

import (
	"fmt"
	"net"
)

// GetIPv4 retrieves the first non-loopback IPv4 address of the host by iterating
// over network interfaces.
func GetIPv4() (net.IP, error) {
	ips, err := GetIPv4s()
	if err != nil {
		return nil, err
	}
	// GetIPv4s ensures that on success, the returned slice is not empty,
	// so we can safely access the first element without a length check.
	return ips[0], nil
}

// GetIPv4s retrieves all non-loopback IPv4 addresses from all network interfaces.
func GetIPv4s() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, fmt.Errorf("failed to get interface addresses: %w", err)
	}

	var ips []net.IP
	for _, addr := range addrs {
		// Check if the address is an IPNet and not a loopback.
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			// Check if it's an IPv4 address.
			if ipv4 := ipnet.IP.To4(); ipv4 != nil {
				ips = append(ips, ipv4)
			}
		}
	}
	if len(ips) == 0 {
		return nil, fmt.Errorf("no non-loopback IPv4 addresses found")
	}

	return ips, nil
}
