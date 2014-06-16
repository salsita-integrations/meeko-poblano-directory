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
	// Instantiate Logging service.
	var log = agent.Logging()

	// Parse the environment.
	var (
		apiBaseURL = os.Getenv("API_BASE_URL")
		apiToken   = os.Getenv("API_TOKEN")
	)
	if apiBaseURL == "" {
		log.Critical("API_BASE_URL is not set")
		log.Close()
		os.Exit(1)
	}
	if apiToken == "" {
		log.Critical("API_TOKEN is not set")
		log.Close()
		os.Exit(1)
	}

	// Instantiate RPC service.
	var rpc = agent.RPC()

	// Export all available methods.
	client := poblano.NewClient(apiBaseURL, apiToken)
	if err := methods.New(log, client).Export(rpc); err != nil {
		log.Critical(err)
		log.Close()
		rpc.Close()
		os.Exit(1)
	}

	// Wait for the termination signal.
	<-agent.Stopped()
	log.Close()
	rpc.Close()
}
