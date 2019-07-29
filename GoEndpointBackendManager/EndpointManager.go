package GoEndpointBackendManager

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	etcdclient "go.etcd.io/etcd/client"
)

type FncProcessEventChange func(ep *EndPoint)

type EndPointManager struct {
	mux          sync.Mutex
	etcdServer   string
	etcdBasePath string
	endPoints    []*EndPoint
	etcdApi      etcdclient.KeysAPI
}

// GetEndPoint get random endpoint from endpoints
func (e *EndPointManager) GetEndPoint() (error, *EndPoint) {
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	index := r.Intn(len(e.endPoints))
	return nil, e.endPoints[index]
}

// GetEndPoints get all endpoint from etcd and base path
func (e *EndPointManager) GetEndPoints() (error, []*EndPoint) {
	if e.endPoints == nil {
		return errors.New("Enpoind nil"), nil
	}
	return nil, e.endPoints
}

func (e *EndPointManager) GetEndPointType(t TType) (error, *EndPoint) {
	for i := 0; i < len(e.endPoints); i++ {
		if e.endPoints[i].Type == t {
			return nil, e.endPoints[i]
		}
	}
	return errors.New("Can not found endpoint service"), nil
}

// LoadEndpoint load all endpoint from etcd and base path
func (e *EndPointManager) LoadEndpoint() error {

	log.Println("Load endpoint from", e.etcdServer, "with base path", e.etcdBasePath)

	return e.doLoadEndpoint()
}

func (e *EndPointManager) LoadEndPointFromServer(etcdServer, basePath string) error {
	e.etcdServer = etcdServer
	e.etcdBasePath = basePath
	return e.doLoadEndpoint()
}

func (e *EndPointManager) doLoadEndpoint() error {
	cfg := etcdclient.Config{
		Endpoints:               []string{e.etcdServer},
		Transport:               etcdclient.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := etcdclient.New(cfg)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("Can not connect to etcd")
	}
	e.etcdApi = etcdclient.NewKeysAPI(c)
	rs, err := e.etcdApi.Get(context.Background(), e.etcdBasePath, nil)
	if err != nil {
		return err
	}
	var listEp []*EndPoint
	for i := 0; i < rs.Node.Nodes.Len(); i++ {
		epPath := rs.Node.Nodes[i].Key
		epValue := rs.Node.Nodes[i].Value

		fmt.Println("enpoint path key :", epPath, "enpoint value :", epValue)
		err, ep := e.parseEndpoint(epPath)
		if err != nil {
			log.Println(err.Error())
		} else {
			listEp = append(listEp, ep)
		}
	}

	if len(listEp) > 0 {
		e.replaceAll(listEp)
	}

	return nil
}

func (e *EndPointManager) parseEndpoint(endPointPath string) (error, *EndPoint) {
	var ep EndPoint
	baseNode := strings.Split(endPointPath, "/")
	if len(baseNode) == 0 {
		return errors.New("Parse endpoint error " + endPointPath), nil
	}
	nodeName := baseNode[len(baseNode)-1]
	token := strings.Split(nodeName, ":")
	if len(token) != 3 {
		return errors.New("Parse endpoint error " + nodeName), nil
	}
	port, err := strconv.Atoi(token[2])
	if err != nil {
		return errors.New("Parse endpoint error " + nodeName), nil
	}
	ep.Type = StringToTType(token[0])
	ep.Host = token[1]
	ep.Port = port
	ep.EnpointFullPath = endPointPath
	return nil, &ep
}

func (e *EndPointManager) removeEndPoint(ep *EndPoint) {
	for i, v := range e.endPoints {
		if v.Host == ep.Host && v.Port == ep.Port && v.Type == ep.Type {
			e.endPoints = append(e.endPoints[:i], e.endPoints[i+1:]...)
			return
		}
	}
}

func (e *EndPointManager) replaceAll(listEndPoints []*EndPoint) {
	e.mux.Lock()
	defer e.mux.Unlock()
	e.endPoints = listEndPoints
}

func (e *EndPointManager) EventChangeEndPoints(fn FncProcessEventChange) {
	watch := e.etcdApi.Watcher(e.etcdBasePath, &etcdclient.WatcherOptions{AfterIndex: 0, Recursive: true})
	go func() {
		for {
			res, err := watch.Next(context.Background())
			if err != nil {
				return
			}
			err, ep := e.parseEndpoint(res.Node.Key)
			if err != nil {
				continue
			}
			fn(ep)
		}
	}()

}

func (e *EndPointManager) TestConnectEtcdServer() error {
	cfg := etcdclient.Config{
		Endpoints:               []string{e.etcdServer},
		Transport:               etcdclient.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := etcdclient.New(cfg)
	if err != nil {
		return err
	}
	if c == nil {
		return errors.New("Can not connect to etcd")
	}
	e.etcdApi = etcdclient.NewKeysAPI(c)
	return nil
}

func NewEndPointManager(aServer, aPath string) *EndPointManager {
	return &EndPointManager{
		etcdServer:   aServer,
		etcdBasePath: aPath,
	}
}
