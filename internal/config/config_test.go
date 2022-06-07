package config

import (
	"os"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestConfigLoading(t *testing.T) {
	os.Setenv("ENVIRONMENT_TYPE", "AUTOTEST")

	want := ""
	cfg, err := Load()

	if cfg.LdapConfig.LdapServer == "" || err != nil {
		t.Fatalf(`TestConfigLoading("LdapServer") = %q, %v, want match for %#q, nil`, cfg.LdapConfig.LdapServer, err, want)
	}
	if cfg.LdapConfig.LdapDomain == "" || err != nil {
		t.Fatalf(`TestConfigLoading("LdapDomain") = %q, %v, want match for %#q, nil`, cfg.LdapConfig.LdapDomain, err, want)
	}
	if cfg.LdapConfig.LdapBase == "" || err != nil {
		t.Fatalf(`TestConfigLoading("LdapBase") = %q, %v, want match for %#q, nil`, cfg.LdapConfig.LdapBase, err, want)
	}
}
