package target

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/BertoldVdb/go-ais"
	"github.com/charmbracelet/log"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/config/network"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/domain"
	pb "gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/pubsub/proto"
	pub "gitlab.com/elcarim-optronic-indonesia/elcas-service-node/internal/service/pubsub/pub"
	"gitlab.com/elcarim-optronic-indonesia/elcas-service-node/pkg"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func PubSubMessageHandler(target domain.Target) error {
	// Get config
	cnf := network.Get()
	opts := []grpc.DialOption{}

	//TODO: MUST SET CREDS/TLS
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// Pubsub grpc broker service address
	addr := cnf.GrpcServer.Ip + ":" + strconv.FormatUint(uint64(cnf.GrpcServer.Port), 10)
	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return pkg.ErrorStatus(pkg.ErrProcessFail, "Cannot connect to broker")
	}
	defer conn.Close()

	c := pb.NewPublisherClient(conn)

	// From object encode to json
	message, err := json.Marshal(target)
	if err != nil {
		return err
	}

	//TODO: Publish message
	topic := "/target"
	pub.Publish(c, topic, &pb.Message{
		Data: message,
	})

	return nil
}

func Handler(packet ais.Packet) error {

	v := packet
	msgId := v.GetHeader().MessageID
	pkg.Log(log.InfoLevel, fmt.Sprintf("Message id: %d", msgId))

	// Target
	var target domain.Target
	target.Timestamp = time.Now().Unix()
	target.Mmsi = v.GetHeader().UserID

	switch msgId {
	case 1, 2, 3, 5, 27: // Class A
		target.Class = domain.ClassA
		if _, ok := v.(ais.PositionReport); ok {
			target.NavigationalStatus = v.(ais.PositionReport).NavigationalStatus
			target.Latitude = ais.FieldLatLonCoarse(v.(ais.PositionReport).Latitude)
			target.Longitude = ais.FieldLatLonCoarse(v.(ais.PositionReport).Longitude)
			target.TrueHeading = v.(ais.PositionReport).TrueHeading
			target.Cog = ais.Field10(v.(ais.PositionReport).Cog)
			target.Sog = ais.Field10(v.(ais.PositionReport).Sog)

		} else if _, ok := v.(ais.ShipStaticData); ok {
			target.ImoNumber = v.(ais.ShipStaticData).ImoNumber
			target.Name = v.(ais.ShipStaticData).Name
			target.CallSign = v.(ais.ShipStaticData).CallSign
			target.ShipType = v.(ais.ShipStaticData).Type
			target.Destination = v.(ais.ShipStaticData).Destination
			target.Dimension = v.(ais.ShipStaticData).Dimension
		} else if _, ok := v.(ais.LongRangeAisBroadcastMessage); ok {
			target.NavigationalStatus = v.(ais.LongRangeAisBroadcastMessage).NavigationalStatus
			target.Latitude = ais.FieldLatLonCoarse(v.(ais.LongRangeAisBroadcastMessage).Latitude)
			target.Longitude = ais.FieldLatLonCoarse(v.(ais.LongRangeAisBroadcastMessage).Longitude)
			target.Cog = ais.Field10(v.(ais.LongRangeAisBroadcastMessage).Cog)
			target.Sog = ais.Field10(v.(ais.LongRangeAisBroadcastMessage).Sog)
		}

	case 18, 19, 24: // Class B
		//! ImoNumber does not exist
		//! NavStatus does not exist
		//! Destination does not exist

		target.Class = domain.ClassB
		if _, ok := v.(ais.StaticDataReport); ok {
			target.Name = v.(ais.StaticDataReport).ReportA.Name
			target.CallSign = v.(ais.StaticDataReport).ReportB.CallSign
			target.ShipType = v.(ais.StaticDataReport).ReportB.ShipType
			target.Dimension = ais.FieldDimension(v.(ais.StaticDataReport).ReportB.Dimension)

		} else if _, ok := v.(ais.StandardClassBPositionReport); ok {
			target.Latitude = ais.FieldLatLonCoarse(v.(ais.StandardClassBPositionReport).Latitude)
			target.Longitude = ais.FieldLatLonCoarse(v.(ais.StandardClassBPositionReport).Longitude)
			target.TrueHeading = v.(ais.StandardClassBPositionReport).TrueHeading
			target.Cog = ais.Field10(v.(ais.StandardClassBPositionReport).Cog)
			target.Sog = ais.Field10(v.(ais.StandardClassBPositionReport).Sog)

		} else if _, ok := v.(ais.ExtendedClassBPositionReport); ok {
			target.Name = v.(ais.ExtendedClassBPositionReport).Name
			target.ShipType = v.(ais.ExtendedClassBPositionReport).Type
			target.Latitude = ais.FieldLatLonCoarse(v.(ais.ExtendedClassBPositionReport).Latitude)
			target.Longitude = ais.FieldLatLonCoarse(v.(ais.ExtendedClassBPositionReport).Longitude)
			target.TrueHeading = v.(ais.ExtendedClassBPositionReport).TrueHeading
			target.Cog = ais.Field10(v.(ais.ExtendedClassBPositionReport).Cog)
			target.Sog = ais.Field10(v.(ais.ExtendedClassBPositionReport).Sog)
			target.Dimension = ais.FieldDimension(v.(ais.ExtendedClassBPositionReport).Dimension)
		}
	case 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 20, 21, 22, 23, 25, 26:
		//TODO: status = NotSupportedAisMSg
		return pkg.ErrorStatus(pkg.ErrProcessFail, "Not supported ais message")

	default:
		//TODO: status = InvalidAisMsg
		return pkg.ErrorStatus(pkg.ErrProcessFail, "Invalid ais message")
	}

	// Publish target to broker
	err := PubSubMessageHandler(target)
	if err != nil {
		return err
	}

	return nil
}
