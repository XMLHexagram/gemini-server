package config

type Config struct {
	ServerMap `toml:"ServerMap"`
	Log       `toml:"Log"`
	Gemini    `toml:"Gemini"`
	TLS       `toml:"TLS"`
}

type TLS struct {
	CertFile string
	KeyFile  string
}

type Server struct {
	Port    string `toml:"Port"`
	Timeout int    `toml:"Timeout"`
}

type ServerMap map[string]Server

type Log struct {
	Level       string
	ToStd       bool
	LogRotate   bool
	Development bool
	Sampling    bool
	Rotate      Rotate
}

type Rotate struct {
	Filename   []string
	MaxSize    int
	MaxAge     int
	MaxBackups int
}

type Gemini struct {
	File []struct {
		Router string `toml:"Router"`
		Path   string `toml:"Path"`
	} `toml:"File"`
	Dir []struct {
		Router        string `toml:"Router"`
		Path          string `toml:"Path"`
		Index         string `toml:"Index"`
		AutoCatalogue bool   `toml:"AutoCatalogue"`
	} `toml:"Dir"`
	Proxy []struct {
		Router    string `toml:"Router"`
		Method    string `toml:"Method"`
		URL       string `toml:"URL"`
		MetaField string `toml:"MetaField"`
		BodyField string `toml:"BodyField"`
	} `toml:"Proxy"`
}
