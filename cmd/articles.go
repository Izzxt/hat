/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// articlesCmd represents the articles command
var articlesCmd = &cobra.Command{
	Use:   "articles",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("articles called")
	},
}

func init() {
	rootCmd.AddCommand(articlesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	articlesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// articlesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
