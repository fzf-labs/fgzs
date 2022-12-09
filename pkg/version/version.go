package version

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

var (
	gitCommit = "" // sha1 from git, output of $(git rev-parse HEAD)
	gitBranch = "" // state of git tree, either "clean" or "dirty"
	buildTime = "" // build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
)

// Info contains versioning information.
type Info struct {
	GitCommit string `json:"git_commit"`
	GitBranch string `json:"git_branch"`
	BuildTime string `json:"build_time"`
	GoVersion string `json:"go_version"`
	Compiler  string `json:"compiler"`
	Platform  string `json:"platform"`
}

// String returns info as a human-friendly version string.
func (info *Info) String() string {
	return fmt.Sprintf("%s (%s) [%s]", info.GitBranch, info.GitCommit, info.BuildTime)
}

// Get 返回详细的版本信息
func Get() Info {
	return Info{
		GitCommit: gitCommit,
		GitBranch: gitBranch,
		BuildTime: buildTime,
		GoVersion: runtime.Version(),
		Compiler:  runtime.Compiler,
		Platform:  fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func Parse(b bool) {
	if b {
		ver := Get()
		marshaled, err := json.MarshalIndent(&ver, "", "  ")
		if err != nil {
			fmt.Printf("get version err:%v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(marshaled))
		os.Exit(1)
		return
	}
}
