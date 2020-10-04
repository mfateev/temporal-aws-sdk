// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestar"
	"github.com/aws/aws-sdk-go/service/codestar/codestariface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type CodeStarActivities struct {
	client codestariface.CodeStarAPI

	sessionFactory SessionFactory
}

func NewCodeStarActivities(sess *session.Session, config ...*aws.Config) *CodeStarActivities {
	client := codestar.New(sess, config...)
	return &CodeStarActivities{client: client}
}

func NewCodeStarActivitiesWithSessionFactory(sessionFactory SessionFactory) *CodeStarActivities {
	return &CodeStarActivities{sessionFactory: sessionFactory}
}

func (a *CodeStarActivities) getClient(ctx context.Context) (codestariface.CodeStarAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return codestar.New(sess), nil
}

func (a *CodeStarActivities) AssociateTeamMember(ctx context.Context, input *codestar.AssociateTeamMemberInput) (*codestar.AssociateTeamMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) CreateProject(ctx context.Context, input *codestar.CreateProjectInput) (*codestar.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) CreateUserProfile(ctx context.Context, input *codestar.CreateUserProfileInput) (*codestar.CreateUserProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DeleteProject(ctx context.Context, input *codestar.DeleteProjectInput) (*codestar.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) DeleteUserProfile(ctx context.Context, input *codestar.DeleteUserProfileInput) (*codestar.DeleteUserProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DescribeProject(ctx context.Context, input *codestar.DescribeProjectInput) (*codestar.DescribeProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) DescribeUserProfile(ctx context.Context, input *codestar.DescribeUserProfileInput) (*codestar.DescribeUserProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeUserProfileWithContext(ctx, input)
}

func (a *CodeStarActivities) DisassociateTeamMember(ctx context.Context, input *codestar.DisassociateTeamMemberInput) (*codestar.DisassociateTeamMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) ListProjects(ctx context.Context, input *codestar.ListProjectsInput) (*codestar.ListProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListProjectsWithContext(ctx, input)
}

func (a *CodeStarActivities) ListResources(ctx context.Context, input *codestar.ListResourcesInput) (*codestar.ListResourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResourcesWithContext(ctx, input)
}

func (a *CodeStarActivities) ListTagsForProject(ctx context.Context, input *codestar.ListTagsForProjectInput) (*codestar.ListTagsForProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) ListTeamMembers(ctx context.Context, input *codestar.ListTeamMembersInput) (*codestar.ListTeamMembersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTeamMembersWithContext(ctx, input)
}

func (a *CodeStarActivities) ListUserProfiles(ctx context.Context, input *codestar.ListUserProfilesInput) (*codestar.ListUserProfilesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListUserProfilesWithContext(ctx, input)
}

func (a *CodeStarActivities) TagProject(ctx context.Context, input *codestar.TagProjectInput) (*codestar.TagProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UntagProject(ctx context.Context, input *codestar.UntagProjectInput) (*codestar.UntagProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateProject(ctx context.Context, input *codestar.UpdateProjectInput) (*codestar.UpdateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateProjectWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateTeamMember(ctx context.Context, input *codestar.UpdateTeamMemberInput) (*codestar.UpdateTeamMemberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTeamMemberWithContext(ctx, input)
}

func (a *CodeStarActivities) UpdateUserProfile(ctx context.Context, input *codestar.UpdateUserProfileInput) (*codestar.UpdateUserProfileOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateUserProfileWithContext(ctx, input)
}
