package config

import "os"

var AppCfg AppConfig

type AppConfig struct {
	SelfHostedDomain string
	GitLabToken      string
	GitLabCfg        GitLabConfig
}

type GitLabConfig struct {
	ProjectIds ProjectIds
}

type ProjectIds struct {
	Dev string
}

func LoadAppConfig() {
	AppCfg = AppConfig{
		SelfHostedDomain: os.Getenv("SELF_HOSTED_DOMAIN"),
		GitLabToken:      os.Getenv("PERSONAL_GITLAB_TOKEN"),
		GitLabCfg: GitLabConfig{
			ProjectIds: ProjectIds{
				Dev: os.Getenv("GITLAB_PROJECT_ID"),
			},
		},
	}
}
