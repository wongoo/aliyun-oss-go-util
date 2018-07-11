// authors: wangoo
// created: 2018-07-11

// aliyun oss golang utility
package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"io/ioutil"
	"flag"
)

var (
	client *oss.Client
	bucket *oss.Bucket

	accessKeyId     string
	accessKeySecret string
	endpoint        string
	bucketName      string
	action          string
	filePath        string
	urlPrefix       string
)

func main() {
	flag.StringVar(&accessKeyId, "access_key_id", "", "access key id")
	flag.StringVar(&accessKeySecret, "access_key_secret", "", "access key secret")
	flag.StringVar(&endpoint, "endpoint", "", "endpoint")
	flag.StringVar(&bucketName, "bucket", "", "bucket")
	flag.StringVar(&action, "action", "", "calling action name")
	flag.StringVar(&filePath, "file_path", "", "file path")
	flag.StringVar(&urlPrefix, "url_prefix", "", "url prefix")

	flag.Parse()

	notNil(accessKeyId, "access_key_id")
	notNil(accessKeySecret, "access_key_secret")
	notNil(endpoint, "endpoint")
	notNil(bucketName, "bucket")
	notNil(action, "action")

	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	checkErr(err)

	bucket, err = client.Bucket(bucketName)
	checkErr(err)

	if action == "batch_delete_url_in_file" {
		batchDeleteUrlInFile()
	}
}

func batchDeleteUrlInFile() {
	notNil(filePath, "file_path")
	notNil(urlPrefix, "url_prefix")

	b, err := ioutil.ReadFile(filePath)
	checkErr(err)

	UriBatchProcess(urlPrefix, b, DeleteObject)
}

func DeleteObject(path string) {
	fmt.Printf("delete object %v\n", path)
	err := bucket.DeleteObject(path)
	if err != nil {
		fmt.Printf("delete object error: %v\n", err)
	}
}

func notNil(field string, fieldName string) {
	if field == "" {
		fmt.Printf("require field %v\n", fieldName)
		os.Exit(-1)
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(-1)
	}
}
