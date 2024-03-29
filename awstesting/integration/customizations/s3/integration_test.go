// +build integration

// Package s3 runs integration tests for S3
package s3

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/shztki/nifcloud-sdk-go/nifcloud"
	"github.com/shztki/nifcloud-sdk-go/service/s3"
)

func TestWriteToObject(t *testing.T) {
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: bucketName,
		Key:    nifcloud.String("key name"),
		Body:   bytes.NewReader([]byte("hello world")),
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	resp, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: bucketName,
		Key:    nifcloud.String("key name"),
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	b, _ := ioutil.ReadAll(resp.Body)
	if e, a := []byte("hello world"), b; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestPresignedGetPut(t *testing.T) {
	putreq, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: bucketName,
		Key:    nifcloud.String("presigned-key"),
	})
	var err error

	// Presign a PUT request
	var puturl string
	puturl, err = putreq.Presign(300 * time.Second)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// PUT to the presigned URL with a body
	var puthttpreq *http.Request
	buf := bytes.NewReader([]byte("hello world"))
	puthttpreq, err = http.NewRequest("PUT", puturl, buf)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	var putresp *http.Response
	putresp, err = http.DefaultClient.Do(puthttpreq)
	if err != nil {
		t.Errorf("expect put with presign url no error, got %v", err)
	}
	if e, a := 200, putresp.StatusCode; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}

	// Presign a GET on the same URL
	getreq, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: bucketName,
		Key:    nifcloud.String("presigned-key"),
	})

	var geturl string
	geturl, err = getreq.Presign(300 * time.Second)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	// Get the body
	var getresp *http.Response
	getresp, err = http.Get(geturl)
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	var b []byte
	defer getresp.Body.Close()
	b, err = ioutil.ReadAll(getresp.Body)
	if e, a := "hello world", string(b); e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
