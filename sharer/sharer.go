package sharer

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

const defaultConfDirectory  = "sharing"

type Sharer interface {
	GetName() string
	InitConfig(confPath string) error
	UploadFile(key string, filePath string) (string, error)
}

func mustGetUserHomeDir() string {
	userHomeDir, _ := os.UserHomeDir()
	return userHomeDir
}

func getSharerConf(sharer Sharer) string {
	return filepath.Join(mustGetUserHomeDir(), defaultConfDirectory, fmt.Sprintf("%s.%s", sharer.GetName(), "json"))
}

func makePublicURL(domain, key string) string {
	u, _ := url.Parse(domain)
	u.Path = path.Join(u.Path, key)
	return u.String()
}

func getFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}