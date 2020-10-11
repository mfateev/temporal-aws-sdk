// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/migrationhub"
	"go.temporal.io/sdk/workflow"
)

type MigrationHubClient interface {
	AssociateCreatedArtifact(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error)
	AssociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) *MigrationHubAssociateCreatedArtifactFuture

	AssociateDiscoveredResource(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error)
	AssociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) *MigrationHubAssociateDiscoveredResourceFuture

	CreateProgressUpdateStream(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error)
	CreateProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) *MigrationHubCreateProgressUpdateStreamFuture

	DeleteProgressUpdateStream(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error)
	DeleteProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) *MigrationHubDeleteProgressUpdateStreamFuture

	DescribeApplicationState(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error)
	DescribeApplicationStateAsync(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) *MigrationHubDescribeApplicationStateFuture

	DescribeMigrationTask(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error)
	DescribeMigrationTaskAsync(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) *MigrationHubDescribeMigrationTaskFuture

	DisassociateCreatedArtifact(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error)
	DisassociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) *MigrationHubDisassociateCreatedArtifactFuture

	DisassociateDiscoveredResource(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error)
	DisassociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) *MigrationHubDisassociateDiscoveredResourceFuture

	ImportMigrationTask(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error)
	ImportMigrationTaskAsync(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) *MigrationHubImportMigrationTaskFuture

	ListApplicationStates(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error)
	ListApplicationStatesAsync(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) *MigrationHubListApplicationStatesFuture

	ListCreatedArtifacts(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error)
	ListCreatedArtifactsAsync(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) *MigrationHubListCreatedArtifactsFuture

	ListDiscoveredResources(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error)
	ListDiscoveredResourcesAsync(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) *MigrationHubListDiscoveredResourcesFuture

	ListMigrationTasks(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error)
	ListMigrationTasksAsync(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) *MigrationHubListMigrationTasksFuture

	ListProgressUpdateStreams(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error)
	ListProgressUpdateStreamsAsync(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) *MigrationHubListProgressUpdateStreamsFuture

	NotifyApplicationState(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error)
	NotifyApplicationStateAsync(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) *MigrationHubNotifyApplicationStateFuture

	NotifyMigrationTaskState(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error)
	NotifyMigrationTaskStateAsync(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) *MigrationHubNotifyMigrationTaskStateFuture

	PutResourceAttributes(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error)
	PutResourceAttributesAsync(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) *MigrationHubPutResourceAttributesFuture
}

type MigrationHubStub struct{}

func NewMigrationHubStub() MigrationHubClient {
	return &MigrationHubStub{}
}

type MigrationHubAssociateCreatedArtifactFuture struct {
	Future workflow.Future
}

