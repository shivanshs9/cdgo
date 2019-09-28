package daemon

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shivanshs9/cdgo/pkg/app"
	log "github.com/sirupsen/logrus"
	"github.com/takama/daemon"
)

const (
	serviceName        = "cdgo"
	serviceDescription = "Decentralized P2P Daemon process"
)

type Daemon struct {
	daemon.Daemon
}

var Service *Daemon

func init() {
	srv, err := daemon.New(serviceName, serviceDescription)
	if err != nil {
		log.Error("Error: ", err)
		os.Exit(1)
	}
	Service = &Daemon{srv}
}

func (service *Daemon) StartDaemon(app *app.App, isBlocking bool) (result string, err error) {
	cfg, err := InitConfig()
	if err != nil {
		return "Error", err
	}

	if !isBlocking {
		result, err = service.Start()
		if err != nil {
			return "Error", err
		}
		return "Started Daemon", nil
	}

	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	var listener net.Listener
	if cfg.UseSocketFile {
		listener, err = net.Listen("unix", cfg.UnixSocket)
		log.Info("Listening on socket ", cfg.UnixSocket)
	} else {
		listener, err = net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
		log.Info("Listening on port ", cfg.Port)
	}
	if err != nil {
		log.Error("Unable to listen: ", err)
		return "Error", err
	}

	srv := &http.Server{Handler: service}
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Errorf("serve: %s\n", err)
		}
	}()
	log.Info("Server started.")

	killSignal := <-interrupt
	if killSignal == os.Interrupt {
		log.Info("Daemon was interrupted by SIGINT")
	} else {
		log.Info("Daemon was killed")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}

	return "Server exited properly", nil
}

func (service *Daemon) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	buf := make([]byte, 4096)
	numBytes, err := req.Body.Read(buf)
	if numBytes == 0 || err != nil {
		return
	}
	// Work with API call
}
