// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/snow/consensus/snowstorm"
	"github.com/Axia-Tech/axiav2/snow/engine/common"
)

// DAGVM defines the minimum functionality that an axia VM must
// implement
type DAGVM interface {
	common.VM

	// Return any transactions that have not been sent to consensus yet
	PendingTxs() []snowstorm.Tx

	// Convert a stream of bytes to a transaction or return an error
	ParseTx(tx []byte) (snowstorm.Tx, error)

	// Retrieve a transaction that was submitted previously
	GetTx(ids.ID) (snowstorm.Tx, error)
}
