package mq

type Config struct {
	Nameserver []string `mapstructure:"name_server"`
	RetryTimes int      `mapstructure:"retry_times"`
	AK         string   `mapstructure:"ak"`
	SK         string   `mapstructure:"sk"`
	GroupName  string   `mapstructure:"group_name"`
}

type ProducerConfig struct {
	Config
}

type ConsumerConfig struct {
	Config
}
