// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/costandusagereportservice"
	"go.temporal.io/sdk/workflow"
)

type CostandUsageReportServiceClient interface {
	DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error)
	DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandUsageReportServiceDeleteReportDefinitionFuture

	DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error)
	DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandUsageReportServiceDescribeReportDefinitionsFuture

	ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error)
	ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandUsageReportServiceModifyReportDefinitionFuture

	PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error)
	PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandUsageReportServicePutReportDefinitionFuture
}

type CostandUsageReportServiceStub struct{}

func NewCostandUsageReportServiceStub() CostandUsageReportServiceClient {
	return &CostandUsageReportServiceStub{}
}

type CostandUsageReportServiceDeleteReportDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostandUsageReportServiceDeleteReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	var output costandusagereportservice.DeleteReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServiceDescribeReportDefinitionsFuture struct {
	Future workflow.Future
}

func (r *CostandUsageReportServiceDescribeReportDefinitionsFuture) Get(ctx workflow.Context) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	var output costandusagereportservice.DescribeReportDefinitionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServiceModifyReportDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostandUsageReportServiceModifyReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	var output costandusagereportservice.ModifyReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CostandUsageReportServicePutReportDefinitionFuture struct {
	Future workflow.Future
}

func (r *CostandUsageReportServicePutReportDefinitionFuture) Get(ctx workflow.Context) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	var output costandusagereportservice.PutReportDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CostandUsageReportServiceStub) DeleteReportDefinition(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) (*costandusagereportservice.DeleteReportDefinitionOutput, error) {
	var output costandusagereportservice.DeleteReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.DeleteReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostandUsageReportServiceStub) DeleteReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.DeleteReportDefinitionInput) *CostandUsageReportServiceDeleteReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.DeleteReportDefinition", input)
	return &CostandUsageReportServiceDeleteReportDefinitionFuture{Future: future}
}

func (a *CostandUsageReportServiceStub) DescribeReportDefinitions(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) (*costandusagereportservice.DescribeReportDefinitionsOutput, error) {
	var output costandusagereportservice.DescribeReportDefinitionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.DescribeReportDefinitions", input).Get(ctx, &output)
	return &output, err
}

func (a *CostandUsageReportServiceStub) DescribeReportDefinitionsAsync(ctx workflow.Context, input *costandusagereportservice.DescribeReportDefinitionsInput) *CostandUsageReportServiceDescribeReportDefinitionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.DescribeReportDefinitions", input)
	return &CostandUsageReportServiceDescribeReportDefinitionsFuture{Future: future}
}

func (a *CostandUsageReportServiceStub) ModifyReportDefinition(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) (*costandusagereportservice.ModifyReportDefinitionOutput, error) {
	var output costandusagereportservice.ModifyReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.ModifyReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostandUsageReportServiceStub) ModifyReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.ModifyReportDefinitionInput) *CostandUsageReportServiceModifyReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.ModifyReportDefinition", input)
	return &CostandUsageReportServiceModifyReportDefinitionFuture{Future: future}
}

func (a *CostandUsageReportServiceStub) PutReportDefinition(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) (*costandusagereportservice.PutReportDefinitionOutput, error) {
	var output costandusagereportservice.PutReportDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.PutReportDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *CostandUsageReportServiceStub) PutReportDefinitionAsync(ctx workflow.Context, input *costandusagereportservice.PutReportDefinitionInput) *CostandUsageReportServicePutReportDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.costandusagereportservice.PutReportDefinition", input)
	return &CostandUsageReportServicePutReportDefinitionFuture{Future: future}
}
