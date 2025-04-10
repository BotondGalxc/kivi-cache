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

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put [key] [value]",
	Short: "Adds a new value to the cache",
	Long:  `Adds a new value to the cache. The key is used to retrieve the value with a Get command.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := internal.NewClient("localhost", "5001", insecure.NewCredentials())
		if err != nil {
			fmt.Printf("Cannot create client to localhost:5001: %v", err)
		} else {
			response, err := client.Put(args[0], args[1])
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(response)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(putCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// putCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// putCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
