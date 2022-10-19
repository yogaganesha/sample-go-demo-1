package handle

import "testing"

type awscreds_t struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	IAMRole  string `json:"iam_role"`
	Password string `json:"password"`
}

func TestVStrings(t *testing.T) {
	creds := &awscreds{
		Access:   "secret",
		Secret:   "secret",
		Password: "secret",
	}

	t.Log(creds)

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			VStrings()
		})
	}
}
