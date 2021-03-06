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

// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

/*
Package lbft is a generated protocol buffer package.

It is generated from these files:
	message.proto

It has these top-level messages:
	Request
	RequestBatch
	PrePrepare
	Prepare
	Commit
	Committed
	FetchCommitted
	ViewChange
	NullRequest
	Message
*/
package lbft

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Time        uint32 `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Transaction []byte `protobuf:"bytes,2,opt,name=transaction,proto3" json:"transaction,omitempty"`
	FromChain   string `protobuf:"bytes,3,opt,name=fromChain" json:"fromChain,omitempty"`
	ToChain     string `protobuf:"bytes,4,opt,name=toChain" json:"toChain,omitempty"`
	Nonce       uint32 `protobuf:"varint,5,opt,name=nonce" json:"nonce,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Request) GetTransaction() []byte {
	if m != nil {
		return m.Transaction
	}
	return nil
}

func (m *Request) GetFromChain() string {
	if m != nil {
		return m.FromChain
	}
	return ""
}

func (m *Request) GetToChain() string {
	if m != nil {
		return m.ToChain
	}
	return ""
}

func (m *Request) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

type RequestBatch struct {
	Time     uint32     `protobuf:"varint,1,opt,name=time" json:"time,omitempty"`
	Requests []*Request `protobuf:"bytes,2,rep,name=requests" json:"requests,omitempty"`
	Id       int64      `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
}

func (m *RequestBatch) Reset()                    { *m = RequestBatch{} }
func (m *RequestBatch) String() string            { return proto.CompactTextString(m) }
func (*RequestBatch) ProtoMessage()               {}
func (*RequestBatch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RequestBatch) GetTime() uint32 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *RequestBatch) GetRequests() []*Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func (m *RequestBatch) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type PrePrepare struct {
	Name      string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PrimaryID string        `protobuf:"bytes,2,opt,name=primaryID" json:"primaryID,omitempty"`
	Chain     string        `protobuf:"bytes,3,opt,name=chain" json:"chain,omitempty"`
	ReplicaID string        `protobuf:"bytes,4,opt,name=replicaID" json:"replicaID,omitempty"`
	SeqNo     uint64        `protobuf:"varint,5,opt,name=seqNo" json:"seqNo,omitempty"`
	Digest    string        `protobuf:"bytes,6,opt,name=digest" json:"digest,omitempty"`
	Quorum    uint64        `protobuf:"varint,7,opt,name=quorum" json:"quorum,omitempty"`
	Requests  *RequestBatch `protobuf:"bytes,8,opt,name=requests" json:"requests,omitempty"`
}

func (m *PrePrepare) Reset()                    { *m = PrePrepare{} }
func (m *PrePrepare) String() string            { return proto.CompactTextString(m) }
func (*PrePrepare) ProtoMessage()               {}
func (*PrePrepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PrePrepare) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PrePrepare) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *PrePrepare) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *PrePrepare) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *PrePrepare) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *PrePrepare) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *PrePrepare) GetQuorum() uint64 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

func (m *PrePrepare) GetRequests() *RequestBatch {
	if m != nil {
		return m.Requests
	}
	return nil
}

type Prepare struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PrimaryID string `protobuf:"bytes,2,opt,name=primaryID" json:"primaryID,omitempty"`
	Chain     string `protobuf:"bytes,3,opt,name=chain" json:"chain,omitempty"`
	ReplicaID string `protobuf:"bytes,4,opt,name=replicaID" json:"replicaID,omitempty"`
	SeqNo     uint64 `protobuf:"varint,5,opt,name=seqNo" json:"seqNo,omitempty"`
	Digest    string `protobuf:"bytes,6,opt,name=digest" json:"digest,omitempty"`
	Quorum    uint64 `protobuf:"varint,7,opt,name=quorum" json:"quorum,omitempty"`
}

func (m *Prepare) Reset()                    { *m = Prepare{} }
func (m *Prepare) String() string            { return proto.CompactTextString(m) }
func (*Prepare) ProtoMessage()               {}
func (*Prepare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Prepare) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Prepare) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *Prepare) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Prepare) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *Prepare) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *Prepare) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Prepare) GetQuorum() uint64 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

type Commit struct {
	Name      string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PrimaryID string `protobuf:"bytes,2,opt,name=primaryID" json:"primaryID,omitempty"`
	Chain     string `protobuf:"bytes,3,opt,name=chain" json:"chain,omitempty"`
	ReplicaID string `protobuf:"bytes,4,opt,name=replicaID" json:"replicaID,omitempty"`
	SeqNo     uint64 `protobuf:"varint,5,opt,name=seqNo" json:"seqNo,omitempty"`
	Digest    string `protobuf:"bytes,6,opt,name=digest" json:"digest,omitempty"`
	Quorum    uint64 `protobuf:"varint,7,opt,name=quorum" json:"quorum,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Commit) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Commit) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *Commit) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Commit) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *Commit) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *Commit) GetDigest() string {
	if m != nil {
		return m.Digest
	}
	return ""
}

