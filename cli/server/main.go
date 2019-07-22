package main

import (
	"flag"
	grpcsetup "github.com/hughselway/m-borderchange/internal/server/grpc"
	"github.com/rs/zerolog/log"
)

func main() {
	var addressPtr = flag.String("address", ":50052", "address to connect with m-borderchange microservice")
	flag.Parse()

	g := grpcsetup.NewServer(*addressPtr)

	err := g.ListenAndServe()

	if err != nil {
		log.Fatal().Err(err).Msg("failed to start grpc server for m-borderchange")
	}
}
