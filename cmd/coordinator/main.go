package main

import (
	"flag"

	"github.com/TheVidz/DistributedTaskScheduler/pkg/common"
	"github.com/TheVidz/DistributedTaskScheduler/pkg/coordinator"
)

var (
	coordinatorPort = flag.String("coordinator_port", ":8080", "Port on which the Coordinator serves requests.")
)

func main() {
	flag.Parse()
	dbConnectionString := common.GetDBConnectionString()
	coordinator := coordinator.NewServer(*coordinatorPort, dbConnectionString)
	coordinator.Start()
}
