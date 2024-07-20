package main

import (
	"log"
	"net"

	"github.com/hebobibun/grpc-auth-crud/pkg/config"
	"github.com/hebobibun/grpc-auth-crud/pkg/db"
	pb "github.com/hebobibun/grpc-auth-crud/pkg/pb"
	"github.com/hebobibun/grpc-auth-crud/pkg/services"
	"github.com/hebobibun/grpc-auth-crud/pkg/utils"
	"google.golang.org/grpc"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	dbHandler := db.InitDB(config.DBUrl)
	log.Println("DB connected", dbHandler.DB)

	jwt := utils.JWTwrapper{
		SecretKey:  config.JWTSecret,
		Issuer:     "grpc-auth-crud",
		Expiration: 24 * 7,
	}
	log.Println("JWT initialized", jwt)

	lis, err := net.Listen("tcp", ":"+config.Port)
	if err != nil {
		log.Fatalln("Failed to listen", err)
	}

	log.Println("Server started on port", config.Port)

	s := services.AuthServer{
		H:   dbHandler,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServer(grpcServer, &s)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln("Failed to start server", err)
	}
}
