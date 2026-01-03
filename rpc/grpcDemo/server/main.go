package main

import (
	context2 "context"
	"fmt"
	"log"
	"net"

	// 导入grpc包
	"google.golang.org/grpc"
	// 导入刚才我们生成的代码所在的proto包。
	pb "GoLearn/rpc/grpcDemo/server/common"
	"google.golang.org/grpc/reflection"
)

// 定义server，用来实现proto文件，里面实现的Greeter服务里面的接口
type server struct{}

func (s *server) Work(ctx context2.Context, request *pb.WorkRequest) (*pb.WorkResponse, error) {

	fmt.Println(request.Tools)
	return &pb.WorkResponse{Money: 4000}, nil

}

func (s *server) Study(ctx context2.Context, request *pb.StudyRequest) (*pb.StudyResponse, error) {

	return &pb.StudyResponse{
		Money:       2555,
		DoSomething: "学会了。。。",
	}, nil
}

func main() {

	// 监听127.0.0.1:50051地址
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc服务端
	s := grpc.NewServer()

	// 注册Greeter服务
	pb.RegisterLifeServer(s, &server{})

	// 往grpc服务端注册反射服务
	reflection.Register(s)

	// 启动grpc服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
