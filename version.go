package version

import (
	"runtime"
)

var (
	version   = "version-dev"
	commit    = "none-dev"
	branch    = "branch-dev"
	codeName  = "name-dev"
	buildDate = "build-date"
	buildTime = "build-time"
)

var info *Info

type Info struct {
	Version    string
	Commit     string
	Branch     string
	BuildDate  string
	BuildTime  string
	CodeName   string
	TargetOs   string
	TargetArch string
	GoVersion  string
}

func GetInfo() *Info {
	if nil == info {
		info = &Info{
			version,
			commit,
			branch,
			buildDate,
			buildTime,
			codeName,
			runtime.GOOS,
			runtime.GOARCH,
			runtime.Version(),
		}
	}
	return info
}
