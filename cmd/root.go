package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Output  string
	Domain  string
	CfgFile string
	Prod    string
)

var rootCmd = &cobra.Command{
	Use:   "hat",
	Short: "Habbo Downloader Tools",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Output, "output", "o", "", "Folder output")
	rootCmd.PersistentFlags().StringVarP(&Domain, "domain", "d", "com", "com.br, com.tr, com, de, es, fi, fr, it, nl")
	// rootCmd.PersistentFlags().StringVarP(&CfgFile, "config", "c", "", "Config file")
	rootCmd.PersistentFlags().StringVarP(&Prod, "production", "p", "", "Habbo gordon production")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".hat" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".hat")

		if _, err := os.Stat(home + "/.hat.json"); errors.Is(err, os.ErrNotExist) {
			fmt.Print("not exists")
			if _, err := os.Create(home + "/.hat.json"); err != nil {
				fmt.Print("created")
				log.Fatal(err)
			}
			viper.SetDefault("test", "hello")
			viper.WriteConfig()
		}

	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Config file not found:", viper.ConfigFileUsed())
	}
}
