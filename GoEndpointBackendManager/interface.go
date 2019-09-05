package GoEndpointBackendManager

type FncProcessEventChange func(ep *EndPoint)

type EndPointManagerIf interface {
	// dat cau hinh enpoint trong ram trong truong hop khong ket noi toi etcdserver
	SetDefaultEnpoint(serviceID, host string, port string, epType TType)
	// lay cau hinh enpoind trong ram trong truong hop khong ket noi to etcdserver
	GetDefaultEndpoint(ServiceID string, epType TType) (*EndPoint, error)
	// Load danh sach endpoint tu etcdserver dua tren basepath
	LoadEndpoints() error
	// Lay 1 enpoint bat ky trong danh sach endpoint da co
	GetEndPoint() (error, *EndPoint)
	// Lay endpoint dua tren kieu endpoind
	GetEndPointType(t TType) (error, *EndPoint)
	// Xu li su kien khi 1 endpoint bi thay doi gia tri host port
	EventChangeEndPoints(fn FncProcessEventChange)
}
