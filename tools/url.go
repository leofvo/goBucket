package tools

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func FormatUrl(url string) string {
	return strings.Replace(url, " ", "%20", -1)
}

func GetUrlContent(url string) {
	fmt.Printf("\033[1;34m%s\033[0m \n", FormatUrl(url))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
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