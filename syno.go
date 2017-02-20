package syno

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Syno struct {
	caller *caller

	Auth *Auth
	DS   *DS
}

func New(base string) (*Syno, error) {
	caller, err := newCaller(base)
	if err != nil {
		return nil, err
	}

	return &Syno{
		caller: caller,

		Auth: &Auth{caller: caller},
		DS:   &DS{caller: caller},
	}, nil
}

func (s *Syno) SetSID(sid string) {
	s.caller.sid = sid
}

type ApiName string

type Params map[string]string

type Query struct {
	Api     ApiName
	Version string
	Path    string
	Method  string
	Params  Params
	SID     string
}

type caller struct {
	sid  string
	base *url.URL
}

func newCaller(base string) (*caller, error) {
	u, err := url.Parse(base)
	if err != nil {
		return nil, err
	}
	return &caller{
		base: u,
	}, nil
}

type synologyResponse struct {
	Success bool
	Error   struct {
		Code int
	}
	Data json.RawMessage
}

func (c *caller) call(q Query, data interface{}) error {
	v := url.Values{}
	v.Add("api", string(q.Api))
	v.Add("version", q.Version)
	v.Add("method", q.Method)

	if q.SID != "" {
		v.Add("_sid", q.SID)
	} else if c.sid != "" {
		v.Add("_sid", c.sid)
	}

	for k, val := range q.Params {
		v.Set(k, val)
	}

	req := &http.Request{
		Method: "GET",
		URL: c.base.ResolveReference(&url.URL{
			Path:     q.Path,
			RawQuery: v.Encode(),
		}),
		Header: make(http.Header),
	}

	dump, _ := httputil.DumpRequest(req, false)
	fmt.Printf("%s\n", dump)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var parsed synologyResponse
	if err := json.NewDecoder(res.Body).Decode(&parsed); err != nil {
		return err
	}
	if !parsed.Success {
		return fmt.Errorf("Error %d", parsed.Error.Code)
	}
	if data != nil {
		if err := json.Unmarshal(parsed.Data, data); err != nil {
			return err
		}
	}
	return nil
}
