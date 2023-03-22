package todo

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type GRPCServerConfig struct {
	Stdout io.Writer

	Stderr io.Writer

	New func(context.Context) (*grpc.Server, error)
}

func defaultListenAddr() string {
	// 環境変数PORTで与えられたポート番号で待ち受ける
	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = "8000"
	}
	return ":" + portStr
}

func defaultLogLevel() string {
	// "ログレベル (choices: trace, debug, info, warn, error)"
	if portStr, exist := os.LookupEnv("LOG_LEVEL"); exist {
		return portStr
	} else {
		return "info"
	}
}

func initLogger(loglevel string, stderr io.Writer) error {
	if lvl, err := logrus.ParseLevel(loglevel); err != nil {
		return err
	} else {
		logrus.SetLevel(lvl)
	}
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	return nil
}

func RunGRPCServer(cfg *GRPCServerConfig) int {
	logrus.SetOutput(cfg.Stdout)

	var (
		listenAddr string = defaultListenAddr()
		logLevel   string = defaultLogLevel()
	)

	if err := initLogger(logLevel, cfg.Stderr); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	initCtx := context.Background()
	initCtx, cancel := context.WithTimeout(initCtx, 10*time.Second)
	defer cancel()

	srv, err := cfg.New(initCtx)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}

	logrus.Infof("listening on %s ...", listenAddr)
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	defer l.Close()

	if err := srv.Serve(l); err != nil {
		fmt.Fprintln(cfg.Stderr, err)
		return 1
	}
	return 0
}
