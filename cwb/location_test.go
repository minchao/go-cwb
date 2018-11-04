package cwb

import (
	"reflect"
	"testing"
)

func TestFindLocationByName(t *testing.T) {
	tests := []struct {
		name string
		want Location
		err  error
	}{
		{
			name: "宜蘭縣",
			want: locations[0],
			err:  nil,
		},
		{
			name: "",
			want: Location{},
			err:  errLocationNotFound,
		},
		{
			name: "Not found",
			want: Location{},
			err:  errLocationNotFound,
		},
	}

	for i, test := range tests {
		got, err := FindLocationByName(test.name)
		if err != test.err {
			t.Errorf("(%v) Expected error: %v, got: %v", i, test.err, err)
		}
		if !reflect.DeepEqual(test.want, got) {
			t.Errorf("(%v) Expected Location: %v, got: %v", i, test.want, got)
		}
	}
}
