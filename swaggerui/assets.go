package swaggerui

import (
	"net/http"
	"path/filepath"

	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/k0kubun/pp"
	"github.com/labstack/echo"
)

func getAssetFS() *assetfs.AssetFS {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
		Prefix:    "",
	}
}

func Handle() echo.HandlerFunc {
	return echo.WrapHandler(http.FileServer(getAssetFS()))
}

func Routes(prefix string, e *echo.Echo) error {
	mimeTypes := map[string]string{
		".js":   echo.MIMEApplicationJavaScript,
		".json": echo.MIMEApplicationJSONCharsetUTF8,
		".html": echo.MIMETextHTMLCharsetUTF8,
		".txt":  echo.MIMETextPlainCharsetUTF8,
		".png":  "image/png",
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".css":  "text/css",
	}
	for _, k := range AssetNames() {
		bts, err := Asset(k)
		if err != nil {
			continue
		}
		k = prefix + "/" + k
		mimeType, ok := mimeTypes[filepath.Ext(k)]
		if !ok {
			mimeType = "application/octet-stream"
		}
		pp.Println(k, "               ", mimeType)
		e.GET(k, func(ctx echo.Context) error {
			return ctx.Blob(http.StatusOK, mimeType, bts)
		})
	}
	return nil
}
