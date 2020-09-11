// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appmesh"
	"github.com/aws/aws-sdk-go/service/appmesh/appmeshiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AppMeshActivities struct {
	client appmeshiface.AppMeshAPI
}

func NewAppMeshActivities(session *session.Session, config ...*aws.Config) *AppMeshActivities {
	client := appmesh.New(session, config...)
	return &AppMeshActivities{client: client}
}

func (a *AppMeshActivities) CreateGatewayRoute(ctx context.Context, input *appmesh.CreateGatewayRouteInput) (*appmesh.CreateGatewayRouteOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateMesh(ctx context.Context, input *appmesh.CreateMeshInput) (*appmesh.CreateMeshOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateRoute(ctx context.Context, input *appmesh.CreateRouteInput) (*appmesh.CreateRouteOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualGateway(ctx context.Context, input *appmesh.CreateVirtualGatewayInput) (*appmesh.CreateVirtualGatewayOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualNode(ctx context.Context, input *appmesh.CreateVirtualNodeInput) (*appmesh.CreateVirtualNodeOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualRouter(ctx context.Context, input *appmesh.CreateVirtualRouterInput) (*appmesh.CreateVirtualRouterOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) CreateVirtualService(ctx context.Context, input *appmesh.CreateVirtualServiceInput) (*appmesh.CreateVirtualServiceOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.CreateVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteGatewayRoute(ctx context.Context, input *appmesh.DeleteGatewayRouteInput) (*appmesh.DeleteGatewayRouteOutput, error) {
	return a.client.DeleteGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteMesh(ctx context.Context, input *appmesh.DeleteMeshInput) (*appmesh.DeleteMeshOutput, error) {
	return a.client.DeleteMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteRoute(ctx context.Context, input *appmesh.DeleteRouteInput) (*appmesh.DeleteRouteOutput, error) {
	return a.client.DeleteRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualGateway(ctx context.Context, input *appmesh.DeleteVirtualGatewayInput) (*appmesh.DeleteVirtualGatewayOutput, error) {
	return a.client.DeleteVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualNode(ctx context.Context, input *appmesh.DeleteVirtualNodeInput) (*appmesh.DeleteVirtualNodeOutput, error) {
	return a.client.DeleteVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualRouter(ctx context.Context, input *appmesh.DeleteVirtualRouterInput) (*appmesh.DeleteVirtualRouterOutput, error) {
	return a.client.DeleteVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) DeleteVirtualService(ctx context.Context, input *appmesh.DeleteVirtualServiceInput) (*appmesh.DeleteVirtualServiceOutput, error) {
	return a.client.DeleteVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeGatewayRoute(ctx context.Context, input *appmesh.DescribeGatewayRouteInput) (*appmesh.DescribeGatewayRouteOutput, error) {
	return a.client.DescribeGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeMesh(ctx context.Context, input *appmesh.DescribeMeshInput) (*appmesh.DescribeMeshOutput, error) {
	return a.client.DescribeMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeRoute(ctx context.Context, input *appmesh.DescribeRouteInput) (*appmesh.DescribeRouteOutput, error) {
	return a.client.DescribeRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualGateway(ctx context.Context, input *appmesh.DescribeVirtualGatewayInput) (*appmesh.DescribeVirtualGatewayOutput, error) {
	return a.client.DescribeVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualNode(ctx context.Context, input *appmesh.DescribeVirtualNodeInput) (*appmesh.DescribeVirtualNodeOutput, error) {
	return a.client.DescribeVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualRouter(ctx context.Context, input *appmesh.DescribeVirtualRouterInput) (*appmesh.DescribeVirtualRouterOutput, error) {
	return a.client.DescribeVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) DescribeVirtualService(ctx context.Context, input *appmesh.DescribeVirtualServiceInput) (*appmesh.DescribeVirtualServiceOutput, error) {
	return a.client.DescribeVirtualServiceWithContext(ctx, input)
}

func (a *AppMeshActivities) ListGatewayRoutes(ctx context.Context, input *appmesh.ListGatewayRoutesInput) (*appmesh.ListGatewayRoutesOutput, error) {
	return a.client.ListGatewayRoutesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListMeshes(ctx context.Context, input *appmesh.ListMeshesInput) (*appmesh.ListMeshesOutput, error) {
	return a.client.ListMeshesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListRoutes(ctx context.Context, input *appmesh.ListRoutesInput) (*appmesh.ListRoutesOutput, error) {
	return a.client.ListRoutesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListTagsForResource(ctx context.Context, input *appmesh.ListTagsForResourceInput) (*appmesh.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualGateways(ctx context.Context, input *appmesh.ListVirtualGatewaysInput) (*appmesh.ListVirtualGatewaysOutput, error) {
	return a.client.ListVirtualGatewaysWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualNodes(ctx context.Context, input *appmesh.ListVirtualNodesInput) (*appmesh.ListVirtualNodesOutput, error) {
	return a.client.ListVirtualNodesWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualRouters(ctx context.Context, input *appmesh.ListVirtualRoutersInput) (*appmesh.ListVirtualRoutersOutput, error) {
	return a.client.ListVirtualRoutersWithContext(ctx, input)
}

func (a *AppMeshActivities) ListVirtualServices(ctx context.Context, input *appmesh.ListVirtualServicesInput) (*appmesh.ListVirtualServicesOutput, error) {
	return a.client.ListVirtualServicesWithContext(ctx, input)
}

func (a *AppMeshActivities) TagResource(ctx context.Context, input *appmesh.TagResourceInput) (*appmesh.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) UntagResource(ctx context.Context, input *appmesh.UntagResourceInput) (*appmesh.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateGatewayRoute(ctx context.Context, input *appmesh.UpdateGatewayRouteInput) (*appmesh.UpdateGatewayRouteOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateGatewayRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateMesh(ctx context.Context, input *appmesh.UpdateMeshInput) (*appmesh.UpdateMeshOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateMeshWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateRoute(ctx context.Context, input *appmesh.UpdateRouteInput) (*appmesh.UpdateRouteOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateRouteWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualGateway(ctx context.Context, input *appmesh.UpdateVirtualGatewayInput) (*appmesh.UpdateVirtualGatewayOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateVirtualGatewayWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualNode(ctx context.Context, input *appmesh.UpdateVirtualNodeInput) (*appmesh.UpdateVirtualNodeOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateVirtualNodeWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualRouter(ctx context.Context, input *appmesh.UpdateVirtualRouterInput) (*appmesh.UpdateVirtualRouterOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateVirtualRouterWithContext(ctx, input)
}

func (a *AppMeshActivities) UpdateVirtualService(ctx context.Context, input *appmesh.UpdateVirtualServiceInput) (*appmesh.UpdateVirtualServiceOutput, error) {
	internal.SetClientToken(ctx, &input.ClientToken)
	return a.client.UpdateVirtualServiceWithContext(ctx, input)
}
