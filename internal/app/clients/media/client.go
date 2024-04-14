package media

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewServiceClient() (MediaServiceClient, error) {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial("media:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return NewMediaServiceClient(cc), nil
}
