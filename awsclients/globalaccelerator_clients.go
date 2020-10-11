// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsclients

import (
	"github.com/aws/aws-sdk-go/service/globalaccelerator"
	"go.temporal.io/sdk/workflow"
)

type GlobalAcceleratorClient interface {
	AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error)
	AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *GlobalAcceleratorAdvertiseByoipCidrFuture

	CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error)
	CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *GlobalAcceleratorCreateAcceleratorFuture

	CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error)
	CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *GlobalAcceleratorCreateEndpointGroupFuture

	CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error)
	CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *GlobalAcceleratorCreateListenerFuture

	DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error)
	DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *GlobalAcceleratorDeleteAcceleratorFuture

	DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error)
	DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *GlobalAcceleratorDeleteEndpointGroupFuture

	DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error)
	DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *GlobalAcceleratorDeleteListenerFuture

	DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error)
	DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *GlobalAcceleratorDeprovisionByoipCidrFuture

	DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error)
	DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *GlobalAcceleratorDescribeAcceleratorFuture

	DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error)
	DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *GlobalAcceleratorDescribeAcceleratorAttributesFuture

	DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error)
	DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *GlobalAcceleratorDescribeEndpointGroupFuture

	DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error)
	DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *GlobalAcceleratorDescribeListenerFuture

	ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error)
	ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *GlobalAcceleratorListAcceleratorsFuture

	ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error)
	ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *GlobalAcceleratorListByoipCidrsFuture

	ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error)
	ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *GlobalAcceleratorListEndpointGroupsFuture

	ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error)
	ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *GlobalAcceleratorListListenersFuture

	ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *GlobalAcceleratorListTagsForResourceFuture

	ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error)
	ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *GlobalAcceleratorProvisionByoipCidrFuture

	TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *GlobalAcceleratorTagResourceFuture

	UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *GlobalAcceleratorUntagResourceFuture

	UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error)
	UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *GlobalAcceleratorUpdateAcceleratorFuture

	UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error)
	UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *GlobalAcceleratorUpdateAcceleratorAttributesFuture

	UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error)
	UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *GlobalAcceleratorUpdateEndpointGroupFuture

	UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error)
	UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *GlobalAcceleratorUpdateListenerFuture

	WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error)
	WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *GlobalAcceleratorWithdrawByoipCidrFuture
}

type GlobalAcceleratorStub struct{}

func NewGlobalAcceleratorStub() GlobalAcceleratorClient {
	return &GlobalAcceleratorStub{}
}

