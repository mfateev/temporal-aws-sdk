package awsclients

import (
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ElasticBeanstalkClient interface {
       AbortEnvironmentUpdate(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error)
       AbortEnvironmentUpdateAsync(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) *ElasticbeanstalkAbortEnvironmentUpdateResult

       ApplyEnvironmentManagedAction(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error)
       ApplyEnvironmentManagedActionAsync(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) *ElasticbeanstalkApplyEnvironmentManagedActionResult

       AssociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error)
       AssociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) *ElasticbeanstalkAssociateEnvironmentOperationsRoleResult

       CheckDNSAvailability(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error)
       CheckDNSAvailabilityAsync(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) *ElasticbeanstalkCheckDNSAvailabilityResult

       ComposeEnvironments(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
       ComposeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) *ElasticbeanstalkComposeEnvironmentsResult

       CreateApplication(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
       CreateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) *ElasticbeanstalkCreateApplicationResult

       CreateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
       CreateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) *ElasticbeanstalkCreateApplicationVersionResult

       CreateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
       CreateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) *ElasticbeanstalkCreateConfigurationTemplateResult

       CreateEnvironment(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
       CreateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) *ElasticbeanstalkCreateEnvironmentResult

       CreatePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error)
       CreatePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) *ElasticbeanstalkCreatePlatformVersionResult

       CreateStorageLocation(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error)
       CreateStorageLocationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) *ElasticbeanstalkCreateStorageLocationResult

       DeleteApplication(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error)
       DeleteApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) *ElasticbeanstalkDeleteApplicationResult

       DeleteApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error)
       DeleteApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) *ElasticbeanstalkDeleteApplicationVersionResult

       DeleteConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error)
       DeleteConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) *ElasticbeanstalkDeleteConfigurationTemplateResult

       DeleteEnvironmentConfiguration(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error)
       DeleteEnvironmentConfigurationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) *ElasticbeanstalkDeleteEnvironmentConfigurationResult

       DeletePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error)
       DeletePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) *ElasticbeanstalkDeletePlatformVersionResult

       DescribeAccountAttributes(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error)
       DescribeAccountAttributesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) *ElasticbeanstalkDescribeAccountAttributesResult

       DescribeApplicationVersions(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error)
       DescribeApplicationVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) *ElasticbeanstalkDescribeApplicationVersionsResult

       DescribeApplications(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error)
       DescribeApplicationsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) *ElasticbeanstalkDescribeApplicationsResult

       DescribeConfigurationOptions(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error)
       DescribeConfigurationOptionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) *ElasticbeanstalkDescribeConfigurationOptionsResult

       DescribeConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error)
       DescribeConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) *ElasticbeanstalkDescribeConfigurationSettingsResult

       DescribeEnvironmentHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error)
       DescribeEnvironmentHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) *ElasticbeanstalkDescribeEnvironmentHealthResult

       DescribeEnvironmentManagedActionHistory(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error)
       DescribeEnvironmentManagedActionHistoryAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) *ElasticbeanstalkDescribeEnvironmentManagedActionHistoryResult

       DescribeEnvironmentManagedActions(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error)
       DescribeEnvironmentManagedActionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) *ElasticbeanstalkDescribeEnvironmentManagedActionsResult

       DescribeEnvironmentResources(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error)
       DescribeEnvironmentResourcesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) *ElasticbeanstalkDescribeEnvironmentResourcesResult

       DescribeEnvironments(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error)
       DescribeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *ElasticbeanstalkDescribeEnvironmentsResult

       DescribeEvents(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error)
       DescribeEventsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) *ElasticbeanstalkDescribeEventsResult

       DescribeInstancesHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error)
       DescribeInstancesHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) *ElasticbeanstalkDescribeInstancesHealthResult

       DescribePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error)
       DescribePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) *ElasticbeanstalkDescribePlatformVersionResult

       DisassociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error)
       DisassociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) *ElasticbeanstalkDisassociateEnvironmentOperationsRoleResult

       ListAvailableSolutionStacks(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error)
       ListAvailableSolutionStacksAsync(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) *ElasticbeanstalkListAvailableSolutionStacksResult

       ListPlatformBranches(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error)
       ListPlatformBranchesAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) *ElasticbeanstalkListPlatformBranchesResult

       ListPlatformVersions(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error)
       ListPlatformVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) *ElasticbeanstalkListPlatformVersionsResult

       ListTagsForResource(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) *ElasticbeanstalkListTagsForResourceResult

       RebuildEnvironment(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error)
       RebuildEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) *ElasticbeanstalkRebuildEnvironmentResult

       RequestEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error)
       RequestEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) *ElasticbeanstalkRequestEnvironmentInfoResult

       RestartAppServer(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error)
       RestartAppServerAsync(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) *ElasticbeanstalkRestartAppServerResult

       RetrieveEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error)
       RetrieveEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) *ElasticbeanstalkRetrieveEnvironmentInfoResult

       SwapEnvironmentCNAMEs(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error)
       SwapEnvironmentCNAMEsAsync(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) *ElasticbeanstalkSwapEnvironmentCNAMEsResult

       TerminateEnvironment(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
       TerminateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) *ElasticbeanstalkTerminateEnvironmentResult

       UpdateApplication(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error)
       UpdateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) *ElasticbeanstalkUpdateApplicationResult

       UpdateApplicationResourceLifecycle(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error)
       UpdateApplicationResourceLifecycleAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) *ElasticbeanstalkUpdateApplicationResourceLifecycleResult

       UpdateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error)
       UpdateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) *ElasticbeanstalkUpdateApplicationVersionResult

       UpdateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error)
       UpdateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) *ElasticbeanstalkUpdateConfigurationTemplateResult

       UpdateEnvironment(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error)
       UpdateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) *ElasticbeanstalkUpdateEnvironmentResult

       UpdateTagsForResource(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error)
       UpdateTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) *ElasticbeanstalkUpdateTagsForResourceResult

       ValidateConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error)
       ValidateConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) *ElasticbeanstalkValidateConfigurationSettingsResult

       WaitUntilEnvironmentExists(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error
       WaitUntilEnvironmentTerminated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error
       WaitUntilEnvironmentUpdated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error}

