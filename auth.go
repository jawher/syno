package syno

type SessionFormat string

const (
	SessionSid    SessionFormat = "sid"
	SessionCookie SessionFormat = "cookie"
)

type LoginParams struct {
	Login    string
	Password string
	Session  string
	Format   SessionFormat
}

func (a LoginParams) Params() map[string]string {
	return map[string]string{
		"account": a.Login,
		"passwd":  a.Password,
	}
}

type LoginResult struct {
	SID string `json:"sid"`
}

type LogoutParams struct {
	Session string
}

type Auth struct {
	caller *caller
}

func (a Auth) Login(p LoginParams) (res LoginResult, err error) {
	err = a.caller.call(Query{
		Api:     "SYNO.API.Auth",
		Version: "2",
		Path:    "auth.cgi",
		Method:  "login",
		Params:  p.Params(),
	}, &res)
	return
}

func (a Auth) Logout(p LogoutParams) error {
	return nil
}
