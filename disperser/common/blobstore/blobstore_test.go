package blobstore_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/zero-gravity-labs/zerog-data-avail/common/aws"
	"github.com/zero-gravity-labs/zerog-data-avail/common/aws/dynamodb"
	test_utils "github.com/zero-gravity-labs/zerog-data-avail/common/aws/dynamodb/utils"

	"github.com/ory/dockertest/v3"
	cmock "github.com/zero-gravity-labs/zerog-data-avail/common/mock"
	"github.com/zero-gravity-labs/zerog-data-avail/core"
	"github.com/zero-gravity-labs/zerog-data-avail/disperser/common/blobstore"
	"github.com/zero-gravity-labs/zerog-data-avail/inabox/deploy"
)

var (
	logger         = &cmock.Logger{}
	securityParams = []*core.SecurityParam{{
		QuorumID:           1,
		AdversaryThreshold: 80,
		QuorumRate:         32000,
	},
	}
	blob = &core.Blob{
		RequestHeader: core.BlobRequestHeader{
			SecurityParams: securityParams,
		},
		Data: []byte("test"),
	}
	s3Client   = cmock.NewS3Client()
	bucketName = "test-zgda-blobstore"
	blobHash   = "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	blobSize   = uint(len(blob.Data))

	dockertestPool     *dockertest.Pool
	dockertestResource *dockertest.Resource

	deployLocalStack bool
	localStackPort   = "4569"

	dynamoClient      *dynamodb.Client
	blobMetadataStore *blobstore.BlobMetadataStore
	sharedStorage     *blobstore.SharedBlobStore

	UUID              = uuid.New()
	metadataTableName = fmt.Sprintf("test-BlobMetadata-%v", UUID)
)

func TestMain(m *testing.M) {
	setup(m)
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup(m *testing.M) {

	deployLocalStack = !(os.Getenv("DEPLOY_LOCALSTACK") == "false")
	if !deployLocalStack {
		localStackPort = os.Getenv("LOCALSTACK_PORT")
	}

	if deployLocalStack {
		var err error
		dockertestPool, dockertestResource, err = deploy.StartDockertestWithLocalstackContainer(localStackPort)
		if err != nil {
			teardown()
			panic("failed to start localstack container")
		}

	}

	cfg := aws.ClientConfig{
		Region:          "us-east-1",
		AccessKey:       "localstack",
		SecretAccessKey: "localstack",
		EndpointURL:     fmt.Sprintf("http://0.0.0.0:%s", localStackPort),
	}

	_, err := test_utils.CreateTable(context.Background(), cfg, metadataTableName, blobstore.GenerateTableSchema(metadataTableName, 10, 10))
	if err != nil {
		teardown()
		panic("failed to create dynamodb table: " + err.Error())
	}

	dynamoClient, err = dynamodb.NewClient(cfg, logger)
	if err != nil {
		teardown()
		panic("failed to create dynamodb client: " + err.Error())
	}

	blobMetadataStore = blobstore.NewBlobMetadataStore(dynamoClient, logger, metadataTableName, time.Hour)
	sharedStorage = blobstore.NewSharedStorage(bucketName, s3Client, blobMetadataStore, logger)
}

func teardown() {
	if deployLocalStack {
		deploy.PurgeDockertestResources(dockertestPool, dockertestResource)
	}
}