type ElasticbeanstalkAbortEnvironmentUpdateResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkAbortEnvironmentUpdateResult) Get(ctx workflow.Context) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
    var output elasticbeanstalk.AbortEnvironmentUpdateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkApplyEnvironmentManagedActionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkApplyEnvironmentManagedActionResult) Get(ctx workflow.Context) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error) {
    var output elasticbeanstalk.ApplyEnvironmentManagedActionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkAssociateEnvironmentOperationsRoleResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkAssociateEnvironmentOperationsRoleResult) Get(ctx workflow.Context) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error) {
    var output elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCheckDNSAvailabilityResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCheckDNSAvailabilityResult) Get(ctx workflow.Context) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error) {
    var output elasticbeanstalk.CheckDNSAvailabilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkComposeEnvironmentsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkComposeEnvironmentsResult) Get(ctx workflow.Context) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    var output elasticbeanstalk.EnvironmentDescriptionsMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreateApplicationResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreateApplicationResult) Get(ctx workflow.Context) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationDescriptionMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreateApplicationVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreateApplicationVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationVersionDescriptionMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreateConfigurationTemplateResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreateConfigurationTemplateResult) Get(ctx workflow.Context) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    var output elasticbeanstalk.ConfigurationSettingsDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreateEnvironmentResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreateEnvironmentResult) Get(ctx workflow.Context) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreatePlatformVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreatePlatformVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.CreatePlatformVersionOutput, error) {
    var output elasticbeanstalk.CreatePlatformVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkCreateStorageLocationResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkCreateStorageLocationResult) Get(ctx workflow.Context) (*elasticbeanstalk.CreateStorageLocationOutput, error) {
    var output elasticbeanstalk.CreateStorageLocationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDeleteApplicationResult) Get(ctx workflow.Context) (*elasticbeanstalk.DeleteApplicationOutput, error) {
    var output elasticbeanstalk.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDeleteApplicationVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDeleteApplicationVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.DeleteApplicationVersionOutput, error) {
    var output elasticbeanstalk.DeleteApplicationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDeleteConfigurationTemplateResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDeleteConfigurationTemplateResult) Get(ctx workflow.Context) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error) {
    var output elasticbeanstalk.DeleteConfigurationTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDeleteEnvironmentConfigurationResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDeleteEnvironmentConfigurationResult) Get(ctx workflow.Context) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error) {
    var output elasticbeanstalk.DeleteEnvironmentConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDeletePlatformVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDeletePlatformVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.DeletePlatformVersionOutput, error) {
    var output elasticbeanstalk.DeletePlatformVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeAccountAttributesResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeAccountAttributesResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
    var output elasticbeanstalk.DescribeAccountAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeApplicationVersionsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeApplicationVersionsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
    var output elasticbeanstalk.DescribeApplicationVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeApplicationsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeApplicationsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
    var output elasticbeanstalk.DescribeApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeConfigurationOptionsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeConfigurationOptionsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
    var output elasticbeanstalk.DescribeConfigurationOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeConfigurationSettingsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeConfigurationSettingsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
    var output elasticbeanstalk.DescribeConfigurationSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEnvironmentHealthResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEnvironmentHealthResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentHealthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEnvironmentManagedActionHistoryResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEnvironmentManagedActionHistoryResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEnvironmentManagedActionsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEnvironmentManagedActionsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentManagedActionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEnvironmentResourcesResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEnvironmentResourcesResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEnvironmentsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEnvironmentsResult) Get(ctx workflow.Context) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    var output elasticbeanstalk.EnvironmentDescriptionsMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeEventsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeEventsResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeEventsOutput, error) {
    var output elasticbeanstalk.DescribeEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribeInstancesHealthResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribeInstancesHealthResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
    var output elasticbeanstalk.DescribeInstancesHealthOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDescribePlatformVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDescribePlatformVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
    var output elasticbeanstalk.DescribePlatformVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkDisassociateEnvironmentOperationsRoleResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkDisassociateEnvironmentOperationsRoleResult) Get(ctx workflow.Context) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error) {
    var output elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkListAvailableSolutionStacksResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkListAvailableSolutionStacksResult) Get(ctx workflow.Context) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
    var output elasticbeanstalk.ListAvailableSolutionStacksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkListPlatformBranchesResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkListPlatformBranchesResult) Get(ctx workflow.Context) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
    var output elasticbeanstalk.ListPlatformBranchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkListPlatformVersionsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkListPlatformVersionsResult) Get(ctx workflow.Context) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
    var output elasticbeanstalk.ListPlatformVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkListTagsForResourceResult) Get(ctx workflow.Context) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
    var output elasticbeanstalk.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkRebuildEnvironmentResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkRebuildEnvironmentResult) Get(ctx workflow.Context) (*elasticbeanstalk.RebuildEnvironmentOutput, error) {
    var output elasticbeanstalk.RebuildEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkRequestEnvironmentInfoResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkRequestEnvironmentInfoResult) Get(ctx workflow.Context) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error) {
    var output elasticbeanstalk.RequestEnvironmentInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkRestartAppServerResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkRestartAppServerResult) Get(ctx workflow.Context) (*elasticbeanstalk.RestartAppServerOutput, error) {
    var output elasticbeanstalk.RestartAppServerOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkRetrieveEnvironmentInfoResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkRetrieveEnvironmentInfoResult) Get(ctx workflow.Context) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error) {
    var output elasticbeanstalk.RetrieveEnvironmentInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkSwapEnvironmentCNAMEsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkSwapEnvironmentCNAMEsResult) Get(ctx workflow.Context) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error) {
    var output elasticbeanstalk.SwapEnvironmentCNAMEsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkTerminateEnvironmentResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkTerminateEnvironmentResult) Get(ctx workflow.Context) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateApplicationResult) Get(ctx workflow.Context) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationDescriptionMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateApplicationResourceLifecycleResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateApplicationResourceLifecycleResult) Get(ctx workflow.Context) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error) {
    var output elasticbeanstalk.UpdateApplicationResourceLifecycleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateApplicationVersionResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateApplicationVersionResult) Get(ctx workflow.Context) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationVersionDescriptionMessage
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateConfigurationTemplateResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateConfigurationTemplateResult) Get(ctx workflow.Context) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    var output elasticbeanstalk.ConfigurationSettingsDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateEnvironmentResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateEnvironmentResult) Get(ctx workflow.Context) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkUpdateTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkUpdateTagsForResourceResult) Get(ctx workflow.Context) (*elasticbeanstalk.UpdateTagsForResourceOutput, error) {
    var output elasticbeanstalk.UpdateTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticbeanstalkValidateConfigurationSettingsResult struct {
	Result workflow.Future
}

func (r *ElasticbeanstalkValidateConfigurationSettingsResult) Get(ctx workflow.Context) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error) {
    var output elasticbeanstalk.ValidateConfigurationSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ElasticBeanstalkStub struct {
    activities awsactivities.ElasticBeanstalkActivities
}

func NewElasticBeanstalkStub() ElasticBeanstalkClient {
    return &ElasticBeanstalkStub{}
}

func (a *ElasticBeanstalkStub) AbortEnvironmentUpdate(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
    var output elasticbeanstalk.AbortEnvironmentUpdateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AbortEnvironmentUpdate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) AbortEnvironmentUpdateAsync(ctx workflow.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) *ElasticbeanstalkAbortEnvironmentUpdateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AbortEnvironmentUpdate, input)
    return &ElasticbeanstalkAbortEnvironmentUpdateResult{Result: future}
}

