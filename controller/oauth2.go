package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"tpler/common"

	"github.com/ohko/hst"
	"golang.org/x/oauth2"
)

// Oauth2Controller 默认主页控制器
type Oauth2Controller struct {
	controller
}

var (
	oauthStateString = "random"
	oauthServerHost  = ""
	oauthRedirectURL = "/oauth2/callback"
	oauthConfig      *oauth2.Config
)

func oauth2Init(o2srv string) {
	oauthServerHost = o2srv
	oauthConfig = &oauth2.Config{
		ClientID:     "your_client_id",
		ClientSecret: "your_client_secret",
		RedirectURL:  oauthRedirectURL,
		Scopes:       []string{"(no scope)"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  oauthServerHost + "/oauth2/auth",
			TokenURL: oauthServerHost + "/oauth2/token",
		},
	}
}

// Login 跳转oauth2登录授权页面
func (o *Oauth2Controller) Login(ctx *hst.Context) {
	if ctx.R.TLS == nil {
		oauthConfig.RedirectURL = "http://" + ctx.R.Host + oauthRedirectURL + "?callback=" + url.QueryEscape(ctx.R.FormValue("callback"))
	} else {
		oauthConfig.RedirectURL = "https://" + ctx.R.Host + oauthRedirectURL + "?callback=" + url.QueryEscape(ctx.R.FormValue("callback"))
	}
	uri := oauthConfig.AuthCodeURL(oauthStateString)
	http.Redirect(ctx.W, ctx.R, uri, http.StatusTemporaryRedirect)
}

// Callback oauth2登录授权返回
func (o *Oauth2Controller) Callback(ctx *hst.Context) {
	common.LL.Log0Debug("Callback:", ctx.R.Method, ctx.R.RequestURI, ctx.R.Form.Encode())
	state := ctx.R.FormValue("state")
	if state != oauthStateString {
		ctx.HTML2(200, "layout/empty.html", fmt.Sprintf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state), "admin/oauth2error.html")
		return
	}

	errorMsg := ctx.R.FormValue("error")
	if errorMsg != "" {
		ctx.HTML2(200, "layout/empty.html", errorMsg, "admin/oauth2error.html")
	}

	code := ctx.R.FormValue("code")
	token, err := oauthConfig.Exchange(oauth2.NoContext, code)
	common.LL.Log0Debug("Token:", fmt.Sprintf("%#v", token))
	if err != nil {
		ctx.Data(200, err.Error())
		return
	}

	response, err := http.Get(oauthServerHost + "/oauth2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		ctx.Data(200, err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.Data(200, err.Error())
	}
	common.LL.Log0Debug("Userinfo:", string(contents))

	var rst struct {
		No   int         `json:"no"`
		Data interface{} `json:"data"`
	}
	if err := json.Unmarshal(contents, &rst); err != nil {
		ctx.Data(200, err.Error())
	}

	u := rst.Data.(map[string]interface{})
	ctx.SessionSet("Member", u["User"].(string))

	if ctx.R.FormValue("callback") != "" {
		http.Redirect(ctx.W, ctx.R, ctx.R.FormValue("callback"), http.StatusFound)
		return
	}
	http.Redirect(ctx.W, ctx.R, "/admin/", http.StatusFound)
}
