package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
)

/*
* @Author:hanyajun
* @Date:2019/5/22 0:01
* @Name:cmd
* @Function:
 */

func main() {
	fileList := readDirectory("need_upload_data")
	printFileList(fileList)
	t, _ := template.ParseFiles("templates/success.html")
	fmt.Println(t.Name())
	fileObj, err := os.OpenFile("success.html", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to open the file", err.Error())
		os.Exit(2)
	}
	_ = t.Execute(fileObj, fileList)
}

func readDirectory(dir string) (fl []string) {

	files, _ := ioutil.ReadDir(dir)

	var fileList []string
	fileList = make([]string, len(files))

	i := 0
	for _, file := range files {
		if file.IsDir() {
			continue
		} else {
			fileList[i] = file.Name()
			i++
		}
	}
	return fileList
}

func printFileList(fl []string) {
	for i := 0; i < len(fl); i++ {
		fmt.Println(fl[i])
	}
}
