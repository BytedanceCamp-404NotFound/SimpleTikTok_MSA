// Code generated by goctl. DO NOT EDIT.
// Source: minio.proto

package minioclient

import (
	"context"

	"SimpleTikTok/internal/MicroServices/pkg/Minio"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	IdRequest     = Minio.IdRequest
	MinioResponse = Minio.MinioResponse

	Minio interface {
		GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error)
	}

	defaultMinio struct {
		cli zrpc.Client
	}
)

func NewMinio(cli zrpc.Client) Minio {
	return &defaultMinio{
		cli: cli,
	}
}

func (m *defaultMinio) GetMinio(ctx context.Context, in *IdRequest, opts ...grpc.CallOption) (*MinioResponse, error) {
	client := Minio.NewMinioClient(m.cli.Conn())
	return client.GetMinio(ctx, in, opts...)
}
