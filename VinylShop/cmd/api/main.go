package main

import (
	vinyls "VinylShop/internal/vinyls/delivery"
	"VinylShop/pkg/configs"
	pg "VinylShop/pkg/drivers/postgressql"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: false,
		DisableColors:    false,
		TimestampFormat:  "2006-01-02 15:04:05",
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "@time",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyLevel: "lvl",
		},
	})
	logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	logrus.SetOutput(logFile)
}
func main() {
	cfg := configs.LoadAppConfig()
	w := logrus.New().WriterLevel(logrus.ErrorLevel)
	defer w.Close()
	db, err := pg.PGConnect(cfg.Db)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"err": err.Error(),
			"dsn": fmt.Sprintf("%s@%s:%s/%s", cfg.Db.Type, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database),
		}).Fatal("Postgres connection failed")
		os.Exit(1)
	}
	logrus.WithField("dsn", fmt.Sprintf("%s@%s:%s/%s", cfg.Db.Type, cfg.Db.Host, cfg.Db.Port, cfg.Db.Database))
	vinylsDelivery := vinyls.NewVinylsDelivery(db)
	router := initRouter()
	router.Route("/", func(router chi.Router) {
		router.Mount("/vinyls", vinylsDelivery.Routes())
	})
	srv := &http.Server{
		Addr:     cfg.Addr,
		Handler:  router,
		ErrorLog: log.New(w, "", 0),
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		logrus.WithField("addr", srv.Addr).Info("server started")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("ListenAndServe: %v", err)
		}
	}()
	<-quit
	logrus.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown: %v", err)
	}
	logrus.Info("Server exiting")
}
