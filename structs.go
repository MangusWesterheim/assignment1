package assignment1

type University struct {
	isocode  string   `json:"alpha_two_code"`
	webPages []string `json:"web_pages"`
	name     string   `json:"name"`
	country  string   `json:"country"`
}

type Countries struct {
	name      string            `json:"name"`
	isocode   string            `json:"isocode"`
	languages map[string]string `json:"languages"`
	maps      string            `json:"maps"`
}
