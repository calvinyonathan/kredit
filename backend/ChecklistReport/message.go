package ChecklistReport

type PpkRequest struct {
	Ppk string `json:"ppk"`
}

type GetSearchRequest struct {
	Branch    string `json:"branch"`
	Company   string `json:"company"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}
