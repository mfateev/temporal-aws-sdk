// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53/route53iface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type Route53Activities struct {
	client route53iface.Route53API
}

func NewRoute53Activities(session *session.Session, config ...*aws.Config) *Route53Activities {
	client := route53.New(session, config...)
	return &Route53Activities{client: client}
}

func (a *Route53Activities) AssociateVPCWithHostedZone(ctx context.Context, input *route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error) {
	return a.client.AssociateVPCWithHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) ChangeResourceRecordSets(ctx context.Context, input *route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error) {
	return a.client.ChangeResourceRecordSetsWithContext(ctx, input)
}

func (a *Route53Activities) ChangeTagsForResource(ctx context.Context, input *route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error) {
	return a.client.ChangeTagsForResourceWithContext(ctx, input)
}

func (a *Route53Activities) CreateHealthCheck(ctx context.Context, input *route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error) {
	return a.client.CreateHealthCheckWithContext(ctx, input)
}

func (a *Route53Activities) CreateHostedZone(ctx context.Context, input *route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error) {
	return a.client.CreateHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) CreateQueryLoggingConfig(ctx context.Context, input *route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error) {
	return a.client.CreateQueryLoggingConfigWithContext(ctx, input)
}

func (a *Route53Activities) CreateReusableDelegationSet(ctx context.Context, input *route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error) {
	return a.client.CreateReusableDelegationSetWithContext(ctx, input)
}

func (a *Route53Activities) CreateTrafficPolicy(ctx context.Context, input *route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error) {
	return a.client.CreateTrafficPolicyWithContext(ctx, input)
}

func (a *Route53Activities) CreateTrafficPolicyInstance(ctx context.Context, input *route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error) {
	return a.client.CreateTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Route53Activities) CreateTrafficPolicyVersion(ctx context.Context, input *route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error) {
	return a.client.CreateTrafficPolicyVersionWithContext(ctx, input)
}

func (a *Route53Activities) CreateVPCAssociationAuthorization(ctx context.Context, input *route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error) {
	return a.client.CreateVPCAssociationAuthorizationWithContext(ctx, input)
}

func (a *Route53Activities) DeleteHealthCheck(ctx context.Context, input *route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error) {
	return a.client.DeleteHealthCheckWithContext(ctx, input)
}

func (a *Route53Activities) DeleteHostedZone(ctx context.Context, input *route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error) {
	return a.client.DeleteHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) DeleteQueryLoggingConfig(ctx context.Context, input *route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error) {
	return a.client.DeleteQueryLoggingConfigWithContext(ctx, input)
}

func (a *Route53Activities) DeleteReusableDelegationSet(ctx context.Context, input *route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error) {
	return a.client.DeleteReusableDelegationSetWithContext(ctx, input)
}

func (a *Route53Activities) DeleteTrafficPolicy(ctx context.Context, input *route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error) {
	return a.client.DeleteTrafficPolicyWithContext(ctx, input)
}

func (a *Route53Activities) DeleteTrafficPolicyInstance(ctx context.Context, input *route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error) {
	return a.client.DeleteTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Route53Activities) DeleteVPCAssociationAuthorization(ctx context.Context, input *route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error) {
	return a.client.DeleteVPCAssociationAuthorizationWithContext(ctx, input)
}

func (a *Route53Activities) DisassociateVPCFromHostedZone(ctx context.Context, input *route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error) {
	return a.client.DisassociateVPCFromHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) GetAccountLimit(ctx context.Context, input *route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error) {
	return a.client.GetAccountLimitWithContext(ctx, input)
}

func (a *Route53Activities) GetChange(ctx context.Context, input *route53.GetChangeInput) (*route53.GetChangeOutput, error) {
	return a.client.GetChangeWithContext(ctx, input)
}

func (a *Route53Activities) GetCheckerIpRanges(ctx context.Context, input *route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error) {
	return a.client.GetCheckerIpRangesWithContext(ctx, input)
}

func (a *Route53Activities) GetGeoLocation(ctx context.Context, input *route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error) {
	return a.client.GetGeoLocationWithContext(ctx, input)
}

func (a *Route53Activities) GetHealthCheck(ctx context.Context, input *route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error) {
	return a.client.GetHealthCheckWithContext(ctx, input)
}

func (a *Route53Activities) GetHealthCheckCount(ctx context.Context, input *route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error) {
	return a.client.GetHealthCheckCountWithContext(ctx, input)
}

func (a *Route53Activities) GetHealthCheckLastFailureReason(ctx context.Context, input *route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error) {
	return a.client.GetHealthCheckLastFailureReasonWithContext(ctx, input)
}

func (a *Route53Activities) GetHealthCheckStatus(ctx context.Context, input *route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error) {
	return a.client.GetHealthCheckStatusWithContext(ctx, input)
}

func (a *Route53Activities) GetHostedZone(ctx context.Context, input *route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error) {
	return a.client.GetHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) GetHostedZoneCount(ctx context.Context, input *route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error) {
	return a.client.GetHostedZoneCountWithContext(ctx, input)
}

