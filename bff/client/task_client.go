package client

import (
	"net/http"
	"os"

	"connectrpc.com/connect"
	"github.com/moonorange/gomicroservice/protogo/gen/genconnect"
	"github.com/sirupsen/logrus"
)

var (
	inventoryClient genconnect.TaskServiceClient
	paymentClient   genconnect.TaskServiceClient
)

func NewQueryServiceClient() genconnect.TaskServiceClient {
	inventoryHost := os.Getenv("INVENTORY_SERVICE_HOST")
	logrus.Info("inventoryHost: ", inventoryHost)
	if inventoryHost == "" {
		logrus.Fatal("empty INVENTORY_SERVICE_HOST")
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	if inventoryClient != nil {
		return inventoryClient
	}
	inventoryClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		"http://"+inventoryHost,
		connect.WithGRPC(),
	)

	return inventoryClient
}

func NewCommandServiceClient() genconnect.TaskServiceClient {
	paymentHost := os.Getenv("PAYMENT_SERVICE_HOST")
	logrus.Info("paymentHost: ", paymentHost)
	if paymentHost == "" {
		logrus.Fatal("empty PAYMENT_SERVICE_HOST")
	}
	if paymentClient != nil {
		return paymentClient
	}
	// Set up a connection to the server.
	// Create a gRPC client using the connect.WithGRPC() option
	paymentClient = genconnect.NewTaskServiceClient(
		http.DefaultClient,
		"http://"+paymentHost,
		connect.WithGRPC(),
	)

	return paymentClient
}
