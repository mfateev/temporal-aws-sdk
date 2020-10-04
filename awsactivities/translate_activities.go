// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/translate"
	"github.com/aws/aws-sdk-go/service/translate/translateiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type TranslateActivities struct {
	client translateiface.TranslateAPI
}

func NewTranslateActivities(session *session.Session, config ...*aws.Config) *TranslateActivities {
	client := translate.New(session, config...)
	return &TranslateActivities{client: client}
}

func (a *TranslateActivities) DeleteTerminology(ctx context.Context, input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
	return a.client.DeleteTerminologyWithContext(ctx, input)
}

func (a *TranslateActivities) DescribeTextTranslationJob(ctx context.Context, input *translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error) {
	return a.client.DescribeTextTranslationJobWithContext(ctx, input)
}

func (a *TranslateActivities) GetTerminology(ctx context.Context, input *translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error) {
	return a.client.GetTerminologyWithContext(ctx, input)
}

func (a *TranslateActivities) ImportTerminology(ctx context.Context, input *translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error) {
	return a.client.ImportTerminologyWithContext(ctx, input)
}

func (a *TranslateActivities) ListTerminologies(ctx context.Context, input *translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error) {
	return a.client.ListTerminologiesWithContext(ctx, input)
}

func (a *TranslateActivities) ListTextTranslationJobs(ctx context.Context, input *translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error) {
	return a.client.ListTextTranslationJobsWithContext(ctx, input)
}

func (a *TranslateActivities) StartTextTranslationJob(ctx context.Context, input *translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.StartTextTranslationJobWithContext(ctx, input)
}

func (a *TranslateActivities) StopTextTranslationJob(ctx context.Context, input *translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error) {
	return a.client.StopTextTranslationJobWithContext(ctx, input)
}

func (a *TranslateActivities) Text(ctx context.Context, input *translate.TextInput) (*translate.TextOutput, error) {
	return a.client.TextWithContext(ctx, input)
}
