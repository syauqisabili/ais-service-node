package network

type DataBits uint8

const (
	DataBits5 DataBits = 5
	DataBits6 DataBits = 6
	DataBits7 DataBits = 7
	DataBits8 DataBits = 8
)

type Baudrate uint32

const (
	Baudrate2400   Baudrate = 2400
	Baudrate4800   Baudrate = 4800
	Baudrate9600   Baudrate = 9600
	Baudrate19200  Baudrate = 19200
	Baudrate1200   Baudrate = 1200
	Baudrate38400  Baudrate = 38500
	Baudrate57600  Baudrate = 57600
	Baudrate115200 Baudrate = 115200
)

type StopBits uint8

const (
	StopBits1 StopBits = 1
	StopBits2 StopBits = 2
)

type FlowControl uint8

const (
	NoneFlow FlowControl = 0
	RtsCts   FlowControl = 1
	DtrDsr   FlowControl = 2
	Rs485Rts FlowControl = 3
)

type Parity uint8

const (
	NoneParity  Parity = 0
	OddParity   Parity = 1
	EvenParity  Parity = 2
	MarkParity  Parity = 3
	SpaceParity Parity = 4
)
