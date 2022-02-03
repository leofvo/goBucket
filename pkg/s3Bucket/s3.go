package s3

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/LeoFVO/goBucket/helpers"
)

type s3ListBucketResult struct {
	XMLName xml.Name `xml:"ListBucketResult"`
	Name    string   `xml:"Name"`
	Prefix  string   `xml:"Prefix"`
	Marker  string   `xml:"Marker"`
	MaxKeys int      `xml:"MaxKeys"`
	IsTruncated bool   `xml:"IsTruncated"`
	Contents  []s3Content  `xml:"Contents"`
}
type s3Content struct {
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int    `xml:"Size"`
	StorageClass string `xml:"StorageClass"`
}

func CheckS3Url(url string, onlyCritical bool) {
	resp, err := http.Get(url)

	if err == nil {
		switch resp.StatusCode {
		case 200:
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				panic(err)
			}

			var result s3ListBucketResult
			if err := xml.Unmarshal(data, &result); err != nil {
        panic(err)
			}
			fmt.Printf("\033[1;32m[%s] Open bucket.\033[0m \n", formatUrl(url))
			for _, content := range result.Contents {
				if onlyCritical {
					if isFileCritical(content.Key) {
						getUrlContent(url + "/" + content.Key)
					}
				} else {
					getUrlContent(url + "/" + content.Key)
				}
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
}

func isFileCritical(file string) bool {
	criticalList, err := os.Open("wordlists/criticals")
	if err != nil {
		log.Fatal(err)
	}
	defer criticalList.Close()

	scanner := bufio.NewScanner(criticalList)
	for scanner.Scan() {
		if strings.Contains(file, scanner.Text()) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}


func getUrlContent(url string) {
	fmt.Printf("\033[1;34m%s\033[0m \n", formatUrl(url))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func formatUrl(url string) string {
	return strings.Replace(url, " ", "%20", -1)
}

func GetS3UrlFromBucket(bucket string) string {
	return fmt.Sprintf("http://%s.s3.amazonaws.com", formatUrl(bucket))
}

func Execute(wordlists string, onlyCritical bool) {
	bucketList := helpers.GetBucketsFromList(wordlists)
	for _, bucket := range bucketList {
		CheckS3Url(GetS3UrlFromBucket(bucket), onlyCritical)
	}
}
