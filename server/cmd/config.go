package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultServerPort  = "5001"
	defaultMetricsPort = "2112"
	defaultLogLevel    = slog.LevelInfo
)

type ServerConfiguration struct {
	serverPort  string
	metricsPort string
	logLevel    slog.Leveler
}

func NewServerConfiguration(ccmd *cobra.Command) ServerConfiguration {

	setupViperConf(ccmd)
	return ServerConfiguration{
		serverPort:  getServerPort(ccmd),
		metricsPort: getMetricsPort(),
		logLevel:    getDefaultLogLevel(ccmd),
	}
}

func setupViperConf(ccmd *cobra.Command) {
	configPath, _ := ccmd.Flags().GetString("config-path")

	// Setup viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("metrics.port", defaultMetricsPort)
	viper.AddConfigPath(configPath)

	viper.ReadInConfig()
}

func getServerPort(cmd *cobra.Command) string {

	// Get config/env
	port := viper.GetString("server.port")

	// Flags can overwrite conf and env
	portFromFlag, _ := cmd.Flags().GetString("port")
	if cmd.Flags().Changed("port") {
		port = portFromFlag
	}

	if port[0] != ':' {
		return ":" + port
	}

	return port
}

func getMetricsPort() string {
	// Get config/env
	port := viper.GetString("metrics.port")

	if port[0] != ':' {
		return ":" + port
	}

	return port
}

func getDefaultLogLevel(cmd *cobra.Command) slog.Leveler {
	value := viper.GetString("log.level")

	valueFromFlag, _ := cmd.Flags().GetString("log-level")
	if cmd.Flags().Changed("log-level") {
		slog.Info("Use value from Flag")
		value = valueFromFlag
	}

	switch value {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "ERROR":
		return slog.LevelError
	case "WARN":
		return slog.LevelWarn
	default:
		return defaultLogLevel
	}

}
