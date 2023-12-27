/*
Copyright © 2023 Max

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (

	// KEY_PDFLATEX is the config key for the PDFLATEX command line tool
	KEY_PDFLATEX = "pdflatex"

	// KEY_OUTPUTDIRNAME is the name of the generated folder for the output files
	KEY_OUTPUTDIRNAME = "outputdirname"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "writings",
	Short: "A tool to generate PDFs and writing templates",
	Long: `writings is a CLI tool for Go that empowers writing.
	This application is used for generating PDFs and 
	writing templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Using config: " + viper.GetString(KEY_PDFLATEX))
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

	viper.SetDefault(KEY_PDFLATEX, "/Library/TeX/texbin/pdflatex")
	viper.SetDefault(KEY_OUTPUTDIRNAME, "generated")

	viper.SetConfigName(".writings") // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")         // optionally look for config in the working directory

	// Read config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Default().Println("Creating default config file...")

			// Save config file if it does not exist
			if writeErr := viper.SafeWriteConfigAs(".writings.yaml"); writeErr != nil {
				log.Fatal(writeErr)
			}
		} else {
			// Config file was found but another error was produced
			log.Fatal(err)
		}
	}

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.writings.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
