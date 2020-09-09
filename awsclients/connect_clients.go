package awsclients

import (
	"github.com/aws/aws-sdk-go/service/connect"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type ConnectClient interface {
       CreateUser(ctx workflow.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error)
       CreateUserAsync(ctx workflow.Context, input *connect.CreateUserInput) *ConnectCreateUserResult

       DeleteUser(ctx workflow.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error)
       DeleteUserAsync(ctx workflow.Context, input *connect.DeleteUserInput) *ConnectDeleteUserResult

       DescribeUser(ctx workflow.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error)
       DescribeUserAsync(ctx workflow.Context, input *connect.DescribeUserInput) *ConnectDescribeUserResult

       DescribeUserHierarchyGroup(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error)
       DescribeUserHierarchyGroupAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) *ConnectDescribeUserHierarchyGroupResult

       DescribeUserHierarchyStructure(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error)
       DescribeUserHierarchyStructureAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) *ConnectDescribeUserHierarchyStructureResult

       GetContactAttributes(ctx workflow.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error)
       GetContactAttributesAsync(ctx workflow.Context, input *connect.GetContactAttributesInput) *ConnectGetContactAttributesResult

       GetCurrentMetricData(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error)
       GetCurrentMetricDataAsync(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) *ConnectGetCurrentMetricDataResult

       GetFederationToken(ctx workflow.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error)
       GetFederationTokenAsync(ctx workflow.Context, input *connect.GetFederationTokenInput) *ConnectGetFederationTokenResult

       GetMetricData(ctx workflow.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error)
       GetMetricDataAsync(ctx workflow.Context, input *connect.GetMetricDataInput) *ConnectGetMetricDataResult

       ListContactFlows(ctx workflow.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error)
       ListContactFlowsAsync(ctx workflow.Context, input *connect.ListContactFlowsInput) *ConnectListContactFlowsResult

       ListHoursOfOperations(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error)
       ListHoursOfOperationsAsync(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) *ConnectListHoursOfOperationsResult

       ListPhoneNumbers(ctx workflow.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error)
       ListPhoneNumbersAsync(ctx workflow.Context, input *connect.ListPhoneNumbersInput) *ConnectListPhoneNumbersResult

       ListQueues(ctx workflow.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error)
       ListQueuesAsync(ctx workflow.Context, input *connect.ListQueuesInput) *ConnectListQueuesResult

       ListRoutingProfiles(ctx workflow.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error)
       ListRoutingProfilesAsync(ctx workflow.Context, input *connect.ListRoutingProfilesInput) *ConnectListRoutingProfilesResult

       ListSecurityProfiles(ctx workflow.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error)
       ListSecurityProfilesAsync(ctx workflow.Context, input *connect.ListSecurityProfilesInput) *ConnectListSecurityProfilesResult

       ListTagsForResource(ctx workflow.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error)
       ListTagsForResourceAsync(ctx workflow.Context, input *connect.ListTagsForResourceInput) *ConnectListTagsForResourceResult

       ListUserHierarchyGroups(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error)
       ListUserHierarchyGroupsAsync(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) *ConnectListUserHierarchyGroupsResult

       ListUsers(ctx workflow.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error)
       ListUsersAsync(ctx workflow.Context, input *connect.ListUsersInput) *ConnectListUsersResult

       ResumeContactRecording(ctx workflow.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error)
       ResumeContactRecordingAsync(ctx workflow.Context, input *connect.ResumeContactRecordingInput) *ConnectResumeContactRecordingResult

       StartChatContact(ctx workflow.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error)
       StartChatContactAsync(ctx workflow.Context, input *connect.StartChatContactInput) *ConnectStartChatContactResult

       StartContactRecording(ctx workflow.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error)
       StartContactRecordingAsync(ctx workflow.Context, input *connect.StartContactRecordingInput) *ConnectStartContactRecordingResult

       StartOutboundVoiceContact(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error)
       StartOutboundVoiceContactAsync(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) *ConnectStartOutboundVoiceContactResult

       StopContact(ctx workflow.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error)
       StopContactAsync(ctx workflow.Context, input *connect.StopContactInput) *ConnectStopContactResult

       StopContactRecording(ctx workflow.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error)
       StopContactRecordingAsync(ctx workflow.Context, input *connect.StopContactRecordingInput) *ConnectStopContactRecordingResult

       SuspendContactRecording(ctx workflow.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error)
       SuspendContactRecordingAsync(ctx workflow.Context, input *connect.SuspendContactRecordingInput) *ConnectSuspendContactRecordingResult

       TagResource(ctx workflow.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error)
       TagResourceAsync(ctx workflow.Context, input *connect.TagResourceInput) *ConnectTagResourceResult

       UntagResource(ctx workflow.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error)
       UntagResourceAsync(ctx workflow.Context, input *connect.UntagResourceInput) *ConnectUntagResourceResult

       UpdateContactAttributes(ctx workflow.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error)
       UpdateContactAttributesAsync(ctx workflow.Context, input *connect.UpdateContactAttributesInput) *ConnectUpdateContactAttributesResult

       UpdateUserHierarchy(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error)
       UpdateUserHierarchyAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) *ConnectUpdateUserHierarchyResult

       UpdateUserIdentityInfo(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error)
       UpdateUserIdentityInfoAsync(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) *ConnectUpdateUserIdentityInfoResult

       UpdateUserPhoneConfig(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error)
       UpdateUserPhoneConfigAsync(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) *ConnectUpdateUserPhoneConfigResult

       UpdateUserRoutingProfile(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error)
       UpdateUserRoutingProfileAsync(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) *ConnectUpdateUserRoutingProfileResult

       UpdateUserSecurityProfiles(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error)
       UpdateUserSecurityProfilesAsync(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) *ConnectUpdateUserSecurityProfilesResult
}