type GlobalAcceleratorAdvertiseByoipCidrFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorAdvertiseByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	var output globalaccelerator.AdvertiseByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorCreateAcceleratorFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorCreateAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateAcceleratorOutput, error) {
	var output globalaccelerator.CreateAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorCreateEndpointGroupFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorCreateEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	var output globalaccelerator.CreateEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorCreateListenerFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorCreateListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.CreateListenerOutput, error) {
	var output globalaccelerator.CreateListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDeleteAcceleratorFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDeleteAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	var output globalaccelerator.DeleteAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDeleteEndpointGroupFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDeleteEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	var output globalaccelerator.DeleteEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDeleteListenerFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDeleteListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.DeleteListenerOutput, error) {
	var output globalaccelerator.DeleteListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDeprovisionByoipCidrFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDeprovisionByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	var output globalaccelerator.DeprovisionByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDescribeAcceleratorFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDescribeAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	var output globalaccelerator.DescribeAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDescribeAcceleratorAttributesFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDescribeAcceleratorAttributesFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	var output globalaccelerator.DescribeAcceleratorAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDescribeEndpointGroupFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDescribeEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	var output globalaccelerator.DescribeEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorDescribeListenerFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorDescribeListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.DescribeListenerOutput, error) {
	var output globalaccelerator.DescribeListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorListAcceleratorsFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorListAcceleratorsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListAcceleratorsOutput, error) {
	var output globalaccelerator.ListAcceleratorsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorListByoipCidrsFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorListByoipCidrsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListByoipCidrsOutput, error) {
	var output globalaccelerator.ListByoipCidrsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorListEndpointGroupsFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorListEndpointGroupsFuture) Get(ctx workflow.Context) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	var output globalaccelerator.ListEndpointGroupsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorListListenersFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorListListenersFuture) Get(ctx workflow.Context) (*globalaccelerator.ListListenersOutput, error) {
	var output globalaccelerator.ListListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorListTagsForResourceFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorListTagsForResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.ListTagsForResourceOutput, error) {
	var output globalaccelerator.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorProvisionByoipCidrFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorProvisionByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	var output globalaccelerator.ProvisionByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorTagResourceFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorTagResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.TagResourceOutput, error) {
	var output globalaccelerator.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorUntagResourceFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorUntagResourceFuture) Get(ctx workflow.Context) (*globalaccelerator.UntagResourceOutput, error) {
	var output globalaccelerator.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorUpdateAcceleratorFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorUpdateAcceleratorFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	var output globalaccelerator.UpdateAcceleratorOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorUpdateAcceleratorAttributesFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorUpdateAcceleratorAttributesFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	var output globalaccelerator.UpdateAcceleratorAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorUpdateEndpointGroupFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorUpdateEndpointGroupFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	var output globalaccelerator.UpdateEndpointGroupOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorUpdateListenerFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorUpdateListenerFuture) Get(ctx workflow.Context) (*globalaccelerator.UpdateListenerOutput, error) {
	var output globalaccelerator.UpdateListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GlobalAcceleratorWithdrawByoipCidrFuture struct {
	Future workflow.Future
}

func (r *GlobalAcceleratorWithdrawByoipCidrFuture) Get(ctx workflow.Context) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	var output globalaccelerator.WithdrawByoipCidrOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) AdvertiseByoipCidr(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) (*globalaccelerator.AdvertiseByoipCidrOutput, error) {
	var output globalaccelerator.AdvertiseByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.AdvertiseByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) AdvertiseByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.AdvertiseByoipCidrInput) *GlobalAcceleratorAdvertiseByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.AdvertiseByoipCidr", input)
	return &GlobalAcceleratorAdvertiseByoipCidrFuture{Future: future}
}

func (a *GlobalAcceleratorStub) CreateAccelerator(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) (*globalaccelerator.CreateAcceleratorOutput, error) {
	var output globalaccelerator.CreateAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) CreateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.CreateAcceleratorInput) *GlobalAcceleratorCreateAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateAccelerator", input)
	return &GlobalAcceleratorCreateAcceleratorFuture{Future: future}
}

func (a *GlobalAcceleratorStub) CreateEndpointGroup(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) (*globalaccelerator.CreateEndpointGroupOutput, error) {
	var output globalaccelerator.CreateEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) CreateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.CreateEndpointGroupInput) *GlobalAcceleratorCreateEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateEndpointGroup", input)
	return &GlobalAcceleratorCreateEndpointGroupFuture{Future: future}
}

