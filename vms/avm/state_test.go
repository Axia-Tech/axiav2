// Copyright (C) 2019-2021, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"math"
	"testing"

	"github.com/Axia-Tech/axiav2/ids"
	"github.com/Axia-Tech/axiav2/snow/choices"
	"github.com/Axia-Tech/axiav2/snow/engine/common"
	"github.com/Axia-Tech/axiav2/utils/crypto"
	"github.com/Axia-Tech/axiav2/utils/units"
	"github.com/Axia-Tech/axiav2/vms/avm/txs"
	"github.com/Axia-Tech/axiav2/vms/components/axc"
	"github.com/Axia-Tech/axiav2/vms/secp256k1fx"
)

func TestSetsAndGets(t *testing.T) {
	_, _, vm, _ := GenesisVMWithArgs(
		t,
		[]*common.Fx{{
			ID: ids.GenerateTestID(),
			Fx: &FxTest{
				InitializeF: func(vmIntf interface{}) error {
					vm := vmIntf.(secp256k1fx.VM)
					return vm.CodecRegistry().RegisterType(&axc.TestVerifiable{})
				},
			},
		}},
		nil,
	)
	ctx := vm.ctx
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		ctx.Lock.Unlock()
	}()

	state := vm.state

	utxo := &axc.UTXO{
		UTXOID: axc.UTXOID{
			TxID:        ids.Empty,
			OutputIndex: 1,
		},
		Asset: axc.Asset{ID: ids.Empty},
		Out:   &axc.TestVerifiable{},
	}

	tx := &txs.Tx{UnsignedTx: &txs.BaseTx{BaseTx: axc.BaseTx{
		NetworkID:    networkID,
		BlockchainID: chainID,
		Ins: []*axc.TransferableInput{{
			UTXOID: axc.UTXOID{
				TxID:        ids.Empty,
				OutputIndex: 0,
			},
			Asset: axc.Asset{ID: assetID},
			In: &secp256k1fx.TransferInput{
				Amt: 20 * units.KiloAxc,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{
						0,
					},
				},
			},
		}},
	}}}
	if err := tx.SignSECP256K1Fx(vm.parser.Codec(), [][]*crypto.PrivateKeySECP256K1R{{keys[0]}}); err != nil {
		t.Fatal(err)
	}

	if err := state.PutUTXO(ids.Empty, utxo); err != nil {
		t.Fatal(err)
	}
	if err := state.PutTx(ids.Empty, tx); err != nil {
		t.Fatal(err)
	}
	if err := state.PutStatus(ids.Empty, choices.Accepted); err != nil {
		t.Fatal(err)
	}

	resultUTXO, err := state.GetUTXO(ids.Empty)
	if err != nil {
		t.Fatal(err)
	}
	resultTx, err := state.GetTx(ids.Empty)
	if err != nil {
		t.Fatal(err)
	}
	resultStatus, err := state.GetStatus(ids.Empty)
	if err != nil {
		t.Fatal(err)
	}

	if resultUTXO.OutputIndex != 1 {
		t.Fatalf("Wrong UTXO returned")
	}
	if resultTx.ID() != tx.ID() {
		t.Fatalf("Wrong Tx returned")
	}
	if resultStatus != choices.Accepted {
		t.Fatalf("Wrong Status returned")
	}
}

func TestFundingNoAddresses(t *testing.T) {
	_, _, vm, _ := GenesisVMWithArgs(
		t,
		[]*common.Fx{{
			ID: ids.GenerateTestID(),
			Fx: &FxTest{
				InitializeF: func(vmIntf interface{}) error {
					vm := vmIntf.(secp256k1fx.VM)
					return vm.CodecRegistry().RegisterType(&axc.TestVerifiable{})
				},
			},
		}},
		nil,
	)
	ctx := vm.ctx
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		ctx.Lock.Unlock()
	}()

	state := vm.state

	utxo := &axc.UTXO{
		UTXOID: axc.UTXOID{
			TxID:        ids.Empty,
			OutputIndex: 1,
		},
		Asset: axc.Asset{ID: ids.Empty},
		Out:   &axc.TestVerifiable{},
	}

	if err := state.PutUTXO(utxo.InputID(), utxo); err != nil {
		t.Fatal(err)
	}
	if err := state.DeleteUTXO(utxo.InputID()); err != nil {
		t.Fatal(err)
	}
}

func TestFundingAddresses(t *testing.T) {
	_, _, vm, _ := GenesisVMWithArgs(
		t,
		[]*common.Fx{{
			ID: ids.GenerateTestID(),
			Fx: &FxTest{
				InitializeF: func(vmIntf interface{}) error {
					vm := vmIntf.(secp256k1fx.VM)
					return vm.CodecRegistry().RegisterType(&axc.TestAddressable{})
				},
			},
		}},
		nil,
	)
	ctx := vm.ctx
	defer func() {
		if err := vm.Shutdown(); err != nil {
			t.Fatal(err)
		}
		ctx.Lock.Unlock()
	}()

	state := vm.state

	utxo := &axc.UTXO{
		UTXOID: axc.UTXOID{
			TxID:        ids.Empty,
			OutputIndex: 1,
		},
		Asset: axc.Asset{ID: ids.Empty},
		Out: &axc.TestAddressable{
			Addrs: [][]byte{{0}},
		},
	}

	if err := state.PutUTXO(utxo.InputID(), utxo); err != nil {
		t.Fatal(err)
	}
	utxos, err := state.UTXOIDs([]byte{0}, ids.Empty, math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	if len(utxos) != 1 {
		t.Fatalf("Should have returned 1 utxoIDs")
	}
	if utxoID := utxos[0]; utxoID != utxo.InputID() {
		t.Fatalf("Returned wrong utxoID")
	}
	if err := state.DeleteUTXO(utxo.InputID()); err != nil {
		t.Fatal(err)
	}
	utxos, err = state.UTXOIDs([]byte{0}, ids.Empty, math.MaxInt32)
	if err != nil {
		t.Fatal(err)
	}
	if len(utxos) != 0 {
		t.Fatalf("Should have returned 0 utxoIDs")
	}
}
