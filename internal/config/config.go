package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

const (
	defaultServerPort  = "8083"
	defaultLdapEnabled = true
)

type LdapCfg struct {
	LdapServer  string `yaml:"ldap_server" env:"LDAP_SERVER"`
	LdapDomain  string `yaml:"ldap_domain" env:"LDAP_DOMAIN"`
	LdapBase    string `yaml:"ldap_base" env:"LDAP_BASE"`
	LdapEnabled bool   `yaml:"ldap_enabled" env:"LDAP_SERVER"`
}

type JwtCfg struct {
	JWTSigningKey string `yaml:"jwt_signing_key" env:"JWT_SIGNING_KEY,secret"`
	JWTExpiration string `yaml:"jwt_expire_time_sec" env:"JWT_EXPIRETIMESEC"`
	JWTIssuer     string `yaml:"jwt_issuer" env:"JWT_ISSUER"`
	JWTRealm      string `yaml:"jwt_realm" env:"JWT_REALM"`
	JWTId         string `yaml:"jwt_id" env:"JWT_ID"`
}

type Config struct {
	ServerPort string  `yaml:"server_port" env:"SERVER_PORT"`
	LdapConfig LdapCfg `yaml:"ldap_config"`
	JwtConfig  JwtCfg  `yaml:"jwt_info"`
}

func Load(file string) (*Config, error) {
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
