package config

type DialogConfig struct {
	Icon string
	Box  string
}

var DialogCfg DialogConfig

func GetDialogConfig() *DialogConfig {
	return &DialogCfg
}
