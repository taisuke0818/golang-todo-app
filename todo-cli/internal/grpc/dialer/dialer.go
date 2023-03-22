package dialer

import (
	"context"
	"fmt"
	"net/url"

	"google.golang.org/api/idtoken"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"
)

// urlStrに対し、TransportCredentialを付与したgRPC Client Connを取得する。
func DialContext(ctx context.Context, urlStr string) (grpc.ClientConnInterface, error) {
	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}
	var (
		defaultPortStr string
		opts           []grpc.DialOption
	)
	switch u.Scheme {
	case "http":
		defaultPortStr = ":80"
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case "https":
		defaultPortStr = ":443"
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
		if l, err := acquireCredentials(ctx, urlStr); err != nil {
			return nil, err
		} else {
			opts = append(opts, l...)
		}
	default:
		return nil, fmt.Errorf("invalid url: %q", urlStr)
	}
	endpoint := u.Host
	if s := u.Port(); s == "" {
		endpoint += defaultPortStr
	}
	return grpc.DialContext(ctx, endpoint, opts...)
}

func acquireCredentials(ctx context.Context, urlStr string) ([]grpc.DialOption, error) {
	ts, err := idtoken.NewTokenSource(ctx, urlStr)
	if err != nil {
		return nil, err
	}
	return []grpc.DialOption{
		grpc.WithPerRPCCredentials(oauth.TokenSource{
			TokenSource: ts,
		}),
	}, nil
}
