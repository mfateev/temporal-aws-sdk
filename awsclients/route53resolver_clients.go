package awsclients

import (
	"github.com/aws/aws-sdk-go/service/route53resolver"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type Route53ResolverClient interface {
       AssociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error)
       AssociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) *Route53resolverAssociateResolverEndpointIpAddressResult

       AssociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error)
       AssociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) *Route53resolverAssociateResolverQueryLogConfigResult

       AssociateResolverRule(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error)
       AssociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) *Route53resolverAssociateResolverRuleResult

       CreateResolverEndpoint(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error)
       CreateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) *Route53resolverCreateResolverEndpointResult

       CreateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error)
       CreateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) *Route53resolverCreateResolverQueryLogConfigResult

       CreateResolverRule(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error)
       CreateResolverRuleAsync(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) *Route53resolverCreateResolverRuleResult

       DeleteResolverEndpoint(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error)
       DeleteResolverEndpointAsync(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) *Route53resolverDeleteResolverEndpointResult

       DeleteResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error)
       DeleteResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) *Route53resolverDeleteResolverQueryLogConfigResult

       DeleteResolverRule(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error)
       DeleteResolverRuleAsync(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) *Route53resolverDeleteResolverRuleResult

       DisassociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error)
       DisassociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) *Route53resolverDisassociateResolverEndpointIpAddressResult

       DisassociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error)
       DisassociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) *Route53resolverDisassociateResolverQueryLogConfigResult

       DisassociateResolverRule(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error)
       DisassociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) *Route53resolverDisassociateResolverRuleResult

       GetResolverEndpoint(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error)
       GetResolverEndpointAsync(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) *Route53resolverGetResolverEndpointResult

       GetResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error)
       GetResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) *Route53resolverGetResolverQueryLogConfigResult

       GetResolverQueryLogConfigAssociation(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error)
       GetResolverQueryLogConfigAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) *Route53resolverGetResolverQueryLogConfigAssociationResult

       GetResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error)
       GetResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) *Route53resolverGetResolverQueryLogConfigPolicyResult

       GetResolverRule(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error)
       GetResolverRuleAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) *Route53resolverGetResolverRuleResult

       GetResolverRuleAssociation(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error)
       GetResolverRuleAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) *Route53resolverGetResolverRuleAssociationResult

       GetResolverRulePolicy(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error)
       GetResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) *Route53resolverGetResolverRulePolicyResult

       ListResolverEndpointIpAddresses(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error)
       ListResolverEndpointIpAddressesAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) *Route53resolverListResolverEndpointIpAddressesResult

       ListResolverEndpoints(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error)
       ListResolverEndpointsAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) *Route53resolverListResolverEndpointsResult

       ListResolverQueryLogConfigAssociations(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error)
       ListResolverQueryLogConfigAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) *Route53resolverListResolverQueryLogConfigAssociationsResult

       ListResolverQueryLogConfigs(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error)
       ListResolverQueryLogConfigsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) *Route53resolverListResolverQueryLogConfigsResult

       ListResolverRuleAssociations(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error)
       ListResolverRuleAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) *Route53resolverListResolverRuleAssociationsResult

       ListResolverRules(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error)
       ListResolverRulesAsync(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) *Route53resolverListResolverRulesResult

       ListTagsForResource(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) *Route53resolverListTagsForResourceResult

       PutResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error)
       PutResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) *Route53resolverPutResolverQueryLogConfigPolicyResult

       PutResolverRulePolicy(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error)
       PutResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) *Route53resolverPutResolverRulePolicyResult

       TagResource(ctx workflow.Context, input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *route53resolver.TagResourceInput) *Route53resolverTagResourceResult

       UntagResource(ctx workflow.Context, input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *route53resolver.UntagResourceInput) *Route53resolverUntagResourceResult

       UpdateResolverEndpoint(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error)
       UpdateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) *Route53resolverUpdateResolverEndpointResult

       UpdateResolverRule(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error)
       UpdateResolverRuleAsync(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) *Route53resolverUpdateResolverRuleResult
}

type Route53resolverAssociateResolverEndpointIpAddressResult struct {
	Result workflow.Future
}

func (r *Route53resolverAssociateResolverEndpointIpAddressResult) Get(ctx workflow.Context) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
    var output route53resolver.AssociateResolverEndpointIpAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverAssociateResolverQueryLogConfigResult struct {
	Result workflow.Future
}

