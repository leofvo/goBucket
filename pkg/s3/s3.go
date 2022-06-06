package s3

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"

	tools "github.com/LeoFVO/goBucket/tools"
)


func checkBucket(url string, onlyCritical bool) {
	resp := tools.GetUrlContent(url)
	defer resp.Body.Close()

	switch resp.StatusCode {
		case 200:
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var result S3Bucketstruct
			if err := xml.Unmarshal(data, &result); err != nil {
				panic(err)
			}

			fmt.Printf("\033[1;32m[%s] Open bucket.\033[0m \n", tools.FormatUrl(url))
			for _, content := range result.Contents {
				if onlyCritical {
					if !tools.UrlLookCritical(content.Key) {
						break
					}
				} 
				tools.AddValidUrl("./found.txt",url + "/" + content.Key)
			}
			// case 401:
			// 	fmt.Printf("[%s] Unauthorized to access bucket.\n", url)
			// case 403:
			// 	fmt.Printf("[%s] Bucket seems to be forbidden.\n", url)
			// case 404:
			// 	fmt.Printf("[%s] No buckets founds.\n", url)
			// default:
			// 	fmt.Printf("[%s] Bucket inaccessible.\n", url)
	}
}

func GenerateS3UrlFromBucket(bucket string) string {
	return fmt.Sprintf("http://%s.s3.amazonaws.com", tools.FormatUrl(bucket))
}

func Execute(mode string, arg string, onlyCritical bool) {
	// Mode url == single check
	if mode == "url" {
		checkBucket(arg, onlyCritical)
	} else if mode == "wordlist" {
		bucketList := tools.GetBucketsFromList(arg)
		for _, bucket := range bucketList {
			checkBucket(GenerateS3UrlFromBucket(bucket), onlyCritical)
		}
	} else {
		fmt.Println("[-] Invalid mode. Use 'url' or 'wordlist'")
		return
	}
}