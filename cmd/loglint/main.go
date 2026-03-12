// Command loglint runs the loglint analyzer as a singlechecker binary.
package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/irumako/loglint"
)

func main() {
	singlechecker.Main(loglint.NewAnalyzer())
}
