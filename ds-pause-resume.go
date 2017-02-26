package syno

import "strings"

type DSPauseResumeParams struct {
	Pause bool
	IDs   []string
}

type DSPauseResumeResult []DSPauseResumeTaskResult

type DSPauseResumeTaskResult struct {
	ID    string
	Error int
}

func (p DSPauseResumeParams) Params() map[string]string {
	res := map[string]string{
		"id": strings.Join(p.IDs, ","),
	}
	return res
}

func (d *DS) PauseResume(p DSPauseResumeParams) (res DSPauseResumeResult, err error) {
	method := "resume"
	if p.Pause {
		method = "pause"
	}
	err = d.caller.call(Query{
		Api:     "SYNO.DownloadStation.Task",
		Version: "1",
		Path:    "DownloadStation/task.cgi",
		Method:  method,
		Params:  p.Params(),
	}, &res)
	return
}
