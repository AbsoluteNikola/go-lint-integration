package main

import (
	"github.com/BelehovEgor/golang-linter/pkg/avoidMutableGlobals"
	"github.com/BelehovEgor/golang-linter/pkg/channelSize"
	"golang.org/x/tools/go/analysis/multichecker"
)

func main() {
	multichecker.Main(
		channelSize.Analyzer,
		avoidMutableGlobals.Analyzer,
	)
}
