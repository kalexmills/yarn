// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yarn_spinner.proto

package bytecode

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

// The type of instruction that this is.
type Instruction_OpCode int32

const (
	// Jumps to a named position in the node.
	// opA = string: label name
	Instruction_JUMP_TO Instruction_OpCode = 0
	// Peeks a string from stack, and jumps to that named position in
	// the node.
	// No operands.
	Instruction_JUMP Instruction_OpCode = 1
	// Delivers a string ID to the client.
	// opA = string: string ID
	Instruction_RUN_LINE Instruction_OpCode = 2
	// Delivers a command to the client.
	// opA = string: command text
	Instruction_RUN_COMMAND Instruction_OpCode = 3
	// Adds an entry to the option list (see ShowOptions).
	// opA = string: string ID for option to add
	Instruction_ADD_OPTION Instruction_OpCode = 4
	// Presents the current list of options to the client, then clears
	// the list. The most recently selected option will be on the top
	// of the stack when execution resumes.
	// No operands.
	Instruction_SHOW_OPTIONS Instruction_OpCode = 5
	// Pushes a string onto the stack.
	// opA = string: the string to push to the stack.
	Instruction_PUSH_STRING Instruction_OpCode = 6
	// Pushes a floating point number onto the stack.
	// opA = float: number to push to stack
	Instruction_PUSH_FLOAT Instruction_OpCode = 7
	// Pushes a boolean onto the stack.
	// opA = bool: the bool to push to stack
	Instruction_PUSH_BOOL Instruction_OpCode = 8
	// Pushes a null value onto the stack.
	// No operands.
	Instruction_PUSH_NULL Instruction_OpCode = 9
	// Jumps to the named position in the the node, if the top of the
	// stack is not null, zero or false.
	// opA = string: label name
	Instruction_JUMP_IF_FALSE Instruction_OpCode = 10
	// Discards top of stack.
	// No operands.
	Instruction_POP Instruction_OpCode = 11
	// opA = string: name of the function
	Instruction_CALL_FUNC Instruction_OpCode = 12
	// Pushes the contents of a variable onto the stack.
	// opA = name of variable
	Instruction_PUSH_VARIABLE Instruction_OpCode = 13
	// Stores the contents of the top of the stack in the named
	// variable.
	// opA = name of variable
	Instruction_STORE_VARIABLE Instruction_OpCode = 14
	// Stops execution of the program.
	// No operands.
	Instruction_STOP Instruction_OpCode = 15
	// Run the node whose name is at the top of the stack.
	// No operands.
	Instruction_RUN_NODE Instruction_OpCode = 16
)

// Enum value maps for Instruction_OpCode.
var (
	Instruction_OpCode_name = map[int32]string{
		0:  "JUMP_TO",
		1:  "JUMP",
		2:  "RUN_LINE",
		3:  "RUN_COMMAND",
		4:  "ADD_OPTION",
		5:  "SHOW_OPTIONS",
		6:  "PUSH_STRING",
		7:  "PUSH_FLOAT",
		8:  "PUSH_BOOL",
		9:  "PUSH_NULL",
		10: "JUMP_IF_FALSE",
		11: "POP",
		12: "CALL_FUNC",
		13: "PUSH_VARIABLE",
		14: "STORE_VARIABLE",
		15: "STOP",
		16: "RUN_NODE",
	}
	Instruction_OpCode_value = map[string]int32{
		"JUMP_TO":        0,
		"JUMP":           1,
		"RUN_LINE":       2,
		"RUN_COMMAND":    3,
		"ADD_OPTION":     4,
		"SHOW_OPTIONS":   5,
		"PUSH_STRING":    6,
		"PUSH_FLOAT":     7,
		"PUSH_BOOL":      8,
		"PUSH_NULL":      9,
		"JUMP_IF_FALSE":  10,
		"POP":            11,
		"CALL_FUNC":      12,
		"PUSH_VARIABLE":  13,
		"STORE_VARIABLE": 14,
		"STOP":           15,
		"RUN_NODE":       16,
	}
)

func (x Instruction_OpCode) Enum() *Instruction_OpCode {
	p := new(Instruction_OpCode)
	*p = x
	return p
}

func (x Instruction_OpCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Instruction_OpCode) Descriptor() protoreflect.EnumDescriptor {
	return file_yarn_spinner_proto_enumTypes[0].Descriptor()
}

func (Instruction_OpCode) Type() protoreflect.EnumType {
	return &file_yarn_spinner_proto_enumTypes[0]
}

