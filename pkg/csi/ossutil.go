package csi

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type myoss struct {
	Endpoint string
	AK       string
	SK       string
	client   *oss.Client
}

func (o myoss) NetCliten() *oss.Client {
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
	return client
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
	var err error
	if exist, _ := isBucketExist(oss.Bucket{
		Client: *o.client, BucketName: bucketName,
	}); !exist {
		// 创建存储空间。
		err := o.client.CreateBucket(bucketName)
		if err != nil {
			handleError(err)
		}
	}
	return err
}

func (o myoss) CreateDir(bucketName string, dir string) error {
	// 填写存储空间名称，例如examplebucket。
	bucket, err := o.client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	err = bucket.PutObject(dir+"/", bytes.NewReader([]byte("")))
	if err != nil {
		fmt.Println("Error:", err)
	}
	return err
}
func (o myoss) DeleteDir(bucketName string, dir string) error {
	// 填写Bucket名称。
	bucket, err := o.client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	marker := oss.Marker("")
	// 填写待删除目录的完整路径，完整路径中不包含Bucket名称。
	prefix := oss.Prefix(dir + "/")
	count := 0
	for {
		lor, err := bucket.ListObjects(marker, prefix)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		objects := []string{}
		for _, object := range lor.Objects {
			objects = append(objects, object.Key)
		}
		// 删除目录及目录下的所有文件。
		// 将oss.DeleteObjectsQuiet设置为true，表示不返回删除结果。
		delRes, err := bucket.DeleteObjects(objects, oss.DeleteObjectsQuiet(true))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}

		if len(delRes.DeletedObjects) > 0 {
			fmt.Println("these objects deleted failure,", delRes.DeletedObjects)
			os.Exit(-1)
		}

		count += len(objects)

		prefix = oss.Prefix(lor.Prefix)
		marker = oss.Marker(lor.NextMarker)
		if !lor.IsTruncated {
			break
		}
	}
	fmt.Printf("success,total delete object count:%d\n", count)
	return err
}
