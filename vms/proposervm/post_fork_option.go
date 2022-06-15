// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"time"

	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/snow/choices"
	"github.com/Axia-Tech/axiav2/vms/proposervm/block"
)

var _ PostForkBlock = &postForkOption{}

// The parent of a *postForkOption must be a *postForkBlock.
type postForkOption struct {
	block.Block
	postForkCommonComponents

	timestamp time.Time
}

func (b *postForkOption) Timestamp() time.Time {
	if b.Status() == choices.Accepted {
		return b.vm.lastAcceptedTime
	}
	return b.timestamp
}

func (b *postForkOption) Accept() error {
	if err := b.acceptOuterBlk(); err != nil {
		return err
	}
	return b.acceptInnerBlk()
}

func (b *postForkOption) acceptOuterBlk() error {
	// Update in-memory references
	b.status = choices.Accepted
	b.vm.lastAcceptedHeight = b.Height()

	blkID := b.ID()
	delete(b.vm.verifiedBlocks, blkID)

	// Persist this block, its height index, and its status
	if err := b.vm.State.SetLastAccepted(blkID); err != nil {
		return err
	}
	return b.vm.storePostForkBlock(b)
}

func (b *postForkOption) acceptInnerBlk() error {
	// mark the inner block as accepted and all conflicting inner blocks as
	// rejected
	return b.vm.Tree.Accept(b.innerBlk)
}

func (b *postForkOption) Reject() error {
	// we do not reject the inner block here because that block may be contained
	// in the proposer block that causing this block to be rejected.

	delete(b.vm.verifiedBlocks, b.ID())
	b.status = choices.Rejected
	return nil
}

func (b *postForkOption) Status() choices.Status {
	if b.status == choices.Accepted && b.Height() > b.vm.lastAcceptedHeight {
		return choices.Processing
	}
	return b.status
}

func (b *postForkOption) Parent() ids.ID {
	return b.ParentID()
}

// If Verify returns nil, Accept or Reject is eventually called on [b] and
// [b.innerBlk].
func (b *postForkOption) Verify() error {
	parent, err := b.vm.getBlock(b.ParentID())
	if err != nil {
		return err
	}
	b.timestamp = parent.Timestamp()
	return parent.verifyPostForkOption(b)
}

func (b *postForkOption) verifyPreForkChild(child *preForkBlock) error {
	// A *preForkBlock's parent must be a *preForkBlock
	return errUnsignedChild
}

func (b *postForkOption) verifyPostForkChild(child *postForkBlock) error {
	parentTimestamp := b.Timestamp()
	parentPChainHeight, err := b.pChainHeight()
	if err != nil {
		return err
	}
	return b.postForkCommonComponents.Verify(
		parentTimestamp,
		parentPChainHeight,
		child,
	)
}

func (b *postForkOption) verifyPostForkOption(child *postForkOption) error {
	// A *postForkOption's parent can't be a *postForkOption
	return errUnexpectedBlockType
}

func (b *postForkOption) buildChild() (Block, error) {
	parentPChainHeight, err := b.pChainHeight()
	if err != nil {
		return nil, err
	}
	return b.postForkCommonComponents.buildChild(
		b.ID(),
		b.Timestamp(),
		parentPChainHeight,
	)
}

// This block's P-Chain height is its parent's P-Chain height
func (b *postForkOption) pChainHeight() (uint64, error) {
	parent, err := b.vm.getBlock(b.ParentID())
	if err != nil {
		return 0, err
	}
	return parent.pChainHeight()
}

func (b *postForkOption) setStatus(status choices.Status) {
	b.status = status
}

func (b *postForkOption) getStatelessBlk() block.Block {
	return b.Block
}
