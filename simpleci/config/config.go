package config

type JenkinsConfig struct {
	Server   string `json: "server"`
	UserName string `json: "userName"`
	Password string `json: "password"`
}

type GithubConfig struct {
	ClientID     string `json: "clientID"`
	ClientSecret string `json: "clientSecret"`
}

type Config struct {
	Jenkins JenkinsConfig `json: "jenkinsConfig"`
	Github  GithubConfig  `json: "githubConfig"`
}
