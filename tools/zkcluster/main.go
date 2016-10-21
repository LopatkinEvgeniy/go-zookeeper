package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/samuel/go-zookeeper/zk"
)

func main() {
	waitsig := make(chan os.Signal, 1)
	signal.Notify(waitsig, syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	cluster, err := zk.StartTestCluster(3, os.Stderr, os.Stderr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer cluster.Stop()

	log.Print("Successfully started test cluster")

	hosts := make([]map[string]interface{}, 0, len(cluster.Servers))
	for _, s := range cluster.Servers {
		hosts = append(hosts, map[string]interface{}{
			"host": "127.0.0.1",
			"port": s.Port,
		})
	}

	b, err := json.MarshalIndent(hosts, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)

	<-waitsig
}
