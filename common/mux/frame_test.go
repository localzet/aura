package mux_test

import (
	"testing"

	"github.com/localzet/aura/common"
	"github.com/localzet/aura/common/buf"
	"github.com/localzet/aura/common/mux"
	"github.com/localzet/aura/common/net"
)

func BenchmarkFrameWrite(b *testing.B) {
	frame := mux.FrameMetadata{
		Target:        net.TCPDestination(net.DomainAddress("www.example.com"), net.Port(80)),
		SessionID:     1,
		SessionStatus: mux.SessionStatusNew,
	}
	writer := buf.New()
	defer writer.Release()

	for i := 0; i < b.N; i++ {
		common.Must(frame.WriteTo(writer))
		writer.Clear()
	}
}
