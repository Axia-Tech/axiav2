// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package router

import (
	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/message"
	"github.com/Axia-Tech/axiav2/version"
)

var _ InboundHandler = InboundHandlerFunc(nil)

// InboundHandler handles inbound messages
type InboundHandler interface {
	HandleInbound(msg message.InboundMessage)
}

// The ExternalRouterFunc type is an adapter to allow the use of ordinary
// functions as ExternalRouters. If f is a function with the appropriate
// signature, ExternalRouterFunc(f) is an ExternalRouter that calls f.
type InboundHandlerFunc func(msg message.InboundMessage)

func (f InboundHandlerFunc) HandleInbound(msg message.InboundMessage) {
	f(msg)
}

// ExternalHandler handles messages from external parties
type ExternalHandler interface {
	InboundHandler

	Connected(nodeID ids.NodeID, nodeVersion version.Application)
	Disconnected(nodeID ids.NodeID)
}
