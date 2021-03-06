// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elastictranscoderstub

import (
	"github.com/aws/aws-sdk-go/service/elastictranscoder"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CancelJob(ctx workflow.Context, input *elastictranscoder.CancelJobInput) (*elastictranscoder.CancelJobOutput, error)
	CancelJobAsync(ctx workflow.Context, input *elastictranscoder.CancelJobInput) *ElasticTranscoderCancelJobFuture

	CreateJob(ctx workflow.Context, input *elastictranscoder.CreateJobInput) (*elastictranscoder.CreateJobResponse, error)
	CreateJobAsync(ctx workflow.Context, input *elastictranscoder.CreateJobInput) *ElasticTranscoderCreateJobFuture

	CreatePipeline(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) (*elastictranscoder.CreatePipelineOutput, error)
	CreatePipelineAsync(ctx workflow.Context, input *elastictranscoder.CreatePipelineInput) *ElasticTranscoderCreatePipelineFuture

	CreatePreset(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) (*elastictranscoder.CreatePresetOutput, error)
	CreatePresetAsync(ctx workflow.Context, input *elastictranscoder.CreatePresetInput) *ElasticTranscoderCreatePresetFuture

	DeletePipeline(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) (*elastictranscoder.DeletePipelineOutput, error)
	DeletePipelineAsync(ctx workflow.Context, input *elastictranscoder.DeletePipelineInput) *ElasticTranscoderDeletePipelineFuture

	DeletePreset(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) (*elastictranscoder.DeletePresetOutput, error)
	DeletePresetAsync(ctx workflow.Context, input *elastictranscoder.DeletePresetInput) *ElasticTranscoderDeletePresetFuture

	ListJobsByPipeline(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) (*elastictranscoder.ListJobsByPipelineOutput, error)
	ListJobsByPipelineAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByPipelineInput) *ElasticTranscoderListJobsByPipelineFuture

	ListJobsByStatus(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) (*elastictranscoder.ListJobsByStatusOutput, error)
	ListJobsByStatusAsync(ctx workflow.Context, input *elastictranscoder.ListJobsByStatusInput) *ElasticTranscoderListJobsByStatusFuture

	ListPipelines(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) (*elastictranscoder.ListPipelinesOutput, error)
	ListPipelinesAsync(ctx workflow.Context, input *elastictranscoder.ListPipelinesInput) *ElasticTranscoderListPipelinesFuture

	ListPresets(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) (*elastictranscoder.ListPresetsOutput, error)
	ListPresetsAsync(ctx workflow.Context, input *elastictranscoder.ListPresetsInput) *ElasticTranscoderListPresetsFuture

	ReadJob(ctx workflow.Context, input *elastictranscoder.ReadJobInput) (*elastictranscoder.ReadJobOutput, error)
	ReadJobAsync(ctx workflow.Context, input *elastictranscoder.ReadJobInput) *ElasticTranscoderReadJobFuture

	ReadPipeline(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) (*elastictranscoder.ReadPipelineOutput, error)
	ReadPipelineAsync(ctx workflow.Context, input *elastictranscoder.ReadPipelineInput) *ElasticTranscoderReadPipelineFuture

	ReadPreset(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) (*elastictranscoder.ReadPresetOutput, error)
	ReadPresetAsync(ctx workflow.Context, input *elastictranscoder.ReadPresetInput) *ElasticTranscoderReadPresetFuture

	TestRole(ctx workflow.Context, input *elastictranscoder.TestRoleInput) (*elastictranscoder.TestRoleOutput, error)
	TestRoleAsync(ctx workflow.Context, input *elastictranscoder.TestRoleInput) *ElasticTranscoderTestRoleFuture

	UpdatePipeline(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) (*elastictranscoder.UpdatePipelineOutput, error)
	UpdatePipelineAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineInput) *ElasticTranscoderUpdatePipelineFuture

	UpdatePipelineNotifications(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) (*elastictranscoder.UpdatePipelineNotificationsOutput, error)
	UpdatePipelineNotificationsAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineNotificationsInput) *ElasticTranscoderUpdatePipelineNotificationsFuture

	UpdatePipelineStatus(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) (*elastictranscoder.UpdatePipelineStatusOutput, error)
	UpdatePipelineStatusAsync(ctx workflow.Context, input *elastictranscoder.UpdatePipelineStatusInput) *ElasticTranscoderUpdatePipelineStatusFuture

	WaitUntilJobComplete(ctx workflow.Context, input *elastictranscoder.ReadJobInput) error
	WaitUntilJobCompleteAsync(ctx workflow.Context, input *elastictranscoder.ReadJobInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
