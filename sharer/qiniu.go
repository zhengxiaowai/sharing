package sharer

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type QiNiuSharer struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Bucket        string `json:"bucket"`
	Zone          string `json:"zone"`
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
}

func (q *QiNiuSharer) InitConfig(qiNiuConf string) error {
	conf, err := ioutil.ReadFile(qiNiuConf)
	if err != nil {
		return err
	}

	err = json.Unmarshal(conf, q)
	if err != nil {
		return err
	}

	return nil

}

func (q *QiNiuSharer) UploadFile(key string, filePath string) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}

	var zone *storage.Zone
	switch strings.ToLower(q.Zone) {
	case "zonehuadong":
		zone = &storage.ZoneHuadong
	case "zonehuabei":
		zone = &storage.ZoneHuabei
	case "zonehuanan":
		zone = &storage.ZoneHuanan
	case "zonebeimei":
		zone = &storage.ZoneBeimei
	default:
		zone = nil
	}

	if zone == nil {
		return "", fmt.Errorf("invalid zone: %s", q.Zone)
	}

	cfg := storage.Config{
		Zone:          zone,
		UseHTTPS:      q.UseHTTPS,
		UseCdnDomains: q.UseCdnDomains,
	}

	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	ret := storage.PutRet{}
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": filepath.Base(filePath),
		},
	}

	uploader := storage.NewFormUploader(&cfg)
	err = uploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(content), int64(len(content)), &putExtra)
	if err != nil {
		return "", err
	}

	return makePublicURL(q.Domain, ret.Key), nil
}

func (q *QiNiuSharer) GetName() string {
	return "qiniu"
}
