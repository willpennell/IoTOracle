package core

import (
	pb "IoToracle/probuf"
	u "IoToracle/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

const (
	port = ":50051"
)

type OracleNodeServer struct {
	pb.UnimplementedOracleNodeServiceServer
}

func (s *OracleNodeServer) CheckForRequest(ctx context.Context, in *pb.AckClient) (*pb.AckServer, error) {
	log.Printf("Recieved client Ack: %v", in.GetAck())
	if in.GetAck() == true {
		return &pb.AckServer{
			Ack:       true,
			RequestId: 0,
		}, nil
	} else {
		return &pb.AckServer{
			Ack:       false,
			RequestId: 0,
		}, nil
	}
}

/*func (s *OracleNodeServer) SendRequest(ctx context.Context, in *pb.RequestToNode) (*pb.ResponseFromNode, error) {
	log.Printf("Received: %v", in.GetRequestId())
	_ = ctx
	// logic to now wait for response from IoT device
	fmt.Printf("outward response sent...: %v\n", in.GetIotId())
	return &probuf.ResponseFromNode{
		IotId:          in.GetIotId(),
		RequestId:      in.GetRequestId(),
		RequiredResult: in.GetRequireResult(),
		ActualResult:   false,
	}, nil
}*/

func ServerSetup(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterOracleNodeServiceServer(s, &OracleNodeServer{})
	u.SERVERLISTENING(lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println(s.GetServiceInfo())

}