func (a *GlobalAcceleratorStub) CreateListener(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) (*globalaccelerator.CreateListenerOutput, error) {
	var output globalaccelerator.CreateListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) CreateListenerAsync(ctx workflow.Context, input *globalaccelerator.CreateListenerInput) *GlobalAcceleratorCreateListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.CreateListener", input)
	return &GlobalAcceleratorCreateListenerFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DeleteAccelerator(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) (*globalaccelerator.DeleteAcceleratorOutput, error) {
	var output globalaccelerator.DeleteAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DeleteAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DeleteAcceleratorInput) *GlobalAcceleratorDeleteAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteAccelerator", input)
	return &GlobalAcceleratorDeleteAcceleratorFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DeleteEndpointGroup(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) (*globalaccelerator.DeleteEndpointGroupOutput, error) {
	var output globalaccelerator.DeleteEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DeleteEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DeleteEndpointGroupInput) *GlobalAcceleratorDeleteEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteEndpointGroup", input)
	return &GlobalAcceleratorDeleteEndpointGroupFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DeleteListener(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) (*globalaccelerator.DeleteListenerOutput, error) {
	var output globalaccelerator.DeleteListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteListener", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DeleteListenerAsync(ctx workflow.Context, input *globalaccelerator.DeleteListenerInput) *GlobalAcceleratorDeleteListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeleteListener", input)
	return &GlobalAcceleratorDeleteListenerFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DeprovisionByoipCidr(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) (*globalaccelerator.DeprovisionByoipCidrOutput, error) {
	var output globalaccelerator.DeprovisionByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeprovisionByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DeprovisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.DeprovisionByoipCidrInput) *GlobalAcceleratorDeprovisionByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DeprovisionByoipCidr", input)
	return &GlobalAcceleratorDeprovisionByoipCidrFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DescribeAccelerator(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) (*globalaccelerator.DescribeAcceleratorOutput, error) {
	var output globalaccelerator.DescribeAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorInput) *GlobalAcceleratorDescribeAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAccelerator", input)
	return &GlobalAcceleratorDescribeAcceleratorFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) (*globalaccelerator.DescribeAcceleratorAttributesOutput, error) {
	var output globalaccelerator.DescribeAcceleratorAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAcceleratorAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DescribeAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.DescribeAcceleratorAttributesInput) *GlobalAcceleratorDescribeAcceleratorAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeAcceleratorAttributes", input)
	return &GlobalAcceleratorDescribeAcceleratorAttributesFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DescribeEndpointGroup(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) (*globalaccelerator.DescribeEndpointGroupOutput, error) {
	var output globalaccelerator.DescribeEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DescribeEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.DescribeEndpointGroupInput) *GlobalAcceleratorDescribeEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeEndpointGroup", input)
	return &GlobalAcceleratorDescribeEndpointGroupFuture{Future: future}
}

func (a *GlobalAcceleratorStub) DescribeListener(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) (*globalaccelerator.DescribeListenerOutput, error) {
	var output globalaccelerator.DescribeListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeListener", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) DescribeListenerAsync(ctx workflow.Context, input *globalaccelerator.DescribeListenerInput) *GlobalAcceleratorDescribeListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.DescribeListener", input)
	return &GlobalAcceleratorDescribeListenerFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ListAccelerators(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) (*globalaccelerator.ListAcceleratorsOutput, error) {
	var output globalaccelerator.ListAcceleratorsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListAccelerators", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ListAcceleratorsAsync(ctx workflow.Context, input *globalaccelerator.ListAcceleratorsInput) *GlobalAcceleratorListAcceleratorsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListAccelerators", input)
	return &GlobalAcceleratorListAcceleratorsFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ListByoipCidrs(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) (*globalaccelerator.ListByoipCidrsOutput, error) {
	var output globalaccelerator.ListByoipCidrsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListByoipCidrs", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ListByoipCidrsAsync(ctx workflow.Context, input *globalaccelerator.ListByoipCidrsInput) *GlobalAcceleratorListByoipCidrsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListByoipCidrs", input)
	return &GlobalAcceleratorListByoipCidrsFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ListEndpointGroups(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) (*globalaccelerator.ListEndpointGroupsOutput, error) {
	var output globalaccelerator.ListEndpointGroupsOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListEndpointGroups", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ListEndpointGroupsAsync(ctx workflow.Context, input *globalaccelerator.ListEndpointGroupsInput) *GlobalAcceleratorListEndpointGroupsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListEndpointGroups", input)
	return &GlobalAcceleratorListEndpointGroupsFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ListListeners(ctx workflow.Context, input *globalaccelerator.ListListenersInput) (*globalaccelerator.ListListenersOutput, error) {
	var output globalaccelerator.ListListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ListListenersAsync(ctx workflow.Context, input *globalaccelerator.ListListenersInput) *GlobalAcceleratorListListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListListeners", input)
	return &GlobalAcceleratorListListenersFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ListTagsForResource(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) (*globalaccelerator.ListTagsForResourceOutput, error) {
	var output globalaccelerator.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ListTagsForResourceAsync(ctx workflow.Context, input *globalaccelerator.ListTagsForResourceInput) *GlobalAcceleratorListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ListTagsForResource", input)
	return &GlobalAcceleratorListTagsForResourceFuture{Future: future}
}

func (a *GlobalAcceleratorStub) ProvisionByoipCidr(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) (*globalaccelerator.ProvisionByoipCidrOutput, error) {
	var output globalaccelerator.ProvisionByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ProvisionByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) ProvisionByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.ProvisionByoipCidrInput) *GlobalAcceleratorProvisionByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.ProvisionByoipCidr", input)
	return &GlobalAcceleratorProvisionByoipCidrFuture{Future: future}
}

