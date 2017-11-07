// Copyright © 2017 Suman Roy

package cmd

import (
  "fmt"
  "os"

  "github.com/mitchellh/go-homedir"
  "github.com/spf13/cobra"
  "github.com/spf13/viper"
)

var cfgFile string
var Verbose bool

// TptCmd represents the base command when called without any subcommands
var TptCmd = &cobra.Command{
  Use:   "tpt",
  Short: "tpt cli",
  Long: `tpt cli`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the TptCmd.
func Execute() {
  if err := TptCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
  cobra.OnInitialize(initConfig)

  // Persistent flags which are global for the application.
  TptCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tpt.yaml)")
  TptCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
  viper.BindPFlag("verbose", TptCmd.Flag("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
  if cfgFile != "" {
    // Use config file from the flag.
    viper.SetConfigFile(cfgFile)
  } else {
    // Find home directory.
    home, err := homedir.Dir()
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    // Search config in home directory with name ".tpt" (without extension).
    viper.AddConfigPath(home)
    viper.SetConfigName(".tpt")
  }

  viper.AutomaticEnv() // read in environment variables that match

  // If a config file is found, read it in.
  if err := viper.ReadInConfig(); err == nil && Verbose {
    fmt.Printf("Using config file:%s\n\n", viper.ConfigFileUsed())
  }
}
