// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"go.temporal.io/sdk/workflow"
)

type TranscribeStreamingServiceClient interface {
	StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
	StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribeStreamingServiceStartStreamTranscriptionFuture
}

type TranscribeStreamingServiceStub struct{}

func NewTranscribeStreamingServiceStub() TranscribeStreamingServiceClient {
	return &TranscribeStreamingServiceStub{}
}

type TranscribeStreamingServiceStartStreamTranscriptionFuture struct {
	Future workflow.Future
}

func (r *TranscribeStreamingServiceStartStreamTranscriptionFuture) Get(ctx workflow.Context) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	var output transcribestreamingservice.StartStreamTranscriptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *TranscribeStreamingServiceStub) StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
	var output transcribestreamingservice.StartStreamTranscriptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.transcribestreamingservice.StartStreamTranscription", input).Get(ctx, &output)
	return &output, err
}

func (a *TranscribeStreamingServiceStub) StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribeStreamingServiceStartStreamTranscriptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.transcribestreamingservice.StartStreamTranscription", input)
	return &TranscribeStreamingServiceStartStreamTranscriptionFuture{Future: future}
}
