package main

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/kudobuilder/shim/shim-controller/pkg/client"
	"github.com/kudobuilder/shim/shim-controller/pkg/controller"
)

func main() {
	log.Infof("bootstrapping KUDO Shim Controller...")
	clientSet, err := client.GetKubeClient()
	if err != nil {
		log.Fatalf("failed to get kube client: %v", err)
		return
	}
	cont := controller.NewController(clientSet)
	cont.Run(context.Background())
}

func init() {
	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	log.SetFormatter(customFormatter)
	customFormatter.FullTimestamp = true
}
