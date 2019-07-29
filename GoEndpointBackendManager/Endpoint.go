package GoEndpointBackendManager

type TType int

const (
	Eunknown       TType = -1
	EAnyType       TType = 0
	EHttp          TType = 1
	EThriftBinary  TType = 2
	EThriftCompact TType = 3
	EGrpc          TType = 4
	EGrpcWeb       TType = 5
)

func (t TType) String() string {
	switch t {
	case Eunknown:
		return "Eunknown"
	case EAnyType:
		return "EAnyType"
	case EHttp:
		return "Ehttp"
	case EThriftBinary:
		return "EThriftBinary"
	case EThriftCompact:
		return "EThriftCompact"
	case EGrpc:
		return "EGrpc"
	case EGrpcWeb:
		return "EGrpcWeb"
	}
	return "UnknownType"
}

func StringToTType(t string) TType {
	switch t {
	case "thrift_compact":
		return EThriftCompact
	case "thrift_binary":
		return EThriftBinary
	case "grpc":
		return EGrpc
	case "grpc_web":
		return EGrpcWeb
	default:
		return Eunknown
	}
}

type EndPoint struct {
	Host            string
	Port            int
	Type            TType
	EnpointFullPath string
}

func (e *EndPoint) IsGoodEndpoint() bool {
	return true
}

func NewEndPoint(aHost string, aPort int, aType TType) *EndPoint {
	return &EndPoint{
		Host: aHost,
		Port: aPort,
		Type: aType,
	}
}