func (a *Route53Activities) GetHostedZoneLimit(ctx context.Context, input *route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error) {
	return a.client.GetHostedZoneLimitWithContext(ctx, input)
}

func (a *Route53Activities) GetQueryLoggingConfig(ctx context.Context, input *route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error) {
	return a.client.GetQueryLoggingConfigWithContext(ctx, input)
}

func (a *Route53Activities) GetReusableDelegationSet(ctx context.Context, input *route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error) {
	return a.client.GetReusableDelegationSetWithContext(ctx, input)
}

func (a *Route53Activities) GetReusableDelegationSetLimit(ctx context.Context, input *route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error) {
	return a.client.GetReusableDelegationSetLimitWithContext(ctx, input)
}

func (a *Route53Activities) GetTrafficPolicy(ctx context.Context, input *route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error) {
	return a.client.GetTrafficPolicyWithContext(ctx, input)
}

func (a *Route53Activities) GetTrafficPolicyInstance(ctx context.Context, input *route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error) {
	return a.client.GetTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Route53Activities) GetTrafficPolicyInstanceCount(ctx context.Context, input *route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error) {
	return a.client.GetTrafficPolicyInstanceCountWithContext(ctx, input)
}

func (a *Route53Activities) ListGeoLocations(ctx context.Context, input *route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error) {
	return a.client.ListGeoLocationsWithContext(ctx, input)
}

func (a *Route53Activities) ListHealthChecks(ctx context.Context, input *route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error) {
	return a.client.ListHealthChecksWithContext(ctx, input)
}

func (a *Route53Activities) ListHostedZones(ctx context.Context, input *route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error) {
	return a.client.ListHostedZonesWithContext(ctx, input)
}

func (a *Route53Activities) ListHostedZonesByName(ctx context.Context, input *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	return a.client.ListHostedZonesByNameWithContext(ctx, input)
}

func (a *Route53Activities) ListHostedZonesByVPC(ctx context.Context, input *route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error) {
	return a.client.ListHostedZonesByVPCWithContext(ctx, input)
}

func (a *Route53Activities) ListQueryLoggingConfigs(ctx context.Context, input *route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error) {
	return a.client.ListQueryLoggingConfigsWithContext(ctx, input)
}

func (a *Route53Activities) ListResourceRecordSets(ctx context.Context, input *route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error) {
	return a.client.ListResourceRecordSetsWithContext(ctx, input)
}

func (a *Route53Activities) ListReusableDelegationSets(ctx context.Context, input *route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error) {
	return a.client.ListReusableDelegationSetsWithContext(ctx, input)
}

func (a *Route53Activities) ListTagsForResource(ctx context.Context, input *route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *Route53Activities) ListTagsForResources(ctx context.Context, input *route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error) {
	return a.client.ListTagsForResourcesWithContext(ctx, input)
}

func (a *Route53Activities) ListTrafficPolicies(ctx context.Context, input *route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error) {
	return a.client.ListTrafficPoliciesWithContext(ctx, input)
}

func (a *Route53Activities) ListTrafficPolicyInstances(ctx context.Context, input *route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error) {
	return a.client.ListTrafficPolicyInstancesWithContext(ctx, input)
}

func (a *Route53Activities) ListTrafficPolicyInstancesByHostedZone(ctx context.Context, input *route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error) {
	return a.client.ListTrafficPolicyInstancesByHostedZoneWithContext(ctx, input)
}

func (a *Route53Activities) ListTrafficPolicyInstancesByPolicy(ctx context.Context, input *route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error) {
	return a.client.ListTrafficPolicyInstancesByPolicyWithContext(ctx, input)
}

func (a *Route53Activities) ListTrafficPolicyVersions(ctx context.Context, input *route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error) {
	return a.client.ListTrafficPolicyVersionsWithContext(ctx, input)
}

func (a *Route53Activities) ListVPCAssociationAuthorizations(ctx context.Context, input *route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error) {
	return a.client.ListVPCAssociationAuthorizationsWithContext(ctx, input)
}

func (a *Route53Activities) TestDNSAnswer(ctx context.Context, input *route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error) {
	return a.client.TestDNSAnswerWithContext(ctx, input)
}

func (a *Route53Activities) UpdateHealthCheck(ctx context.Context, input *route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error) {
	return a.client.UpdateHealthCheckWithContext(ctx, input)
}

func (a *Route53Activities) UpdateHostedZoneComment(ctx context.Context, input *route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error) {
	return a.client.UpdateHostedZoneCommentWithContext(ctx, input)
}

func (a *Route53Activities) UpdateTrafficPolicyComment(ctx context.Context, input *route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error) {
	return a.client.UpdateTrafficPolicyCommentWithContext(ctx, input)
}

func (a *Route53Activities) UpdateTrafficPolicyInstance(ctx context.Context, input *route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error) {
	return a.client.UpdateTrafficPolicyInstanceWithContext(ctx, input)
}

func (a *Route53Activities) WaitUntilResourceRecordSetsChanged(ctx context.Context, input *route53.GetChangeInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilResourceRecordSetsChangedWithContext(ctx, input, options...)
	})
}
