package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"qiniuyun_upload_tools/pkg"
	"time"
)

/*
* @Author:hanyajun
* @Date:2019/5/22 0:01
* @Name:cmd
* @Function:
 */

func main() {
	config := pkg.LoadConfig()
	client := pkg.NewClient(config.AccessKey, config.SecretKey, config.Bucket, config.Zone, config.UseHTTPS, config.UseCdnDomains, config.Domain)
	fileList := pkg.ReadDirectory("need_upload_data")
	pkg.PrintFileList(fileList)
	successFileList := client.UploadFile(fileList)
	t, _ := template.ParseFiles("templates/success.html")
	successFileName := time.Now().Format("20060102150405") + "_success.html"
	fileObj, err := os.OpenFile(successFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	_ = t.Execute(fileObj, successFileList)
	cmd := exec.Command("explorer", successFileName)
	cmd.Start()
}
