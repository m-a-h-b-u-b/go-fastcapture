//go:build !dpdk && !pfring
// +build !dpdk,!pfring

package afpacket

import (
	"context"
	"fmt"
	"net"
	"syscall"
	"time"

	"github.com/m-a-h-b-u-b/high-speed-packet-capture/internal/capture"
)

type afpacketBackend struct {
	iface string
}

func newAFPacket(ctx context.Context, cfg capture.Config) (capture.Backend, error) {
	return &afpacketBackend{iface: cfg.Iface}, nil
}

func (a *afpacketBackend) Run(handler func(capture.Packet) error) error {
	// Simplest possible example: open raw socket and read packets.
	conn, err := net.ListenPacket("ip4:icmp", a.iface)
	if err != nil {
		return err
	}
	defer conn.Close()

	buf := make([]byte, 65535)
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			n, _, err := conn.ReadFrom(buf)
			if err != nil {
				return err
			}
			pkt := capture.Packet{
				Data: append([]byte(nil), buf[:n]...),
				Ts:   time.Now().UnixNano(),
			}
			if err := handler(pkt); err != nil {
				return err
			}
		}
	}
}

func (a *afpacketBackend) Close() error { return nil }
