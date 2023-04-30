package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	svc "github.com/bmoffatt/smithy-lambda-url-example/model"
)

func main() {
	endpoint := os.Args[1]
	cfg, _ := config.LoadDefaultConfig(context.Background())
	client := svc.New(svc.Options{
		EndpointResolver: svc.EndpointResolverFromURL(endpoint),
		Credentials:      cfg.Credentials,
		Region:           "us-west-2",
	})
	res, err := client.Wave(context.Background(), &svc.WaveInput{
		Name: ref("bruh"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("response\n----------\ntime: %v\ntext: %s\n", *res.Time, *res.Text)
}

func ref[T any](t T) *T {
	return &t
}
