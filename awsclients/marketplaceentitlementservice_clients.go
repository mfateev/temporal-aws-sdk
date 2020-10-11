// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/marketplaceentitlementservice"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceEntitlementServiceClient interface {
	GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error)
	GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceEntitlementServiceGetEntitlementsFuture
}

type MarketplaceEntitlementServiceStub struct{}

func NewMarketplaceEntitlementServiceStub() MarketplaceEntitlementServiceClient {
	return &MarketplaceEntitlementServiceStub{}
}

type MarketplaceEntitlementServiceGetEntitlementsFuture struct {
	Future workflow.Future
}

func (r *MarketplaceEntitlementServiceGetEntitlementsFuture) Get(ctx workflow.Context) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceEntitlementServiceStub) GetEntitlements(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) (*marketplaceentitlementservice.GetEntitlementsOutput, error) {
	var output marketplaceentitlementservice.GetEntitlementsOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplaceentitlementservice.GetEntitlements", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceEntitlementServiceStub) GetEntitlementsAsync(ctx workflow.Context, input *marketplaceentitlementservice.GetEntitlementsInput) *MarketplaceEntitlementServiceGetEntitlementsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplaceentitlementservice.GetEntitlements", input)
	return &MarketplaceEntitlementServiceGetEntitlementsFuture{Future: future}
}