func (a *ElasticBeanstalkStub) ApplyEnvironmentManagedAction(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error) {
    var output elasticbeanstalk.ApplyEnvironmentManagedActionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ApplyEnvironmentManagedAction, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ApplyEnvironmentManagedActionAsync(ctx workflow.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) *ElasticbeanstalkApplyEnvironmentManagedActionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ApplyEnvironmentManagedAction, input)
    return &ElasticbeanstalkApplyEnvironmentManagedActionResult{Result: future}
}

func (a *ElasticBeanstalkStub) AssociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error) {
    var output elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateEnvironmentOperationsRole, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) AssociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) *ElasticbeanstalkAssociateEnvironmentOperationsRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateEnvironmentOperationsRole, input)
    return &ElasticbeanstalkAssociateEnvironmentOperationsRoleResult{Result: future}
}

func (a *ElasticBeanstalkStub) CheckDNSAvailability(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error) {
    var output elasticbeanstalk.CheckDNSAvailabilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CheckDNSAvailability, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CheckDNSAvailabilityAsync(ctx workflow.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) *ElasticbeanstalkCheckDNSAvailabilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CheckDNSAvailability, input)
    return &ElasticbeanstalkCheckDNSAvailabilityResult{Result: future}
}

func (a *ElasticBeanstalkStub) ComposeEnvironments(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    var output elasticbeanstalk.EnvironmentDescriptionsMessage
    err := workflow.ExecuteActivity(ctx, a.activities.ComposeEnvironments, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ComposeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) *ElasticbeanstalkComposeEnvironmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ComposeEnvironments, input)
    return &ElasticbeanstalkComposeEnvironmentsResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreateApplication(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationDescriptionMessage
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationInput) *ElasticbeanstalkCreateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input)
    return &ElasticbeanstalkCreateApplicationResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationVersionDescriptionMessage
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreateApplicationVersionInput) *ElasticbeanstalkCreateApplicationVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApplicationVersion, input)
    return &ElasticbeanstalkCreateApplicationVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    var output elasticbeanstalk.ConfigurationSettingsDescription
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) *ElasticbeanstalkCreateConfigurationTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationTemplate, input)
    return &ElasticbeanstalkCreateConfigurationTemplateResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreateEnvironment(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.CreateEnvironmentInput) *ElasticbeanstalkCreateEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironment, input)
    return &ElasticbeanstalkCreateEnvironmentResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreatePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error) {
    var output elasticbeanstalk.CreatePlatformVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePlatformVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreatePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.CreatePlatformVersionInput) *ElasticbeanstalkCreatePlatformVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreatePlatformVersion, input)
    return &ElasticbeanstalkCreatePlatformVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) CreateStorageLocation(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error) {
    var output elasticbeanstalk.CreateStorageLocationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateStorageLocation, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) CreateStorageLocationAsync(ctx workflow.Context, input *elasticbeanstalk.CreateStorageLocationInput) *ElasticbeanstalkCreateStorageLocationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateStorageLocation, input)
    return &ElasticbeanstalkCreateStorageLocationResult{Result: future}
}

