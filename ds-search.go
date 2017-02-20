package syno

import "strconv"

type DSSearchStartParams struct {
	Keyword string
}

type DSSearchStartResult struct {
	SearchID string `json:"taskid"`
}

func (p DSSearchStartParams) Params() map[string]string {
	res := map[string]string{
		"keyword": p.Keyword,
		"module":  "all",
	}
	return res
}

func (d *DS) SearchStart(p DSSearchStartParams) (res DSSearchStartResult, err error) {
	err = d.caller.call(Query{
		Api:     "SYNO.DownloadStation.BTSearch",
		Version: "1",
		Path:    "DownloadStation/btsearch.cgi",
		Method:  "start",
		Params:  p.Params(),
	}, &res)
	return
}

type DSSearchListParams struct {
	SearchID       string
	Offset         int
	Limit          int
	SortBy         DSSortBy
	SortDirection  DSSortDir
	FilterTitle    string
	FilterCategory string
}

type DSSortBy string

const (
	DS_SORT_BY_TITLE    DSSortBy = "title"
	DS_SORT_BY_SIZE     DSSortBy = "size"
	DS_SORT_BY_DATE     DSSortBy = "date"
	DS_SORT_BY_PEERS    DSSortBy = "peers"
	DS_SORT_BY_PROVIDER DSSortBy = "provider"
	DS_SORT_BY_SEEDS    DSSortBy = "seeds"
	DS_SORT_BY_LEECHS   DSSortBy = "leechs"
)

type DSSortDir string

const (
	DS_SORT_DIR_ASC  DSSortDir = "ASC"
	DS_SORT_DIR_DESC DSSortDir = "DESC"
)

type DSSearchListResult struct {
	SearchID string `json:"taskid"`
	Finished bool
	Total    int
	Items    []DSSearchResult
}

func (p DSSearchListParams) Params() map[string]string {
	if p.Limit == 0 {
		p.Limit = -1
	}
	return map[string]string{
		"taskid":          p.SearchID,
		"offset":          strconv.Itoa(p.Offset),
		"limit":           strconv.Itoa(p.Limit),
		"sort_by":         string(p.SortBy),
		"sort_direction":  string(p.SortDirection),
		"filter_category": p.FilterCategory,
		"filter_title":    p.FilterTitle,
	}
}

func (d *DS) SearchList(p DSSearchListParams) (res DSSearchListResult, err error) {
	err = d.caller.call(Query{
		Api:     "SYNO.DownloadStation.BTSearch",
		Version: "1",
		Path:    "DownloadStation/btsearch.cgi",
		Method:  "list",
		Params:  p.Params(),
	}, &res)
	return
}

type DSSearchResult struct {
	Date         string `json:"date"`
	DownloadURI  string `json:"download_uri"`
	ExternalLink string `json:"external_link"`
	ID           int    `json:"id"`
	Leechs       int    `json:"leechs"`
	ModuleID     string `json:"module_id"`
	ModuleTitle  string `json:"module_title"`
	Peers        int    `json:"peers"`
	Seeds        int    `json:"seeds"`
	Size         string `json:"size"`
	Title        string `json:"title"`
}
