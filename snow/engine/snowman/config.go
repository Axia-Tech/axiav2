// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/Axia-Tech/axiav2/snow"
	"github.com/Axia-Tech/axiav2/snow/consensus/snowball"
	"github.com/Axia-Tech/axiav2/snow/consensus/snowman"
	"github.com/Axia-Tech/axiav2/snow/engine/common"
	"github.com/Axia-Tech/axiav2/snow/engine/snowman/block"
	"github.com/Axia-Tech/axiav2/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
