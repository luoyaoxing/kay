package main

import (
	"context"
	"google.golang.org/grpc"
	pb "kay/api/product/v1"
	"kay/pkg/utils/dec"
	"log"
)

func main() {
	conn, err := grpc.Dial(":19000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect： %v", err)
	}
	defer conn.Close()

	client := pb.NewProductClient(conn)

	CreateProduct(client)

	//GetProductInfo(client)

	//ListProductInfo(client)

	//DeStock(client)
}

func DeStock(client pb.ProductClient) {
	req := &pb.DeStockRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.DeStockRequest_Body{
			SkuId: 2,
		},
	}

	reply, _ := client.DeStock(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
}

func ListProductInfo(client pb.ProductClient) {
	req := &pb.ListProductByProductIdRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.ListProductByProductIdRequest_Body{
			ProductId: 825,
		},
	}

	reply, _ := client.ListProductByProductId(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
	log.Println(dec.JsonEncode(reply.Body.ProductList))
}

func GetProductInfo(client pb.ProductClient) {
	req := &pb.GetProductInfoRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.GetProductInfoRequest_Body{
			SkuId: 1,
		},
	}

	reply, _ := client.GetProductInfo(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
	log.Println(dec.JsonEncode(reply.Body.Product))
}

func CreateProduct(client pb.ProductClient) {
	req := &pb.CreateProductRequest{
		Header: &pb.RequestHeader{
			Source: "kay",
			Chn:    "kay",
		},
		Body: &pb.CreateProductRequest_Body{
			Stock: 1000,
			Product: &pb.ProductDTO{
				SkuId:       9,
				VersionId:   107,
				VersionName: "华贵版",
				ProductId:   17598,
				ProductName: "kay保险",
				Price:       5000,
				ProductDesc: "性价比高",
				Attr:        "哈哈哈哈",
			},
		},
	}

	reply, _ := client.CreateProduct(context.Background(), req)
	if reply.Header.Errno > 0 {
		log.Fatal(reply.Header.Errmsg)
	}
	log.Println("success")
}
