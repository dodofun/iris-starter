package config

type Server struct {
	HttpProtocol              string
	HttpPort                  string
	HttpRequestData           string
	HttpResponseData          string
	RpcPort                   string
	RpcRequestData            string
	RpcResponseData           string
	AccessControlAllowOrigin  string
	AccessControlAllowHeaders string
	AccessControlAllowMethods string
	HttpAuthDuration          uint64
}
