// Copyright (C) 2017, Beijing Bochen Technology Co.,Ltd.  All rights reserved.
//
// This file is part of L0
//
// The L0 is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The L0 is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package consensus

import "github.com/bocheninc/L0/core/types"

//BroadcastConsensus Define consensus data for broadcast
type BroadcastConsensus struct {
	To      string
	Payload []byte
}

// CommittedTxs Consensus output object
type CommittedTxs struct {
	SeqNos       []uint64
	Time         uint32
	Transactions []*types.Transaction
}

// Consenter Interface for plugin consenser
type Consenter interface {
	Start()
	Stop()
	RecvConsensus([]byte)
	BroadcastConsensusChannel() <-chan *BroadcastConsensus
	CommittedTxsChannel() <-chan *CommittedTxs
}

// ITxPool Interface for tx containter, input
type ITxPool interface {
	IterTransaction(func(*types.Transaction) bool)
	GetGroupingTxs(maxSize, maxGroup uint64) [][]*types.Transaction
	Removes([]*types.Transaction)
	Len() int
}

// IStack Interface for other function for plugin consenser
type IStack interface {
	VerifyTxsInConsensus(txs []*types.Transaction, primary bool) []*types.Transaction
	GetLastSeqNo() uint64
	ITxPool
}
