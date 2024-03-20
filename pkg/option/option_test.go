package option

import (
	"errors"
	"testing"
)

var errTest = errors.New("test error")

type testStruct struct {
	Value int
	Text  string
}

func withValue(val int) Option[testStruct] {
	return func(ts *testStruct) error {
		ts.Value = val
		return nil
	}
}

func withText(text string) Option[testStruct] {
	return func(ts *testStruct) error {
		ts.Text = text
		return nil
	}
}

func failingOption() Option[testStruct] {
	return func(_ *testStruct) error {
		return errTest
	}
}

func TestApply(t *testing.T) {
	testCases := []struct {
		name      string
		opts      []Option[testStruct]
		wantValue int
		wantText  string
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "SingleOptionAppliesCorrectly",
			opts:      []Option[testStruct]{withValue(1)},
			wantValue: 1,
		},
		{
			name:      "MultipleOptionsApplyCorrectly",
			opts:      []Option[testStruct]{withValue(69), withText("test")},
			wantValue: 69,
			wantText:  "test",
		},
		{
			name:    "OptionReturnsError",
			opts:    []Option[testStruct]{failingOption()},
			wantErr: true,
			errMsg:  errTest.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := testStruct{}
			err := Apply(&ts, tc.opts...)

			if (err != nil) != tc.wantErr {
				assertError(t, err, tc.errMsg)
			}

			assertTestStructState(t, ts, tc.wantValue, tc.wantText)
		})
	}
}

func TestGroup(t *testing.T) {
	testCases := []struct {
		name      string
		opts      []Option[testStruct]
		wantValue int
		wantText  string
		wantErr   bool
		errMsg    string
	}{
		{
			name:      "GroupAndApplyMultipleOptionsCorrectly",
			opts:      []Option[testStruct]{withValue(69), withText("test")},
			wantValue: 69,
			wantText:  "test",
		},
		{
			name:    "GroupOptionReturnsError",
			opts:    []Option[testStruct]{failingOption()},
			wantErr: true,
			errMsg:  errTest.Error(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := testStruct{}
			groupedOption := Group(tc.opts...)
			err := Apply(&ts, groupedOption)

			if (err != nil) != tc.wantErr {
				assertError(t, err, tc.errMsg)
			}

			assertTestStructState(t, ts, tc.wantValue, tc.wantText)
		})
	}
}

func assertTestStructState(t testing.TB, got testStruct, wantValue int, wantText string) {
	t.Helper()

	if got.Value != wantValue || got.Text != wantText {
		t.Errorf("got '%+v', want '{Value:%d Text:%s}'", got, wantValue, wantText)
	}
}

func assertError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		t.Fatal("expected an error but got none")
	}

	if got.Error() != want {
		t.Errorf("got error '%s', want '%s'", got, want)
	}
}
