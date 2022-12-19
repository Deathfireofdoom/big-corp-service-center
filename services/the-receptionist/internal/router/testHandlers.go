package router

import (
	"context"
	"log"
	"time"

	pb "github.com/Deathfireofdoom/big-corp-service-center/grpc/themanager"
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

func theManagerTest(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewTheManagerClient(conn)

	ctxTemp, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var new_messages = make(map[string]int32)
	new_messages["Test1"] = 43
	new_messages["Test2"] = 42
	for message, number := range new_messages {
		r, err := c.TestFunction(ctxTemp, &pb.TestMessage{Message: message, Number: number})
		if err != nil {
			log.Println("could not test grpc: %v", err)
		}
		log.Printf(`
		 Details:
		 Message: %s,
		 Number: %d,
		 Random: %s
		`, r.GetMessage(), r.GetNumber(), r.GetRandom())
	}

}
