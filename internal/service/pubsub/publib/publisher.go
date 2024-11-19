package publisher

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	pb "gitlab.com/elcarim-optronic-indonesia/ais-service-node/internal/service/pubsub/proto"
	"gitlab.com/elcarim-optronic-indonesia/ais-service-node/pkg"
)

func Publish(client pb.PublisherClient, key string, msg *pb.Message) {
	request := &pb.PublishRequest{Key: key, Messages: []*pb.Message{msg}}
	_, err := client.Publish(context.Background(), request)
	if err != nil {
		pkg.Log(log.ErrorLevel, fmt.Sprintf("Error publish: %s", err.Error()))
	}
}
