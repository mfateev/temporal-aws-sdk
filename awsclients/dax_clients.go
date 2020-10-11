// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/dax"
	"go.temporal.io/sdk/workflow"
)

type DAXClient interface {
	CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *DAXCreateClusterFuture

	CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error)
	CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *DAXCreateParameterGroupFuture

	CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error)
	CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *DAXCreateSubnetGroupFuture

	DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error)
	DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DAXDecreaseReplicationFactorFuture

	DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DAXDeleteClusterFuture

	DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error)
	DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DAXDeleteParameterGroupFuture

	DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error)
	DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DAXDeleteSubnetGroupFuture

	DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DAXDescribeClustersFuture

	DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error)
	DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DAXDescribeDefaultParametersFuture

	DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error)
	DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DAXDescribeEventsFuture

	DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error)
	DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DAXDescribeParameterGroupsFuture

	DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error)
	DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DAXDescribeParametersFuture

	DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error)
	DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DAXDescribeSubnetGroupsFuture

	IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error)
	IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *DAXIncreaseReplicationFactorFuture

	ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *DAXListTagsFuture

	RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error)
	RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *DAXRebootNodeFuture

	TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *DAXTagResourceFuture

	UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *DAXUntagResourceFuture

	UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error)
	UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *DAXUpdateClusterFuture

	UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error)
	UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *DAXUpdateParameterGroupFuture

	UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error)
	UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *DAXUpdateSubnetGroupFuture
}

type DAXStub struct{}

func NewDAXStub() DAXClient {
	return &DAXStub{}
}

type DAXCreateClusterFuture struct {
	Future workflow.Future
}

func (r *DAXCreateClusterFuture) Get(ctx workflow.Context) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXCreateParameterGroupFuture struct {
	Future workflow.Future
}

func (r *DAXCreateParameterGroupFuture) Get(ctx workflow.Context) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXCreateSubnetGroupFuture struct {
	Future workflow.Future
}

func (r *DAXCreateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDecreaseReplicationFactorFuture struct {
	Future workflow.Future
}

func (r *DAXDecreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteClusterFuture struct {
	Future workflow.Future
}

func (r *DAXDeleteClusterFuture) Get(ctx workflow.Context) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteParameterGroupFuture struct {
	Future workflow.Future
}

func (r *DAXDeleteParameterGroupFuture) Get(ctx workflow.Context) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDeleteSubnetGroupFuture struct {
	Future workflow.Future
}

func (r *DAXDeleteSubnetGroupFuture) Get(ctx workflow.Context) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeClustersFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeClustersFuture) Get(ctx workflow.Context) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeDefaultParametersFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeDefaultParametersFuture) Get(ctx workflow.Context) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeEventsFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeEventsFuture) Get(ctx workflow.Context) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeParameterGroupsFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeParameterGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeParametersFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeParametersFuture) Get(ctx workflow.Context) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXDescribeSubnetGroupsFuture struct {
	Future workflow.Future
}

func (r *DAXDescribeSubnetGroupsFuture) Get(ctx workflow.Context) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXIncreaseReplicationFactorFuture struct {
	Future workflow.Future
}

func (r *DAXIncreaseReplicationFactorFuture) Get(ctx workflow.Context) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXListTagsFuture struct {
	Future workflow.Future
}

func (r *DAXListTagsFuture) Get(ctx workflow.Context) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXRebootNodeFuture struct {
	Future workflow.Future
}

func (r *DAXRebootNodeFuture) Get(ctx workflow.Context) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXTagResourceFuture struct {
	Future workflow.Future
}

func (r *DAXTagResourceFuture) Get(ctx workflow.Context) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUntagResourceFuture struct {
	Future workflow.Future
}

func (r *DAXUntagResourceFuture) Get(ctx workflow.Context) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateClusterFuture struct {
	Future workflow.Future
}

func (r *DAXUpdateClusterFuture) Get(ctx workflow.Context) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateParameterGroupFuture struct {
	Future workflow.Future
}

func (r *DAXUpdateParameterGroupFuture) Get(ctx workflow.Context) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DAXUpdateSubnetGroupFuture struct {
	Future workflow.Future
}

