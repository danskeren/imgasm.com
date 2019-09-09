package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"

	"github.com/danskeren/imgasm.com/db"
	"github.com/gorilla/csrf"
	"github.com/packago/config"
	"github.com/packago/cookie"
)

type CommonData struct {
	Lang            string
	MetaTitle       string
	MetaDescription string
	CSRF            template.HTML
	Flashes         []interface{}
	CurrentURL      template.URL
}

func ReadCommonData(w http.ResponseWriter, r *http.Request) CommonData {
	currentURL := "/"
	if r.URL.Path != "" {
		currentURL = r.URL.Path
		if r.URL.RawQuery != "" {
			currentURL = currentURL + "?" + r.URL.RawQuery
		}
	}
	sess, err := cookie.GetSession(r, config.File().GetString("session.key"))
	if err != nil {
		db.NewCookie()
	}
	flashes := sess.Flashes()
	if err := sess.Save(r, w); err != nil {
		fmt.Println(err)
	}
	return CommonData{
		CSRF:       csrf.TemplateField(r),
		CurrentURL: template.URL(url.QueryEscape(currentURL)),
		Flashes:    flashes,
	}
}
