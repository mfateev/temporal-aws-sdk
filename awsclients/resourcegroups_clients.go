// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/resourcegroups"
	"go.temporal.io/sdk/workflow"
)

type ResourceGroupsClient interface {
	CreateGroup(ctx workflow.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error)
	CreateGroupAsync(ctx workflow.Context, input *resourcegroups.CreateGroupInput) *ResourceGroupsCreateGroupFuture

	DeleteGroup(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error)
	DeleteGroupAsync(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) *ResourceGroupsDeleteGroupFuture

	GetGroup(ctx workflow.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error)
	GetGroupAsync(ctx workflow.Context, input *resourcegroups.GetGroupInput) *ResourceGroupsGetGroupFuture

	GetGroupConfiguration(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error)
	GetGroupConfigurationAsync(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) *ResourceGroupsGetGroupConfigurationFuture

	GetGroupQuery(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error)
	GetGroupQueryAsync(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) *ResourceGroupsGetGroupQueryFuture

	GetTags(ctx workflow.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error)
	GetTagsAsync(ctx workflow.Context, input *resourcegroups.GetTagsInput) *ResourceGroupsGetTagsFuture

	GroupResources(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error)
	GroupResourcesAsync(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) *ResourceGroupsGroupResourcesFuture

	ListGroupResources(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error)
	ListGroupResourcesAsync(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) *ResourceGroupsListGroupResourcesFuture

	ListGroups(ctx workflow.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error)
	ListGroupsAsync(ctx workflow.Context, input *resourcegroups.ListGroupsInput) *ResourceGroupsListGroupsFuture

	SearchResources(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error)
	SearchResourcesAsync(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) *ResourceGroupsSearchResourcesFuture

	Tag(ctx workflow.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error)
	TagAsync(ctx workflow.Context, input *resourcegroups.TagInput) *ResourceGroupsTagFuture

	UngroupResources(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error)
	UngroupResourcesAsync(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) *ResourceGroupsUngroupResourcesFuture

	Untag(ctx workflow.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error)
	UntagAsync(ctx workflow.Context, input *resourcegroups.UntagInput) *ResourceGroupsUntagFuture

	UpdateGroup(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error)
	UpdateGroupAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) *ResourceGroupsUpdateGroupFuture

	UpdateGroupQuery(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error)
	UpdateGroupQueryAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) *ResourceGroupsUpdateGroupQueryFuture
}

type ResourceGroupsStub struct{}

func NewResourceGroupsStub() ResourceGroupsClient {
	return &ResourceGroupsStub{}
}

type ResourceGroupsCreateGroupFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsCreateGroupFuture) Get(ctx workflow.Context) (*resourcegroups.CreateGroupOutput, error) {
	var output resourcegroups.CreateGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsDeleteGroupFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsDeleteGroupFuture) Get(ctx workflow.Context) (*resourcegroups.DeleteGroupOutput, error) {
	var output resourcegroups.DeleteGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsGetGroupFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsGetGroupFuture) Get(ctx workflow.Context) (*resourcegroups.GetGroupOutput, error) {
	var output resourcegroups.GetGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsGetGroupConfigurationFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsGetGroupConfigurationFuture) Get(ctx workflow.Context) (*resourcegroups.GetGroupConfigurationOutput, error) {
	var output resourcegroups.GetGroupConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsGetGroupQueryFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsGetGroupQueryFuture) Get(ctx workflow.Context) (*resourcegroups.GetGroupQueryOutput, error) {
	var output resourcegroups.GetGroupQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsGetTagsFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsGetTagsFuture) Get(ctx workflow.Context) (*resourcegroups.GetTagsOutput, error) {
	var output resourcegroups.GetTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsGroupResourcesFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsGroupResourcesFuture) Get(ctx workflow.Context) (*resourcegroups.GroupResourcesOutput, error) {
	var output resourcegroups.GroupResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsListGroupResourcesFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsListGroupResourcesFuture) Get(ctx workflow.Context) (*resourcegroups.ListGroupResourcesOutput, error) {
	var output resourcegroups.ListGroupResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsListGroupsFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsListGroupsFuture) Get(ctx workflow.Context) (*resourcegroups.ListGroupsOutput, error) {
	var output resourcegroups.ListGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsSearchResourcesFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsSearchResourcesFuture) Get(ctx workflow.Context) (*resourcegroups.SearchResourcesOutput, error) {
	var output resourcegroups.SearchResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsTagFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsTagFuture) Get(ctx workflow.Context) (*resourcegroups.TagOutput, error) {
	var output resourcegroups.TagOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsUngroupResourcesFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsUngroupResourcesFuture) Get(ctx workflow.Context) (*resourcegroups.UngroupResourcesOutput, error) {
	var output resourcegroups.UngroupResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsUntagFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsUntagFuture) Get(ctx workflow.Context) (*resourcegroups.UntagOutput, error) {
	var output resourcegroups.UntagOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsUpdateGroupFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsUpdateGroupFuture) Get(ctx workflow.Context) (*resourcegroups.UpdateGroupOutput, error) {
	var output resourcegroups.UpdateGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ResourceGroupsUpdateGroupQueryFuture struct {
	Future workflow.Future
}

func (r *ResourceGroupsUpdateGroupQueryFuture) Get(ctx workflow.Context) (*resourcegroups.UpdateGroupQueryOutput, error) {
	var output resourcegroups.UpdateGroupQueryOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) CreateGroup(ctx workflow.Context, input *resourcegroups.CreateGroupInput) (*resourcegroups.CreateGroupOutput, error) {
	var output resourcegroups.CreateGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.CreateGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) CreateGroupAsync(ctx workflow.Context, input *resourcegroups.CreateGroupInput) *ResourceGroupsCreateGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.CreateGroup", input)
	return &ResourceGroupsCreateGroupFuture{Future: future}
}

func (a *ResourceGroupsStub) DeleteGroup(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) (*resourcegroups.DeleteGroupOutput, error) {
	var output resourcegroups.DeleteGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.DeleteGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) DeleteGroupAsync(ctx workflow.Context, input *resourcegroups.DeleteGroupInput) *ResourceGroupsDeleteGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.DeleteGroup", input)
	return &ResourceGroupsDeleteGroupFuture{Future: future}
}

func (a *ResourceGroupsStub) GetGroup(ctx workflow.Context, input *resourcegroups.GetGroupInput) (*resourcegroups.GetGroupOutput, error) {
	var output resourcegroups.GetGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) GetGroupAsync(ctx workflow.Context, input *resourcegroups.GetGroupInput) *ResourceGroupsGetGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroup", input)
	return &ResourceGroupsGetGroupFuture{Future: future}
}

func (a *ResourceGroupsStub) GetGroupConfiguration(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) (*resourcegroups.GetGroupConfigurationOutput, error) {
	var output resourcegroups.GetGroupConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroupConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) GetGroupConfigurationAsync(ctx workflow.Context, input *resourcegroups.GetGroupConfigurationInput) *ResourceGroupsGetGroupConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroupConfiguration", input)
	return &ResourceGroupsGetGroupConfigurationFuture{Future: future}
}

func (a *ResourceGroupsStub) GetGroupQuery(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) (*resourcegroups.GetGroupQueryOutput, error) {
	var output resourcegroups.GetGroupQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroupQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) GetGroupQueryAsync(ctx workflow.Context, input *resourcegroups.GetGroupQueryInput) *ResourceGroupsGetGroupQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetGroupQuery", input)
	return &ResourceGroupsGetGroupQueryFuture{Future: future}
}

func (a *ResourceGroupsStub) GetTags(ctx workflow.Context, input *resourcegroups.GetTagsInput) (*resourcegroups.GetTagsOutput, error) {
	var output resourcegroups.GetTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetTags", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) GetTagsAsync(ctx workflow.Context, input *resourcegroups.GetTagsInput) *ResourceGroupsGetTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GetTags", input)
	return &ResourceGroupsGetTagsFuture{Future: future}
}

