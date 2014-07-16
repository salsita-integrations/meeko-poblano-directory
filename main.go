// Copyright (c) 2014 The meeko-poblano-directory AUTHORS
//
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/salsita-integrations/meeko-poblano-directory/methods"

	"github.com/meeko/go-meeko/agent"
)

func main() {
	// Get the Logging service.
	var log = agent.Logging()

	// Parse the environment.
	var (
		apiBaseURL = os.Getenv("API_BASE_URL")
		apiToken   = os.Getenv("API_TOKEN")
		rpcToken   = os.Getenv("RPC_TOKEN")
	)
	if apiBaseURL == "" {
		log.Critical("API_BASE_URL is not set")
		agent.Terminate(2)
	}
	if apiToken == "" {
		log.Critical("API_TOKEN is not set")
		agent.Terminate(2)
	}
	if rpcToken == "" {
		log.Critical("RPC_TOKEN is not set")
		agent.Terminate(2)
	}

	// Get the RPC service.
	var rpc = agent.RPC()

	// Export all available methods.
	client := poblano.NewClient(apiBaseURL, apiToken)
	if err := methods.New(log, client, rpcToken).Export(rpc); err != nil {
		log.Critical(err)
		agent.Terminate(1)
	}

	// Wait for the termination signal.
	<-agent.Stopped()
	agent.Terminate(0)
}
