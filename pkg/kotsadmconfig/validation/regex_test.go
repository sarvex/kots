package validation

import (
	kotsv1beta1 "github.com/replicatedhq/kots/kotskinds/apis/kots/v1beta1"
	configtypes "github.com/replicatedhq/kots/pkg/kotsadmconfig/types"
	"reflect"
	"testing"
)

func Test_regexValidator_Validate(t *testing.T) {
	type fields struct {
		RegexValidator *kotsv1beta1.RegexValidator
	}
	type args struct {
		input string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *configtypes.ValidationError
	}{
		{
			name: "valid regex",
			fields: fields{
				RegexValidator: &kotsv1beta1.RegexValidator{
					Pattern: ".*",
				},
			},
			args: args{
				input: "test",
			},
			want: nil,
		}, {
			name: "invalid regex",
			fields: fields{
				RegexValidator: &kotsv1beta1.RegexValidator{
					Pattern: "[",
				},
			},
			args: args{
				input: "test",
			},
			want: &configtypes.ValidationError{
				ValidationErrorMessage: "Invalid regex: error parsing regexp: missing closing ]: `[`",
				RegexValidator: &kotsv1beta1.RegexValidator{
					Pattern: "[",
				},
			},
		}, {
			name: "invalid input",
			fields: fields{
				RegexValidator: &kotsv1beta1.RegexValidator{
					Pattern: "test",
				},
			},
			args: args{
				input: "foo",
			},
			want: &configtypes.ValidationError{
				ValidationErrorMessage: regexMatchError,
				RegexValidator: &kotsv1beta1.RegexValidator{
					Pattern: "test",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &regexValidator{
				RegexValidator: tt.fields.RegexValidator,
			}
			if got := v.Validate(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("regexValidator.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
