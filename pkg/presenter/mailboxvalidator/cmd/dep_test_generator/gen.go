// +build ignore
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common/dep_fixture_generator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/mailboxvalidator"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	presenterName = "DepPresenterForView"
	packageName   = "mailboxvalidator"
)

func main() {
	var bodyBytes []byte
	var err error
	var dep mailboxvalidator.DepPresenterForView
	emails := common.EmailsForTests()[0:1]
	deps := make([]interface{}, len(emails))

	err = godotenv.Load()
	die(err)

	apiKey := os.Getenv("MAIL_BOX_VALIDATOR_API")
	if apiKey == "" {
		panic("MAIL_BOX_VALIDATOR_API should be set")
	}

	for i, email := range emails {
		req, err := http.NewRequest(
			"GET",
			"https://api.mailboxvalidator.com/v1/validation/single?email="+url.QueryEscape(email)+"&key="+url.QueryEscape(apiKey),
			nil,
		)
		die(err)

		func() {
			resp, err := http.DefaultClient.Do(req)
			die(err)
			defer resp.Body.Close()
			bodyBytes, err = ioutil.ReadAll(resp.Body)
			die(err)
		}()

		err = json.Unmarshal(bodyBytes, &dep)
		die(err)

		if dep.ErrorCode != "" {
			panic(fmt.Sprint(email, dep.ErrorMessage))
		}

		deps[i] = dep
	}

	// TODO need convert from DepPresenterForView to DepPresenter
	f, err := os.Create("dep_fixture_test2.go")
	die(err)
	defer f.Close()

	dep_fixture_generator.PackageTemplate.Execute(f, dep_fixture_generator.Template{
		Timestamp:     time.Now(),
		PackageName:   packageName,
		PresenterName: presenterName,
		Presenters:    deps,
	})
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