func (r *Route53resolverAssociateResolverQueryLogConfigResult) Get(ctx workflow.Context) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
    var output route53resolver.AssociateResolverQueryLogConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverAssociateResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverAssociateResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.AssociateResolverRuleOutput, error) {
    var output route53resolver.AssociateResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverCreateResolverEndpointResult struct {
	Result workflow.Future
}

func (r *Route53resolverCreateResolverEndpointResult) Get(ctx workflow.Context) (*route53resolver.CreateResolverEndpointOutput, error) {
    var output route53resolver.CreateResolverEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverCreateResolverQueryLogConfigResult struct {
	Result workflow.Future
}

func (r *Route53resolverCreateResolverQueryLogConfigResult) Get(ctx workflow.Context) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
    var output route53resolver.CreateResolverQueryLogConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverCreateResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverCreateResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.CreateResolverRuleOutput, error) {
    var output route53resolver.CreateResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDeleteResolverEndpointResult struct {
	Result workflow.Future
}

func (r *Route53resolverDeleteResolverEndpointResult) Get(ctx workflow.Context) (*route53resolver.DeleteResolverEndpointOutput, error) {
    var output route53resolver.DeleteResolverEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDeleteResolverQueryLogConfigResult struct {
	Result workflow.Future
}

func (r *Route53resolverDeleteResolverQueryLogConfigResult) Get(ctx workflow.Context) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
    var output route53resolver.DeleteResolverQueryLogConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDeleteResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverDeleteResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.DeleteResolverRuleOutput, error) {
    var output route53resolver.DeleteResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDisassociateResolverEndpointIpAddressResult struct {
	Result workflow.Future
}

func (r *Route53resolverDisassociateResolverEndpointIpAddressResult) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
    var output route53resolver.DisassociateResolverEndpointIpAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDisassociateResolverQueryLogConfigResult struct {
	Result workflow.Future
}

func (r *Route53resolverDisassociateResolverQueryLogConfigResult) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
    var output route53resolver.DisassociateResolverQueryLogConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverDisassociateResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverDisassociateResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.DisassociateResolverRuleOutput, error) {
    var output route53resolver.DisassociateResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverEndpointResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverEndpointResult) Get(ctx workflow.Context) (*route53resolver.GetResolverEndpointOutput, error) {
    var output route53resolver.GetResolverEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverQueryLogConfigResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverQueryLogConfigResult) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverQueryLogConfigAssociationResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverQueryLogConfigAssociationResult) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverQueryLogConfigPolicyResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverQueryLogConfigPolicyResult) Get(ctx workflow.Context) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.GetResolverRuleOutput, error) {
    var output route53resolver.GetResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverRuleAssociationResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverRuleAssociationResult) Get(ctx workflow.Context) (*route53resolver.GetResolverRuleAssociationOutput, error) {
    var output route53resolver.GetResolverRuleAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverGetResolverRulePolicyResult struct {
	Result workflow.Future
}

func (r *Route53resolverGetResolverRulePolicyResult) Get(ctx workflow.Context) (*route53resolver.GetResolverRulePolicyOutput, error) {
    var output route53resolver.GetResolverRulePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverEndpointIpAddressesResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverEndpointIpAddressesResult) Get(ctx workflow.Context) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
    var output route53resolver.ListResolverEndpointIpAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverEndpointsResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverEndpointsResult) Get(ctx workflow.Context) (*route53resolver.ListResolverEndpointsOutput, error) {
    var output route53resolver.ListResolverEndpointsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverQueryLogConfigAssociationsResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverQueryLogConfigAssociationsResult) Get(ctx workflow.Context) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
    var output route53resolver.ListResolverQueryLogConfigAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverQueryLogConfigsResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverQueryLogConfigsResult) Get(ctx workflow.Context) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
    var output route53resolver.ListResolverQueryLogConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverRuleAssociationsResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverRuleAssociationsResult) Get(ctx workflow.Context) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
    var output route53resolver.ListResolverRuleAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListResolverRulesResult struct {
	Result workflow.Future
}

