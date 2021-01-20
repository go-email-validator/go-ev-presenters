package check_if_email_exist

import (
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/evmail"
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

	tests := make([]DepPresenter, 0)
	presenter_test.TestDepPresenters(t, &tests, "")

	// Some data or functional cannot be matched, see more nearby DepPresenter of emails
	skipEmail := hashset.New(
		// TODO problem with SMTP, CIEE think that email is not is_catch_all. Need to run and research source code on RUST
		"sewag33689@itymail.com",
		/* TODO add proxy to test
		5.7.1 Service unavailable, Client host [94.181.152.110] blocked using Spamhaus. To request removal from this list see https://www.spamhaus.org/query/ip/94.181.152.110 (AS3130). [BN8NAM12FT053.eop-nam12.prod.protection.outlook.com]
		*/
		"salestrade86@hotmail.com",
		"monicaramirezrestrepo@hotmail.com",
		// TODO CIEE banned
		"credit@mail.ru",
		// TODO need check source code
		"asdasd@tradepro.net",
		// TODO need check source code
		"y-numata@senko.ed.jp",
	)

	opts := preparer.NewOptions(0)
	for _, tt := range tests {
		tt := tt
		if skipEmail.Contains(tt.Input) {
			t.Logf("skipped %v", tt.Input)
			continue
		}
		t.Run(tt.Input, func(t *testing.T) {
			t.Parallel()
			email := evmail.FromString(tt.Input)

			resultValidator := validator.Validate(ev.NewInput(email))
			got := d.Prepare(email, resultValidator, opts)
			gotPresenter := got.(DepPresenter)

			sort.Strings(gotPresenter.MX.Records)
			sort.Strings(tt.MX.Records)
			if !reflect.DeepEqual(got, tt) {
				t.Errorf("Prepare()\n%#v, \n want\n%#v", got, tt)
			}
		})
	}
}
