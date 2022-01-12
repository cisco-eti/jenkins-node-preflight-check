package server

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/models"
	"wwwin-github.cisco.com/eti/sre-go-helloworld/pkg/utils"
)

// Get godoc
// @Summary Get Home
// @Description get a response from home endpoint
// @ID get-home
// @Tags Home
// @Error 401
// @Router / [get]
func (s *Server) RootHandler(w http.ResponseWriter, _ *http.Request) {
	s.log.Info("/ request received")

	_ = utils.OKResponse(w, "root")
	return
}

// Get godoc
// @Summary Get Ping
// @Description get helloworld status
// @Produce json
// @Success 200 {object} models.PingResponse
// @Router /ping [get]
func (s *Server) PingHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/ping request received")

	hostname := os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "HOSTNAME env missing"
	}

	// Pure sample data, value need to be changed
	myService := models.Service{
		ServiceName:  "HelloWorld",
		ServiceType:  "OPTIONAL",
		ServiceState: "online",
		Message:      "Healthy",
		ServiceInstance: models.ServiceInstance{
			InstanceID: hostname,
			Host:       "172.18.231.5",
			Port:       21455,
		},
		LastUpdated:    "2020-10-20T08:42:07.290Z",
		BaseURL:        "http://helloworld.int.scratch-aws-1.prod.eticloud.io/",
		DurationPretty: "91ms",
		Duration:       91350005,
		UpstreamServices: []models.Service{
			{
				ServiceName:  "Postgres",
				ServiceType:  "REQUIRED",
				ServiceState: "online",
				ServiceInstance: models.ServiceInstance{
					InstanceID: "e3c16830-2c65-c6c8-68ab-30d728d6179e[9]",
					Host:       "172.18.244.25",
					Port:       5432,
				},
				Message:          "PostgresDataSource is online",
				LastUpdated:      "2021-01-06T13:43:42.984Z",
				DurationPretty:   "3ms",
				Duration:         3631684,
				UpstreamServices: []models.Service{},
				DefaultCharset:   "UTF-8",
			},
		},
		DefaultCharset: "UTF-8",
	}
	_ = utils.OKResponse(w, myService)
	return
}

// Get godoc
// @Summary Get API Docs
// @Description get Swagger API documentation
// @Produce Yaml
// @Success 200
// @Router /docs [get]
func (s *Server) DocsHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/docs request received")

	filePath, err := filepath.Abs("./docs/openapi.yaml")
	if err != nil {
		utils.ServerErrorResponse(w, err.Error())
		return
	}

	s.log.Info("filePath: %s", filePath)
	http.ServeFile(w, r, filePath)
}

// Get godoc
// @Summary Get Prometheus Metrics
// @Description get helloworld status
// @Produce Yaml
// @Success 200
// @Router /metrics [get]
func (s *Server) MetricsHandler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/metrics request received")

	promhttp.Handler().ServeHTTP(w, r)
}

// Get godoc
// @Summary Get S3
// @Description get S3 objects
// @Produce json
// @Success 200
// @Router /s3 [get]
func (s *Server) S3Handler(w http.ResponseWriter, r *http.Request) {
	s.log.Info("/s3 request received")

	var (
		bucket   string = "eks-sre-1-test"
		key      string = "sre-go-helloworld"
		filename string = "s3_object.txt"
	)

	// All clients require a Session. The Session provides the client with
	// shared configuration such as region, endpoint, and credentials. A
	// Session should be shared where possible to take advantage of
	// configuration and credential caching. See the session package for
	// more information.
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-2"),
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	timeout, _ := time.ParseDuration("1m")
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
		defer cancelFn()
	}
	wd, err := os.Getwd()
	if err != nil {
		s.log.Error(err.Error())
	}
	s.log.Info("current working directory: %s", wd)
	files, err := ioutil.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}

	// s.log.Info("Printing files in '/'...")
	// for _, file := range files {
	// 	s.log.Info("Name: %s, IsDir: %t", file.Name(), file.IsDir())
	// }

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	f, err := os.Open("/" + filename)
	if err != nil {
		s.log.Error("failed to open file %q, %v", filename, err)
		return
	}
	result, err := uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   f,
	})
	if err != nil {
		s.log.Error("failed to upload file, %v", err)
		return
	}
	s.log.Info("file uploaded to, %s\n", result.Location)
	filePath, err := filepath.Abs("./web/s3.html")
	if err != nil {
		utils.ServerErrorResponse(w, err.Error())
		return
	}

	s.log.Info("filePath: %s", filePath)
	http.ServeFile(w, r, filePath)
}