type ConnectCreateUserResult struct {
	Result workflow.Future
}

func (r *ConnectCreateUserResult) Get(ctx workflow.Context) (*connect.CreateUserOutput, error) {
    var output connect.CreateUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectDeleteUserResult struct {
	Result workflow.Future
}

func (r *ConnectDeleteUserResult) Get(ctx workflow.Context) (*connect.DeleteUserOutput, error) {
    var output connect.DeleteUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectDescribeUserResult struct {
	Result workflow.Future
}

func (r *ConnectDescribeUserResult) Get(ctx workflow.Context) (*connect.DescribeUserOutput, error) {
    var output connect.DescribeUserOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectDescribeUserHierarchyGroupResult struct {
	Result workflow.Future
}

func (r *ConnectDescribeUserHierarchyGroupResult) Get(ctx workflow.Context) (*connect.DescribeUserHierarchyGroupOutput, error) {
    var output connect.DescribeUserHierarchyGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectDescribeUserHierarchyStructureResult struct {
	Result workflow.Future
}

func (r *ConnectDescribeUserHierarchyStructureResult) Get(ctx workflow.Context) (*connect.DescribeUserHierarchyStructureOutput, error) {
    var output connect.DescribeUserHierarchyStructureOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectGetContactAttributesResult struct {
	Result workflow.Future
}

func (r *ConnectGetContactAttributesResult) Get(ctx workflow.Context) (*connect.GetContactAttributesOutput, error) {
    var output connect.GetContactAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectGetCurrentMetricDataResult struct {
	Result workflow.Future
}

func (r *ConnectGetCurrentMetricDataResult) Get(ctx workflow.Context) (*connect.GetCurrentMetricDataOutput, error) {
    var output connect.GetCurrentMetricDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectGetFederationTokenResult struct {
	Result workflow.Future
}

func (r *ConnectGetFederationTokenResult) Get(ctx workflow.Context) (*connect.GetFederationTokenOutput, error) {
    var output connect.GetFederationTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectGetMetricDataResult struct {
	Result workflow.Future
}

func (r *ConnectGetMetricDataResult) Get(ctx workflow.Context) (*connect.GetMetricDataOutput, error) {
    var output connect.GetMetricDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListContactFlowsResult struct {
	Result workflow.Future
}

func (r *ConnectListContactFlowsResult) Get(ctx workflow.Context) (*connect.ListContactFlowsOutput, error) {
    var output connect.ListContactFlowsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListHoursOfOperationsResult struct {
	Result workflow.Future
}

func (r *ConnectListHoursOfOperationsResult) Get(ctx workflow.Context) (*connect.ListHoursOfOperationsOutput, error) {
    var output connect.ListHoursOfOperationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListPhoneNumbersResult struct {
	Result workflow.Future
}

func (r *ConnectListPhoneNumbersResult) Get(ctx workflow.Context) (*connect.ListPhoneNumbersOutput, error) {
    var output connect.ListPhoneNumbersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListQueuesResult struct {
	Result workflow.Future
}

func (r *ConnectListQueuesResult) Get(ctx workflow.Context) (*connect.ListQueuesOutput, error) {
    var output connect.ListQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListRoutingProfilesResult struct {
	Result workflow.Future
}

func (r *ConnectListRoutingProfilesResult) Get(ctx workflow.Context) (*connect.ListRoutingProfilesOutput, error) {
    var output connect.ListRoutingProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListSecurityProfilesResult struct {
	Result workflow.Future
}

func (r *ConnectListSecurityProfilesResult) Get(ctx workflow.Context) (*connect.ListSecurityProfilesOutput, error) {
    var output connect.ListSecurityProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *ConnectListTagsForResourceResult) Get(ctx workflow.Context) (*connect.ListTagsForResourceOutput, error) {
    var output connect.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListUserHierarchyGroupsResult struct {
	Result workflow.Future
}

func (r *ConnectListUserHierarchyGroupsResult) Get(ctx workflow.Context) (*connect.ListUserHierarchyGroupsOutput, error) {
    var output connect.ListUserHierarchyGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectListUsersResult struct {
	Result workflow.Future
}

func (r *ConnectListUsersResult) Get(ctx workflow.Context) (*connect.ListUsersOutput, error) {
    var output connect.ListUsersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectResumeContactRecordingResult struct {
	Result workflow.Future
}

func (r *ConnectResumeContactRecordingResult) Get(ctx workflow.Context) (*connect.ResumeContactRecordingOutput, error) {
    var output connect.ResumeContactRecordingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStartChatContactResult struct {
	Result workflow.Future
}

func (r *ConnectStartChatContactResult) Get(ctx workflow.Context) (*connect.StartChatContactOutput, error) {
    var output connect.StartChatContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStartContactRecordingResult struct {
	Result workflow.Future
}

func (r *ConnectStartContactRecordingResult) Get(ctx workflow.Context) (*connect.StartContactRecordingOutput, error) {
    var output connect.StartContactRecordingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStartOutboundVoiceContactResult struct {
	Result workflow.Future
}

func (r *ConnectStartOutboundVoiceContactResult) Get(ctx workflow.Context) (*connect.StartOutboundVoiceContactOutput, error) {
    var output connect.StartOutboundVoiceContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStopContactResult struct {
	Result workflow.Future
}

func (r *ConnectStopContactResult) Get(ctx workflow.Context) (*connect.StopContactOutput, error) {
    var output connect.StopContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStopContactRecordingResult struct {
	Result workflow.Future
}

func (r *ConnectStopContactRecordingResult) Get(ctx workflow.Context) (*connect.StopContactRecordingOutput, error) {
    var output connect.StopContactRecordingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectSuspendContactRecordingResult struct {
	Result workflow.Future
}

func (r *ConnectSuspendContactRecordingResult) Get(ctx workflow.Context) (*connect.SuspendContactRecordingOutput, error) {
    var output connect.SuspendContactRecordingOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectTagResourceResult struct {
	Result workflow.Future
}

func (r *ConnectTagResourceResult) Get(ctx workflow.Context) (*connect.TagResourceOutput, error) {
    var output connect.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUntagResourceResult struct {
	Result workflow.Future
}

func (r *ConnectUntagResourceResult) Get(ctx workflow.Context) (*connect.UntagResourceOutput, error) {
    var output connect.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateContactAttributesResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateContactAttributesResult) Get(ctx workflow.Context) (*connect.UpdateContactAttributesOutput, error) {
    var output connect.UpdateContactAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateUserHierarchyResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateUserHierarchyResult) Get(ctx workflow.Context) (*connect.UpdateUserHierarchyOutput, error) {
    var output connect.UpdateUserHierarchyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateUserIdentityInfoResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateUserIdentityInfoResult) Get(ctx workflow.Context) (*connect.UpdateUserIdentityInfoOutput, error) {
    var output connect.UpdateUserIdentityInfoOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateUserPhoneConfigResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateUserPhoneConfigResult) Get(ctx workflow.Context) (*connect.UpdateUserPhoneConfigOutput, error) {
    var output connect.UpdateUserPhoneConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateUserRoutingProfileResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateUserRoutingProfileResult) Get(ctx workflow.Context) (*connect.UpdateUserRoutingProfileOutput, error) {
    var output connect.UpdateUserRoutingProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectUpdateUserSecurityProfilesResult struct {
	Result workflow.Future
}

func (r *ConnectUpdateUserSecurityProfilesResult) Get(ctx workflow.Context) (*connect.UpdateUserSecurityProfilesOutput, error) {
    var output connect.UpdateUserSecurityProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ConnectStub struct {
    activities awsactivities.ConnectActivities
}

func NewConnectStub() ConnectClient {
    return &ConnectStub{}
}

func (a *ConnectStub) CreateUser(ctx workflow.Context, input *connect.CreateUserInput) (*connect.CreateUserOutput, error) {
    var output connect.CreateUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) CreateUserAsync(ctx workflow.Context, input *connect.CreateUserInput) *ConnectCreateUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateUser, input)
    return &ConnectCreateUserResult{Result: future}
}

func (a *ConnectStub) DeleteUser(ctx workflow.Context, input *connect.DeleteUserInput) (*connect.DeleteUserOutput, error) {
    var output connect.DeleteUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) DeleteUserAsync(ctx workflow.Context, input *connect.DeleteUserInput) *ConnectDeleteUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteUser, input)
    return &ConnectDeleteUserResult{Result: future}
}

func (a *ConnectStub) DescribeUser(ctx workflow.Context, input *connect.DescribeUserInput) (*connect.DescribeUserOutput, error) {
    var output connect.DescribeUserOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) DescribeUserAsync(ctx workflow.Context, input *connect.DescribeUserInput) *ConnectDescribeUserResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUser, input)
    return &ConnectDescribeUserResult{Result: future}
}

func (a *ConnectStub) DescribeUserHierarchyGroup(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) (*connect.DescribeUserHierarchyGroupOutput, error) {
    var output connect.DescribeUserHierarchyGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserHierarchyGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) DescribeUserHierarchyGroupAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyGroupInput) *ConnectDescribeUserHierarchyGroupResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUserHierarchyGroup, input)
    return &ConnectDescribeUserHierarchyGroupResult{Result: future}
}

func (a *ConnectStub) DescribeUserHierarchyStructure(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) (*connect.DescribeUserHierarchyStructureOutput, error) {
    var output connect.DescribeUserHierarchyStructureOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeUserHierarchyStructure, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) DescribeUserHierarchyStructureAsync(ctx workflow.Context, input *connect.DescribeUserHierarchyStructureInput) *ConnectDescribeUserHierarchyStructureResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeUserHierarchyStructure, input)
    return &ConnectDescribeUserHierarchyStructureResult{Result: future}
}

func (a *ConnectStub) GetContactAttributes(ctx workflow.Context, input *connect.GetContactAttributesInput) (*connect.GetContactAttributesOutput, error) {
    var output connect.GetContactAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetContactAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) GetContactAttributesAsync(ctx workflow.Context, input *connect.GetContactAttributesInput) *ConnectGetContactAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetContactAttributes, input)
    return &ConnectGetContactAttributesResult{Result: future}
}

func (a *ConnectStub) GetCurrentMetricData(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) (*connect.GetCurrentMetricDataOutput, error) {
    var output connect.GetCurrentMetricDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCurrentMetricData, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) GetCurrentMetricDataAsync(ctx workflow.Context, input *connect.GetCurrentMetricDataInput) *ConnectGetCurrentMetricDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetCurrentMetricData, input)
    return &ConnectGetCurrentMetricDataResult{Result: future}
}

func (a *ConnectStub) GetFederationToken(ctx workflow.Context, input *connect.GetFederationTokenInput) (*connect.GetFederationTokenOutput, error) {
    var output connect.GetFederationTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetFederationToken, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) GetFederationTokenAsync(ctx workflow.Context, input *connect.GetFederationTokenInput) *ConnectGetFederationTokenResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetFederationToken, input)
    return &ConnectGetFederationTokenResult{Result: future}
}

func (a *ConnectStub) GetMetricData(ctx workflow.Context, input *connect.GetMetricDataInput) (*connect.GetMetricDataOutput, error) {
    var output connect.GetMetricDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMetricData, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) GetMetricDataAsync(ctx workflow.Context, input *connect.GetMetricDataInput) *ConnectGetMetricDataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetMetricData, input)
    return &ConnectGetMetricDataResult{Result: future}
}

func (a *ConnectStub) ListContactFlows(ctx workflow.Context, input *connect.ListContactFlowsInput) (*connect.ListContactFlowsOutput, error) {
    var output connect.ListContactFlowsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListContactFlows, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListContactFlowsAsync(ctx workflow.Context, input *connect.ListContactFlowsInput) *ConnectListContactFlowsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListContactFlows, input)
    return &ConnectListContactFlowsResult{Result: future}
}

func (a *ConnectStub) ListHoursOfOperations(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) (*connect.ListHoursOfOperationsOutput, error) {
    var output connect.ListHoursOfOperationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHoursOfOperations, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListHoursOfOperationsAsync(ctx workflow.Context, input *connect.ListHoursOfOperationsInput) *ConnectListHoursOfOperationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListHoursOfOperations, input)
    return &ConnectListHoursOfOperationsResult{Result: future}
}

func (a *ConnectStub) ListPhoneNumbers(ctx workflow.Context, input *connect.ListPhoneNumbersInput) (*connect.ListPhoneNumbersOutput, error) {
    var output connect.ListPhoneNumbersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPhoneNumbers, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListPhoneNumbersAsync(ctx workflow.Context, input *connect.ListPhoneNumbersInput) *ConnectListPhoneNumbersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListPhoneNumbers, input)
    return &ConnectListPhoneNumbersResult{Result: future}
}

func (a *ConnectStub) ListQueues(ctx workflow.Context, input *connect.ListQueuesInput) (*connect.ListQueuesOutput, error) {
    var output connect.ListQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListQueuesAsync(ctx workflow.Context, input *connect.ListQueuesInput) *ConnectListQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input)
    return &ConnectListQueuesResult{Result: future}
}

func (a *ConnectStub) ListRoutingProfiles(ctx workflow.Context, input *connect.ListRoutingProfilesInput) (*connect.ListRoutingProfilesOutput, error) {
    var output connect.ListRoutingProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRoutingProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListRoutingProfilesAsync(ctx workflow.Context, input *connect.ListRoutingProfilesInput) *ConnectListRoutingProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRoutingProfiles, input)
    return &ConnectListRoutingProfilesResult{Result: future}
}

func (a *ConnectStub) ListSecurityProfiles(ctx workflow.Context, input *connect.ListSecurityProfilesInput) (*connect.ListSecurityProfilesOutput, error) {
    var output connect.ListSecurityProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListSecurityProfilesAsync(ctx workflow.Context, input *connect.ListSecurityProfilesInput) *ConnectListSecurityProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListSecurityProfiles, input)
    return &ConnectListSecurityProfilesResult{Result: future}
}

func (a *ConnectStub) ListTagsForResource(ctx workflow.Context, input *connect.ListTagsForResourceInput) (*connect.ListTagsForResourceOutput, error) {
    var output connect.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListTagsForResourceAsync(ctx workflow.Context, input *connect.ListTagsForResourceInput) *ConnectListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &ConnectListTagsForResourceResult{Result: future}
}

func (a *ConnectStub) ListUserHierarchyGroups(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) (*connect.ListUserHierarchyGroupsOutput, error) {
    var output connect.ListUserHierarchyGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUserHierarchyGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListUserHierarchyGroupsAsync(ctx workflow.Context, input *connect.ListUserHierarchyGroupsInput) *ConnectListUserHierarchyGroupsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUserHierarchyGroups, input)
    return &ConnectListUserHierarchyGroupsResult{Result: future}
}

func (a *ConnectStub) ListUsers(ctx workflow.Context, input *connect.ListUsersInput) (*connect.ListUsersOutput, error) {
    var output connect.ListUsersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ListUsersAsync(ctx workflow.Context, input *connect.ListUsersInput) *ConnectListUsersResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListUsers, input)
    return &ConnectListUsersResult{Result: future}
}

