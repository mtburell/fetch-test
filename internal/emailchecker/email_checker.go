package emailchecker

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"fetch-test/internal/server"

	"github.com/gorilla/mux"
	"github.com/pquerna/ffjson/ffjson"
)

// From https://www.w3.org/TR/2016/REC-html51-20161101/sec-forms.html#email-state-typeemail
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type RequestBody struct {
	Emails []string `json:"emails"`
}

type ResponseBody struct {
	Targets int `json:"unique_email_addresses"`
}

func RegisterRoutes(router *mux.Router) {

	router.Path("/email-checker").
		HandlerFunc(EmailCheckerHandler).
		Methods("GET")
}

func EmailCheckerHandler(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		server.WriteErrorResponse(res, 500, "Im sorry - something went wrong. Please contact our support.")
		return
	}

	var data RequestBody
	err = ffjson.Unmarshal(body, &data)
	if err != nil {
		server.WriteErrorResponse(res, 400, "Error reading request body. Please check the docs")
		return
	}

	// map for storing unique destination emails
	targets := make(map[string]bool)
	validateTargets(data.Emails, targets)

	response := &ResponseBody{
		Targets: len(targets),
	}

	rbody, err := ffjson.Marshal(response)
	if err != nil {
		server.WriteErrorResponse(res, 500, "Im sorry - something went wrong. Please contact our support.")
		return
	}

	server.WriteResponse(res, rbody)
}

func isEmailValid(e string) bool {
	if len(e) < 3 && len(e) > 254 {
		return false
	}
	return emailRegex.MatchString(e)
}

func validateTargets(emails []string, targets map[string]bool) {
	for _, email := range emails {
		// continue if not valid email address
		if !isEmailValid(email) {
			log.Println(email + " is not a valid email address")
			continue
		}

		// get username
		parts := strings.Split(email, "@")
		username := parts[0]
		domain := parts[1]

		//	get partial before '+'
		username = strings.Split(username, "+")[0]

		// remove "."
		username = strings.ReplaceAll(username, ".", "")

		// add to map if unique
		if _, ok := targets[username+"@"+domain]; !ok {
			targets[username+"@"+domain] = true
		}
	}
}
