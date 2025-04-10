/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"kivi-cache/client/internal"

	"github.com/spf13/cobra"
	"google.golang.org/grpc/credentials/insecure"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Retrieves the value for a key from a kivi-server",
	Long:  `Retrieves the value for a key from a kivi-server`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("host")
		port, _ := cmd.Flags().GetString("port")

		client, err := internal.NewClient(host, port, insecure.NewCredentials())
		if err != nil {
			fmt.Printf("Cannot create client to %s:%s: %v", host, port, err)
		} else {
			kv := client.Get(args[0])
			fmt.Printf("Received key %s=%s", kv.Key, kv.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
