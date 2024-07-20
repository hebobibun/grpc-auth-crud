package services

import (
	"github.com/hebobibun/grpc-auth-crud/pkg/db"
	"github.com/hebobibun/grpc-auth-crud/pkg/utils"
)

type Server struct {
	H   db.Handler
	Jwt utils.JWTwrapper
}
