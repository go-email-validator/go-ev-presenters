// +build ignore
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	ciee "github.com/go-email-validator/go-ev-presenters/pkg/presenter/check_if_email_exist"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common/dep_fixture_generator"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	presenterName = "DepPresenter"
	packageName   = "check_if_email_exist"
)

func main() {
	var bodyBytes []byte
	var err error
	var dep ciee.DepPresenter
	emails := common.EmailsForTests()
	deps := make([]interface{}, len(emails))

	apiKey := os.Getenv("CHECK_IF_EMAIL_EXIST")

	for i, email := range emails {
		message := map[string]interface{}{
			"to_email": email,
		}

		bytesRepresentation, _ := json.Marshal(message)
		req, err := http.NewRequest(
			"POST",
			"https://ssfy.sh/amaurymartiny/reacher@2d2ce35c/check_email",
			bytes.NewBuffer(bytesRepresentation),
		)
		die(err)
		req.Header.Set("Content-Type", "application/json")
		if apiKey != "" {
			req.Header.Set("authorization", apiKey)
		}

		func() {
			resp, err := http.DefaultClient.Do(req)
			die(err)
			defer resp.Body.Close()
			bodyBytes, err = ioutil.ReadAll(resp.Body)
			die(err)
		}()

		err = json.Unmarshal(bodyBytes, &dep)
		die(err)

		if dep.Error != "" {
			panic(fmt.Sprint(email, dep.Error))
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
		Import:        "import \"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common\"",
	}
	die(dep_fixture_generator.PackageTemplate.Execute(f, data))
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
