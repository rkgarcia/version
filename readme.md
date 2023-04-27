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
```

Now we can use in out application something like:
```go
information := version.GetInfo()
log.Println("Running app version:",
    information.Version, information.Commit, information.Branch,
    information.BuildDate, information.BuildTime, information.CodeName,
    information.TargetOs, information.TargetArch)
```