func (a *ConnectStub) ResumeContactRecording(ctx workflow.Context, input *connect.ResumeContactRecordingInput) (*connect.ResumeContactRecordingOutput, error) {
    var output connect.ResumeContactRecordingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ResumeContactRecording, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) ResumeContactRecordingAsync(ctx workflow.Context, input *connect.ResumeContactRecordingInput) *ConnectResumeContactRecordingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ResumeContactRecording, input)
    return &ConnectResumeContactRecordingResult{Result: future}
}

func (a *ConnectStub) StartChatContact(ctx workflow.Context, input *connect.StartChatContactInput) (*connect.StartChatContactOutput, error) {
    var output connect.StartChatContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartChatContact, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) StartChatContactAsync(ctx workflow.Context, input *connect.StartChatContactInput) *ConnectStartChatContactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartChatContact, input)
    return &ConnectStartChatContactResult{Result: future}
}

func (a *ConnectStub) StartContactRecording(ctx workflow.Context, input *connect.StartContactRecordingInput) (*connect.StartContactRecordingOutput, error) {
    var output connect.StartContactRecordingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartContactRecording, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) StartContactRecordingAsync(ctx workflow.Context, input *connect.StartContactRecordingInput) *ConnectStartContactRecordingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartContactRecording, input)
    return &ConnectStartContactRecordingResult{Result: future}
}

