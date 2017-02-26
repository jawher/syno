package syno

import (
	"strconv"
	"strings"
)

type DSListParams struct {
	Offset   int
	Limit    int
	Detail   bool
	Transfer bool
	File     bool
	Tracker  bool
	Peer     bool
}

type DSListResult struct {
	Total  int
	Offset int
	Tasks  []DSTask
}

func (p DSListParams) Params() map[string]string {
	res := map[string]string{
		"offset": strconv.Itoa(p.Offset),
		"limit":  strconv.Itoa(p.Limit),
	}

	additional := []string{}
	if p.Detail {
		additional = append(additional, "detail")
	}
	if p.Transfer {
		additional = append(additional, "transfer")
	}
	if p.File {
		additional = append(additional, "file")
	}
	if p.Tracker {
		additional = append(additional, "tracker")
	}
	if p.Peer {
		additional = append(additional, "peer")
	}

	a := strings.Join(additional, ",")
	if a != "" {
		res["additional"] = a
	}
	return res
}

func (d *DS) List(p DSListParams) (res DSListResult, err error) {
	err = d.caller.call(Query{
		Api:     "SYNO.DownloadStation.Task",
		Version: "1",
		Path:    "DownloadStation/task.cgi",
		Method:  "list",
		Params:  p.Params(),
	}, &res)
	return
}

type DSTask struct {
	Additional *DSTaskAdditional `json:"additional"`
	ID         string            `json:"id"`
	Size       int64             `json:"size"`
	Status     DSTaskStatus      `json:"status"`
	Title      string            `json:"title"`
	Type       string            `json:"type"`
	Username   string            `json:"username"`
}

type DSTaskStatus string

const (
	DS_TASK_DOWNLOADING         DSTaskStatus = "downloading"
	DS_TASK_PAUSED              DSTaskStatus = "paused"
	DS_TASK_SEEDING             DSTaskStatus = "seeding"
	DS_TASK_WAITING             DSTaskStatus = "waiting"
	DS_TASK_ERROR               DSTaskStatus = "error"
	DS_TASK_FINISHED            DSTaskStatus = "finished"
	DS_TASK_FINISHING           DSTaskStatus = "finishing"
	DS_TASK_HASH_CHECKING       DSTaskStatus = "hash_checking"
	DS_TASK_EXTRACTING          DSTaskStatus = "extracting"
	DS_TASK_FILEHOSTING_WAITING DSTaskStatus = "filehosting_waiting"
)

type DSTaskAdditional struct {
	Detail   *DSTaskDetail   `json:"detail"`
	File     []DSTaskFile    `json:"file"`
	Peer     []DSPeer        `json:"peer"`
	Tracker  []DSTracker     `json:"tracker"`
	Transfer *DSTaskTransfer `json:"transfer"`
}

type DSTaskDetail struct {
	CompletedTime     int    `json:"completed_time"`
	ConnectedLeechers int    `json:"connected_leechers"`
	ConnectedPeers    int    `json:"connected_peers"`
	ConnectedSeeders  int    `json:"connected_seeders"`
	CreateTime        int64  `json:"create_time"`
	Destination       string `json:"destination"`
	Seedelapsed       int    `json:"seedelapsed"`
	StartedTime       int64  `json:"started_time"`
	TotalPeers        int    `json:"total_peers"`
	TotalPieces       int    `json:"total_pieces"`
	UnzipPassword     string `json:"unzip_password"`
	URI               string `json:"uri"`
	WaitingSeconds    int    `json:"waiting_seconds"`
}

type DSTaskFile struct {
	Filename       string `json:"filename"`
	Index          int    `json:"index"`
	Priority       string `json:"priority"`
	Size           int    `json:"size"`
	SizeDownloaded int    `json:"size_downloaded"`
	Wanted         bool   `json:"wanted"`
}

type DSTracker struct {
	Peers       int    `json:"peers"`
	Seeds       int    `json:"seeds"`
	Status      string `json:"status"`
	UpdateTimer int    `json:"update_timer"`
	URL         string `json:"url"`
}

type DSTaskTransfer struct {
	DownloadedPieces int   `json:"downloaded_pieces"`
	SizeDownloaded   int64 `json:"size_downloaded"`
	SizeUploaded     int64 `json:"size_uploaded"`
	SpeedDownload    int   `json:"speed_download"`
	SpeedUpload      int   `json:"speed_upload"`
}

type DSPeer struct {
	Address       string  `json:"address"`
	Agent         string  `json:"agent"`
	Progress      float64 `json:"progress"`
	SpeedDownload int     `json:"speed_download"`
	SpeedUpload   int     `json:"speed_upload"`
}
