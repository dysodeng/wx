package open_platform

type config struct {
	appId string
}

type Config func(*config)

type option struct {
}

type Option func(*option)