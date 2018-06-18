package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

func main() {
	cfg := client.Config{
		Endpoints: []string{"http://192.168.1.61:2379", "http://192.168.1.62:2379", "http://192.168.1.63:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	kapi := client.NewKeysAPI(c)

	watcherOptions := client.WatcherOptions{Recursive: true}
	watcher := kapi.Watcher("/foo", &watcherOptions)

	go onChanged(&watcher)

	fmt.Print("Press 'Enter' to exit...")
	for {
		time.Sleep(200)
	}
}

func interactWithEtcd(kapi client.KeysAPI) {
	// set "/foo" key with "bar" value
	log.Print("Setting '/foo' key with 'bar' value")
	resp, err := kapi.Set(context.Background(), "/foo", "bar", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Set is done. Metadata is %q\n", resp)
	}
	// get "/foo" key's value
	log.Print("Getting '/foo' key value")
	resp, err = kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		// print common key info
		log.Printf("Get is done. Metadata is %q\n", resp)
		// print value
		log.Printf("%q key has %q value\n", resp.Node.Key, resp.Node.Value)
	}
}

func onChanged(watcher *client.Watcher) {
	for {
		resp, err := (*watcher).Next(context.Background())
		if err != nil {
			log.Println("Failed to get next event on watcher : ", err)
			break
		}

		log.Println(fmt.Sprintf("Action: [%s] on Node [%s] (Value: %s).", resp.Action, resp.Node.Key, resp.Node.Value))
		if resp.Action == "expire" {
			log.Println(fmt.Sprintf("Node [%s] (Value: %s) has expired.", resp.Node.Key, resp.Node.Value))
			continue
		}

		if resp.Action == "set" {
			log.Println(fmt.Sprintf("Set node [%s] (Value: %s).", resp.Node.Key, resp.Node.Value))
			continue
		}

		if resp.Action == "update" {
			log.Println(fmt.Sprintf("Node [%s] updated, new value: %s", resp.Node.Key, resp.Node.Value))
			continue
		}

		if resp.Action == "delete" {
			log.Println(fmt.Sprintf("Node [%s] (Value: %s) has been deleted.", resp.Node.Key, resp.Node.Value))
			continue
		}
	}
}
