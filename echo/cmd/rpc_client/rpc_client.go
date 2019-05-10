package main

import (
	"context"
	"fmt"
	pbecho "github.com/wncbb/grpc_pg/echo/pb"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := pbecho.NewEchoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Echo(ctx, &pbecho.StringMessage{
		Value: "haha",
	})

	fmt.Printf("resp:%#v, err:%#v\n", resp, err)

	fmt.Println("vim-go")
}