func (x Instruction_OpCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Instruction_OpCode.Descriptor instead.
func (Instruction_OpCode) EnumDescriptor() ([]byte, []int) {
	return file_yarn_spinner_proto_rawDescGZIP(), []int{2, 0}
}

// A complete Yarn program.
type Program struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the program.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The collection of nodes in this program.
	Nodes map[string]*Node `protobuf:"bytes,2,rep,name=nodes,proto3" json:"nodes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Program) Reset() {
	*x = Program{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_spinner_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Program) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Program) ProtoMessage() {}

func (x *Program) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_spinner_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Program.ProtoReflect.Descriptor instead.
func (*Program) Descriptor() ([]byte, []int) {
	return file_yarn_spinner_proto_rawDescGZIP(), []int{0}
}

func (x *Program) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Program) GetNodes() map[string]*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

// A collection of instructions
type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of this node.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The list of instructions in this node.
	Instructions []*Instruction `protobuf:"bytes,2,rep,name=instructions,proto3" json:"instructions,omitempty"`
	// A jump table, mapping the names of labels to positions in the
	// instructions list.
	Labels map[string]int32 `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// The tags associated with this node.
	Tags []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	// the entry in the program's string table that contains the original
	// text of this node; null if this is not available
	SourceTextStringID string `protobuf:"bytes,5,opt,name=sourceTextStringID,proto3" json:"sourceTextStringID,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_spinner_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_spinner_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_yarn_spinner_proto_rawDescGZIP(), []int{1}
}

func (x *Node) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Node) GetInstructions() []*Instruction {
	if x != nil {
		return x.Instructions
	}
	return nil
}

func (x *Node) GetLabels() map[string]int32 {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Node) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Node) GetSourceTextStringID() string {
	if x != nil {
		return x.SourceTextStringID
	}
	return ""
}

// A single Yarn instruction.
type Instruction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The operation that this instruction will perform.
	Opcode Instruction_OpCode `protobuf:"varint,1,opt,name=opcode,proto3,enum=Yarn.Instruction_OpCode" json:"opcode,omitempty"`
	// The list of operands, if any, that this instruction uses.
	Operands []*Operand `protobuf:"bytes,2,rep,name=operands,proto3" json:"operands,omitempty"`
}

func (x *Instruction) Reset() {
	*x = Instruction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_spinner_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Instruction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Instruction) ProtoMessage() {}

func (x *Instruction) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_spinner_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Instruction.ProtoReflect.Descriptor instead.
func (*Instruction) Descriptor() ([]byte, []int) {
	return file_yarn_spinner_proto_rawDescGZIP(), []int{2}
}

func (x *Instruction) GetOpcode() Instruction_OpCode {
	if x != nil {
		return x.Opcode
	}
	return Instruction_JUMP_TO
}

func (x *Instruction) GetOperands() []*Operand {
	if x != nil {
		return x.Operands
	}
	return nil
}

// A value used by an Instruction.
type Operand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of operand this is.
	//
	// Types that are assignable to Value:
	//	*Operand_StringValue
	//	*Operand_BoolValue
	//	*Operand_FloatValue
	Value isOperand_Value `protobuf_oneof:"value"`
}

func (x *Operand) Reset() {
	*x = Operand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yarn_spinner_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operand) ProtoMessage() {}

func (x *Operand) ProtoReflect() protoreflect.Message {
	mi := &file_yarn_spinner_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operand.ProtoReflect.Descriptor instead.
func (*Operand) Descriptor() ([]byte, []int) {
	return file_yarn_spinner_proto_rawDescGZIP(), []int{3}
}

func (m *Operand) GetValue() isOperand_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Operand) GetStringValue() string {
	if x, ok := x.GetValue().(*Operand_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *Operand) GetBoolValue() bool {
	if x, ok := x.GetValue().(*Operand_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (x *Operand) GetFloatValue() float32 {
	if x, ok := x.GetValue().(*Operand_FloatValue); ok {
		return x.FloatValue
	}
	return 0
}

type isOperand_Value interface {
	isOperand_Value()
}

type Operand_StringValue struct {
	// A string.
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Operand_BoolValue struct {
	// A boolean (true or false).
	BoolValue bool `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Operand_FloatValue struct {
	// A floating point number.
	FloatValue float32 `protobuf:"fixed32,3,opt,name=float_value,json=floatValue,proto3,oneof"`
}

func (*Operand_StringValue) isOperand_Value() {}

func (*Operand_BoolValue) isOperand_Value() {}

func (*Operand_FloatValue) isOperand_Value() {}

var File_yarn_spinner_proto protoreflect.FileDescriptor

var file_yarn_spinner_proto_rawDesc = []byte{
	0x0a, 0x12, 0x79, 0x61, 0x72, 0x6e, 0x5f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x59, 0x61, 0x72, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6e, 0x6f,
	0x64, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x59, 0x61, 0x72, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x1a, 0x44, 0x0a, 0x0a, 0x4e, 0x6f,
	0x64, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x59, 0x61, 0x72, 0x6e,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x80, 0x02, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x59, 0x61, 0x72, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x59, 0x61, 0x72, 0x6e, 0x2e, 0x4e, 0x6f, 0x64, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x65, 0x78, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x65, 0x78, 0x74,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xf6, 0x02, 0x0a, 0x0b, 0x49, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x06, 0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x59, 0x61, 0x72, 0x6e, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x6f,
	0x70, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x59, 0x61, 0x72, 0x6e, 0x2e, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x73,
	0x22, 0x89, 0x02, 0x0a, 0x06, 0x4f, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4a,
	0x55, 0x4d, 0x50, 0x5f, 0x54, 0x4f, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x55, 0x4d, 0x50,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x55, 0x4e, 0x5f, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x02,
	0x12, 0x0f, 0x0a, 0x0b, 0x52, 0x55, 0x4e, 0x5f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x10,
	0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x44, 0x44, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x04, 0x12, 0x10, 0x0a, 0x0c, 0x53, 0x48, 0x4f, 0x57, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e,
	0x53, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x53, 0x54, 0x52, 0x49,
	0x4e, 0x47, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x46, 0x4c, 0x4f,
	0x41, 0x54, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x42, 0x4f, 0x4f,
	0x4c, 0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x4e, 0x55, 0x4c, 0x4c,
	0x10, 0x09, 0x12, 0x11, 0x0a, 0x0d, 0x4a, 0x55, 0x4d, 0x50, 0x5f, 0x49, 0x46, 0x5f, 0x46, 0x41,
	0x4c, 0x53, 0x45, 0x10, 0x0a, 0x12, 0x07, 0x0a, 0x03, 0x50, 0x4f, 0x50, 0x10, 0x0b, 0x12, 0x0d,
	0x0a, 0x09, 0x43, 0x41, 0x4c, 0x4c, 0x5f, 0x46, 0x55, 0x4e, 0x43, 0x10, 0x0c, 0x12, 0x11, 0x0a,
	0x0d, 0x50, 0x55, 0x53, 0x48, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x0d,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x5f, 0x56, 0x41, 0x52, 0x49, 0x41, 0x42,
	0x4c, 0x45, 0x10, 0x0e, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x0f, 0x12, 0x0c,
	0x0a, 0x08, 0x52, 0x55, 0x4e, 0x5f, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x10, 0x22, 0x7b, 0x0a, 0x07,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0a,
	0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a,
	0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x72, 0x4a, 0x6f, 0x73, 0x68, 0x39, 0x30,
	0x30, 0x30, 0x2f, 0x79, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x79, 0x74, 0x65, 0x63, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yarn_spinner_proto_rawDescOnce sync.Once
	file_yarn_spinner_proto_rawDescData = file_yarn_spinner_proto_rawDesc
)

func file_yarn_spinner_proto_rawDescGZIP() []byte {
	file_yarn_spinner_proto_rawDescOnce.Do(func() {
		file_yarn_spinner_proto_rawDescData = protoimpl.X.CompressGZIP(file_yarn_spinner_proto_rawDescData)
	})
	return file_yarn_spinner_proto_rawDescData
}

var file_yarn_spinner_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yarn_spinner_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yarn_spinner_proto_goTypes = []interface{}{
	(Instruction_OpCode)(0), // 0: Yarn.Instruction.OpCode
	(*Program)(nil),         // 1: Yarn.Program
	(*Node)(nil),            // 2: Yarn.Node
	(*Instruction)(nil),     // 3: Yarn.Instruction
	(*Operand)(nil),         // 4: Yarn.Operand
	nil,                     // 5: Yarn.Program.NodesEntry
	nil,                     // 6: Yarn.Node.LabelsEntry
}
var file_yarn_spinner_proto_depIdxs = []int32{
	5, // 0: Yarn.Program.nodes:type_name -> Yarn.Program.NodesEntry
	3, // 1: Yarn.Node.instructions:type_name -> Yarn.Instruction
	6, // 2: Yarn.Node.labels:type_name -> Yarn.Node.LabelsEntry
	0, // 3: Yarn.Instruction.opcode:type_name -> Yarn.Instruction.OpCode
	4, // 4: Yarn.Instruction.operands:type_name -> Yarn.Operand
	2, // 5: Yarn.Program.NodesEntry.value:type_name -> Yarn.Node
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yarn_spinner_proto_init() }
func file_yarn_spinner_proto_init() {
	if File_yarn_spinner_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yarn_spinner_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Program); i {
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
		file_yarn_spinner_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_yarn_spinner_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Instruction); i {
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
		file_yarn_spinner_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operand); i {
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
	file_yarn_spinner_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Operand_StringValue)(nil),
		(*Operand_BoolValue)(nil),
		(*Operand_FloatValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yarn_spinner_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yarn_spinner_proto_goTypes,
		DependencyIndexes: file_yarn_spinner_proto_depIdxs,
		EnumInfos:         file_yarn_spinner_proto_enumTypes,
		MessageInfos:      file_yarn_spinner_proto_msgTypes,
	}.Build()
	File_yarn_spinner_proto = out.File
	file_yarn_spinner_proto_rawDesc = nil
	file_yarn_spinner_proto_goTypes = nil
	file_yarn_spinner_proto_depIdxs = nil
}