package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	functionPrefix := "forward_EmailValidation_SingleValidation_"
	functionArguments := "(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)"

	placeToOpenApiPreparation := "(" +
		regexp.QuoteMeta(functionPrefix) +
		"\\d+" +
		regexp.QuoteMeta(functionArguments) +
		")"

	openApiPreparation :=
		`respResult := resp.(*EmailResponse).GetResult()
		switch result := respResult.(type) {
		case *EmailResponse_CheckIfEmailExist:
			resp = proto.Message(result.CheckIfEmailExist)
		case *EmailResponse_MailBoxValidator:
			resp = proto.Message(result.MailBoxValidator)
		}
		$1`

	re := regexp.MustCompile(placeToOpenApiPreparation)

	filePath, err := filepath.Abs("pkg/api/v1/ev.pb.gw.go")
	check(err)
	fileInfo, _ := os.Stat(filePath)
	data, err := ioutil.ReadFile(filePath)
	check(err)

	str := re.ReplaceAllString(string(data), openApiPreparation)

	err = ioutil.WriteFile(filePath, []byte(str), fileInfo.Mode())
	check(err)
}
