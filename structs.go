package assignment1

type University struct {
	Isocode  string   `json:"alpha_two_code"`
	WebPages []string `json:"web_pages"`
	Name     string   `json:"name"`
	Country  string   `json:"country"`
}

type Countries struct {
	Isocode   string            `json:"cca2"`
	Cca3      string            `json:"cca3"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"maps"`
}

type Diag struct {
	Universitiesapi string `json:"universities_api"`
	Countriespai    string `json:"countries_api"`
	Version         string `json:"version"`
	Uptime          string `json:"uptime"`
}

type Response struct {
	Name      string            `json:"name"`
	Country   string            `json:"country"`
	Isocode   string            `json:"isocode"`
	Webpages  []string          `json:"webpages"`
	Languages map[string]string `json:"languages"`
	Maps      map[string]string `json:"map"`
}