func (r *Route53resolverListResolverRulesResult) Get(ctx workflow.Context) (*route53resolver.ListResolverRulesOutput, error) {
    var output route53resolver.ListResolverRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *Route53resolverListTagsForResourceResult) Get(ctx workflow.Context) (*route53resolver.ListTagsForResourceOutput, error) {
    var output route53resolver.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverPutResolverQueryLogConfigPolicyResult struct {
	Result workflow.Future
}

func (r *Route53resolverPutResolverQueryLogConfigPolicyResult) Get(ctx workflow.Context) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
    var output route53resolver.PutResolverQueryLogConfigPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverPutResolverRulePolicyResult struct {
	Result workflow.Future
}

func (r *Route53resolverPutResolverRulePolicyResult) Get(ctx workflow.Context) (*route53resolver.PutResolverRulePolicyOutput, error) {
    var output route53resolver.PutResolverRulePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverTagResourceResult struct {
	Result workflow.Future
}

func (r *Route53resolverTagResourceResult) Get(ctx workflow.Context) (*route53resolver.TagResourceOutput, error) {
    var output route53resolver.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverUntagResourceResult struct {
	Result workflow.Future
}

func (r *Route53resolverUntagResourceResult) Get(ctx workflow.Context) (*route53resolver.UntagResourceOutput, error) {
    var output route53resolver.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverUpdateResolverEndpointResult struct {
	Result workflow.Future
}

func (r *Route53resolverUpdateResolverEndpointResult) Get(ctx workflow.Context) (*route53resolver.UpdateResolverEndpointOutput, error) {
    var output route53resolver.UpdateResolverEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53resolverUpdateResolverRuleResult struct {
	Result workflow.Future
}

func (r *Route53resolverUpdateResolverRuleResult) Get(ctx workflow.Context) (*route53resolver.UpdateResolverRuleOutput, error) {
    var output route53resolver.UpdateResolverRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type Route53ResolverStub struct {
    activities awsactivities.Route53ResolverActivities
}

func NewRoute53ResolverStub() Route53ResolverClient {
    return &Route53ResolverStub{}
}

func (a *Route53ResolverStub) AssociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
    var output route53resolver.AssociateResolverEndpointIpAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverEndpointIpAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) AssociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.AssociateResolverEndpointIpAddressInput) *Route53resolverAssociateResolverEndpointIpAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverEndpointIpAddress, input)
    return &Route53resolverAssociateResolverEndpointIpAddressResult{Result: future}
}

func (a *Route53ResolverStub) AssociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error) {
    var output route53resolver.AssociateResolverQueryLogConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverQueryLogConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) AssociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.AssociateResolverQueryLogConfigInput) *Route53resolverAssociateResolverQueryLogConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverQueryLogConfig, input)
    return &Route53resolverAssociateResolverQueryLogConfigResult{Result: future}
}

func (a *Route53ResolverStub) AssociateResolverRule(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error) {
    var output route53resolver.AssociateResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) AssociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.AssociateResolverRuleInput) *Route53resolverAssociateResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AssociateResolverRule, input)
    return &Route53resolverAssociateResolverRuleResult{Result: future}
}

func (a *Route53ResolverStub) CreateResolverEndpoint(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error) {
    var output route53resolver.CreateResolverEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResolverEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) CreateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.CreateResolverEndpointInput) *Route53resolverCreateResolverEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateResolverEndpoint, input)
    return &Route53resolverCreateResolverEndpointResult{Result: future}
}

func (a *Route53ResolverStub) CreateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error) {
    var output route53resolver.CreateResolverQueryLogConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResolverQueryLogConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) CreateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.CreateResolverQueryLogConfigInput) *Route53resolverCreateResolverQueryLogConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateResolverQueryLogConfig, input)
    return &Route53resolverCreateResolverQueryLogConfigResult{Result: future}
}

func (a *Route53ResolverStub) CreateResolverRule(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error) {
    var output route53resolver.CreateResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) CreateResolverRuleAsync(ctx workflow.Context, input *route53resolver.CreateResolverRuleInput) *Route53resolverCreateResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateResolverRule, input)
    return &Route53resolverCreateResolverRuleResult{Result: future}
}

func (a *Route53ResolverStub) DeleteResolverEndpoint(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error) {
    var output route53resolver.DeleteResolverEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DeleteResolverEndpointAsync(ctx workflow.Context, input *route53resolver.DeleteResolverEndpointInput) *Route53resolverDeleteResolverEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverEndpoint, input)
    return &Route53resolverDeleteResolverEndpointResult{Result: future}
}