func (a *ConnectStub) StartOutboundVoiceContact(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) (*connect.StartOutboundVoiceContactOutput, error) {
    var output connect.StartOutboundVoiceContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartOutboundVoiceContact, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) StartOutboundVoiceContactAsync(ctx workflow.Context, input *connect.StartOutboundVoiceContactInput) *ConnectStartOutboundVoiceContactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartOutboundVoiceContact, input)
    return &ConnectStartOutboundVoiceContactResult{Result: future}
}

func (a *ConnectStub) StopContact(ctx workflow.Context, input *connect.StopContactInput) (*connect.StopContactOutput, error) {
    var output connect.StopContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopContact, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) StopContactAsync(ctx workflow.Context, input *connect.StopContactInput) *ConnectStopContactResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopContact, input)
    return &ConnectStopContactResult{Result: future}
}

func (a *ConnectStub) StopContactRecording(ctx workflow.Context, input *connect.StopContactRecordingInput) (*connect.StopContactRecordingOutput, error) {
    var output connect.StopContactRecordingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopContactRecording, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) StopContactRecordingAsync(ctx workflow.Context, input *connect.StopContactRecordingInput) *ConnectStopContactRecordingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopContactRecording, input)
    return &ConnectStopContactRecordingResult{Result: future}
}

