// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/connectparticipant"
	"github.com/aws/aws-sdk-go/service/connectparticipant/connectparticipantiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ConnectParticipantActivities struct {
	client connectparticipantiface.ConnectParticipantAPI

	sessionFactory SessionFactory
}

func NewConnectParticipantActivities(sess *session.Session, config ...*aws.Config) *ConnectParticipantActivities {
	client := connectparticipant.New(sess, config...)
	return &ConnectParticipantActivities{client: client}
}

func NewConnectParticipantActivitiesWithSessionFactory(sessionFactory SessionFactory) *ConnectParticipantActivities {
	return &ConnectParticipantActivities{sessionFactory: sessionFactory}
}

func (a *ConnectParticipantActivities) getClient(ctx context.Context) (connectparticipantiface.ConnectParticipantAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return connectparticipant.New(sess), nil
}

func (a *ConnectParticipantActivities) CreateParticipantConnection(ctx context.Context, input *connectparticipant.CreateParticipantConnectionInput) (*connectparticipant.CreateParticipantConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateParticipantConnectionWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) DisconnectParticipant(ctx context.Context, input *connectparticipant.DisconnectParticipantInput) (*connectparticipant.DisconnectParticipantOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.DisconnectParticipantWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) GetTranscript(ctx context.Context, input *connectparticipant.GetTranscriptInput) (*connectparticipant.GetTranscriptOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTranscriptWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) SendEvent(ctx context.Context, input *connectparticipant.SendEventInput) (*connectparticipant.SendEventOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.SendEventWithContext(ctx, input)
}

func (a *ConnectParticipantActivities) SendMessage(ctx context.Context, input *connectparticipant.SendMessageInput) (*connectparticipant.SendMessageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	internal.SetClientToken(ctx, &input.ClientToken)
	return client.SendMessageWithContext(ctx, input)
}
