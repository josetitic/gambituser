package awsgo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func StartAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-2"))
	fmt.Println(" > Cfg " + Cfg)
	if err != nil {
		panic("Error load configurations .aws/config " + err.Error())
	}
}
