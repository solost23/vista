package base

import (
	"context"
	"github.com/gookit/slog"
	"github.com/minio/minio-go"

	"gorm.io/gorm"

	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/solost23/protopb/gen/go/protos/common"
)

type Action struct {
	ctx           context.Context
	sl            *slog.SugaredLogger
	mysqlConnect  *gorm.DB
	redisClient   *redis.Client
	kafkaProducer sarama.SyncProducer
	minioClient   *minio.Client
	traceId       int64
	operator      int32
}

func (a *Action) SetContext(ctx context.Context) {
	a.ctx = ctx
	return
}

func (a *Action) SetSl(sl *slog.SugaredLogger) {
	a.sl = sl
	return
}

func (a *Action) SetHeader(header *common.RequestHeader) {
	a.traceId = header.TraceId
	a.operator = header.OperatorUid
	return
}

func (a *Action) SetMysql(mysqlConn *gorm.DB) {
	a.mysqlConnect = mysqlConn.WithContext(a.ctx)
	return
}

func (a *Action) SetMinio(minioClient *minio.Client) {
	a.minioClient = minioClient
}

func (a *Action) SetkafkaProducer(kafkaProducer sarama.SyncProducer) {
	a.kafkaProducer = kafkaProducer
	return
}

func (a *Action) GetSl() *slog.SugaredLogger {
	return a.sl
}

func (a *Action) GetMinioClient() *minio.Client {
	return a.minioClient
}

func (a *Action) GetTraceId() int64 {
	return a.traceId
}

func (a *Action) GetOperator() int32 {
	return a.operator
}

func (this *Action) GetMysqlConnect() *gorm.DB {
	return this.mysqlConnect
}

func (this *Action) GetRedisClient() *redis.Client {
	return this.redisClient
}

func (this *Action) GetKafkaProducer() sarama.SyncProducer {
	return this.kafkaProducer
}

func (*Action) BuildError(code int32, msg string, header *common.RequestHeader) *common.ErrorInfo {
	if header == nil {
		header = new(common.RequestHeader)
	}
	return &common.ErrorInfo{
		Requester:   header.Requester,
		OperatorUid: header.OperatorUid,
		TraceId:     header.TraceId,
		Code:        code,
		Msg:         msg,
	}
}
