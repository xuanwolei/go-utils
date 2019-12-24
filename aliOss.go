package goutils

import(
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliOssConf struct{
	AccessKeyId string
	AccessKeySecret string
	Endpoint string
	Bucket   string
	Prefix   string
}

type AliOss struct{
	Client *oss.Client	
	Bucket *oss.Bucket
	Conf   *AliOssConf
}

func NewAliOss(conf *AliOssConf)(*AliOss, error){

	client, err := oss.New(conf.Endpoint, conf.AccessKeyId, conf.AccessKeySecret)
    if err != nil {
        return nil, err
    }
	//select bucket
    bucket, err := client.Bucket(conf.Bucket)
    if err != nil {
        return nil, err
	}
	
	return &AliOss{Client:client, Bucket:bucket, Conf: conf,}, nil
}

func (this *AliOss) PutObjectFromFile(object string, filePath string) (error){
	//上传文件到bucket
	err := this.Bucket.PutObjectFromFile(this.ParseObject(object), filePath)
	return err
}

func (this *AliOss) ParseObject(object string) (string){
	return this.Conf.Prefix + object
}

