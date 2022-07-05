package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"github.com/url-shortener/models"
	"github.com/url-shortener/types"
)

var log = logrus.New()

func init() {
	log.Formatter = new(logrus.JSONFormatter)
}

func Routes(urlModel models.UrlIface) chi.Router {
	router := chi.NewRouter()

	router.Use(chiMiddleware.AllowContentType("application/json"))

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		Create(urlModel, w, r)
	})

	return router
}

func RedirectRoutes(urlModel models.UrlIface) chi.Router {
	router := chi.NewRouter()

	router.Route("/{urlID}", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return UrlCtx(urlModel, next)
		})
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			RedirectUrl(urlModel, w, r)
		})
	})
	return router
}

func Create(urlModel models.UrlIface, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	parent := opentracing.GlobalTracer().StartSpan("POST /urls")

	defer parent.Finish()

	b, err := ioutil.ReadAll(r.Body)

	defer r.Body.Close()
	if err != nil {
		log.Error(err)
		return
	}

	var url types.Url
	err = json.Unmarshal(b, &url)
	if err != nil {
		log.Error(err)
		return
	}

	url.ShortenUrl = ""
	url, err = urlModel.CreateShortUrl(url)

	if err != nil {
		log.Error(err)
		return
	}

	output, err := json.Marshal(url)
	if err != nil {
		log.Error(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(output)
}

func UrlCtx(urlModel models.UrlIface, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		shorten_url := "https://dpurl.dev/" + chi.URLParam(r, "urlID")
		url, err := urlModel.GetLongUrl(shorten_url)
		if err != nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "url", url)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func RedirectUrl(urlModel models.UrlIface, w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	url, ok := ctx.Value("url").(types.Url)
	if !ok {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	http.Redirect(w, r, url.Url, http.StatusSeeOther)
}
