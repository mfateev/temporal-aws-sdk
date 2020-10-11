// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/kinesisanalytics"
	"go.temporal.io/sdk/workflow"
)

type KinesisAnalyticsClient interface {
	AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error)
	AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsAddApplicationCloudWatchLoggingOptionFuture

	AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error)
	AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisAnalyticsAddApplicationInputFuture

	AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error)
	AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisAnalyticsAddApplicationInputProcessingConfigurationFuture

	AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error)
	AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisAnalyticsAddApplicationOutputFuture

	AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error)
	AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisAnalyticsAddApplicationReferenceDataSourceFuture

	CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error)
	CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisAnalyticsCreateApplicationFuture

	DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error)
	DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisAnalyticsDeleteApplicationFuture

	DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error)
	DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsDeleteApplicationCloudWatchLoggingOptionFuture

	DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error)
	DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisAnalyticsDeleteApplicationInputProcessingConfigurationFuture

	DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error)
	DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisAnalyticsDeleteApplicationOutputFuture

	DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error)
	DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisAnalyticsDeleteApplicationReferenceDataSourceFuture

	DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error)
	DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisAnalyticsDescribeApplicationFuture

	DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error)
	DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisAnalyticsDiscoverInputSchemaFuture

	ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error)
	ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisAnalyticsListApplicationsFuture

	ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisAnalyticsListTagsForResourceFuture

	StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error)
	StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisAnalyticsStartApplicationFuture

	StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error)
	StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisAnalyticsStopApplicationFuture

	TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisAnalyticsTagResourceFuture

	UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisAnalyticsUntagResourceFuture

	UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error)
	UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisAnalyticsUpdateApplicationFuture
}

type KinesisAnalyticsStub struct{}

func NewKinesisAnalyticsStub() KinesisAnalyticsClient {
	return &KinesisAnalyticsStub{}
}

type KinesisAnalyticsAddApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsAddApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsAddApplicationInputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsAddApplicationInputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsAddApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsAddApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsAddApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsAddApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsAddApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsAddApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsCreateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsCreateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDeleteApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDeleteApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDeleteApplicationCloudWatchLoggingOptionFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDeleteApplicationCloudWatchLoggingOptionFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDeleteApplicationInputProcessingConfigurationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDeleteApplicationInputProcessingConfigurationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDeleteApplicationOutputFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDeleteApplicationOutputFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDeleteApplicationReferenceDataSourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDeleteApplicationReferenceDataSourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDescribeApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDescribeApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsDiscoverInputSchemaFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsDiscoverInputSchemaFuture) Get(ctx workflow.Context) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsListApplicationsFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsListApplicationsFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsListTagsForResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsStartApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsStartApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsStopApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsStopApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsTagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsTagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsUntagResourceFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsUntagResourceFuture) Get(ctx workflow.Context) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type KinesisAnalyticsUpdateApplicationFuture struct {
	Future workflow.Future
}

