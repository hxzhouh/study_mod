package server

import (
	pb "github.com/hxzhouh/study_mod/grpc/protobuf"
	"io"
	"log"
	"strconv"
)

//双向流

// StreamService 定义我们的服务
type StreamService struct{}

func (s *StreamService)Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}