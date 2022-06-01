package config

import (
	"os"
	"testing"
)

func TestGetEnvType(t *testing.T) {
	wants := "../../config/dev.yml"
	env, err := GetEnvType()

	if env != wants {
		t.Fatalf(`TestGetEnvType() = %q, %v, want match for %#q, nil`, env, err, wants)
	}
	os.Setenv("ENVIRONMENT_TYPE", "DEV")

	env, err = GetEnvType()
	wants = "./config/dev.yml"
	if env != wants || err != nil {
		t.Fatalf(`TestGetEnvType() = %q, %v, want match for %#q, nil`, env, err, wants)
	}

	os.Setenv("ENVIRONMENT_TYPE", "PROD")

	env, err = GetEnvType()
	wants = "./config/prod.yml"
	if env != wants || err != nil {
		t.Fatalf(`TestGetEnvType() = %q, %v, want match for %#q, nil`, env, err, wants)
	}
}
