package tools

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func FormatUrl(url string) string {
	return strings.Replace(url, " ", "%20", -1)
}

func GetUrlContent(url string) *http.Response {
	fmt.Printf("\033[1;34m%s\033[0m \n", FormatUrl(url))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func UrlLookCritical(url string) bool {
	criticalList, err := os.Open("wordlists/criticals")
	if err != nil {
		log.Fatal(err)
	}
	defer criticalList.Close()

	scanner := bufio.NewScanner(criticalList)
	for scanner.Scan() {
		if strings.Contains(url, scanner.Text()) {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return false
}

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