func (m *Commit) GetQuorum() uint64 {
	if m != nil {
		return m.Quorum
	}
	return 0
}

type Committed struct {
	Name         string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	PrimaryID    string        `protobuf:"bytes,2,opt,name=primaryID" json:"primaryID,omitempty"`
	Chain        string        `protobuf:"bytes,3,opt,name=chain" json:"chain,omitempty"`
	ReplicaID    string        `protobuf:"bytes,4,opt,name=replicaID" json:"replicaID,omitempty"`
	SeqNo        uint64        `protobuf:"varint,5,opt,name=seqNo" json:"seqNo,omitempty"`
	RequestBatch *RequestBatch `protobuf:"bytes,6,opt,name=requestBatch" json:"requestBatch,omitempty"`
}

func (m *Committed) Reset()                    { *m = Committed{} }
func (m *Committed) String() string            { return proto.CompactTextString(m) }
func (*Committed) ProtoMessage()               {}
func (*Committed) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Committed) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Committed) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *Committed) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *Committed) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *Committed) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

func (m *Committed) GetRequestBatch() *RequestBatch {
	if m != nil {
		return m.RequestBatch
	}
	return nil
}

type FetchCommitted struct {
	Chain     string `protobuf:"bytes,1,opt,name=chain" json:"chain,omitempty"`
	ReplicaID string `protobuf:"bytes,2,opt,name=replicaID" json:"replicaID,omitempty"`
	SeqNo     uint64 `protobuf:"varint,3,opt,name=seqNo" json:"seqNo,omitempty"`
}

func (m *FetchCommitted) Reset()                    { *m = FetchCommitted{} }
func (m *FetchCommitted) String() string            { return proto.CompactTextString(m) }
func (*FetchCommitted) ProtoMessage()               {}
func (*FetchCommitted) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FetchCommitted) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *FetchCommitted) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *FetchCommitted) GetSeqNo() uint64 {
	if m != nil {
		return m.SeqNo
	}
	return 0
}

