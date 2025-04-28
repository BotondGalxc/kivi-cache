/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"kivi-cache/cache"
	"kivi-cache/server/internal"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cacheEntriesCount = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Namespace: "kivicache",
			Name:      "entries_count",
			Help:      "Indicates, how much keys are recoreded in the cache",
		})
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		conf := NewServerConfiguration(cmd)
		listener, err := net.Listen("tcp", conf.serverPort)
		if err != nil {
			logger.Error(fmt.Sprintf("failed to listen on port %s: %v", conf.serverPort, err))
		}

		http.Handle("/metrics", promhttp.Handler())
		go http.ListenAndServe(conf.metricsPort, nil)

		logger.Info(fmt.Sprintf("Metrics available at [::]%v/metrics", conf.metricsPort))

		cacheServer := internal.NewCacheServer(logger)

		prometheus.MustRegister(cacheEntriesCount)
		go func() {
			for {
				cacheEntriesCount.Set(float64(cacheServer.Count()))

				time.Sleep(time.Second)
			}
		}()

		grpcSrv := grpc.NewServer()
		cache.RegisterKiviCacheServiceServer(grpcSrv, cacheServer)
		logger.Info(fmt.Sprintf("gRPC server listening at %v", listener.Addr()))

		if err := grpcSrv.Serve(listener); err != nil {
			logger.Error(fmt.Sprintf("Failed to serve gRPC: %s", err))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.server.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("port", "p", defaultServerPort, "The gRPC endpoint will be exposed on this port")
	rootCmd.Flags().String("config-path", ".", "The gRPC endpoint will be exposed on this port")

}
