package fingerspot

type DefaultMessage struct {
	Type    string `json:"type"`
	CloudID string `json:"cloud_id"`
	TransID int    `json:"trans_id,omitempty"`
}

type RealtimeAttLogMessage struct {
	DefaultMessage
	Data struct {
		Pin  string `json:"pin"`
		Scan string `json:"scan"`
		//Verifikasi saat user scan di mesin absensi. 1: finger, 2: password, 3:
		//card, 4: face, 5: gps, 6:vein
		Verify int `json:"verify"`
		// Verifikasi saat user scan di mesin absensi. 0: scan in, 1: scan out, 2:
		// break in, 3: break out, 4: overtime in, 5: overtime out, 6: rapat in, 7:
		// rapat out, 8: custome1, 9: custome2
		StatusScan int `json:"status_scan"`
	} `json:"data"`
}

type GetUserInfoMessage struct {
	DefaultMessage
	Data struct {
		Pin       string `json:"pin"`
		Name      string `json:"name"`
		Password  string `json:"password"`
		Privilege int    `json:"privilege"`
		Finger    int    `json:"finger"`
		Face      int    `json:"face"`
		Rfid      int    `json:"rfid"`
		Vein      int    `json:"vein"`
		Template  string `json:"template"`
	} `json:"data"`
}

type SetUserInfoMessage struct {
	DefaultMessage
	Data struct {
		//Status respon dari mesin.
		//1: user info berhasil ditambahkan ke mesin,
		//2: user info gagal ditambahkan
		Status int `json:"status"`
	} `json:"data"`
}

type DeleteUserInfoMessage struct {
	DefaultMessage
	Data struct {
		//Status respon dari mesin.
		//1: user info berhasil ditambahkan ke mesin,
		//2: user info gagal ditambahkan
		Status int `json:"status"`
	} `json:"data"`
}