func (r *MigrationHubAssociateCreatedArtifactFuture) Get(ctx workflow.Context) (*migrationhub.AssociateCreatedArtifactOutput, error) {
	var output migrationhub.AssociateCreatedArtifactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubAssociateDiscoveredResourceFuture struct {
	Future workflow.Future
}

func (r *MigrationHubAssociateDiscoveredResourceFuture) Get(ctx workflow.Context) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
	var output migrationhub.AssociateDiscoveredResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubCreateProgressUpdateStreamFuture struct {
	Future workflow.Future
}

func (r *MigrationHubCreateProgressUpdateStreamFuture) Get(ctx workflow.Context) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
	var output migrationhub.CreateProgressUpdateStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubDeleteProgressUpdateStreamFuture struct {
	Future workflow.Future
}

func (r *MigrationHubDeleteProgressUpdateStreamFuture) Get(ctx workflow.Context) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
	var output migrationhub.DeleteProgressUpdateStreamOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubDescribeApplicationStateFuture struct {
	Future workflow.Future
}

func (r *MigrationHubDescribeApplicationStateFuture) Get(ctx workflow.Context) (*migrationhub.DescribeApplicationStateOutput, error) {
	var output migrationhub.DescribeApplicationStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubDescribeMigrationTaskFuture struct {
	Future workflow.Future
}

func (r *MigrationHubDescribeMigrationTaskFuture) Get(ctx workflow.Context) (*migrationhub.DescribeMigrationTaskOutput, error) {
	var output migrationhub.DescribeMigrationTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubDisassociateCreatedArtifactFuture struct {
	Future workflow.Future
}

func (r *MigrationHubDisassociateCreatedArtifactFuture) Get(ctx workflow.Context) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
	var output migrationhub.DisassociateCreatedArtifactOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubDisassociateDiscoveredResourceFuture struct {
	Future workflow.Future
}

func (r *MigrationHubDisassociateDiscoveredResourceFuture) Get(ctx workflow.Context) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
	var output migrationhub.DisassociateDiscoveredResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubImportMigrationTaskFuture struct {
	Future workflow.Future
}

func (r *MigrationHubImportMigrationTaskFuture) Get(ctx workflow.Context) (*migrationhub.ImportMigrationTaskOutput, error) {
	var output migrationhub.ImportMigrationTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubListApplicationStatesFuture struct {
	Future workflow.Future
}

func (r *MigrationHubListApplicationStatesFuture) Get(ctx workflow.Context) (*migrationhub.ListApplicationStatesOutput, error) {
	var output migrationhub.ListApplicationStatesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubListCreatedArtifactsFuture struct {
	Future workflow.Future
}

func (r *MigrationHubListCreatedArtifactsFuture) Get(ctx workflow.Context) (*migrationhub.ListCreatedArtifactsOutput, error) {
	var output migrationhub.ListCreatedArtifactsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubListDiscoveredResourcesFuture struct {
	Future workflow.Future
}

func (r *MigrationHubListDiscoveredResourcesFuture) Get(ctx workflow.Context) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	var output migrationhub.ListDiscoveredResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubListMigrationTasksFuture struct {
	Future workflow.Future
}

func (r *MigrationHubListMigrationTasksFuture) Get(ctx workflow.Context) (*migrationhub.ListMigrationTasksOutput, error) {
	var output migrationhub.ListMigrationTasksOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubListProgressUpdateStreamsFuture struct {
	Future workflow.Future
}

func (r *MigrationHubListProgressUpdateStreamsFuture) Get(ctx workflow.Context) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	var output migrationhub.ListProgressUpdateStreamsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubNotifyApplicationStateFuture struct {
	Future workflow.Future
}

func (r *MigrationHubNotifyApplicationStateFuture) Get(ctx workflow.Context) (*migrationhub.NotifyApplicationStateOutput, error) {
	var output migrationhub.NotifyApplicationStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubNotifyMigrationTaskStateFuture struct {
	Future workflow.Future
}

func (r *MigrationHubNotifyMigrationTaskStateFuture) Get(ctx workflow.Context) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
	var output migrationhub.NotifyMigrationTaskStateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type MigrationHubPutResourceAttributesFuture struct {
	Future workflow.Future
}

func (r *MigrationHubPutResourceAttributesFuture) Get(ctx workflow.Context) (*migrationhub.PutResourceAttributesOutput, error) {
	var output migrationhub.PutResourceAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) AssociateCreatedArtifact(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) (*migrationhub.AssociateCreatedArtifactOutput, error) {
	var output migrationhub.AssociateCreatedArtifactOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.AssociateCreatedArtifact", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) AssociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.AssociateCreatedArtifactInput) *MigrationHubAssociateCreatedArtifactFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.AssociateCreatedArtifact", input)
	return &MigrationHubAssociateCreatedArtifactFuture{Future: future}
}

func (a *MigrationHubStub) AssociateDiscoveredResource(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) (*migrationhub.AssociateDiscoveredResourceOutput, error) {
	var output migrationhub.AssociateDiscoveredResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.AssociateDiscoveredResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) AssociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.AssociateDiscoveredResourceInput) *MigrationHubAssociateDiscoveredResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.AssociateDiscoveredResource", input)
	return &MigrationHubAssociateDiscoveredResourceFuture{Future: future}
}

func (a *MigrationHubStub) CreateProgressUpdateStream(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) (*migrationhub.CreateProgressUpdateStreamOutput, error) {
	var output migrationhub.CreateProgressUpdateStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.CreateProgressUpdateStream", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) CreateProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.CreateProgressUpdateStreamInput) *MigrationHubCreateProgressUpdateStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.CreateProgressUpdateStream", input)
	return &MigrationHubCreateProgressUpdateStreamFuture{Future: future}
}

func (a *MigrationHubStub) DeleteProgressUpdateStream(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) (*migrationhub.DeleteProgressUpdateStreamOutput, error) {
	var output migrationhub.DeleteProgressUpdateStreamOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.DeleteProgressUpdateStream", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) DeleteProgressUpdateStreamAsync(ctx workflow.Context, input *migrationhub.DeleteProgressUpdateStreamInput) *MigrationHubDeleteProgressUpdateStreamFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.DeleteProgressUpdateStream", input)
	return &MigrationHubDeleteProgressUpdateStreamFuture{Future: future}
}

func (a *MigrationHubStub) DescribeApplicationState(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) (*migrationhub.DescribeApplicationStateOutput, error) {
	var output migrationhub.DescribeApplicationStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.DescribeApplicationState", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) DescribeApplicationStateAsync(ctx workflow.Context, input *migrationhub.DescribeApplicationStateInput) *MigrationHubDescribeApplicationStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.DescribeApplicationState", input)
	return &MigrationHubDescribeApplicationStateFuture{Future: future}
}

func (a *MigrationHubStub) DescribeMigrationTask(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) (*migrationhub.DescribeMigrationTaskOutput, error) {
	var output migrationhub.DescribeMigrationTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.DescribeMigrationTask", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) DescribeMigrationTaskAsync(ctx workflow.Context, input *migrationhub.DescribeMigrationTaskInput) *MigrationHubDescribeMigrationTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.DescribeMigrationTask", input)
	return &MigrationHubDescribeMigrationTaskFuture{Future: future}
}

func (a *MigrationHubStub) DisassociateCreatedArtifact(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) (*migrationhub.DisassociateCreatedArtifactOutput, error) {
	var output migrationhub.DisassociateCreatedArtifactOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.DisassociateCreatedArtifact", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) DisassociateCreatedArtifactAsync(ctx workflow.Context, input *migrationhub.DisassociateCreatedArtifactInput) *MigrationHubDisassociateCreatedArtifactFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.DisassociateCreatedArtifact", input)
	return &MigrationHubDisassociateCreatedArtifactFuture{Future: future}
}

