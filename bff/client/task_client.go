package client

import (
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/moonorange/gomicroservice/protogo/gen/genconnect"
	"github.com/sirupsen/logrus"
)

var (
	queryClient   genconnect.TaskServiceClient
	commandClient genconnect.TaskServiceClient
)

func NewQueryServiceClient() genconnect.TaskServiceClient {
	inventoryHost := os.Getenv("INVENTORY_SERVICE_HOST")
	logrus.Info("inventoryHost: ", inventoryHost)
	if inventoryHost == "" {
		logrus.Fatal("empty INVENTORY_SERVICE_HOST")
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	if queryClient != nil {
		return queryClient
	}
	queryClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		"http://"+inventoryHost,
		connect.WithGRPC(),
	)

	return queryClient
}

func NewCommandServiceClient() genconnect.TaskServiceClient {
	commandHost := os.Getenv("PAYMENT_SERVICE_HOST")
	logrus.Info("commandHost: ", commandHost)
	if commandHost == "" {
		logrus.Fatal("empty PAYMENT_SERVICE_HOST")
	}
	if commandClient != nil {
		return commandClient
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	commandClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		"http://"+commandHost,
		connect.WithGRPC(),
	)

	return commandClient
}
