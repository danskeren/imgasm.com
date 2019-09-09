package models

type File struct {
	Body      []byte
	MD5Hash   string
	MimeType  string
	Extension string
}

type FileData struct {
	ID             string `json:"id"`
	Filename       string `json:"filename"`
	DeletePassword string `json:"deletePassword"`
}

type FileMD5 struct {
	Processed string `json:"processed"`
	Extension string `json:"extension"`
}
