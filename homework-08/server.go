package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"strings"
	"time"
)

type ServerConfig struct {
	URL 				string `yaml:"url"`
	ListenHTTP 			string `yaml:"listen-http"`
	TemplateFilesDir 	string `yaml:"template-files-dir"`
	StaticFilesDir 		string `yaml:"static-files-dir"`
	StaticFilesURL 		string `yaml:"static-files-url"`
	MongoURL 			string `yaml:"mongodb-url"`
	MongoDATABASE 		string `yaml:"mongodb-database"`
	PinCodeForUpload 	string `yaml:"pin-for-upload"`
	PinCodeForDelete 	string `yaml:"pin-for-delete"`
}

// Server - Server instance
type Server struct {
	config   ServerConfig
	lg       *logrus.Logger
	ctx      context.Context
	mux      *chi.Mux
	server   *http.Server
	db		 *mongo.Database
}

// ServErr - Server error response
type ServErr struct {
	Code     int         `json:"code"`
	Err      string      `json:"err"`
	Desc     string      `json:"desc"`
	Internal interface{} `json:"internal"`
}

// NewServer - Create a new server instance
func NewServer(ctx context.Context, lg *logrus.Logger) *Server {
	serv := &Server{
		lg:  lg,
		ctx: ctx,
		mux: chi.NewMux(),
	}
	return serv
}


// Start - Starts the server
func (serv *Server) Start() *Server {
	// DATABASE
	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(serv.config.MongoURL))
	if err != nil {
		serv.lg.Errorf("database: %s", err)
	}
	serv.db = client.Database(serv.config.MongoDATABASE)


	// HTTP
	serv.server = &http.Server{
		Addr:    serv.config.ListenHTTP,
		Handler: serv.mux,
	}
	serv.configureRoutes()
	go func() {
		serv.lg.Info("starting server")
		if err := serv.server.ListenAndServe(); err != http.ErrServerClosed {
			serv.lg.Errorf("server: %s", err)
		}
	}()

	return serv
}

// Stop - Stops the server
func (serv *Server) Stop() error {
	serv.lg.Info("stopping server")
	return serv.server.Shutdown(serv.ctx)
}

// SendErr - Sends error
func (serv *Server) SendErr(w http.ResponseWriter, err error, code int, obj ...interface{}) {
	serv.lg.WithField("data", obj).WithError(err).Error("server error")
	w.WriteHeader(code)
	errModel := ServErr{
		Code:     code,
		Err:      err.Error(),
		Desc:     "server error",
		Internal: obj,
	}
	data, _ := json.Marshal(errModel)
	w.Write(data)
}

// SetConfig - Setting config
func (serv *Server) SetConfig(conf ServerConfig) {
	serv.config = conf
}

// SendInternalErr - Sends internal server error
func (serv *Server) SendInternalErr(w http.ResponseWriter, err error, obj ...interface{}) {
	serv.SendErr(w, err, http.StatusInternalServerError, obj)
}

// SendJSON - Wrapper for structure sending
func (serv *Server) SendJSON(w http.ResponseWriter, obj interface{}) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	w.Write(data)
	return nil
}

// FileServer conveniently sets up a http.FileServer handler to serve static files from a http.FileSystem
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func GetIP(r *http.Request) string {
	var ip string
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		ip = forwarded
	}
	ip = r.RemoteAddr

	if strings.Contains(ip,":") {
		arr := strings.Split(ip,":")
		return arr[0]
	}
	return ip
}