type ViewChange struct {
	ReplicaID string `protobuf:"bytes,1,opt,name=replicaID" json:"replicaID,omitempty"`
	Chain     string `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	Priority  int64  `protobuf:"varint,3,opt,name=priority" json:"priority,omitempty"`
	PrimaryID string `protobuf:"bytes,4,opt,name=primaryID" json:"primaryID,omitempty"`
	H         uint64 `protobuf:"varint,5,opt,name=h" json:"h,omitempty"`
}

func (m *ViewChange) Reset()                    { *m = ViewChange{} }
func (m *ViewChange) String() string            { return proto.CompactTextString(m) }
func (*ViewChange) ProtoMessage()               {}
func (*ViewChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ViewChange) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *ViewChange) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *ViewChange) GetPriority() int64 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *ViewChange) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *ViewChange) GetH() uint64 {
	if m != nil {
		return m.H
	}
	return 0
}

type NullRequest struct {
	ReplicaID string `protobuf:"bytes,1,opt,name=replicaID" json:"replicaID,omitempty"`
	Chain     string `protobuf:"bytes,2,opt,name=chain" json:"chain,omitempty"`
	PrimaryID string `protobuf:"bytes,3,opt,name=primaryID" json:"primaryID,omitempty"`
	H         uint64 `protobuf:"varint,4,opt,name=h" json:"h,omitempty"`
}

func (m *NullRequest) Reset()                    { *m = NullRequest{} }
func (m *NullRequest) String() string            { return proto.CompactTextString(m) }
func (*NullRequest) ProtoMessage()               {}
func (*NullRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *NullRequest) GetReplicaID() string {
	if m != nil {
		return m.ReplicaID
	}
	return ""
}

func (m *NullRequest) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *NullRequest) GetPrimaryID() string {
	if m != nil {
		return m.PrimaryID
	}
	return ""
}

func (m *NullRequest) GetH() uint64 {
	if m != nil {
		return m.H
	}
	return 0
}

type Message struct {
	// Types that are valid to be assigned to Payload:
	//	*Message_RequestBatch
	//	*Message_PrePrepare
	//	*Message_Prepare
	//	*Message_Commit
	//	*Message_Committed
	//	*Message_FetchCommitted
	//	*Message_Viewchange
	//	*Message_NullReqest
	Payload isMessage_Payload `protobuf_oneof:"payload"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_RequestBatch struct {
	RequestBatch *RequestBatch `protobuf:"bytes,1,opt,name=requestBatch,oneof"`
}
type Message_PrePrepare struct {
	PrePrepare *PrePrepare `protobuf:"bytes,2,opt,name=prePrepare,oneof"`
}
type Message_Prepare struct {
	Prepare *Prepare `protobuf:"bytes,3,opt,name=prepare,oneof"`
}
type Message_Commit struct {
	Commit *Commit `protobuf:"bytes,4,opt,name=commit,oneof"`
}
type Message_Committed struct {
	Committed *Committed `protobuf:"bytes,5,opt,name=committed,oneof"`
}
type Message_FetchCommitted struct {
	FetchCommitted *FetchCommitted `protobuf:"bytes,6,opt,name=fetchCommitted,oneof"`
}
type Message_Viewchange struct {
	Viewchange *ViewChange `protobuf:"bytes,7,opt,name=viewchange,oneof"`
}
type Message_NullReqest struct {
	NullReqest *NullRequest `protobuf:"bytes,8,opt,name=nullReqest,oneof"`
}

func (*Message_RequestBatch) isMessage_Payload()   {}
func (*Message_PrePrepare) isMessage_Payload()     {}
func (*Message_Prepare) isMessage_Payload()        {}
func (*Message_Commit) isMessage_Payload()         {}
func (*Message_Committed) isMessage_Payload()      {}
func (*Message_FetchCommitted) isMessage_Payload() {}
func (*Message_Viewchange) isMessage_Payload()     {}
func (*Message_NullReqest) isMessage_Payload()     {}

func (m *Message) GetPayload() isMessage_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Message) GetRequestBatch() *RequestBatch {
	if x, ok := m.GetPayload().(*Message_RequestBatch); ok {
		return x.RequestBatch
	}
	return nil
}

func (m *Message) GetPrePrepare() *PrePrepare {
	if x, ok := m.GetPayload().(*Message_PrePrepare); ok {
		return x.PrePrepare
	}
	return nil
}

func (m *Message) GetPrepare() *Prepare {
	if x, ok := m.GetPayload().(*Message_Prepare); ok {
		return x.Prepare
	}
	return nil
}

func (m *Message) GetCommit() *Commit {
	if x, ok := m.GetPayload().(*Message_Commit); ok {
		return x.Commit
	}
	return nil
}

func (m *Message) GetCommitted() *Committed {
	if x, ok := m.GetPayload().(*Message_Committed); ok {
		return x.Committed
	}
	return nil
}

func (m *Message) GetFetchCommitted() *FetchCommitted {
	if x, ok := m.GetPayload().(*Message_FetchCommitted); ok {
		return x.FetchCommitted
	}
	return nil
}

