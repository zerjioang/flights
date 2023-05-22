package embed

import (
	_ "embed"
)

var (
	//go:embed VERSION.txt
	Version string
)
