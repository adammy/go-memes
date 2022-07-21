package httpapi

const (
	// ChiServerType denotes the Server was implemented with Chi.
	ChiServerType ServerType = "Chi"

	// GinServerType denotes the Server was implemented with Gin.
	GinServerType ServerType = "Gin"

	// StdLibServerType denotes the Server was implemented with the standard library.
	StdLibServerType ServerType = "StdLib"

	// ContentTypeHeader is the "Content-Type" header key.
	ContentTypeHeader = "Content-Type"

	// ApplicationJson is the "application/json" header value for ContentTypeHeader.
	ApplicationJson = "application/json"

	// TraceLogLevel denotes to show trace and all log levels below.
	TraceLogLevel LogLevel = "trace"

	// DebugLogLevel denotes to show debug and all log levels below.
	DebugLogLevel LogLevel = "debug"

	// InfoLogLevel denotes to show info and all log levels below.
	InfoLogLevel LogLevel = "info"

	// WarnLogLevel denotes to show warn and all log levels below.
	WarnLogLevel LogLevel = "warn"

	// ErrorLogLevel denotes to show error and all log levels below.
	ErrorLogLevel LogLevel = "error"

	// CriticalLogLevel denotes to show critical logs only.
	CriticalLogLevel LogLevel = "critical"
)
