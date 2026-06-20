//go:build linux && !android

package dns

import (
	"path/filepath"
	"testing"
)

func TestSetSystemDNSMissingResolvConf(t *testing.T) {
	oldPath := resolvConfPath
	resolvConfPath = filepath.Join(t.TempDir(), "missing-resolv.conf")
	t.Cleanup(func() { resolvConfPath = oldPath })

	err := SetSystemDNS()
	if err == nil {
		t.Fatal("SetSystemDNS() error = nil, want missing resolv.conf error")
	}
	if !IsResolvConfNotFound(err) {
		t.Fatalf("IsResolvConfNotFound(%v) = false, want true", err)
	}
}

func TestRestoreSystemDNSMissingResolvConf(t *testing.T) {
	oldPath := resolvConfPath
	resolvConfPath = filepath.Join(t.TempDir(), "missing-resolv.conf")
	t.Cleanup(func() { resolvConfPath = oldPath })

	if err := RestoreSystemDNS(); err != nil {
		t.Fatalf("RestoreSystemDNS() error = %v, want nil", err)
	}
}
