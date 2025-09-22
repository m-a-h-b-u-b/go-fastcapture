//go:build pfring
// +build pfring

package pfring

/*
#cgo LDFLAGS: -lpfring
#include <pfring_zc.h>
*/
import "C"
import (
	"context"
	"fmt"

	"github.com/m-a-h-b-u-b/high-speed-packet-capture/internal/capture"
)

type pfringBackend struct {
	iface string
}

func newPFRing(ctx context.Context, cfg capture.Config) (capture.Backend, error) {
	fmt.Println("PF_RING backend stub — implement initialization here")
	return &pfringBackend{iface: cfg.Iface}, nil
}

func (p *pfringBackend) Run(handler func(capture.Packet) error) error {
	// TODO: PF_RING capture loop
	return fmt.Errorf("PF_RING capture not implemented")
}

func (p *pfringBackend) Close() error {
	return nil
}
