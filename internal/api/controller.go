package api

import (
	"net/http"

	"github.com/dewep-online/uri-one/pkg/utils"
	"github.com/deweppro/go-http/web/routes"
)

//Index controller
func (v *API) Index(w http.ResponseWriter, r *http.Request) {
	filename := r.RequestURI
	switch filename {
	case "", "/":
		filename = "/index.html"
	default:
	}

	if err := v.cache.Write(filename, w); err != nil {
		v.log.Errorf("static response: %s", err.Error())
	}
}

var (
	invalidRequest = []byte("Invalid request")
)

func (v *API) Add(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if !utils.IsValidUrl(query) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(403)
		w.Write(invalidRequest) //nolint: errcheck
		return
	}
	id, err := v.db.SetUrl(query)
	if err != nil {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(403)
		w.Write(invalidRequest) //nolint: errcheck
	}
	hash := "https://" + r.Host + "/" + v.enc.Marshal(id)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte(hash)) //nolint: errcheck
}

var skip = map[string]struct{}{
	"":           {},
	"index.html": {},
	"404.html":   {},
	"+":          {},
}

func (v *API) DetectLinkMiddleware() func(c routes.CtrlFunc) routes.CtrlFunc {
	return func(c routes.CtrlFunc) routes.CtrlFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			code := r.URL.Path[1:]
			if _, ok := skip[code]; !ok {
				id := v.enc.Unmarshal(code)
				if data, err := v.db.GetUrl(id); err == nil {
					w.Header().Set("Location", data)
				} else {
					w.Header().Set("Location", "/404.html")
				}
				w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
				w.WriteHeader(301)
				return
			}
			c(w, r)
		}
	}
}
