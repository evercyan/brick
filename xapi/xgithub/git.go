package xgithub

// Client ...
type Client struct {
	Repo        string `json:"repo" `
	Owner       string `json:"owner"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

// New ...
func New(repo string, owner string, email string, accessToken string) *Client {
	return &Client{
		Repo:        repo,
		Owner:       owner,
		Email:       email,
		AccessToken: accessToken,
	}
}
