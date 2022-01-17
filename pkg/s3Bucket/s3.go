package s3

import (
	"fmt"
	"net/http"

	"github.com/LeoFVO/goBucket/helpers"
)

func CheckS3Url(url string) {
	resp, err := http.Get(url)
	
	if err == nil {	
		switch resp.StatusCode {
			case 200:
				fmt.Printf("[%s] Open bucket.\n", url)
			case 401:
				fmt.Printf("[%s] Unauthorized to access bucket.\n", url)
			case 403:
				fmt.Printf("[%s] Bucket seems to be forbidden.\n", url)
			case 404:
				fmt.Printf("[%s] No buckets founds.\n", url)
			default:
				fmt.Printf("[%s] Bucket inaccessible.\n", url)
		}
	}
}


func GetS3UrlFromBucket(bucket string) string {
	return fmt.Sprintf("http://%s.s3.amazonaws.com", bucket)
}

func Execute(wordlists string) {
	bucketList := helpers.GetBucketsFromList(wordlists)
	for _, bucket := range bucketList {
		CheckS3Url(GetS3UrlFromBucket(bucket))
	}
}