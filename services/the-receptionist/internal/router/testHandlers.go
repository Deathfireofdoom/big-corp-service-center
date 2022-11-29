package router

import (
	"context"
	"log"
	"time"

	pb "github.com/Deathfireofdoom/big-corp-service-center/the-manager/protobuf/themanager"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func connectionTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func theManageTest(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTheManagerClient(conn)

}
