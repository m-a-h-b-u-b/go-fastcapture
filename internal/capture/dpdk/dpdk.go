//go:build dpdk
// +build dpdk

package dpdk

/*
#cgo CFLAGS: -I/usr/local/include/dpdk
#cgo LDFLAGS: -ldpdk
#include <rte_eal.h>
#include <rte_ethdev.h>
*/
import "C"
import (
	"context"
	"fmt"

	"github.com/m-a-h-b-u-b/high-speed-packet-capture/internal/capture"
)

type dpdkBackend struct {
	iface string
}

func newDPDK(ctx context.Context, cfg capture.Config) (capture.Backend, error) {
	fmt.Println("DPDK backend stub — implement initialization here")
	return &dpdkBackend{iface: cfg.Iface}, nil
}

func (d *dpdkBackend) Run(handler func(capture.Packet) error) error {
	// TODO: initialize DPDK EAL, configure ports/queues,
	// poll RX queues, call handler for each packet.
	return fmt.Errorf("DPDK capture not implemented")
}

func (d *dpdkBackend) Close() error {
	// TODO: free DPDK resources
	return nil
}
