package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/etcdv3"
)

func UploadFile(header *multipart.FileHeader) (string, error) {
	filename := time.Now().Format("20060102150405") + header.Filename
	// 尝试创建此路径
	mkdirErr := os.MkdirAll("./testdata/", os.ModePerm)
	if mkdirErr != nil {
		return "", errors.New("function os.MkdirAll() Filed, err:" + mkdirErr.Error())
	}
	newName := "./testdata/" + filename

	f, openError := header.Open() // 读取文件
	if openError != nil {
		return "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(newName)
	if createErr != nil {
		return "", errors.New("function os.Create() Filed, err:" + createErr.Error())
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		return "", errors.New("function io.Copy() Filed, err:" + copyErr.Error())
	}
	return newName, nil
}

func UploadHandler(c *gin.Context) {
	fmt.Println("Receive")
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传文件失败",
		})
		return
	}
	name := ""
	name, err = UploadFile(header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "上传文件失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": name,
	})
}

func main() {
	// TODO: use config file.
	var (
		etcdAddrs = []string{"127.0.0.1:2379"}
		serName   = "svc.file"
		grpcAddr  = "127.0.0.1:8882"
		ttl       = 5 * time.Second
	)

	// 初始化etcd客户端
	options := etcdv3.ClientOptions{
		DialTimeout:   ttl,
		DialKeepAlive: ttl,
	}
	etcdClient, err := etcdv3.NewClient(context.Background(), etcdAddrs, options)
	if err != nil {
		fmt.Println(err)
		return
	}
	Registar := etcdv3.NewRegistrar(etcdClient, etcdv3.Service{
		Key:   serName,
		Value: grpcAddr,
	}, log.NewNopLogger())

	// 注册 etcd
	Registar.Register()

	var Router = gin.Default()
	ApiGroup := Router.Group("")

	FileUploadAndDownloadGroup := ApiGroup.Group("file")
	{
		FileUploadAndDownloadGroup.POST("/upload", UploadHandler) // 上传文件
	}

	Router.Run(":8882")
}
