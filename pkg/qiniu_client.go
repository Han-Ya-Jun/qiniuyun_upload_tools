package pkg

import (
	"context"
	"fmt"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/5/22 15:16
* @Name:pkg
* @Function:
 */

type QiNiuClient struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Bucket        string `json:"bucket"`
	Zone          int    `json:"zone"` //0:华东, 1:华北, 2:华南, 3:北美
	UseHTTPS      bool   `json:"use_https"`
	UseCdnDomains bool   `json:"use_cdn_domains"`
	Domain        string `json:"domain"`
}

func NewClient(accessKey, secretKey, bucket string, Zone int, useHttps, useCdnDomains bool, domain string) *QiNiuClient {
	return &QiNiuClient{
		AccessKey:     accessKey,
		SecretKey:     secretKey,
		Bucket:        bucket,
		Zone:          Zone,
		UseHTTPS:      useHttps,
		UseCdnDomains: useCdnDomains,
		Domain:        domain,
	}
}

func (client *QiNiuClient) UploadFile(fileList []string) {
	putPolicy := storage.PutPolicy{
		Scope: client.Bucket,
	}
	mac := qbox.NewMac(client.AccessKey, client.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = client.UseHTTPS
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = client.UseCdnDomains
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		//Params: map[string]string{
		//	"x:name": "github logo",
		//},
	}
	key := time.Now().Unix() + "_"
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}
