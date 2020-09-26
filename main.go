package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Setup logging
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func handleEcho(res http.ResponseWriter, req *http.Request) {
	log.Info("Handling request for /echo")
	reqBody := new(bytes.Buffer)
	reqBody.ReadFrom(req.Body)
	url, err := url.Parse(req.URL.RequestURI())

	if err != nil {
		res.WriteHeader(503)
		return
	}

	params := url.Query()

	// Set latency
	latency := params.Get("latency")
	if latency != "" {
		latentDuration, err := time.ParseDuration(latency)
		if err != nil {
			res.WriteHeader(400)
			return
		}
		log.WithField("duration", latentDuration).Debug("Adding latency to response")
		time.Sleep(latentDuration)
	}

	// Set status code
	status := params.Get("status")
	if status != "" {
		statusCode, err := strconv.Atoi(status)
		if err != nil {
			res.WriteHeader(400)
			return
		}
		log.WithField("statusCode", statusCode).Debug("Setting HTTP reponse code")
		res.WriteHeader(statusCode)
	}

	res.Write(reqBody.Bytes())
}

func handleOk(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	res.Write([]byte("chirp"))
}

func main() {
	port := getEnvOrDefault("HTTP_PORT", "8080")
	log.Info("Starting Mockingbird on port ", port)
	http.HandleFunc("/echo", handleEcho)
	http.HandleFunc("/", handleOk)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func getEnvOrDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
