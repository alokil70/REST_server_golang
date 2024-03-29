package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"test_RESTserver_01/internal/author"
	authorRepo "test_RESTserver_01/internal/author/db"
	"test_RESTserver_01/internal/config"
	"test_RESTserver_01/internal/user"
	"test_RESTserver_01/pkg/client/postgresql"
	"test_RESTserver_01/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger()

	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("error client: %v", err)
	}

	repository := authorRepo.NewRepository(postgreSQLClient, logger)

	// newAuthor := author.Author{
	// 	Name: "Маяковский",
	// }
	// err = repository.Create(context.TODO(), &newAuthor)
	// if err != nil {
	// 	logger.Fatalf("error create author: %v", err)
	// }

	one, err := repository.FindOne(context.TODO(), "f8a6a73d-b720-4249-a59f-590d254cc185")
	if err != nil {
		logger.Fatalf("error findone: %v", err)
	}
	logger.Infof("%v", one)

	all, err := repository.FindAll(context.TODO())
	if err != nil {
		logger.Fatalf("error findall: %v", err)
	}

	for _, a := range all {
		logger.Infof("%v", a)
	}

	logger.Info("register author handler")
	authorHandler := author.NewHandler(repository, logger)
	authorHandler.Register(router)

	logger.Info("register user handler")
	userHandler := user.NewHandler(logger)
	userHandler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenError error

	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info(appDir) // tett
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenError = net.Listen("unix", socketPath)
		logger.Infof("server listen unix socket %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server listen ip %s port %s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenError != nil {
		logger.Fatal(listenError)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
