// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"go.temporal.io/sdk/workflow"
)

type SavingsPlansClient interface {
	CreateSavingsPlan(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error)
	CreateSavingsPlanAsync(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) *SavingsPlansCreateSavingsPlanFuture

	DeleteQueuedSavingsPlan(ctx workflow.Context, input *savingsplans.DeleteQueuedSavingsPlanInput) (*savingsplans.DeleteQueuedSavingsPlanOutput, error)
	DeleteQueuedSavingsPlanAsync(ctx workflow.Context, input *savingsplans.DeleteQueuedSavingsPlanInput) *SavingsPlansDeleteQueuedSavingsPlanFuture

	DescribeSavingsPlanRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
	DescribeSavingsPlanRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) *SavingsPlansDescribeSavingsPlanRatesFuture

	DescribeSavingsPlans(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error)
	DescribeSavingsPlansAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) *SavingsPlansDescribeSavingsPlansFuture

	DescribeSavingsPlansOfferingRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
	DescribeSavingsPlansOfferingRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) *SavingsPlansDescribeSavingsPlansOfferingRatesFuture

	DescribeSavingsPlansOfferings(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
	DescribeSavingsPlansOfferingsAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) *SavingsPlansDescribeSavingsPlansOfferingsFuture

	ListTagsForResource(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) *SavingsPlansListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *savingsplans.TagResourceInput) *SavingsPlansTagResourceFuture

	UntagResource(ctx workflow.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *savingsplans.UntagResourceInput) *SavingsPlansUntagResourceFuture
}

type SavingsPlansStub struct{}

func NewSavingsPlansStub() SavingsPlansClient {
	return &SavingsPlansStub{}
}

type SavingsPlansCreateSavingsPlanFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansCreateSavingsPlanFuture) Get(ctx workflow.Context) (*savingsplans.CreateSavingsPlanOutput, error) {
	var output savingsplans.CreateSavingsPlanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansDeleteQueuedSavingsPlanFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansDeleteQueuedSavingsPlanFuture) Get(ctx workflow.Context) (*savingsplans.DeleteQueuedSavingsPlanOutput, error) {
	var output savingsplans.DeleteQueuedSavingsPlanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansDescribeSavingsPlanRatesFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansDescribeSavingsPlanRatesFuture) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	var output savingsplans.DescribeSavingsPlanRatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansDescribeSavingsPlansFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansDescribeSavingsPlansFuture) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOutput, error) {
	var output savingsplans.DescribeSavingsPlansOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansDescribeSavingsPlansOfferingRatesFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansDescribeSavingsPlansOfferingRatesFuture) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	var output savingsplans.DescribeSavingsPlansOfferingRatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansDescribeSavingsPlansOfferingsFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansDescribeSavingsPlansOfferingsFuture) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	var output savingsplans.DescribeSavingsPlansOfferingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansListTagsForResourceFuture) Get(ctx workflow.Context) (*savingsplans.ListTagsForResourceOutput, error) {
	var output savingsplans.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansTagResourceFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansTagResourceFuture) Get(ctx workflow.Context) (*savingsplans.TagResourceOutput, error) {
	var output savingsplans.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SavingsPlansUntagResourceFuture struct {
	Future workflow.Future
}

func (r *SavingsPlansUntagResourceFuture) Get(ctx workflow.Context) (*savingsplans.UntagResourceOutput, error) {
	var output savingsplans.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) CreateSavingsPlan(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
	var output savingsplans.CreateSavingsPlanOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.CreateSavingsPlan", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) CreateSavingsPlanAsync(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) *SavingsPlansCreateSavingsPlanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.CreateSavingsPlan", input)
	return &SavingsPlansCreateSavingsPlanFuture{Future: future}
}

func (a *SavingsPlansStub) DeleteQueuedSavingsPlan(ctx workflow.Context, input *savingsplans.DeleteQueuedSavingsPlanInput) (*savingsplans.DeleteQueuedSavingsPlanOutput, error) {
	var output savingsplans.DeleteQueuedSavingsPlanOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.DeleteQueuedSavingsPlan", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) DeleteQueuedSavingsPlanAsync(ctx workflow.Context, input *savingsplans.DeleteQueuedSavingsPlanInput) *SavingsPlansDeleteQueuedSavingsPlanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.DeleteQueuedSavingsPlan", input)
	return &SavingsPlansDeleteQueuedSavingsPlanFuture{Future: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlanRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
	var output savingsplans.DescribeSavingsPlanRatesOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlanRates", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlanRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) *SavingsPlansDescribeSavingsPlanRatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlanRates", input)
	return &SavingsPlansDescribeSavingsPlanRatesFuture{Future: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlans(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
	var output savingsplans.DescribeSavingsPlansOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlans", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) *SavingsPlansDescribeSavingsPlansFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlans", input)
	return &SavingsPlansDescribeSavingsPlansFuture{Future: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
	var output savingsplans.DescribeSavingsPlansOfferingRatesOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlansOfferingRates", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) *SavingsPlansDescribeSavingsPlansOfferingRatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlansOfferingRates", input)
	return &SavingsPlansDescribeSavingsPlansOfferingRatesFuture{Future: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferings(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
	var output savingsplans.DescribeSavingsPlansOfferingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlansOfferings", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingsAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) *SavingsPlansDescribeSavingsPlansOfferingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.DescribeSavingsPlansOfferings", input)
	return &SavingsPlansDescribeSavingsPlansOfferingsFuture{Future: future}
}

func (a *SavingsPlansStub) ListTagsForResource(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
	var output savingsplans.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) ListTagsForResourceAsync(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) *SavingsPlansListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.ListTagsForResource", input)
	return &SavingsPlansListTagsForResourceFuture{Future: future}
}

func (a *SavingsPlansStub) TagResource(ctx workflow.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error) {
	var output savingsplans.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) TagResourceAsync(ctx workflow.Context, input *savingsplans.TagResourceInput) *SavingsPlansTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.TagResource", input)
	return &SavingsPlansTagResourceFuture{Future: future}
}

func (a *SavingsPlansStub) UntagResource(ctx workflow.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error) {
	var output savingsplans.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.savingsplans.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *SavingsPlansStub) UntagResourceAsync(ctx workflow.Context, input *savingsplans.UntagResourceInput) *SavingsPlansUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.savingsplans.UntagResource", input)
	return &SavingsPlansUntagResourceFuture{Future: future}
}
