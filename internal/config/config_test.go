package config

import (
	"flag"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestConfigLoading(t *testing.T) {
	confFileName := "../../config/local.yml"

	var flagConfig = flag.String("config", confFileName, "path to the config file")

	flag.Parse()
	want := ""
	cfg, err := Load(*flagConfig)

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
