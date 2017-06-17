package swaggerui

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
)

func AssetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "",
	}
}

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(http.FileServer(AssetFS()))
}
