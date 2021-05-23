package main

import (
	"encoding/json"
	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/converter"
	"github.com/go-email-validator/go-ev-presenters/pkg/presentation/presentation_test"
	"log"
	"os"
)

// TODO DRY for gen fixtures
func main() {
	var err error
	emails := converter.EmailsForTests()
	deps := make([]interface{}, len(emails))

	verifier := emailverifier.NewVerifier().
		EnableGravatarCheck().
		EnableSMTPCheck()
	// .EnableDomainSuggest()

	es := emails

	skipEmail := hashset.New(
		/* TODO add proxy to test
		5.7.1 Service unavailable, Client host [94.181.152.110] blocked using Spamhaus. To request removal from this list see https://www.spamhaus.org/query/ip/94.181.152.110 (AS3130). [BN8NAM12FT053.eop-nam12.prod.protection.outlook.com]
		*/
		"salestrade86@hotmail.com",
		"monicaramirezrestrepo@hotmail.com",
	)

	for i, email := range es {
		if skipEmail.Contains(email) {
			log.Printf("skipped %v", email)
			continue
		}
		verifyResult, _ := verifier.Verify(email)

		deps[i] = verifyResult
	}

	f, err := os.Create(presentation_test.DefaultDepFixtureFile)
	die(err)
	defer f.Close()

	bytes, err := json.MarshalIndent(deps, "", "  ")
	die(err)
	_, err = f.Write(bytes)
	die(err)
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
