package model

type HTTP struct {
	Hostname      string `json:"hostname"`
	URL           string `json:"url"`
	HTTPUserAgent string `json:"http_user_agent"`
	HTTPMethod    string `json:"http_method"`
	Protocol      string `json:"protocol"`
	Status        int    `json:"status"`
	Length        int    `json:"length"`
}
