// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalyticsv2"
	"go.temporal.io/sdk/workflow"
)

type KinesisAnalyticsV2Client interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture

	AddApplicationInput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) *KinesisAnalyticsV2AddApplicationInputFuture

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) *KinesisAnalyticsV2AddApplicationOutputFuture

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) *KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture

	AddApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error)
	AddApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) *KinesisAnalyticsV2AddApplicationVpcConfigurationFuture

	CreateApplication(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) *KinesisAnalyticsV2CreateApplicationFuture

	CreateApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error)
	CreateApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) *KinesisAnalyticsV2CreateApplicationSnapshotFuture

	DeleteApplication(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) *KinesisAnalyticsV2DeleteApplicationFuture

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) *KinesisAnalyticsV2DeleteApplicationOutputFuture

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) *KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture

	DeleteApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error)
	DeleteApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) *KinesisAnalyticsV2DeleteApplicationSnapshotFuture

	DeleteApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error)
	DeleteApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) *KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture

	DescribeApplication(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) *KinesisAnalyticsV2DescribeApplicationFuture

	DescribeApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error)
	DescribeApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) *KinesisAnalyticsV2DescribeApplicationSnapshotFuture

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) *KinesisAnalyticsV2DiscoverInputSchemaFuture

	ListApplicationSnapshots(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error)
	ListApplicationSnapshotsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) *KinesisAnalyticsV2ListApplicationSnapshotsFuture

	ListApplications(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) *KinesisAnalyticsV2ListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) *KinesisAnalyticsV2ListTagsForResourceFuture

	StartApplication(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) *KinesisAnalyticsV2StartApplicationFuture

	StopApplication(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) *KinesisAnalyticsV2StopApplicationFuture

	TagResource(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) *KinesisAnalyticsV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) *KinesisAnalyticsV2UntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) *KinesisAnalyticsV2UpdateApplicationFuture
}

type KinesisAnalyticsV2Stub struct{}

func NewKinesisAnalyticsV2Stub() KinesisAnalyticsV2Client {
	return &KinesisAnalyticsV2Stub{}
}

type KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2AddApplicationInputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationInputFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
	var output kinesisanalyticsv2.AddApplicationInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2AddApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
	var output kinesisanalyticsv2.AddApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2AddApplicationVpcConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2AddApplicationVpcConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
	var output kinesisanalyticsv2.AddApplicationVpcConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2CreateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2CreateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
	var output kinesisanalyticsv2.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2CreateApplicationSnapshotFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2CreateApplicationSnapshotFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.CreateApplicationSnapshotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationSnapshotFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationSnapshotFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationSnapshotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DescribeApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DescribeApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
	var output kinesisanalyticsv2.DescribeApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DescribeApplicationSnapshotFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DescribeApplicationSnapshotFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.DescribeApplicationSnapshotOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2DiscoverInputSchemaFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2DiscoverInputSchemaFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
	var output kinesisanalyticsv2.DiscoverInputSchemaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2ListApplicationSnapshotsFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2ListApplicationSnapshotsFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
	var output kinesisanalyticsv2.ListApplicationSnapshotsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2ListApplicationsFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2ListApplicationsFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
	var output kinesisanalyticsv2.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2ListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2ListTagsForResourceFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
	var output kinesisanalyticsv2.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2StartApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2StartApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.StartApplicationOutput, error) {
	var output kinesisanalyticsv2.StartApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2StopApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2StopApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.StopApplicationOutput, error) {
	var output kinesisanalyticsv2.StopApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2TagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2TagResourceFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.TagResourceOutput, error) {
	var output kinesisanalyticsv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2UntagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2UntagResourceFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.UntagResourceOutput, error) {
	var output kinesisanalyticsv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsV2UpdateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsV2UpdateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
	var output kinesisanalyticsv2.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationCloudWatchLoggingOption", input)
	return &KinesisAnalyticsV2AddApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) (*kinesisanalyticsv2.AddApplicationInputOutput, error) {
	var output kinesisanalyticsv2.AddApplicationInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationInput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputInput) *KinesisAnalyticsV2AddApplicationInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationInput", input)
	return &KinesisAnalyticsV2AddApplicationInputFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalyticsv2.AddApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationInputProcessingConfiguration", input)
	return &KinesisAnalyticsV2AddApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) (*kinesisanalyticsv2.AddApplicationOutputOutput, error) {
	var output kinesisanalyticsv2.AddApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationOutputInput) *KinesisAnalyticsV2AddApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationOutput", input)
	return &KinesisAnalyticsV2AddApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalyticsv2.AddApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationReferenceDataSourceInput) *KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationReferenceDataSource", input)
	return &KinesisAnalyticsV2AddApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) AddApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) (*kinesisanalyticsv2.AddApplicationVpcConfigurationOutput, error) {
	var output kinesisanalyticsv2.AddApplicationVpcConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationVpcConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) AddApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.AddApplicationVpcConfigurationInput) *KinesisAnalyticsV2AddApplicationVpcConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.AddApplicationVpcConfiguration", input)
	return &KinesisAnalyticsV2AddApplicationVpcConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) CreateApplication(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) (*kinesisanalyticsv2.CreateApplicationOutput, error) {
	var output kinesisanalyticsv2.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationInput) *KinesisAnalyticsV2CreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.CreateApplication", input)
	return &KinesisAnalyticsV2CreateApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) CreateApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) (*kinesisanalyticsv2.CreateApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.CreateApplicationSnapshotOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.CreateApplicationSnapshot", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) CreateApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.CreateApplicationSnapshotInput) *KinesisAnalyticsV2CreateApplicationSnapshotFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.CreateApplicationSnapshot", input)
	return &KinesisAnalyticsV2CreateApplicationSnapshotFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplication(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) (*kinesisanalyticsv2.DeleteApplicationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInput) *KinesisAnalyticsV2DeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplication", input)
	return &KinesisAnalyticsV2DeleteApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationCloudWatchLoggingOption", input)
	return &KinesisAnalyticsV2DeleteApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationInputProcessingConfigurationInput) *KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationInputProcessingConfiguration", input)
	return &KinesisAnalyticsV2DeleteApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) (*kinesisanalyticsv2.DeleteApplicationOutputOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationOutputInput) *KinesisAnalyticsV2DeleteApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationOutput", input)
	return &KinesisAnalyticsV2DeleteApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) (*kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationReferenceDataSourceInput) *KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationReferenceDataSource", input)
	return &KinesisAnalyticsV2DeleteApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) (*kinesisanalyticsv2.DeleteApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationSnapshotOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationSnapshot", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationSnapshotInput) *KinesisAnalyticsV2DeleteApplicationSnapshotFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationSnapshot", input)
	return &KinesisAnalyticsV2DeleteApplicationSnapshotFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationVpcConfiguration(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) (*kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput, error) {
	var output kinesisanalyticsv2.DeleteApplicationVpcConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationVpcConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DeleteApplicationVpcConfigurationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DeleteApplicationVpcConfigurationInput) *KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DeleteApplicationVpcConfiguration", input)
	return &KinesisAnalyticsV2DeleteApplicationVpcConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DescribeApplication(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) (*kinesisanalyticsv2.DescribeApplicationOutput, error) {
	var output kinesisanalyticsv2.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DescribeApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationInput) *KinesisAnalyticsV2DescribeApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DescribeApplication", input)
	return &KinesisAnalyticsV2DescribeApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DescribeApplicationSnapshot(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) (*kinesisanalyticsv2.DescribeApplicationSnapshotOutput, error) {
	var output kinesisanalyticsv2.DescribeApplicationSnapshotOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DescribeApplicationSnapshot", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DescribeApplicationSnapshotAsync(ctx workflow.Context, input *kinesisanalyticsv2.DescribeApplicationSnapshotInput) *KinesisAnalyticsV2DescribeApplicationSnapshotFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DescribeApplicationSnapshot", input)
	return &KinesisAnalyticsV2DescribeApplicationSnapshotFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) (*kinesisanalyticsv2.DiscoverInputSchemaOutput, error) {
	var output kinesisanalyticsv2.DiscoverInputSchemaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DiscoverInputSchema", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalyticsv2.DiscoverInputSchemaInput) *KinesisAnalyticsV2DiscoverInputSchemaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.DiscoverInputSchema", input)
	return &KinesisAnalyticsV2DiscoverInputSchemaFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) ListApplicationSnapshots(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) (*kinesisanalyticsv2.ListApplicationSnapshotsOutput, error) {
	var output kinesisanalyticsv2.ListApplicationSnapshotsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListApplicationSnapshots", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListApplicationSnapshotsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationSnapshotsInput) *KinesisAnalyticsV2ListApplicationSnapshotsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListApplicationSnapshots", input)
	return &KinesisAnalyticsV2ListApplicationSnapshotsFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) ListApplications(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) (*kinesisanalyticsv2.ListApplicationsOutput, error) {
	var output kinesisanalyticsv2.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListApplicationsInput) *KinesisAnalyticsV2ListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListApplications", input)
	return &KinesisAnalyticsV2ListApplicationsFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) ListTagsForResource(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) (*kinesisanalyticsv2.ListTagsForResourceOutput, error) {
	var output kinesisanalyticsv2.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.ListTagsForResourceInput) *KinesisAnalyticsV2ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.ListTagsForResource", input)
	return &KinesisAnalyticsV2ListTagsForResourceFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) StartApplication(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) (*kinesisanalyticsv2.StartApplicationOutput, error) {
	var output kinesisanalyticsv2.StartApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.StartApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StartApplicationInput) *KinesisAnalyticsV2StartApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.StartApplication", input)
	return &KinesisAnalyticsV2StartApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) StopApplication(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) (*kinesisanalyticsv2.StopApplicationOutput, error) {
	var output kinesisanalyticsv2.StopApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.StopApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.StopApplicationInput) *KinesisAnalyticsV2StopApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.StopApplication", input)
	return &KinesisAnalyticsV2StopApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) TagResource(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) (*kinesisanalyticsv2.TagResourceOutput, error) {
	var output kinesisanalyticsv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) TagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.TagResourceInput) *KinesisAnalyticsV2TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.TagResource", input)
	return &KinesisAnalyticsV2TagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) UntagResource(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) (*kinesisanalyticsv2.UntagResourceOutput, error) {
	var output kinesisanalyticsv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalyticsv2.UntagResourceInput) *KinesisAnalyticsV2UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.UntagResource", input)
	return &KinesisAnalyticsV2UntagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsV2Stub) UpdateApplication(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) (*kinesisanalyticsv2.UpdateApplicationOutput, error) {
	var output kinesisanalyticsv2.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsV2Stub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalyticsv2.UpdateApplicationInput) *KinesisAnalyticsV2UpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalyticsv2.UpdateApplication", input)
	return &KinesisAnalyticsV2UpdateApplicationFuture{Future: future}
}
