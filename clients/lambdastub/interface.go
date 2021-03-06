// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package lambdastub

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddLayerVersionPermission(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) (*lambda.AddLayerVersionPermissionOutput, error)
	AddLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.AddLayerVersionPermissionInput) *LambdaAddLayerVersionPermissionFuture

	AddPermission(ctx workflow.Context, input *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error)
	AddPermissionAsync(ctx workflow.Context, input *lambda.AddPermissionInput) *LambdaAddPermissionFuture

	CreateAlias(ctx workflow.Context, input *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error)
	CreateAliasAsync(ctx workflow.Context, input *lambda.CreateAliasInput) *LambdaCreateAliasFuture

	CreateEventSourceMapping(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	CreateEventSourceMappingAsync(ctx workflow.Context, input *lambda.CreateEventSourceMappingInput) *LambdaCreateEventSourceMappingFuture

	CreateFunction(ctx workflow.Context, input *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error)
	CreateFunctionAsync(ctx workflow.Context, input *lambda.CreateFunctionInput) *LambdaCreateFunctionFuture

	DeleteAlias(ctx workflow.Context, input *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error)
	DeleteAliasAsync(ctx workflow.Context, input *lambda.DeleteAliasInput) *LambdaDeleteAliasFuture

	DeleteEventSourceMapping(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	DeleteEventSourceMappingAsync(ctx workflow.Context, input *lambda.DeleteEventSourceMappingInput) *LambdaDeleteEventSourceMappingFuture

	DeleteFunction(ctx workflow.Context, input *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error)
	DeleteFunctionAsync(ctx workflow.Context, input *lambda.DeleteFunctionInput) *LambdaDeleteFunctionFuture

	DeleteFunctionConcurrency(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) (*lambda.DeleteFunctionConcurrencyOutput, error)
	DeleteFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.DeleteFunctionConcurrencyInput) *LambdaDeleteFunctionConcurrencyFuture

	DeleteFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) (*lambda.DeleteFunctionEventInvokeConfigOutput, error)
	DeleteFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.DeleteFunctionEventInvokeConfigInput) *LambdaDeleteFunctionEventInvokeConfigFuture

	DeleteLayerVersion(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) (*lambda.DeleteLayerVersionOutput, error)
	DeleteLayerVersionAsync(ctx workflow.Context, input *lambda.DeleteLayerVersionInput) *LambdaDeleteLayerVersionFuture

	DeleteProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) (*lambda.DeleteProvisionedConcurrencyConfigOutput, error)
	DeleteProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.DeleteProvisionedConcurrencyConfigInput) *LambdaDeleteProvisionedConcurrencyConfigFuture

	GetAccountSettings(ctx workflow.Context, input *lambda.GetAccountSettingsInput) (*lambda.GetAccountSettingsOutput, error)
	GetAccountSettingsAsync(ctx workflow.Context, input *lambda.GetAccountSettingsInput) *LambdaGetAccountSettingsFuture

	GetAlias(ctx workflow.Context, input *lambda.GetAliasInput) (*lambda.AliasConfiguration, error)
	GetAliasAsync(ctx workflow.Context, input *lambda.GetAliasInput) *LambdaGetAliasFuture

	GetEventSourceMapping(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	GetEventSourceMappingAsync(ctx workflow.Context, input *lambda.GetEventSourceMappingInput) *LambdaGetEventSourceMappingFuture

	GetFunction(ctx workflow.Context, input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
	GetFunctionAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *LambdaGetFunctionFuture

	GetFunctionConcurrency(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) (*lambda.GetFunctionConcurrencyOutput, error)
	GetFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.GetFunctionConcurrencyInput) *LambdaGetFunctionConcurrencyFuture

	GetFunctionConfiguration(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	GetFunctionConfigurationAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *LambdaGetFunctionConfigurationFuture

	GetFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) (*lambda.GetFunctionEventInvokeConfigOutput, error)
	GetFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.GetFunctionEventInvokeConfigInput) *LambdaGetFunctionEventInvokeConfigFuture

	GetLayerVersion(ctx workflow.Context, input *lambda.GetLayerVersionInput) (*lambda.GetLayerVersionOutput, error)
	GetLayerVersionAsync(ctx workflow.Context, input *lambda.GetLayerVersionInput) *LambdaGetLayerVersionFuture

	GetLayerVersionByArn(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) (*lambda.GetLayerVersionByArnOutput, error)
	GetLayerVersionByArnAsync(ctx workflow.Context, input *lambda.GetLayerVersionByArnInput) *LambdaGetLayerVersionByArnFuture

	GetLayerVersionPolicy(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) (*lambda.GetLayerVersionPolicyOutput, error)
	GetLayerVersionPolicyAsync(ctx workflow.Context, input *lambda.GetLayerVersionPolicyInput) *LambdaGetLayerVersionPolicyFuture

	GetPolicy(ctx workflow.Context, input *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error)
	GetPolicyAsync(ctx workflow.Context, input *lambda.GetPolicyInput) *LambdaGetPolicyFuture

	GetProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) (*lambda.GetProvisionedConcurrencyConfigOutput, error)
	GetProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.GetProvisionedConcurrencyConfigInput) *LambdaGetProvisionedConcurrencyConfigFuture

	Invoke(ctx workflow.Context, input *lambda.InvokeInput) (*lambda.InvokeOutput, error)
	InvokeAsync(ctx workflow.Context, input *lambda.InvokeInput) *LambdaInvokeFuture

	ListAliases(ctx workflow.Context, input *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error)
	ListAliasesAsync(ctx workflow.Context, input *lambda.ListAliasesInput) *LambdaListAliasesFuture

	ListEventSourceMappings(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error)
	ListEventSourceMappingsAsync(ctx workflow.Context, input *lambda.ListEventSourceMappingsInput) *LambdaListEventSourceMappingsFuture

	ListFunctionEventInvokeConfigs(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) (*lambda.ListFunctionEventInvokeConfigsOutput, error)
	ListFunctionEventInvokeConfigsAsync(ctx workflow.Context, input *lambda.ListFunctionEventInvokeConfigsInput) *LambdaListFunctionEventInvokeConfigsFuture

	ListFunctions(ctx workflow.Context, input *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error)
	ListFunctionsAsync(ctx workflow.Context, input *lambda.ListFunctionsInput) *LambdaListFunctionsFuture

	ListLayerVersions(ctx workflow.Context, input *lambda.ListLayerVersionsInput) (*lambda.ListLayerVersionsOutput, error)
	ListLayerVersionsAsync(ctx workflow.Context, input *lambda.ListLayerVersionsInput) *LambdaListLayerVersionsFuture

	ListLayers(ctx workflow.Context, input *lambda.ListLayersInput) (*lambda.ListLayersOutput, error)
	ListLayersAsync(ctx workflow.Context, input *lambda.ListLayersInput) *LambdaListLayersFuture

	ListProvisionedConcurrencyConfigs(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) (*lambda.ListProvisionedConcurrencyConfigsOutput, error)
	ListProvisionedConcurrencyConfigsAsync(ctx workflow.Context, input *lambda.ListProvisionedConcurrencyConfigsInput) *LambdaListProvisionedConcurrencyConfigsFuture

	ListTags(ctx workflow.Context, input *lambda.ListTagsInput) (*lambda.ListTagsOutput, error)
	ListTagsAsync(ctx workflow.Context, input *lambda.ListTagsInput) *LambdaListTagsFuture

	ListVersionsByFunction(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error)
	ListVersionsByFunctionAsync(ctx workflow.Context, input *lambda.ListVersionsByFunctionInput) *LambdaListVersionsByFunctionFuture

	PublishLayerVersion(ctx workflow.Context, input *lambda.PublishLayerVersionInput) (*lambda.PublishLayerVersionOutput, error)
	PublishLayerVersionAsync(ctx workflow.Context, input *lambda.PublishLayerVersionInput) *LambdaPublishLayerVersionFuture

	PublishVersion(ctx workflow.Context, input *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error)
	PublishVersionAsync(ctx workflow.Context, input *lambda.PublishVersionInput) *LambdaPublishVersionFuture

	PutFunctionConcurrency(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) (*lambda.PutFunctionConcurrencyOutput, error)
	PutFunctionConcurrencyAsync(ctx workflow.Context, input *lambda.PutFunctionConcurrencyInput) *LambdaPutFunctionConcurrencyFuture

	PutFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) (*lambda.PutFunctionEventInvokeConfigOutput, error)
	PutFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.PutFunctionEventInvokeConfigInput) *LambdaPutFunctionEventInvokeConfigFuture

	PutProvisionedConcurrencyConfig(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) (*lambda.PutProvisionedConcurrencyConfigOutput, error)
	PutProvisionedConcurrencyConfigAsync(ctx workflow.Context, input *lambda.PutProvisionedConcurrencyConfigInput) *LambdaPutProvisionedConcurrencyConfigFuture

	RemoveLayerVersionPermission(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) (*lambda.RemoveLayerVersionPermissionOutput, error)
	RemoveLayerVersionPermissionAsync(ctx workflow.Context, input *lambda.RemoveLayerVersionPermissionInput) *LambdaRemoveLayerVersionPermissionFuture

	RemovePermission(ctx workflow.Context, input *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error)
	RemovePermissionAsync(ctx workflow.Context, input *lambda.RemovePermissionInput) *LambdaRemovePermissionFuture

	TagResource(ctx workflow.Context, input *lambda.TagResourceInput) (*lambda.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *lambda.TagResourceInput) *LambdaTagResourceFuture

	UntagResource(ctx workflow.Context, input *lambda.UntagResourceInput) (*lambda.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *lambda.UntagResourceInput) *LambdaUntagResourceFuture

	UpdateAlias(ctx workflow.Context, input *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error)
	UpdateAliasAsync(ctx workflow.Context, input *lambda.UpdateAliasInput) *LambdaUpdateAliasFuture

	UpdateEventSourceMapping(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error)
	UpdateEventSourceMappingAsync(ctx workflow.Context, input *lambda.UpdateEventSourceMappingInput) *LambdaUpdateEventSourceMappingFuture

	UpdateFunctionCode(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionCodeAsync(ctx workflow.Context, input *lambda.UpdateFunctionCodeInput) *LambdaUpdateFunctionCodeFuture

	UpdateFunctionConfiguration(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error)
	UpdateFunctionConfigurationAsync(ctx workflow.Context, input *lambda.UpdateFunctionConfigurationInput) *LambdaUpdateFunctionConfigurationFuture

	UpdateFunctionEventInvokeConfig(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) (*lambda.UpdateFunctionEventInvokeConfigOutput, error)
	UpdateFunctionEventInvokeConfigAsync(ctx workflow.Context, input *lambda.UpdateFunctionEventInvokeConfigInput) *LambdaUpdateFunctionEventInvokeConfigFuture

	WaitUntilFunctionActive(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionActiveAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *clients.VoidFuture

	WaitUntilFunctionExists(ctx workflow.Context, input *lambda.GetFunctionInput) error
	WaitUntilFunctionExistsAsync(ctx workflow.Context, input *lambda.GetFunctionInput) *clients.VoidFuture

	WaitUntilFunctionUpdated(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) error
	WaitUntilFunctionUpdatedAsync(ctx workflow.Context, input *lambda.GetFunctionConfigurationInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
