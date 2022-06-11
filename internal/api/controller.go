package api

import (
	"net/http"
	"regexp"

	"github.com/deweppro/go-logger"

	"github.com/dewep-online/uri-one/pkg/utils"
	"github.com/deweppro/go-badges"
	"github.com/deweppro/go-http/pkg/routes"
)

//Index controller
func (v *API) Index(w http.ResponseWriter, r *http.Request) {
	filename := r.RequestURI
	switch filename {
	case "", "/":
		filename = "/index.html"
	default:
	}

	if err := v.cache.ResponseWrite(w, filename); err != nil {
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
	"faq.html":   {},
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

var (
	rex = regexp.MustCompile(`^\/badge\/([a-z]+)\/([^\/]+)\/([^\/]+)\/image.svg$`)

	colors = map[string]badges.Color{
		"blue":   badges.ColorPrimary,
		"red":    badges.ColorDanger,
		"yellow": badges.ColorWarning,
		"green":  badges.ColorSuccess,
		"light":  badges.ColorLight,
	}
)

func (v *API) BadgesMiddleware() func(c routes.CtrlFunc) routes.CtrlFunc {
	return func(c routes.CtrlFunc) routes.CtrlFunc {
		return func(w http.ResponseWriter, r *http.Request) {

			values := rex.FindStringSubmatch(r.URL.Path)
			if len(values) == 4 {
				c, ok := colors[values[1]]
				if !ok {
					c = badges.ColorLight
				}
				err := v.badges.WriteResponse(w, c, values[2], values[3])
				if err != nil {
					logger.WithFields(logger.Fields{"err": err}).Errorf("Badges generate")
				}
				return
			}

			c(w, r)
		}
	}
}
