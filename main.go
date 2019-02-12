package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"./models"
	"./routes"

	"github.com/gorilla/handlers"
)

func main() {
	logs := &models.Logger{}
	logs.InitLogging("GoGraphQL", os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr, os.Stdout)

	api := routes.NewRouter() // create routes
	defer api.Database.Close()

	// These two lines are important in order to allow access from the front-end side to the methods
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
	}

	srv := &http.Server{
		Addr:         ":443",
		Handler:      handlers.CORS(allowedOrigins, allowedMethods)(api.Router),
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	// Launch server with CORS validations
	logs.Fatal.Println(srv.ListenAndServeTLS("certs/server.rsa.crt", "certs/server.rsa.key"))
}
