package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"shortLink/internal/config"
	appMetrics "shortLink/internal/metrics"
	"shortLink/internal/service"
	"shortLink/internal/store"
	"shortLink/internal/store/memory"
	_ "shortLink/internal/store/memory"
	appTracing "shortLink/internal/tracing"
	"strings"

	_ "go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

type App struct {
	appName        string
	serviceVersion string
	logger         *slog.Logger
	store          store.Store
	cfg            *config.Config
	httpServer     *http.Server
	shortenerSvc   service.ShortenerService
	tracerProvider *sdktrace.TracerProvider
}

func New(cfg *config.Config, logger *slog.Logger, appName string, version string) (*App, error) {
	app := &App{
		cfg:            cfg,
		logger:         logger.With("component", "appcore"),
		appName:        appName,
		serviceVersion: version,
	}
	app.logger.Info("Creating new application instance...",
		slog.String("appName", app.appName),
		slog.String("version", app.serviceVersion),
	)
	if app.cfg.Tracing.Enabled {

		var errInitTracer error
		app.tracerProvider, errInitTracer = appTracing.InitTracerProvider(
			app.appName,
			app.serviceVersion,
			app.cfg.Tracing.Enabled,
			app.cfg.Tracing.SampleRatio,
		)

		if errInitTracer != nil {
			app.logger.Error("Failed to initialize TracerProvider", slog.Any("error", errInitTracer))
			return nil, fmt.Errorf("app.New: failed to init tracer provider: %w", errInitTracer)
		}
	}
	// 1b. 初始化 Metrics (Prometheus Go运行时等)
	appMetrics.Init()
	app.logger.Info("Prometheus Go runtime metrics collectors registered.")

	// --- [App.New - 初始化阶段 2] ---
	// 初始化核心业务依赖 (层级：Store -> Service -> Handler)
	app.logger.Debug("Initializing core dependencies...")
	// 2a. 初始化 Store 层
	switch strings.ToLower(app.cfg.Store.Type) {
	case "memory":
		app.store = memory.NewStore(app.logger.With("datastore", "memory"))
		app.logger.Info("Initialized in-memory store.")
	default:
		err := fmt.Errorf("unsupported store type from config: %s", app.cfg.Store.Type)
		app.logger.Error("Failed to initialize store", slog.Any("error", err))
		return nil, err
	}

	return app, nil
}
