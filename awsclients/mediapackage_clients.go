// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.temporal.io/sdk/workflow"
)

type MediaPackageClient interface {
	ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error)
	ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *MediaPackageConfigureLogsFuture

	CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error)
	CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediaPackageCreateChannelFuture

	CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error)
	CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediaPackageCreateHarvestJobFuture

	CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error)
	CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediaPackageCreateOriginEndpointFuture

	DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error)
	DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediaPackageDeleteChannelFuture

	DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error)
	DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediaPackageDeleteOriginEndpointFuture

	DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error)
	DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediaPackageDescribeChannelFuture

	DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error)
	DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediaPackageDescribeHarvestJobFuture

	DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error)
	DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediaPackageDescribeOriginEndpointFuture

	ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error)
	ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediaPackageListChannelsFuture

	ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error)
	ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediaPackageListHarvestJobsFuture

	ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error)
	ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediaPackageListOriginEndpointsFuture

	ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediaPackageListTagsForResourceFuture

	RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error)
	RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediaPackageRotateChannelCredentialsFuture

	RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error)
	RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediaPackageRotateIngestEndpointCredentialsFuture

	TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediaPackageTagResourceFuture

	UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediaPackageUntagResourceFuture

	UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error)
	UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediaPackageUpdateChannelFuture

	UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error)
	UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediaPackageUpdateOriginEndpointFuture
}

type MediaPackageStub struct{}

func NewMediaPackageStub() MediaPackageClient {
	return &MediaPackageStub{}
}

type MediaPackageConfigureLogsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageConfigureLogsFuture) Get(ctx workflow.Context) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageCreateChannelFuture struct {
	Future workflow.Future
}

func (r *MediaPackageCreateChannelFuture) Get(ctx workflow.Context) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageCreateHarvestJobFuture struct {
	Future workflow.Future
}

func (r *MediaPackageCreateHarvestJobFuture) Get(ctx workflow.Context) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageCreateOriginEndpointFuture struct {
	Future workflow.Future
}

func (r *MediaPackageCreateOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageDeleteChannelFuture struct {
	Future workflow.Future
}

func (r *MediaPackageDeleteChannelFuture) Get(ctx workflow.Context) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageDeleteOriginEndpointFuture struct {
	Future workflow.Future
}

func (r *MediaPackageDeleteOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageDescribeChannelFuture struct {
	Future workflow.Future
}

func (r *MediaPackageDescribeChannelFuture) Get(ctx workflow.Context) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageDescribeHarvestJobFuture struct {
	Future workflow.Future
}

func (r *MediaPackageDescribeHarvestJobFuture) Get(ctx workflow.Context) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageDescribeOriginEndpointFuture struct {
	Future workflow.Future
}

func (r *MediaPackageDescribeOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageListChannelsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageListChannelsFuture) Get(ctx workflow.Context) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageListHarvestJobsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageListHarvestJobsFuture) Get(ctx workflow.Context) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageListOriginEndpointsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageListOriginEndpointsFuture) Get(ctx workflow.Context) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *MediaPackageListTagsForResourceFuture) Get(ctx workflow.Context) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageRotateChannelCredentialsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageRotateChannelCredentialsFuture) Get(ctx workflow.Context) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageRotateIngestEndpointCredentialsFuture struct {
	Future workflow.Future
}

func (r *MediaPackageRotateIngestEndpointCredentialsFuture) Get(ctx workflow.Context) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageTagResourceFuture struct {
	Future workflow.Future
}

func (r *MediaPackageTagResourceFuture) Get(ctx workflow.Context) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageUntagResourceFuture struct {
	Future workflow.Future
}

func (r *MediaPackageUntagResourceFuture) Get(ctx workflow.Context) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageUpdateChannelFuture struct {
	Future workflow.Future
}

func (r *MediaPackageUpdateChannelFuture) Get(ctx workflow.Context) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MediaPackageUpdateOriginEndpointFuture struct {
	Future workflow.Future
}

func (r *MediaPackageUpdateOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *MediaPackageConfigureLogsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input)
	return &MediaPackageConfigureLogsFuture{Future: future}
}

func (a *MediaPackageStub) CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *MediaPackageCreateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input)
	return &MediaPackageCreateChannelFuture{Future: future}
}

func (a *MediaPackageStub) CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *MediaPackageCreateHarvestJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input)
	return &MediaPackageCreateHarvestJobFuture{Future: future}
}

func (a *MediaPackageStub) CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *MediaPackageCreateOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input)
	return &MediaPackageCreateOriginEndpointFuture{Future: future}
}

func (a *MediaPackageStub) DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *MediaPackageDeleteChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input)
	return &MediaPackageDeleteChannelFuture{Future: future}
}

func (a *MediaPackageStub) DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *MediaPackageDeleteOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input)
	return &MediaPackageDeleteOriginEndpointFuture{Future: future}
}

func (a *MediaPackageStub) DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *MediaPackageDescribeChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input)
	return &MediaPackageDescribeChannelFuture{Future: future}
}

func (a *MediaPackageStub) DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *MediaPackageDescribeHarvestJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input)
	return &MediaPackageDescribeHarvestJobFuture{Future: future}
}

func (a *MediaPackageStub) DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *MediaPackageDescribeOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input)
	return &MediaPackageDescribeOriginEndpointFuture{Future: future}
}

func (a *MediaPackageStub) ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *MediaPackageListChannelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input)
	return &MediaPackageListChannelsFuture{Future: future}
}

func (a *MediaPackageStub) ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *MediaPackageListHarvestJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input)
	return &MediaPackageListHarvestJobsFuture{Future: future}
}

func (a *MediaPackageStub) ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *MediaPackageListOriginEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input)
	return &MediaPackageListOriginEndpointsFuture{Future: future}
}

func (a *MediaPackageStub) ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *MediaPackageListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input)
	return &MediaPackageListTagsForResourceFuture{Future: future}
}

func (a *MediaPackageStub) RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *MediaPackageRotateChannelCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input)
	return &MediaPackageRotateChannelCredentialsFuture{Future: future}
}

func (a *MediaPackageStub) RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *MediaPackageRotateIngestEndpointCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input)
	return &MediaPackageRotateIngestEndpointCredentialsFuture{Future: future}
}

func (a *MediaPackageStub) TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *MediaPackageTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input)
	return &MediaPackageTagResourceFuture{Future: future}
}

func (a *MediaPackageStub) UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *MediaPackageUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input)
	return &MediaPackageUntagResourceFuture{Future: future}
}

func (a *MediaPackageStub) UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *MediaPackageUpdateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input)
	return &MediaPackageUpdateChannelFuture{Future: future}
}

func (a *MediaPackageStub) UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *MediaPackageStub) UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *MediaPackageUpdateOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input)
	return &MediaPackageUpdateOriginEndpointFuture{Future: future}
}
