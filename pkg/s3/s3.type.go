package s3

import "encoding/xml"

type S3Bucketstruct struct {
	XMLName xml.Name `xml:"ListBucketResult"`
	Name    string   `xml:"Name"`
	Prefix  string   `xml:"Prefix"`
	Marker  string   `xml:"Marker"`
	MaxKeys int      `xml:"MaxKeys"`
	IsTruncated bool   `xml:"IsTruncated"`
	Contents  []S3BucketContent  `xml:"Contents"`
}

type S3BucketContent struct {
	Key          string `xml:"Key"`
	LastModified string `xml:"LastModified"`
	ETag         string `xml:"ETag"`
	Size         int    `xml:"Size"`
	StorageClass string `xml:"StorageClass"`
}