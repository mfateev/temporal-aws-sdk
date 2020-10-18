// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package datapipeline

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"github.com/aws/aws-sdk-go/service/datapipeline/datapipelineiface"
	"go.temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client datapipelineiface.DataPipelineAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := datapipeline.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (datapipelineiface.DataPipelineAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return datapipeline.New(sess), nil
}

func (a *Activities) ActivatePipeline(ctx context.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ActivatePipelineWithContext(ctx, input)
}

func (a *Activities) AddTags(ctx context.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AddTagsWithContext(ctx, input)
}

func (a *Activities) CreatePipeline(ctx context.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePipelineWithContext(ctx, input)
}

func (a *Activities) DeactivatePipeline(ctx context.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeactivatePipelineWithContext(ctx, input)
}

func (a *Activities) DeletePipeline(ctx context.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePipelineWithContext(ctx, input)
}

func (a *Activities) DescribeObjects(ctx context.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeObjectsWithContext(ctx, input)
}

func (a *Activities) DescribePipelines(ctx context.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePipelinesWithContext(ctx, input)
}

func (a *Activities) EvaluateExpression(ctx context.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.EvaluateExpressionWithContext(ctx, input)
}

func (a *Activities) GetPipelineDefinition(ctx context.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPipelineDefinitionWithContext(ctx, input)
}

func (a *Activities) ListPipelines(ctx context.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPipelinesWithContext(ctx, input)
}

func (a *Activities) PollForTask(ctx context.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PollForTaskWithContext(ctx, input)
}

func (a *Activities) PutPipelineDefinition(ctx context.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPipelineDefinitionWithContext(ctx, input)
}

func (a *Activities) QueryObjects(ctx context.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.QueryObjectsWithContext(ctx, input)
}

func (a *Activities) RemoveTags(ctx context.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.RemoveTagsWithContext(ctx, input)
}

func (a *Activities) ReportTaskProgress(ctx context.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReportTaskProgressWithContext(ctx, input)
}

func (a *Activities) ReportTaskRunnerHeartbeat(ctx context.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ReportTaskRunnerHeartbeatWithContext(ctx, input)
}

func (a *Activities) SetStatus(ctx context.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetStatusWithContext(ctx, input)
}

func (a *Activities) SetTaskStatus(ctx context.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.SetTaskStatusWithContext(ctx, input)
}

func (a *Activities) ValidatePipelineDefinition(ctx context.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ValidatePipelineDefinitionWithContext(ctx, input)
}