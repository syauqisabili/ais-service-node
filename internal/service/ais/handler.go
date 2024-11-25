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

func PubSubMessageHandler(target domain.AisTarget) error {
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

	// Target
	var aisTarget domain.AisTarget
	aisTarget.MessageID = msgId
	aisTarget.Target.Timestamp = time.Now().Unix()
	aisTarget.Target.Mmsi = v.GetHeader().UserID

	switch msgId {
	case 1, 2, 3, 5, 27: // Class A
		aisTarget.Target.Class = domain.ClassA
		if _, ok := v.(ais.PositionReport); ok {
			aisTarget.Target.NavigationalStatus = v.(ais.PositionReport).NavigationalStatus
			aisTarget.Target.Latitude = ais.FieldLatLonCoarse(v.(ais.PositionReport).Latitude)
			aisTarget.Target.Longitude = ais.FieldLatLonCoarse(v.(ais.PositionReport).Longitude)
			aisTarget.Target.TrueHeading = v.(ais.PositionReport).TrueHeading
			aisTarget.Target.Cog = ais.Field10(v.(ais.PositionReport).Cog)
			aisTarget.Target.Sog = ais.Field10(v.(ais.PositionReport).Sog)

		} else if _, ok := v.(ais.ShipStaticData); ok {
			aisTarget.Target.ImoNumber = v.(ais.ShipStaticData).ImoNumber
			aisTarget.Target.Name = v.(ais.ShipStaticData).Name
			aisTarget.Target.CallSign = v.(ais.ShipStaticData).CallSign
			aisTarget.Target.ShipType = v.(ais.ShipStaticData).Type
			aisTarget.Target.Destination = v.(ais.ShipStaticData).Destination
			aisTarget.Target.Dimension = v.(ais.ShipStaticData).Dimension
		} else if _, ok := v.(ais.LongRangeAisBroadcastMessage); ok {
			aisTarget.Target.NavigationalStatus = v.(ais.LongRangeAisBroadcastMessage).NavigationalStatus
			aisTarget.Target.Latitude = ais.FieldLatLonCoarse(v.(ais.LongRangeAisBroadcastMessage).Latitude)
			aisTarget.Target.Longitude = ais.FieldLatLonCoarse(v.(ais.LongRangeAisBroadcastMessage).Longitude)
			aisTarget.Target.Cog = ais.Field10(v.(ais.LongRangeAisBroadcastMessage).Cog)
			aisTarget.Target.Sog = ais.Field10(v.(ais.LongRangeAisBroadcastMessage).Sog)
		}

	case 18, 19, 24: // Class B
		//! ImoNumber does not exist
		//! NavStatus does not exist
		//! Destination does not exist

		aisTarget.Target.Class = domain.ClassB
		if _, ok := v.(ais.StaticDataReport); ok {
			aisTarget.Target.Name = v.(ais.StaticDataReport).ReportA.Name
			aisTarget.Target.CallSign = v.(ais.StaticDataReport).ReportB.CallSign
			aisTarget.Target.ShipType = v.(ais.StaticDataReport).ReportB.ShipType
			aisTarget.Target.Dimension = ais.FieldDimension(v.(ais.StaticDataReport).ReportB.Dimension)

		} else if _, ok := v.(ais.StandardClassBPositionReport); ok {
			aisTarget.Target.Latitude = ais.FieldLatLonCoarse(v.(ais.StandardClassBPositionReport).Latitude)
			aisTarget.Target.Longitude = ais.FieldLatLonCoarse(v.(ais.StandardClassBPositionReport).Longitude)
			aisTarget.Target.TrueHeading = v.(ais.StandardClassBPositionReport).TrueHeading
			aisTarget.Target.Cog = ais.Field10(v.(ais.StandardClassBPositionReport).Cog)
			aisTarget.Target.Sog = ais.Field10(v.(ais.StandardClassBPositionReport).Sog)

		} else if _, ok := v.(ais.ExtendedClassBPositionReport); ok {
			aisTarget.Target.Name = v.(ais.ExtendedClassBPositionReport).Name
			aisTarget.Target.ShipType = v.(ais.ExtendedClassBPositionReport).Type
			aisTarget.Target.Latitude = ais.FieldLatLonCoarse(v.(ais.ExtendedClassBPositionReport).Latitude)
			aisTarget.Target.Longitude = ais.FieldLatLonCoarse(v.(ais.ExtendedClassBPositionReport).Longitude)
			aisTarget.Target.TrueHeading = v.(ais.ExtendedClassBPositionReport).TrueHeading
			aisTarget.Target.Cog = ais.Field10(v.(ais.ExtendedClassBPositionReport).Cog)
			aisTarget.Target.Sog = ais.Field10(v.(ais.ExtendedClassBPositionReport).Sog)
			aisTarget.Target.Dimension = ais.FieldDimension(v.(ais.ExtendedClassBPositionReport).Dimension)
		}
	case 4, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 20, 21, 22, 23, 25, 26:
		//TODO: status = NotSupportedAisMSg
		return pkg.ErrorStatus(pkg.ErrProcessFail, "Not supported ais message")

	default:
		//TODO: status = InvalidAisMsg
		return pkg.ErrorStatus(pkg.ErrProcessFail, "Invalid ais message")
	}

	// Publish target to broker
	pkg.Log(log.InfoLevel, fmt.Sprintf("Target: %v", aisTarget))
	err := PubSubMessageHandler(aisTarget)
	if err != nil {
		return err
	}

	return nil
}
