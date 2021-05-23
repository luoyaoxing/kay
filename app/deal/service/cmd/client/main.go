package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "kay/api/deal/v1"
	"kay/pkg/utils/dec"
	"log"
)

func main() {
	conn, err := grpc.Dial(":19001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect： %v", err)
	}
	defer conn.Close()

	client := pb.NewDealClient(conn)

	//CreateDeal(client)

	GetDeal(client)
}

func GetDeal(client pb.DealClient) {
	req := &pb.GetDealRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.GetDealRequest_Body{
			DealId: 1,
		},
	}

	reply, _ := client.GetDeal(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
	fmt.Println(dec.JsonEncode(reply.Body.Deal))
}

func CreateDeal(client pb.DealClient) {
	req := &pb.CreateDealRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.CreateDealRequest_Body{
			Deal: &pb.DealDTO{
				DealId:      1,
				SkuId:       1,
				ProductId:   825,
				ProductName: "kay保险",
				Uid:         7200601,
				CustomName:  "kay",
				Phone:       "13077065592",
				Price:       1000,
			},
		},
	}

	reply, _ := client.CreateDeal(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
	fmt.Println("success")
}