func (a *GlobalAcceleratorStub) TagResource(ctx workflow.Context, input *globalaccelerator.TagResourceInput) (*globalaccelerator.TagResourceOutput, error) {
	var output globalaccelerator.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) TagResourceAsync(ctx workflow.Context, input *globalaccelerator.TagResourceInput) *GlobalAcceleratorTagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.TagResource", input)
	return &GlobalAcceleratorTagResourceFuture{Future: future}
}

func (a *GlobalAcceleratorStub) UntagResource(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) (*globalaccelerator.UntagResourceOutput, error) {
	var output globalaccelerator.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) UntagResourceAsync(ctx workflow.Context, input *globalaccelerator.UntagResourceInput) *GlobalAcceleratorUntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UntagResource", input)
	return &GlobalAcceleratorUntagResourceFuture{Future: future}
}

func (a *GlobalAcceleratorStub) UpdateAccelerator(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) (*globalaccelerator.UpdateAcceleratorOutput, error) {
	var output globalaccelerator.UpdateAcceleratorOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAccelerator", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorInput) *GlobalAcceleratorUpdateAcceleratorFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAccelerator", input)
	return &GlobalAcceleratorUpdateAcceleratorFuture{Future: future}
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAttributes(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) (*globalaccelerator.UpdateAcceleratorAttributesOutput, error) {
	var output globalaccelerator.UpdateAcceleratorAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAcceleratorAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) UpdateAcceleratorAttributesAsync(ctx workflow.Context, input *globalaccelerator.UpdateAcceleratorAttributesInput) *GlobalAcceleratorUpdateAcceleratorAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateAcceleratorAttributes", input)
	return &GlobalAcceleratorUpdateAcceleratorAttributesFuture{Future: future}
}

func (a *GlobalAcceleratorStub) UpdateEndpointGroup(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) (*globalaccelerator.UpdateEndpointGroupOutput, error) {
	var output globalaccelerator.UpdateEndpointGroupOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateEndpointGroup", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) UpdateEndpointGroupAsync(ctx workflow.Context, input *globalaccelerator.UpdateEndpointGroupInput) *GlobalAcceleratorUpdateEndpointGroupFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateEndpointGroup", input)
	return &GlobalAcceleratorUpdateEndpointGroupFuture{Future: future}
}

func (a *GlobalAcceleratorStub) UpdateListener(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) (*globalaccelerator.UpdateListenerOutput, error) {
	var output globalaccelerator.UpdateListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateListener", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) UpdateListenerAsync(ctx workflow.Context, input *globalaccelerator.UpdateListenerInput) *GlobalAcceleratorUpdateListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.UpdateListener", input)
	return &GlobalAcceleratorUpdateListenerFuture{Future: future}
}

func (a *GlobalAcceleratorStub) WithdrawByoipCidr(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) (*globalaccelerator.WithdrawByoipCidrOutput, error) {
	var output globalaccelerator.WithdrawByoipCidrOutput
	err := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.WithdrawByoipCidr", input).Get(ctx, &output)
	return &output, err
}

func (a *GlobalAcceleratorStub) WithdrawByoipCidrAsync(ctx workflow.Context, input *globalaccelerator.WithdrawByoipCidrInput) *GlobalAcceleratorWithdrawByoipCidrFuture {
	future := workflow.ExecuteActivity(ctx, "aws.globalaccelerator.WithdrawByoipCidr", input)
	return &GlobalAcceleratorWithdrawByoipCidrFuture{Future: future}
}
