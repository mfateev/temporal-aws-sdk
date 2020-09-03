package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/eventbridge"
	"go.temporal.io/sdk/workflow"
)

type EventBridgeClient interface {
    ActivateEventSource(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error)
    ActivateEventSourceAsync(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) *EventbridgeActivateEventSourceResult

    CreateEventBus(ctx workflow.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error)
    CreateEventBusAsync(ctx workflow.Context, input *eventbridge.CreateEventBusInput) *EventbridgeCreateEventBusResult

    CreatePartnerEventSource(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error)
    CreatePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) *EventbridgeCreatePartnerEventSourceResult

    DeactivateEventSource(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error)
    DeactivateEventSourceAsync(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) *EventbridgeDeactivateEventSourceResult

    DeleteEventBus(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error)
    DeleteEventBusAsync(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) *EventbridgeDeleteEventBusResult

    DeletePartnerEventSource(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error)
    DeletePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) *EventbridgeDeletePartnerEventSourceResult

    DeleteRule(ctx workflow.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error)
    DeleteRuleAsync(ctx workflow.Context, input *eventbridge.DeleteRuleInput) *EventbridgeDeleteRuleResult

    DescribeEventBus(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error)
    DescribeEventBusAsync(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) *EventbridgeDescribeEventBusResult

    DescribeEventSource(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error)
    DescribeEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) *EventbridgeDescribeEventSourceResult

    DescribePartnerEventSource(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error)
    DescribePartnerEventSourceAsync(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) *EventbridgeDescribePartnerEventSourceResult

    DescribeRule(ctx workflow.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error)
    DescribeRuleAsync(ctx workflow.Context, input *eventbridge.DescribeRuleInput) *EventbridgeDescribeRuleResult

    DisableRule(ctx workflow.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error)
    DisableRuleAsync(ctx workflow.Context, input *eventbridge.DisableRuleInput) *EventbridgeDisableRuleResult

    EnableRule(ctx workflow.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error)
    EnableRuleAsync(ctx workflow.Context, input *eventbridge.EnableRuleInput) *EventbridgeEnableRuleResult

    ListEventBuses(ctx workflow.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error)
    ListEventBusesAsync(ctx workflow.Context, input *eventbridge.ListEventBusesInput) *EventbridgeListEventBusesResult

    ListEventSources(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error)
    ListEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) *EventbridgeListEventSourcesResult

    ListPartnerEventSourceAccounts(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error)
    ListPartnerEventSourceAccountsAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) *EventbridgeListPartnerEventSourceAccountsResult

    ListPartnerEventSources(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error)
    ListPartnerEventSourcesAsync(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) *EventbridgeListPartnerEventSourcesResult

    ListRuleNamesByTarget(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error)
    ListRuleNamesByTargetAsync(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) *EventbridgeListRuleNamesByTargetResult

    ListRules(ctx workflow.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error)
    ListRulesAsync(ctx workflow.Context, input *eventbridge.ListRulesInput) *EventbridgeListRulesResult

    ListTagsForResource(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) *EventbridgeListTagsForResourceResult

    ListTargetsByRule(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error)
    ListTargetsByRuleAsync(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) *EventbridgeListTargetsByRuleResult

    PutEvents(ctx workflow.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error)
    PutEventsAsync(ctx workflow.Context, input *eventbridge.PutEventsInput) *EventbridgePutEventsResult

    PutPartnerEvents(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error)
    PutPartnerEventsAsync(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) *EventbridgePutPartnerEventsResult

    PutPermission(ctx workflow.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error)
    PutPermissionAsync(ctx workflow.Context, input *eventbridge.PutPermissionInput) *EventbridgePutPermissionResult

    PutRule(ctx workflow.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error)
    PutRuleAsync(ctx workflow.Context, input *eventbridge.PutRuleInput) *EventbridgePutRuleResult

    PutTargets(ctx workflow.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error)
    PutTargetsAsync(ctx workflow.Context, input *eventbridge.PutTargetsInput) *EventbridgePutTargetsResult

    RemovePermission(ctx workflow.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error)
    RemovePermissionAsync(ctx workflow.Context, input *eventbridge.RemovePermissionInput) *EventbridgeRemovePermissionResult

    RemoveTargets(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error)
    RemoveTargetsAsync(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) *EventbridgeRemoveTargetsResult

    TagResource(ctx workflow.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *eventbridge.TagResourceInput) *EventbridgeTagResourceResult

    TestEventPattern(ctx workflow.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error)
    TestEventPatternAsync(ctx workflow.Context, input *eventbridge.TestEventPatternInput) *EventbridgeTestEventPatternResult

    UntagResource(ctx workflow.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *eventbridge.UntagResourceInput) *EventbridgeUntagResourceResult
}
type EventbridgeActivateEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeActivateEventSourceResult) Get(ctx workflow.Context) (*eventbridge.ActivateEventSourceOutput, error) {
    var output eventbridge.ActivateEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeCreateEventBusResult struct {
	Result workflow.Future
}

func (r *EventbridgeCreateEventBusResult) Get(ctx workflow.Context) (*eventbridge.CreateEventBusOutput, error) {
    var output eventbridge.CreateEventBusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeCreatePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeCreatePartnerEventSourceResult) Get(ctx workflow.Context) (*eventbridge.CreatePartnerEventSourceOutput, error) {
    var output eventbridge.CreatePartnerEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDeactivateEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeDeactivateEventSourceResult) Get(ctx workflow.Context) (*eventbridge.DeactivateEventSourceOutput, error) {
    var output eventbridge.DeactivateEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDeleteEventBusResult struct {
	Result workflow.Future
}

func (r *EventbridgeDeleteEventBusResult) Get(ctx workflow.Context) (*eventbridge.DeleteEventBusOutput, error) {
    var output eventbridge.DeleteEventBusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDeletePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeDeletePartnerEventSourceResult) Get(ctx workflow.Context) (*eventbridge.DeletePartnerEventSourceOutput, error) {
    var output eventbridge.DeletePartnerEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDeleteRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgeDeleteRuleResult) Get(ctx workflow.Context) (*eventbridge.DeleteRuleOutput, error) {
    var output eventbridge.DeleteRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDescribeEventBusResult struct {
	Result workflow.Future
}

func (r *EventbridgeDescribeEventBusResult) Get(ctx workflow.Context) (*eventbridge.DescribeEventBusOutput, error) {
    var output eventbridge.DescribeEventBusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDescribeEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeDescribeEventSourceResult) Get(ctx workflow.Context) (*eventbridge.DescribeEventSourceOutput, error) {
    var output eventbridge.DescribeEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDescribePartnerEventSourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeDescribePartnerEventSourceResult) Get(ctx workflow.Context) (*eventbridge.DescribePartnerEventSourceOutput, error) {
    var output eventbridge.DescribePartnerEventSourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDescribeRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgeDescribeRuleResult) Get(ctx workflow.Context) (*eventbridge.DescribeRuleOutput, error) {
    var output eventbridge.DescribeRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeDisableRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgeDisableRuleResult) Get(ctx workflow.Context) (*eventbridge.DisableRuleOutput, error) {
    var output eventbridge.DisableRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeEnableRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgeEnableRuleResult) Get(ctx workflow.Context) (*eventbridge.EnableRuleOutput, error) {
    var output eventbridge.EnableRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListEventBusesResult struct {
	Result workflow.Future
}

func (r *EventbridgeListEventBusesResult) Get(ctx workflow.Context) (*eventbridge.ListEventBusesOutput, error) {
    var output eventbridge.ListEventBusesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListEventSourcesResult struct {
	Result workflow.Future
}

func (r *EventbridgeListEventSourcesResult) Get(ctx workflow.Context) (*eventbridge.ListEventSourcesOutput, error) {
    var output eventbridge.ListEventSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListPartnerEventSourceAccountsResult struct {
	Result workflow.Future
}

func (r *EventbridgeListPartnerEventSourceAccountsResult) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
    var output eventbridge.ListPartnerEventSourceAccountsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListPartnerEventSourcesResult struct {
	Result workflow.Future
}

func (r *EventbridgeListPartnerEventSourcesResult) Get(ctx workflow.Context) (*eventbridge.ListPartnerEventSourcesOutput, error) {
    var output eventbridge.ListPartnerEventSourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListRuleNamesByTargetResult struct {
	Result workflow.Future
}

func (r *EventbridgeListRuleNamesByTargetResult) Get(ctx workflow.Context) (*eventbridge.ListRuleNamesByTargetOutput, error) {
    var output eventbridge.ListRuleNamesByTargetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListRulesResult struct {
	Result workflow.Future
}

func (r *EventbridgeListRulesResult) Get(ctx workflow.Context) (*eventbridge.ListRulesOutput, error) {
    var output eventbridge.ListRulesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeListTagsForResourceResult) Get(ctx workflow.Context) (*eventbridge.ListTagsForResourceOutput, error) {
    var output eventbridge.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeListTargetsByRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgeListTargetsByRuleResult) Get(ctx workflow.Context) (*eventbridge.ListTargetsByRuleOutput, error) {
    var output eventbridge.ListTargetsByRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgePutEventsResult struct {
	Result workflow.Future
}

func (r *EventbridgePutEventsResult) Get(ctx workflow.Context) (*eventbridge.PutEventsOutput, error) {
    var output eventbridge.PutEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgePutPartnerEventsResult struct {
	Result workflow.Future
}

func (r *EventbridgePutPartnerEventsResult) Get(ctx workflow.Context) (*eventbridge.PutPartnerEventsOutput, error) {
    var output eventbridge.PutPartnerEventsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgePutPermissionResult struct {
	Result workflow.Future
}

func (r *EventbridgePutPermissionResult) Get(ctx workflow.Context) (*eventbridge.PutPermissionOutput, error) {
    var output eventbridge.PutPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgePutRuleResult struct {
	Result workflow.Future
}

func (r *EventbridgePutRuleResult) Get(ctx workflow.Context) (*eventbridge.PutRuleOutput, error) {
    var output eventbridge.PutRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgePutTargetsResult struct {
	Result workflow.Future
}

func (r *EventbridgePutTargetsResult) Get(ctx workflow.Context) (*eventbridge.PutTargetsOutput, error) {
    var output eventbridge.PutTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeRemovePermissionResult struct {
	Result workflow.Future
}

func (r *EventbridgeRemovePermissionResult) Get(ctx workflow.Context) (*eventbridge.RemovePermissionOutput, error) {
    var output eventbridge.RemovePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeRemoveTargetsResult struct {
	Result workflow.Future
}

func (r *EventbridgeRemoveTargetsResult) Get(ctx workflow.Context) (*eventbridge.RemoveTargetsOutput, error) {
    var output eventbridge.RemoveTargetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeTagResourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeTagResourceResult) Get(ctx workflow.Context) (*eventbridge.TagResourceOutput, error) {
    var output eventbridge.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeTestEventPatternResult struct {
	Result workflow.Future
}

func (r *EventbridgeTestEventPatternResult) Get(ctx workflow.Context) (*eventbridge.TestEventPatternOutput, error) {
    var output eventbridge.TestEventPatternOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EventbridgeUntagResourceResult struct {
	Result workflow.Future
}

func (r *EventbridgeUntagResourceResult) Get(ctx workflow.Context) (*eventbridge.UntagResourceOutput, error) {
    var output eventbridge.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type EventBridgeStub struct {
    activities EventBridgeClient
}

func NewEventBridgeStub() EventBridgeClient {
    return &EventBridgeStub{}
}

func (a *EventBridgeStub) ActivateEventSource(ctx workflow.Context, input *eventbridge.ActivateEventSourceInput) (*eventbridge.ActivateEventSourceOutput, error) {
    var output eventbridge.ActivateEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ActivateEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) CreateEventBus(ctx workflow.Context, input *eventbridge.CreateEventBusInput) (*eventbridge.CreateEventBusOutput, error) {
    var output eventbridge.CreateEventBusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEventBus, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) CreatePartnerEventSource(ctx workflow.Context, input *eventbridge.CreatePartnerEventSourceInput) (*eventbridge.CreatePartnerEventSourceOutput, error) {
    var output eventbridge.CreatePartnerEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreatePartnerEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DeactivateEventSource(ctx workflow.Context, input *eventbridge.DeactivateEventSourceInput) (*eventbridge.DeactivateEventSourceOutput, error) {
    var output eventbridge.DeactivateEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeactivateEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DeleteEventBus(ctx workflow.Context, input *eventbridge.DeleteEventBusInput) (*eventbridge.DeleteEventBusOutput, error) {
    var output eventbridge.DeleteEventBusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEventBus, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DeletePartnerEventSource(ctx workflow.Context, input *eventbridge.DeletePartnerEventSourceInput) (*eventbridge.DeletePartnerEventSourceOutput, error) {
    var output eventbridge.DeletePartnerEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePartnerEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DeleteRule(ctx workflow.Context, input *eventbridge.DeleteRuleInput) (*eventbridge.DeleteRuleOutput, error) {
    var output eventbridge.DeleteRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DescribeEventBus(ctx workflow.Context, input *eventbridge.DescribeEventBusInput) (*eventbridge.DescribeEventBusOutput, error) {
    var output eventbridge.DescribeEventBusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventBus, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DescribeEventSource(ctx workflow.Context, input *eventbridge.DescribeEventSourceInput) (*eventbridge.DescribeEventSourceOutput, error) {
    var output eventbridge.DescribeEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DescribePartnerEventSource(ctx workflow.Context, input *eventbridge.DescribePartnerEventSourceInput) (*eventbridge.DescribePartnerEventSourceOutput, error) {
    var output eventbridge.DescribePartnerEventSourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePartnerEventSource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DescribeRule(ctx workflow.Context, input *eventbridge.DescribeRuleInput) (*eventbridge.DescribeRuleOutput, error) {
    var output eventbridge.DescribeRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) DisableRule(ctx workflow.Context, input *eventbridge.DisableRuleInput) (*eventbridge.DisableRuleOutput, error) {
    var output eventbridge.DisableRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisableRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) EnableRule(ctx workflow.Context, input *eventbridge.EnableRuleInput) (*eventbridge.EnableRuleOutput, error) {
    var output eventbridge.EnableRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.EnableRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListEventBuses(ctx workflow.Context, input *eventbridge.ListEventBusesInput) (*eventbridge.ListEventBusesOutput, error) {
    var output eventbridge.ListEventBusesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventBuses, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListEventSources(ctx workflow.Context, input *eventbridge.ListEventSourcesInput) (*eventbridge.ListEventSourcesOutput, error) {
    var output eventbridge.ListEventSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEventSources, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListPartnerEventSourceAccounts(ctx workflow.Context, input *eventbridge.ListPartnerEventSourceAccountsInput) (*eventbridge.ListPartnerEventSourceAccountsOutput, error) {
    var output eventbridge.ListPartnerEventSourceAccountsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPartnerEventSourceAccounts, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListPartnerEventSources(ctx workflow.Context, input *eventbridge.ListPartnerEventSourcesInput) (*eventbridge.ListPartnerEventSourcesOutput, error) {
    var output eventbridge.ListPartnerEventSourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPartnerEventSources, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListRuleNamesByTarget(ctx workflow.Context, input *eventbridge.ListRuleNamesByTargetInput) (*eventbridge.ListRuleNamesByTargetOutput, error) {
    var output eventbridge.ListRuleNamesByTargetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRuleNamesByTarget, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListRules(ctx workflow.Context, input *eventbridge.ListRulesInput) (*eventbridge.ListRulesOutput, error) {
    var output eventbridge.ListRulesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRules, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListTagsForResource(ctx workflow.Context, input *eventbridge.ListTagsForResourceInput) (*eventbridge.ListTagsForResourceOutput, error) {
    var output eventbridge.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) ListTargetsByRule(ctx workflow.Context, input *eventbridge.ListTargetsByRuleInput) (*eventbridge.ListTargetsByRuleOutput, error) {
    var output eventbridge.ListTargetsByRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTargetsByRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) PutEvents(ctx workflow.Context, input *eventbridge.PutEventsInput) (*eventbridge.PutEventsOutput, error) {
    var output eventbridge.PutEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) PutPartnerEvents(ctx workflow.Context, input *eventbridge.PutPartnerEventsInput) (*eventbridge.PutPartnerEventsOutput, error) {
    var output eventbridge.PutPartnerEventsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPartnerEvents, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) PutPermission(ctx workflow.Context, input *eventbridge.PutPermissionInput) (*eventbridge.PutPermissionOutput, error) {
    var output eventbridge.PutPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) PutRule(ctx workflow.Context, input *eventbridge.PutRuleInput) (*eventbridge.PutRuleOutput, error) {
    var output eventbridge.PutRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRule, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) PutTargets(ctx workflow.Context, input *eventbridge.PutTargetsInput) (*eventbridge.PutTargetsOutput, error) {
    var output eventbridge.PutTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) RemovePermission(ctx workflow.Context, input *eventbridge.RemovePermissionInput) (*eventbridge.RemovePermissionOutput, error) {
    var output eventbridge.RemovePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) RemoveTargets(ctx workflow.Context, input *eventbridge.RemoveTargetsInput) (*eventbridge.RemoveTargetsOutput, error) {
    var output eventbridge.RemoveTargetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemoveTargets, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) TagResource(ctx workflow.Context, input *eventbridge.TagResourceInput) (*eventbridge.TagResourceOutput, error) {
    var output eventbridge.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) TestEventPattern(ctx workflow.Context, input *eventbridge.TestEventPatternInput) (*eventbridge.TestEventPatternOutput, error) {
    var output eventbridge.TestEventPatternOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestEventPattern, input).Get(ctx, &output)
    return &output, err
}

func (a *EventBridgeStub) UntagResource(ctx workflow.Context, input *eventbridge.UntagResourceInput) (*eventbridge.UntagResourceOutput, error) {
    var output eventbridge.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}
