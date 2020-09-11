// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi"
	"github.com/aws/aws-sdk-go/service/apigatewaymanagementapi/apigatewaymanagementapiiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ApiGatewayManagementApiActivities struct {
	client apigatewaymanagementapiiface.ApiGatewayManagementApiAPI
}

func NewApiGatewayManagementApiActivities(session *session.Session, config ...*aws.Config) *ApiGatewayManagementApiActivities {
	client := apigatewaymanagementapi.New(session, config...)
	return &ApiGatewayManagementApiActivities{client: client}
}

func (a *ApiGatewayManagementApiActivities) DeleteConnection(ctx context.Context, input *apigatewaymanagementapi.DeleteConnectionInput) (*apigatewaymanagementapi.DeleteConnectionOutput, error) {
	return a.client.DeleteConnectionWithContext(ctx, input)
}

func (a *ApiGatewayManagementApiActivities) GetConnection(ctx context.Context, input *apigatewaymanagementapi.GetConnectionInput) (*apigatewaymanagementapi.GetConnectionOutput, error) {
	return a.client.GetConnectionWithContext(ctx, input)
}

func (a *ApiGatewayManagementApiActivities) PostToConnection(ctx context.Context, input *apigatewaymanagementapi.PostToConnectionInput) (*apigatewaymanagementapi.PostToConnectionOutput, error) {
	return a.client.PostToConnectionWithContext(ctx, input)
}
