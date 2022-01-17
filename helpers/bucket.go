package helpers

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	str "strings"
)
var s3BucketNamePattern = regexp.MustCompile(`[a-zA-Z0-9\-\.]{3,63}`)

func ExtractBucketName(bucket string) (string, error) {
	switch {
		case str.HasPrefix(bucket, "http://s3.amazonaws.com/"):
			bucket = str.TrimPrefix(bucket, "http://s3.amazonaws.com/")
		case str.HasSuffix(bucket, ".amazonaws.com"):
			bucket = str.TrimSuffix(bucket, ".s3.amazonaws.com")
		case str.Contains(bucket, ":"):
			bucket = str.Split(bucket, ":")[0]
	}
	bucket = str.TrimPrefix(bucket, "http://")

	if !s3BucketNamePattern.MatchString(bucket) {
		return "", fmt.Errorf("invalid bucket name")
	}

	return bucket, nil
}

func GetBucketsFromList(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bucketList := make([]string, 0)

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		bucket, err := ExtractBucketName(scanner.Text())
		if err == nil {
			bucketList = append(bucketList, bucket)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return bucketList
}