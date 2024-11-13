package main

import (
	"log"
	"net/http"

	"github.com/gabriel-vasile/mimetype"
)

func main() {
	url := "https://cos-console-dev.yuni.vip/v1/cos/file/6dba558a-33ac-4468-ac65-216fbb618f2f?ApiKey=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzM1NjM1MjIsInVpZCI6ImNzX2VicmNnRFM3MDIiLCJzaWQiOiIzMDlhNDEzMTZlMmQ2ZjE4ZmFiYzRkNGZhNWJiZGIxYmU3N2FjYjQ0NjVjMWQyNjM4NzRiMGYxMDMwNDRmOTdmIn0.6aFVN1eGqj3jxsF_bLcHBoUGlxz2EGQWtoVUn-vSc3w&partid=0"
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	// 读取512字节
	buffer := make([]byte, 512)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	contentType := mimetype.Detect(buffer).String()
	log.Println(contentType)
}