func (a *MigrationHubStub) DisassociateDiscoveredResource(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) (*migrationhub.DisassociateDiscoveredResourceOutput, error) {
	var output migrationhub.DisassociateDiscoveredResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.DisassociateDiscoveredResource", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) DisassociateDiscoveredResourceAsync(ctx workflow.Context, input *migrationhub.DisassociateDiscoveredResourceInput) *MigrationHubDisassociateDiscoveredResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.DisassociateDiscoveredResource", input)
	return &MigrationHubDisassociateDiscoveredResourceFuture{Future: future}
}

func (a *MigrationHubStub) ImportMigrationTask(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) (*migrationhub.ImportMigrationTaskOutput, error) {
	var output migrationhub.ImportMigrationTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ImportMigrationTask", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ImportMigrationTaskAsync(ctx workflow.Context, input *migrationhub.ImportMigrationTaskInput) *MigrationHubImportMigrationTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ImportMigrationTask", input)
	return &MigrationHubImportMigrationTaskFuture{Future: future}
}

func (a *MigrationHubStub) ListApplicationStates(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) (*migrationhub.ListApplicationStatesOutput, error) {
	var output migrationhub.ListApplicationStatesOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListApplicationStates", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ListApplicationStatesAsync(ctx workflow.Context, input *migrationhub.ListApplicationStatesInput) *MigrationHubListApplicationStatesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListApplicationStates", input)
	return &MigrationHubListApplicationStatesFuture{Future: future}
}

