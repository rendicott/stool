package inspec

type InspecOutputModel struct {
	Version  string `json:"version"`
	Controls []struct {
		Status    string  `json:"status,omitempty"`
		CodeDesc  string  `json:"code_desc"`
		RunTime   float64 `json:"run_time"`
		StartTime string  `json:"start_time"`
		Message   string  `json:"message,omitempty"`
	} `json:"controls"`
	OtherChecks []interface{} `json:"other_checks"`
	Profiles    []struct {
		Name           string        `json:"name"`
		Title          string        `json:"title"`
		Maintainer     string        `json:"maintainer"`
		Copyright      string        `json:"copyright"`
		CopyrightEmail string        `json:"copyright_email"`
		License        string        `json:"license"`
		Summary        string        `json:"summary"`
		Version        string        `json:"version"`
		Supports       []interface{} `json:"supports"`
		Controls       []struct {
			Title  string        `json:"title"`
			Desc   string        `json:"desc,omitempty"`
			Impact float64       `json:"impact"`
			Refs   []interface{} `json:"refs"`
			Tags   struct {
			} `json:"tags"`
			Code           string `json:"code"`
			SourceLocation struct {
				Ref  string `json:"ref"`
				Line int    `json:"line"`
			} `json:"source_location"`
			ID      string `json:"id"`
			Results []struct {
				Status    string  `json:"status,omitempty"`
				CodeDesc  string  `json:"code_desc,omitempty"`
				RunTime   float64 `json:"run_time"`
				StartTime string  `json:"start_time"`
				Message   string  `json:"message"`
			} `json:"results"`
		} `json:"controls"`
		Groups []struct {
			Title    string   `json:"title"`
			Controls []string `json:"controls"`
			ID       string   `json:"id"`
		} `json:"groups"`
		Attributes []interface{} `json:"attributes"`
		Sha256     string        `json:"sha256"`
	} `json:"profiles"`
	Platform struct {
		Name    string `json:"name"`
		Release string `json:"release"`
	} `json:"platform"`
	Statistics struct {
		Duration float64 `json:"duration"`
	} `json:"statistics"`
}

type Controls struct {
	Status    string  `json:"status,omitempty"`
	CodeDesc  string  `json:"code_desc"`
	RunTime   float64 `json:"run_time"`
	StartTime string  `json:"start_time"`
	Message   string  `json:"message,omitempty"`
}
