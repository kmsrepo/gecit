package app

import (
	"fmt"
	"os"

	bpf "github.com/boratanrikulu/gecit/pkg/ebpf"
)

func printPlatformStatus() {
	fmt.Printf("  engine:     ebpf-sockops (Android)\n")
	fmt.Printf("  dns:        Android resolver via ndc (no /etc/resolv.conf)\n")

	if os.Geteuid() != 0 {
		fmt.Printf("  (run with su/root for accurate capability detection)\n")
		return
	}

	fmt.Printf("  sock_ops:   %s\n", boolStatus(bpf.HaveSockOps()))
	fmt.Printf("  setsockopt: %s\n", boolStatus(bpf.HaveSockOpsSetsockopt()))
}

func boolStatus(ok bool) string {
	if ok {
		return "supported"
	}
	return "NOT supported"
}