func (a *ConnectStub) SuspendContactRecording(ctx workflow.Context, input *connect.SuspendContactRecordingInput) (*connect.SuspendContactRecordingOutput, error) {
    var output connect.SuspendContactRecordingOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SuspendContactRecording, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) SuspendContactRecordingAsync(ctx workflow.Context, input *connect.SuspendContactRecordingInput) *ConnectSuspendContactRecordingResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SuspendContactRecording, input)
    return &ConnectSuspendContactRecordingResult{Result: future}
}

func (a *ConnectStub) TagResource(ctx workflow.Context, input *connect.TagResourceInput) (*connect.TagResourceOutput, error) {
    var output connect.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) TagResourceAsync(ctx workflow.Context, input *connect.TagResourceInput) *ConnectTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &ConnectTagResourceResult{Result: future}
}

func (a *ConnectStub) UntagResource(ctx workflow.Context, input *connect.UntagResourceInput) (*connect.UntagResourceOutput, error) {
    var output connect.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UntagResourceAsync(ctx workflow.Context, input *connect.UntagResourceInput) *ConnectUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &ConnectUntagResourceResult{Result: future}
}

func (a *ConnectStub) UpdateContactAttributes(ctx workflow.Context, input *connect.UpdateContactAttributesInput) (*connect.UpdateContactAttributesOutput, error) {
    var output connect.UpdateContactAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateContactAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateContactAttributesAsync(ctx workflow.Context, input *connect.UpdateContactAttributesInput) *ConnectUpdateContactAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateContactAttributes, input)
    return &ConnectUpdateContactAttributesResult{Result: future}
}

