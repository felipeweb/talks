m.Use(oauth2.Google(
    &goauth2.Config{
        ClientID:     "client_id",
        ClientSecret: "client_secret",
        Scopes:       []string{"https://www.googleapis.com/auth/drive"},
        RedirectURL:  "redirect_url",
    },
))
// URLs usadas pelo middleware
const (
	KEY_TOKEN     = "oauth2_token"
	KEY_NEXT_PAGE = "next"
)

var (
	// PathLogin is the path to handle OAuth 2.0 logins.
	PathLogin = "/login"
	// PathLogout is the path to handle OAuth 2.0 logouts.
	PathLogout = "/logout"
	// PathCallback is the path to handle callback from OAuth 2.0 backend
	// to exchange credentials.
	PathCallback = "/oauth2callback"
	// PathError is the path to handle error cases.
	PathError = "/oauth2error"
)
