package main

import (
	"log"

	"github.com/hhyouke/server/cmd"
)

// a http server
// server shall hold a bunch of REST endpoints
// also a websocket serve: clazz chat, course-group chat and notifications
// course-material type: plain text/images, premade audio, premade video and live streaming

// must-have-components: cmd, log, config, db, ws, http mux, auth

func main() {
	rootCmd := cmd.NewRootCmd()
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("startup failed, %v\n", err.Error())
	}
}