func (r *KinesisAnalyticsUpdateApplicationFuture) Get(ctx workflow.Context) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.AddApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsAddApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationCloudWatchLoggingOption", input)
	return &KinesisAnalyticsAddApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInput(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) (*kinesisanalytics.AddApplicationInputOutput, error) {
	var output kinesisanalytics.AddApplicationInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputInput) *KinesisAnalyticsAddApplicationInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInput", input)
	return &KinesisAnalyticsAddApplicationInputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) (*kinesisanalytics.AddApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.AddApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationInputProcessingConfigurationInput) *KinesisAnalyticsAddApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationInputProcessingConfiguration", input)
	return &KinesisAnalyticsAddApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationOutput(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) (*kinesisanalytics.AddApplicationOutputOutput, error) {
	var output kinesisanalytics.AddApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationOutputInput) *KinesisAnalyticsAddApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationOutput", input)
	return &KinesisAnalyticsAddApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) (*kinesisanalytics.AddApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.AddApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) AddApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.AddApplicationReferenceDataSourceInput) *KinesisAnalyticsAddApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.AddApplicationReferenceDataSource", input)
	return &KinesisAnalyticsAddApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) CreateApplication(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) (*kinesisanalytics.CreateApplicationOutput, error) {
	var output kinesisanalytics.CreateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) CreateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.CreateApplicationInput) *KinesisAnalyticsCreateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.CreateApplication", input)
	return &KinesisAnalyticsCreateApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplication(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) (*kinesisanalytics.DeleteApplicationOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInput) *KinesisAnalyticsDeleteApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplication", input)
	return &KinesisAnalyticsDeleteApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOption(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) (*kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput, error) {
	var output kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationCloudWatchLoggingOptionAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationCloudWatchLoggingOptionInput) *KinesisAnalyticsDeleteApplicationCloudWatchLoggingOptionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationCloudWatchLoggingOption", input)
	return &KinesisAnalyticsDeleteApplicationCloudWatchLoggingOptionFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfiguration(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) (*kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput, error) {
	var output kinesisanalytics.DeleteApplicationInputProcessingConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationInputProcessingConfigurationAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationInputProcessingConfigurationInput) *KinesisAnalyticsDeleteApplicationInputProcessingConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationInputProcessingConfiguration", input)
	return &KinesisAnalyticsDeleteApplicationInputProcessingConfigurationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutput(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) (*kinesisanalytics.DeleteApplicationOutputOutput, error) {
	var output kinesisanalytics.DeleteApplicationOutputOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationOutputAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationOutputInput) *KinesisAnalyticsDeleteApplicationOutputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationOutput", input)
	return &KinesisAnalyticsDeleteApplicationOutputFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSource(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) (*kinesisanalytics.DeleteApplicationReferenceDataSourceOutput, error) {
	var output kinesisanalytics.DeleteApplicationReferenceDataSourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DeleteApplicationReferenceDataSourceAsync(ctx workflow.Context, input *kinesisanalytics.DeleteApplicationReferenceDataSourceInput) *KinesisAnalyticsDeleteApplicationReferenceDataSourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DeleteApplicationReferenceDataSource", input)
	return &KinesisAnalyticsDeleteApplicationReferenceDataSourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DescribeApplication(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) (*kinesisanalytics.DescribeApplicationOutput, error) {
	var output kinesisanalytics.DescribeApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DescribeApplicationAsync(ctx workflow.Context, input *kinesisanalytics.DescribeApplicationInput) *KinesisAnalyticsDescribeApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DescribeApplication", input)
	return &KinesisAnalyticsDescribeApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) DiscoverInputSchema(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) (*kinesisanalytics.DiscoverInputSchemaOutput, error) {
	var output kinesisanalytics.DiscoverInputSchemaOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) DiscoverInputSchemaAsync(ctx workflow.Context, input *kinesisanalytics.DiscoverInputSchemaInput) *KinesisAnalyticsDiscoverInputSchemaFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.DiscoverInputSchema", input)
	return &KinesisAnalyticsDiscoverInputSchemaFuture{Future: future}
}

func (a *KinesisAnalyticsStub) ListApplications(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) (*kinesisanalytics.ListApplicationsOutput, error) {
	var output kinesisanalytics.ListApplicationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListApplicationsAsync(ctx workflow.Context, input *kinesisanalytics.ListApplicationsInput) *KinesisAnalyticsListApplicationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListApplications", input)
	return &KinesisAnalyticsListApplicationsFuture{Future: future}
}

func (a *KinesisAnalyticsStub) ListTagsForResource(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) (*kinesisanalytics.ListTagsForResourceOutput, error) {
	var output kinesisanalytics.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) ListTagsForResourceAsync(ctx workflow.Context, input *kinesisanalytics.ListTagsForResourceInput) *KinesisAnalyticsListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.ListTagsForResource", input)
	return &KinesisAnalyticsListTagsForResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) StartApplication(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) (*kinesisanalytics.StartApplicationOutput, error) {
	var output kinesisanalytics.StartApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StartApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StartApplicationInput) *KinesisAnalyticsStartApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StartApplication", input)
	return &KinesisAnalyticsStartApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) StopApplication(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) (*kinesisanalytics.StopApplicationOutput, error) {
	var output kinesisanalytics.StopApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) StopApplicationAsync(ctx workflow.Context, input *kinesisanalytics.StopApplicationInput) *KinesisAnalyticsStopApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.StopApplication", input)
	return &KinesisAnalyticsStopApplicationFuture{Future: future}
}

func (a *KinesisAnalyticsStub) TagResource(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) (*kinesisanalytics.TagResourceOutput, error) {
	var output kinesisanalytics.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) TagResourceAsync(ctx workflow.Context, input *kinesisanalytics.TagResourceInput) *KinesisAnalyticsTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.TagResource", input)
	return &KinesisAnalyticsTagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) UntagResource(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) (*kinesisanalytics.UntagResourceOutput, error) {
	var output kinesisanalytics.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UntagResourceAsync(ctx workflow.Context, input *kinesisanalytics.UntagResourceInput) *KinesisAnalyticsUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UntagResource", input)
	return &KinesisAnalyticsUntagResourceFuture{Future: future}
}

func (a *KinesisAnalyticsStub) UpdateApplication(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) (*kinesisanalytics.UpdateApplicationOutput, error) {
	var output kinesisanalytics.UpdateApplicationOutput
	err := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input).Get(ctx, &output)
	return &output, err
}

func (a *KinesisAnalyticsStub) UpdateApplicationAsync(ctx workflow.Context, input *kinesisanalytics.UpdateApplicationInput) *KinesisAnalyticsUpdateApplicationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.kinesisanalytics.UpdateApplication", input)
	return &KinesisAnalyticsUpdateApplicationFuture{Future: future}
}
