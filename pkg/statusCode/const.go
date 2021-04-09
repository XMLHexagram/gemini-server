package statusCode

type Status int

const (
	Input                              Status = 10
	Success                            Status = 20
	SuccessEndClientCertificateSession Status = 21
	RedirectTemporary                  Status = 30
	RedirectPermanent                  Status = 31
	TemporaryFailure                   Status = 40
	ServerUnavailable                  Status = 41
	CGIError                           Status = 42
	ProxyError                         Status = 43
	SlowDown                           Status = 44
	PermanentFailure                   Status = 50
	NotFound                           Status = 51
	Gone                               Status = 52
	ProxyRequestRefused                Status = 53
	BadRequest                         Status = 59
	ClientCertRequired                 Status = 60
	TransientCertRequested             Status = 61
	AuthorisedCertRequired             Status = 62
	CertNotAccepted                    Status = 63
	FutureCertRejected                 Status = 64
	ExpiredCertRejected                Status = 65
)
