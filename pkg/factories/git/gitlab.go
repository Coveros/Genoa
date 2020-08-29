package git

import (
	"coveros.com/pkg/utils"
	"os"
)

type gitlab struct{}

func NewGitlab() *gitlab {
	return &gitlab{}
}

func (g gitlab) GetAccessToken() (string, error) {
	token, ok := os.LookupEnv(utils.EnvVarGitlabPersonalAccessToken)
	if !ok || token == "" {
		return "", ErrorEnvVarNotFound{Message: "Gitlab personal access token not specified in env vars"}
	}
	return token, nil
}

func (g gitlab) GetDeployDir() (string, error) {
	deployDir, ok := os.LookupEnv(utils.EnvVarGitlabReleaseFilesDir)
	if !ok || deployDir == "" {
		return "", ErrorEnvVarNotFound{Message: "Gitlab deploy dir not specified in env vars"}
	}
	return deployDir, nil
}

func (g gitlab) GetWebhookSecret() (string, error) {
	secret, ok := os.LookupEnv(utils.EnvVarGitlabWebhookSecret)
	if !ok || secret == "" {
		return "", ErrorEnvVarNotFound{Message: "Gitlab webhook secret not specified in env vars"}
	}
	return secret, nil
}

func (g gitlab) GetSelfHostedUrl() (string, error) {
	url, ok := os.LookupEnv(utils.EnvVarGitlabSelfHostedUrl)
	if !ok || url == "" {
		return "", ErrorEnvVarNotFound{Message: "Gitlab self-hosted url not specified in env vars"}
	}
	return url, nil
}
