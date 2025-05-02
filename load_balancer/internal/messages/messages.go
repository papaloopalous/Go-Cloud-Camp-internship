package messages

// error messages
const (
	ErrLoadConfig         = "failed to load config file"
	ErrLAS                = "listenAndServe failed"
	ErrShutdown           = "server Shutdown Failed"
	ErrAllAttemptsFailed  = "all attempts failed"
	ErrNoBackends         = "no backends are reachable"
	ErrAttemptFailed      = "attempt failed"
	ErrServiceUnavailable = "service unavailable"
	ErrResponse           = "failed to write a response"
	ErrInvalidBackendURL  = "invalid backend url"
	ErrReadConfig         = "unable to read config file: %v"
)

// info messages
const (
	InfoBalancerON         = "load Balancer is on"
	InfoGracefulStopStart  = "shutting down gracefully"
	InfoGracefulStopFinish = "server gracefully stopped"
	InfoForwardingURL      = "forwarding to"
	InfoForwardingActive   = "active"
	InfoSuccessfulProxy    = "successfully proxied to"
	InfoShutdownHealth     = "shutting down health checks"
	InfoUnreachable        = "server is unreachable"
	InfoReachable          = "server is reachable"
)

// misc
const (
	URL    = "URL"
	Port   = "Port"
	Number = "Number"
	Code   = "Code"
	Status = "Status"
)
