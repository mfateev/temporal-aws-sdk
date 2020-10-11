// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
	"go.temporal.io/sdk/workflow"
)

type CloudHSMV2Client interface {
	CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error)
	CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *CloudHSMV2CopyBackupToRegionFuture

	CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *CloudHSMV2CreateClusterFuture

	CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error)
	CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *CloudHSMV2CreateHsmFuture

	DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error)
	DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *CloudHSMV2DeleteBackupFuture

	DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *CloudHSMV2DeleteClusterFuture

	DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error)
	DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *CloudHSMV2DeleteHsmFuture

	DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *CloudHSMV2DescribeBackupsFuture

	DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *CloudHSMV2DescribeClustersFuture

	InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error)
	InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *CloudHSMV2InitializeClusterFuture

	ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *CloudHSMV2ListTagsFuture

	RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error)
	RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *CloudHSMV2RestoreBackupFuture

	TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *CloudHSMV2TagResourceFuture

	UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *CloudHSMV2UntagResourceFuture
}

type CloudHSMV2Stub struct{}

func NewCloudHSMV2Stub() CloudHSMV2Client {
	return &CloudHSMV2Stub{}
}

type CloudHSMV2CopyBackupToRegionFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2CopyBackupToRegionFuture) Get(ctx workflow.Context) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	var output cloudhsmv2.CopyBackupToRegionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2CreateClusterFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2CreateClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.CreateClusterOutput, error) {
	var output cloudhsmv2.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2CreateHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2CreateHsmFuture) Get(ctx workflow.Context) (*cloudhsmv2.CreateHsmOutput, error) {
	var output cloudhsmv2.CreateHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2DeleteBackupFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2DeleteBackupFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteBackupOutput, error) {
	var output cloudhsmv2.DeleteBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2DeleteClusterFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2DeleteClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteClusterOutput, error) {
	var output cloudhsmv2.DeleteClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2DeleteHsmFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2DeleteHsmFuture) Get(ctx workflow.Context) (*cloudhsmv2.DeleteHsmOutput, error) {
	var output cloudhsmv2.DeleteHsmOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2DescribeBackupsFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2DescribeBackupsFuture) Get(ctx workflow.Context) (*cloudhsmv2.DescribeBackupsOutput, error) {
	var output cloudhsmv2.DescribeBackupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2DescribeClustersFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2DescribeClustersFuture) Get(ctx workflow.Context) (*cloudhsmv2.DescribeClustersOutput, error) {
	var output cloudhsmv2.DescribeClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2InitializeClusterFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2InitializeClusterFuture) Get(ctx workflow.Context) (*cloudhsmv2.InitializeClusterOutput, error) {
	var output cloudhsmv2.InitializeClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2ListTagsFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2ListTagsFuture) Get(ctx workflow.Context) (*cloudhsmv2.ListTagsOutput, error) {
	var output cloudhsmv2.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2RestoreBackupFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2RestoreBackupFuture) Get(ctx workflow.Context) (*cloudhsmv2.RestoreBackupOutput, error) {
	var output cloudhsmv2.RestoreBackupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2TagResourceFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2TagResourceFuture) Get(ctx workflow.Context) (*cloudhsmv2.TagResourceOutput, error) {
	var output cloudhsmv2.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudHSMV2UntagResourceFuture struct {
	Future workflow.Future
}

func (r *CloudHSMV2UntagResourceFuture) Get(ctx workflow.Context) (*cloudhsmv2.UntagResourceOutput, error) {
	var output cloudhsmv2.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) CopyBackupToRegion(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
	var output cloudhsmv2.CopyBackupToRegionOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CopyBackupToRegion", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) CopyBackupToRegionAsync(ctx workflow.Context, input *cloudhsmv2.CopyBackupToRegionInput) *CloudHSMV2CopyBackupToRegionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CopyBackupToRegion", input)
	return &CloudHSMV2CopyBackupToRegionFuture{Future: future}
}

func (a *CloudHSMV2Stub) CreateCluster(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error) {
	var output cloudhsmv2.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) CreateClusterAsync(ctx workflow.Context, input *cloudhsmv2.CreateClusterInput) *CloudHSMV2CreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateCluster", input)
	return &CloudHSMV2CreateClusterFuture{Future: future}
}

