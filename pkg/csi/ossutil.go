package csi

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type myoss struct {
	Endpoint string
	AK       string
	SK       string
}

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func isBucketExist(bucket oss.Bucket) (bool, error) {
	isExist, err := bucket.Client.IsBucketExist(bucket.BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	return isExist, err
}
func (o myoss) CreateBucekt(bucketName string) error {
	// yourBucketName填写Bucket名称。
	_ = os.Setenv("OSS_ACCESS_KEY_ID", o.AK)
	_ = os.Setenv("OSS_ACCESS_KEY_SECRET", o.SK)
	// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New(o.Endpoint, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	if exist, _ := isBucketExist(oss.Bucket{
		Client: *client, BucketName: bucketName,
	}); !exist {
		// 创建存储空间。
		err = client.CreateBucket(bucketName)
		if err != nil {
			handleError(err)
		}
	}
	return err
}
