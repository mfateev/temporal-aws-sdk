// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/ssooidc"
	"go.temporal.io/sdk/workflow"
)

type SSOOIDCClient interface {
	CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error)
	CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *SSOOIDCCreateTokenFuture

	RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error)
	RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *SSOOIDCRegisterClientFuture

	StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error)
	StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *SSOOIDCStartDeviceAuthorizationFuture
}

type SSOOIDCStub struct{}

func NewSSOOIDCStub() SSOOIDCClient {
	return &SSOOIDCStub{}
}

type SSOOIDCCreateTokenFuture struct {
	Future workflow.Future
}

func (r *SSOOIDCCreateTokenFuture) Get(ctx workflow.Context) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOOIDCRegisterClientFuture struct {
	Future workflow.Future
}

func (r *SSOOIDCRegisterClientFuture) Get(ctx workflow.Context) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SSOOIDCStartDeviceAuthorizationFuture struct {
	Future workflow.Future
}

func (r *SSOOIDCStartDeviceAuthorizationFuture) Get(ctx workflow.Context) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) CreateToken(ctx workflow.Context, input *ssooidc.CreateTokenInput) (*ssooidc.CreateTokenOutput, error) {
	var output ssooidc.CreateTokenOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.CreateToken", input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) CreateTokenAsync(ctx workflow.Context, input *ssooidc.CreateTokenInput) *SSOOIDCCreateTokenFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.CreateToken", input)
	return &SSOOIDCCreateTokenFuture{Future: future}
}

func (a *SSOOIDCStub) RegisterClient(ctx workflow.Context, input *ssooidc.RegisterClientInput) (*ssooidc.RegisterClientOutput, error) {
	var output ssooidc.RegisterClientOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.RegisterClient", input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) RegisterClientAsync(ctx workflow.Context, input *ssooidc.RegisterClientInput) *SSOOIDCRegisterClientFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.RegisterClient", input)
	return &SSOOIDCRegisterClientFuture{Future: future}
}

func (a *SSOOIDCStub) StartDeviceAuthorization(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) (*ssooidc.StartDeviceAuthorizationOutput, error) {
	var output ssooidc.StartDeviceAuthorizationOutput
	err := workflow.ExecuteActivity(ctx, "aws.ssooidc.StartDeviceAuthorization", input).Get(ctx, &output)
	return &output, err
}

func (a *SSOOIDCStub) StartDeviceAuthorizationAsync(ctx workflow.Context, input *ssooidc.StartDeviceAuthorizationInput) *SSOOIDCStartDeviceAuthorizationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.ssooidc.StartDeviceAuthorization", input)
	return &SSOOIDCStartDeviceAuthorizationFuture{Future: future}
}
