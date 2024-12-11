package config

import (
	"fmt"
	"os"
)

// READ ENV DETAILS
func LoadStorageBucketDetails(key string) (string, error){
	bucketKey , ok := os.LookupEnv(key)
	if !ok{
		return "", fmt.Errorf("no %s key set in the Env Variable", key)
	}
	if bucketKey == ""{
		return "", fmt.Errorf("no %s key set in the Env Variable", key)
	}
	return bucketKey, nil
}

