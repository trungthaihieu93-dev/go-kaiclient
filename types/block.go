/*
 *  Copyright 2020 KardiaChain
 *  This file is part of the go-kardia library.
 *
 *  The go-kardia library is free software: you can redistribute it and/or modify
 *  it under the terms of the GNU Lesser General Public License as published by
 *  the Free Software Foundation, either version 3 of the License, or
 *  (at your option) any later version.
 *
 *  The go-kardia library is distributed in the hope that it will be useful,
 *  but WITHOUT ANY WARRANTY; without even the implied warranty of
 *  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 *  GNU Lesser General Public License for more details.
 *
 *  You should have received a copy of the GNU Lesser General Public License
 *  along with the go-kardia library. If not, see <http://www.gnu.org/licenses/>.
 */

package types

import (
	"encoding/json"
	"time"
)

type Header struct {
	Hash   string `json:"hash" bson:"blockHash"`
	Height uint64 `json:"height" bson:"height"`

	CommitHash      string    `json:"commitHash" bson:"commitHash"`
	GasLimit        uint64    `json:"gasLimit" bson:"gasLimit"`
	GasUsed         uint64    `json:"gasUsed" bson:"gasUsed"`
	Rewards         string    `json:"rewards" bson:"rewards"`
	NumTxs          uint64    `json:"numTxs" bson:"numTxs"`
	Time            time.Time `json:"time" bson:"time"`
	ProposerAddress string    `json:"proposerAddress" bson:"proposerAddress"`

	LastBlock string `json:"lastBlock" bson:"lastBlock"`

	DataHash     string `json:"dataHash" bson:"dataHash"`
	ReceiptsRoot string `json:"receiptsRoot" bson:"receiptsRoot"`
	LogsBloom    string `json:"logsBloom" bson:"logsBloom"`

	ValidatorHash     string `json:"validatorHash" bson:"validatorHash"`
	NextValidatorHash string `json:"nextValidatorHash" bson:"nextValidatorHash"` // validators for the next block
	ConsensusHash     string `json:"consensusHash" bson:"consensusHash"`
	AppHash           string `json:"appHash" bson:"appHash"`
	EvidenceHash      string `json:"evidenceHash" bson:"evidenceHash"`

	// Dual nodes
	NumDualEvents  uint64 `json:"numDualEvents" bson:"numDualEvents"`
	DualEventsHash string `json:"dualEventsHash" bson:"dualEventsHash"`
}

type Block struct {
	Hash   string `json:"hash,omitempty" bson:"hash"`
	Height uint64 `json:"height,omitempty" bson:"height"`

	CommitHash      string    `json:"commitHash,omitempty" bson:"commitHash"`
	GasLimit        uint64    `json:"gasLimit,omitempty" bson:"gasLimit"`
	GasUsed         uint64    `json:"gasUsed" bson:"gasUsed"`
	Rewards         string    `json:"rewards" bson:"rewards"`
	NumTxs          uint64    `json:"numTxs" bson:"numTxs"`
	Time            time.Time `json:"time,omitempty" bson:"time"`
	ProposerAddress string    `json:"proposerAddress,omitempty" bson:"proposerAddress"`

	LastBlock string `json:"lastBlock,omitempty" bson:"lastBlock"`

	DataHash     string `json:"dataHash,omitempty" bson:"dataHash"`
	ReceiptsRoot string `json:"receiptsRoot,omitempty" bson:"receiptsRoot"`
	LogsBloom    string `json:"logsBloom,omitempty" bson:"logsBloom"`

	ValidatorHash     string `json:"validatorHash,omitempty" bson:"validatorHash"`
	NextValidatorHash string `json:"nextValidatorHash,omitempty" bson:"nextValidatorHash"` // validators for the next block
	ConsensusHash     string `json:"consensusHash,omitempty" bson:"consensusHash"`
	AppHash           string `json:"appHash,omitempty" bson:"appHash"`
	EvidenceHash      string `json:"evidenceHash,omitempty" bson:"evidenceHash"`

	// Dual nodes
	NumDualEvents  uint64 `json:"numDualEvents,omitempty" bson:"numDualEvents"`
	DualEventsHash string `json:"dualEventsHash,omitempty" bson:"dualEventsHash"`

	Txs      []*Transaction `json:"txs,omitempty" bson:"-"`
	Receipts []*Receipt     `json:"receipts,omitempty" bson:"-"`
}

type VerifyBlockParam struct {
	VerifyTxCount   bool
	VerifyBlockHash bool
}

func (b *Block) String() string {
	data, err := json.Marshal(b)
	if err != nil {
		return ""
	}
	return string(data)
}
