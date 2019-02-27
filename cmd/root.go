// Copyright © 2019 Sooz Richman (FavoredFortune - GitHub)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

//use constructor from kvchache package for global access to struct and it's private fields per Troy
var cache = kvcache.NewSimpleKVCache()

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cli",
	Short: "a simple key-value cache cli",
	Long: `This is a CLI app that allows you to input your action and your key;value pair of strings.

	Actions available include put, read,  update or delete your content from the cache.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {

		case "put":
			if len(args) < 3 {
				return errors.New("put failed: put command and both key and value strings required")
			}
			cache.Put(args[1], args[2])

		case "read":
			if len(args) < 2 {
				return fmt.Errorf("READ command incomplete: %v", args)
			}
			_, readResult := cache.Read(args[1])
			fmt.Println(">", readResult)

		case "update":
			if len(args) < 3 {
				return fmt.Errorf("UPDATE command incomplete: %v", args)
			}
			err := cache.Update(args[1], args[2])
			if err != nil {
				return err
			}

		case "delete":
			if len(args) < 2 {
				return fmt.Errorf("DELETE command incomplete: %v", args)
			}
			err := cache.Delete(args[1])
			if err != nil {
				return err
			}
		}
		return nil
	},
}

//var putCmd = &cobra.Command{
//	Use:   "put",
//	Short: "put key-value pair",
//	Long:  "put key value strings into the key-value cache",
//	RunE: func(cmd *cobra.Command, args []string) error {
//
//		if len(args) < 3 {
//			return errors.New("put failed: put command and both key and value strings required")
//		}
//		cache.Put(args[1], args[2])
//		return nil
//	},
//}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.KVCache.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")


	//maybe don't need flags at all can put logic all in initial command?
	//persistent flags
	//putCmd.PersistentFlags().StringSlice("put","put data in cache")
	//
	//putCmd.Flags().String
	//
	////bind flags
	//viper.BindPFlag("put", putCmd.PersistentFlags().Lookup("put"))

	//local flags
	//do I need them?
	//
	////commands
	//putCmd.AddCommand(putCmd)

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

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".KVCache")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
