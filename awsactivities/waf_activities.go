// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/waf/wafiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type WAFActivities struct {
	client wafiface.WAFAPI

	sessionFactory SessionFactory
}

func NewWAFActivities(sess *session.Session, config ...*aws.Config) *WAFActivities {
	client := waf.New(sess, config...)
	return &WAFActivities{client: client}
}

func NewWAFActivitiesWithSessionFactory(sessionFactory SessionFactory) *WAFActivities {
	return &WAFActivities{sessionFactory: sessionFactory}
}

func (a *WAFActivities) getClient(ctx context.Context) (wafiface.WAFAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return waf.New(sess), nil
}

func (a *WAFActivities) CreateByteMatchSet(ctx context.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateGeoMatchSet(ctx context.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateIPSet(ctx context.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateIPSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRateBasedRule(ctx context.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) CreateRegexMatchSet(ctx context.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRegexPatternSet(ctx context.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateRule(ctx context.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRuleWithContext(ctx, input)
}

func (a *WAFActivities) CreateRuleGroup(ctx context.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) CreateSizeConstraintSet(ctx context.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateSqlInjectionMatchSet(ctx context.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) CreateWebACL(ctx context.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateWebACLWithContext(ctx, input)
}

func (a *WAFActivities) CreateWebACLMigrationStack(ctx context.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateWebACLMigrationStackWithContext(ctx, input)
}

func (a *WAFActivities) CreateXssMatchSet(ctx context.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteByteMatchSet(ctx context.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteGeoMatchSet(ctx context.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteIPSet(ctx context.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteIPSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteLoggingConfiguration(ctx context.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) DeletePermissionPolicy(ctx context.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRateBasedRule(ctx context.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRegexMatchSet(ctx context.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRegexPatternSet(ctx context.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRule(ctx context.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRuleWithContext(ctx, input)
}

func (a *WAFActivities) DeleteRuleGroup(ctx context.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) DeleteSizeConstraintSet(ctx context.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteSqlInjectionMatchSet(ctx context.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) DeleteWebACL(ctx context.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteWebACLWithContext(ctx, input)
}

func (a *WAFActivities) DeleteXssMatchSet(ctx context.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetByteMatchSet(ctx context.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetChangeToken(ctx context.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetChangeTokenWithContext(ctx, input)
}

func (a *WAFActivities) GetChangeTokenStatus(ctx context.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetChangeTokenStatusWithContext(ctx, input)
}

func (a *WAFActivities) GetGeoMatchSet(ctx context.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetIPSet(ctx context.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetIPSetWithContext(ctx, input)
}

func (a *WAFActivities) GetLoggingConfiguration(ctx context.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) GetPermissionPolicy(ctx context.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetPermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) GetRateBasedRule(ctx context.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) GetRateBasedRuleManagedKeys(ctx context.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRateBasedRuleManagedKeysWithContext(ctx, input)
}

func (a *WAFActivities) GetRegexMatchSet(ctx context.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetRegexPatternSet(ctx context.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) GetRule(ctx context.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRuleWithContext(ctx, input)
}

func (a *WAFActivities) GetRuleGroup(ctx context.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) GetSampledRequests(ctx context.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSampledRequestsWithContext(ctx, input)
}

func (a *WAFActivities) GetSizeConstraintSet(ctx context.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) GetSqlInjectionMatchSet(ctx context.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) GetWebACL(ctx context.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetWebACLWithContext(ctx, input)
}

func (a *WAFActivities) GetXssMatchSet(ctx context.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.GetXssMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) ListActivatedRulesInRuleGroup(ctx context.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListActivatedRulesInRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) ListByteMatchSets(ctx context.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListByteMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListGeoMatchSets(ctx context.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListGeoMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListIPSets(ctx context.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListIPSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListLoggingConfigurations(ctx context.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListLoggingConfigurationsWithContext(ctx, input)
}

func (a *WAFActivities) ListRateBasedRules(ctx context.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRateBasedRulesWithContext(ctx, input)
}

func (a *WAFActivities) ListRegexMatchSets(ctx context.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRegexMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListRegexPatternSets(ctx context.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRegexPatternSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListRuleGroups(ctx context.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRuleGroupsWithContext(ctx, input)
}

func (a *WAFActivities) ListRules(ctx context.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListRulesWithContext(ctx, input)
}

func (a *WAFActivities) ListSizeConstraintSets(ctx context.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSizeConstraintSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListSqlInjectionMatchSets(ctx context.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSqlInjectionMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) ListSubscribedRuleGroups(ctx context.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSubscribedRuleGroupsWithContext(ctx, input)
}

func (a *WAFActivities) ListTagsForResource(ctx context.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WAFActivities) ListWebACLs(ctx context.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListWebACLsWithContext(ctx, input)
}

func (a *WAFActivities) ListXssMatchSets(ctx context.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListXssMatchSetsWithContext(ctx, input)
}

func (a *WAFActivities) PutLoggingConfiguration(ctx context.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFActivities) PutPermissionPolicy(ctx context.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.PutPermissionPolicyWithContext(ctx, input)
}

func (a *WAFActivities) TagResource(ctx context.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *WAFActivities) UntagResource(ctx context.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *WAFActivities) UpdateByteMatchSet(ctx context.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateByteMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateGeoMatchSet(ctx context.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateIPSet(ctx context.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateIPSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRateBasedRule(ctx context.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRegexMatchSet(ctx context.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRegexPatternSet(ctx context.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRule(ctx context.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRuleWithContext(ctx, input)
}

func (a *WAFActivities) UpdateRuleGroup(ctx context.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateRuleGroupWithContext(ctx, input)
}

func (a *WAFActivities) UpdateSizeConstraintSet(ctx context.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateSqlInjectionMatchSet(ctx context.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFActivities) UpdateWebACL(ctx context.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateWebACLWithContext(ctx, input)
}

func (a *WAFActivities) UpdateXssMatchSet(ctx context.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateXssMatchSetWithContext(ctx, input)
}
