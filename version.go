package version

import (
	"runtime"
	"time"
)

var (
	version  = "version-dev"
	commit   = "none-dev"
	branch   = "branch-dev"
	codeName = "name-dev"
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
		today := time.Now()
		info = &Info{
			version,
			commit,
			branch,
			today.Format("02-01-2006"),
			today.Format("15:04:05"),
			codeName,
			runtime.GOOS,
			runtime.GOARCH,
			runtime.Version(),
		}
	}
	return info
}
