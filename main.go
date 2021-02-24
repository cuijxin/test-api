package main

import (
	"errors"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/cuijxin/mysql-cluster-presslabs/pkg/config"
	"github.com/cuijxin/mysql-cluster-presslabs/pkg/kube"
	"github.com/cuijxin/mysql-cluster-presslabs/routers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg        = pflag.StringP("config", "c", "", "mysql-cluster apiserver config file path.")
	kubeconfig = pflag.StringP("kubeconfig", "k", "", "optional absolute path to the kubeconfig file.")
)

func main() {
	pflag.Parse()

	// init config
	if err := config.Init(*cfg); err != nil {
		log.Fatal(err)
	}

	if err := kube.Init(*kubeconfig); err != nil {
		log.Fatal(err)
	}

	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	routers.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Info("The router has no response, or it might took too long to start up. %v", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Printf("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Printf(http.ListenAndServe(viper.GetString("addr"), g).Error())

}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("url") + "/healthcheck/ping")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Printf("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("cannot connect to the router")
}
