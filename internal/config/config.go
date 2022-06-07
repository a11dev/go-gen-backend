package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	defaultServerPort  = "8083"
	defaultLdapEnabled = true
)

// Ldap standard configuration
type LdapCfg struct {
	LdapServer  string `yaml:"ldap_server" env:"LDAP_SERVER"`
	LdapDomain  string `yaml:"ldap_domain" env:"LDAP_DOMAIN"`
	LdapBase    string `yaml:"ldap_base" env:"LDAP_BASE"`
	LdapEnabled bool   `yaml:"ldap_enabled" env:"LDAP_SERVER"`
}

// Json web token commons configuration parameters
type JwtCfg struct {
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	JWTExpiration string `yaml:"jwt_expire_time_sec" env:"JWT_EXPIRETIMESEC"`
	JWTIssuer     string `yaml:"jwt_issuer" env:"JWT_ISSUER"`
	JWTRealm      string `yaml:"jwt_realm" env:"JWT_REALM"`
	JWTId         string `yaml:"jwt_id" env:"JWT_ID"`
}

// generic configuration structure containing all configuration part
type Config struct {
	ServerPort string  `yaml:"server_port" env:"SERVER_PORT"`
	LdapConfig LdapCfg `yaml:"ldap_config"`
	JwtConfig  JwtCfg  `yaml:"jwt_info"`
}

// Receive a yaml configuration file name as input, unmarshall it into the Config structure
// ToDo: configuration validation
func Load() (*Config, error) {
	file, err := GetEnvType()

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	l := LdapCfg{
		LdapEnabled: true,
	}
	c := Config{
		ServerPort: defaultServerPort,
		LdapConfig: l,
	}

	bytesv, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(bytesv, &c); err != nil {
		return nil, err
	}

	return &c, err
}
