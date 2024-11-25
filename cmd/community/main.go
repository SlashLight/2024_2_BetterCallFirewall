package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/2024_2_BetterCallFirewall/internal/app/community"
	"github.com/2024_2_BetterCallFirewall/internal/config"
)

func main() {
	confPath := flag.String("c", ".env", "path to config file")
	flag.Parse()

	cfg, err := config.GetConfig(*confPath)
	if err != nil {
		panic(err)
	}

	httpServer, grpcServer, err := community.GetServers(cfg)
	if err != nil {
		panic(err)
	}

	go func() {
		l, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.COMMUNITYGRPC.Port))
		if err != nil {
			panic(err)
		}

		log.Printf("Listening on :%s with protocol gRPC", cfg.COMMUNITYGRPC.Port)
		if err := grpcServer.Serve(l); err != nil {
			panic(err)
		}
	}()

	log.Printf("Starting server on posrt %s", cfg.COMMUNITY.Port)
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}
}
