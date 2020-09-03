package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/waf"
	"go.temporal.io/sdk/workflow"
)

type WAFRegionalClient interface {
    AssociateWebACL(ctx workflow.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error)
    AssociateWebACLAsync(ctx workflow.Context, input *wafregional.AssociateWebACLInput) *WafregionalAssociateWebACLResult

    CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error)
    CreateByteMatchSetAsync(ctx workflow.Context, input *waf.CreateByteMatchSetInput) *WafCreateByteMatchSetResult

    CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error)
    CreateGeoMatchSetAsync(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) *WafCreateGeoMatchSetResult

    CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error)
    CreateIPSetAsync(ctx workflow.Context, input *waf.CreateIPSetInput) *WafCreateIPSetResult

    CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error)
    CreateRateBasedRuleAsync(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) *WafCreateRateBasedRuleResult

    CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error)
    CreateRegexMatchSetAsync(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) *WafCreateRegexMatchSetResult

    CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error)
    CreateRegexPatternSetAsync(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) *WafCreateRegexPatternSetResult

    CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error)
    CreateRuleAsync(ctx workflow.Context, input *waf.CreateRuleInput) *WafCreateRuleResult

    CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error)
    CreateRuleGroupAsync(ctx workflow.Context, input *waf.CreateRuleGroupInput) *WafCreateRuleGroupResult

    CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error)
    CreateSizeConstraintSetAsync(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) *WafCreateSizeConstraintSetResult

    CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error)
    CreateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) *WafCreateSqlInjectionMatchSetResult

    CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error)
    CreateWebACLAsync(ctx workflow.Context, input *waf.CreateWebACLInput) *WafCreateWebACLResult

    CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error)
    CreateWebACLMigrationStackAsync(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) *WafCreateWebACLMigrationStackResult

    CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error)
    CreateXssMatchSetAsync(ctx workflow.Context, input *waf.CreateXssMatchSetInput) *WafCreateXssMatchSetResult

    DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error)
    DeleteByteMatchSetAsync(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) *WafDeleteByteMatchSetResult

    DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error)
    DeleteGeoMatchSetAsync(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) *WafDeleteGeoMatchSetResult

    DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error)
    DeleteIPSetAsync(ctx workflow.Context, input *waf.DeleteIPSetInput) *WafDeleteIPSetResult

    DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error)
    DeleteLoggingConfigurationAsync(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) *WafDeleteLoggingConfigurationResult

    DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error)
    DeletePermissionPolicyAsync(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) *WafDeletePermissionPolicyResult

    DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error)
    DeleteRateBasedRuleAsync(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) *WafDeleteRateBasedRuleResult

    DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error)
    DeleteRegexMatchSetAsync(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) *WafDeleteRegexMatchSetResult

    DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error)
    DeleteRegexPatternSetAsync(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) *WafDeleteRegexPatternSetResult

    DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error)
    DeleteRuleAsync(ctx workflow.Context, input *waf.DeleteRuleInput) *WafDeleteRuleResult

    DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error)
    DeleteRuleGroupAsync(ctx workflow.Context, input *waf.DeleteRuleGroupInput) *WafDeleteRuleGroupResult

    DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error)
    DeleteSizeConstraintSetAsync(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) *WafDeleteSizeConstraintSetResult

    DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error)
    DeleteSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) *WafDeleteSqlInjectionMatchSetResult

    DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error)
    DeleteWebACLAsync(ctx workflow.Context, input *waf.DeleteWebACLInput) *WafDeleteWebACLResult

    DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error)
    DeleteXssMatchSetAsync(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) *WafDeleteXssMatchSetResult

    DisassociateWebACL(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error)
    DisassociateWebACLAsync(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) *WafregionalDisassociateWebACLResult

    GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error)
    GetByteMatchSetAsync(ctx workflow.Context, input *waf.GetByteMatchSetInput) *WafGetByteMatchSetResult

    GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error)
    GetChangeTokenAsync(ctx workflow.Context, input *waf.GetChangeTokenInput) *WafGetChangeTokenResult

    GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error)
    GetChangeTokenStatusAsync(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) *WafGetChangeTokenStatusResult

    GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error)
    GetGeoMatchSetAsync(ctx workflow.Context, input *waf.GetGeoMatchSetInput) *WafGetGeoMatchSetResult

    GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error)
    GetIPSetAsync(ctx workflow.Context, input *waf.GetIPSetInput) *WafGetIPSetResult

    GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error)
    GetLoggingConfigurationAsync(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) *WafGetLoggingConfigurationResult

    GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error)
    GetPermissionPolicyAsync(ctx workflow.Context, input *waf.GetPermissionPolicyInput) *WafGetPermissionPolicyResult

    GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error)
    GetRateBasedRuleAsync(ctx workflow.Context, input *waf.GetRateBasedRuleInput) *WafGetRateBasedRuleResult

    GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error)
    GetRateBasedRuleManagedKeysAsync(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) *WafGetRateBasedRuleManagedKeysResult

    GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error)
    GetRegexMatchSetAsync(ctx workflow.Context, input *waf.GetRegexMatchSetInput) *WafGetRegexMatchSetResult

    GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error)
    GetRegexPatternSetAsync(ctx workflow.Context, input *waf.GetRegexPatternSetInput) *WafGetRegexPatternSetResult

    GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error)
    GetRuleAsync(ctx workflow.Context, input *waf.GetRuleInput) *WafGetRuleResult

    GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error)
    GetRuleGroupAsync(ctx workflow.Context, input *waf.GetRuleGroupInput) *WafGetRuleGroupResult

    GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error)
    GetSampledRequestsAsync(ctx workflow.Context, input *waf.GetSampledRequestsInput) *WafGetSampledRequestsResult

    GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error)
    GetSizeConstraintSetAsync(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) *WafGetSizeConstraintSetResult

    GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error)
    GetSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) *WafGetSqlInjectionMatchSetResult

    GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error)
    GetWebACLAsync(ctx workflow.Context, input *waf.GetWebACLInput) *WafGetWebACLResult

    GetWebACLForResource(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error)
    GetWebACLForResourceAsync(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) *WafregionalGetWebACLForResourceResult

    GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error)
    GetXssMatchSetAsync(ctx workflow.Context, input *waf.GetXssMatchSetInput) *WafGetXssMatchSetResult

    ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error)
    ListActivatedRulesInRuleGroupAsync(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) *WafListActivatedRulesInRuleGroupResult

    ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error)
    ListByteMatchSetsAsync(ctx workflow.Context, input *waf.ListByteMatchSetsInput) *WafListByteMatchSetsResult

    ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error)
    ListGeoMatchSetsAsync(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) *WafListGeoMatchSetsResult

    ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error)
    ListIPSetsAsync(ctx workflow.Context, input *waf.ListIPSetsInput) *WafListIPSetsResult

    ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error)
    ListLoggingConfigurationsAsync(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) *WafListLoggingConfigurationsResult

    ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error)
    ListRateBasedRulesAsync(ctx workflow.Context, input *waf.ListRateBasedRulesInput) *WafListRateBasedRulesResult

    ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error)
    ListRegexMatchSetsAsync(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) *WafListRegexMatchSetsResult

    ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error)
    ListRegexPatternSetsAsync(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) *WafListRegexPatternSetsResult

    ListResourcesForWebACL(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error)
    ListResourcesForWebACLAsync(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) *WafregionalListResourcesForWebACLResult

    ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error)
    ListRuleGroupsAsync(ctx workflow.Context, input *waf.ListRuleGroupsInput) *WafListRuleGroupsResult

    ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error)
    ListRulesAsync(ctx workflow.Context, input *waf.ListRulesInput) *WafListRulesResult

    ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error)
    ListSizeConstraintSetsAsync(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) *WafListSizeConstraintSetsResult

    ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error)
    ListSqlInjectionMatchSetsAsync(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) *WafListSqlInjectionMatchSetsResult

    ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error)
    ListSubscribedRuleGroupsAsync(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) *WafListSubscribedRuleGroupsResult

    ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *waf.ListTagsForResourceInput) *WafListTagsForResourceResult

    ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error)
    ListWebACLsAsync(ctx workflow.Context, input *waf.ListWebACLsInput) *WafListWebACLsResult

    ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error)
    ListXssMatchSetsAsync(ctx workflow.Context, input *waf.ListXssMatchSetsInput) *WafListXssMatchSetsResult

    PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error)
    PutLoggingConfigurationAsync(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) *WafPutLoggingConfigurationResult

    PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error)
    PutPermissionPolicyAsync(ctx workflow.Context, input *waf.PutPermissionPolicyInput) *WafPutPermissionPolicyResult

    TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *waf.TagResourceInput) *WafTagResourceResult

    UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *waf.UntagResourceInput) *WafUntagResourceResult

    UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error)
    UpdateByteMatchSetAsync(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) *WafUpdateByteMatchSetResult

    UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error)
    UpdateGeoMatchSetAsync(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) *WafUpdateGeoMatchSetResult

    UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error)
    UpdateIPSetAsync(ctx workflow.Context, input *waf.UpdateIPSetInput) *WafUpdateIPSetResult

    UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error)
    UpdateRateBasedRuleAsync(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) *WafUpdateRateBasedRuleResult

    UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error)
    UpdateRegexMatchSetAsync(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) *WafUpdateRegexMatchSetResult

    UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error)
    UpdateRegexPatternSetAsync(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) *WafUpdateRegexPatternSetResult

    UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error)
    UpdateRuleAsync(ctx workflow.Context, input *waf.UpdateRuleInput) *WafUpdateRuleResult

    UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error)
    UpdateRuleGroupAsync(ctx workflow.Context, input *waf.UpdateRuleGroupInput) *WafUpdateRuleGroupResult

    UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error)
    UpdateSizeConstraintSetAsync(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) *WafUpdateSizeConstraintSetResult

    UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error)
    UpdateSqlInjectionMatchSetAsync(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) *WafUpdateSqlInjectionMatchSetResult

    UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error)
    UpdateWebACLAsync(ctx workflow.Context, input *waf.UpdateWebACLInput) *WafUpdateWebACLResult

    UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error)
    UpdateXssMatchSetAsync(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) *WafUpdateXssMatchSetResult
}
type WafregionalAssociateWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalAssociateWebACLResult) Get(ctx workflow.Context) (*wafregional.AssociateWebACLOutput, error) {
    var output wafregional.AssociateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateByteMatchSetResult struct {
	Result workflow.Future
}

func (r *WafCreateByteMatchSetResult) Get(ctx workflow.Context) (*waf.CreateByteMatchSetOutput, error) {
    var output waf.CreateByteMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateGeoMatchSetResult struct {
	Result workflow.Future
}

func (r *WafCreateGeoMatchSetResult) Get(ctx workflow.Context) (*waf.CreateGeoMatchSetOutput, error) {
    var output waf.CreateGeoMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateIPSetResult struct {
	Result workflow.Future
}

func (r *WafCreateIPSetResult) Get(ctx workflow.Context) (*waf.CreateIPSetOutput, error) {
    var output waf.CreateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateRateBasedRuleResult struct {
	Result workflow.Future
}

func (r *WafCreateRateBasedRuleResult) Get(ctx workflow.Context) (*waf.CreateRateBasedRuleOutput, error) {
    var output waf.CreateRateBasedRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateRegexMatchSetResult struct {
	Result workflow.Future
}

func (r *WafCreateRegexMatchSetResult) Get(ctx workflow.Context) (*waf.CreateRegexMatchSetOutput, error) {
    var output waf.CreateRegexMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *WafCreateRegexPatternSetResult) Get(ctx workflow.Context) (*waf.CreateRegexPatternSetOutput, error) {
    var output waf.CreateRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateRuleResult struct {
	Result workflow.Future
}

func (r *WafCreateRuleResult) Get(ctx workflow.Context) (*waf.CreateRuleOutput, error) {
    var output waf.CreateRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateRuleGroupResult struct {
	Result workflow.Future
}

func (r *WafCreateRuleGroupResult) Get(ctx workflow.Context) (*waf.CreateRuleGroupOutput, error) {
    var output waf.CreateRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateSizeConstraintSetResult struct {
	Result workflow.Future
}

func (r *WafCreateSizeConstraintSetResult) Get(ctx workflow.Context) (*waf.CreateSizeConstraintSetOutput, error) {
    var output waf.CreateSizeConstraintSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateSqlInjectionMatchSetResult struct {
	Result workflow.Future
}

func (r *WafCreateSqlInjectionMatchSetResult) Get(ctx workflow.Context) (*waf.CreateSqlInjectionMatchSetOutput, error) {
    var output waf.CreateSqlInjectionMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateWebACLResult struct {
	Result workflow.Future
}

func (r *WafCreateWebACLResult) Get(ctx workflow.Context) (*waf.CreateWebACLOutput, error) {
    var output waf.CreateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateWebACLMigrationStackResult struct {
	Result workflow.Future
}

func (r *WafCreateWebACLMigrationStackResult) Get(ctx workflow.Context) (*waf.CreateWebACLMigrationStackOutput, error) {
    var output waf.CreateWebACLMigrationStackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafCreateXssMatchSetResult struct {
	Result workflow.Future
}

func (r *WafCreateXssMatchSetResult) Get(ctx workflow.Context) (*waf.CreateXssMatchSetOutput, error) {
    var output waf.CreateXssMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteByteMatchSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteByteMatchSetResult) Get(ctx workflow.Context) (*waf.DeleteByteMatchSetOutput, error) {
    var output waf.DeleteByteMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteGeoMatchSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteGeoMatchSetResult) Get(ctx workflow.Context) (*waf.DeleteGeoMatchSetOutput, error) {
    var output waf.DeleteGeoMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteIPSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteIPSetResult) Get(ctx workflow.Context) (*waf.DeleteIPSetOutput, error) {
    var output waf.DeleteIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *WafDeleteLoggingConfigurationResult) Get(ctx workflow.Context) (*waf.DeleteLoggingConfigurationOutput, error) {
    var output waf.DeleteLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeletePermissionPolicyResult struct {
	Result workflow.Future
}

func (r *WafDeletePermissionPolicyResult) Get(ctx workflow.Context) (*waf.DeletePermissionPolicyOutput, error) {
    var output waf.DeletePermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteRateBasedRuleResult struct {
	Result workflow.Future
}

func (r *WafDeleteRateBasedRuleResult) Get(ctx workflow.Context) (*waf.DeleteRateBasedRuleOutput, error) {
    var output waf.DeleteRateBasedRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteRegexMatchSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteRegexMatchSetResult) Get(ctx workflow.Context) (*waf.DeleteRegexMatchSetOutput, error) {
    var output waf.DeleteRegexMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteRegexPatternSetResult) Get(ctx workflow.Context) (*waf.DeleteRegexPatternSetOutput, error) {
    var output waf.DeleteRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteRuleResult struct {
	Result workflow.Future
}

func (r *WafDeleteRuleResult) Get(ctx workflow.Context) (*waf.DeleteRuleOutput, error) {
    var output waf.DeleteRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteRuleGroupResult struct {
	Result workflow.Future
}

func (r *WafDeleteRuleGroupResult) Get(ctx workflow.Context) (*waf.DeleteRuleGroupOutput, error) {
    var output waf.DeleteRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteSizeConstraintSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteSizeConstraintSetResult) Get(ctx workflow.Context) (*waf.DeleteSizeConstraintSetOutput, error) {
    var output waf.DeleteSizeConstraintSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteSqlInjectionMatchSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteSqlInjectionMatchSetResult) Get(ctx workflow.Context) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
    var output waf.DeleteSqlInjectionMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteWebACLResult struct {
	Result workflow.Future
}

func (r *WafDeleteWebACLResult) Get(ctx workflow.Context) (*waf.DeleteWebACLOutput, error) {
    var output waf.DeleteWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafDeleteXssMatchSetResult struct {
	Result workflow.Future
}

func (r *WafDeleteXssMatchSetResult) Get(ctx workflow.Context) (*waf.DeleteXssMatchSetOutput, error) {
    var output waf.DeleteXssMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafregionalDisassociateWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalDisassociateWebACLResult) Get(ctx workflow.Context) (*wafregional.DisassociateWebACLOutput, error) {
    var output wafregional.DisassociateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetByteMatchSetResult struct {
	Result workflow.Future
}

func (r *WafGetByteMatchSetResult) Get(ctx workflow.Context) (*waf.GetByteMatchSetOutput, error) {
    var output waf.GetByteMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetChangeTokenResult struct {
	Result workflow.Future
}

func (r *WafGetChangeTokenResult) Get(ctx workflow.Context) (*waf.GetChangeTokenOutput, error) {
    var output waf.GetChangeTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetChangeTokenStatusResult struct {
	Result workflow.Future
}

func (r *WafGetChangeTokenStatusResult) Get(ctx workflow.Context) (*waf.GetChangeTokenStatusOutput, error) {
    var output waf.GetChangeTokenStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetGeoMatchSetResult struct {
	Result workflow.Future
}

func (r *WafGetGeoMatchSetResult) Get(ctx workflow.Context) (*waf.GetGeoMatchSetOutput, error) {
    var output waf.GetGeoMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetIPSetResult struct {
	Result workflow.Future
}

func (r *WafGetIPSetResult) Get(ctx workflow.Context) (*waf.GetIPSetOutput, error) {
    var output waf.GetIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *WafGetLoggingConfigurationResult) Get(ctx workflow.Context) (*waf.GetLoggingConfigurationOutput, error) {
    var output waf.GetLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetPermissionPolicyResult struct {
	Result workflow.Future
}

func (r *WafGetPermissionPolicyResult) Get(ctx workflow.Context) (*waf.GetPermissionPolicyOutput, error) {
    var output waf.GetPermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRateBasedRuleResult struct {
	Result workflow.Future
}

func (r *WafGetRateBasedRuleResult) Get(ctx workflow.Context) (*waf.GetRateBasedRuleOutput, error) {
    var output waf.GetRateBasedRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRateBasedRuleManagedKeysResult struct {
	Result workflow.Future
}

func (r *WafGetRateBasedRuleManagedKeysResult) Get(ctx workflow.Context) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
    var output waf.GetRateBasedRuleManagedKeysOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRegexMatchSetResult struct {
	Result workflow.Future
}

func (r *WafGetRegexMatchSetResult) Get(ctx workflow.Context) (*waf.GetRegexMatchSetOutput, error) {
    var output waf.GetRegexMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *WafGetRegexPatternSetResult) Get(ctx workflow.Context) (*waf.GetRegexPatternSetOutput, error) {
    var output waf.GetRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRuleResult struct {
	Result workflow.Future
}

func (r *WafGetRuleResult) Get(ctx workflow.Context) (*waf.GetRuleOutput, error) {
    var output waf.GetRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetRuleGroupResult struct {
	Result workflow.Future
}

func (r *WafGetRuleGroupResult) Get(ctx workflow.Context) (*waf.GetRuleGroupOutput, error) {
    var output waf.GetRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetSampledRequestsResult struct {
	Result workflow.Future
}

func (r *WafGetSampledRequestsResult) Get(ctx workflow.Context) (*waf.GetSampledRequestsOutput, error) {
    var output waf.GetSampledRequestsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetSizeConstraintSetResult struct {
	Result workflow.Future
}

func (r *WafGetSizeConstraintSetResult) Get(ctx workflow.Context) (*waf.GetSizeConstraintSetOutput, error) {
    var output waf.GetSizeConstraintSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetSqlInjectionMatchSetResult struct {
	Result workflow.Future
}

func (r *WafGetSqlInjectionMatchSetResult) Get(ctx workflow.Context) (*waf.GetSqlInjectionMatchSetOutput, error) {
    var output waf.GetSqlInjectionMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetWebACLResult struct {
	Result workflow.Future
}

func (r *WafGetWebACLResult) Get(ctx workflow.Context) (*waf.GetWebACLOutput, error) {
    var output waf.GetWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafregionalGetWebACLForResourceResult struct {
	Result workflow.Future
}

func (r *WafregionalGetWebACLForResourceResult) Get(ctx workflow.Context) (*wafregional.GetWebACLForResourceOutput, error) {
    var output wafregional.GetWebACLForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafGetXssMatchSetResult struct {
	Result workflow.Future
}

func (r *WafGetXssMatchSetResult) Get(ctx workflow.Context) (*waf.GetXssMatchSetOutput, error) {
    var output waf.GetXssMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListActivatedRulesInRuleGroupResult struct {
	Result workflow.Future
}

func (r *WafListActivatedRulesInRuleGroupResult) Get(ctx workflow.Context) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
    var output waf.ListActivatedRulesInRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListByteMatchSetsResult struct {
	Result workflow.Future
}

func (r *WafListByteMatchSetsResult) Get(ctx workflow.Context) (*waf.ListByteMatchSetsOutput, error) {
    var output waf.ListByteMatchSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListGeoMatchSetsResult struct {
	Result workflow.Future
}

func (r *WafListGeoMatchSetsResult) Get(ctx workflow.Context) (*waf.ListGeoMatchSetsOutput, error) {
    var output waf.ListGeoMatchSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListIPSetsResult struct {
	Result workflow.Future
}

func (r *WafListIPSetsResult) Get(ctx workflow.Context) (*waf.ListIPSetsOutput, error) {
    var output waf.ListIPSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListLoggingConfigurationsResult struct {
	Result workflow.Future
}

func (r *WafListLoggingConfigurationsResult) Get(ctx workflow.Context) (*waf.ListLoggingConfigurationsOutput, error) {
    var output waf.ListLoggingConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListRateBasedRulesResult struct {
	Result workflow.Future
}

func (r *WafListRateBasedRulesResult) Get(ctx workflow.Context) (*waf.ListRateBasedRulesOutput, error) {
    var output waf.ListRateBasedRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListRegexMatchSetsResult struct {
	Result workflow.Future
}

func (r *WafListRegexMatchSetsResult) Get(ctx workflow.Context) (*waf.ListRegexMatchSetsOutput, error) {
    var output waf.ListRegexMatchSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListRegexPatternSetsResult struct {
	Result workflow.Future
}

func (r *WafListRegexPatternSetsResult) Get(ctx workflow.Context) (*waf.ListRegexPatternSetsOutput, error) {
    var output waf.ListRegexPatternSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafregionalListResourcesForWebACLResult struct {
	Result workflow.Future
}

func (r *WafregionalListResourcesForWebACLResult) Get(ctx workflow.Context) (*wafregional.ListResourcesForWebACLOutput, error) {
    var output wafregional.ListResourcesForWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListRuleGroupsResult struct {
	Result workflow.Future
}

func (r *WafListRuleGroupsResult) Get(ctx workflow.Context) (*waf.ListRuleGroupsOutput, error) {
    var output waf.ListRuleGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListRulesResult struct {
	Result workflow.Future
}

func (r *WafListRulesResult) Get(ctx workflow.Context) (*waf.ListRulesOutput, error) {
    var output waf.ListRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListSizeConstraintSetsResult struct {
	Result workflow.Future
}

func (r *WafListSizeConstraintSetsResult) Get(ctx workflow.Context) (*waf.ListSizeConstraintSetsOutput, error) {
    var output waf.ListSizeConstraintSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListSqlInjectionMatchSetsResult struct {
	Result workflow.Future
}

func (r *WafListSqlInjectionMatchSetsResult) Get(ctx workflow.Context) (*waf.ListSqlInjectionMatchSetsOutput, error) {
    var output waf.ListSqlInjectionMatchSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListSubscribedRuleGroupsResult struct {
	Result workflow.Future
}

func (r *WafListSubscribedRuleGroupsResult) Get(ctx workflow.Context) (*waf.ListSubscribedRuleGroupsOutput, error) {
    var output waf.ListSubscribedRuleGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *WafListTagsForResourceResult) Get(ctx workflow.Context) (*waf.ListTagsForResourceOutput, error) {
    var output waf.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListWebACLsResult struct {
	Result workflow.Future
}

func (r *WafListWebACLsResult) Get(ctx workflow.Context) (*waf.ListWebACLsOutput, error) {
    var output waf.ListWebACLsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafListXssMatchSetsResult struct {
	Result workflow.Future
}

func (r *WafListXssMatchSetsResult) Get(ctx workflow.Context) (*waf.ListXssMatchSetsOutput, error) {
    var output waf.ListXssMatchSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafPutLoggingConfigurationResult struct {
	Result workflow.Future
}

func (r *WafPutLoggingConfigurationResult) Get(ctx workflow.Context) (*waf.PutLoggingConfigurationOutput, error) {
    var output waf.PutLoggingConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafPutPermissionPolicyResult struct {
	Result workflow.Future
}

func (r *WafPutPermissionPolicyResult) Get(ctx workflow.Context) (*waf.PutPermissionPolicyOutput, error) {
    var output waf.PutPermissionPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafTagResourceResult struct {
	Result workflow.Future
}

func (r *WafTagResourceResult) Get(ctx workflow.Context) (*waf.TagResourceOutput, error) {
    var output waf.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUntagResourceResult struct {
	Result workflow.Future
}

func (r *WafUntagResourceResult) Get(ctx workflow.Context) (*waf.UntagResourceOutput, error) {
    var output waf.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateByteMatchSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateByteMatchSetResult) Get(ctx workflow.Context) (*waf.UpdateByteMatchSetOutput, error) {
    var output waf.UpdateByteMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateGeoMatchSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateGeoMatchSetResult) Get(ctx workflow.Context) (*waf.UpdateGeoMatchSetOutput, error) {
    var output waf.UpdateGeoMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateIPSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateIPSetResult) Get(ctx workflow.Context) (*waf.UpdateIPSetOutput, error) {
    var output waf.UpdateIPSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateRateBasedRuleResult struct {
	Result workflow.Future
}

func (r *WafUpdateRateBasedRuleResult) Get(ctx workflow.Context) (*waf.UpdateRateBasedRuleOutput, error) {
    var output waf.UpdateRateBasedRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateRegexMatchSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateRegexMatchSetResult) Get(ctx workflow.Context) (*waf.UpdateRegexMatchSetOutput, error) {
    var output waf.UpdateRegexMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateRegexPatternSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateRegexPatternSetResult) Get(ctx workflow.Context) (*waf.UpdateRegexPatternSetOutput, error) {
    var output waf.UpdateRegexPatternSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateRuleResult struct {
	Result workflow.Future
}

func (r *WafUpdateRuleResult) Get(ctx workflow.Context) (*waf.UpdateRuleOutput, error) {
    var output waf.UpdateRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateRuleGroupResult struct {
	Result workflow.Future
}

func (r *WafUpdateRuleGroupResult) Get(ctx workflow.Context) (*waf.UpdateRuleGroupOutput, error) {
    var output waf.UpdateRuleGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateSizeConstraintSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateSizeConstraintSetResult) Get(ctx workflow.Context) (*waf.UpdateSizeConstraintSetOutput, error) {
    var output waf.UpdateSizeConstraintSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateSqlInjectionMatchSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateSqlInjectionMatchSetResult) Get(ctx workflow.Context) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
    var output waf.UpdateSqlInjectionMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateWebACLResult struct {
	Result workflow.Future
}

func (r *WafUpdateWebACLResult) Get(ctx workflow.Context) (*waf.UpdateWebACLOutput, error) {
    var output waf.UpdateWebACLOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type WafUpdateXssMatchSetResult struct {
	Result workflow.Future
}

func (r *WafUpdateXssMatchSetResult) Get(ctx workflow.Context) (*waf.UpdateXssMatchSetOutput, error) {
    var output waf.UpdateXssMatchSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type WAFRegionalStub struct {
    activities WAFRegionalClient
}

func NewWAFRegionalStub() WAFRegionalClient {
    return &WAFRegionalStub{}
}

func (a *WAFRegionalStub) AssociateWebACL(ctx workflow.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
    var output wafregional.AssociateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateByteMatchSet(ctx workflow.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
    var output waf.CreateByteMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateByteMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateGeoMatchSet(ctx workflow.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
    var output waf.CreateGeoMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateGeoMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateIPSet(ctx workflow.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
    var output waf.CreateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateRateBasedRule(ctx workflow.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
    var output waf.CreateRateBasedRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRateBasedRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateRegexMatchSet(ctx workflow.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
    var output waf.CreateRegexMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRegexMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateRegexPatternSet(ctx workflow.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
    var output waf.CreateRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateRule(ctx workflow.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
    var output waf.CreateRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateRuleGroup(ctx workflow.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
    var output waf.CreateRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateSizeConstraintSet(ctx workflow.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
    var output waf.CreateSizeConstraintSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSizeConstraintSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateSqlInjectionMatchSet(ctx workflow.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
    var output waf.CreateSqlInjectionMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSqlInjectionMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateWebACL(ctx workflow.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
    var output waf.CreateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateWebACLMigrationStack(ctx workflow.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
    var output waf.CreateWebACLMigrationStackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWebACLMigrationStack, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) CreateXssMatchSet(ctx workflow.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
    var output waf.CreateXssMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateXssMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteByteMatchSet(ctx workflow.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
    var output waf.DeleteByteMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteByteMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteGeoMatchSet(ctx workflow.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
    var output waf.DeleteGeoMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteGeoMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteIPSet(ctx workflow.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
    var output waf.DeleteIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteLoggingConfiguration(ctx workflow.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
    var output waf.DeleteLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeletePermissionPolicy(ctx workflow.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
    var output waf.DeletePermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteRateBasedRule(ctx workflow.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
    var output waf.DeleteRateBasedRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRateBasedRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteRegexMatchSet(ctx workflow.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
    var output waf.DeleteRegexMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteRegexPatternSet(ctx workflow.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
    var output waf.DeleteRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteRule(ctx workflow.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
    var output waf.DeleteRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteRuleGroup(ctx workflow.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
    var output waf.DeleteRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteSizeConstraintSet(ctx workflow.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
    var output waf.DeleteSizeConstraintSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSizeConstraintSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteSqlInjectionMatchSet(ctx workflow.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
    var output waf.DeleteSqlInjectionMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteSqlInjectionMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteWebACL(ctx workflow.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
    var output waf.DeleteWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DeleteXssMatchSet(ctx workflow.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
    var output waf.DeleteXssMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteXssMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) DisassociateWebACL(ctx workflow.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error) {
    var output wafregional.DisassociateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetByteMatchSet(ctx workflow.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
    var output waf.GetByteMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetByteMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetChangeToken(ctx workflow.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
    var output waf.GetChangeTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetChangeToken, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetChangeTokenStatus(ctx workflow.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
    var output waf.GetChangeTokenStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetChangeTokenStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetGeoMatchSet(ctx workflow.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
    var output waf.GetGeoMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetGeoMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetIPSet(ctx workflow.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
    var output waf.GetIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetLoggingConfiguration(ctx workflow.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
    var output waf.GetLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetPermissionPolicy(ctx workflow.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
    var output waf.GetPermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRateBasedRule(ctx workflow.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
    var output waf.GetRateBasedRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRateBasedRuleManagedKeys(ctx workflow.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
    var output waf.GetRateBasedRuleManagedKeysOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRateBasedRuleManagedKeys, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRegexMatchSet(ctx workflow.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
    var output waf.GetRegexMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRegexMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRegexPatternSet(ctx workflow.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
    var output waf.GetRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRule(ctx workflow.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
    var output waf.GetRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetRuleGroup(ctx workflow.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
    var output waf.GetRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetSampledRequests(ctx workflow.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
    var output waf.GetSampledRequestsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSampledRequests, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetSizeConstraintSet(ctx workflow.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
    var output waf.GetSizeConstraintSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSizeConstraintSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetSqlInjectionMatchSet(ctx workflow.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
    var output waf.GetSqlInjectionMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSqlInjectionMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetWebACL(ctx workflow.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
    var output waf.GetWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetWebACLForResource(ctx workflow.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
    var output wafregional.GetWebACLForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWebACLForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) GetXssMatchSet(ctx workflow.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
    var output waf.GetXssMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetXssMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListActivatedRulesInRuleGroup(ctx workflow.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
    var output waf.ListActivatedRulesInRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActivatedRulesInRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListByteMatchSets(ctx workflow.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
    var output waf.ListByteMatchSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListByteMatchSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListGeoMatchSets(ctx workflow.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
    var output waf.ListGeoMatchSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGeoMatchSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListIPSets(ctx workflow.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
    var output waf.ListIPSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIPSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListLoggingConfigurations(ctx workflow.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
    var output waf.ListLoggingConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListLoggingConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListRateBasedRules(ctx workflow.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
    var output waf.ListRateBasedRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRateBasedRules, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListRegexMatchSets(ctx workflow.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
    var output waf.ListRegexMatchSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRegexMatchSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListRegexPatternSets(ctx workflow.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
    var output waf.ListRegexPatternSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRegexPatternSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListResourcesForWebACL(ctx workflow.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
    var output wafregional.ListResourcesForWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResourcesForWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListRuleGroups(ctx workflow.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
    var output waf.ListRuleGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRuleGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListRules(ctx workflow.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
    var output waf.ListRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRules, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListSizeConstraintSets(ctx workflow.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
    var output waf.ListSizeConstraintSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSizeConstraintSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListSqlInjectionMatchSets(ctx workflow.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
    var output waf.ListSqlInjectionMatchSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSqlInjectionMatchSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListSubscribedRuleGroups(ctx workflow.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
    var output waf.ListSubscribedRuleGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSubscribedRuleGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListTagsForResource(ctx workflow.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
    var output waf.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListWebACLs(ctx workflow.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
    var output waf.ListWebACLsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebACLs, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) ListXssMatchSets(ctx workflow.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
    var output waf.ListXssMatchSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListXssMatchSets, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) PutLoggingConfiguration(ctx workflow.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
    var output waf.PutLoggingConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutLoggingConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) PutPermissionPolicy(ctx workflow.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
    var output waf.PutPermissionPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPermissionPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) TagResource(ctx workflow.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
    var output waf.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UntagResource(ctx workflow.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
    var output waf.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateByteMatchSet(ctx workflow.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
    var output waf.UpdateByteMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateByteMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateGeoMatchSet(ctx workflow.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
    var output waf.UpdateGeoMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateGeoMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateIPSet(ctx workflow.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
    var output waf.UpdateIPSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateIPSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateRateBasedRule(ctx workflow.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
    var output waf.UpdateRateBasedRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRateBasedRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateRegexMatchSet(ctx workflow.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
    var output waf.UpdateRegexMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateRegexPatternSet(ctx workflow.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
    var output waf.UpdateRegexPatternSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegexPatternSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateRule(ctx workflow.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
    var output waf.UpdateRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRule, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateRuleGroup(ctx workflow.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
    var output waf.UpdateRuleGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRuleGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateSizeConstraintSet(ctx workflow.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
    var output waf.UpdateSizeConstraintSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSizeConstraintSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateSqlInjectionMatchSet(ctx workflow.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
    var output waf.UpdateSqlInjectionMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateSqlInjectionMatchSet, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateWebACL(ctx workflow.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
    var output waf.UpdateWebACLOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWebACL, input).Get(ctx, &output)
    return &output, err
}

func (a *WAFRegionalStub) UpdateXssMatchSet(ctx workflow.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
    var output waf.UpdateXssMatchSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateXssMatchSet, input).Get(ctx, &output)
    return &output, err
}
