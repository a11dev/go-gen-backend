package config

import (
	"os"
	"testing"
)

func TestGetEnvType(t *testing.T) {
	// test setup
	os.Setenv("ENVIRONMENT_TYPE", "AUTOTEST")

	// tests execution
	wants := "../../config/dev.yml"
	env, err := GetEnvType()

	if env != wants {
		t.Fatalf(`TestGetEnvType() = %q, %v, want match for %#q, nil`, env, err, wants)
	}
	os.Setenv("ENVIRONMENT_TYPE", "DEV")

	env, err = GetEnvType()
	wants = "./config/dev.yml does not exist ; env type: DEV"
	if err.Error() != wants {
		t.Fatalf(`%v, want match for %#q, nil`, err, wants)
	}

	os.Setenv("ENVIRONMENT_TYPE", "PROD")

	env, err = GetEnvType()
	wants = "./config/prod.yml does not exist ; env type: PROD"
	if err.Error() != wants {
		t.Fatalf(`%v, want match for %#q, nil`, err, wants)
	}

	os.Setenv("ENVIRONMENT_TYPE", "AUTOTEST")

	env, err = GetEnvType()
	wants = "../../config/dev.yml"
	if env != wants || err != nil {
		t.Fatalf(`TestGetEnvType() = %q, %v, want match for %#q, nil`, env, err, wants)
	}
}
