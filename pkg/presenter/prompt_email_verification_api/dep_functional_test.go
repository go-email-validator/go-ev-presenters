package prompt_email_verification_api

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evtests"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/presenter_test"
	"reflect"
	"sort"
	"testing"
)

func TestMain(m *testing.M) {
	evtests.TestMain(m)
}

func TestDepPreparer_Prepare(t *testing.T) {
	evtests.FunctionalSkip(t)

	validator := NewDepValidator(nil)
	d := NewDepPreparerDefault()
	tests := make([]DepPresenterTest, 0)
	presenter_test.TestDepPresenters(t, &tests, "")

	// Some data or functional cannot be matched, see more nearby DepPresenter of emails
	skipEmail := hashset.New(
		// Banned on disposable domain
		"sewag33689@itymail.com",
		// Banned on disposable domain
		"sewag33689@itymail.com",
		// Banned on disposable domain
		"asdasd@tradepro.net",
		// Banned on disposable domain
		"tvzamhkdc@emlhub.com",
		// Banned
		"credit@mail.ru",
		// Banned
		"salestrade86@hotmail.com",
		// Banned
		"monicaramirezrestrepo@hotmail.com",
		// TODO Cannot connect from my hosts pc
		"y-numata@senko.ed.jp",
	)

	opts := preparer.NewOptions(0)
	for _, tt := range tests {
		tt := tt
		if skipEmail.Contains(tt.Email) {
			t.Logf("skipped %v", tt.Email)
			continue
		}

		t.Run(tt.Email, func(t *testing.T) {
			t.Parallel()

			email := EmailFromString(tt.Email)

			resultValidator := validator.Validate(ev.NewInput(email))
			got := d.Prepare(email, resultValidator, opts)
			gotPresenter := got.(DepPresenter)

			sort.Strings(gotPresenter.MxRecords.Records)
			sort.Strings(tt.Dep.MxRecords.Records)
			if !reflect.DeepEqual(got, tt.Dep) {
				t.Errorf("Prepare()\n%#v, \n want\n%#v", got, tt.Dep)
			}
		})
	}
}

// see /pkg/presenter/prompt_email_verification_api/cmd/dep_test_generator/gen.go
type DepPresenterTest struct {
	Email string
	Dep   DepPresenter
}
