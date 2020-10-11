// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/cloudtrail"
	"go.temporal.io/sdk/workflow"
)

type CloudTrailClient interface {
	AddTags(ctx workflow.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *cloudtrail.AddTagsInput) *CloudTrailAddTagsFuture

	CreateTrail(ctx workflow.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error)
	CreateTrailAsync(ctx workflow.Context, input *cloudtrail.CreateTrailInput) *CloudTrailCreateTrailFuture

	DeleteTrail(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error)
	DeleteTrailAsync(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) *CloudTrailDeleteTrailFuture

	DescribeTrails(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error)
	DescribeTrailsAsync(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) *CloudTrailDescribeTrailsFuture

	GetEventSelectors(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error)
	GetEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) *CloudTrailGetEventSelectorsFuture

	GetInsightSelectors(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error)
	GetInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) *CloudTrailGetInsightSelectorsFuture

	GetTrail(ctx workflow.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error)
	GetTrailAsync(ctx workflow.Context, input *cloudtrail.GetTrailInput) *CloudTrailGetTrailFuture

	GetTrailStatus(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error)
	GetTrailStatusAsync(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) *CloudTrailGetTrailStatusFuture

	ListPublicKeys(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error)
	ListPublicKeysAsync(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) *CloudTrailListPublicKeysFuture

	ListTags(ctx workflow.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *cloudtrail.ListTagsInput) *CloudTrailListTagsFuture

	ListTrails(ctx workflow.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error)
	ListTrailsAsync(ctx workflow.Context, input *cloudtrail.ListTrailsInput) *CloudTrailListTrailsFuture

	LookupEvents(ctx workflow.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error)
	LookupEventsAsync(ctx workflow.Context, input *cloudtrail.LookupEventsInput) *CloudTrailLookupEventsFuture

	PutEventSelectors(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error)
	PutEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) *CloudTrailPutEventSelectorsFuture

	PutInsightSelectors(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error)
	PutInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) *CloudTrailPutInsightSelectorsFuture

	RemoveTags(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) *CloudTrailRemoveTagsFuture

	StartLogging(ctx workflow.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error)
	StartLoggingAsync(ctx workflow.Context, input *cloudtrail.StartLoggingInput) *CloudTrailStartLoggingFuture

	StopLogging(ctx workflow.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error)
	StopLoggingAsync(ctx workflow.Context, input *cloudtrail.StopLoggingInput) *CloudTrailStopLoggingFuture

	UpdateTrail(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error)
	UpdateTrailAsync(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) *CloudTrailUpdateTrailFuture
}

type CloudTrailStub struct{}

func NewCloudTrailStub() CloudTrailClient {
	return &CloudTrailStub{}
}

type CloudTrailAddTagsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailAddTagsFuture) Get(ctx workflow.Context) (*cloudtrail.AddTagsOutput, error) {
	var output cloudtrail.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailCreateTrailFuture struct {
	Future workflow.Future
}

func (r *CloudTrailCreateTrailFuture) Get(ctx workflow.Context) (*cloudtrail.CreateTrailOutput, error) {
	var output cloudtrail.CreateTrailOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailDeleteTrailFuture struct {
	Future workflow.Future
}

func (r *CloudTrailDeleteTrailFuture) Get(ctx workflow.Context) (*cloudtrail.DeleteTrailOutput, error) {
	var output cloudtrail.DeleteTrailOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailDescribeTrailsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailDescribeTrailsFuture) Get(ctx workflow.Context) (*cloudtrail.DescribeTrailsOutput, error) {
	var output cloudtrail.DescribeTrailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailGetEventSelectorsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailGetEventSelectorsFuture) Get(ctx workflow.Context) (*cloudtrail.GetEventSelectorsOutput, error) {
	var output cloudtrail.GetEventSelectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailGetInsightSelectorsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailGetInsightSelectorsFuture) Get(ctx workflow.Context) (*cloudtrail.GetInsightSelectorsOutput, error) {
	var output cloudtrail.GetInsightSelectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailGetTrailFuture struct {
	Future workflow.Future
}

func (r *CloudTrailGetTrailFuture) Get(ctx workflow.Context) (*cloudtrail.GetTrailOutput, error) {
	var output cloudtrail.GetTrailOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailGetTrailStatusFuture struct {
	Future workflow.Future
}

func (r *CloudTrailGetTrailStatusFuture) Get(ctx workflow.Context) (*cloudtrail.GetTrailStatusOutput, error) {
	var output cloudtrail.GetTrailStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailListPublicKeysFuture struct {
	Future workflow.Future
}

func (r *CloudTrailListPublicKeysFuture) Get(ctx workflow.Context) (*cloudtrail.ListPublicKeysOutput, error) {
	var output cloudtrail.ListPublicKeysOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailListTagsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailListTagsFuture) Get(ctx workflow.Context) (*cloudtrail.ListTagsOutput, error) {
	var output cloudtrail.ListTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailListTrailsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailListTrailsFuture) Get(ctx workflow.Context) (*cloudtrail.ListTrailsOutput, error) {
	var output cloudtrail.ListTrailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailLookupEventsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailLookupEventsFuture) Get(ctx workflow.Context) (*cloudtrail.LookupEventsOutput, error) {
	var output cloudtrail.LookupEventsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailPutEventSelectorsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailPutEventSelectorsFuture) Get(ctx workflow.Context) (*cloudtrail.PutEventSelectorsOutput, error) {
	var output cloudtrail.PutEventSelectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailPutInsightSelectorsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailPutInsightSelectorsFuture) Get(ctx workflow.Context) (*cloudtrail.PutInsightSelectorsOutput, error) {
	var output cloudtrail.PutInsightSelectorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailRemoveTagsFuture struct {
	Future workflow.Future
}

func (r *CloudTrailRemoveTagsFuture) Get(ctx workflow.Context) (*cloudtrail.RemoveTagsOutput, error) {
	var output cloudtrail.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailStartLoggingFuture struct {
	Future workflow.Future
}

func (r *CloudTrailStartLoggingFuture) Get(ctx workflow.Context) (*cloudtrail.StartLoggingOutput, error) {
	var output cloudtrail.StartLoggingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailStopLoggingFuture struct {
	Future workflow.Future
}

func (r *CloudTrailStopLoggingFuture) Get(ctx workflow.Context) (*cloudtrail.StopLoggingOutput, error) {
	var output cloudtrail.StopLoggingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CloudTrailUpdateTrailFuture struct {
	Future workflow.Future
}

func (r *CloudTrailUpdateTrailFuture) Get(ctx workflow.Context) (*cloudtrail.UpdateTrailOutput, error) {
	var output cloudtrail.UpdateTrailOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) AddTags(ctx workflow.Context, input *cloudtrail.AddTagsInput) (*cloudtrail.AddTagsOutput, error) {
	var output cloudtrail.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) AddTagsAsync(ctx workflow.Context, input *cloudtrail.AddTagsInput) *CloudTrailAddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.AddTags", input)
	return &CloudTrailAddTagsFuture{Future: future}
}

func (a *CloudTrailStub) CreateTrail(ctx workflow.Context, input *cloudtrail.CreateTrailInput) (*cloudtrail.CreateTrailOutput, error) {
	var output cloudtrail.CreateTrailOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.CreateTrail", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) CreateTrailAsync(ctx workflow.Context, input *cloudtrail.CreateTrailInput) *CloudTrailCreateTrailFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.CreateTrail", input)
	return &CloudTrailCreateTrailFuture{Future: future}
}

func (a *CloudTrailStub) DeleteTrail(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) (*cloudtrail.DeleteTrailOutput, error) {
	var output cloudtrail.DeleteTrailOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.DeleteTrail", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) DeleteTrailAsync(ctx workflow.Context, input *cloudtrail.DeleteTrailInput) *CloudTrailDeleteTrailFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.DeleteTrail", input)
	return &CloudTrailDeleteTrailFuture{Future: future}
}

func (a *CloudTrailStub) DescribeTrails(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) (*cloudtrail.DescribeTrailsOutput, error) {
	var output cloudtrail.DescribeTrailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.DescribeTrails", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) DescribeTrailsAsync(ctx workflow.Context, input *cloudtrail.DescribeTrailsInput) *CloudTrailDescribeTrailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.DescribeTrails", input)
	return &CloudTrailDescribeTrailsFuture{Future: future}
}

func (a *CloudTrailStub) GetEventSelectors(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) (*cloudtrail.GetEventSelectorsOutput, error) {
	var output cloudtrail.GetEventSelectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetEventSelectors", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) GetEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetEventSelectorsInput) *CloudTrailGetEventSelectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetEventSelectors", input)
	return &CloudTrailGetEventSelectorsFuture{Future: future}
}

func (a *CloudTrailStub) GetInsightSelectors(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) (*cloudtrail.GetInsightSelectorsOutput, error) {
	var output cloudtrail.GetInsightSelectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetInsightSelectors", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) GetInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.GetInsightSelectorsInput) *CloudTrailGetInsightSelectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetInsightSelectors", input)
	return &CloudTrailGetInsightSelectorsFuture{Future: future}
}

func (a *CloudTrailStub) GetTrail(ctx workflow.Context, input *cloudtrail.GetTrailInput) (*cloudtrail.GetTrailOutput, error) {
	var output cloudtrail.GetTrailOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetTrail", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) GetTrailAsync(ctx workflow.Context, input *cloudtrail.GetTrailInput) *CloudTrailGetTrailFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetTrail", input)
	return &CloudTrailGetTrailFuture{Future: future}
}

func (a *CloudTrailStub) GetTrailStatus(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) (*cloudtrail.GetTrailStatusOutput, error) {
	var output cloudtrail.GetTrailStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetTrailStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) GetTrailStatusAsync(ctx workflow.Context, input *cloudtrail.GetTrailStatusInput) *CloudTrailGetTrailStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.GetTrailStatus", input)
	return &CloudTrailGetTrailStatusFuture{Future: future}
}

func (a *CloudTrailStub) ListPublicKeys(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) (*cloudtrail.ListPublicKeysOutput, error) {
	var output cloudtrail.ListPublicKeysOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListPublicKeys", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) ListPublicKeysAsync(ctx workflow.Context, input *cloudtrail.ListPublicKeysInput) *CloudTrailListPublicKeysFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListPublicKeys", input)
	return &CloudTrailListPublicKeysFuture{Future: future}
}