func (a *Route53ResolverStub) DeleteResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error) {
    var output route53resolver.DeleteResolverQueryLogConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverQueryLogConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DeleteResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DeleteResolverQueryLogConfigInput) *Route53resolverDeleteResolverQueryLogConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverQueryLogConfig, input)
    return &Route53resolverDeleteResolverQueryLogConfigResult{Result: future}
}

func (a *Route53ResolverStub) DeleteResolverRule(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error) {
    var output route53resolver.DeleteResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DeleteResolverRuleAsync(ctx workflow.Context, input *route53resolver.DeleteResolverRuleInput) *Route53resolverDeleteResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteResolverRule, input)
    return &Route53resolverDeleteResolverRuleResult{Result: future}
}

func (a *Route53ResolverStub) DisassociateResolverEndpointIpAddress(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error) {
    var output route53resolver.DisassociateResolverEndpointIpAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverEndpointIpAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DisassociateResolverEndpointIpAddressAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverEndpointIpAddressInput) *Route53resolverDisassociateResolverEndpointIpAddressResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverEndpointIpAddress, input)
    return &Route53resolverDisassociateResolverEndpointIpAddressResult{Result: future}
}

func (a *Route53ResolverStub) DisassociateResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error) {
    var output route53resolver.DisassociateResolverQueryLogConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverQueryLogConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DisassociateResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverQueryLogConfigInput) *Route53resolverDisassociateResolverQueryLogConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverQueryLogConfig, input)
    return &Route53resolverDisassociateResolverQueryLogConfigResult{Result: future}
}

func (a *Route53ResolverStub) DisassociateResolverRule(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error) {
    var output route53resolver.DisassociateResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) DisassociateResolverRuleAsync(ctx workflow.Context, input *route53resolver.DisassociateResolverRuleInput) *Route53resolverDisassociateResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DisassociateResolverRule, input)
    return &Route53resolverDisassociateResolverRuleResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverEndpoint(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error) {
    var output route53resolver.GetResolverEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverEndpointAsync(ctx workflow.Context, input *route53resolver.GetResolverEndpointInput) *Route53resolverGetResolverEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverEndpoint, input)
    return &Route53resolverGetResolverEndpointResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverQueryLogConfig(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverQueryLogConfigAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigInput) *Route53resolverGetResolverQueryLogConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfig, input)
    return &Route53resolverGetResolverQueryLogConfigResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverQueryLogConfigAssociation(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfigAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverQueryLogConfigAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigAssociationInput) *Route53resolverGetResolverQueryLogConfigAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfigAssociation, input)
    return &Route53resolverGetResolverQueryLogConfigAssociationResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error) {
    var output route53resolver.GetResolverQueryLogConfigPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfigPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverQueryLogConfigPolicyInput) *Route53resolverGetResolverQueryLogConfigPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverQueryLogConfigPolicy, input)
    return &Route53resolverGetResolverQueryLogConfigPolicyResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverRule(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error) {
    var output route53resolver.GetResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverRuleAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleInput) *Route53resolverGetResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverRule, input)
    return &Route53resolverGetResolverRuleResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverRuleAssociation(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error) {
    var output route53resolver.GetResolverRuleAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverRuleAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverRuleAssociationAsync(ctx workflow.Context, input *route53resolver.GetResolverRuleAssociationInput) *Route53resolverGetResolverRuleAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverRuleAssociation, input)
    return &Route53resolverGetResolverRuleAssociationResult{Result: future}
}

