package pwned

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"

	emailvalidation "bitbucket.com/phanorcoll/clipwned/lib/emailValidation"

	"github.com/fatih/color"
)

type PwnedData struct {
	Domain string `json:"Domain"`
}

const URL_API string = "https://haveibeenpwned.com/api/v2/breachedaccount/"

/**
 * verifies that the email is well formatted
 */
func GetEmail(e string) {
	ErrorMessage := color.New(color.Bold, color.FgRed).PrintlnFunc()
	if e != "" {
		if !emailvalidation.Validate(e) {
			ErrorMessage("\n\nThe email [ " + e + " ] is not valid, please verify and try again! \n\n")
		} else {
			getApiData(e)
		}
	} else {
		ErrorMessage("\n\nYou must specify an email account, run gopwned -h for more information. \n\n")
	}
}

/**
 * gets the data from the API and returns the content to the users
 */
func getApiData(e string) {
	noBreaches := color.New(color.Bold, color.FgGreen).PrintlnFunc()
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

	var breaches []PwnedData

	if err := json.NewDecoder(resp.Body).Decode(&breaches); err != nil {
		log.Println(err)
	}

	if len(breaches) > 0 {

		breachTemplate := template.Must(template.New("breachMessage").Parse(breachTmpl))
		if err := breachTemplate.Execute(os.Stdout, breaches); err != nil {
			panic(err)
		}

	} else {
		noBreaches("\n\nThe email [ " + e + " ] is safe for now, update your passwords often! \n\n")
	}
}
