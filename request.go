package fingerspot

type DefaultField struct {
	TransID string `json:"trans_id,omitempty"`
	CloudID string `json:"cloud_id,omitempty"`
}

type GetAttLogRequest struct {
	DefaultField
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type GetUserInfoRequest struct {
	DefaultField
	Pin string `json:"pin,omitempty"`
}

type setUserData struct {
	Pin       string `json:"pin"`
	Name      string `json:"name"`
	Privilege string `json:"privilege"`
	Password  string `json:"password"`
	Template  string `json:"template"`
}

type SetUserInfoRequest struct {
	DefaultField
	Data setUserData `json:"data"`
}
