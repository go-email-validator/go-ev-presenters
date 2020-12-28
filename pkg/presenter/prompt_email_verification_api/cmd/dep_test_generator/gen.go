// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common/dep_fixture_generator"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/prompt_email_verification_api"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	presenterName = "DepPresenter"
	packageName   = "prompt_email_verification_api"
)

func main() {
	var bodyBytes []byte
	var err error
	var dep prompt_email_verification_api.DepPresenter
	emails := common.EmailsForTests()
	deps := make([]interface{}, len(emails))

	apiKey := os.Getenv("PROMPT_EMAIL_VERIFICATION_API")
	if apiKey == "" {
		panic("PROMPT_EMAIL_VERIFICATION_API should be set")
	}

	for i, email := range emails {
		req, err := http.NewRequest(
			"GET",
			fmt.Sprintf("https://api.promptapi.com/email_verification/%s", email),
			nil,
		)
		die(err)
		req.Header.Set("apikey", apiKey)

		func() {
			resp, err := http.DefaultClient.Do(req)
			die(err)
			defer resp.Body.Close()
			bodyBytes, err = ioutil.ReadAll(resp.Body)
			die(err)
		}()

		err = json.Unmarshal(bodyBytes, &dep)
		die(err)

		if !strings.Contains(dep.Message, "API rate limit") {
			panic(fmt.Sprint(email, dep.Message))
		}

		deps[i] = dep
	}

	f, err := os.Create("dep_fixture_test.go")
	die(err)
	defer f.Close()

	data := dep_fixture_generator.Template{
		Timestamp:     time.Now(),
		PackageName:   packageName,
		PresenterName: presenterName,
		Presenters:    deps,
	}
	die(dep_fixture_generator.PackageTemplate.Execute(f, data))
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
