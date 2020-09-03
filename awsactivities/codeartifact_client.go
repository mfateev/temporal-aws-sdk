package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codeartifact"
	"go.temporal.io/sdk/workflow"
)

type CodeArtifactClient interface {
    AssociateExternalConnection(ctx workflow.Context, input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error)
    AssociateExternalConnectionAsync(ctx workflow.Context, input *codeartifact.AssociateExternalConnectionInput) *CodeartifactAssociateExternalConnectionResult

    CopyPackageVersions(ctx workflow.Context, input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error)
    CopyPackageVersionsAsync(ctx workflow.Context, input *codeartifact.CopyPackageVersionsInput) *CodeartifactCopyPackageVersionsResult

    CreateDomain(ctx workflow.Context, input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error)
    CreateDomainAsync(ctx workflow.Context, input *codeartifact.CreateDomainInput) *CodeartifactCreateDomainResult

    CreateRepository(ctx workflow.Context, input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error)
    CreateRepositoryAsync(ctx workflow.Context, input *codeartifact.CreateRepositoryInput) *CodeartifactCreateRepositoryResult

    DeleteDomain(ctx workflow.Context, input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error)
    DeleteDomainAsync(ctx workflow.Context, input *codeartifact.DeleteDomainInput) *CodeartifactDeleteDomainResult

    DeleteDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error)
    DeleteDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) *CodeartifactDeleteDomainPermissionsPolicyResult

    DeletePackageVersions(ctx workflow.Context, input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error)
    DeletePackageVersionsAsync(ctx workflow.Context, input *codeartifact.DeletePackageVersionsInput) *CodeartifactDeletePackageVersionsResult

    DeleteRepository(ctx workflow.Context, input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error)
    DeleteRepositoryAsync(ctx workflow.Context, input *codeartifact.DeleteRepositoryInput) *CodeartifactDeleteRepositoryResult

    DeleteRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error)
    DeleteRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) *CodeartifactDeleteRepositoryPermissionsPolicyResult

    DescribeDomain(ctx workflow.Context, input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error)
    DescribeDomainAsync(ctx workflow.Context, input *codeartifact.DescribeDomainInput) *CodeartifactDescribeDomainResult

    DescribePackageVersion(ctx workflow.Context, input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error)
    DescribePackageVersionAsync(ctx workflow.Context, input *codeartifact.DescribePackageVersionInput) *CodeartifactDescribePackageVersionResult

    DescribeRepository(ctx workflow.Context, input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error)
    DescribeRepositoryAsync(ctx workflow.Context, input *codeartifact.DescribeRepositoryInput) *CodeartifactDescribeRepositoryResult

    DisassociateExternalConnection(ctx workflow.Context, input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error)
    DisassociateExternalConnectionAsync(ctx workflow.Context, input *codeartifact.DisassociateExternalConnectionInput) *CodeartifactDisassociateExternalConnectionResult

    DisposePackageVersions(ctx workflow.Context, input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error)
    DisposePackageVersionsAsync(ctx workflow.Context, input *codeartifact.DisposePackageVersionsInput) *CodeartifactDisposePackageVersionsResult

    GetAuthorizationToken(ctx workflow.Context, input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error)
    GetAuthorizationTokenAsync(ctx workflow.Context, input *codeartifact.GetAuthorizationTokenInput) *CodeartifactGetAuthorizationTokenResult

    GetDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error)
    GetDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.GetDomainPermissionsPolicyInput) *CodeartifactGetDomainPermissionsPolicyResult

    GetPackageVersionAsset(ctx workflow.Context, input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error)
    GetPackageVersionAssetAsync(ctx workflow.Context, input *codeartifact.GetPackageVersionAssetInput) *CodeartifactGetPackageVersionAssetResult

    GetPackageVersionReadme(ctx workflow.Context, input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error)
    GetPackageVersionReadmeAsync(ctx workflow.Context, input *codeartifact.GetPackageVersionReadmeInput) *CodeartifactGetPackageVersionReadmeResult

    GetRepositoryEndpoint(ctx workflow.Context, input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error)
    GetRepositoryEndpointAsync(ctx workflow.Context, input *codeartifact.GetRepositoryEndpointInput) *CodeartifactGetRepositoryEndpointResult

    GetRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error)
    GetRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) *CodeartifactGetRepositoryPermissionsPolicyResult

    ListDomains(ctx workflow.Context, input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error)
    ListDomainsAsync(ctx workflow.Context, input *codeartifact.ListDomainsInput) *CodeartifactListDomainsResult

    ListPackageVersionAssets(ctx workflow.Context, input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error)
    ListPackageVersionAssetsAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionAssetsInput) *CodeartifactListPackageVersionAssetsResult

    ListPackageVersionDependencies(ctx workflow.Context, input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error)
    ListPackageVersionDependenciesAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionDependenciesInput) *CodeartifactListPackageVersionDependenciesResult

    ListPackageVersions(ctx workflow.Context, input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error)
    ListPackageVersionsAsync(ctx workflow.Context, input *codeartifact.ListPackageVersionsInput) *CodeartifactListPackageVersionsResult

    ListPackages(ctx workflow.Context, input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error)
    ListPackagesAsync(ctx workflow.Context, input *codeartifact.ListPackagesInput) *CodeartifactListPackagesResult

    ListRepositories(ctx workflow.Context, input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error)
    ListRepositoriesAsync(ctx workflow.Context, input *codeartifact.ListRepositoriesInput) *CodeartifactListRepositoriesResult

    ListRepositoriesInDomain(ctx workflow.Context, input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error)
    ListRepositoriesInDomainAsync(ctx workflow.Context, input *codeartifact.ListRepositoriesInDomainInput) *CodeartifactListRepositoriesInDomainResult

    PutDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error)
    PutDomainPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.PutDomainPermissionsPolicyInput) *CodeartifactPutDomainPermissionsPolicyResult

    PutRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error)
    PutRepositoryPermissionsPolicyAsync(ctx workflow.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) *CodeartifactPutRepositoryPermissionsPolicyResult

    UpdatePackageVersionsStatus(ctx workflow.Context, input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error)
    UpdatePackageVersionsStatusAsync(ctx workflow.Context, input *codeartifact.UpdatePackageVersionsStatusInput) *CodeartifactUpdatePackageVersionsStatusResult

    UpdateRepository(ctx workflow.Context, input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error)
    UpdateRepositoryAsync(ctx workflow.Context, input *codeartifact.UpdateRepositoryInput) *CodeartifactUpdateRepositoryResult
}
type CodeartifactAssociateExternalConnectionResult struct {
	Result workflow.Future
}

