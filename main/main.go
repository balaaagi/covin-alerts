package main

import (
	"encoding/json"
	"fmt"
	"github.com/martinlindhe/notify"
	_ "github.com/martinlindhe/notify"
	"github.com/robfig/cron/v3"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

type CovidCentresResponse struct {
	Centers []struct {
		CenterID     int    `json:"center_id"`
		Name         string `json:"name"`
		StateName    string `json:"state_name"`
		DistrictName string `json:"district_name"`
		BlockName    string `json:"block_name"`
		Pincode      int    `json:"pincode"`
		Lat          int    `json:"lat"`
		Long         int    `json:"long"`
		From         string `json:"from"`
		To           string `json:"to"`
		FeeType      string `json:"fee_type"`
		Sessions     []struct {
			SessionID         string   `json:"session_id"`
			Date              string   `json:"date"`
			AvailableCapacity int      `json:"available_capacity"`
			MinAgeLimit       int      `json:"min_age_limit"`
			Vaccine           string   `json:"vaccine"`
			Slots             []string `json:"slots"`
		} `json:"sessions"`
	} `json:"centers"`
}

func main() {
	print(len(os.Args))
	print(os.Args)
	if len(os.Args) > 2 {
		covinArgs := os.Args[0:]
		searchBy, _ := strconv.Atoi(covinArgs[1])
		date:= covinArgs[2]
		minAge, _ := strconv.Atoi(covinArgs[3])
		thirdParam:= covinArgs[4]
		cronScheduler := cron.New()

		if searchBy == 1 {
			cronScheduler.AddFunc("@every 15m", func() { findSlotsByPinCode(date,thirdParam,minAge) })

		}else{
			cronScheduler.AddFunc("@every 15s", func() { findSlotsByDistrictId(date,thirdParam,minAge) })

		}
		cronScheduler.Start()

	}else{
			print("Please give Arguments of this format - covin ")
	}


	select {}
}

func findSlotsByDistrictId(date string, districtId string, minAge int) {
	getDetailsFromCoWin("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByDistrict?district_id="+districtId+"&date="+date,minAge)
}

func findSlotsByPinCode(date string, pincode string, minAge int) {
	getDetailsFromCoWin("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode="+pincode+"&date="+date,minAge)
}

func getDetailsFromCoWin(url string,minAge int)  {
	coWinResponse := CovidCentresResponse{}
	var DefaultClient = &http.Client{}

	response, err := DefaultClient.Get(url)
	if err!=nil {
		fmt.Println("Error No Response from CoWin")
	}else{

		responseData, err1 := ioutil.ReadAll(response.Body)
		if err1==nil {
			err3 :=json.Unmarshal(responseData,&coWinResponse)
			if err3 ==nil{
				processResponseAndAlertIfPresent(coWinResponse,minAge)

			}else{

				fmt.Println("Error While Processing Cowin")
			}
		}else{
			fmt.Println("Error While Processing CoWin")
		}

	}
	
}

func processResponseAndAlertIfPresent(response CovidCentresResponse, age int) {
	totalNoOfCenters:=len(response.Centers)
	if totalNoOfCenters>0 {
		for i := 0; i < totalNoOfCenters; i++{
			var center = response.Centers[i]
			var noOfSessionsInCenter=len(center.Sessions)
			if noOfSessionsInCenter>0 {
				for j:=0;j<noOfSessionsInCenter;j++{

					var minimumAgeLimit=center.Sessions[j].MinAgeLimit
					var availablity=center.Sessions[j].AvailableCapacity
					if minimumAgeLimit== minimumAgeLimit && availablity>5 {
						alertMessage:= fmt.Sprintf("Center: %s , Date: %s , Slot Info: %s Count: %d", center.Name, center.Sessions[j].Date, center.Sessions[j].Slots,center.Sessions[j].AvailableCapacity)
						fmt.Printf("Center Name %s \n", center.Name)
						fmt.Printf("Date %s \n", center.Sessions[j].Date)
						fmt.Printf("Slot Information %s",center.Sessions[j].Slots)
						fmt.Printf("Count Available: %d Minimum Age: %d",center.Sessions[j].AvailableCapacity,minimumAgeLimit)
						notify.Alert("Slot Alert", "Slot Available", alertMessage, "https://png.pngtree.com/element_pic/17/04/07/628c04fea84856c8d04b3878eb989009.jpg")
					}
				}
			}
		}
	}else{
		fmt.Println("No Centers Found")
	}
}
