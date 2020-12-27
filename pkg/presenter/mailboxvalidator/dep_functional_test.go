package mailboxvalidator

import (
	"flag"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"reflect"
	"testing"
)

var functional = flag.Bool("functional", false, "run functional tests")

func FunctionalSkip(t *testing.T) {
	flag.Parse()
	if !*functional {
		t.Skip()
	}
}

func TestDepPreparer_Functional_Prepare(t *testing.T) {
	FunctionalSkip(t)

	validator := NewDepValidator()
	d := NewDepPreparerDefault()

	// Some data or functional cannot be matched, see more nearby DepPresenter of emails
	skipEmail := hashset.New(
		"zxczxczxc@joycasinoru", // TODO syntax is valid
		"sewag33689@itymail.com",
		"derduzikne@nedoz.com",
		"tvzamhkdc@emlhub.com",
		"admin@gmail.com",
		"salestrade86@hotmail.com",
		"monicaramirezrestrepo@hotmail.com",
		"y-numata@senko.ed.jp",
		"pr@yandex-team.ru",
	)

	tests := getPresenters()
	for _, tt := range tests {
		if skipEmail.Contains(tt.EmailAddress) {
			t.Logf("skipped %v", tt.EmailAddress)
			continue
		}

		t.Run(tt.EmailAddress, func(t *testing.T) {
			email := EmailFromString(tt.EmailAddress)
			opts := preparer.NewOptions(tt.TimeTaken)

			resultValidator := validator.Validate(email)
			if gotResult := d.Prepare(email, resultValidator, opts); !reflect.DeepEqual(gotResult, tt) {
				t.Errorf("Prepare()\n%v, \n want\n%v", gotResult, tt)
			}
		})
	}
}