func (m *Message) GetViewchange() *ViewChange {
	if x, ok := m.GetPayload().(*Message_Viewchange); ok {
		return x.Viewchange
	}
	return nil
}

func (m *Message) GetNullReqest() *NullRequest {
	if x, ok := m.GetPayload().(*Message_NullReqest); ok {
		return x.NullReqest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Message) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Message_OneofMarshaler, _Message_OneofUnmarshaler, _Message_OneofSizer, []interface{}{
		(*Message_RequestBatch)(nil),
		(*Message_PrePrepare)(nil),
		(*Message_Prepare)(nil),
		(*Message_Commit)(nil),
		(*Message_Committed)(nil),
		(*Message_FetchCommitted)(nil),
		(*Message_Viewchange)(nil),
		(*Message_NullReqest)(nil),
	}
}

func _Message_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.RequestBatch); err != nil {
			return err
		}
	case *Message_PrePrepare:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.PrePrepare); err != nil {
			return err
		}
	case *Message_Prepare:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Prepare); err != nil {
			return err
		}
	case *Message_Commit:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Commit); err != nil {
			return err
		}
	case *Message_Committed:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Committed); err != nil {
			return err
		}
	case *Message_FetchCommitted:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FetchCommitted); err != nil {
			return err
		}
	case *Message_Viewchange:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Viewchange); err != nil {
			return err
		}
	case *Message_NullReqest:
		b.EncodeVarint(8<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.NullReqest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Message.Payload has unexpected type %T", x)
	}
	return nil
}

