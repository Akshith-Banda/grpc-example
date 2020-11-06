package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"example/models"
	"example/proto"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

/*
var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
)
*/
type server struct{}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	pemClientCA, err := ioutil.ReadFile("certs/ca-cert.crt")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemClientCA) {
		return nil, fmt.Errorf("failed to add client CA's certificate")
	}

	serverCert, err := tls.LoadX509KeyPair("certs/server-cert.crt", "certs/server-unencrypted-key.key")
	if err != nil {
		return nil, err
	}
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	models.Database()
	flag.Parse()
	listner, err := net.Listen("tcp", ":4040")
	if err != nil {
		fmt.Sprint("we got error while listening!")
	}
	/*
		var opts []grpc.ServerOption
		if *tls {
			if *certFile == "" {
				*certFile = data.Path("x509/server_cert.pem")
			}
			if *keyFile == "" {
				*keyFile = data.Path("x509/server_key.pem")
			}
			creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
			if err != nil {
				log.Fatalf("Failed to generate credentials %v", err)
			}
			opts = []grpc.ServerOption{grpc.Creds(creds)}
		}
	*/
	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	srv := grpc.NewServer(grpc.Creds(tlsCredentials))
	proto.RegisterTrialServer(srv, &server{})
	if e := srv.Serve(listner); e != nil {
		fmt.Sprint("hey we got an error in server!")
	}
}

func (s *server) Add_Thermo(ctx context.Context, req *proto.Thermostat) (*proto.Response, error) {
	name := req.GetName()
	ctemp := req.GetCurrentTemp()
	tempdsply := req.GetTempDsply()
	result := models.Thermo(name, ctemp, tempdsply)

	return &proto.Response{Result: result}, nil
}

func (s *server) Add_Group(ctx context.Context, req *proto.Group) (*proto.Response, error) {
	name := req.GetName()
	result := models.Grp(name)

	return &proto.Response{Result: result}, nil
}

func (s *server) Add_Thermo_Grp(ctx context.Context, req *proto.ThermoGrp) (*proto.Response, error) {
	tname := req.GetThermoName()
	gname := req.GetGroupName()
	result := models.AddTG(tname, gname)

	return &proto.Response{Result: result}, nil
}

func (s *server) Delete_Group(ctx context.Context, req *proto.Group) (*proto.Response, error) {
	name := req.GetName()
	result := models.DeleteGroup(name)

	return &proto.Response{Result: result}, nil
}

func (s *server) Get_Thermo(ctx context.Context, thermo *proto.Thermostat) (*proto.GetThermostat, error) {
	name := thermo.GetName()
	thermostat := models.GetThermo(name)
	return &proto.GetThermostat{ThermostatGet: thermostat}, nil
}
