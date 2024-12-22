package DTO

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
)

type ConnectionConfig struct {
	ConnectionId string
	DomainName   string
	Stage        string
	AwsConfig    aws.Config
	ApiEndpoint  string
}

func (connectionConfig *ConnectionConfig) NewConnectionConfig(context context.Context) {
	if context.Value("connectionId") != nil {
		connectionConfig.ConnectionId = context.Value("connectionId").(string)
	}
	if context.Value("awsConfig") != nil {
		connectionConfig.AwsConfig = context.Value("awsConfig").(aws.Config)
	}
	if context.Value("domainName") != nil {
		connectionConfig.DomainName = context.Value("domainName").(string)
	}
	if context.Value("stage") != nil {
		connectionConfig.Stage = context.Value("stage").(string)
	}
	if connectionConfig.DomainName != "" && connectionConfig.Stage != "" {
		connectionConfig.ApiEndpoint = fmt.Sprintf("https://%v/%v", connectionConfig.DomainName, connectionConfig.Stage)
	}
}

func (connectionConfig *ConnectionConfig) GetApiGatewayClient() *apigatewaymanagementapi.Client {
	return apigatewaymanagementapi.NewFromConfig(connectionConfig.AwsConfig, func(o *apigatewaymanagementapi.Options) {
		o.EndpointResolver = apigatewaymanagementapi.EndpointResolverFromURL(connectionConfig.ApiEndpoint)
	})
}
