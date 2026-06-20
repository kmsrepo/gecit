package dns

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const androidNetIDEnv = "GECIT_ANDROID_NETID"

var androidSavedDNS []string

// SetSystemDNS points Android's active resolver at gecit's local DoH server.
// Android does not use /etc/resolv.conf; on rooted devices the resolver is
// controlled through netd's `ndc resolver` interface. Users can override the
// detected netId with GECIT_ANDROID_NETID when Android's shell output differs
// across releases or vendor builds.
func SetSystemDNS() error {
	netID, err := androidNetID()
	if err != nil {
		return err
	}

	androidSavedDNS = androidCurrentDNS(netID)
	if err := androidSetDNS(netID, []string{"127.0.0.1"}); err != nil {
		return fmt.Errorf("set Android DNS for netId %s: %w", netID, err)
	}
	return nil
}

func RestoreSystemDNS() error {
	netID, err := androidNetID()
	if err != nil {
		return nil
	}
	servers := androidSavedDNS
	if len(servers) == 0 {
		servers = []string{"8.8.8.8", "1.1.1.1"}
	}
	return androidSetDNS(netID, servers)
}

func androidNetID() (string, error) {
	if netID := strings.TrimSpace(os.Getenv(androidNetIDEnv)); netID != "" {
		return netID, nil
	}

	out, err := exec.Command("cmd", "connectivity", "get", "active-network").CombinedOutput()
	if err == nil {
		for _, field := range strings.Fields(string(out)) {
			field = strings.Trim(field, "[](),")
			if isDecimal(field) {
				return field, nil
			}
		}
	}

	// Most Android devices use netId 100 for the default Wi-Fi/mobile network.
	// Keep this as a fallback so rooted shells without `cmd connectivity` can
	// still use `ndc resolver`.
	return "100", nil
}

func androidCurrentDNS(netID string) []string {
	out, err := exec.Command("ndc", "resolver", "getnetdns", netID).CombinedOutput()
	if err != nil {
		return nil
	}
	var servers []string
	for _, field := range strings.Fields(string(out)) {
		if strings.Count(field, ".") == 3 || strings.Contains(field, ":") {
			servers = append(servers, strings.Trim(field, ",;"))
		}
	}
	return servers
}

func androidSetDNS(netID string, servers []string) error {
	args := append([]string{"resolver", "setnetdns", netID, ""}, servers...)
	out, err := exec.Command("ndc", args...).CombinedOutput()
	if err != nil {
		return fmt.Errorf("ndc %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(string(out)))
	}
	return nil
}

func isDecimal(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
