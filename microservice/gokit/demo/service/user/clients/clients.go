package clients

import "google.golang.org/grpc"

var clientsMap map[string]*grpc.ClientConn
