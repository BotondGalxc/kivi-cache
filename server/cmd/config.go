package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	defaultServerPort  = "5001"
	defaultMetricsPort = "2112"
)

type ServerConfiguration struct {
	serverPort  string
	metricsPort string
}

func NewServerConfiguration(ccmd *cobra.Command) ServerConfiguration {

	setupViperConf(ccmd)
	return ServerConfiguration{
		serverPort:  getServerPort(ccmd),
		metricsPort: getMetricsPort(),
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
	portFromFlag, err := cmd.Flags().GetString("port")
	if err != nil {
		logger.Debug("Get port from flag")
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
