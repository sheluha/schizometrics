package config

type Llm struct {
	ApiKey string `env:"LLM_TOKEN,required"`
	Url    string `env:"LLM_URL,required"`
}
