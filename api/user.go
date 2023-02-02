package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ehsanghaffar/github-stats-golang/utils"
)

type User struct {
	Login             string      `json:"login"`
	ID                int64       `json:"id"`
	NodeID            string      `json:"node_id"`
	AvatarURL         string      `json:"avatar_url"`
	GravatarID        string      `json:"gravatar_id"`
	URL               string      `json:"url"`
	HTMLURL           string      `json:"html_url"`
	FollowersURL      string      `json:"followers_url"`
	FollowingURL      string      `json:"following_url"`
	GistsURL          string      `json:"gists_url"`
	StarredURL        string      `json:"starred_url"`
	SubscriptionsURL  string      `json:"subscriptions_url"`
	OrganizationsURL  string      `json:"organizations_url"`
	ReposURL          string      `json:"repos_url"`
	EventsURL         string      `json:"events_url"`
	ReceivedEventsURL string      `json:"received_events_url"`
	Type              string      `json:"type"`
	SiteAdmin         bool        `json:"site_admin"`
	Name              string      `json:"name"`
	Company           string      `json:"company"`
	Blog              string      `json:"blog"`
	Location          string      `json:"location"`
	Email             interface{} `json:"email"`
	Hireable          interface{} `json:"hireable"`
	Bio               string      `json:"bio"`
	TwitterUsername   string      `json:"twitter_username"`
	PublicRepos       int64       `json:"public_repos"`
	PublicGists       int64       `json:"public_gists"`
	Followers         int64       `json:"followers"`
	Following         int64       `json:"following"`
	CreatedAt         string      `json:"created_at"`
	UpdatedAt         string      `json:"updated_at"`
}

// Get Github user data user using Github REST-API
func GetUserData(username string) User {

	res, dataErr := http.Get("https://api.github.com/users/" + username)
	if dataErr != nil {
		utils.Perror(dataErr)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		utils.Perror(err)
	}

	var data User
	decodeErr := json.Unmarshal(body, &data)
	if decodeErr != nil {
		utils.Perror(err)
	}
	return data
}
