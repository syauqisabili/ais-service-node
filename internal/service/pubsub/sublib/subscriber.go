package subscriber

import (
	"context"
	"io"
	"sync"

	pb "gitlab.com/elcarim-optronic-indonesia/ais-service-node/internal/service/pubsub/proto"
)

type SubscriberContext struct {
	Client   pb.SubscriberClient
	identity *pb.Identity
	Speed    int
	Size     int
	Mutex    sync.Mutex
}

func (s *SubscriberContext) Authenticate(name string) error {
	if identity, error := s.Client.Authenticate(context.Background(), &pb.Identity{Name: name}); error == nil {
		s.identity = identity
		return nil
	}
	return nil
}

func (s *SubscriberContext) Subscribe(key string) error {
	request := &pb.SubscribeRequest{Identity: s.identity, Subscription: &pb.Subscription{Key: key}}
	if _, error := s.Client.Subscribe(context.Background(), request); error == nil {
		return nil
	}
	return nil
}

func (s *SubscriberContext) Pull() error {
	if stream, error := s.Client.Pull(context.Background(), s.identity); error == nil {

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
			}
			s.Mutex.Lock()
			s.Speed++
			s.Size += len(msg.Data)
			s.Mutex.Unlock()
		}

		return nil
	}
	return nil
}

func NewSubscriberContext() *SubscriberContext {
	s := new(SubscriberContext)
	s.Speed = 0
	return s
}