func (a *CloudTrailStub) ListTags(ctx workflow.Context, input *cloudtrail.ListTagsInput) (*cloudtrail.ListTagsOutput, error) {
	var output cloudtrail.ListTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) ListTagsAsync(ctx workflow.Context, input *cloudtrail.ListTagsInput) *CloudTrailListTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListTags", input)
	return &CloudTrailListTagsFuture{Future: future}
}

func (a *CloudTrailStub) ListTrails(ctx workflow.Context, input *cloudtrail.ListTrailsInput) (*cloudtrail.ListTrailsOutput, error) {
	var output cloudtrail.ListTrailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListTrails", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) ListTrailsAsync(ctx workflow.Context, input *cloudtrail.ListTrailsInput) *CloudTrailListTrailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.ListTrails", input)
	return &CloudTrailListTrailsFuture{Future: future}
}

func (a *CloudTrailStub) LookupEvents(ctx workflow.Context, input *cloudtrail.LookupEventsInput) (*cloudtrail.LookupEventsOutput, error) {
	var output cloudtrail.LookupEventsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.LookupEvents", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) LookupEventsAsync(ctx workflow.Context, input *cloudtrail.LookupEventsInput) *CloudTrailLookupEventsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.LookupEvents", input)
	return &CloudTrailLookupEventsFuture{Future: future}
}

func (a *CloudTrailStub) PutEventSelectors(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) (*cloudtrail.PutEventSelectorsOutput, error) {
	var output cloudtrail.PutEventSelectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.PutEventSelectors", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) PutEventSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutEventSelectorsInput) *CloudTrailPutEventSelectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.PutEventSelectors", input)
	return &CloudTrailPutEventSelectorsFuture{Future: future}
}

func (a *CloudTrailStub) PutInsightSelectors(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) (*cloudtrail.PutInsightSelectorsOutput, error) {
	var output cloudtrail.PutInsightSelectorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.PutInsightSelectors", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) PutInsightSelectorsAsync(ctx workflow.Context, input *cloudtrail.PutInsightSelectorsInput) *CloudTrailPutInsightSelectorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.PutInsightSelectors", input)
	return &CloudTrailPutInsightSelectorsFuture{Future: future}
}

func (a *CloudTrailStub) RemoveTags(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) (*cloudtrail.RemoveTagsOutput, error) {
	var output cloudtrail.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) RemoveTagsAsync(ctx workflow.Context, input *cloudtrail.RemoveTagsInput) *CloudTrailRemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.RemoveTags", input)
	return &CloudTrailRemoveTagsFuture{Future: future}
}

func (a *CloudTrailStub) StartLogging(ctx workflow.Context, input *cloudtrail.StartLoggingInput) (*cloudtrail.StartLoggingOutput, error) {
	var output cloudtrail.StartLoggingOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.StartLogging", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) StartLoggingAsync(ctx workflow.Context, input *cloudtrail.StartLoggingInput) *CloudTrailStartLoggingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.StartLogging", input)
	return &CloudTrailStartLoggingFuture{Future: future}
}

func (a *CloudTrailStub) StopLogging(ctx workflow.Context, input *cloudtrail.StopLoggingInput) (*cloudtrail.StopLoggingOutput, error) {
	var output cloudtrail.StopLoggingOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.StopLogging", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) StopLoggingAsync(ctx workflow.Context, input *cloudtrail.StopLoggingInput) *CloudTrailStopLoggingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.StopLogging", input)
	return &CloudTrailStopLoggingFuture{Future: future}
}

func (a *CloudTrailStub) UpdateTrail(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) (*cloudtrail.UpdateTrailOutput, error) {
	var output cloudtrail.UpdateTrailOutput
	err := workflow.ExecuteActivity(ctx, "aws.cloudtrail.UpdateTrail", input).Get(ctx, &output)
	return &output, err
}

func (a *CloudTrailStub) UpdateTrailAsync(ctx workflow.Context, input *cloudtrail.UpdateTrailInput) *CloudTrailUpdateTrailFuture {
	future := workflow.ExecuteActivity(ctx, "aws.cloudtrail.UpdateTrail", input)
	return &CloudTrailUpdateTrailFuture{Future: future}
}
