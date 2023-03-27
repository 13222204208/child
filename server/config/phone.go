package config

type Phone struct {
	PhoneKeyId        string `mapstructure:"phone-key-id" json:"phone-key-id" yaml:"phone-key-id"`
	PhoneKeySecret    string `mapstructure:"phone-key-secret" json:"phone-key-secret" yaml:"phone-key-secret"`
	PhoneSignName     string `mapstructure:"phone-sign-name" json:"phone-sign-name" yaml:"phone-sign-name"`
	PhoneTemplateCode string `mapstructure:"phone-template-code" json:"phone-template-code" yaml:"phone-template-code"`
}
