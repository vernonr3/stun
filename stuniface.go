package stun

import (
	"io"
)

type stunvarsIF interface {
	GetTransactionID() [TransactionIDSize]byte
	GetMessage() *Message
	SetMessage(m []byte) bool
}

type accessMessageType interface {
	GetTypeMethod() Method
	GetTypeClass() MessageClass
}

// StunMessageIF is a stupidly big interface and ought to be decomposed...
type StunMessageIF interface {
	stunvarsIF
	accessMessageType

	ApplyBuf(buff ...byte)
	//	setInternalMessage(m *Message)
	Build(setters ...Setter) error
	Check(checkers ...Checker)
	Parse(getters ...Getter) error
	// the next function is the exception as it takes a pointer to a message..
	ForEach(t AttrType, f func(m *Message) error) error
	//ForEach(t AttrType, f func(m StunMessageIF) error) error

	CloneTo(b *StunMessageIF) error
	Contains(t AttrType) bool

	UnmarshalBinary(data []byte) error
	GobEncode() ([]byte, error)
	GobDecode(data []byte) error
	AddTo(b *StunMessageIF) error
	Equal(b *StunMessageIF) bool
	NewTransactionID() error
	String() string
	Reset()
	Grow(n int)
	Add(t AttrType, v []byte)

	WriteLength()
	WriteHeader()
	WriteTransactionID()
	WriteAttributes()
	WriteType()
	SetType(t MessageType)
	Encode()
	Decode() error
	WriteTo(w io.Writer) (int64, error)
	ReadFrom(r io.Reader) (int64, error)
	Write(tBuf []byte) (int, error)
}

type StunMessage struct {
	stMessage Message
	stIF      StunMessageIF
}

func NewStunMessage() *StunMessage {
	return &StunMessage{}
}

func (m *StunMessage) ApplyBuf(buff ...byte) {
	m.stMessage = Message{Raw: append([]byte{}, buff...)}
}

func (m *StunMessage) SetMessage(msg []byte) bool {
	// ought to arrange to pick other bits out of the raw buffer and fill in the fields in the StunMessage struct..
	m.stMessage.Raw = msg
	return true
}

func (m *StunMessage) GetTypeMethod() Method {
	return m.stMessage.Type.Method
}
func (m *StunMessage) GetTypeClass() MessageClass {
	return m.stMessage.Type.Class
}

func (m *StunMessage) Decode() error {
	return m.stMessage.Decode()
}

func (m *StunMessage) Encode() {
	m.stMessage.Encode()
}

func (m *StunMessage) Add(t AttrType, v []byte) {
	m.stMessage.Add(t, v)
}

func (m *StunMessage) AddTo(b *StunMessageIF) error {
	return m.stMessage.AddTo((*b).GetMessage())
}

func (m *StunMessage) Equal(b *StunMessageIF) bool {
	return m.stMessage.Equal((*b).GetMessage())
}

func (m *StunMessage) ForEach(t AttrType, f func(m *Message) error) error {
	return m.stMessage.ForEach(t, f)
}

func (m *StunMessage) Build(setters ...Setter) error {
	return m.stMessage.Build(setters...)
}

func (m *StunMessage) Check(checkers ...Checker) {
	m.stMessage.Check(checkers...)
}

func (m *StunMessage) Parse(getters ...Getter) error {
	return m.stMessage.Parse(getters...)
}

func (m *StunMessage) CloneTo(b *StunMessageIF) error {
	return m.stMessage.CloneTo((*b).GetMessage())
}

func (m *StunMessage) Contains(t AttrType) bool {
	return m.stMessage.Contains(t)
}

func (m *StunMessage) WriteTo(w io.Writer) (int64, error) {
	return m.stMessage.WriteTo(w)
}
func (m *StunMessage) ReadFrom(r io.Reader) (int64, error) {
	return m.stMessage.ReadFrom(r)
}
func (m *StunMessage) Write(tBuf []byte) (int, error) {
	return m.stMessage.Write(tBuf)
}

func (m *StunMessage) GetMessage() *Message {
	return &m.stMessage
}

func (m *StunMessage) GetTransactionID() [TransactionIDSize]byte {
	return m.stMessage.TransactionID
}

func (m *StunMessage) UnmarshalBinary(data []byte) error {
	return m.stMessage.UnmarshalBinary((data))
}
func (m *StunMessage) GobEncode() ([]byte, error) {
	return m.stMessage.GobEncode()
}
func (m *StunMessage) GobDecode(data []byte) error {
	return m.stMessage.GobDecode(data)
}

func (m *StunMessage) String() string {
	return m.stMessage.String()
}
func (m *StunMessage) Reset() {
	m.stMessage.Reset()
}
func (m *StunMessage) Grow(n int) {
	m.stMessage.grow(n)
}

func (m *StunMessage) NewTransactionID() error {
	return m.stMessage.NewTransactionID()
}

func (m *StunMessage) WriteLength() {
	m.stMessage.WriteLength()
}

func (m *StunMessage) WriteHeader() {
	m.stMessage.WriteHeader()
}
func (m *StunMessage) WriteTransactionID() {
	m.stMessage.WriteTransactionID()
}
func (m *StunMessage) WriteAttributes() {
	m.stMessage.WriteAttributes()
}
func (m *StunMessage) WriteType() {
	m.stMessage.WriteType()
}
func (m *StunMessage) SetType(t MessageType) {
	m.stMessage.SetType(t)
}