func (a *ElasticBeanstalkStub) DeleteApplication(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error) {
    var output elasticbeanstalk.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DeleteApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationInput) *ElasticbeanstalkDeleteApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input)
    return &ElasticbeanstalkDeleteApplicationResult{Result: future}
}

func (a *ElasticBeanstalkStub) DeleteApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error) {
    var output elasticbeanstalk.DeleteApplicationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DeleteApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) *ElasticbeanstalkDeleteApplicationVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApplicationVersion, input)
    return &ElasticbeanstalkDeleteApplicationVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) DeleteConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error) {
    var output elasticbeanstalk.DeleteConfigurationTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DeleteConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) *ElasticbeanstalkDeleteConfigurationTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationTemplate, input)
    return &ElasticbeanstalkDeleteConfigurationTemplateResult{Result: future}
}

func (a *ElasticBeanstalkStub) DeleteEnvironmentConfiguration(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error) {
    var output elasticbeanstalk.DeleteEnvironmentConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironmentConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DeleteEnvironmentConfigurationAsync(ctx workflow.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) *ElasticbeanstalkDeleteEnvironmentConfigurationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironmentConfiguration, input)
    return &ElasticbeanstalkDeleteEnvironmentConfigurationResult{Result: future}
}