func (a *MigrationHubStub) ListCreatedArtifacts(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) (*migrationhub.ListCreatedArtifactsOutput, error) {
	var output migrationhub.ListCreatedArtifactsOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListCreatedArtifacts", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ListCreatedArtifactsAsync(ctx workflow.Context, input *migrationhub.ListCreatedArtifactsInput) *MigrationHubListCreatedArtifactsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListCreatedArtifacts", input)
	return &MigrationHubListCreatedArtifactsFuture{Future: future}
}

func (a *MigrationHubStub) ListDiscoveredResources(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) (*migrationhub.ListDiscoveredResourcesOutput, error) {
	var output migrationhub.ListDiscoveredResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListDiscoveredResources", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ListDiscoveredResourcesAsync(ctx workflow.Context, input *migrationhub.ListDiscoveredResourcesInput) *MigrationHubListDiscoveredResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListDiscoveredResources", input)
	return &MigrationHubListDiscoveredResourcesFuture{Future: future}
}

func (a *MigrationHubStub) ListMigrationTasks(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) (*migrationhub.ListMigrationTasksOutput, error) {
	var output migrationhub.ListMigrationTasksOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListMigrationTasks", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ListMigrationTasksAsync(ctx workflow.Context, input *migrationhub.ListMigrationTasksInput) *MigrationHubListMigrationTasksFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListMigrationTasks", input)
	return &MigrationHubListMigrationTasksFuture{Future: future}
}

func (a *MigrationHubStub) ListProgressUpdateStreams(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) (*migrationhub.ListProgressUpdateStreamsOutput, error) {
	var output migrationhub.ListProgressUpdateStreamsOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListProgressUpdateStreams", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) ListProgressUpdateStreamsAsync(ctx workflow.Context, input *migrationhub.ListProgressUpdateStreamsInput) *MigrationHubListProgressUpdateStreamsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.ListProgressUpdateStreams", input)
	return &MigrationHubListProgressUpdateStreamsFuture{Future: future}
}

func (a *MigrationHubStub) NotifyApplicationState(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) (*migrationhub.NotifyApplicationStateOutput, error) {
	var output migrationhub.NotifyApplicationStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.NotifyApplicationState", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) NotifyApplicationStateAsync(ctx workflow.Context, input *migrationhub.NotifyApplicationStateInput) *MigrationHubNotifyApplicationStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.NotifyApplicationState", input)
	return &MigrationHubNotifyApplicationStateFuture{Future: future}
}

func (a *MigrationHubStub) NotifyMigrationTaskState(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) (*migrationhub.NotifyMigrationTaskStateOutput, error) {
	var output migrationhub.NotifyMigrationTaskStateOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.NotifyMigrationTaskState", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) NotifyMigrationTaskStateAsync(ctx workflow.Context, input *migrationhub.NotifyMigrationTaskStateInput) *MigrationHubNotifyMigrationTaskStateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.NotifyMigrationTaskState", input)
	return &MigrationHubNotifyMigrationTaskStateFuture{Future: future}
}

func (a *MigrationHubStub) PutResourceAttributes(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) (*migrationhub.PutResourceAttributesOutput, error) {
	var output migrationhub.PutResourceAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.migrationhub.PutResourceAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *MigrationHubStub) PutResourceAttributesAsync(ctx workflow.Context, input *migrationhub.PutResourceAttributesInput) *MigrationHubPutResourceAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.migrationhub.PutResourceAttributes", input)
	return &MigrationHubPutResourceAttributesFuture{Future: future}
}
