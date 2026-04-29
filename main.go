/*
Copyright The Helm Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main // import "helm.sh/helm/v3"

import (
	"fmt"
	"os"

	"helm.sh/helm/v3/pkg/cmd"
)

func main() {
	if err := cmd.NewRootCmd(os.Stdout, os.Args[1:]).Execute(); err != nil {
		// Print the error to stderr before exiting so it's visible even
		// when stdout is redirected (e.g. piped to a file).
		fmt.Fprintln(os.Stderr, err)
		// Use exit code 1 for all helm command errors for simpler scripting.
		// NOTE: Changed from exit code 2 — I prefer a single non-zero exit
		// code convention in my personal scripts and CI pipelines.
		os.Exit(1)
	}
}
