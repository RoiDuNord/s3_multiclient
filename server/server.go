package server

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"s3_multiclient/config"
	"syscall"

	_ "s3_multiclient/docs"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

type LoadManager interface {
	Upload(r *http.Request, ctx context.Context, data *UploadRequestMetadata) error
	Download(w http.ResponseWriter, ctx context.Context, data *DownloadRequestMetadata) error
	Delete(w http.ResponseWriter, r *http.Request, ctx context.Context) error
}

// type DBManager interface{
// 	Upload(ctx context.Context, w http.ResponseWriter, r *http.Request)
// 	Download(ctx context.Context, w http.ResponseWriter, r *http.Request)
// 	Delete(ctx context.Context, w http.ResponseWriter, r *http.Request)
// }

// type DefaultLoadManager struct {
// 	fileManager load.FileManager
// 	dbManager   db.DBManager
// }

// func NewDefaultLoadManager(fm load.FileManager, dm db.DBManager) *DefaultLoadManager {
// 	return &DefaultLoadManager{
// 		fileManager: fm,
// 		dbManager:   dm,
// 	}
// }

type Server struct {
	ctx         context.Context
	loadManager LoadManager
	// dbManager DBManager
}

func Init(ctx context.Context, lm LoadManager) *Server {
	return &Server{
		ctx:         ctx,
		loadManager: lm,
	}
}

//	@title			S3 Multiclient API
//	@version		1.0
//	@description	API for uploading and downloading files from S3 storage (MinIO, Ceph)
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.example.com/support
//	@contact.email	support@example.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		localhost:8080
// @BasePath	/
func (s *Server) setupRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Post("/{storage_name}/{relative_path}/objects/{object_id}/content", s.Upload)
	router.Get("/{storage_name}/{relative_path}/objects/{object_id}/content", s.Download)
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	return router
}

func (s *Server) MustStart(srvcfg config.ServerConfig) error {
	router := s.setupRouter()
	address := fmt.Sprintf("%s:%d", srvcfg.Host, srvcfg.Port)

	slog.Info("запуск HTTP сервера", "address", address)
	httpServer := &http.Server{
		Address: address,
		Handler: router,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			// slog.Error("Ошибка запуска сервера", "error", err)
			log.Fatalf("Сервер не стартовал: %s", err.Error())
		}
	}()

	return s.gracefulShutdown(httpServer)
}

func (s *Server) gracefulShutdown(server *http.Server) error {
	shutdownSignals := make(chan os.Signal, 1)
	signal.Notify(shutdownSignals, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)

	select {
	case <-shutdownSignals:
		slog.Info("Получен сигнал завершения работы")
	case <-s.ctx.Done():
		slog.Info("Истекло время ожидания контекста")
	}

	if err := server.Shutdown(s.ctx); err != nil {
		slog.Error("Не удалось корректно завершить работу сервера", "error", err)
		return err
	}

	slog.Info("Сервер успешно завершил работу")
	return nil
}
