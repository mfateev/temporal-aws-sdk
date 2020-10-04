// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"github.com/aws/aws-sdk-go/service/comprehend/comprehendiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type ComprehendActivities struct {
	client comprehendiface.ComprehendAPI

	sessionFactory SessionFactory
}

func NewComprehendActivities(sess *session.Session, config ...*aws.Config) *ComprehendActivities {
	client := comprehend.New(sess, config...)
	return &ComprehendActivities{client: client}
}

func NewComprehendActivitiesWithSessionFactory(sessionFactory SessionFactory) *ComprehendActivities {
	return &ComprehendActivities{sessionFactory: sessionFactory}
}

func (a *ComprehendActivities) getClient(ctx context.Context) (comprehendiface.ComprehendAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return comprehend.New(sess), nil
}

func (a *ComprehendActivities) BatchDetectDominantLanguage(ctx context.Context, input *comprehend.BatchDetectDominantLanguageInput) (*comprehend.BatchDetectDominantLanguageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDetectDominantLanguageWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectEntities(ctx context.Context, input *comprehend.BatchDetectEntitiesInput) (*comprehend.BatchDetectEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDetectEntitiesWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectKeyPhrases(ctx context.Context, input *comprehend.BatchDetectKeyPhrasesInput) (*comprehend.BatchDetectKeyPhrasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDetectKeyPhrasesWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectSentiment(ctx context.Context, input *comprehend.BatchDetectSentimentInput) (*comprehend.BatchDetectSentimentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDetectSentimentWithContext(ctx, input)
}

func (a *ComprehendActivities) BatchDetectSyntax(ctx context.Context, input *comprehend.BatchDetectSyntaxInput) (*comprehend.BatchDetectSyntaxOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.BatchDetectSyntaxWithContext(ctx, input)
}

func (a *ComprehendActivities) ClassifyDocument(ctx context.Context, input *comprehend.ClassifyDocumentInput) (*comprehend.ClassifyDocumentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ClassifyDocumentWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateDocumentClassifier(ctx context.Context, input *comprehend.CreateDocumentClassifierInput) (*comprehend.CreateDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateEndpoint(ctx context.Context, input *comprehend.CreateEndpointInput) (*comprehend.CreateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) CreateEntityRecognizer(ctx context.Context, input *comprehend.CreateEntityRecognizerInput) (*comprehend.CreateEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteDocumentClassifier(ctx context.Context, input *comprehend.DeleteDocumentClassifierInput) (*comprehend.DeleteDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteEndpoint(ctx context.Context, input *comprehend.DeleteEndpointInput) (*comprehend.DeleteEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) DeleteEntityRecognizer(ctx context.Context, input *comprehend.DeleteEntityRecognizerInput) (*comprehend.DeleteEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDocumentClassificationJob(ctx context.Context, input *comprehend.DescribeDocumentClassificationJobInput) (*comprehend.DescribeDocumentClassificationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDocumentClassificationJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDocumentClassifier(ctx context.Context, input *comprehend.DescribeDocumentClassifierInput) (*comprehend.DescribeDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeDominantLanguageDetectionJob(ctx context.Context, input *comprehend.DescribeDominantLanguageDetectionJobInput) (*comprehend.DescribeDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEndpoint(ctx context.Context, input *comprehend.DescribeEndpointInput) (*comprehend.DescribeEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEndpointWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEntitiesDetectionJob(ctx context.Context, input *comprehend.DescribeEntitiesDetectionJobInput) (*comprehend.DescribeEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeEntityRecognizer(ctx context.Context, input *comprehend.DescribeEntityRecognizerInput) (*comprehend.DescribeEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.DescribeKeyPhrasesDetectionJobInput) (*comprehend.DescribeKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeSentimentDetectionJob(ctx context.Context, input *comprehend.DescribeSentimentDetectionJobInput) (*comprehend.DescribeSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DescribeTopicsDetectionJob(ctx context.Context, input *comprehend.DescribeTopicsDetectionJobInput) (*comprehend.DescribeTopicsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTopicsDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectDominantLanguage(ctx context.Context, input *comprehend.DetectDominantLanguageInput) (*comprehend.DetectDominantLanguageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectDominantLanguageWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectEntities(ctx context.Context, input *comprehend.DetectEntitiesInput) (*comprehend.DetectEntitiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectEntitiesWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectKeyPhrases(ctx context.Context, input *comprehend.DetectKeyPhrasesInput) (*comprehend.DetectKeyPhrasesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectKeyPhrasesWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectSentiment(ctx context.Context, input *comprehend.DetectSentimentInput) (*comprehend.DetectSentimentOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectSentimentWithContext(ctx, input)
}

func (a *ComprehendActivities) DetectSyntax(ctx context.Context, input *comprehend.DetectSyntaxInput) (*comprehend.DetectSyntaxOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DetectSyntaxWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDocumentClassificationJobs(ctx context.Context, input *comprehend.ListDocumentClassificationJobsInput) (*comprehend.ListDocumentClassificationJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDocumentClassificationJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDocumentClassifiers(ctx context.Context, input *comprehend.ListDocumentClassifiersInput) (*comprehend.ListDocumentClassifiersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDocumentClassifiersWithContext(ctx, input)
}

func (a *ComprehendActivities) ListDominantLanguageDetectionJobs(ctx context.Context, input *comprehend.ListDominantLanguageDetectionJobsInput) (*comprehend.ListDominantLanguageDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListDominantLanguageDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEndpoints(ctx context.Context, input *comprehend.ListEndpointsInput) (*comprehend.ListEndpointsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEndpointsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEntitiesDetectionJobs(ctx context.Context, input *comprehend.ListEntitiesDetectionJobsInput) (*comprehend.ListEntitiesDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEntitiesDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListEntityRecognizers(ctx context.Context, input *comprehend.ListEntityRecognizersInput) (*comprehend.ListEntityRecognizersOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListEntityRecognizersWithContext(ctx, input)
}

func (a *ComprehendActivities) ListKeyPhrasesDetectionJobs(ctx context.Context, input *comprehend.ListKeyPhrasesDetectionJobsInput) (*comprehend.ListKeyPhrasesDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListKeyPhrasesDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListSentimentDetectionJobs(ctx context.Context, input *comprehend.ListSentimentDetectionJobsInput) (*comprehend.ListSentimentDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListSentimentDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) ListTagsForResource(ctx context.Context, input *comprehend.ListTagsForResourceInput) (*comprehend.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) ListTopicsDetectionJobs(ctx context.Context, input *comprehend.ListTopicsDetectionJobsInput) (*comprehend.ListTopicsDetectionJobsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTopicsDetectionJobsWithContext(ctx, input)
}

func (a *ComprehendActivities) StartDocumentClassificationJob(ctx context.Context, input *comprehend.StartDocumentClassificationJobInput) (*comprehend.StartDocumentClassificationJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDocumentClassificationJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StartDominantLanguageDetectionJobInput) (*comprehend.StartDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartEntitiesDetectionJob(ctx context.Context, input *comprehend.StartEntitiesDetectionJobInput) (*comprehend.StartEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StartKeyPhrasesDetectionJobInput) (*comprehend.StartKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartSentimentDetectionJob(ctx context.Context, input *comprehend.StartSentimentDetectionJobInput) (*comprehend.StartSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StartTopicsDetectionJob(ctx context.Context, input *comprehend.StartTopicsDetectionJobInput) (*comprehend.StartTopicsDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartTopicsDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopDominantLanguageDetectionJob(ctx context.Context, input *comprehend.StopDominantLanguageDetectionJobInput) (*comprehend.StopDominantLanguageDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopDominantLanguageDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopEntitiesDetectionJob(ctx context.Context, input *comprehend.StopEntitiesDetectionJobInput) (*comprehend.StopEntitiesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopEntitiesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopKeyPhrasesDetectionJob(ctx context.Context, input *comprehend.StopKeyPhrasesDetectionJobInput) (*comprehend.StopKeyPhrasesDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopKeyPhrasesDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopSentimentDetectionJob(ctx context.Context, input *comprehend.StopSentimentDetectionJobInput) (*comprehend.StopSentimentDetectionJobOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopSentimentDetectionJobWithContext(ctx, input)
}

func (a *ComprehendActivities) StopTrainingDocumentClassifier(ctx context.Context, input *comprehend.StopTrainingDocumentClassifierInput) (*comprehend.StopTrainingDocumentClassifierOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopTrainingDocumentClassifierWithContext(ctx, input)
}

func (a *ComprehendActivities) StopTrainingEntityRecognizer(ctx context.Context, input *comprehend.StopTrainingEntityRecognizerInput) (*comprehend.StopTrainingEntityRecognizerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopTrainingEntityRecognizerWithContext(ctx, input)
}

func (a *ComprehendActivities) TagResource(ctx context.Context, input *comprehend.TagResourceInput) (*comprehend.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) UntagResource(ctx context.Context, input *comprehend.UntagResourceInput) (*comprehend.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *ComprehendActivities) UpdateEndpoint(ctx context.Context, input *comprehend.UpdateEndpointInput) (*comprehend.UpdateEndpointOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateEndpointWithContext(ctx, input)
}
