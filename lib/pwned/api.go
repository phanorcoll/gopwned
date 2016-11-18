package pwnedapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"

	emailvalidation "bitbucket.com/phanorcoll/clipwned/lib/emailValidation"
)

var URL_API string = "https://haveibeenpwned.com/api/v2/breachedaccount/"

type PwnedData struct {
	Title      string `json:"Title"`
	BreachDate string `json:"BreachDate"`
	Domain     string `json:"Domain"`
}

func GetEmail(e string) {
	if e != "" {
		if !emailvalidation.Validate(e) {
			fmt.Printf("the email %v is not valid, please verify and try again \n", e)
		} else {
			getApiData(e)
		}
	} else {
		fmt.Println("You must specify an email account [ USAGE -> gopwned verify <email> ]")
	}
}

func getApiData(e string) {
	req, err := http.NewRequest("GET", URL_API+e, nil)
	if err != nil {
		log.Fatal("Error getting the API: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var records []PwnedData

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

	fmt.Printf("********************* Breaches where the email %v was found ********************* \n", e)
	red := color.New(color.FgRed).PrintfFunc()
	for _, breach := range records {
		color.Blue("Company\n")
		red("%v -  Domain %v \n\n", breach.Title, breach.Domain)
	}
	fmt.Printf("You can get detail information using gopwned verify user@example.com --domain adobe.com \n")
	fmt.Printf("*********************           End of Breaches             ********************* \n")
}