func (a *ElasticBeanstalkStub) DeletePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error) {
    var output elasticbeanstalk.DeletePlatformVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePlatformVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DeletePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DeletePlatformVersionInput) *ElasticbeanstalkDeletePlatformVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeletePlatformVersion, input)
    return &ElasticbeanstalkDeletePlatformVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeAccountAttributes(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
    var output elasticbeanstalk.DescribeAccountAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeAccountAttributesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) *ElasticbeanstalkDescribeAccountAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeAccountAttributes, input)
    return &ElasticbeanstalkDescribeAccountAttributesResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeApplicationVersions(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
    var output elasticbeanstalk.DescribeApplicationVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicationVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeApplicationVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) *ElasticbeanstalkDescribeApplicationVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplicationVersions, input)
    return &ElasticbeanstalkDescribeApplicationVersionsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeApplications(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
    var output elasticbeanstalk.DescribeApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeApplicationsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeApplicationsInput) *ElasticbeanstalkDescribeApplicationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeApplications, input)
    return &ElasticbeanstalkDescribeApplicationsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeConfigurationOptions(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
    var output elasticbeanstalk.DescribeConfigurationOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeConfigurationOptionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) *ElasticbeanstalkDescribeConfigurationOptionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationOptions, input)
    return &ElasticbeanstalkDescribeConfigurationOptionsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
    var output elasticbeanstalk.DescribeConfigurationSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) *ElasticbeanstalkDescribeConfigurationSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationSettings, input)
    return &ElasticbeanstalkDescribeConfigurationSettingsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentHealthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentHealth, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) *ElasticbeanstalkDescribeEnvironmentHealthResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentHealth, input)
    return &ElasticbeanstalkDescribeEnvironmentHealthResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentManagedActionHistory(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentManagedActionHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentManagedActionHistoryAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) *ElasticbeanstalkDescribeEnvironmentManagedActionHistoryResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentManagedActionHistory, input)
    return &ElasticbeanstalkDescribeEnvironmentManagedActionHistoryResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentManagedActions(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentManagedActionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentManagedActions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentManagedActionsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) *ElasticbeanstalkDescribeEnvironmentManagedActionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentManagedActions, input)
    return &ElasticbeanstalkDescribeEnvironmentManagedActionsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentResources(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
    var output elasticbeanstalk.DescribeEnvironmentResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentResources, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentResourcesAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) *ElasticbeanstalkDescribeEnvironmentResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironmentResources, input)
    return &ElasticbeanstalkDescribeEnvironmentResourcesResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEnvironments(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
    var output elasticbeanstalk.EnvironmentDescriptionsMessage
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironments, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEnvironmentsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) *ElasticbeanstalkDescribeEnvironmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEnvironments, input)
    return &ElasticbeanstalkDescribeEnvironmentsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeEvents(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error) {
    var output elasticbeanstalk.DescribeEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeEventsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEventsInput) *ElasticbeanstalkDescribeEventsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeEvents, input)
    return &ElasticbeanstalkDescribeEventsResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribeInstancesHealth(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
    var output elasticbeanstalk.DescribeInstancesHealthOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeInstancesHealth, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribeInstancesHealthAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) *ElasticbeanstalkDescribeInstancesHealthResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeInstancesHealth, input)
    return &ElasticbeanstalkDescribeInstancesHealthResult{Result: future}
}

