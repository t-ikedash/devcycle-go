package main

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env"

	devcycle "github.com/devcyclehq/go-server-sdk/v2"
)

type devCycle struct {
	SDKKey    string `env:"DEVCYCLE_SERVER_SDK_KEY"`
	SampleKey string `env:"SAMPLE_KEY"`
}

func main() {
	cfg := &devCycle{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}
	sdkKey := cfg.SDKKey

	options := devcycle.Options{
		EnableEdgeDB:                 false,
		EnableCloudBucketing:         false,
		EventFlushIntervalMS:         30 * time.Second,
		ConfigPollingIntervalMS:      1 * time.Minute,
		RequestTimeout:               30 * time.Second,
		DisableAutomaticEventLogging: false,
		DisableCustomEventLogging:    false,
	}
	devcycleClient, err := devcycle.NewClient(sdkKey, &options)
	if err != nil {
		log.Fatalf("Error initializing DevCycle client: %v", err)
	}
	user := devcycle.User{
		UserId: "example_user_id",
	}
	key := cfg.SampleKey
	variableValue, err := devcycleClient.VariableValue(user, key, false)
	fmt.Printf("Variable value: %v\n", variableValue)
	if variableValue.(bool) {
		// Put feature code here, or launch feature from here
		fmt.Println("Variable" + key + " is enabled")
	} else {
		// Put feature code here, or launch feature from here
		fmt.Println("Variable " + key + " is disabled")
	}
}
