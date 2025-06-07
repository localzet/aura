package inbound

import (
	"context"

	"github.com/localzet/aura/common"
	"github.com/localzet/aura/common/net"
	"github.com/localzet/aura/common/serial"
	"github.com/localzet/aura/features"
)

// Handler is the interface for handlers that process inbound connections.
//
// aura:api:stable
type Handler interface {
	common.Runnable
	// The tag of this handler.
	Tag() string
	// Returns the active receiver settings.
	ReceiverSettings() *serial.TypedMessage
	// Returns the active proxy settings.
	ProxySettings() *serial.TypedMessage

	// Deprecated: Do not use in new code.
	GetRandomInboundProxy() (interface{}, net.Port, int)
}

// Manager is a feature that manages InboundHandlers.
//
// aura:api:stable
type Manager interface {
	features.Feature
	// GetHandler returns an InboundHandler for the given tag.
	GetHandler(ctx context.Context, tag string) (Handler, error)
	// AddHandler adds the given handler into this Manager.
	AddHandler(ctx context.Context, handler Handler) error

	// RemoveHandler removes a handler from Manager.
	RemoveHandler(ctx context.Context, tag string) error

	// ListHandlers returns a list of inbound.Handler.
	ListHandlers(ctx context.Context) []Handler
}

// ManagerType returns the type of Manager interface. Can be used for implementing common.HasType.
//
// aura:api:stable
func ManagerType() interface{} {
	return (*Manager)(nil)
}
