package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/multicharts/go-zookeeper/zk"
)

func main() {
	log.Printf("Starting test cluster")
	cluster, err := zk.StartTestCluster(3, nil, os.Stderr)
	if err != nil {
		panic(err)
	}
	defer cluster.Stop()

	hosts := make([]string, 0, len(cluster.Servers))
	for _, s := range cluster.Servers {
		hosts = append(hosts, fmt.Sprintf("127.0.0.1:%d", s.Port))
	}

	log.Printf("Establishing connection")
	conn, events, err := zk.Connect(hosts, time.Second)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	log.Printf("Waiting StateHasSession event")
	select {
	case event := <-events:
		log.Printf("Received event from zk: %v", event)
		if event.State == zk.StateHasSession {
			log.Printf("Session established")
		}
	case <-time.After(10 * time.Second):
		log.Printf("Timeout expired, closing connection")
		return
	}

	log.Printf("Creating node /myznode")
	var path string
	if path, err = conn.Create("/myznode", []byte{}, 0, zk.WorldACL(zk.PermAll)); err != nil {
		panic(err)
	}

	log.Printf("Checking if node %s exists", path)
	if yes, _, err := conn.Exists(path); !yes || err != nil {
		panic(err)
	}

	log.Printf("Setting node %s data", path)
	if _, err = conn.Set(path, []byte("hello"), -1); err != nil {
		panic(err)
	}

	log.Printf("Getting node %s data", path)
	if data, _, err := conn.Get(path); err != nil || string(data) != "hello" {
		panic(err)
	} else {
		log.Printf("Node data: %v", string(data))
	}

	log.Printf("Deleting node %s", path)
	if err = conn.Delete(path, -1); err != nil {
		panic(err)
	}
}
