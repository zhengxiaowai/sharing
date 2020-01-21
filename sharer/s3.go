package sharer

import (
	"context"
	"encoding/json"
	"github.com/minio/minio-go/v6"
	"io/ioutil"
	"mime"
	"os"
	"path/filepath"
)

type S3Sharer struct {
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	EndPoint   string `json:"endpoint"`
	BucketName string `json:"bucket_name"`
	UseSSL     bool   `json:"use_ssl"`
	UseDomain  string `json:"use_domain"`
}

func (s *S3Sharer) GetName() string {
	return "s3"
}

func (s *S3Sharer) InitConfig(confPath string) error {
	conf, err := ioutil.ReadFile(confPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(conf, s)
	if err != nil {
		return err
	}

	return nil
}

func (s *S3Sharer) UploadFile(key string, filePath string) (string, error) {
	minioClient, err := minio.New(s.EndPoint, s.AccessKey, s.SecretKey, s.UseSSL)
	if err != nil {
		return "", err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return "", err
	}

	opts := minio.PutObjectOptions{UserMetadata: map[string]string{"x-amz-acl": "public-read"}}
	if opts.ContentType == "" {
		if opts.ContentType = mime.TypeByExtension(filepath.Ext(filePath)); opts.ContentType == "" {
			opts.ContentType = "application/octet-stream"
		}
	}

	_, err = minioClient.PutObjectWithContext(context.Background(), s.BucketName, key, file, fileStat.Size(), opts)
	if err != nil {
		return "", err
	}

	return makePublicURL(s.UseDomain, key), nil
}
