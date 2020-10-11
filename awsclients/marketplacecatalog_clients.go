// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/marketplacecatalog"
	"go.temporal.io/sdk/workflow"
)

type MarketplaceCatalogClient interface {
	CancelChangeSet(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error)
	CancelChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) *MarketplaceCatalogCancelChangeSetFuture

	DescribeChangeSet(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error)
	DescribeChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) *MarketplaceCatalogDescribeChangeSetFuture

	DescribeEntity(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error)
	DescribeEntityAsync(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) *MarketplaceCatalogDescribeEntityFuture

	ListChangeSets(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error)
	ListChangeSetsAsync(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) *MarketplaceCatalogListChangeSetsFuture

	ListEntities(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error)
	ListEntitiesAsync(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) *MarketplaceCatalogListEntitiesFuture

	StartChangeSet(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error)
	StartChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) *MarketplaceCatalogStartChangeSetFuture
}

type MarketplaceCatalogStub struct{}

func NewMarketplaceCatalogStub() MarketplaceCatalogClient {
	return &MarketplaceCatalogStub{}
}

type MarketplaceCatalogCancelChangeSetFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogCancelChangeSetFuture) Get(ctx workflow.Context) (*marketplacecatalog.CancelChangeSetOutput, error) {
	var output marketplacecatalog.CancelChangeSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MarketplaceCatalogDescribeChangeSetFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogDescribeChangeSetFuture) Get(ctx workflow.Context) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	var output marketplacecatalog.DescribeChangeSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MarketplaceCatalogDescribeEntityFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogDescribeEntityFuture) Get(ctx workflow.Context) (*marketplacecatalog.DescribeEntityOutput, error) {
	var output marketplacecatalog.DescribeEntityOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MarketplaceCatalogListChangeSetsFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogListChangeSetsFuture) Get(ctx workflow.Context) (*marketplacecatalog.ListChangeSetsOutput, error) {
	var output marketplacecatalog.ListChangeSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MarketplaceCatalogListEntitiesFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogListEntitiesFuture) Get(ctx workflow.Context) (*marketplacecatalog.ListEntitiesOutput, error) {
	var output marketplacecatalog.ListEntitiesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MarketplaceCatalogStartChangeSetFuture struct {
	Future workflow.Future
}

func (r *MarketplaceCatalogStartChangeSetFuture) Get(ctx workflow.Context) (*marketplacecatalog.StartChangeSetOutput, error) {
	var output marketplacecatalog.StartChangeSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) CancelChangeSet(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) (*marketplacecatalog.CancelChangeSetOutput, error) {
	var output marketplacecatalog.CancelChangeSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.CancelChangeSet", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) CancelChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.CancelChangeSetInput) *MarketplaceCatalogCancelChangeSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.CancelChangeSet", input)
	return &MarketplaceCatalogCancelChangeSetFuture{Future: future}
}

func (a *MarketplaceCatalogStub) DescribeChangeSet(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) (*marketplacecatalog.DescribeChangeSetOutput, error) {
	var output marketplacecatalog.DescribeChangeSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.DescribeChangeSet", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) DescribeChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.DescribeChangeSetInput) *MarketplaceCatalogDescribeChangeSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.DescribeChangeSet", input)
	return &MarketplaceCatalogDescribeChangeSetFuture{Future: future}
}

func (a *MarketplaceCatalogStub) DescribeEntity(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) (*marketplacecatalog.DescribeEntityOutput, error) {
	var output marketplacecatalog.DescribeEntityOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.DescribeEntity", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) DescribeEntityAsync(ctx workflow.Context, input *marketplacecatalog.DescribeEntityInput) *MarketplaceCatalogDescribeEntityFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.DescribeEntity", input)
	return &MarketplaceCatalogDescribeEntityFuture{Future: future}
}

func (a *MarketplaceCatalogStub) ListChangeSets(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) (*marketplacecatalog.ListChangeSetsOutput, error) {
	var output marketplacecatalog.ListChangeSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.ListChangeSets", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) ListChangeSetsAsync(ctx workflow.Context, input *marketplacecatalog.ListChangeSetsInput) *MarketplaceCatalogListChangeSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.ListChangeSets", input)
	return &MarketplaceCatalogListChangeSetsFuture{Future: future}
}

func (a *MarketplaceCatalogStub) ListEntities(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) (*marketplacecatalog.ListEntitiesOutput, error) {
	var output marketplacecatalog.ListEntitiesOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.ListEntities", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) ListEntitiesAsync(ctx workflow.Context, input *marketplacecatalog.ListEntitiesInput) *MarketplaceCatalogListEntitiesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.ListEntities", input)
	return &MarketplaceCatalogListEntitiesFuture{Future: future}
}

func (a *MarketplaceCatalogStub) StartChangeSet(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) (*marketplacecatalog.StartChangeSetOutput, error) {
	var output marketplacecatalog.StartChangeSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.StartChangeSet", input).Get(ctx, &output)
	return &output, err
}

func (a *MarketplaceCatalogStub) StartChangeSetAsync(ctx workflow.Context, input *marketplacecatalog.StartChangeSetInput) *MarketplaceCatalogStartChangeSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecatalog.StartChangeSet", input)
	return &MarketplaceCatalogStartChangeSetFuture{Future: future}
}
