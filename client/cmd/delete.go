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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Remove value from cache",
	Long:  `Remove value from cache addressed by a given key`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		client, err := internal.NewClient("localhost", "5001", insecure.NewCredentials())
		if err != nil {
			fmt.Printf("Cannot create client to localhost:5001: %v", err)
		} else {
			res, err := client.Delete(args[0])
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println(res)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