func (r *DAXUpdateSubnetGroupFuture) Get(ctx workflow.Context) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateCluster(ctx workflow.Context, input *dax.CreateClusterInput) (*dax.CreateClusterOutput, error) {
	var output dax.CreateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateClusterAsync(ctx workflow.Context, input *dax.CreateClusterInput) *DAXCreateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateCluster", input)
	return &DAXCreateClusterFuture{Future: future}
}

func (a *DAXStub) CreateParameterGroup(ctx workflow.Context, input *dax.CreateParameterGroupInput) (*dax.CreateParameterGroupOutput, error) {
	var output dax.CreateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateParameterGroupAsync(ctx workflow.Context, input *dax.CreateParameterGroupInput) *DAXCreateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateParameterGroup", input)
	return &DAXCreateParameterGroupFuture{Future: future}
}

func (a *DAXStub) CreateSubnetGroup(ctx workflow.Context, input *dax.CreateSubnetGroupInput) (*dax.CreateSubnetGroupOutput, error) {
	var output dax.CreateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) CreateSubnetGroupAsync(ctx workflow.Context, input *dax.CreateSubnetGroupInput) *DAXCreateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.CreateSubnetGroup", input)
	return &DAXCreateSubnetGroupFuture{Future: future}
}

func (a *DAXStub) DecreaseReplicationFactor(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) (*dax.DecreaseReplicationFactorOutput, error) {
	var output dax.DecreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DecreaseReplicationFactorAsync(ctx workflow.Context, input *dax.DecreaseReplicationFactorInput) *DAXDecreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DecreaseReplicationFactor", input)
	return &DAXDecreaseReplicationFactorFuture{Future: future}
}

func (a *DAXStub) DeleteCluster(ctx workflow.Context, input *dax.DeleteClusterInput) (*dax.DeleteClusterOutput, error) {
	var output dax.DeleteClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteClusterAsync(ctx workflow.Context, input *dax.DeleteClusterInput) *DAXDeleteClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteCluster", input)
	return &DAXDeleteClusterFuture{Future: future}
}

func (a *DAXStub) DeleteParameterGroup(ctx workflow.Context, input *dax.DeleteParameterGroupInput) (*dax.DeleteParameterGroupOutput, error) {
	var output dax.DeleteParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteParameterGroupAsync(ctx workflow.Context, input *dax.DeleteParameterGroupInput) *DAXDeleteParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteParameterGroup", input)
	return &DAXDeleteParameterGroupFuture{Future: future}
}

func (a *DAXStub) DeleteSubnetGroup(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) (*dax.DeleteSubnetGroupOutput, error) {
	var output dax.DeleteSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DeleteSubnetGroupAsync(ctx workflow.Context, input *dax.DeleteSubnetGroupInput) *DAXDeleteSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DeleteSubnetGroup", input)
	return &DAXDeleteSubnetGroupFuture{Future: future}
}

func (a *DAXStub) DescribeClusters(ctx workflow.Context, input *dax.DescribeClustersInput) (*dax.DescribeClustersOutput, error) {
	var output dax.DescribeClustersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeClustersAsync(ctx workflow.Context, input *dax.DescribeClustersInput) *DAXDescribeClustersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeClusters", input)
	return &DAXDescribeClustersFuture{Future: future}
}

func (a *DAXStub) DescribeDefaultParameters(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) (*dax.DescribeDefaultParametersOutput, error) {
	var output dax.DescribeDefaultParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeDefaultParametersAsync(ctx workflow.Context, input *dax.DescribeDefaultParametersInput) *DAXDescribeDefaultParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeDefaultParameters", input)
	return &DAXDescribeDefaultParametersFuture{Future: future}
}

func (a *DAXStub) DescribeEvents(ctx workflow.Context, input *dax.DescribeEventsInput) (*dax.DescribeEventsOutput, error) {
	var output dax.DescribeEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeEventsAsync(ctx workflow.Context, input *dax.DescribeEventsInput) *DAXDescribeEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeEvents", input)
	return &DAXDescribeEventsFuture{Future: future}
}

