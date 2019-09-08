package backblaze

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/packago/config"
)

type AuthorizeAccount struct {
	AbsoluteMinimumPartSize int    `json:"absoluteMinimumPartSize"`
	AccountID               string `json:"accountId"`
	Allowed                 struct {
		BucketID     string      `json:"bucketId"`
		BucketName   string      `json:"bucketName"`
		Capabilities []string    `json:"capabilities"`
		NamePrefix   interface{} `json:"namePrefix"`
	} `json:"allowed"`
	APIURL              string `json:"apiUrl"`
	AuthorizationToken  string `json:"authorizationToken"`
	DownloadURL         string `json:"downloadUrl"`
	RecommendedPartSize int    `json:"recommendedPartSize"`
}

func init() {
	authorizeAccount, err := B2AuthorizeAccount()
	if err != nil {
		panic(err)
	}
	B2UploadURL, err = B2GetUploadURL(authorizeAccount)
	if err != nil {
		panic(err)
	}
}

func B2AuthorizeAccount() (AuthorizeAccount, error) {
	var authorizeAccount AuthorizeAccount
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/b2api/v2/b2_authorize_account", config.File().GetString("backblaze.rootUrl")), nil)
	if err != nil {
		return authorizeAccount, err
	}
	req.SetBasicAuth(config.File().GetString("backblaze.keyID"), config.File().GetString("backblaze.applicationKey"))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return authorizeAccount, err
	}
	defer resp.Body.Close()
	if err = json.NewDecoder(resp.Body).Decode(&authorizeAccount); err != nil {
		return authorizeAccount, err
	}
	return authorizeAccount, nil
}
