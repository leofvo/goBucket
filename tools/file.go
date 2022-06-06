package tools

import "os"

func AddValidUrl(file string, url string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString(url + "\n"); err != nil {	// Write url to file
		panic(err)
	}
}