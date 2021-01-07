package mailboxvalidator

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evtests"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/common"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	evtests.TestMain(m)
}

func TestDepPreparer_Functional_Prepare(t *testing.T) {
	evtests.FunctionalSkip(t)

	validator := NewDepValidator(nil)
	d := NewDepPreparerDefault()

	tests := detPresenters(t)

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
				t.Errorf("Prepare()\n%#v, \n want\n%#v", gotResult, tt)
			}
		})
	}
}

func detPresenters(t *testing.T) []DepPresenter {
	tests := make([]DepPresenter, 0)
	common.TestDepPresenters(t, &tests, "")

	return tests
}
