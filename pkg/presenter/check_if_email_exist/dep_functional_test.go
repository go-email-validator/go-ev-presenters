package check_if_email_exist

import (
	"github.com/go-email-validator/go-email-validator/pkg/ev"
	"github.com/go-email-validator/go-email-validator/pkg/ev/ev_email"
	"github.com/go-email-validator/go-ev-presenters/pkg/presenter/preparer"
	"reflect"
	"testing"
)

func TestDepPreparer_Prepare(t *testing.T) {
	type fields struct {
		preparer              preparer.MultiplePreparer
		calculateAvailability FuncAvailability
	}
	type args struct {
		email  ev_email.EmailAddress
		result ev.ValidationResult
		opts   preparer.Options
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := DepPreparer{
				preparer:              tt.fields.preparer,
				calculateAvailability: tt.fields.calculateAvailability,
			}
			if got := s.Prepare(tt.args.email, tt.args.result, tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}
