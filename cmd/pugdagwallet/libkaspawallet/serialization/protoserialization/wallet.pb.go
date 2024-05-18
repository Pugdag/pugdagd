// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: wallet.proto

package protoserialization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PartiallySignedTransaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tx                    *TransactionMessage     `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
	PartiallySignedInputs []*PartiallySignedInput `protobuf:"bytes,2,rep,name=partiallySignedInputs,proto3" json:"partiallySignedInputs,omitempty"`
}

func (x *PartiallySignedTransaction) Reset() {
	*x = PartiallySignedTransaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartiallySignedTransaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartiallySignedTransaction) ProtoMessage() {}

func (x *PartiallySignedTransaction) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartiallySignedTransaction.ProtoReflect.Descriptor instead.
func (*PartiallySignedTransaction) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *PartiallySignedTransaction) GetTx() *TransactionMessage {
	if x != nil {
		return x.Tx
	}
	return nil
}

func (x *PartiallySignedTransaction) GetPartiallySignedInputs() []*PartiallySignedInput {
	if x != nil {
		return x.PartiallySignedInputs
	}
	return nil
}

type PartiallySignedInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RedeemScript         []byte                 `protobuf:"bytes,1,opt,name=redeemScript,proto3" json:"redeemScript,omitempty"`
	PrevOutput           *TransactionOutput     `protobuf:"bytes,2,opt,name=prevOutput,proto3" json:"prevOutput,omitempty"`
	MinimumSignatures    uint32                 `protobuf:"varint,3,opt,name=minimumSignatures,proto3" json:"minimumSignatures,omitempty"`
	PubKeySignaturePairs []*PubKeySignaturePair `protobuf:"bytes,4,rep,name=pubKeySignaturePairs,proto3" json:"pubKeySignaturePairs,omitempty"`
	DerivationPath       string                 `protobuf:"bytes,5,opt,name=derivationPath,proto3" json:"derivationPath,omitempty"`
}

func (x *PartiallySignedInput) Reset() {
	*x = PartiallySignedInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartiallySignedInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartiallySignedInput) ProtoMessage() {}

func (x *PartiallySignedInput) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartiallySignedInput.ProtoReflect.Descriptor instead.
func (*PartiallySignedInput) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *PartiallySignedInput) GetRedeemScript() []byte {
	if x != nil {
		return x.RedeemScript
	}
	return nil
}

func (x *PartiallySignedInput) GetPrevOutput() *TransactionOutput {
	if x != nil {
		return x.PrevOutput
	}
	return nil
}

func (x *PartiallySignedInput) GetMinimumSignatures() uint32 {
	if x != nil {
		return x.MinimumSignatures
	}
	return 0
}

func (x *PartiallySignedInput) GetPubKeySignaturePairs() []*PubKeySignaturePair {
	if x != nil {
		return x.PubKeySignaturePairs
	}
	return nil
}

func (x *PartiallySignedInput) GetDerivationPath() string {
	if x != nil {
		return x.DerivationPath
	}
	return ""
}

type PubKeySignaturePair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExtendedPubKey string `protobuf:"bytes,1,opt,name=extendedPubKey,proto3" json:"extendedPubKey,omitempty"`
	Signature      []byte `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *PubKeySignaturePair) Reset() {
	*x = PubKeySignaturePair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKeySignaturePair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKeySignaturePair) ProtoMessage() {}

func (x *PubKeySignaturePair) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PubKeySignaturePair.ProtoReflect.Descriptor instead.
func (*PubKeySignaturePair) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *PubKeySignaturePair) GetExtendedPubKey() string {
	if x != nil {
		return x.ExtendedPubKey
	}
	return ""
}

func (x *PubKeySignaturePair) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type SubnetworkId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *SubnetworkId) Reset() {
	*x = SubnetworkId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubnetworkId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubnetworkId) ProtoMessage() {}