func _Message_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Message)
	switch tag {
	case 1: // payload.requestBatch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RequestBatch)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_RequestBatch{msg}
		return true, err
	case 2: // payload.prePrepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(PrePrepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_PrePrepare{msg}
		return true, err
	case 3: // payload.prepare
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Prepare)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Prepare{msg}
		return true, err
	case 4: // payload.commit
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Commit)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Commit{msg}
		return true, err
	case 5: // payload.committed
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Committed)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Committed{msg}
		return true, err
	case 6: // payload.fetchCommitted
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(FetchCommitted)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_FetchCommitted{msg}
		return true, err
	case 7: // payload.viewchange
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ViewChange)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_Viewchange{msg}
		return true, err
	case 8: // payload.nullReqest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NullRequest)
		err := b.DecodeMessage(msg)
		m.Payload = &Message_NullReqest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Message_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Message)
	// payload
	switch x := m.Payload.(type) {
	case *Message_RequestBatch:
		s := proto.Size(x.RequestBatch)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_PrePrepare:
		s := proto.Size(x.PrePrepare)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Prepare:
		s := proto.Size(x.Prepare)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Commit:
		s := proto.Size(x.Commit)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Committed:
		s := proto.Size(x.Committed)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_FetchCommitted:
		s := proto.Size(x.FetchCommitted)
		n += proto.SizeVarint(6<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_Viewchange:
		s := proto.Size(x.Viewchange)
		n += proto.SizeVarint(7<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Message_NullReqest:
		s := proto.Size(x.NullReqest)
		n += proto.SizeVarint(8<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Request)(nil), "lbft.Request")
	proto.RegisterType((*RequestBatch)(nil), "lbft.RequestBatch")
	proto.RegisterType((*PrePrepare)(nil), "lbft.PrePrepare")
	proto.RegisterType((*Prepare)(nil), "lbft.Prepare")
	proto.RegisterType((*Commit)(nil), "lbft.Commit")
	proto.RegisterType((*Committed)(nil), "lbft.Committed")
	proto.RegisterType((*FetchCommitted)(nil), "lbft.FetchCommitted")
	proto.RegisterType((*ViewChange)(nil), "lbft.ViewChange")
	proto.RegisterType((*NullRequest)(nil), "lbft.NullRequest")
	proto.RegisterType((*Message)(nil), "lbft.Message")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 607 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xbd, 0x6e, 0xdb, 0x3c,
	0x14, 0x35, 0x25, 0xc5, 0x8a, 0xae, 0xec, 0x7c, 0x5f, 0x89, 0xa0, 0x10, 0x8a, 0x0e, 0x82, 0x86,
	0xc2, 0x59, 0x5c, 0x40, 0x01, 0x8a, 0x4e, 0x1d, 0x92, 0xa0, 0x70, 0x86, 0x06, 0x01, 0x87, 0x0e,
	0x05, 0x3a, 0x30, 0x12, 0x13, 0x11, 0xb0, 0x7e, 0x4c, 0xd1, 0x0d, 0xfc, 0x06, 0x5d, 0xfa, 0x2c,
	0x5d, 0xda, 0xb9, 0xaf, 0xd3, 0xc7, 0x28, 0x44, 0xd2, 0xfa, 0x31, 0xec, 0xa5, 0x4b, 0x91, 0x4d,
	0x97, 0x3c, 0x87, 0xf7, 0x1c, 0xde, 0xcb, 0x2b, 0x98, 0xe6, 0xac, 0xae, 0xe9, 0x03, 0x9b, 0x57,
	0xa2, 0x94, 0x25, 0x76, 0x96, 0x77, 0xf7, 0x32, 0xfa, 0x86, 0xc0, 0x25, 0x6c, 0xb5, 0x66, 0xb5,
	0xc4, 0x18, 0x1c, 0xc9, 0x73, 0x16, 0xa0, 0x10, 0xcd, 0xa6, 0x44, 0x7d, 0xe3, 0x10, 0x7c, 0x29,
	0x68, 0x51, 0xd3, 0x44, 0xf2, 0xb2, 0x08, 0xac, 0x10, 0xcd, 0x26, 0xa4, 0xbf, 0x84, 0x5f, 0x82,
	0x77, 0x2f, 0xca, 0xfc, 0x32, 0xa3, 0xbc, 0x08, 0xec, 0x10, 0xcd, 0x3c, 0xd2, 0x2d, 0xe0, 0x00,
	0x5c, 0x59, 0xea, 0x3d, 0x47, 0xed, 0x6d, 0x43, 0x7c, 0x0a, 0x47, 0x45, 0x59, 0x24, 0x2c, 0x38,
	0x52, 0xe9, 0x74, 0x10, 0x7d, 0x86, 0x89, 0x91, 0x73, 0x41, 0x65, 0x92, 0xed, 0xd5, 0x74, 0x06,
	0xc7, 0x42, 0x63, 0xea, 0xc0, 0x0a, 0xed, 0x99, 0x1f, 0x4f, 0xe7, 0x8d, 0x99, 0xb9, 0x61, 0x92,
	0x76, 0x1b, 0x9f, 0x80, 0xc5, 0x53, 0xa5, 0xca, 0x26, 0x16, 0x4f, 0xa3, 0xdf, 0x08, 0xe0, 0x56,
	0xb0, 0x5b, 0xc1, 0x2a, 0x2a, 0x58, 0x73, 0x7a, 0x41, 0xcd, 0xe9, 0x1e, 0x51, 0xdf, 0x8d, 0x9f,
	0x4a, 0xf0, 0x9c, 0x8a, 0xcd, 0xf5, 0x95, 0xf2, 0xeb, 0x91, 0x6e, 0xa1, 0x51, 0x9d, 0xf4, 0x9c,
	0xea, 0xa0, 0xe1, 0x08, 0x56, 0x2d, 0x79, 0x42, 0xaf, 0xaf, 0x8c, 0xcf, 0x6e, 0xa1, 0xe1, 0xd4,
	0x6c, 0x75, 0x53, 0x2a, 0xa7, 0x0e, 0xd1, 0x01, 0x7e, 0x0e, 0xe3, 0x94, 0x3f, 0xb0, 0x5a, 0x06,
	0x63, 0x45, 0x30, 0x51, 0xb3, 0xbe, 0x5a, 0x97, 0x62, 0x9d, 0x07, 0xae, 0x82, 0x9b, 0x08, 0xcf,
	0x7b, 0xae, 0x8f, 0x43, 0x34, 0xf3, 0x63, 0x3c, 0x70, 0xad, 0xee, 0xab, 0xb3, 0x1e, 0xfd, 0x44,
	0xe0, 0x3e, 0x41, 0x9f, 0xd1, 0x0f, 0x04, 0xe3, 0xcb, 0x32, 0xcf, 0xb9, 0x7c, 0x52, 0xb2, 0x7f,
	0x21, 0xf0, 0xb4, 0x6c, 0xc9, 0xd2, 0x7f, 0xaa, 0xfc, 0x0d, 0x4c, 0x44, 0xaf, 0x25, 0x94, 0xfe,
	0xfd, 0xcd, 0x32, 0xc0, 0x45, 0x9f, 0xe0, 0xe4, 0x3d, 0x93, 0x49, 0xd6, 0xb9, 0x68, 0x35, 0xa1,
	0x83, 0x9a, 0xac, 0x83, 0x9a, 0xec, 0x9e, 0xa6, 0xe8, 0x2b, 0x02, 0xf8, 0xc8, 0xd9, 0xe3, 0x65,
	0x46, 0x8b, 0x07, 0x36, 0x3c, 0x02, 0xed, 0x39, 0x42, 0xa7, 0xb5, 0xfa, 0x69, 0x5f, 0xc0, 0x71,
	0x25, 0x78, 0x29, 0xb8, 0xdc, 0x98, 0x07, 0xdd, 0xc6, 0xc3, 0xab, 0x75, 0x76, 0xaf, 0x76, 0x02,
	0x28, 0x33, 0x57, 0x84, 0xb2, 0x28, 0x07, 0xff, 0x66, 0xbd, 0x5c, 0x6e, 0x87, 0xde, 0xdf, 0x48,
	0x19, 0xa4, 0xb3, 0xf7, 0xa6, 0x73, 0xb6, 0xe9, 0xbe, 0xdb, 0xe0, 0x7e, 0xd0, 0x83, 0x17, 0xbf,
	0xdd, 0xa9, 0x0c, 0x3a, 0x54, 0x99, 0xc5, 0x68, 0x58, 0x1b, 0x1c, 0x03, 0x54, 0xed, 0xd8, 0x52,
	0x62, 0xfc, 0xf8, 0x7f, 0xcd, 0xeb, 0xc6, 0xd9, 0x62, 0x44, 0x7a, 0x28, 0x7c, 0x06, 0x6e, 0x65,
	0x08, 0xb6, 0x22, 0x4c, 0x5b, 0x82, 0x41, 0x6f, 0xf7, 0xf1, 0x2b, 0x18, 0x27, 0xaa, 0xea, 0x4a,
	0xb7, 0x1f, 0x4f, 0x34, 0x52, 0x77, 0xc2, 0x62, 0x44, 0xcc, 0x2e, 0x7e, 0x0d, 0x5e, 0xb2, 0xed,
	0x0e, 0x75, 0xa3, 0x7e, 0xfc, 0x5f, 0x1f, 0x2a, 0x59, 0xba, 0x18, 0x91, 0x0e, 0x83, 0xdf, 0xc1,
	0xc9, 0xfd, 0xa0, 0xa7, 0x4c, 0x37, 0x9e, 0x6a, 0xd6, 0xb0, 0xdf, 0x16, 0x23, 0xb2, 0x83, 0x6e,
	0x7c, 0x7f, 0xe1, 0xec, 0x31, 0x51, 0x6d, 0xa3, 0x5e, 0x5c, 0xeb, 0xbb, 0x6b, 0xa7, 0xc6, 0x77,
	0x87, 0xc2, 0xe7, 0x00, 0x85, 0x2e, 0x70, 0xf3, 0x7a, 0xf5, 0xa8, 0x7c, 0xa6, 0x39, 0xbd, 0xc2,
	0x37, 0xa4, 0x0e, 0x76, 0xe1, 0x81, 0x5b, 0xd1, 0xcd, 0xb2, 0xa4, 0xe9, 0xdd, 0x58, 0xfd, 0x1f,
	0xcf, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x38, 0x59, 0xf8, 0x30, 0x07, 0x00, 0x00,
}
