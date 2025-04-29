package city

import (
	"testing"
)

func TestCity_GetAdcode(t *testing.T) {
	err := CityClient.LoadCodeMap()
	if err != nil {
		t.Errorf("LoadCodeMap() error = %v", err)
		return
	}

	type args struct {
		cityName string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "Test with existing city name",
			args: args{
				cityName: "吴江",
			},
			want:  "320509",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, got1 := CityClient.GetAdcode(tt.args.cityName)
			if got != tt.want {
				t.Errorf("GetAdcode() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetAdcode() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
