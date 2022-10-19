package handle

import (
	"github.com/calculi-corp/log"
)

type awscreds struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	IAMRole  string `json:"iam_role"`
	Password string `json:"password"`
}

func VStrings() {
	creds := &awscreds{
		Access:   "somekey",
		Secret:   "secret",
		Password: "myawesomepassword",
	}

	log.Infof("creds: %v", creds)
}
