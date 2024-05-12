package media

import (
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type Config struct {
	Address   string
	EnableTLS bool
}

func NewServiceClient(cfg Config) (c MediaServiceClient, err error) {
	var cc *grpc.ClientConn
	if cfg.EnableTLS {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		cc, err = grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}
	if err != nil {
		return nil, err
	}
	return NewMediaServiceClient(cc), nil
}