func (x *SubnetworkId) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubnetworkId.ProtoReflect.Descriptor instead.
func (*SubnetworkId) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *SubnetworkId) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type TransactionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version      uint32               `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Inputs       []*TransactionInput  `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty"`
	Outputs      []*TransactionOutput `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty"`
	LockTime     uint64               `protobuf:"varint,4,opt,name=lockTime,proto3" json:"lockTime,omitempty"`
	SubnetworkId *SubnetworkId        `protobuf:"bytes,5,opt,name=subnetworkId,proto3" json:"subnetworkId,omitempty"`
	Gas          uint64               `protobuf:"varint,6,opt,name=gas,proto3" json:"gas,omitempty"`
	Payload      []byte               `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (x *TransactionMessage) Reset() {
	*x = TransactionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionMessage) ProtoMessage() {}

func (x *TransactionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionMessage.ProtoReflect.Descriptor instead.
func (*TransactionMessage) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *TransactionMessage) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *TransactionMessage) GetInputs() []*TransactionInput {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *TransactionMessage) GetOutputs() []*TransactionOutput {
	if x != nil {
		return x.Outputs
	}
	return nil
}

func (x *TransactionMessage) GetLockTime() uint64 {
	if x != nil {
		return x.LockTime
	}
	return 0
}

func (x *TransactionMessage) GetSubnetworkId() *SubnetworkId {
	if x != nil {
		return x.SubnetworkId
	}
	return nil
}

func (x *TransactionMessage) GetGas() uint64 {
	if x != nil {
		return x.Gas
	}
	return 0
}

func (x *TransactionMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type TransactionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PreviousOutpoint *Outpoint `protobuf:"bytes,1,opt,name=previousOutpoint,proto3" json:"previousOutpoint,omitempty"`
	SignatureScript  []byte    `protobuf:"bytes,2,opt,name=signatureScript,proto3" json:"signatureScript,omitempty"`
	Sequence         uint64    `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
	SigOpCount       uint32    `protobuf:"varint,4,opt,name=sigOpCount,proto3" json:"sigOpCount,omitempty"`
}

func (x *TransactionInput) Reset() {
	*x = TransactionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInput) ProtoMessage() {}

func (x *TransactionInput) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInput.ProtoReflect.Descriptor instead.
func (*TransactionInput) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *TransactionInput) GetPreviousOutpoint() *Outpoint {
	if x != nil {
		return x.PreviousOutpoint
	}
	return nil
}

func (x *TransactionInput) GetSignatureScript() []byte {
	if x != nil {
		return x.SignatureScript
	}
	return nil
}

func (x *TransactionInput) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *TransactionInput) GetSigOpCount() uint32 {
	if x != nil {
		return x.SigOpCount
	}
	return 0
}

type Outpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionId *TransactionId `protobuf:"bytes,1,opt,name=transactionId,proto3" json:"transactionId,omitempty"`
	Index         uint32         `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *Outpoint) Reset() {
	*x = Outpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outpoint) ProtoMessage() {}

func (x *Outpoint) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outpoint.ProtoReflect.Descriptor instead.
func (*Outpoint) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *Outpoint) GetTransactionId() *TransactionId {
	if x != nil {
		return x.TransactionId
	}
	return nil
}

func (x *Outpoint) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type TransactionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (x *TransactionId) Reset() {
	*x = TransactionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionId) ProtoMessage() {}

func (x *TransactionId) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionId.ProtoReflect.Descriptor instead.
func (*TransactionId) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{7}
}

func (x *TransactionId) GetBytes() []byte {
	if x != nil {
		return x.Bytes
	}
	return nil
}

type ScriptPublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Script  []byte `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	Version uint32 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ScriptPublicKey) Reset() {
	*x = ScriptPublicKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScriptPublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScriptPublicKey) ProtoMessage() {}

func (x *ScriptPublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScriptPublicKey.ProtoReflect.Descriptor instead.
func (*ScriptPublicKey) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{8}
}

func (x *ScriptPublicKey) GetScript() []byte {
	if x != nil {
		return x.Script
	}
	return nil
}

