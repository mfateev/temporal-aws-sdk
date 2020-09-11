package awsclients

import (
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"go.temporal.io/sdk/workflow"
)

type RedshiftDataAPIServiceClient interface {
       CancelStatement(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error)
       CancelStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) *RedshiftdataapiserviceCancelStatementResult

       DescribeStatement(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error)
       DescribeStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) *RedshiftdataapiserviceDescribeStatementResult

       DescribeTable(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error)
       DescribeTableAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) *RedshiftdataapiserviceDescribeTableResult

       ExecuteStatement(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error)
       ExecuteStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) *RedshiftdataapiserviceExecuteStatementResult

       GetStatementResult(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error)
       GetStatementResultAsync(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) *RedshiftdataapiserviceGetStatementResultResult

       ListDatabases(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error)
       ListDatabasesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) *RedshiftdataapiserviceListDatabasesResult

       ListSchemas(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error)
       ListSchemasAsync(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) *RedshiftdataapiserviceListSchemasResult

       ListStatements(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error)
       ListStatementsAsync(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) *RedshiftdataapiserviceListStatementsResult

       ListTables(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error)
       ListTablesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) *RedshiftdataapiserviceListTablesResult
}

type RedshiftDataAPIServiceStub struct {
}

func NewRedshiftDataAPIServiceStub() RedshiftDataAPIServiceClient {
    return &RedshiftDataAPIServiceStub{}
}

type RedshiftdataapiserviceCancelStatementResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceCancelStatementResult) Get(ctx workflow.Context) (*redshiftdataapiservice.CancelStatementOutput, error) {
    var output redshiftdataapiservice.CancelStatementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceDescribeStatementResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceDescribeStatementResult) Get(ctx workflow.Context) (*redshiftdataapiservice.DescribeStatementOutput, error) {
    var output redshiftdataapiservice.DescribeStatementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceDescribeTableResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceDescribeTableResult) Get(ctx workflow.Context) (*redshiftdataapiservice.DescribeTableOutput, error) {
    var output redshiftdataapiservice.DescribeTableOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceExecuteStatementResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceExecuteStatementResult) Get(ctx workflow.Context) (*redshiftdataapiservice.ExecuteStatementOutput, error) {
    var output redshiftdataapiservice.ExecuteStatementOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceGetStatementResultResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceGetStatementResultResult) Get(ctx workflow.Context) (*redshiftdataapiservice.GetStatementResultOutput, error) {
    var output redshiftdataapiservice.GetStatementResultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceListDatabasesResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceListDatabasesResult) Get(ctx workflow.Context) (*redshiftdataapiservice.ListDatabasesOutput, error) {
    var output redshiftdataapiservice.ListDatabasesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceListSchemasResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceListSchemasResult) Get(ctx workflow.Context) (*redshiftdataapiservice.ListSchemasOutput, error) {
    var output redshiftdataapiservice.ListSchemasOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceListStatementsResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceListStatementsResult) Get(ctx workflow.Context) (*redshiftdataapiservice.ListStatementsOutput, error) {
    var output redshiftdataapiservice.ListStatementsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}
type RedshiftdataapiserviceListTablesResult struct {
	Result workflow.Future
}

func (r *RedshiftdataapiserviceListTablesResult) Get(ctx workflow.Context) (*redshiftdataapiservice.ListTablesOutput, error) {
    var output redshiftdataapiservice.ListTablesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) CancelStatement(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) (*redshiftdataapiservice.CancelStatementOutput, error) {
    var output redshiftdataapiservice.CancelStatementOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.CancelStatement", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) CancelStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.CancelStatementInput) *RedshiftdataapiserviceCancelStatementResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.CancelStatement", input)
    return &RedshiftdataapiserviceCancelStatementResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) DescribeStatement(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) (*redshiftdataapiservice.DescribeStatementOutput, error) {
    var output redshiftdataapiservice.DescribeStatementOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.DescribeStatement", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) DescribeStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeStatementInput) *RedshiftdataapiserviceDescribeStatementResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.DescribeStatement", input)
    return &RedshiftdataapiserviceDescribeStatementResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) DescribeTable(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) (*redshiftdataapiservice.DescribeTableOutput, error) {
    var output redshiftdataapiservice.DescribeTableOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.DescribeTable", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) DescribeTableAsync(ctx workflow.Context, input *redshiftdataapiservice.DescribeTableInput) *RedshiftdataapiserviceDescribeTableResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.DescribeTable", input)
    return &RedshiftdataapiserviceDescribeTableResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) ExecuteStatement(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) (*redshiftdataapiservice.ExecuteStatementOutput, error) {
    var output redshiftdataapiservice.ExecuteStatementOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ExecuteStatement", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) ExecuteStatementAsync(ctx workflow.Context, input *redshiftdataapiservice.ExecuteStatementInput) *RedshiftdataapiserviceExecuteStatementResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ExecuteStatement", input)
    return &RedshiftdataapiserviceExecuteStatementResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) GetStatementResult(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) (*redshiftdataapiservice.GetStatementResultOutput, error) {
    var output redshiftdataapiservice.GetStatementResultOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.GetStatementResult", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) GetStatementResultAsync(ctx workflow.Context, input *redshiftdataapiservice.GetStatementResultInput) *RedshiftdataapiserviceGetStatementResultResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.GetStatementResult", input)
    return &RedshiftdataapiserviceGetStatementResultResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) ListDatabases(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) (*redshiftdataapiservice.ListDatabasesOutput, error) {
    var output redshiftdataapiservice.ListDatabasesOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListDatabases", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) ListDatabasesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListDatabasesInput) *RedshiftdataapiserviceListDatabasesResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListDatabases", input)
    return &RedshiftdataapiserviceListDatabasesResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) ListSchemas(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) (*redshiftdataapiservice.ListSchemasOutput, error) {
    var output redshiftdataapiservice.ListSchemasOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListSchemas", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) ListSchemasAsync(ctx workflow.Context, input *redshiftdataapiservice.ListSchemasInput) *RedshiftdataapiserviceListSchemasResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListSchemas", input)
    return &RedshiftdataapiserviceListSchemasResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) ListStatements(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) (*redshiftdataapiservice.ListStatementsOutput, error) {
    var output redshiftdataapiservice.ListStatementsOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListStatements", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) ListStatementsAsync(ctx workflow.Context, input *redshiftdataapiservice.ListStatementsInput) *RedshiftdataapiserviceListStatementsResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListStatements", input)
    return &RedshiftdataapiserviceListStatementsResult{Result: future}
}

func (a *RedshiftDataAPIServiceStub) ListTables(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) (*redshiftdataapiservice.ListTablesOutput, error) {
    var output redshiftdataapiservice.ListTablesOutput
    err := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListTables", input).Get(ctx, &output)
    return &output, err
}

func (a *RedshiftDataAPIServiceStub) ListTablesAsync(ctx workflow.Context, input *redshiftdataapiservice.ListTablesInput) *RedshiftdataapiserviceListTablesResult {
    future := workflow.ExecuteActivity(ctx, "RedshiftDataAPIService.ListTables", input)
    return &RedshiftdataapiserviceListTablesResult{Result: future}
}