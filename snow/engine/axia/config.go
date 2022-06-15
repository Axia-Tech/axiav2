// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package axia

import (
	"github.com/Axia-Tech/axiav2/snow"
	"github.com/Axia-Tech/axiav2/snow/consensus/axia"
	"github.com/Axia-Tech/axiav2/snow/engine/axia/vertex"
	"github.com/Axia-Tech/axiav2/snow/engine/common"
	"github.com/Axia-Tech/axiav2/snow/validators"
)

// Config wraps all the parameters needed for an axia engine
type Config struct {
	Ctx *snow.ConsensusContext
	common.AllGetsServer
	VM         vertex.DAGVM
	Manager    vertex.Manager
	Sender     common.Sender
	Validators validators.Set

	Params    axia.Parameters
	Consensus axia.Consensus
}
