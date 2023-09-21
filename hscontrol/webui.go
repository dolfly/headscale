package hscontrol

import (
	_ "embed"
	"net/http"

	"github.com/gnue/httpfs/zipfs"
	"github.com/gorilla/mux"
)

//go:embed templates/headscale-ui-0.1.4.zip
var zipdata []byte

func web(router *mux.Router) {
	router.PathPrefix("/web/").Handler(func() http.Handler {
		zfs, err := zipfs.New(zipdata, &zipfs.Options{Prefix: "headscale-ui"})
		if err != nil {
			return nil
		}
		return http.StripPrefix("/web/", http.FileServer(zfs))
	}())
}
