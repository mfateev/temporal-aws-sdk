// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package appstream

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/appstream"
	"github.com/aws/aws-sdk-go/service/appstream/appstreamiface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client appstreamiface.AppStreamAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := appstream.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (appstreamiface.AppStreamAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return appstream.New(sess), nil
}

func (a *Activities) AssociateFleet(ctx context.Context, input *appstream.AssociateFleetInput) (*appstream.AssociateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.AssociateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchAssociateUserStack(ctx context.Context, input *appstream.BatchAssociateUserStackInput) (*appstream.BatchAssociateUserStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchAssociateUserStackWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) BatchDisassociateUserStack(ctx context.Context, input *appstream.BatchDisassociateUserStackInput) (*appstream.BatchDisassociateUserStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.BatchDisassociateUserStackWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CopyImage(ctx context.Context, input *appstream.CopyImageInput) (*appstream.CopyImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CopyImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateDirectoryConfig(ctx context.Context, input *appstream.CreateDirectoryConfigInput) (*appstream.CreateDirectoryConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateDirectoryConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateFleet(ctx context.Context, input *appstream.CreateFleetInput) (*appstream.CreateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateImageBuilder(ctx context.Context, input *appstream.CreateImageBuilderInput) (*appstream.CreateImageBuilderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateImageBuilderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateImageBuilderStreamingURL(ctx context.Context, input *appstream.CreateImageBuilderStreamingURLInput) (*appstream.CreateImageBuilderStreamingURLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateImageBuilderStreamingURLWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStack(ctx context.Context, input *appstream.CreateStackInput) (*appstream.CreateStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStackWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStreamingURL(ctx context.Context, input *appstream.CreateStreamingURLInput) (*appstream.CreateStreamingURLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStreamingURLWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUsageReportSubscription(ctx context.Context, input *appstream.CreateUsageReportSubscriptionInput) (*appstream.CreateUsageReportSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUsageReportSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateUser(ctx context.Context, input *appstream.CreateUserInput) (*appstream.CreateUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteDirectoryConfig(ctx context.Context, input *appstream.DeleteDirectoryConfigInput) (*appstream.DeleteDirectoryConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteDirectoryConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFleet(ctx context.Context, input *appstream.DeleteFleetInput) (*appstream.DeleteFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteImage(ctx context.Context, input *appstream.DeleteImageInput) (*appstream.DeleteImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteImageBuilder(ctx context.Context, input *appstream.DeleteImageBuilderInput) (*appstream.DeleteImageBuilderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteImageBuilderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteImagePermissions(ctx context.Context, input *appstream.DeleteImagePermissionsInput) (*appstream.DeleteImagePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteImagePermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStack(ctx context.Context, input *appstream.DeleteStackInput) (*appstream.DeleteStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStackWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUsageReportSubscription(ctx context.Context, input *appstream.DeleteUsageReportSubscriptionInput) (*appstream.DeleteUsageReportSubscriptionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUsageReportSubscriptionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteUser(ctx context.Context, input *appstream.DeleteUserInput) (*appstream.DeleteUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeDirectoryConfigs(ctx context.Context, input *appstream.DescribeDirectoryConfigsInput) (*appstream.DescribeDirectoryConfigsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeDirectoryConfigsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeFleets(ctx context.Context, input *appstream.DescribeFleetsInput) (*appstream.DescribeFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeFleetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeImageBuilders(ctx context.Context, input *appstream.DescribeImageBuildersInput) (*appstream.DescribeImageBuildersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeImageBuildersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeImagePermissions(ctx context.Context, input *appstream.DescribeImagePermissionsInput) (*appstream.DescribeImagePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeImagePermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeImages(ctx context.Context, input *appstream.DescribeImagesInput) (*appstream.DescribeImagesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeImagesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeSessions(ctx context.Context, input *appstream.DescribeSessionsInput) (*appstream.DescribeSessionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeSessionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStacks(ctx context.Context, input *appstream.DescribeStacksInput) (*appstream.DescribeStacksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStacksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUsageReportSubscriptions(ctx context.Context, input *appstream.DescribeUsageReportSubscriptionsInput) (*appstream.DescribeUsageReportSubscriptionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUsageReportSubscriptionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUserStackAssociations(ctx context.Context, input *appstream.DescribeUserStackAssociationsInput) (*appstream.DescribeUserStackAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUserStackAssociationsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeUsers(ctx context.Context, input *appstream.DescribeUsersInput) (*appstream.DescribeUsersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeUsersWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisableUser(ctx context.Context, input *appstream.DisableUserInput) (*appstream.DisableUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisableUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DisassociateFleet(ctx context.Context, input *appstream.DisassociateFleetInput) (*appstream.DisassociateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DisassociateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) EnableUser(ctx context.Context, input *appstream.EnableUserInput) (*appstream.EnableUserOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.EnableUserWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ExpireSession(ctx context.Context, input *appstream.ExpireSessionInput) (*appstream.ExpireSessionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ExpireSessionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAssociatedFleets(ctx context.Context, input *appstream.ListAssociatedFleetsInput) (*appstream.ListAssociatedFleetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAssociatedFleetsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListAssociatedStacks(ctx context.Context, input *appstream.ListAssociatedStacksInput) (*appstream.ListAssociatedStacksOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListAssociatedStacksWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListTagsForResource(ctx context.Context, input *appstream.ListTagsForResourceInput) (*appstream.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListTagsForResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartFleet(ctx context.Context, input *appstream.StartFleetInput) (*appstream.StartFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartImageBuilder(ctx context.Context, input *appstream.StartImageBuilderInput) (*appstream.StartImageBuilderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartImageBuilderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopFleet(ctx context.Context, input *appstream.StopFleetInput) (*appstream.StopFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopImageBuilder(ctx context.Context, input *appstream.StopImageBuilderInput) (*appstream.StopImageBuilderOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopImageBuilderWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) TagResource(ctx context.Context, input *appstream.TagResourceInput) (*appstream.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.TagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UntagResource(ctx context.Context, input *appstream.UntagResourceInput) (*appstream.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UntagResourceWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateDirectoryConfig(ctx context.Context, input *appstream.UpdateDirectoryConfigInput) (*appstream.UpdateDirectoryConfigOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateDirectoryConfigWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateFleet(ctx context.Context, input *appstream.UpdateFleetInput) (*appstream.UpdateFleetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateFleetWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateImagePermissions(ctx context.Context, input *appstream.UpdateImagePermissionsInput) (*appstream.UpdateImagePermissionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateImagePermissionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) UpdateStack(ctx context.Context, input *appstream.UpdateStackInput) (*appstream.UpdateStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.UpdateStackWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilFleetStarted(ctx context.Context, input *appstream.DescribeFleetsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilFleetStartedWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilFleetStopped(ctx context.Context, input *appstream.DescribeFleetsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilFleetStoppedWithContext(ctx, input, options...))
	})
}