func (a *ConnectStub) UpdateUserHierarchy(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) (*connect.UpdateUserHierarchyOutput, error) {
    var output connect.UpdateUserHierarchyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserHierarchy, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateUserHierarchyAsync(ctx workflow.Context, input *connect.UpdateUserHierarchyInput) *ConnectUpdateUserHierarchyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserHierarchy, input)
    return &ConnectUpdateUserHierarchyResult{Result: future}
}

func (a *ConnectStub) UpdateUserIdentityInfo(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) (*connect.UpdateUserIdentityInfoOutput, error) {
    var output connect.UpdateUserIdentityInfoOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserIdentityInfo, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateUserIdentityInfoAsync(ctx workflow.Context, input *connect.UpdateUserIdentityInfoInput) *ConnectUpdateUserIdentityInfoResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserIdentityInfo, input)
    return &ConnectUpdateUserIdentityInfoResult{Result: future}
}

func (a *ConnectStub) UpdateUserPhoneConfig(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) (*connect.UpdateUserPhoneConfigOutput, error) {
    var output connect.UpdateUserPhoneConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserPhoneConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateUserPhoneConfigAsync(ctx workflow.Context, input *connect.UpdateUserPhoneConfigInput) *ConnectUpdateUserPhoneConfigResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserPhoneConfig, input)
    return &ConnectUpdateUserPhoneConfigResult{Result: future}
}

func (a *ConnectStub) UpdateUserRoutingProfile(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) (*connect.UpdateUserRoutingProfileOutput, error) {
    var output connect.UpdateUserRoutingProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserRoutingProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateUserRoutingProfileAsync(ctx workflow.Context, input *connect.UpdateUserRoutingProfileInput) *ConnectUpdateUserRoutingProfileResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserRoutingProfile, input)
    return &ConnectUpdateUserRoutingProfileResult{Result: future}
}

func (a *ConnectStub) UpdateUserSecurityProfiles(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) (*connect.UpdateUserSecurityProfilesOutput, error) {
    var output connect.UpdateUserSecurityProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateUserSecurityProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *ConnectStub) UpdateUserSecurityProfilesAsync(ctx workflow.Context, input *connect.UpdateUserSecurityProfilesInput) *ConnectUpdateUserSecurityProfilesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateUserSecurityProfiles, input)
    return &ConnectUpdateUserSecurityProfilesResult{Result: future}
}