// Package version contains build information such as the git commit, app version, build date.
//
// This info generated at build time and compiled into the binary.
//
// Usage:
// go build -o ${BIN_OUT} "${GO_BUILD_PACKAGE}"

package version

import (
	"runtime/debug"
)

const unset = "unset"

var ( // build info
	version     = unset
	builddate   = unset
	commit      = unset
	shortcommit = unset
	appname     = unset
	goversion   = unset
)

func init() {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	goversion = info.GoVersion

	var modified bool

	for _, setting := range info.Settings {
		switch setting.Key {
		case "vcs.revision":
			commit = setting.Value
		case "vcs.time":
			builddate = setting.Value
		case "vcs.modified":
			modified = true
		}
	}

	if modified {
		commit += "+CHANGES"
	}

	appname = info.Main.Path
}

// GetGoVersion returns the go version
func GetGoVersion() string {
	return goversion
}

// GetVersion returns the app version
func GetVersion() string {
	return version
}

// GetBuildDate returns the build date
func GetBuildDate() string {
	return builddate
}

// GetCommit returns the git commit
func GetCommit() string {
	return commit
}

// GetAppName returns the app name
func GetAppName() string {
	return appname
}

// GetShortCommit returns the short git commit
func GetShortCommit() string {
	return shortcommit
}
