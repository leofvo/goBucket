package s3

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	tools "github.com/LeoFVO/goBucket/tools"
)


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

				var result S3Bucketstruct
				if err := xml.Unmarshal(data, &result); err != nil {
					panic(err)
				}
				fmt.Printf("\033[1;32m[%s] Open bucket.\033[0m \n", tools.FormatUrl(url))
				for _, content := range result.Contents {
					if onlyCritical {
						if tools.UrlLookCritical(content.Key) {
							tools.GetUrlContent(url + "/" + content.Key)
						}
					} else {
						tools.GetUrlContent(url + "/" + content.Key)
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

// func main() {
// 	fileUrl := "https://golangcode.com/logo.svg"
// 	err := DownloadFile("logo.svg", fileUrl)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Downloaded: " + fileUrl)
// }

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error { // TODO: Should be in helper

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GenerateS3UrlFromBucket(bucket string) string {
	return fmt.Sprintf("http://%s.s3.amazonaws.com", tools.FormatUrl(bucket))
}

func Execute(mode string, arg string, onlyCritical bool) {
	// Mode url == single check
	if mode == "url" {
		CheckS3Url(arg, onlyCritical)
	} else if mode == "wordlist" {
		bucketList := tools.GetBucketsFromList(arg)
		for _, bucket := range bucketList {
			CheckS3Url(GenerateS3UrlFromBucket(bucket), onlyCritical)
		}
	} else {
		fmt.Println("[-] Invalid mode. Use 'url' or 'wordlist'")
		return
	}
}