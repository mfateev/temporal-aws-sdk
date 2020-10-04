// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/budgets"
	"github.com/aws/aws-sdk-go/service/budgets/budgetsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type BudgetsActivities struct {
	client budgetsiface.BudgetsAPI

	sessionFactory SessionFactory
}

func NewBudgetsActivities(sess *session.Session, config ...*aws.Config) *BudgetsActivities {
	client := budgets.New(sess, config...)
	return &BudgetsActivities{client: client}
}

func NewBudgetsActivitiesWithSessionFactory(sessionFactory SessionFactory) *BudgetsActivities {
	return &BudgetsActivities{sessionFactory: sessionFactory}
}

func (a *BudgetsActivities) getClient(ctx context.Context) (budgetsiface.BudgetsAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return budgets.New(sess), nil
}

func (a *BudgetsActivities) CreateBudget(ctx context.Context, input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) CreateNotification(ctx context.Context, input *budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) CreateSubscriber(ctx context.Context, input *budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSubscriberWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteBudget(ctx context.Context, input *budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteNotification(ctx context.Context, input *budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) DeleteSubscriber(ctx context.Context, input *budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSubscriberWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudget(ctx context.Context, input *budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudgetPerformanceHistory(ctx context.Context, input *budgets.DescribeBudgetPerformanceHistoryInput) (*budgets.DescribeBudgetPerformanceHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBudgetPerformanceHistoryWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeBudgets(ctx context.Context, input *budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeBudgetsWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeNotificationsForBudget(ctx context.Context, input *budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeNotificationsForBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) DescribeSubscribersForNotification(ctx context.Context, input *budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSubscribersForNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateBudget(ctx context.Context, input *budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateBudgetWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateNotification(ctx context.Context, input *budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateNotificationWithContext(ctx, input)
}

func (a *BudgetsActivities) UpdateSubscriber(ctx context.Context, input *budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSubscriberWithContext(ctx, input)
}
