package myFavoriteArtistService

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Invoke(RegisterMyFavoriteArtistformEndPoint),
	fx.Provide(NewMyFavoriteArtistService),
)
