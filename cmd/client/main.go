package main

import (
	"fmt"
)

var (
	buildVersion string
	buildDate    string
)

func main() {
	fmt.Println(buildVersion)
	fmt.Println(buildDate)
	fmt.Println("----")
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
