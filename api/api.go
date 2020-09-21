package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/hhyouke/server/auth"
	"github.com/hhyouke/server/auth/providers/password"
	"github.com/hhyouke/server/conf"
	"github.com/hhyouke/server/logger"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"
	"github.com/sebest/xff"
)

var defaultWait time.Duration = time.Second * 30
var machineID string

// API the rest api
type API struct {
	Authentication *auth.Auth
	handler        http.Handler
	config         *conf.APIConfiguration
	logger         *logger.AppLogger
	DB             *gorm.DB
}

// NewAPI create a new api instance
func NewAPI(ctx context.Context, apiServerConfig *conf.APIConfiguration, appLogger *logger.AppLogger, db *gorm.DB) *API {
	api := &API{
		config: apiServerConfig,
		logger: appLogger,
		DB:     db,
	}

	machineID = apiServerConfig.MachineID

	xff, _ := xff.Default()
	apiLogger := newStructuredLogger(appLogger)

	authentication := auth.New(&auth.Config{
		DB:        db,
		Logger:    api.logger,
		MachineID: machineID,
	})
	authentication.RegisterProvider(password.New(&password.Config{}))
	api.Authentication = authentication

	r := newRouter()

	r.UseBypass(xff.Handler)
	r.UseBypass(apiLogger)
	r.UseBypass(recoverer)
	r.Use(withRequestID)
	r.Use(withMachineID)
	r.Get("/health", api.HealthCheck)

	r.Route("/auth/{provider}/{action}", func(r *router) {
		r.Handle("/", api.Auth)
	})

	corsHandler := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PATCH", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link", "X-Total-Count"},
		AllowCredentials: true,
	})

	api.handler = corsHandler.Handler(chi.ServerBaseContext(ctx, r))

	// http.ListenAndServe()

	return api
}

// ListenAndServe starts the api server
func (a *API) ListenAndServe(hostAndPort string) {
	server := &http.Server{
		Addr:    hostAndPort,
		Handler: a.handler,
	}
	// Run server in a goroutine so that it doesn't block.
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("api server startup failed, %v\n", err.Error())
		}
	}()
	log.Printf("server started at %s\n", hostAndPort)
	gracefulShutdown(server)
}

func gracefulShutdown(httpServer *http.Server) {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	log.Printf("server shutting down triggered by %v\n", sig.String())
	// srv.l.Info("Server is shutting down", zap.String("reason", sig.String()))

	ctx, cancel := context.WithTimeout(context.Background(), defaultWait)
	defer cancel()

	// srv.SetKeepAlivesEnabled(false)
	if err := httpServer.Shutdown(ctx); err != nil {
		// srv.l.Fatal("Could not gracefully shutdown the server", zap.Error(err))
		log.Fatalf("Could not gracefully shutdown the server %v\n", err.Error())
	}
	log.Println("Server shutted down")
}

func withRequestID(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()
	id, err := uuid.NewRandom()
	if err != nil {
		return ctx, err
	}
	idString := id.String()
	ctx = WithRequestID(ctx, idString)
	return ctx, nil
}

func withMachineID(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()
	ctx = WithMachineID(ctx, machineID)
	return ctx, nil
}
