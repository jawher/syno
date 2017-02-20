package syno

type DSCreateParams struct {
	URI         string
	Destination string
}

func (p DSCreateParams) Params() map[string]string {
	res := map[string]string{
		"uri":         p.URI,
		"destination": p.Destination,
	}
	return res
}

func (d *DS) Create(p DSCreateParams) error {
	return d.caller.call(Query{
		Api:     "SYNO.DownloadStation.Task",
		Version: "3",
		Path:    "DownloadStation/task.cgi",
		Method:  "create",
		Params:  p.Params(),
	}, nil)
}
