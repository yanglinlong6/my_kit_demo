package main

import service "users/cmd/service"

func main() {
	// kit new service users
	// kit generate service users --dmw
	// kit generate service notificator -t grpc --dmw
	// grpcui -plaintext 127.0.0.1:8081
	service.Run()
}
