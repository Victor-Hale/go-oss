package OSS

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
	"os"
)
func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}

var red string
func Oss(file multipart.File,Header *multipart.FileHeader,endpoint , accessKeyId ,accessKeySecret,bucketName,objectName string) string {
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。

	err = bucket.PutObject(objectName, file, oss.ContentType("image/jpg"))
	if err != nil {
		handleError(err)
	}
	//获取url并返回
	signedURL, err := bucket.SignURL(objectName, oss.HTTPPut, 60)
	for k, v := range signedURL {
		if v == '?' {
			red = signedURL[0:k]
		}
	}
	if err != nil {
		handleError(err)
	} else {
		return red
	}
	return red
}

