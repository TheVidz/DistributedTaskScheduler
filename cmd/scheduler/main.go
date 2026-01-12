package main

import (
	"flag"
	"log"

	"github.com/TheVidz/DistributedTaskScheduler/pkg/common"
	"github.com/TheVidz/DistributedTaskScheduler/pkg/scheduler"
)

var (
	schedulerPort = flag.String("scheduler_port", ":8081", "Port on which the Scheduler serves requests.")
)

func main() {
	dbConnectionString := common.GetDBConnectionString()
	schedulerServer := scheduler.NewServer(*schedulerPort, dbConnectionString)
	err := schedulerServer.Start()
	if err != nil {
		log.Fatalf("Error while starting server: %+v", err)
	}
}
