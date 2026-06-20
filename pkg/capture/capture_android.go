package capture

import "fmt"

func NewCapture(iface string, ports []uint16) (Detector, error) {
	return nil, fmt.Errorf("pcap capture is not used on Android TUN mode")
}
