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

package main

import (
	"KVCache/kvcache"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

func main() {

	//use constructor from kvchache package for global access to struct and it's private fields per Troy
	var cache = kvcache.NewSimpleKVCache()

	//make root command not executable without subcommand by not providing a 'Run' for the 'rootCmd'
	var RootCmd = &cobra.Command{Use:"cli"}
	var putCmd = &cobra.Command{
		Use:   "put",
		Short: "put key-value pair",
		Long:  "put key value strings into the key-value cache",
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) < 3 {
				return errors.New(">>put failed: put command and both key and value strings required")
			}
			cache.Put(args[1], args[2])
			fmt.Printf(" key '%v' and value '%v' put into the cache", args[1], args[2])
			return nil
		},
	}

	var readCmd = &cobra.Command{
		Use:"read",
		Short: "read value given key",
		Long: "read value string out to command line from key-value cache given key string input from command line",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 2 {
				return fmt.Errorf(">>read failed: %v", args)
			}
			_, readResult := cache.Read(args[2])
			fmt.Println(">>", readResult)
			return nil
		},
	}

	//attach subcommands to rootcommand
	RootCmd.AddCommand(putCmd, readCmd)
	RootCmd.Execute()
}



