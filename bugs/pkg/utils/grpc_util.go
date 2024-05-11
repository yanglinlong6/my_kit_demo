package utils

//package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	npb "notificator/pkg/grpc/pb"
	"users/pkg/grpc/pb"
)

func CallOtherGRPCService() pb.UsersClient {
	// 创建 gRPC 连接
	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure())
	if err != nil {
		// 错误处理
		return nil
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := pb.NewUsersClient(conn)

	hello, err := client.Hello(context.Background(), &pb.HelloRequest{})
	fmt.Println("hello===" + hello.String())
	return client
}

func CallOtherGRPCService01() npb.NotificatorClient {
	// 创建 gRPC 连接
	conn, err := grpc.Dial("127.0.0.1:8087", grpc.WithInsecure())
	if err != nil {
		// 错误处理
		return nil
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := npb.NewNotificatorClient(conn)

	//hello, err := client.Hello(context.Background(), &pb.HelloRequest{})
	//fmt.Println("hello===" + hello.String())
	return client
}

//func main() {
//	service := CallOtherGRPCService()
//	hello, err := service.Hello(context.Background(), &pb.HelloRequest{})
//	if err != nil {
//		return
//	}
//	fmt.Println(hello.String())
//}