func (r *CodeartifactAssociateExternalConnectionResult) Get(ctx workflow.Context) (*codeartifact.AssociateExternalConnectionOutput, error) {
    var output codeartifact.AssociateExternalConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactCopyPackageVersionsResult struct {
	Result workflow.Future
}

func (r *CodeartifactCopyPackageVersionsResult) Get(ctx workflow.Context) (*codeartifact.CopyPackageVersionsOutput, error) {
    var output codeartifact.CopyPackageVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactCreateDomainResult struct {
	Result workflow.Future
}

func (r *CodeartifactCreateDomainResult) Get(ctx workflow.Context) (*codeartifact.CreateDomainOutput, error) {
    var output codeartifact.CreateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactCreateRepositoryResult struct {
	Result workflow.Future
}

func (r *CodeartifactCreateRepositoryResult) Get(ctx workflow.Context) (*codeartifact.CreateRepositoryOutput, error) {
    var output codeartifact.CreateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDeleteDomainResult struct {
	Result workflow.Future
}

func (r *CodeartifactDeleteDomainResult) Get(ctx workflow.Context) (*codeartifact.DeleteDomainOutput, error) {
    var output codeartifact.DeleteDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDeleteDomainPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactDeleteDomainPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error) {
    var output codeartifact.DeleteDomainPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDeletePackageVersionsResult struct {
	Result workflow.Future
}

func (r *CodeartifactDeletePackageVersionsResult) Get(ctx workflow.Context) (*codeartifact.DeletePackageVersionsOutput, error) {
    var output codeartifact.DeletePackageVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDeleteRepositoryResult struct {
	Result workflow.Future
}

func (r *CodeartifactDeleteRepositoryResult) Get(ctx workflow.Context) (*codeartifact.DeleteRepositoryOutput, error) {
    var output codeartifact.DeleteRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDeleteRepositoryPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactDeleteRepositoryPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.DeleteRepositoryPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDescribeDomainResult struct {
	Result workflow.Future
}

func (r *CodeartifactDescribeDomainResult) Get(ctx workflow.Context) (*codeartifact.DescribeDomainOutput, error) {
    var output codeartifact.DescribeDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDescribePackageVersionResult struct {
	Result workflow.Future
}

func (r *CodeartifactDescribePackageVersionResult) Get(ctx workflow.Context) (*codeartifact.DescribePackageVersionOutput, error) {
    var output codeartifact.DescribePackageVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDescribeRepositoryResult struct {
	Result workflow.Future
}

func (r *CodeartifactDescribeRepositoryResult) Get(ctx workflow.Context) (*codeartifact.DescribeRepositoryOutput, error) {
    var output codeartifact.DescribeRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDisassociateExternalConnectionResult struct {
	Result workflow.Future
}

func (r *CodeartifactDisassociateExternalConnectionResult) Get(ctx workflow.Context) (*codeartifact.DisassociateExternalConnectionOutput, error) {
    var output codeartifact.DisassociateExternalConnectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactDisposePackageVersionsResult struct {
	Result workflow.Future
}

func (r *CodeartifactDisposePackageVersionsResult) Get(ctx workflow.Context) (*codeartifact.DisposePackageVersionsOutput, error) {
    var output codeartifact.DisposePackageVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetAuthorizationTokenResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetAuthorizationTokenResult) Get(ctx workflow.Context) (*codeartifact.GetAuthorizationTokenOutput, error) {
    var output codeartifact.GetAuthorizationTokenOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetDomainPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetDomainPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
    var output codeartifact.GetDomainPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetPackageVersionAssetResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetPackageVersionAssetResult) Get(ctx workflow.Context) (*codeartifact.GetPackageVersionAssetOutput, error) {
    var output codeartifact.GetPackageVersionAssetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetPackageVersionReadmeResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetPackageVersionReadmeResult) Get(ctx workflow.Context) (*codeartifact.GetPackageVersionReadmeOutput, error) {
    var output codeartifact.GetPackageVersionReadmeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetRepositoryEndpointResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetRepositoryEndpointResult) Get(ctx workflow.Context) (*codeartifact.GetRepositoryEndpointOutput, error) {
    var output codeartifact.GetRepositoryEndpointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactGetRepositoryPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactGetRepositoryPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.GetRepositoryPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListDomainsResult struct {
	Result workflow.Future
}

func (r *CodeartifactListDomainsResult) Get(ctx workflow.Context) (*codeartifact.ListDomainsOutput, error) {
    var output codeartifact.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListPackageVersionAssetsResult struct {
	Result workflow.Future
}

func (r *CodeartifactListPackageVersionAssetsResult) Get(ctx workflow.Context) (*codeartifact.ListPackageVersionAssetsOutput, error) {
    var output codeartifact.ListPackageVersionAssetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListPackageVersionDependenciesResult struct {
	Result workflow.Future
}

func (r *CodeartifactListPackageVersionDependenciesResult) Get(ctx workflow.Context) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
    var output codeartifact.ListPackageVersionDependenciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListPackageVersionsResult struct {
	Result workflow.Future
}

func (r *CodeartifactListPackageVersionsResult) Get(ctx workflow.Context) (*codeartifact.ListPackageVersionsOutput, error) {
    var output codeartifact.ListPackageVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListPackagesResult struct {
	Result workflow.Future
}

func (r *CodeartifactListPackagesResult) Get(ctx workflow.Context) (*codeartifact.ListPackagesOutput, error) {
    var output codeartifact.ListPackagesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListRepositoriesResult struct {
	Result workflow.Future
}

func (r *CodeartifactListRepositoriesResult) Get(ctx workflow.Context) (*codeartifact.ListRepositoriesOutput, error) {
    var output codeartifact.ListRepositoriesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactListRepositoriesInDomainResult struct {
	Result workflow.Future
}

func (r *CodeartifactListRepositoriesInDomainResult) Get(ctx workflow.Context) (*codeartifact.ListRepositoriesInDomainOutput, error) {
    var output codeartifact.ListRepositoriesInDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactPutDomainPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactPutDomainPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.PutDomainPermissionsPolicyOutput, error) {
    var output codeartifact.PutDomainPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactPutRepositoryPermissionsPolicyResult struct {
	Result workflow.Future
}

func (r *CodeartifactPutRepositoryPermissionsPolicyResult) Get(ctx workflow.Context) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.PutRepositoryPermissionsPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactUpdatePackageVersionsStatusResult struct {
	Result workflow.Future
}

func (r *CodeartifactUpdatePackageVersionsStatusResult) Get(ctx workflow.Context) (*codeartifact.UpdatePackageVersionsStatusOutput, error) {
    var output codeartifact.UpdatePackageVersionsStatusOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodeartifactUpdateRepositoryResult struct {
	Result workflow.Future
}

func (r *CodeartifactUpdateRepositoryResult) Get(ctx workflow.Context) (*codeartifact.UpdateRepositoryOutput, error) {
    var output codeartifact.UpdateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CodeArtifactStub struct {
    activities CodeArtifactClient
}

func NewCodeArtifactStub() CodeArtifactClient {
    return &CodeArtifactStub{}
}

func (a *CodeArtifactStub) AssociateExternalConnection(ctx workflow.Context, input *codeartifact.AssociateExternalConnectionInput) (*codeartifact.AssociateExternalConnectionOutput, error) {
    var output codeartifact.AssociateExternalConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateExternalConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) CopyPackageVersions(ctx workflow.Context, input *codeartifact.CopyPackageVersionsInput) (*codeartifact.CopyPackageVersionsOutput, error) {
    var output codeartifact.CopyPackageVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CopyPackageVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) CreateDomain(ctx workflow.Context, input *codeartifact.CreateDomainInput) (*codeartifact.CreateDomainOutput, error) {
    var output codeartifact.CreateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) CreateRepository(ctx workflow.Context, input *codeartifact.CreateRepositoryInput) (*codeartifact.CreateRepositoryOutput, error) {
    var output codeartifact.CreateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DeleteDomain(ctx workflow.Context, input *codeartifact.DeleteDomainInput) (*codeartifact.DeleteDomainOutput, error) {
    var output codeartifact.DeleteDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DeleteDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteDomainPermissionsPolicyInput) (*codeartifact.DeleteDomainPermissionsPolicyOutput, error) {
    var output codeartifact.DeleteDomainPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DeletePackageVersions(ctx workflow.Context, input *codeartifact.DeletePackageVersionsInput) (*codeartifact.DeletePackageVersionsOutput, error) {
    var output codeartifact.DeletePackageVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeletePackageVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DeleteRepository(ctx workflow.Context, input *codeartifact.DeleteRepositoryInput) (*codeartifact.DeleteRepositoryOutput, error) {
    var output codeartifact.DeleteRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DeleteRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.DeleteRepositoryPermissionsPolicyInput) (*codeartifact.DeleteRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.DeleteRepositoryPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRepositoryPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DescribeDomain(ctx workflow.Context, input *codeartifact.DescribeDomainInput) (*codeartifact.DescribeDomainOutput, error) {
    var output codeartifact.DescribeDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DescribePackageVersion(ctx workflow.Context, input *codeartifact.DescribePackageVersionInput) (*codeartifact.DescribePackageVersionOutput, error) {
    var output codeartifact.DescribePackageVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribePackageVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DescribeRepository(ctx workflow.Context, input *codeartifact.DescribeRepositoryInput) (*codeartifact.DescribeRepositoryOutput, error) {
    var output codeartifact.DescribeRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DisassociateExternalConnection(ctx workflow.Context, input *codeartifact.DisassociateExternalConnectionInput) (*codeartifact.DisassociateExternalConnectionOutput, error) {
    var output codeartifact.DisassociateExternalConnectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateExternalConnection, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) DisposePackageVersions(ctx workflow.Context, input *codeartifact.DisposePackageVersionsInput) (*codeartifact.DisposePackageVersionsOutput, error) {
    var output codeartifact.DisposePackageVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisposePackageVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetAuthorizationToken(ctx workflow.Context, input *codeartifact.GetAuthorizationTokenInput) (*codeartifact.GetAuthorizationTokenOutput, error) {
    var output codeartifact.GetAuthorizationTokenOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAuthorizationToken, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetDomainPermissionsPolicyInput) (*codeartifact.GetDomainPermissionsPolicyOutput, error) {
    var output codeartifact.GetDomainPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetPackageVersionAsset(ctx workflow.Context, input *codeartifact.GetPackageVersionAssetInput) (*codeartifact.GetPackageVersionAssetOutput, error) {
    var output codeartifact.GetPackageVersionAssetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPackageVersionAsset, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetPackageVersionReadme(ctx workflow.Context, input *codeartifact.GetPackageVersionReadmeInput) (*codeartifact.GetPackageVersionReadmeOutput, error) {
    var output codeartifact.GetPackageVersionReadmeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetPackageVersionReadme, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetRepositoryEndpoint(ctx workflow.Context, input *codeartifact.GetRepositoryEndpointInput) (*codeartifact.GetRepositoryEndpointOutput, error) {
    var output codeartifact.GetRepositoryEndpointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRepositoryEndpoint, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) GetRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.GetRepositoryPermissionsPolicyInput) (*codeartifact.GetRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.GetRepositoryPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRepositoryPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListDomains(ctx workflow.Context, input *codeartifact.ListDomainsInput) (*codeartifact.ListDomainsOutput, error) {
    var output codeartifact.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListPackageVersionAssets(ctx workflow.Context, input *codeartifact.ListPackageVersionAssetsInput) (*codeartifact.ListPackageVersionAssetsOutput, error) {
    var output codeartifact.ListPackageVersionAssetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackageVersionAssets, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListPackageVersionDependencies(ctx workflow.Context, input *codeartifact.ListPackageVersionDependenciesInput) (*codeartifact.ListPackageVersionDependenciesOutput, error) {
    var output codeartifact.ListPackageVersionDependenciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackageVersionDependencies, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListPackageVersions(ctx workflow.Context, input *codeartifact.ListPackageVersionsInput) (*codeartifact.ListPackageVersionsOutput, error) {
    var output codeartifact.ListPackageVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackageVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListPackages(ctx workflow.Context, input *codeartifact.ListPackagesInput) (*codeartifact.ListPackagesOutput, error) {
    var output codeartifact.ListPackagesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListPackages, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListRepositories(ctx workflow.Context, input *codeartifact.ListRepositoriesInput) (*codeartifact.ListRepositoriesOutput, error) {
    var output codeartifact.ListRepositoriesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRepositories, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) ListRepositoriesInDomain(ctx workflow.Context, input *codeartifact.ListRepositoriesInDomainInput) (*codeartifact.ListRepositoriesInDomainOutput, error) {
    var output codeartifact.ListRepositoriesInDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRepositoriesInDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) PutDomainPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutDomainPermissionsPolicyInput) (*codeartifact.PutDomainPermissionsPolicyOutput, error) {
    var output codeartifact.PutDomainPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutDomainPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) PutRepositoryPermissionsPolicy(ctx workflow.Context, input *codeartifact.PutRepositoryPermissionsPolicyInput) (*codeartifact.PutRepositoryPermissionsPolicyOutput, error) {
    var output codeartifact.PutRepositoryPermissionsPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRepositoryPermissionsPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) UpdatePackageVersionsStatus(ctx workflow.Context, input *codeartifact.UpdatePackageVersionsStatusInput) (*codeartifact.UpdatePackageVersionsStatusOutput, error) {
    var output codeartifact.UpdatePackageVersionsStatusOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdatePackageVersionsStatus, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeArtifactStub) UpdateRepository(ctx workflow.Context, input *codeartifact.UpdateRepositoryInput) (*codeartifact.UpdateRepositoryOutput, error) {
    var output codeartifact.UpdateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRepository, input).Get(ctx, &output)
    return &output, err
}
