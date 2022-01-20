package main

import (
	"fmt"
	"runtime"
)

// All of this added by goreleaser using ldflags
var (
	// Short Git Commit Hash
	CommitHash string
	// Version vx.x.x
	Version string
	// Date of build
	BuildDate string
)

func versionString() string {
	return fmt.Sprintf("%s on %s (%s), %s", Version, BuildDate, CommitHash, runtime.Version())
}
