package dynamodb

import (
	"context"
	"testing"

	SDK "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/matryer/is"

	"github.com/evalphobia/aws-sdk-go-v2-wrapper/config"
	"github.com/evalphobia/aws-sdk-go-v2-wrapper/log"
)

const (
	defaultEndpoint = "http://localhost:8000"

	testTableName = "test-table"
)

func TestNew(t *testing.T) {
	is := is.NewRelaxed(t)

	// custom endpoint
	svc, err := New(getTestConfig())
	is.NoErr(err)
	is.True(svc.client != nil)
	is.Equal(SDK.ServiceName, svc.client.Metadata.ServiceName)
	is.Equal(SDK.ServiceID, svc.client.Metadata.ServiceID)
	is.Equal(SDK.EndpointsID, svc.client.Metadata.EndpointsID)

	ep, err := svc.client.EndpointResolver.ResolveEndpoint("", "")
	is.NoErr(err)
	is.Equal(defaultEndpoint, ep.URL)

	// plain region
	svc, err = New(config.Config{})
	is.NoErr(err)
	is.True(svc.client != nil)

	region := "us-west-1"
	expectedEndpoint := "https://dynamodb." + region + ".amazonaws.com"
	ep, err = svc.client.EndpointResolver.ResolveEndpoint(SDK.EndpointsID, region)
	is.NoErr(err)
	is.Equal(expectedEndpoint, ep.URL)

	_, err = svc.client.EndpointResolver.ResolveEndpoint("", "")
	is.True(err != nil)
}

func TestSetLogger(t *testing.T) {
	is := is.NewRelaxed(t)
	svc := getTestClient(t)
	is.Equal(log.DefaultLogger, svc.logger)

	stdLogger := &log.StdLogger{}
	svc.SetLogger(stdLogger)
	is.Equal(stdLogger, svc.logger)
}

func getTestConfig() config.Config {
	return config.Config{
		AccessKey:      "access",
		SecretKey:      "secret",
		CommonEndpoint: defaultEndpoint,
	}
}

func getTestClient(t *testing.T) *DynamoDB {
	svc, err := New(getTestConfig())
	if err != nil {
		t.Errorf("error on create client; error=%s;", err.Error())
		t.FailNow()
	}
	return svc
}

func createTestTable(name string) error {
	svc, err := New(getTestConfig())
	if err != nil {
		return err
	}

	ctx := context.Background()
	ok, err := svc.XExistTable(ctx, name)
	switch {
	case err != nil:
		return err
	case ok:
		return nil
	}

	_, err = svc.CreateTable(ctx, CreateTableRequest{
		TableName: name,
		AttributeDefinitions: []AttributeDefinition{
			{AttributeName: "hashkey",
				AttributeType: ScalarAttributeTypeS},
		},
		KeySchema: []KeySchemaElement{
			{AttributeName: "hashkey",
				KeyType: KeyTypeHash},
		},
		ProvisionedThroughput: ProvisionedThroughput{
			ReadCapacityUnits:  5,
			WriteCapacityUnits: 5,
		},
	})
	return err
}
