package jaegerui

import (
	"net/http"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rai-project/web/jaegerui/static"
)

func AssetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "jaeger",
	}
}

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(http.FileServer(AssetFS()))
}
func Handle2() echo.HandlerFunc {
	return echo.WrapHandler(static.Handler)
}

func Routes(e *echo.Echo, prefix string) error {

	assetGroup := e.Group(prefix,
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
	)
	assetGroup.Any("/*", Handle2())
	return nil
}

func Routes2(e *echo.Echo, prefix string) error {

	assetGroup := e.Group(prefix,
		middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}),
	)

	index := func(c echo.Context) error {
		html, err := jaegerIndexHtmlBytes()
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, string(html))
	}

	favicon := func(c echo.Context) error {
		ico, err := jaegerFaviconIcoBytes()
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "image/x-icon", ico)
	}
	assetManifest := func(c echo.Context) error {
		js, err := jaegerAssetManifestJsonBytes()
		if err != nil {
			return err
		}
		return c.Blob(http.StatusOK, "application/javascript", js)
	}

	assetGroup.GET("/", index)
	assetGroup.GET("/*", index)
	assetGroup.GET("/index.html", index)
	assetGroup.GET("/favicon.ico", favicon)
	assetGroup.HEAD("/favicon.ico", favicon)
	assetGroup.GET("/asset-manifest.json", assetManifest)
	assetGroup.GET("/vendor/*", Handle())
	assetGroup.GET("/static/*", Handle())
	assetGroup.Any("/*", Handle())
	return nil
}
