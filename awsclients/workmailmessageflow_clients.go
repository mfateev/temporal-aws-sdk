// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/workmailmessageflow"
	"go.temporal.io/sdk/workflow"
)

type WorkMailMessageFlowClient interface {
	GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error)
	GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkMailMessageFlowGetRawMessageContentFuture
}

type WorkMailMessageFlowStub struct{}

func NewWorkMailMessageFlowStub() WorkMailMessageFlowClient {
	return &WorkMailMessageFlowStub{}
}

type WorkMailMessageFlowGetRawMessageContentFuture struct {
	Future workflow.Future
}

func (r *WorkMailMessageFlowGetRawMessageContentFuture) Get(ctx workflow.Context) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *WorkMailMessageFlowStub) GetRawMessageContent(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) (*workmailmessageflow.GetRawMessageContentOutput, error) {
	var output workmailmessageflow.GetRawMessageContentOutput
	err := workflow.ExecuteActivity(ctx, "aws.workmailmessageflow.GetRawMessageContent", input).Get(ctx, &output)
	return &output, err
}

func (a *WorkMailMessageFlowStub) GetRawMessageContentAsync(ctx workflow.Context, input *workmailmessageflow.GetRawMessageContentInput) *WorkMailMessageFlowGetRawMessageContentFuture {
	future := workflow.ExecuteActivity(ctx, "aws.workmailmessageflow.GetRawMessageContent", input)
	return &WorkMailMessageFlowGetRawMessageContentFuture{Future: future}
}
