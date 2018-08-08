package main

import (
	"log"

	"github.com/waltzofpearls/otto/docker"
	"github.com/waltzofpearls/otto/logger"
	"go.uber.org/zap"
)

func main() {
	if err := logger.Init(); err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	engine, err := docker.New()
	if err != nil {
		logger.Fatal("failed to create docker engine",
			zap.Error(err))
	}
	log.Printf("%#v", engine)
	// output, outdated, err := goDep()
	// if err != nil {
	// 	log.Fatalf("%s%s", output, err)
	// }
	// fmt.Println("There are", outdated, "outdated deps")
	// if outdated > 0 {
	// 	fmt.Println(string(output))
	// 	os.Exit(1)
	// }
}
