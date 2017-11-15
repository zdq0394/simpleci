package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"github.com/zdq0394/simpleci/simpleci/config"
	"github.com/zdq0394/simpleci/simpleci/githubclient"
)

const (
	GithubAuthorizeURL   = "https://github.com/login/oauth/authorize"
	GithubAccessTokenURL = "https://github.com/login/oauth/access_token"
	AuthScope            = "repo,user:email,admin:repo_hook"
)

type CIService struct {
	Conf *config.Config
}

func NewCIService(conf *config.Config) *CIService {
	return &CIService{
		Conf: conf,
	}
}

func (s *CIService) ping(c *gin.Context) {
	c.String(http.StatusOK, "%s", "pong")
}

func (s *CIService) authurlHanlder(c *gin.Context) {
	authorizeURL := fmt.Sprintf("%s?redirect_uri=%s&client_id=%s&scope=%s",
		GithubAuthorizeURL, s.Conf.Github.AuthRedirectURL, s.Conf.Github.ClientID, AuthScope)
	//respText := fmt.Sprintf("<html><head><title>auth</title></head><body><a href=\"%s\">Click here</a> to begin!</a></body></html>", authorizeURL)
	c.String(http.StatusOK, authorizeURL)
}

func (s *CIService) codeGotHandler(c *gin.Context) {
	fmt.Println("codeGotHandler....")
	code := c.Query("code")
	fmt.Println(code)
	s.httpPost(code)
}

func (s *CIService) httpPost(code string) {
	params := fmt.Sprintf("client_id=%s&client_secret=%s&code=%s", s.Conf.Github.ClientID, s.Conf.Github.ClientSecret, code)
	resp, err := http.Post(GithubAccessTokenURL,
		"application/x-www-form-urlencoded",
		strings.NewReader(params))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	respText := string(body)
	fmt.Println(respText)
	accessTokenStr := strings.Split(respText, "&")[0]
	AccessToken = strings.Split(accessTokenStr, "=")[1]
	fmt.Println(AccessToken)
}

func (s *CIService) getAccessTokenOfUser(user string) string {
	return AccessToken
}

func (s *CIService) createHookHandler(c *gin.Context) {
	ctx := context.Background()
	owner := c.Param("owner")
	repo := c.Param("repo")
	accessToken := s.getAccessTokenOfUser(owner)
	client := githubclient.GetClient(accessToken)
	// v := new(github.Hook)
	// json.NewDecoder(c.Request.Body).Decode(v)
	var v github.Hook
	v.Events = []string{"push"}
	jenkinsURL := fmt.Sprintf("%s/github-webhook/", s.Conf.Jenkins.Server)
	v.URL = &jenkinsURL
	name := "web"
	v.Name = &name
	v.Config = map[string]interface{}{
		"content_type": "json",
		"url":          jenkinsURL,
	}
	hook, _, err := client.Repositories.CreateHook(ctx, owner, repo, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hook.GetName())
}
