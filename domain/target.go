package domain

import "github.com/BertoldVdb/go-ais"

const MaxActiveTarget = 500

type ActiveTargetSet bool

const (
	Increment ActiveTargetSet = true
	Decrement ActiveTargetSet = false
)

type Class string

const (
	ClassA Class = "A"
	ClassB Class = "B"
	ClassC Class = "C" // Aton
)

type AisTarget struct {
	MessageID uint8  `json:"message_id"`
	Target    Target `json:"target"`
}

type Target struct {
	Timestamp          int64                 `json:"timestamp"`
	Status             string                `json:"status"`
	Mmsi               uint32                `json:"mmsi"`
	ImoNumber          uint32                `json:"imo"`
	Name               string                `json:"name"`
	CallSign           string                `json:"call_sign"`
	Class              Class                 `json:"class"`
	ShipType           uint8                 `json:"ship_type"`
	NavigationalStatus uint8                 `json:"nav_status"`
	Latitude           ais.FieldLatLonCoarse `json:"latitude"`
	Longitude          ais.FieldLatLonCoarse `json:"longitude"`
	Cog                ais.Field10           `json:"cog"`
	Sog                ais.Field10           `json:"sog"`
	Dimension          ais.FieldDimension    `json:"dimension"`
	TrueHeading        uint16                `json:"true_heading"`
	Destination        string                `json:"destination"`
}