func (a *Route53ResolverStub) GetResolverRulePolicy(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error) {
    var output route53resolver.GetResolverRulePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetResolverRulePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) GetResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.GetResolverRulePolicyInput) *Route53resolverGetResolverRulePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetResolverRulePolicy, input)
    return &Route53resolverGetResolverRulePolicyResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverEndpointIpAddresses(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error) {
    var output route53resolver.ListResolverEndpointIpAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverEndpointIpAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverEndpointIpAddressesAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointIpAddressesInput) *Route53resolverListResolverEndpointIpAddressesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverEndpointIpAddresses, input)
    return &Route53resolverListResolverEndpointIpAddressesResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverEndpoints(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error) {
    var output route53resolver.ListResolverEndpointsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverEndpoints, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverEndpointsAsync(ctx workflow.Context, input *route53resolver.ListResolverEndpointsInput) *Route53resolverListResolverEndpointsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverEndpoints, input)
    return &Route53resolverListResolverEndpointsResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverQueryLogConfigAssociations(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error) {
    var output route53resolver.ListResolverQueryLogConfigAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverQueryLogConfigAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverQueryLogConfigAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigAssociationsInput) *Route53resolverListResolverQueryLogConfigAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverQueryLogConfigAssociations, input)
    return &Route53resolverListResolverQueryLogConfigAssociationsResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverQueryLogConfigs(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error) {
    var output route53resolver.ListResolverQueryLogConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverQueryLogConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverQueryLogConfigsAsync(ctx workflow.Context, input *route53resolver.ListResolverQueryLogConfigsInput) *Route53resolverListResolverQueryLogConfigsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverQueryLogConfigs, input)
    return &Route53resolverListResolverQueryLogConfigsResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverRuleAssociations(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error) {
    var output route53resolver.ListResolverRuleAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverRuleAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverRuleAssociationsAsync(ctx workflow.Context, input *route53resolver.ListResolverRuleAssociationsInput) *Route53resolverListResolverRuleAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverRuleAssociations, input)
    return &Route53resolverListResolverRuleAssociationsResult{Result: future}
}

func (a *Route53ResolverStub) ListResolverRules(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error) {
    var output route53resolver.ListResolverRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListResolverRules, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListResolverRulesAsync(ctx workflow.Context, input *route53resolver.ListResolverRulesInput) *Route53resolverListResolverRulesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListResolverRules, input)
    return &Route53resolverListResolverRulesResult{Result: future}
}

func (a *Route53ResolverStub) ListTagsForResource(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error) {
    var output route53resolver.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) ListTagsForResourceAsync(ctx workflow.Context, input *route53resolver.ListTagsForResourceInput) *Route53resolverListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &Route53resolverListTagsForResourceResult{Result: future}
}

func (a *Route53ResolverStub) PutResolverQueryLogConfigPolicy(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error) {
    var output route53resolver.PutResolverQueryLogConfigPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResolverQueryLogConfigPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) PutResolverQueryLogConfigPolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverQueryLogConfigPolicyInput) *Route53resolverPutResolverQueryLogConfigPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResolverQueryLogConfigPolicy, input)
    return &Route53resolverPutResolverQueryLogConfigPolicyResult{Result: future}
}

func (a *Route53ResolverStub) PutResolverRulePolicy(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error) {
    var output route53resolver.PutResolverRulePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutResolverRulePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) PutResolverRulePolicyAsync(ctx workflow.Context, input *route53resolver.PutResolverRulePolicyInput) *Route53resolverPutResolverRulePolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutResolverRulePolicy, input)
    return &Route53resolverPutResolverRulePolicyResult{Result: future}
}

func (a *Route53ResolverStub) TagResource(ctx workflow.Context, input *route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error) {
    var output route53resolver.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) TagResourceAsync(ctx workflow.Context, input *route53resolver.TagResourceInput) *Route53resolverTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &Route53resolverTagResourceResult{Result: future}
}

func (a *Route53ResolverStub) UntagResource(ctx workflow.Context, input *route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error) {
    var output route53resolver.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) UntagResourceAsync(ctx workflow.Context, input *route53resolver.UntagResourceInput) *Route53resolverUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &Route53resolverUntagResourceResult{Result: future}
}

func (a *Route53ResolverStub) UpdateResolverEndpoint(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error) {
    var output route53resolver.UpdateResolverEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResolverEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) UpdateResolverEndpointAsync(ctx workflow.Context, input *route53resolver.UpdateResolverEndpointInput) *Route53resolverUpdateResolverEndpointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateResolverEndpoint, input)
    return &Route53resolverUpdateResolverEndpointResult{Result: future}
}

func (a *Route53ResolverStub) UpdateResolverRule(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error) {
    var output route53resolver.UpdateResolverRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateResolverRule, input).Get(ctx, &output)
    return &output, err
}

func (a *Route53ResolverStub) UpdateResolverRuleAsync(ctx workflow.Context, input *route53resolver.UpdateResolverRuleInput) *Route53resolverUpdateResolverRuleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateResolverRule, input)
    return &Route53resolverUpdateResolverRuleResult{Result: future}
}
