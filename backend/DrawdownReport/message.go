package DrawdownReport

type GetSearchRequest struct {
	Branch    string `json:"branch"`
	Company   string `json:"company"`
	StartDate string `json:"startdate"`
	EndDate   string `json:"enddate"`
}
