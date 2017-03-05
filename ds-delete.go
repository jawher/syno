package syno

import (
	"strings"
	"strconv"
)

type DSDeleteParams struct {
	IDs   []string
	ForceComplete bool
}

type DSDeleteResult []DSDeleteTaskResult

type DSDeleteTaskResult struct {
	ID    string
	Error int
}

func (p DSDeleteParams) Params() map[string]string {
	res := map[string]string{
		"id": strings.Join(p.IDs, ","),
		"force_complete": strconv.FormatBool(p.ForceComplete),
	}
	return res
}

func (d *DS) Delete(p DSDeleteParams) (res DSDeleteResult, err error) {
	err = d.caller.call(Query{
		Api:     "SYNO.DownloadStation.Task",
		Version: "1",
		Path:    "DownloadStation/task.cgi",
		Method:  "delete",
		Params:  p.Params(),
	}, &res)
	return
}
