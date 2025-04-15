/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"kivi-cache/cache"
	"kivi-cache/server/internal"
	"log"
	"net"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

const defaultPort = "5001"

func getPort(cmd *cobra.Command) string {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		return ":" + defaultPort
	}

	if port[0] != ':' {
		return ":" + port
	}

	return port
}

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
		port := getPort(cmd)
		listener, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen on port %s: %v", port, err)
		}

		grpcSrv := grpc.NewServer()
		cache.RegisterKiviCacheServiceServer(grpcSrv, internal.NewCacheServer())
		log.Printf("gRPC server listening at %v", listener.Addr())

		if err := grpcSrv.Serve(listener); err != nil {
			log.Fatalf("Failed to serve %s", err)
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
	rootCmd.Flags().StringP("port", "p", defaultPort, "The gRPC endpoint will be exposed on this port")
}