func (x *ScriptPublicKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type TransactionOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value           uint64           `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	ScriptPublicKey *ScriptPublicKey `protobuf:"bytes,2,opt,name=scriptPublicKey,proto3" json:"scriptPublicKey,omitempty"`
}

func (x *TransactionOutput) Reset() {
	*x = TransactionOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_wallet_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionOutput) ProtoMessage() {}

func (x *TransactionOutput) ProtoReflect() protoreflect.Message {
	mi := &file_wallet_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionOutput.ProtoReflect.Descriptor instead.
func (*TransactionOutput) Descriptor() ([]byte, []int) {
	return file_wallet_proto_rawDescGZIP(), []int{9}
}

func (x *TransactionOutput) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *TransactionOutput) GetScriptPublicKey() *ScriptPublicKey {
	if x != nil {
		return x.ScriptPublicKey
	}
	return nil
}

var File_wallet_proto protoreflect.FileDescriptor

var file_wallet_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xb4, 0x01, 0x0a, 0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x36, 0x0a, 0x02, 0x74, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x02, 0x74, 0x78, 0x12, 0x5e, 0x0a, 0x15, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x15, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67,
	0x6e, 0x65, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x22, 0xb4, 0x02, 0x0a, 0x14, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x6c, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d,
	0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x52, 0x0a, 0x70, 0x72, 0x65, 0x76, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x2c, 0x0a,
	0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75, 0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x6d, 0x69, 0x6e, 0x69, 0x6d, 0x75,
	0x6d, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x14, 0x70,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x75, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61,
	0x69, 0x72, 0x52, 0x14, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x50, 0x61, 0x69, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x64, 0x65, 0x72, 0x69, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x22, 0x5b, 0x0a, 0x13, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x50, 0x61, 0x69, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x24, 0x0a,
	0x0c, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x3f, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x44, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72,
	0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x67, 0x61, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x22, 0xc2, 0x01, 0x0a, 0x10, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x48, 0x0a, 0x10, 0x70, 0x72, 0x65, 0x76, 0x69, 0x6f,
	0x75, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x10,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x6f, 0x75, 0x73, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x28, 0x0a, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x69, 0x67, 0x4f, 0x70, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x73, 0x69, 0x67, 0x4f,
	0x70, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x69, 0x0a, 0x08, 0x4f, 0x75, 0x74, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x47, 0x0a, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x52, 0x0d, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x22, 0x25, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x62, 0x79, 0x74, 0x65, 0x73, 0x22, 0x43, 0x0a, 0x0f, 0x53, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x78, 0x0a,
	0x11, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x4d, 0x0a, 0x0f, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x0f, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x5c, 0x5a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x6e, 0x65, 0x74, 0x2f, 0x6b,
	0x61, 0x73, 0x70, 0x61, 0x64, 0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x6c, 0x69, 0x62, 0x6b, 0x61, 0x73, 0x70, 0x61, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_wallet_proto_rawDescOnce sync.Once
	file_wallet_proto_rawDescData = file_wallet_proto_rawDesc
)

func file_wallet_proto_rawDescGZIP() []byte {
	file_wallet_proto_rawDescOnce.Do(func() {
		file_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_wallet_proto_rawDescData)
	})
	return file_wallet_proto_rawDescData
}

var file_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_wallet_proto_goTypes = []interface{}{
	(*PartiallySignedTransaction)(nil), // 0: protoserialization.PartiallySignedTransaction
	(*PartiallySignedInput)(nil),       // 1: protoserialization.PartiallySignedInput
	(*PubKeySignaturePair)(nil),        // 2: protoserialization.PubKeySignaturePair
	(*SubnetworkId)(nil),               // 3: protoserialization.SubnetworkId
	(*TransactionMessage)(nil),         // 4: protoserialization.TransactionMessage
	(*TransactionInput)(nil),           // 5: protoserialization.TransactionInput
	(*Outpoint)(nil),                   // 6: protoserialization.Outpoint
	(*TransactionId)(nil),              // 7: protoserialization.TransactionId
	(*ScriptPublicKey)(nil),            // 8: protoserialization.ScriptPublicKey
	(*TransactionOutput)(nil),          // 9: protoserialization.TransactionOutput
}
var file_wallet_proto_depIdxs = []int32{
	4,  // 0: protoserialization.PartiallySignedTransaction.tx:type_name -> protoserialization.TransactionMessage
	1,  // 1: protoserialization.PartiallySignedTransaction.partiallySignedInputs:type_name -> protoserialization.PartiallySignedInput
	9,  // 2: protoserialization.PartiallySignedInput.prevOutput:type_name -> protoserialization.TransactionOutput
	2,  // 3: protoserialization.PartiallySignedInput.pubKeySignaturePairs:type_name -> protoserialization.PubKeySignaturePair
	5,  // 4: protoserialization.TransactionMessage.inputs:type_name -> protoserialization.TransactionInput
	9,  // 5: protoserialization.TransactionMessage.outputs:type_name -> protoserialization.TransactionOutput
	3,  // 6: protoserialization.TransactionMessage.subnetworkId:type_name -> protoserialization.SubnetworkId
	6,  // 7: protoserialization.TransactionInput.previousOutpoint:type_name -> protoserialization.Outpoint
	7,  // 8: protoserialization.Outpoint.transactionId:type_name -> protoserialization.TransactionId
	8,  // 9: protoserialization.TransactionOutput.scriptPublicKey:type_name -> protoserialization.ScriptPublicKey
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_wallet_proto_init() }
func file_wallet_proto_init() {
	if File_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartiallySignedTransaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PartiallySignedInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKeySignaturePair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubnetworkId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScriptPublicKey); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_wallet_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionOutput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wallet_proto_goTypes,
		DependencyIndexes: file_wallet_proto_depIdxs,
		MessageInfos:      file_wallet_proto_msgTypes,
	}.Build()
	File_wallet_proto = out.File
	file_wallet_proto_rawDesc = nil
	file_wallet_proto_goTypes = nil
	file_wallet_proto_depIdxs = nil
}
