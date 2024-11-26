package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/ShatteredRealms/dimension-service/pkg/config"
	"github.com/ShatteredRealms/dimension-service/pkg/pb"
	"github.com/ShatteredRealms/dimension-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/log"
	commonpb "github.com/ShatteredRealms/go-common-service/pkg/pb"
	commonsrv "github.com/ShatteredRealms/go-common-service/pkg/srv"
	"github.com/ShatteredRealms/go-common-service/pkg/telemetry"
	"github.com/ShatteredRealms/go-common-service/pkg/util"
	"github.com/WilSimpson/gocloak/v13"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Load configuration and setup server context
	cfg, err := config.NewDimensionConfig(ctx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("loading config: %v", err)
		return
	}

	srvCtx, err := srv.NewDimensionContext(ctx, cfg, config.ServiceName)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("dimension server context: %v", err)
		return
	}
	ctx, span := srvCtx.Tracer.Start(ctx, "main")
	defer span.End()

	log.Logger.WithContext(ctx).Infof("Starting %s service", config.ServiceName)

	// OpenTelemetry setup
	otelShutdown, err := telemetry.SetupOTelSDK(ctx, "dimension", config.Version, cfg.OpenTelemtryAddress)
	defer func() {
		log.Logger.Infof("Shutting down")
		err = otelShutdown(context.Background())
		if err != nil {
			log.Logger.Warnf("Error shutting down: %v", err)
		}
	}()

	if err != nil {
		log.Logger.WithContext(ctx).Errorf("connecting to otel: %v", err)
		return
	}

	// Common gRPC server setup
	keycloakClient := gocloak.NewClient(cfg.Keycloak.BaseURL)
	grpcServer, gwmux := util.InitServerDefaults(keycloakClient, cfg.Keycloak.Realm)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Health Service
	commonpb.RegisterHealthServiceServer(grpcServer, commonsrv.NewHealthServiceServer())
	err = commonpb.RegisterHealthServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register health service handler endpoint: %v", err)
		return
	}

	// Dimension Service
	dimensionService, err := srv.NewDimensionServiceServer(ctx, srvCtx)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("creating dimension service: %v", err)
		return
	}
	pb.RegisterDimensionServiceServer(grpcServer, dimensionService)
	err = pb.RegisterDimensionServiceHandlerFromEndpoint(ctx, gwmux, cfg.Server.Address(), opts)
	if err != nil {
		log.Logger.WithContext(ctx).Errorf("register dimension service handler endpoint: %v", err)
		return
	}

	// Setup Complete
	log.Logger.WithContext(ctx).Info("Initializtion complete")
	span.End()

	srvErr := make(chan error, 1)
	go func() {
		srvErr <- util.StartServer(ctx, grpcServer, gwmux, cfg.Server.Address())
	}()

	select {
	case err := <-srvErr:
		if err != nil {
			log.Logger.Error(err)
		}

	case <-ctx.Done():
		log.Logger.Info("Server canceled by user input.")
		stop()
	}

	log.Logger.Info("Server stopped.")

}
