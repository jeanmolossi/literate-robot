package client

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"strconv"

	"github.com/jeanmolossi/literate-robot/src/common/genproto/job"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

func NewJobClient() (client job.JobServiceClient, close func() error, err error) {
	grpcAddr := os.Getenv("JOB_GRPC_ADDR")
	if grpcAddr == "" {
		return nil, func() error { return nil }, errors.New("empty env JOB_GRPC_ADDR")
	}

	opts, err := grpcDialOpts(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.Dial(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	return job.NewJobServiceClient(conn), conn.Close, nil
}

func grpcDialOpts(grpcAddr string) ([]grpc.DialOption, error) {
	if noTLS, _ := strconv.ParseBool(os.Getenv("GRPC_NO_TLS")); noTLS {
		return []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}, nil
	}

	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil, errors.Wrap(err, "cannot load root CA cert")
	}

	creds := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})

	return []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(newMetadataServerToken(grpcAddr)),
	}, nil
}
