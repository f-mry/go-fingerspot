package fingerspot

type Response interface {
	GetAttLogRequest | GetUserInfoResponse
}

type DefaultResponseField struct {
	Success bool   `json:"success,omitempty"`
	TransID string `json:"trans_id,omitempty"`
}

type attLogData struct {
	Pin      string `json:"pin"`
	ScanDate string `json:"scan_date"`
	//Verifikasi saat user scan di mesin absensi. 1: finger, 2: password, 3:
	//card, 4: face, 5: gps, 6:vein
	Verify int `json:"verify"`
	// Verifikasi saat user scan di mesin absensi. 0: scan in, 1: scan out, 2:
	// break in, 3: break out, 4: overtime in, 5: overtime out, 6: rapat in, 7:
	// rapat out, 8: custome1, 9: custome2
	StatusScan int `json:"status_scan"`
}

type GetAttLogResponse struct {
	DefaultResponseField
	Data []attLogData `json:"data"`
}

type GetUserInfoResponse struct {
	DefaultResponseField
}

type SetUserInfoResponse struct {
	DefaultResponseField
}

type DeleteUserInfoResponse struct {
	DefaultResponseField
}
