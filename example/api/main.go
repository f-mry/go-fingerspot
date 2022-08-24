package main

import (
	"fmt"
	"log"

	"github.com/farhanmry/go-fingerspot"
)

func main() {
	cl := fingerspot.NewClient("ISQI95BC26PI0TJU")

	req := fingerspot.GetUserInfoRequest{}
	req.TransID = "1"
	req.CloudID = "C269784293101835"
	req.Pin = "1"

	user, err := cl.GetUserInfo(req)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("User: %+v\n", user)

	req.TransID = "1"
	req.CloudID = "C269784293132433"
	req.Pin = "1"
	user, err = cl.GetUserInfo(req)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("User: %+v\n", user)

	attReq := fingerspot.GetAttLogRequest{
		DefaultField: fingerspot.DefaultField{
			TransID: "1",
			CloudID: "C269784293132433",
		},
		StartDate: "2022-08-01",
		EndDate:   "2022-08-03",
	}

	attLog, err := cl.GetAttLog(attReq)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Log: %+v\n", attLog)
}
