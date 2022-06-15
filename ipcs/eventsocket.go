// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package ipcs

import (
	"errors"
	"os"
	"syscall"

	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/ipcs/socket"
	"github.com/Axia-Tech/axiav2/snow"
	"github.com/Axia-Tech/axiav2/utils/logging"
	"github.com/Axia-Tech/axiav2/utils/wrappers"
)

var _ snow.Acceptor = &EventSockets{}

// EventSockets is a set of named eventSockets
type EventSockets struct {
	consensusSocket *eventSocket
	decisionsSocket *eventSocket
}

// newEventSockets creates a *ChainIPCs with both consensus and decisions IPCs
func newEventSockets(ctx context, chainID ids.ID, consensusAcceptorGroup, decisionAcceptorGroup snow.AcceptorGroup) (*EventSockets, error) {
	consensusIPC, err := newEventIPCSocket(ctx, chainID, ipcConsensusIdentifier, consensusAcceptorGroup)
	if err != nil {
		return nil, err
	}

	decisionsIPC, err := newEventIPCSocket(ctx, chainID, ipcDecisionsIdentifier, decisionAcceptorGroup)
	if err != nil {
		return nil, err
	}

	return &EventSockets{
		consensusSocket: consensusIPC,
		decisionsSocket: decisionsIPC,
	}, nil
}

// Accept delivers a message to the underlying eventSockets
func (ipcs *EventSockets) Accept(ctx *snow.ConsensusContext, containerID ids.ID, container []byte) error {
	if ipcs.consensusSocket != nil {
		if err := ipcs.consensusSocket.Accept(ctx, containerID, container); err != nil {
			return err
		}
	}

	if ipcs.decisionsSocket != nil {
		if err := ipcs.decisionsSocket.Accept(ctx, containerID, container); err != nil {
			return err
		}
	}

	return nil
}

// stop closes the underlying eventSockets
func (ipcs *EventSockets) stop() error {
	errs := wrappers.Errs{}

	if ipcs.consensusSocket != nil {
		errs.Add(ipcs.consensusSocket.stop())
	}

	if ipcs.decisionsSocket != nil {
		errs.Add(ipcs.decisionsSocket.stop())
	}

	return errs.Err
}

// ConsensusURL returns the URL of socket receiving consensus events
func (ipcs *EventSockets) ConsensusURL() string {
	return ipcs.consensusSocket.URL()
}

// DecisionsURL returns the URL of socket receiving decisions events
func (ipcs *EventSockets) DecisionsURL() string {
	return ipcs.decisionsSocket.URL()
}

// eventSocket is a single IPC socket for a single chain
type eventSocket struct {
	url          string
	log          logging.Logger
	socket       *socket.Socket
	unregisterFn func() error
}

// newEventIPCSocket creates a *eventSocket for the given chain and
// EventDispatcher that writes to a local IPC socket
func newEventIPCSocket(ctx context, chainID ids.ID, name string, acceptorGroup snow.AcceptorGroup) (*eventSocket, error) {
	var (
		url     = ipcURL(ctx, chainID, name)
		ipcName = ipcIdentifierPrefix + "-" + name
	)

	err := os.Remove(url)
	if err != nil && !errors.Is(err, syscall.ENOENT) {
		return nil, err
	}

	eis := &eventSocket{
		log:    ctx.log,
		url:    url,
		socket: socket.NewSocket(url, ctx.log),
		unregisterFn: func() error {
			return acceptorGroup.DeregisterAcceptor(chainID, ipcName)
		},
	}

	if err := eis.socket.Listen(); err != nil {
		if err := eis.socket.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}

	if err := acceptorGroup.RegisterAcceptor(chainID, ipcName, eis, false); err != nil {
		if err := eis.stop(); err != nil {
			return nil, err
		}
		return nil, err
	}

	return eis, nil
}

// Accept delivers a message to the eventSocket
func (eis *eventSocket) Accept(_ *snow.ConsensusContext, _ ids.ID, container []byte) error {
	eis.socket.Send(container)
	return nil
}

// stop unregisters the event handler and closes the eventSocket
func (eis *eventSocket) stop() error {
	eis.log.Info("closing Chain IPC")
	errs := wrappers.Errs{}
	errs.Add(eis.unregisterFn(), eis.socket.Close())
	return errs.Err
}

// URL returns the URL of the socket
func (eis *eventSocket) URL() string {
	return eis.url
}