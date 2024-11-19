package network

const MaxIpLen = 16

type NetConn struct {
	Ip   string `json:"ip"`   //  IP Address
	Port uint16 `json:"port"` // Port
}

type GrpcServer struct {
	Ip   string `json:"ip"`   //  IP Address
	Port uint16 `json:"port"` // Port
}

type Serial struct {
	ComPort     string      `json:"serial_port"` // COM0-COM999
	Baudrate    Baudrate    `json:"baudrate"`
	Parity      Parity      `json:"parity"`
	DataBits    DataBits    `json:"databits"`
	StopBits    StopBits    `json:"stopbits"`
	FlowControl FlowControl `json:"flow_control"`
}

type MysqlServer struct {
	Ip           string `json:"ip"`            // IP
	Port         uint16 `json:"port"`          // Port
	Username     string `json:"username"`      // Username MySQL
	Password     string `json:"password"`      // Password  MySQL
	DatabaseName string `json:"database_name"` // Db name
}

type MongoDbServer struct {
	Ip           string `json:"ip"`   // IP
	Port         uint16 `json:"port"` // Port
	DatabaseName string `json:"database_name"`
}

type Redis struct {
	Ip            string `json:"ip"`   // IP
	Port          uint16 `json:"port"` // Port
	Password      string `json:"password"`
	DatabaseIndex uint8  `json:"database_index"`
}

type NetworkConfig struct {
	DebugMode  DebugMode  `json:"debug_mode"`
	SourceAis  SourceGps  `json:"source_ais"`
	UdpNet     NetConn    `json:"udp_net"`
	GrpcServer GrpcServer `json:"grpc_server"`
	Serial     Serial     `json:"serial"`
}