func (a *ElasticBeanstalkStub) DescribePlatformVersion(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
    var output elasticbeanstalk.DescribePlatformVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePlatformVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DescribePlatformVersionAsync(ctx workflow.Context, input *elasticbeanstalk.DescribePlatformVersionInput) *ElasticbeanstalkDescribePlatformVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribePlatformVersion, input)
    return &ElasticbeanstalkDescribePlatformVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) DisassociateEnvironmentOperationsRole(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error) {
    var output elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateEnvironmentOperationsRole, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) DisassociateEnvironmentOperationsRoleAsync(ctx workflow.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) *ElasticbeanstalkDisassociateEnvironmentOperationsRoleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateEnvironmentOperationsRole, input)
    return &ElasticbeanstalkDisassociateEnvironmentOperationsRoleResult{Result: future}
}

func (a *ElasticBeanstalkStub) ListAvailableSolutionStacks(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
    var output elasticbeanstalk.ListAvailableSolutionStacksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListAvailableSolutionStacks, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ListAvailableSolutionStacksAsync(ctx workflow.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) *ElasticbeanstalkListAvailableSolutionStacksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListAvailableSolutionStacks, input)
    return &ElasticbeanstalkListAvailableSolutionStacksResult{Result: future}
}

func (a *ElasticBeanstalkStub) ListPlatformBranches(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
    var output elasticbeanstalk.ListPlatformBranchesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlatformBranches, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ListPlatformBranchesAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformBranchesInput) *ElasticbeanstalkListPlatformBranchesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPlatformBranches, input)
    return &ElasticbeanstalkListPlatformBranchesResult{Result: future}
}

func (a *ElasticBeanstalkStub) ListPlatformVersions(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
    var output elasticbeanstalk.ListPlatformVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPlatformVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ListPlatformVersionsAsync(ctx workflow.Context, input *elasticbeanstalk.ListPlatformVersionsInput) *ElasticbeanstalkListPlatformVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPlatformVersions, input)
    return &ElasticbeanstalkListPlatformVersionsResult{Result: future}
}

func (a *ElasticBeanstalkStub) ListTagsForResource(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
    var output elasticbeanstalk.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ListTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.ListTagsForResourceInput) *ElasticbeanstalkListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ElasticbeanstalkListTagsForResourceResult{Result: future}
}

func (a *ElasticBeanstalkStub) RebuildEnvironment(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error) {
    var output elasticbeanstalk.RebuildEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RebuildEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) RebuildEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.RebuildEnvironmentInput) *ElasticbeanstalkRebuildEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RebuildEnvironment, input)
    return &ElasticbeanstalkRebuildEnvironmentResult{Result: future}
}

func (a *ElasticBeanstalkStub) RequestEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error) {
    var output elasticbeanstalk.RequestEnvironmentInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestEnvironmentInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) RequestEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) *ElasticbeanstalkRequestEnvironmentInfoResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RequestEnvironmentInfo, input)
    return &ElasticbeanstalkRequestEnvironmentInfoResult{Result: future}
}

func (a *ElasticBeanstalkStub) RestartAppServer(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error) {
    var output elasticbeanstalk.RestartAppServerOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RestartAppServer, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) RestartAppServerAsync(ctx workflow.Context, input *elasticbeanstalk.RestartAppServerInput) *ElasticbeanstalkRestartAppServerResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RestartAppServer, input)
    return &ElasticbeanstalkRestartAppServerResult{Result: future}
}

