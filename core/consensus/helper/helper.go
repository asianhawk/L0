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

package helper

import (
	"github.com/bocheninc/L0/core/types"
)

// NewStack Create Stack instance
func NewStack() *Stack {
	return &Stack{}
}

// Stack Implenment consensus.IStack
type Stack struct {
}

// GetLastSeqNo Implenment consensus.IStack
func (stack *Stack) GetLastSeqNo() uint64 {
	return 0
}

// VerifyTxsInConsensus Implenment consensus.IStack
func (stack *Stack) VerifyTxsInConsensus(txs []*types.Transaction, primary bool) []*types.Transaction {
	return txs
}

// IterTransaction  Implenment consensus.IStack
func (stack *Stack) IterTransaction(func(*types.Transaction) bool) {

}

// Removes  Implenment consensus.IStack
func (stack *Stack) Removes([]*types.Transaction) {

}

// Len  Implenment consensus.IStack
func (stack *Stack) Len() int {
	return 0
}

// GetGroupingTxs  Implenment consensus.IStack
func (stack *Stack) GetGroupingTxs(maxSize, maxGroup uint64) [][]*types.Transaction {
	return nil
}
