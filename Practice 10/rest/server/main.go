package main

import (
	pb "de_pract/grpc/protobuf"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"

	//	"fmt"
	"github.com/gin-gonic/gin"
	//	"github.com/joho/godotenv"
	"log"

	//	"net"
	"net/http"
	//	"os"
	"sync"
)

var lock = &sync.Mutex{}
var grpcSumClient pb.DeLogicClient
var grpcMultClient pb.DeLogicClient

func PostMult(c *gin.Context) {
	var input *pb.Input

	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := grpcMultClient.Calc(ctx, input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func PostSum(c *gin.Context) {
	var input *pb.Input

	if err := c.BindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	result, err := grpcSumClient.Calc(ctx, input)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}

func init() {
	connSum, err := grpc.Dial("service2:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcSumClient = pb.NewDeLogicClient(connSum)

	connMult, err := grpc.Dial("service1:9002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	grpcMultClient = pb.NewDeLogicClient(connMult)
}

func main() {
	router := gin.Default()

	router.POST("/sum", PostSum)
	router.POST("/mult", PostMult)

	router.Run()
}
