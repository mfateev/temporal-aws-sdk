// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/aws/aws-sdk-go/service/sts/stsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type STSActivities struct {
	client stsiface.STSAPI
}

func NewSTSActivities(session *session.Session, config ...*aws.Config) *STSActivities {
	client := sts.New(session, config...)
	return &STSActivities{client: client}
}

func (a *STSActivities) AssumeRole(ctx context.Context, input *sts.AssumeRoleInput) (*sts.AssumeRoleOutput, error) {
	return a.client.AssumeRoleWithContext(ctx, input)
}

func (a *STSActivities) AssumeRoleWithSAML(ctx context.Context, input *sts.AssumeRoleWithSAMLInput) (*sts.AssumeRoleWithSAMLOutput, error) {
	return a.client.AssumeRoleWithSAMLWithContext(ctx, input)
}

func (a *STSActivities) AssumeRoleWithWebIdentity(ctx context.Context, input *sts.AssumeRoleWithWebIdentityInput) (*sts.AssumeRoleWithWebIdentityOutput, error) {
	return a.client.AssumeRoleWithWebIdentityWithContext(ctx, input)
}

func (a *STSActivities) DecodeAuthorizationMessage(ctx context.Context, input *sts.DecodeAuthorizationMessageInput) (*sts.DecodeAuthorizationMessageOutput, error) {
	return a.client.DecodeAuthorizationMessageWithContext(ctx, input)
}

func (a *STSActivities) GetAccessKeyInfo(ctx context.Context, input *sts.GetAccessKeyInfoInput) (*sts.GetAccessKeyInfoOutput, error) {
	return a.client.GetAccessKeyInfoWithContext(ctx, input)
}

func (a *STSActivities) GetCallerIdentity(ctx context.Context, input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return a.client.GetCallerIdentityWithContext(ctx, input)
}

func (a *STSActivities) GetFederationToken(ctx context.Context, input *sts.GetFederationTokenInput) (*sts.GetFederationTokenOutput, error) {
	return a.client.GetFederationTokenWithContext(ctx, input)
}

func (a *STSActivities) GetSessionToken(ctx context.Context, input *sts.GetSessionTokenInput) (*sts.GetSessionTokenOutput, error) {
	return a.client.GetSessionTokenWithContext(ctx, input)
}
