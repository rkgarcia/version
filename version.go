package version

import "runtime"

var (
	version    = "version-dev"
	commit     = "none-dev"
	branch     = "branch-dev"
	buildDate  = "date-dev"
	buildTime  = "time-dev"
	codeName   = "name-dev"
	targetOs   = "os-dev"
	targetArch = "arch-dev"
	goVersion  = "go-dev"
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
			targetOs,
			targetArch,
			runtime.Version(),
		}
	}
	return info
}
