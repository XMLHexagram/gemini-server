[ServerMap]
  [ServerMap.Gemini]
    Port = ":1965"
    Timeout = 5

[TLS]
  CertFile = "test/MyCertificate.crt" # relative or absolute path
  KeyFile = "test/MyKey.key"

[Log]
  Level = "info"
  ToStd = true # todo: coming soon
  LogRotate = true
  Development = false
  Sampling = true # todo: coming soon
  [Log.Rotate]
    Filename = ["log", "test"]
    MaxSize = 100
    MaxAge = 30
    MaxBackups = 30 # max number of log file

[Gemini]
  DefaultLang = "zh"
  AutoRedirect = true
  AutoRedirectUrl = "gemini://localhost/dir"
  [[Gemini.File]]
    Router = "/file" # the router will be registed
    Path = "test.gmi" # relative or absolute path

  [[Gemini.Dir]]
    Router = "/dir"
    Path = "test/dir"
    Index = "index.gmi"
    AutoCatalogue = false # todo: coming soon

# need json response and will use body field in json struct as response body
  [[Gemini.Proxy]]
    Router = "/url"
    Method = "GET" # todo: coming soon
    URL = "http://127.0.0.1:12222"
    MetaField = "meta" # todo: coming soon
    BodyField = "body"