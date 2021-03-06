package config

import (
	"fmt"
	"os"
)

type Config struct {
	Region 			string
	TempBucketName 	string
	LocalRun		bool
}

func FromEnvironment () (Config, error) {
	region := os.Getenv("MW_AWS_REGION")
	if region == "" {
		return Config{}, fmt.Errorf("MW_AWS_REGION env variable is not defined")
	}

	tempBucketName := os.Getenv("MW_TEMP_BUCKET_NAME")
	if tempBucketName == "" {
		return Config{}, fmt.Errorf("MW_TEMP_BUCKET_NAME env variable is not defined")
	}

	localRun := os.Getenv("MW_LOCAL_RUN") != ""

	return Config{
		Region: region,
		TempBucketName: tempBucketName,
		LocalRun: localRun,
	}, nil
}
