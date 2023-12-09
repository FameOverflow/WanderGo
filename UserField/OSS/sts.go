package OSS

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var accessKeyId = "LTAI5t5rPhn5UbdoLLQncn8r"
var accessKeySecret = "KGJn7jyzCUBbX7wIgCFpVhngQhY7JK"

func STSUpload() {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "spark-forge"
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := "test/test.txt"
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	filepath := "D:\\music.txt"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println(err)
	}
	// 通过STS授权第三方上传文件。
	err = bucket.PutObjectFromFile(objectName, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("upload success")
}

func STSDownload() {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("http://oss-cn-hangzhou.aliyuncs.com", accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 填写Bucket名称，例如examplebucket。
	bucketName := "spark-forge"
	// 填写Object的完整路径，完整路径中不能包含Bucket名称，例如exampledir/exampleobject.txt。
	objectName := "test/test.txt"
	// 填写本地文件的完整路径，例如D:\\localpath\\examplefile.txt。
	filepath := "D:\\music.txt"
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		log.Println(err)
		return
	}
	// 通过STS授权第三方下载文件。
	err = bucket.GetObjectToFile(objectName, filepath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	fmt.Println("download success")
}
