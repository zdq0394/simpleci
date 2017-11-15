package jenkinsclient

import (
	"fmt"

	"github.com/zdq0394/simpleci/simpleci/jenkinsclient/gojenkins"
)

type JenkinsClient struct {
	JenkinsURL string
	UserName   string
	Password   string
	Jenkins    *gojenkins.Jenkins
}

func NewJenkinsClient(jenkinsURL, username, password string) *JenkinsClient {
	jenkins := gojenkins.CreateJenkins(nil, jenkinsURL, username, password)
	return &JenkinsClient{
		JenkinsURL: jenkinsURL,
		UserName:   username,
		Password:   password,
		Jenkins:    jenkins,
	}
}

func (c *JenkinsClient) Refresh() {
	c.Jenkins = gojenkins.CreateJenkins(nil, c.JenkinsURL, c.UserName, c.Password)
}

func (c *JenkinsClient) CreateFolder(projectOwner, projectName string) (*gojenkins.Folder, error) {
	_, err := c.Jenkins.CreateFolder(projectOwner, "hub")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	projectNameFolder, err := c.Jenkins.CreateFolder(projectName, projectOwner)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return projectNameFolder, nil
}

func (c *JenkinsClient) CreateBranchJob(projectOwner, projectName, branch, accessToken string) {

}

func (c *JenkinsClient) CreateTagJob(projectOwner, projectName, tag, accessToken string) {

}

func (c *JenkinsClient) CreateJob(config, jobName, folder) {

}
