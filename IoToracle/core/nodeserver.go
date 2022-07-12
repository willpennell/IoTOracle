package core

import (
	"IoToracle/probuf"
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
	probuf.UnimplementedOracleNodeServiceServer
}

func (s *OracleNodeServer) SendRequest(ctx context.Context, in *probuf.RequestToNode) (*probuf.ResponseFromNode, error) {
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
}

func ServerSetup(wg *sync.WaitGroup) {
	defer wg.Done()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	probuf.RegisterOracleNodeServiceServer(s, &OracleNodeServer{})
	u.SERVERLISTENING(lis)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println(s.GetServiceInfo())

}