func (a *DAXStub) DescribeParameterGroups(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) (*dax.DescribeParameterGroupsOutput, error) {
	var output dax.DescribeParameterGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeParameterGroupsAsync(ctx workflow.Context, input *dax.DescribeParameterGroupsInput) *DAXDescribeParameterGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameterGroups", input)
	return &DAXDescribeParameterGroupsFuture{Future: future}
}

func (a *DAXStub) DescribeParameters(ctx workflow.Context, input *dax.DescribeParametersInput) (*dax.DescribeParametersOutput, error) {
	var output dax.DescribeParametersOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeParametersAsync(ctx workflow.Context, input *dax.DescribeParametersInput) *DAXDescribeParametersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeParameters", input)
	return &DAXDescribeParametersFuture{Future: future}
}

func (a *DAXStub) DescribeSubnetGroups(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) (*dax.DescribeSubnetGroupsOutput, error) {
	var output dax.DescribeSubnetGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) DescribeSubnetGroupsAsync(ctx workflow.Context, input *dax.DescribeSubnetGroupsInput) *DAXDescribeSubnetGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.DescribeSubnetGroups", input)
	return &DAXDescribeSubnetGroupsFuture{Future: future}
}

func (a *DAXStub) IncreaseReplicationFactor(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) (*dax.IncreaseReplicationFactorOutput, error) {
	var output dax.IncreaseReplicationFactorOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) IncreaseReplicationFactorAsync(ctx workflow.Context, input *dax.IncreaseReplicationFactorInput) *DAXIncreaseReplicationFactorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.IncreaseReplicationFactor", input)
	return &DAXIncreaseReplicationFactorFuture{Future: future}
}

func (a *DAXStub) ListTags(ctx workflow.Context, input *dax.ListTagsInput) (*dax.ListTagsOutput, error) {
	var output dax.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) ListTagsAsync(ctx workflow.Context, input *dax.ListTagsInput) *DAXListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.ListTags", input)
	return &DAXListTagsFuture{Future: future}
}

func (a *DAXStub) RebootNode(ctx workflow.Context, input *dax.RebootNodeInput) (*dax.RebootNodeOutput, error) {
	var output dax.RebootNodeOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) RebootNodeAsync(ctx workflow.Context, input *dax.RebootNodeInput) *DAXRebootNodeFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.RebootNode", input)
	return &DAXRebootNodeFuture{Future: future}
}

func (a *DAXStub) TagResource(ctx workflow.Context, input *dax.TagResourceInput) (*dax.TagResourceOutput, error) {
	var output dax.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) TagResourceAsync(ctx workflow.Context, input *dax.TagResourceInput) *DAXTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.TagResource", input)
	return &DAXTagResourceFuture{Future: future}
}

func (a *DAXStub) UntagResource(ctx workflow.Context, input *dax.UntagResourceInput) (*dax.UntagResourceOutput, error) {
	var output dax.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UntagResourceAsync(ctx workflow.Context, input *dax.UntagResourceInput) *DAXUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UntagResource", input)
	return &DAXUntagResourceFuture{Future: future}
}

func (a *DAXStub) UpdateCluster(ctx workflow.Context, input *dax.UpdateClusterInput) (*dax.UpdateClusterOutput, error) {
	var output dax.UpdateClusterOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateClusterAsync(ctx workflow.Context, input *dax.UpdateClusterInput) *DAXUpdateClusterFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateCluster", input)
	return &DAXUpdateClusterFuture{Future: future}
}

func (a *DAXStub) UpdateParameterGroup(ctx workflow.Context, input *dax.UpdateParameterGroupInput) (*dax.UpdateParameterGroupOutput, error) {
	var output dax.UpdateParameterGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateParameterGroupAsync(ctx workflow.Context, input *dax.UpdateParameterGroupInput) *DAXUpdateParameterGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateParameterGroup", input)
	return &DAXUpdateParameterGroupFuture{Future: future}
}

func (a *DAXStub) UpdateSubnetGroup(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) (*dax.UpdateSubnetGroupOutput, error) {
	var output dax.UpdateSubnetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *DAXStub) UpdateSubnetGroupAsync(ctx workflow.Context, input *dax.UpdateSubnetGroupInput) *DAXUpdateSubnetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.dax.UpdateSubnetGroup", input)
	return &DAXUpdateSubnetGroupFuture{Future: future}
}
