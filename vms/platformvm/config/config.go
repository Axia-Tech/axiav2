// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package config

import (
	"time"

	"github.com/Axia-Tech/axiav2/chains"
	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/snow/engine/common"
	"github.com/Axia-Tech/axiav2/snow/uptime"
	"github.com/Axia-Tech/axiav2/snow/validators"
	"github.com/Axia-Tech/axiav2/vms/platformvm/reward"
)

// Struct collecting all foundational parameters of PlatformVM
type Config struct {
	// The node's chain manager
	Chains chains.Manager

	// Node's validator set maps subnetID -> validators of the subnet
	Validators validators.Manager

	// Provides access to subnet tracking
	SubnetTracker common.SubnetTracker

	// Provides access to the uptime manager as a thread safe data structure
	UptimeLockedCalculator uptime.LockedCalculator

	// True if the node is being run with staking enabled
	StakingEnabled bool

	// Set of subnets that this node is validating
	WhitelistedSubnets ids.Set

	// Fee that must be burned by every create staker transaction
	AddStakerTxFee uint64

	// Fee that is burned by every non-state creating transaction
	TxFee uint64

	// Fee that must be burned by every state creating transaction before AP3
	CreateAssetTxFee uint64

	// Fee that must be burned by every subnet creating transaction after AP3
	CreateSubnetTxFee uint64

	// Fee that must be burned by every blockchain creating transaction after AP3
	CreateBlockchainTxFee uint64

	// The minimum amount of tokens one must bond to be a validator
	MinValidatorStake uint64

	// The maximum amount of tokens that can be bonded on a validator
	MaxValidatorStake uint64

	// Minimum stake, in nAXC, that can be delegated on the primary network
	MinDelegatorStake uint64

	// Minimum fee that can be charged for delegation
	MinDelegationFee uint32

	// UptimePercentage is the minimum uptime required to be rewarded for staking
	UptimePercentage float64

	// Minimum amount of time to allow a staker to stake
	MinStakeDuration time.Duration

	// Maximum amount of time to allow a staker to stake
	MaxStakeDuration time.Duration

	// Config for the minting function
	RewardConfig reward.Config

	// Time of the AP3 network upgrade
	ApricotPhase3Time time.Time

	// Time of the AP4 network upgrade
	ApricotPhase4Time time.Time

	// Time of the AP5 network upgrade
	ApricotPhase5Time time.Time
}
