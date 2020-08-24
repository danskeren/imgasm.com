package backblaze

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/packago/config"
	"imgasm.com/models"
)

type UploadURL struct {
	AuthorizationToken string `json:"authorizationToken"`
	BucketID           string `json:"bucketId"`
	UploadURL          string `json:"uploadUrl"`
}

func B2GetUploadURL(authorizeAccount AuthorizeAccount) (UploadURL, error) {
	var b2UploadURL UploadURL
	body := strings.NewReader(fmt.Sprintf("{\"bucketId\": \"%s\"}", authorizeAccount.Allowed.BucketID))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/b2api/v2/b2_get_upload_url", config.File().GetString("backblaze.rootUrl")), body)
	if err != nil {
		return b2UploadURL, err
	}
	req.Header.Set("Authorization", authorizeAccount.AuthorizationToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return b2UploadURL, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&b2UploadURL); err != nil {
		return b2UploadURL, err
	}
	return b2UploadURL, nil
}

func B2Upload(file models.File) error {
	authorizeAccount, err := B2AuthorizeAccount()
	if err != nil {
		return err
	}
	uploadURL, err := B2GetUploadURL(authorizeAccount)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", uploadURL.UploadURL, bytes.NewReader(file.Body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", uploadURL.AuthorizationToken)
	req.Header.Set("X-Bz-File-Name", fmt.Sprintf("%x.%s", file.MD5Hash, file.Extension))
	req.Header.Set("Content-Type", file.MimeType)
	req.Header.Set("X-Bz-Content-Sha1", fmt.Sprintf("%x", sha1.Sum(file.Body)))
	req.Header.Set("X-Bz-Info-Author", "unknown")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("returned status code %d", resp.StatusCode))
	}
	return nil
}
