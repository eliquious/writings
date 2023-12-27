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
	"errors"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// generateCmdCmd represents the generateCmd command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// cwd, getwdErr := os.Getwd()
		// if getwdErr != nil {
		// 	log.Fatal(getwdErr)
		// }

		if inputfilename, absErr := filepath.Abs(args[0]); absErr != nil {

			// failure to get absolute path
			log.Fatal(absErr)
		} else {

			// use absolute path for input file
			if generateErr := generateLatex(inputfilename); generateErr != nil {
				log.Fatal(generateErr)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}

func generateLatex(filename string) error {
	directory := filepath.Dir(filename)
	generated := filepath.Join(directory, viper.GetString(KEY_OUTPUTDIRNAME))
	log.Println("Output directory: " + generated)

	// Make output directory
	if mkdirErr := os.MkdirAll(generated, os.ModePerm); mkdirErr != nil {
		log.Fatal(mkdirErr)
	}

	// Prepare command arguments
	outputArg := strings.Join([]string{"-output-dir", generated}, "=")
	var args = []string{"-jobname=generated", "-halt-on-error", outputArg, filename}

	// Prepare the command.
	var cmd = exec.Command(viper.GetString(KEY_PDFLATEX), args...)
	log.Println(cmd.String())

	// Set the cwd to the parent directory of the input file; this is for relative processing paths
	cmd.Dir = directory

	// // Set $TEXINPUTS if requested. The trailing colon means that LaTeX should
	// // include the normal asset directories as well.
	// if options.Texinputs != "" {
	// 	cmd.Env = append(os.Environ(), "TEXINPUTS="+options.Texinputs+":")
	// }

	// Launch and let it finish.
	var err = cmd.Start()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		// The actual error is useless, do provide a better one.
		return errors.New("LaTeX error. Check " + path.Join(generated, "generated.log"))
	}

	return nil
}
