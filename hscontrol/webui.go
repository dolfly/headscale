package hscontrol

import (
	_ "embed"
	"net/http"

	"github.com/gnue/httpfs/zipfs"
	"github.com/gorilla/mux"
)

//go:embed templates/headscale-ui.zip
var zipdata []byte

func web(router *mux.Router) {
	router.PathPrefix("/web").Handler(func() http.Handler {
		zfs, err := zipfs.New(zipdata, &zipfs.Options{Prefix: "web"})
		if err != nil {
			return nil
		}

		return http.StripPrefix("/web", http.FileServer(zfs))
	}())
}
