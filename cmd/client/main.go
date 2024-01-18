package main

import (
	"context"
	"github.com/Genry72/GophKeeper/internal/client/usecase/grpcclient"
	"github.com/Genry72/GophKeeper/internal/client/usecase/tuiclient"
	"github.com/Genry72/GophKeeper/pkg/logger"
	"go.uber.org/zap"
)

const (
	grpcServerAddress = ":3200"
)

var (
	buildVersion string
	buildDate    string
)

func main() {
	// todo нужен ли?
	zapLogger := logger.NewZapLogger("info")

	defer func() {
		_ = zapLogger.Sync()
	}()

	zapLogger.Info("build version:\t" + buildVersion)
	zapLogger.Info("build date:\t" + buildDate)

	ctxMain, cancelMain := context.WithCancel(context.Background())

	grpcClient, err := grpcclient.NewClient(grpcServerAddress, zapLogger)
	if err != nil {
		zapLogger.Fatal("grpcclient.NewClient", zap.Error(err))
	}

	client := tuiclient.NewApp(grpcClient, zapLogger)

	if err := client.Run(ctxMain); err != nil {
		zapLogger.Fatal("client.Run", zap.Error(err))
	}
	cancelMain()
	//grpcconn, err := grpc.Dial(hostPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//client := proto.NewServerClient(grpcconn)
	//for {
	//	msg, err := client.Hello(context.Background(), &proto.HelloMsg{
	//		Msg: "Сообщение от клиента",
	//	})
	//	if err != nil {
	//		status, _ := status2.FromError(err)
	//		fmt.Printf("%+v\n", status.Message())
	//		time.Sleep(5 * time.Second)
	//		continue
	//	}
	//
	//	fmt.Println("Получено сообщение от сервера " + msg.String())
	//	time.Sleep(5 * time.Second)
	//}

}
