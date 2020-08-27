package config

import (
	"os"
	"testing"

	"github.com/matryer/is"
)

func TestEnvRegion(t *testing.T) {
	is := is.NewRelaxed(t)

	tests := []struct {
		isSuccess bool
		envName   string
		region    string
	}{
		{true, "AWS_REGION", "foo"},
		{true, "AWS_REGION", "bar"},
		{false, "AWS_REGION1", "xxx"},
		{false, "AWS_REGION2", "xxx"},
		{false, "AWS_REGIO", "xxx"},
	}

	defer os.Clearenv()
	for _, tt := range tests {
		os.Clearenv()

		is.Equal("", EnvRegion()) // initial env should be empty

		os.Setenv(tt.envName, tt.region)
		if !tt.isSuccess {
			is.Equal("", EnvRegion()) // invalid env should be empty
			return
		}

		is.Equal(tt.region, EnvRegion()) // valid env should be match
		os.Clearenv()
		is.Equal("", EnvRegion()) // cleared env should be empty
	}
}

func TestEnvEndpoint(t *testing.T) {
	is := is.NewRelaxed(t)

	tests := []struct {
		isSuccess bool
		envName   string
		endpoint  string
	}{
		{true, "AWS_ENDPOINT", "foo"},
		{true, "AWS_ENDPOINT", "bar"},
		{false, "AWS_ENDPOINT1", "xxx"},
		{false, "AWS_ENDPOINT2", "xxx"},
		{false, "AWS_ENDPOIN", "xxx"},
	}

	defer os.Clearenv()
	for _, tt := range tests {
		os.Clearenv()

		is.Equal("", EnvEndpoint()) // initial env should be empty

		os.Setenv(tt.envName, tt.endpoint)
		if !tt.isSuccess {
			is.Equal("", EnvEndpoint()) // invalid env should be empty
			return
		}

		is.Equal(tt.endpoint, EnvEndpoint()) // valid env should be match
		os.Clearenv()
		is.Equal("", EnvEndpoint()) // cleared env should be empty
	}
}