func (a *ResourceGroupsStub) GroupResources(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) (*resourcegroups.GroupResourcesOutput, error) {
	var output resourcegroups.GroupResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GroupResources", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) GroupResourcesAsync(ctx workflow.Context, input *resourcegroups.GroupResourcesInput) *ResourceGroupsGroupResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.GroupResources", input)
	return &ResourceGroupsGroupResourcesFuture{Future: future}
}

func (a *ResourceGroupsStub) ListGroupResources(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) (*resourcegroups.ListGroupResourcesOutput, error) {
	var output resourcegroups.ListGroupResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.ListGroupResources", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) ListGroupResourcesAsync(ctx workflow.Context, input *resourcegroups.ListGroupResourcesInput) *ResourceGroupsListGroupResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.ListGroupResources", input)
	return &ResourceGroupsListGroupResourcesFuture{Future: future}
}

func (a *ResourceGroupsStub) ListGroups(ctx workflow.Context, input *resourcegroups.ListGroupsInput) (*resourcegroups.ListGroupsOutput, error) {
	var output resourcegroups.ListGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.ListGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) ListGroupsAsync(ctx workflow.Context, input *resourcegroups.ListGroupsInput) *ResourceGroupsListGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.ListGroups", input)
	return &ResourceGroupsListGroupsFuture{Future: future}
}

func (a *ResourceGroupsStub) SearchResources(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) (*resourcegroups.SearchResourcesOutput, error) {
	var output resourcegroups.SearchResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.SearchResources", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) SearchResourcesAsync(ctx workflow.Context, input *resourcegroups.SearchResourcesInput) *ResourceGroupsSearchResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.SearchResources", input)
	return &ResourceGroupsSearchResourcesFuture{Future: future}
}

func (a *ResourceGroupsStub) Tag(ctx workflow.Context, input *resourcegroups.TagInput) (*resourcegroups.TagOutput, error) {
	var output resourcegroups.TagOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.Tag", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) TagAsync(ctx workflow.Context, input *resourcegroups.TagInput) *ResourceGroupsTagFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.Tag", input)
	return &ResourceGroupsTagFuture{Future: future}
}

func (a *ResourceGroupsStub) UngroupResources(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) (*resourcegroups.UngroupResourcesOutput, error) {
	var output resourcegroups.UngroupResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UngroupResources", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) UngroupResourcesAsync(ctx workflow.Context, input *resourcegroups.UngroupResourcesInput) *ResourceGroupsUngroupResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UngroupResources", input)
	return &ResourceGroupsUngroupResourcesFuture{Future: future}
}

func (a *ResourceGroupsStub) Untag(ctx workflow.Context, input *resourcegroups.UntagInput) (*resourcegroups.UntagOutput, error) {
	var output resourcegroups.UntagOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.Untag", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) UntagAsync(ctx workflow.Context, input *resourcegroups.UntagInput) *ResourceGroupsUntagFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.Untag", input)
	return &ResourceGroupsUntagFuture{Future: future}
}

func (a *ResourceGroupsStub) UpdateGroup(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) (*resourcegroups.UpdateGroupOutput, error) {
	var output resourcegroups.UpdateGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UpdateGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) UpdateGroupAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupInput) *ResourceGroupsUpdateGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UpdateGroup", input)
	return &ResourceGroupsUpdateGroupFuture{Future: future}
}

func (a *ResourceGroupsStub) UpdateGroupQuery(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) (*resourcegroups.UpdateGroupQueryOutput, error) {
	var output resourcegroups.UpdateGroupQueryOutput
	err := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UpdateGroupQuery", input).Get(ctx, &output)
	return &output, err
}

func (a *ResourceGroupsStub) UpdateGroupQueryAsync(ctx workflow.Context, input *resourcegroups.UpdateGroupQueryInput) *ResourceGroupsUpdateGroupQueryFuture {
	future := workflow.ExecuteActivity(ctx, "aws.resourcegroups.UpdateGroupQuery", input)
	return &ResourceGroupsUpdateGroupQueryFuture{Future: future}
}
