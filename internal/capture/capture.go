package capture

import (
	"context"
	"errors"
)

// Config holds runtime configuration for a capture backend.
type Config struct {
	Backend string   // dpdk | pfring | afpacket
	Iface   string   // NIC name or PCI address
	Cores   []int    // optional list of CPU cores to pin
}

// Packet represents a captured network packet.
// Release is optional and used for zero-copy buffers.
type Packet struct {
	Data    []byte
	Ts      int64
	Release func() // call when done if zero-copy
}

// Backend defines the methods all backends must implement.
type Backend interface {
	Run(handler func(Packet) error) error
	Close() error
}

// Factory chooses the correct backend based on cfg.Backend.
func New(ctx context.Context, cfg Config) (Backend, error) {
	switch cfg.Backend {
	case "dpdk":
		return newDPDK(ctx, cfg)
	case "pfring":
		return newPFRing(ctx, cfg)
	case "afpacket":
		return newAFPacket(ctx, cfg)
	default:
		return nil, errors.New("unsupported backend: " + cfg.Backend)
	}
}
