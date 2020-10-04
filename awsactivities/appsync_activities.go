// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/appsync/appsynciface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type AppSyncActivities struct {
	client appsynciface.AppSyncAPI

	sessionFactory SessionFactory
}

func NewAppSyncActivities(sess *session.Session, config ...*aws.Config) *AppSyncActivities {
	client := appsync.New(sess, config...)
	return &AppSyncActivities{client: client}
}

func NewAppSyncActivitiesWithSessionFactory(sessionFactory SessionFactory) *AppSyncActivities {
	return &AppSyncActivities{sessionFactory: sessionFactory}
}

func (a *AppSyncActivities) getClient(ctx context.Context) (appsynciface.AppSyncAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return appsync.New(sess), nil
}

func (a *AppSyncActivities) CreateApiCache(ctx context.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateApiKey(ctx context.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateDataSource(ctx context.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateFunction(ctx context.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateGraphqlApi(ctx context.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateResolver(ctx context.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateType(ctx context.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteApiCache(ctx context.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteApiKey(ctx context.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteDataSource(ctx context.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteFunction(ctx context.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteGraphqlApi(ctx context.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteResolver(ctx context.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteType(ctx context.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) FlushApiCache(ctx context.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.FlushApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) GetApiCache(ctx context.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) GetDataSource(ctx context.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) GetFunction(ctx context.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) GetGraphqlApi(ctx context.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) GetIntrospectionSchema(ctx context.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIntrospectionSchemaWithContext(ctx, input)
}

func (a *AppSyncActivities) GetResolver(ctx context.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) GetSchemaCreationStatus(ctx context.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSchemaCreationStatusWithContext(ctx, input)
}

func (a *AppSyncActivities) GetType(ctx context.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) ListApiKeys(ctx context.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListApiKeysWithContext(ctx, input)
}

func (a *AppSyncActivities) ListDataSources(ctx context.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDataSourcesWithContext(ctx, input)
}

func (a *AppSyncActivities) ListFunctions(ctx context.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListFunctionsWithContext(ctx, input)
}

func (a *AppSyncActivities) ListGraphqlApis(ctx context.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGraphqlApisWithContext(ctx, input)
}

func (a *AppSyncActivities) ListResolvers(ctx context.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolversWithContext(ctx, input)
}

func (a *AppSyncActivities) ListResolversByFunction(ctx context.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListResolversByFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) ListTagsForResource(ctx context.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) ListTypes(ctx context.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTypesWithContext(ctx, input)
}

func (a *AppSyncActivities) StartSchemaCreation(ctx context.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartSchemaCreationWithContext(ctx, input)
}

func (a *AppSyncActivities) TagResource(ctx context.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UntagResource(ctx context.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateApiCache(ctx context.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateApiKey(ctx context.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateDataSource(ctx context.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateFunction(ctx context.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateGraphqlApi(ctx context.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateResolver(ctx context.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateType(ctx context.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateTypeWithContext(ctx, input)
}
