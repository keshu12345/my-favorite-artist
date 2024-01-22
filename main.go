package main

import (
	"flag"
	"log"
	"os"

	"github.com/keshu12345/my-favorite-artist/config"
	"github.com/keshu12345/my-favorite-artist/myFavoriteArtistService"
	"github.com/keshu12345/my-favorite-artist/server/router"
	"go.uber.org/fx"
)

var configDirPath = flag.String("config", "", "path for config dir")

func main() {

	flag.Parse()
	log.New(os.Stdout, "", 0)
	app := fx.New(
		config.NewFxModule(*configDirPath, ""),
		router.Module,
		myFavoriteArtistService.Module,
	)
	app.Run()
}
