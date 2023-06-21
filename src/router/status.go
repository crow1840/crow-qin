package router

import (
	"github.com/gin-gonic/gin"
)

type VersionResponse struct {
	Version   string `json:"version"`
	BuildTime string `json:"build_time"`
	BuildHost string `json:"build_host"`
	GOVersion string `json:"go_version"`
	GitLog    string `json:"git_log"`
	GitBranch string `json:"git_branch"`
}

var (
	CrowVersion string
	BuildTime   string
	BuildHost   string
	GOVersion   string
	GitBranch   string
	GitLog      string
)

func Ping(c *gin.Context) {
	SetOK(c, "pong")
}

func Version(c *gin.Context) {
	var versionResponse VersionResponse
	versionResponse.Version = CrowVersion
	versionResponse.BuildTime = BuildTime
	versionResponse.BuildHost = BuildHost
	versionResponse.GOVersion = GOVersion
	versionResponse.GitLog = GitLog
	versionResponse.GitBranch = GitBranch

	SetOK(c, versionResponse)
}

type UserName struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Age      int64  `json:"age"`
}
