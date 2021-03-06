// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package cognitosyncstub

import (
	"github.com/aws/aws-sdk-go/service/cognitosync"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CognitoSyncBulkPublishFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncBulkPublishFuture) Get(ctx workflow.Context) (*cognitosync.BulkPublishOutput, error) {
	var output cognitosync.BulkPublishOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncDeleteDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncDeleteDatasetFuture) Get(ctx workflow.Context) (*cognitosync.DeleteDatasetOutput, error) {
	var output cognitosync.DeleteDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncDescribeDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncDescribeDatasetFuture) Get(ctx workflow.Context) (*cognitosync.DescribeDatasetOutput, error) {
	var output cognitosync.DescribeDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncDescribeIdentityPoolUsageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncDescribeIdentityPoolUsageFuture) Get(ctx workflow.Context) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	var output cognitosync.DescribeIdentityPoolUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncDescribeIdentityUsageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncDescribeIdentityUsageFuture) Get(ctx workflow.Context) (*cognitosync.DescribeIdentityUsageOutput, error) {
	var output cognitosync.DescribeIdentityUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncGetBulkPublishDetailsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncGetBulkPublishDetailsFuture) Get(ctx workflow.Context) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	var output cognitosync.GetBulkPublishDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncGetCognitoEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncGetCognitoEventsFuture) Get(ctx workflow.Context) (*cognitosync.GetCognitoEventsOutput, error) {
	var output cognitosync.GetCognitoEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncGetIdentityPoolConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncGetIdentityPoolConfigurationFuture) Get(ctx workflow.Context) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	var output cognitosync.GetIdentityPoolConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncListDatasetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncListDatasetsFuture) Get(ctx workflow.Context) (*cognitosync.ListDatasetsOutput, error) {
	var output cognitosync.ListDatasetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncListIdentityPoolUsageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncListIdentityPoolUsageFuture) Get(ctx workflow.Context) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	var output cognitosync.ListIdentityPoolUsageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncListRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncListRecordsFuture) Get(ctx workflow.Context) (*cognitosync.ListRecordsOutput, error) {
	var output cognitosync.ListRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncRegisterDeviceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncRegisterDeviceFuture) Get(ctx workflow.Context) (*cognitosync.RegisterDeviceOutput, error) {
	var output cognitosync.RegisterDeviceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncSetCognitoEventsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncSetCognitoEventsFuture) Get(ctx workflow.Context) (*cognitosync.SetCognitoEventsOutput, error) {
	var output cognitosync.SetCognitoEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncSetIdentityPoolConfigurationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncSetIdentityPoolConfigurationFuture) Get(ctx workflow.Context) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	var output cognitosync.SetIdentityPoolConfigurationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncSubscribeToDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncSubscribeToDatasetFuture) Get(ctx workflow.Context) (*cognitosync.SubscribeToDatasetOutput, error) {
	var output cognitosync.SubscribeToDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncUnsubscribeFromDatasetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncUnsubscribeFromDatasetFuture) Get(ctx workflow.Context) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	var output cognitosync.UnsubscribeFromDatasetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CognitoSyncUpdateRecordsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CognitoSyncUpdateRecordsFuture) Get(ctx workflow.Context) (*cognitosync.UpdateRecordsOutput, error) {
	var output cognitosync.UpdateRecordsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) BulkPublish(ctx workflow.Context, input *cognitosync.BulkPublishInput) (*cognitosync.BulkPublishOutput, error) {
	var output cognitosync.BulkPublishOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.BulkPublish", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) BulkPublishAsync(ctx workflow.Context, input *cognitosync.BulkPublishInput) *CognitoSyncBulkPublishFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.BulkPublish", input)
	return &CognitoSyncBulkPublishFuture{Future: future}
}

func (a *stub) DeleteDataset(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) (*cognitosync.DeleteDatasetOutput, error) {
	var output cognitosync.DeleteDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.DeleteDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDatasetAsync(ctx workflow.Context, input *cognitosync.DeleteDatasetInput) *CognitoSyncDeleteDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.DeleteDataset", input)
	return &CognitoSyncDeleteDatasetFuture{Future: future}
}

func (a *stub) DescribeDataset(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) (*cognitosync.DescribeDatasetOutput, error) {
	var output cognitosync.DescribeDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDatasetAsync(ctx workflow.Context, input *cognitosync.DescribeDatasetInput) *CognitoSyncDescribeDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeDataset", input)
	return &CognitoSyncDescribeDatasetFuture{Future: future}
}

func (a *stub) DescribeIdentityPoolUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) (*cognitosync.DescribeIdentityPoolUsageOutput, error) {
	var output cognitosync.DescribeIdentityPoolUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeIdentityPoolUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityPoolUsageInput) *CognitoSyncDescribeIdentityPoolUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeIdentityPoolUsage", input)
	return &CognitoSyncDescribeIdentityPoolUsageFuture{Future: future}
}

func (a *stub) DescribeIdentityUsage(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) (*cognitosync.DescribeIdentityUsageOutput, error) {
	var output cognitosync.DescribeIdentityUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeIdentityUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeIdentityUsageAsync(ctx workflow.Context, input *cognitosync.DescribeIdentityUsageInput) *CognitoSyncDescribeIdentityUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.DescribeIdentityUsage", input)
	return &CognitoSyncDescribeIdentityUsageFuture{Future: future}
}

