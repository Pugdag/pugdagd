// Copyright (c) 2014-2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package dagconfig

import (
	"math/big"

	"github.com/pugdag/pugdagd/domain/consensus/model/externalapi"
	"github.com/pugdag/pugdagd/domain/consensus/utils/blockheader"
	"github.com/pugdag/pugdagd/domain/consensus/utils/subnetworks"
	"github.com/pugdag/pugdagd/domain/consensus/utils/transactionhelper"
	"github.com/kaspanet/go-muhash"
)

var genesisTxOuts = []*externalapi.DomainTransactionOutput{}

var genesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                           // Varint
	0x00,                                           // OP-FALSE
	0x00, 0x00, 0x00, 0x00, 0x00, 0x19, 0xd6, 0x68, // Bitcoin genesis block hash 000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f
	0x9c, 0x09, 0x5a, 0xe1, 0x65, 0x83, 0x1e, 0x93, 
	0x4f, 0xf7, 0x63, 0xae, 0x46, 0xa2, 0xa6, 0xc1,
	0x72, 0xb3, 0xf1, 0xb6, 0x0a, 0x8c, 0xe2, 0x6f,
	0x46, 0x6f, 0x72, 0x30, 0x63, 0x6c, 0x6c, 0x20,
	0x74, 0x68, 0x65, 0x20, 0x73, 0x6d, 0x63, 0x6c,
	0x6c, 0x20, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x20, 0x6d, 0x69, 0x6e, 0x65, 0x72, 0x73, 0x20,
	0x69, 0x6e, 0x20, 0x74, 0x68, 0x65, 0x20, 0x77,
	0x6f, 0x72, 0x6b, 0x64, 0x2c, 0x20, 0x77, 0x65,
	0x20, 0x6b, 0x6f, 0x76, 0x65, 0x20, 0x62, 0x74,
	0x63, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6b, 0x61,
	0x73, 0x70, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x2e, 0x20, 0x52, 0x65, 0x61,
	0x61, 0x72, 0x65, 0x20, 0x6f, 0x6e, 0x65, 0x2e,
}

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network.
var genesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0, []*externalapi.DomainTransactionInput{}, genesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, genesisTxPayload)

// genesisHash is the hash of the first block in the block DAG for the main
// network (genesis block).
var genesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xe0, 0xf2, 0x59, 0x78, 0x24, 0xbe, 0xad, 0xd8, 0x0c, 0x16, 0x16, 0xb8, 0xeb, 0xc7, 0xbd, 0x63, 0x76, 0x1c, 0x68, 0xeb, 0x6d, 0xb4, 0xed, 0xb4, 0x5f, 0x86, 0xcf, 0x1a, 0x53, 0xef, 0xe9, 0x1b,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.

var genesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x0d, 0x4c, 0x07, 0x90, 0x33, 0xe4, 0x8a, 0xde, 0xc9, 0x4d, 0x2e, 0x86, 0xb6, 0xdc, 0x21, 0x29, 0x98, 0xc5, 0x1b, 0x9c, 0x8b, 0xea, 0xed, 0x75, 0x4b, 0x94, 0x18, 0x60, 0x90, 0xaf, 0x96, 0x46,
})

// genesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the main network.
var genesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		genesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1716053644637,
		0x1e7fffff,
		0x14582,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{genesisCoinbaseTx},
}

var devnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var devnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                               // Varint
	0x00,                                                                               // OP-FALSE
	0x6B, 0x61, 0x72, 0x6C, 0x73, 0x65, 0x6E, 0x2D, 0x64, 0x65, 0x76, 0x6E, 0x65, 0x74, // pugdag-devnet
}

// devnetGenesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the development network.
var devnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, devnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, devnetGenesisTxPayload)

// devGenesisHash is the hash of the first block in the block DAG for the development
// network (genesis block).
// pugdag devnet genesis hash
var devnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0xcb, 0x1b, 0x9e, 0x97, 0x2c, 0x04, 0x3e, 0xc9, 0x98, 0xc4, 0x36, 0x13, 0x46, 0x45, 0x04, 0xe1, 0x7d, 0xf2, 0xa4, 0x5a, 0x8a, 0x6a, 0xa1, 0x16, 0x21, 0xd9, 0x4b, 0x87, 0x6d, 0x69, 0xe0, 0xd4,
})

// devnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
// pugdag devnet genesit merkleroot
var devnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x5e, 0xab, 0x60, 0xd4, 0xaa, 0x01, 0x02, 0x97, 0x8b, 0xc6, 0x8b, 0x43, 0xc5, 0x4d, 0x22, 0x8b, 0x71, 0x38, 0xa4, 0x20, 0x54, 0x48, 0x84, 0x31, 0x96, 0x7b, 0xc7, 0xaa, 0x86, 0x51, 0xb0, 0xe9,
})

// devnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var devnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		devnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x11e9db49828,
		525264379,
		0x48e5e,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{devnetGenesisCoinbaseTx},
}

var simnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var simnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                               // Varint
	0x00,                                                                               // OP-FALSE
	0x6B, 0x61, 0x72, 0x6C, 0x73, 0x65, 0x6E, 0x2D, 0x73, 0x69, 0x6D, 0x6E, 0x65, 0x74, // pugdag-simnet
}

// simnetGenesisCoinbaseTx is the coinbase transaction for the simnet genesis block.
var simnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, simnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, simnetGenesisTxPayload)

// simnetGenesisHash is the hash of the first block in the block DAG for
// the simnet (genesis block).
// pugdag simnet genesis hash
var simnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x8f, 0xe8, 0xb0, 0xf8, 0x04, 0x32, 0x52, 0xfd, 0xe9, 0x27, 0x09, 0x26, 0x33, 0x93, 0x79, 0x20, 0x94, 0x79, 0x5f, 0x34, 0x4e, 0xc2, 0x52, 0xc9, 0xb7, 0x56, 0xd1, 0xd1, 0x3e, 0x0d, 0xfe, 0x11,
})

// simnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for the devopment network.
// pugdag simnet merkel root
var simnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x04, 0xcf, 0x01, 0xcf, 0xc2, 0x9e, 0xce, 0x66, 0x55, 0x43, 0xd6, 0xbf, 0x5e, 0xc0, 0x99, 0x98, 0x8d, 0x4d, 0x3b, 0xaf, 0x19, 0xf2, 0x8f, 0xb0, 0xf9, 0xd4, 0xfa, 0xe3, 0x41, 0x20, 0x85, 0x17,
})

// simnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for the development network.
var simnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		simnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		0x17c5f62fbb6,
		0x207fffff,
		0x2,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{simnetGenesisCoinbaseTx},
}

var testnetGenesisTxOuts = []*externalapi.DomainTransactionOutput{}

var testnetGenesisTxPayload = []byte{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Blue score
	0x00, 0xE1, 0xF5, 0x05, 0x00, 0x00, 0x00, 0x00, // Subsidy
	0x00, 0x00, // Script version
	0x01,                                                                                     // Varint
	0x00,                                                                                     // OP-FALSE
	0x6B, 0x61, 0x72, 0x6C, 0x73, 0x65, 0x6E, 0x2D, 0x74, 0x65, 0x73, 0x74, 0x6E, 0x65, 0x74, // pugdag-testnet
}

// testnetGenesisCoinbaseTx is the coinbase transaction for the testnet genesis block.
var testnetGenesisCoinbaseTx = transactionhelper.NewSubnetworkTransaction(0,
	[]*externalapi.DomainTransactionInput{}, testnetGenesisTxOuts,
	&subnetworks.SubnetworkIDCoinbase, 0, testnetGenesisTxPayload)

// testnetGenesisHash is the hash of the first block in the block DAG for the test
// network (genesis block).
// pugdag testnet genesis hash
var testnetGenesisHash = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x26, 0xe1, 0x79, 0x6d, 0x6d, 0x9a, 0x12, 0xb9, 0x64, 0xc0, 0x5a, 0xf8, 0xf3, 0xac, 0x6f, 0xfd, 0xd8, 0x28, 0xb4, 0x3a, 0x09, 0x95, 0xeb, 0xdd, 0xcd, 0x26, 0x07, 0x1a, 0x61, 0x04, 0xb8, 0xdc,
})

// testnetGenesisMerkleRoot is the hash of the first transaction in the genesis block
// for testnet.
// pugdag testnet merkle root
var testnetGenesisMerkleRoot = externalapi.NewDomainHashFromByteArray(&[externalapi.DomainHashSize]byte{
	0x06, 0x8e, 0x09, 0xad, 0xab, 0x75, 0x3b, 0x8c, 0x0d, 0x91, 0x61, 0xb9, 0xde, 0x39, 0x5a, 0x4a, 0xa2, 0x38, 0xcb, 0xa8, 0x9b, 0xdc, 0x9b, 0x03, 0x67, 0xf6, 0xab, 0xdf, 0xe9, 0xd0, 0x0b, 0xe0,
})

// testnetGenesisBlock defines the genesis block of the block DAG which serves as the
// public transaction ledger for testnet.
var testnetGenesisBlock = externalapi.DomainBlock{
	Header: blockheader.NewImmutableBlockHeader(
		0,
		[]externalapi.BlockLevelParents{},
		testnetGenesisMerkleRoot,
		&externalapi.DomainHash{},
		externalapi.NewDomainHashFromByteArray(muhash.EmptyMuHashHash.AsArray()),
		1706347351518,
		511705087,
		83330,
		0,
		0,
		big.NewInt(0),
		&externalapi.DomainHash{},
	),
	Transactions: []*externalapi.DomainTransaction{testnetGenesisCoinbaseTx},
}
