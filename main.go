package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func ImagesToBase64(str_images string) []byte {
	//读原图片
	ff, _ := os.Open(str_images)
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])
	return []byte(sourcestring)
}

func Base64ToImage(sourcestring []byte) {
	// 写入临时文件
	ioutil.WriteFile("a.png.txt", sourcestring, 0667)
	// 读取临时文件
	cc, _ := ioutil.ReadFile("a.png.txt")

	// 解压
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	// 写入新文件
	f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	return
}

func main() {
	base64Bytes := ImagesToBase64("/home/fangman/Downloads/1564733186636.jpg")
	fmt.Printf("%v\n",string(base64Bytes))
}