func (a *CloudHSMV2Stub) CreateHsm(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error) {
	var output cloudhsmv2.CreateHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) CreateHsmAsync(ctx workflow.Context, input *cloudhsmv2.CreateHsmInput) *CloudHSMV2CreateHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.CreateHsm", input)
	return &CloudHSMV2CreateHsmFuture{Future: future}
}

func (a *CloudHSMV2Stub) DeleteBackup(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error) {
	var output cloudhsmv2.DeleteBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) DeleteBackupAsync(ctx workflow.Context, input *cloudhsmv2.DeleteBackupInput) *CloudHSMV2DeleteBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteBackup", input)
	return &CloudHSMV2DeleteBackupFuture{Future: future}
}

func (a *CloudHSMV2Stub) DeleteCluster(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error) {
	var output cloudhsmv2.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) DeleteClusterAsync(ctx workflow.Context, input *cloudhsmv2.DeleteClusterInput) *CloudHSMV2DeleteClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteCluster", input)
	return &CloudHSMV2DeleteClusterFuture{Future: future}
}

func (a *CloudHSMV2Stub) DeleteHsm(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error) {
	var output cloudhsmv2.DeleteHsmOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteHsm", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) DeleteHsmAsync(ctx workflow.Context, input *cloudhsmv2.DeleteHsmInput) *CloudHSMV2DeleteHsmFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DeleteHsm", input)
	return &CloudHSMV2DeleteHsmFuture{Future: future}
}

func (a *CloudHSMV2Stub) DescribeBackups(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error) {
	var output cloudhsmv2.DescribeBackupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeBackups", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) DescribeBackupsAsync(ctx workflow.Context, input *cloudhsmv2.DescribeBackupsInput) *CloudHSMV2DescribeBackupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeBackups", input)
	return &CloudHSMV2DescribeBackupsFuture{Future: future}
}

func (a *CloudHSMV2Stub) DescribeClusters(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error) {
	var output cloudhsmv2.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) DescribeClustersAsync(ctx workflow.Context, input *cloudhsmv2.DescribeClustersInput) *CloudHSMV2DescribeClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.DescribeClusters", input)
	return &CloudHSMV2DescribeClustersFuture{Future: future}
}

func (a *CloudHSMV2Stub) InitializeCluster(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error) {
	var output cloudhsmv2.InitializeClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.InitializeCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) InitializeClusterAsync(ctx workflow.Context, input *cloudhsmv2.InitializeClusterInput) *CloudHSMV2InitializeClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.InitializeCluster", input)
	return &CloudHSMV2InitializeClusterFuture{Future: future}
}

func (a *CloudHSMV2Stub) ListTags(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error) {
	var output cloudhsmv2.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) ListTagsAsync(ctx workflow.Context, input *cloudhsmv2.ListTagsInput) *CloudHSMV2ListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.ListTags", input)
	return &CloudHSMV2ListTagsFuture{Future: future}
}

func (a *CloudHSMV2Stub) RestoreBackup(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error) {
	var output cloudhsmv2.RestoreBackupOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.RestoreBackup", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) RestoreBackupAsync(ctx workflow.Context, input *cloudhsmv2.RestoreBackupInput) *CloudHSMV2RestoreBackupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.RestoreBackup", input)
	return &CloudHSMV2RestoreBackupFuture{Future: future}
}

func (a *CloudHSMV2Stub) TagResource(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error) {
	var output cloudhsmv2.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) TagResourceAsync(ctx workflow.Context, input *cloudhsmv2.TagResourceInput) *CloudHSMV2TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.TagResource", input)
	return &CloudHSMV2TagResourceFuture{Future: future}
}

func (a *CloudHSMV2Stub) UntagResource(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error) {
	var output cloudhsmv2.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudHSMV2Stub) UntagResourceAsync(ctx workflow.Context, input *cloudhsmv2.UntagResourceInput) *CloudHSMV2UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudhsmv2.UntagResource", input)
	return &CloudHSMV2UntagResourceFuture{Future: future}
}
