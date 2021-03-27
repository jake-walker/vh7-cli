package vh7

type Response struct {
	Created UtcTime `json:"created"`
	Updated UtcTime `json:"updated"`
	Url     Url     `json:"url,omitempty"`
	Paste   Paste   `json:"paste,omitempty"`
	Upload  Upload  `json:"upload,omitempty"`
	Link    string  `json:"link"`
	Expiry  UtcTime `json:"expiry"`
}

func (r *Response) GetType() string {
	if (r.Url != Url{}) {
		return "url"
	} else if (r.Paste != Paste{}) {
		return "paste"
	} else if (r.Upload != Upload{}) {
		return "upload"
	} else {
		return ""
	}
}

func (r *Response) GetSummary() string {
	responseType := r.GetType()

	if responseType == "url" {
		return r.Url.Url
	} else if responseType == "paste" {
		return r.Paste.Language + " code"
	} else if responseType == "upload" {
		return r.Upload.OriginalFilename
	} else {
		return "Unknown"
	}
}

type Url struct {
	Url string `json:"url"`
}

type Paste struct {
	Language string `json:"language"`
	Code     string `json:"code"`
	Hash     string `json:"hash,omitempty"`
}

type Upload struct {
	Mimetype         string `json:"mimetype"`
	OriginalFilename string `json:"original_filename"`
	Hash             string `json:"hash"`
}

type Language struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}
