// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package pinpointsmsvoicestub

import (
	"github.com/aws/aws-sdk-go/service/pinpointsmsvoice"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type PinpointSMSVoiceCreateConfigurationSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceCreateConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceCreateConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceCreateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceDeleteConfigurationSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceDeleteConfigurationSetFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceDeleteConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceDeleteConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceGetConfigurationSetEventDestinationsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceGetConfigurationSetEventDestinationsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceListConfigurationSetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceListConfigurationSetsFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceSendVoiceMessageFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceSendVoiceMessageFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PinpointSMSVoiceUpdateConfigurationSetEventDestinationFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PinpointSMSVoiceUpdateConfigurationSetEventDestinationFuture) Get(ctx workflow.Context) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) (*pinpointsmsvoice.CreateConfigurationSetOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetInput) *PinpointSMSVoiceCreateConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSet", input)
	return &PinpointSMSVoiceCreateConfigurationSetFuture{Future: future}
}

func (a *stub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.CreateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.CreateConfigurationSetEventDestinationInput) *PinpointSMSVoiceCreateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.CreateConfigurationSetEventDestination", input)
	return &PinpointSMSVoiceCreateConfigurationSetEventDestinationFuture{Future: future}
}

func (a *stub) DeleteConfigurationSet(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) (*pinpointsmsvoice.DeleteConfigurationSetOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigurationSetAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetInput) *PinpointSMSVoiceDeleteConfigurationSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSet", input)
	return &PinpointSMSVoiceDeleteConfigurationSetFuture{Future: future}
}

func (a *stub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) (*pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.DeleteConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.DeleteConfigurationSetEventDestinationInput) *PinpointSMSVoiceDeleteConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.DeleteConfigurationSetEventDestination", input)
	return &PinpointSMSVoiceDeleteConfigurationSetEventDestinationFuture{Future: future}
}

func (a *stub) GetConfigurationSetEventDestinations(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) (*pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput, error) {
	var output pinpointsmsvoice.GetConfigurationSetEventDestinationsOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.GetConfigurationSetEventDestinations", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetConfigurationSetEventDestinationsAsync(ctx workflow.Context, input *pinpointsmsvoice.GetConfigurationSetEventDestinationsInput) *PinpointSMSVoiceGetConfigurationSetEventDestinationsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.GetConfigurationSetEventDestinations", input)
	return &PinpointSMSVoiceGetConfigurationSetEventDestinationsFuture{Future: future}
}

func (a *stub) ListConfigurationSets(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) (*pinpointsmsvoice.ListConfigurationSetsOutput, error) {
	var output pinpointsmsvoice.ListConfigurationSetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.ListConfigurationSets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListConfigurationSetsAsync(ctx workflow.Context, input *pinpointsmsvoice.ListConfigurationSetsInput) *PinpointSMSVoiceListConfigurationSetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.ListConfigurationSets", input)
	return &PinpointSMSVoiceListConfigurationSetsFuture{Future: future}
}

func (a *stub) SendVoiceMessage(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) (*pinpointsmsvoice.SendVoiceMessageOutput, error) {
	var output pinpointsmsvoice.SendVoiceMessageOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.SendVoiceMessage", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SendVoiceMessageAsync(ctx workflow.Context, input *pinpointsmsvoice.SendVoiceMessageInput) *PinpointSMSVoiceSendVoiceMessageFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.SendVoiceMessage", input)
	return &PinpointSMSVoiceSendVoiceMessageFuture{Future: future}
}

func (a *stub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) (*pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput, error) {
	var output pinpointsmsvoice.UpdateConfigurationSetEventDestinationOutput
	err := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.UpdateConfigurationSetEventDestination", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *pinpointsmsvoice.UpdateConfigurationSetEventDestinationInput) *PinpointSMSVoiceUpdateConfigurationSetEventDestinationFuture {
	future := workflow.ExecuteActivity(ctx, "aws.pinpointsmsvoice.UpdateConfigurationSetEventDestination", input)
	return &PinpointSMSVoiceUpdateConfigurationSetEventDestinationFuture{Future: future}
}
