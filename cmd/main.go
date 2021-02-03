package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/titohazin/imersao-fullstack-fullcycle/application/grpc"
	"github.com/titohazin/imersao-fullstack-fullcycle/infra/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
