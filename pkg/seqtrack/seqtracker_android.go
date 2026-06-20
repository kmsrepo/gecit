package seqtrack

import (
	"fmt"
	"net"
	"time"

	"github.com/boratanrikulu/gecit/pkg/capture"
)

type SeqTracker struct{}

func NewSeqTracker(iface string, ports []uint16) (*SeqTracker, error) {
	return nil, fmt.Errorf("seq/ack tracker is not available on Android")
}

func (st *SeqTracker) WaitForSeqAck(localPort uint16, timeout time.Duration) *capture.ConnectionEvent {
	return nil
}

func (st *SeqTracker) Stop() {}

func SetSeqTracker(st *SeqTracker) {}

func GetSeqAck(conn net.Conn) (seq, ack uint32) {
	return 1, 1
}