func (a *ElasticBeanstalkStub) RetrieveEnvironmentInfo(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error) {
    var output elasticbeanstalk.RetrieveEnvironmentInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RetrieveEnvironmentInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) RetrieveEnvironmentInfoAsync(ctx workflow.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) *ElasticbeanstalkRetrieveEnvironmentInfoResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RetrieveEnvironmentInfo, input)
    return &ElasticbeanstalkRetrieveEnvironmentInfoResult{Result: future}
}

func (a *ElasticBeanstalkStub) SwapEnvironmentCNAMEs(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error) {
    var output elasticbeanstalk.SwapEnvironmentCNAMEsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SwapEnvironmentCNAMEs, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) SwapEnvironmentCNAMEsAsync(ctx workflow.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) *ElasticbeanstalkSwapEnvironmentCNAMEsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SwapEnvironmentCNAMEs, input)
    return &ElasticbeanstalkSwapEnvironmentCNAMEsResult{Result: future}
}

func (a *ElasticBeanstalkStub) TerminateEnvironment(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) TerminateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.TerminateEnvironmentInput) *ElasticbeanstalkTerminateEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TerminateEnvironment, input)
    return &ElasticbeanstalkTerminateEnvironmentResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateApplication(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationDescriptionMessage
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateApplicationAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationInput) *ElasticbeanstalkUpdateApplicationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input)
    return &ElasticbeanstalkUpdateApplicationResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateApplicationResourceLifecycle(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error) {
    var output elasticbeanstalk.UpdateApplicationResourceLifecycleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationResourceLifecycle, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateApplicationResourceLifecycleAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) *ElasticbeanstalkUpdateApplicationResourceLifecycleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationResourceLifecycle, input)
    return &ElasticbeanstalkUpdateApplicationResourceLifecycleResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateApplicationVersion(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
    var output elasticbeanstalk.ApplicationVersionDescriptionMessage
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateApplicationVersionAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) *ElasticbeanstalkUpdateApplicationVersionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApplicationVersion, input)
    return &ElasticbeanstalkUpdateApplicationVersionResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateConfigurationTemplate(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
    var output elasticbeanstalk.ConfigurationSettingsDescription
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateConfigurationTemplateAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) *ElasticbeanstalkUpdateConfigurationTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationTemplate, input)
    return &ElasticbeanstalkUpdateConfigurationTemplateResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateEnvironment(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
    var output elasticbeanstalk.EnvironmentDescription
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateEnvironmentAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateEnvironmentInput) *ElasticbeanstalkUpdateEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironment, input)
    return &ElasticbeanstalkUpdateEnvironmentResult{Result: future}
}

func (a *ElasticBeanstalkStub) UpdateTagsForResource(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error) {
    var output elasticbeanstalk.UpdateTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) UpdateTagsForResourceAsync(ctx workflow.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) *ElasticbeanstalkUpdateTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateTagsForResource, input)
    return &ElasticbeanstalkUpdateTagsForResourceResult{Result: future}
}

func (a *ElasticBeanstalkStub) ValidateConfigurationSettings(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error) {
    var output elasticbeanstalk.ValidateConfigurationSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateConfigurationSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *ElasticBeanstalkStub) ValidateConfigurationSettingsAsync(ctx workflow.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) *ElasticbeanstalkValidateConfigurationSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ValidateConfigurationSettings, input)
    return &ElasticbeanstalkValidateConfigurationSettingsResult{Result: future}
}

func (a *ElasticBeanstalkStub) WaitUntilEnvironmentExists(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentExists, input).Get(ctx, nil)
}

func (a *ElasticBeanstalkStub) WaitUntilEnvironmentExistsAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentExists, input)
}


func (a *ElasticBeanstalkStub) WaitUntilEnvironmentTerminated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentTerminated, input).Get(ctx, nil)
}

func (a *ElasticBeanstalkStub) WaitUntilEnvironmentTerminatedAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentTerminated, input)
}


func (a *ElasticBeanstalkStub) WaitUntilEnvironmentUpdated(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentUpdated, input).Get(ctx, nil)
}

func (a *ElasticBeanstalkStub) WaitUntilEnvironmentUpdatedAsync(ctx workflow.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) workflow.Future {
    return workflow.ExecuteActivity(ctx, a.activities.WaitUntilEnvironmentUpdated, input)
}
