package pwnedapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	raven "github.com/getsentry/raven-go"

	emailvalidation "bitbucket.com/phanorcoll/clipwned/lib/emailValidation"
)

var URL_API string = "https://haveibeenpwned.com/api/v2/breachedaccount/"

type PwnedData struct {
	Title      string `json:"Title"`
	BreachDate string `json:"BreachDate"`
	Domain     string `json:"Domain"`
}

func GetEmail(e string) {
	red := color.New(color.FgRed).PrintfFunc()
	if e != "" {
		if !emailvalidation.Validate(e) {
			fmt.Printf("the email %v is not valid, please verify and try again \n", e)
		} else {
			getApiData(e)
		}
	} else {
		red("You must specify an email account [ USAGE -> gopwned verify <email> ]\n")
	}
}

func getApiData(e string) {
	req, err := http.NewRequest("GET", URL_API+e, nil)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal("Error getting the API: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()

	var records []PwnedData

	if err := json.NewDecoder(resp.Body).Decode(&records); err != nil {
		log.Println(err)
	}

	c := color.New(color.FgRed).Add(color.Underline).Add(color.Bold)
	c.Printf("\nBreaches for %v : \n\n", e)

	const padding = 10
	w := tabwriter.NewWriter(os.Stdout, 0, 0, padding, ' ', tabwriter.DiscardEmptyColumns)
	fmt.Fprintf(w, "%v\t%v\n", color.RedString("Company"), color.RedString("Domain"))

	for _, breach := range records {
		fmt.Fprintln(w, "-"+color.WhiteString(breach.Title)+"\t"+" -"+color.WhiteString(breach.Domain))
	}
	w.Flush()

	notice := color.New(color.Bold, color.FgRed).PrintlnFunc()
	notice("\n\nTIP: You can get detail information using -> gopwned verify user@example.com --domain adobe.com \n\n")
}
