package cmd

import (
	"fmt"
	"net/http"
	"time"
	
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/log"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.app_name}}/internal/routers"
	"github.com/gin-gonic/gin"
)



func Setup() error {

	err := HttpServer()
	if err != nil {
		return err
	}

	return err
}


func HttpServer() error {

	/*
	 * config
	 */

	log.Debug("Starting server...")
	mode := config.Config().GetString("http.mode")
	read_timeout := config.Config().GetDuration("http.read_timeout")
	write_timeout := config.Config().GetDuration("http.write_timeout")
	port := config.Config().GetInt("http.port")

	gin.SetMode(mode)

	router := routers.InitRouter()

	readTimeout := read_timeout * time.Second
	writeTimeout := write_timeout * time.Second
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	err := s.ListenAndServe()
	return err
}
