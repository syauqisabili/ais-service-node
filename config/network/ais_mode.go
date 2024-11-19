package network

type SourceGps int

const (
	SourceUdpMulticast SourceGps = 0
	SourceSerial       SourceGps = 1
)

type DebugMode int

const (
	DisableDebug DebugMode = 0
	DebugAll     DebugMode = 1
)
