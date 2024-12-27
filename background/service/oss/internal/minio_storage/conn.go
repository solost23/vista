package minio_storage

import (
	"oss_service/configs"

	"github.com/minio/minio-go"
)

func NewMinio(config *configs.MinioConf) (client *minio.Client, err error) {
	client, err = minio.New(
		config.EndPoint,
		config.SecretAccessKey,
		config.AccessKeyId,
		config.UserSsl,
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}
