// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk/elasticbeanstalkiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client elasticbeanstalkiface.ElasticBeanstalkAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := elasticbeanstalk.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (elasticbeanstalkiface.ElasticBeanstalkAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return elasticbeanstalk.New(sess), nil
}

func (a *Activities) AbortEnvironmentUpdate(ctx context.Context, input *elasticbeanstalk.AbortEnvironmentUpdateInput) (*elasticbeanstalk.AbortEnvironmentUpdateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AbortEnvironmentUpdateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ApplyEnvironmentManagedAction(ctx context.Context, input *elasticbeanstalk.ApplyEnvironmentManagedActionInput) (*elasticbeanstalk.ApplyEnvironmentManagedActionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ApplyEnvironmentManagedActionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) AssociateEnvironmentOperationsRole(ctx context.Context, input *elasticbeanstalk.AssociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.AssociateEnvironmentOperationsRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateEnvironmentOperationsRoleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CheckDNSAvailability(ctx context.Context, input *elasticbeanstalk.CheckDNSAvailabilityInput) (*elasticbeanstalk.CheckDNSAvailabilityOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CheckDNSAvailabilityWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ComposeEnvironments(ctx context.Context, input *elasticbeanstalk.ComposeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ComposeEnvironmentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateApplication(ctx context.Context, input *elasticbeanstalk.CreateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateApplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateApplicationVersion(ctx context.Context, input *elasticbeanstalk.CreateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateApplicationVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.CreateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateConfigurationTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateEnvironment(ctx context.Context, input *elasticbeanstalk.CreateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateEnvironmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreatePlatformVersion(ctx context.Context, input *elasticbeanstalk.CreatePlatformVersionInput) (*elasticbeanstalk.CreatePlatformVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreatePlatformVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStorageLocation(ctx context.Context, input *elasticbeanstalk.CreateStorageLocationInput) (*elasticbeanstalk.CreateStorageLocationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStorageLocationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteApplication(ctx context.Context, input *elasticbeanstalk.DeleteApplicationInput) (*elasticbeanstalk.DeleteApplicationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteApplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteApplicationVersion(ctx context.Context, input *elasticbeanstalk.DeleteApplicationVersionInput) (*elasticbeanstalk.DeleteApplicationVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteApplicationVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.DeleteConfigurationTemplateInput) (*elasticbeanstalk.DeleteConfigurationTemplateOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteConfigurationTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteEnvironmentConfiguration(ctx context.Context, input *elasticbeanstalk.DeleteEnvironmentConfigurationInput) (*elasticbeanstalk.DeleteEnvironmentConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteEnvironmentConfigurationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeletePlatformVersion(ctx context.Context, input *elasticbeanstalk.DeletePlatformVersionInput) (*elasticbeanstalk.DeletePlatformVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeletePlatformVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeAccountAttributes(ctx context.Context, input *elasticbeanstalk.DescribeAccountAttributesInput) (*elasticbeanstalk.DescribeAccountAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeAccountAttributesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeApplicationVersions(ctx context.Context, input *elasticbeanstalk.DescribeApplicationVersionsInput) (*elasticbeanstalk.DescribeApplicationVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeApplicationVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeApplications(ctx context.Context, input *elasticbeanstalk.DescribeApplicationsInput) (*elasticbeanstalk.DescribeApplicationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeApplicationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConfigurationOptions(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationOptionsInput) (*elasticbeanstalk.DescribeConfigurationOptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConfigurationOptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeConfigurationSettings(ctx context.Context, input *elasticbeanstalk.DescribeConfigurationSettingsInput) (*elasticbeanstalk.DescribeConfigurationSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeConfigurationSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEnvironmentHealth(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentHealthInput) (*elasticbeanstalk.DescribeEnvironmentHealthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEnvironmentHealthWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEnvironmentManagedActionHistory(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionHistoryInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEnvironmentManagedActionHistoryWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEnvironmentManagedActions(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentManagedActionsInput) (*elasticbeanstalk.DescribeEnvironmentManagedActionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEnvironmentManagedActionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEnvironmentResources(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentResourcesInput) (*elasticbeanstalk.DescribeEnvironmentResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEnvironmentResourcesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEnvironments(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) (*elasticbeanstalk.EnvironmentDescriptionsMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEnvironmentsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeEvents(ctx context.Context, input *elasticbeanstalk.DescribeEventsInput) (*elasticbeanstalk.DescribeEventsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeEventsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeInstancesHealth(ctx context.Context, input *elasticbeanstalk.DescribeInstancesHealthInput) (*elasticbeanstalk.DescribeInstancesHealthOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeInstancesHealthWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribePlatformVersion(ctx context.Context, input *elasticbeanstalk.DescribePlatformVersionInput) (*elasticbeanstalk.DescribePlatformVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribePlatformVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateEnvironmentOperationsRole(ctx context.Context, input *elasticbeanstalk.DisassociateEnvironmentOperationsRoleInput) (*elasticbeanstalk.DisassociateEnvironmentOperationsRoleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateEnvironmentOperationsRoleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAvailableSolutionStacks(ctx context.Context, input *elasticbeanstalk.ListAvailableSolutionStacksInput) (*elasticbeanstalk.ListAvailableSolutionStacksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAvailableSolutionStacksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPlatformBranches(ctx context.Context, input *elasticbeanstalk.ListPlatformBranchesInput) (*elasticbeanstalk.ListPlatformBranchesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPlatformBranchesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListPlatformVersions(ctx context.Context, input *elasticbeanstalk.ListPlatformVersionsInput) (*elasticbeanstalk.ListPlatformVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListPlatformVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *elasticbeanstalk.ListTagsForResourceInput) (*elasticbeanstalk.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RebuildEnvironment(ctx context.Context, input *elasticbeanstalk.RebuildEnvironmentInput) (*elasticbeanstalk.RebuildEnvironmentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RebuildEnvironmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RequestEnvironmentInfo(ctx context.Context, input *elasticbeanstalk.RequestEnvironmentInfoInput) (*elasticbeanstalk.RequestEnvironmentInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RequestEnvironmentInfoWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RestartAppServer(ctx context.Context, input *elasticbeanstalk.RestartAppServerInput) (*elasticbeanstalk.RestartAppServerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RestartAppServerWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RetrieveEnvironmentInfo(ctx context.Context, input *elasticbeanstalk.RetrieveEnvironmentInfoInput) (*elasticbeanstalk.RetrieveEnvironmentInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RetrieveEnvironmentInfoWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SwapEnvironmentCNAMEs(ctx context.Context, input *elasticbeanstalk.SwapEnvironmentCNAMEsInput) (*elasticbeanstalk.SwapEnvironmentCNAMEsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SwapEnvironmentCNAMEsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TerminateEnvironment(ctx context.Context, input *elasticbeanstalk.TerminateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TerminateEnvironmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateApplication(ctx context.Context, input *elasticbeanstalk.UpdateApplicationInput) (*elasticbeanstalk.ApplicationDescriptionMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateApplicationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateApplicationResourceLifecycle(ctx context.Context, input *elasticbeanstalk.UpdateApplicationResourceLifecycleInput) (*elasticbeanstalk.UpdateApplicationResourceLifecycleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateApplicationResourceLifecycleWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateApplicationVersion(ctx context.Context, input *elasticbeanstalk.UpdateApplicationVersionInput) (*elasticbeanstalk.ApplicationVersionDescriptionMessage, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateApplicationVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateConfigurationTemplate(ctx context.Context, input *elasticbeanstalk.UpdateConfigurationTemplateInput) (*elasticbeanstalk.ConfigurationSettingsDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateConfigurationTemplateWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateEnvironment(ctx context.Context, input *elasticbeanstalk.UpdateEnvironmentInput) (*elasticbeanstalk.EnvironmentDescription, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateEnvironmentWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateTagsForResource(ctx context.Context, input *elasticbeanstalk.UpdateTagsForResourceInput) (*elasticbeanstalk.UpdateTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ValidateConfigurationSettings(ctx context.Context, input *elasticbeanstalk.ValidateConfigurationSettingsInput) (*elasticbeanstalk.ValidateConfigurationSettingsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ValidateConfigurationSettingsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilEnvironmentExists(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilEnvironmentExistsWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilEnvironmentTerminated(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilEnvironmentTerminatedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilEnvironmentUpdated(ctx context.Context, input *elasticbeanstalk.DescribeEnvironmentsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilEnvironmentUpdatedWithContext(ctx, input, options...))
	})
}
