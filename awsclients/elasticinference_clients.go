// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elasticinference"
	"go.temporal.io/sdk/workflow"
)

type ElasticInferenceClient interface {
	DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error)
	DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticInferenceDescribeAcceleratorOfferingsFuture

	DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error)
	DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticInferenceDescribeAcceleratorTypesFuture

	DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error)
	DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticInferenceDescribeAcceleratorsFuture

	ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticInferenceListTagsForResourceFuture

	TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticInferenceTagResourceFuture

	UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticInferenceUntagResourceFuture
}

type ElasticInferenceStub struct{}

func NewElasticInferenceStub() ElasticInferenceClient {
	return &ElasticInferenceStub{}
}

type ElasticInferenceDescribeAcceleratorOfferingsFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorOfferingsFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	var output elasticinference.DescribeAcceleratorOfferingsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceDescribeAcceleratorTypesFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorTypesFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	var output elasticinference.DescribeAcceleratorTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceDescribeAcceleratorsFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceDescribeAcceleratorsFuture) Get(ctx workflow.Context) (*elasticinference.DescribeAcceleratorsOutput, error) {
	var output elasticinference.DescribeAcceleratorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceListTagsForResourceFuture) Get(ctx workflow.Context) (*elasticinference.ListTagsForResourceOutput, error) {
	var output elasticinference.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceTagResourceFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceTagResourceFuture) Get(ctx workflow.Context) (*elasticinference.TagResourceOutput, error) {
	var output elasticinference.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ElasticInferenceUntagResourceFuture struct {
	Future workflow.Future
}

func (r *ElasticInferenceUntagResourceFuture) Get(ctx workflow.Context) (*elasticinference.UntagResourceOutput, error) {
	var output elasticinference.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorOfferings(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) (*elasticinference.DescribeAcceleratorOfferingsOutput, error) {
	var output elasticinference.DescribeAcceleratorOfferingsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAcceleratorOfferings", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorOfferingsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorOfferingsInput) *ElasticInferenceDescribeAcceleratorOfferingsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAcceleratorOfferings", input)
	return &ElasticInferenceDescribeAcceleratorOfferingsFuture{Future: future}
}

func (a *ElasticInferenceStub) DescribeAcceleratorTypes(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) (*elasticinference.DescribeAcceleratorTypesOutput, error) {
	var output elasticinference.DescribeAcceleratorTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAcceleratorTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorTypesAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorTypesInput) *ElasticInferenceDescribeAcceleratorTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAcceleratorTypes", input)
	return &ElasticInferenceDescribeAcceleratorTypesFuture{Future: future}
}

func (a *ElasticInferenceStub) DescribeAccelerators(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) (*elasticinference.DescribeAcceleratorsOutput, error) {
	var output elasticinference.DescribeAcceleratorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAccelerators", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) DescribeAcceleratorsAsync(ctx workflow.Context, input *elasticinference.DescribeAcceleratorsInput) *ElasticInferenceDescribeAcceleratorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.DescribeAccelerators", input)
	return &ElasticInferenceDescribeAcceleratorsFuture{Future: future}
}

func (a *ElasticInferenceStub) ListTagsForResource(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) (*elasticinference.ListTagsForResourceOutput, error) {
	var output elasticinference.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) ListTagsForResourceAsync(ctx workflow.Context, input *elasticinference.ListTagsForResourceInput) *ElasticInferenceListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.ListTagsForResource", input)
	return &ElasticInferenceListTagsForResourceFuture{Future: future}
}

func (a *ElasticInferenceStub) TagResource(ctx workflow.Context, input *elasticinference.TagResourceInput) (*elasticinference.TagResourceOutput, error) {
	var output elasticinference.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) TagResourceAsync(ctx workflow.Context, input *elasticinference.TagResourceInput) *ElasticInferenceTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.TagResource", input)
	return &ElasticInferenceTagResourceFuture{Future: future}
}

func (a *ElasticInferenceStub) UntagResource(ctx workflow.Context, input *elasticinference.UntagResourceInput) (*elasticinference.UntagResourceOutput, error) {
	var output elasticinference.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.elasticinference.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *ElasticInferenceStub) UntagResourceAsync(ctx workflow.Context, input *elasticinference.UntagResourceInput) *ElasticInferenceUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elasticinference.UntagResource", input)
	return &ElasticInferenceUntagResourceFuture{Future: future}
}
