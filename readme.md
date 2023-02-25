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
      - -X github.com/rkgarcia/version.buildDate={{time "2006-01-02"}}
      - -X github.com/rkgarcia/version.buildTime={{time "15:04:05"}}
      - -X github.com/rkgarcia/version.codeName={{.Env.CODENAME}}
      - -X github.com/rkgarcia/version.targetOs={{.Runtime.Goos}}
      - -X github.com/rkgarcia/version.targetArch={{.Runtime.Goarch}}
```