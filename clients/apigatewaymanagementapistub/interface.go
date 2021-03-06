// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package apigatewaymanagementapistub

import (
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DeleteConnection(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error)
	DeleteConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.DeleteConnectionInput) *ApiGatewayManagementApiDeleteConnectionFuture

	GetConnection(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error)
	GetConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.GetConnectionInput) *ApiGatewayManagementApiGetConnectionFuture

	PostToConnection(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error)
	PostToConnectionAsync(ctx workflow.Context, input *apigatewaymanagementapi.PostToConnectionInput) *ApiGatewayManagementApiPostToConnectionFuture
}

func NewClient() Client {
	return &stub{}
}
