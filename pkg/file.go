package pkg

import (
	"fmt"
	"io/ioutil"
)

/*
* @Author:hanyajun
* @Date:2019/5/25 0:48
* @Name:pkg
* @Function:
 */

/*
 * 读取目录内的文件
 */
func ReadDirectory(dir string) (fl []string) {

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

func PrintFileList(fl []string) {
	for i := 0; i < len(fl); i++ {
		fmt.Println(fl[i])
	}
}