func (a *stub) GetBulkPublishDetails(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) (*cognitosync.GetBulkPublishDetailsOutput, error) {
	var output cognitosync.GetBulkPublishDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetBulkPublishDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetBulkPublishDetailsAsync(ctx workflow.Context, input *cognitosync.GetBulkPublishDetailsInput) *CognitoSyncGetBulkPublishDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetBulkPublishDetails", input)
	return &CognitoSyncGetBulkPublishDetailsFuture{Future: future}
}

func (a *stub) GetCognitoEvents(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) (*cognitosync.GetCognitoEventsOutput, error) {
	var output cognitosync.GetCognitoEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetCognitoEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.GetCognitoEventsInput) *CognitoSyncGetCognitoEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetCognitoEvents", input)
	return &CognitoSyncGetCognitoEventsFuture{Future: future}
}

func (a *stub) GetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) (*cognitosync.GetIdentityPoolConfigurationOutput, error) {
	var output cognitosync.GetIdentityPoolConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetIdentityPoolConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.GetIdentityPoolConfigurationInput) *CognitoSyncGetIdentityPoolConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.GetIdentityPoolConfiguration", input)
	return &CognitoSyncGetIdentityPoolConfigurationFuture{Future: future}
}

func (a *stub) ListDatasets(ctx workflow.Context, input *cognitosync.ListDatasetsInput) (*cognitosync.ListDatasetsOutput, error) {
	var output cognitosync.ListDatasetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListDatasets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDatasetsAsync(ctx workflow.Context, input *cognitosync.ListDatasetsInput) *CognitoSyncListDatasetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListDatasets", input)
	return &CognitoSyncListDatasetsFuture{Future: future}
}

func (a *stub) ListIdentityPoolUsage(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) (*cognitosync.ListIdentityPoolUsageOutput, error) {
	var output cognitosync.ListIdentityPoolUsageOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListIdentityPoolUsage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListIdentityPoolUsageAsync(ctx workflow.Context, input *cognitosync.ListIdentityPoolUsageInput) *CognitoSyncListIdentityPoolUsageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListIdentityPoolUsage", input)
	return &CognitoSyncListIdentityPoolUsageFuture{Future: future}
}

func (a *stub) ListRecords(ctx workflow.Context, input *cognitosync.ListRecordsInput) (*cognitosync.ListRecordsOutput, error) {
	var output cognitosync.ListRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListRecordsAsync(ctx workflow.Context, input *cognitosync.ListRecordsInput) *CognitoSyncListRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.ListRecords", input)
	return &CognitoSyncListRecordsFuture{Future: future}
}

func (a *stub) RegisterDevice(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) (*cognitosync.RegisterDeviceOutput, error) {
	var output cognitosync.RegisterDeviceOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.RegisterDevice", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterDeviceAsync(ctx workflow.Context, input *cognitosync.RegisterDeviceInput) *CognitoSyncRegisterDeviceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.RegisterDevice", input)
	return &CognitoSyncRegisterDeviceFuture{Future: future}
}

func (a *stub) SetCognitoEvents(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) (*cognitosync.SetCognitoEventsOutput, error) {
	var output cognitosync.SetCognitoEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.SetCognitoEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetCognitoEventsAsync(ctx workflow.Context, input *cognitosync.SetCognitoEventsInput) *CognitoSyncSetCognitoEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.SetCognitoEvents", input)
	return &CognitoSyncSetCognitoEventsFuture{Future: future}
}

func (a *stub) SetIdentityPoolConfiguration(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) (*cognitosync.SetIdentityPoolConfigurationOutput, error) {
	var output cognitosync.SetIdentityPoolConfigurationOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.SetIdentityPoolConfiguration", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetIdentityPoolConfigurationAsync(ctx workflow.Context, input *cognitosync.SetIdentityPoolConfigurationInput) *CognitoSyncSetIdentityPoolConfigurationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.SetIdentityPoolConfiguration", input)
	return &CognitoSyncSetIdentityPoolConfigurationFuture{Future: future}
}

func (a *stub) SubscribeToDataset(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) (*cognitosync.SubscribeToDatasetOutput, error) {
	var output cognitosync.SubscribeToDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.SubscribeToDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SubscribeToDatasetAsync(ctx workflow.Context, input *cognitosync.SubscribeToDatasetInput) *CognitoSyncSubscribeToDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.SubscribeToDataset", input)
	return &CognitoSyncSubscribeToDatasetFuture{Future: future}
}

func (a *stub) UnsubscribeFromDataset(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) (*cognitosync.UnsubscribeFromDatasetOutput, error) {
	var output cognitosync.UnsubscribeFromDatasetOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.UnsubscribeFromDataset", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UnsubscribeFromDatasetAsync(ctx workflow.Context, input *cognitosync.UnsubscribeFromDatasetInput) *CognitoSyncUnsubscribeFromDatasetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.UnsubscribeFromDataset", input)
	return &CognitoSyncUnsubscribeFromDatasetFuture{Future: future}
}

func (a *stub) UpdateRecords(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) (*cognitosync.UpdateRecordsOutput, error) {
	var output cognitosync.UpdateRecordsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cognitosync.UpdateRecords", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateRecordsAsync(ctx workflow.Context, input *cognitosync.UpdateRecordsInput) *CognitoSyncUpdateRecordsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cognitosync.UpdateRecords", input)
	return &CognitoSyncUpdateRecordsFuture{Future: future}
}
