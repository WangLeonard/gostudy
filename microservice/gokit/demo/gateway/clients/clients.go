package clients

import (
	"fmt"
	"google.golang.org/grpc"
)

//TODO: key shoule use service name or instanceAddr ?
//fixme: use concurrent map
//fixme: need runtime.SetFinalizer to clean?
var clientsMap = make(map[string]*grpc.ClientConn, 10)

func GetClient(instanceAddr string) (*grpc.ClientConn, error) {
	fmt.Println("clientsMap.Len:", len(clientsMap))
	if c, ok := clientsMap[instanceAddr]; ok {
		return c, nil
	}

	if cl, err := grpc.Dial(instanceAddr, grpc.WithInsecure()); err != nil {
		return nil, err
	} else {
		clientsMap[instanceAddr] = cl
		return cl, nil
	}
}
