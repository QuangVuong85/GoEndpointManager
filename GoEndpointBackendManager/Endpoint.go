package GoEndpointBackendManager

import (
	"net"
	"time"
)

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

func ParseProtocol(name string) TType {
	switch name {
	case "binary":
		return EThriftBinary
	case "compact":
		return EThriftCompact
	}
	return Eunknown
}

type EndPoint struct {
	Host      string
	Port      string
	Type      TType
	ServiceID string
}

func (e *EndPoint) IsGoodEndpoint() bool {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(e.Host, e.Port), 5*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func NewEndPoint(aHost string, aPort string, aType TType) *EndPoint {
	return &EndPoint{
		Host: aHost,
		Port: aPort,
		Type: aType,
	}
}
