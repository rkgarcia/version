# Version

This struct only is for organization to use with goreleaser.
Adding ldflags to yaml definition:

```markdown
builds:
  - ldflags:
      - -s -w
      - -X github.com/rkgarcia/version.version={{.Version}}
      - -X github.com/rkgarcia/version.commit={{.ShortCommit}}
      - -X github.com/rkgarcia/version.branch={{.Branch}}
      - -X github.com/rkgarcia/version.codeName={{.Env.CODENAME}}
      - -X github.com/rkgarcia/version.buildDate={{ .Now.Format "02-01-2006" }}
      - -X github.com/rkgarcia/version.buildTime={{ .Now.Format "15:04:05" }}
```

Now we can use in out application something like:
```go
information := version.GetInfo()
log.Println("Running app version:",
    information.Version, information.Commit, information.Branch,
    information.BuildDate, information.BuildTime, information.CodeName,
    information.TargetOs, information.TargetArch)
```