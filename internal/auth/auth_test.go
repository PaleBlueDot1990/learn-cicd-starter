package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

var ErrFooNoAuthHeaderTest = errors.New("no authorization header included")
var ErrFooMalformedAuthHeaderTest = errors.New("malformed authorization header")

func TestGetApiKey(t *testing.T) {
	type Output struct {
		apiKey string
		err    error
	}

	type Test struct {
		input          http.Header
		expectedOutput Output
	}

	tests := []Test{
		{input: http.Header{"Authorization": []string{"ApiKey qwertyuiop"}}, expectedOutput: Output{apiKey: "qwertyuiop", err: nil}},
		{input: http.Header{"Authentication": []string{"ApiKey qwertyuiop"}}, expectedOutput: Output{apiKey: "", err: ErrFooNoAuthHeaderTest}},
		{input: http.Header{"Authorization": []string{"qwertyuiop"}}, expectedOutput: Output{apiKey: "", err: ErrFooMalformedAuthHeaderTest}},
	}

	for _, test := range tests {
		gotApiKey, gotErr := GetAPIKey(test.input)
		wantApiKey, wantErr := test.expectedOutput.apiKey, test.expectedOutput.err

		if gotErr != nil && wantErr == nil {
			t.Fatalf("Got an error, but expecting no error")
			continue
		}

		if gotErr == nil && wantErr != nil {
			t.Fatalf("Got no error, but expecting an error")
		}

		if gotErr != nil && wantErr != nil && !reflect.DeepEqual(gotErr.Error(), wantErr.Error()) {
			t.Fatalf("expected error: %v, got: %v", wantErr, gotErr)
		}

		if gotErr == nil && wantErr == nil && !reflect.DeepEqual(gotApiKey, wantApiKey) {
			t.Fatalf("expected api key: %s, got: %s", wantApiKey, gotApiKey)
		}
	}
}
