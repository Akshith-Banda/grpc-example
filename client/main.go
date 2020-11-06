package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"example/proto"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var client proto.TrialClient

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("ca-cert.crt")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	clientCert, err := tls.LoadX509KeyPair("client-cert.crt", "client-unencrypted-key.key")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates:       []tls.Certificate{clientCert},
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	flag.Parse()

	tlsCredentials, err := loadTLSCredentials()
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	conn, err := grpc.Dial("52.66.126.83:4040", grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client = proto.NewTrialClient(conn)
	//g := gin.Default()

	createThermostat("orient", 75, true)
	createGroup("heater")
	addThermostatToGroup("orient", "heater")
	//deleteGroup("boiler")
	getThermostat("orient")

}

func createThermostat(name string, ctemp int64, tempdsply bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	req := &proto.Thermostat{Name: name, CurrentTemp: ctemp, TempDsply: tempdsply}
	response, err := client.Add_Thermo(ctx, req)
	if err != nil {
		log.Fatalf("%v.Add_Thermo(_) = _, %v: ", client, err)
	}
	log.Println(response.Result)
}

func createGroup(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &proto.Group{Name: string(name)}
	response, err := client.Add_Group(ctx, req)
	if err != nil {
		log.Fatalf("%v.Add_Group(_) = _, %v: ", client, err)
	}
	log.Println(response.Result)
}

func addThermostatToGroup(tname, gname string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &proto.ThermoGrp{ThermoName: string(tname), GroupName: string(gname)}
	response, err := client.Add_Thermo_Grp(ctx, req)
	if err != nil {
		log.Fatalf("%v.Add_Thermo_Grp(_) = _, %v: ", client, err)
	}
	log.Println(response.Result)
}

func deleteGroup(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &proto.Group{Name: string(name)}
	response, err := client.Delete_Group(ctx, req)
	if err != nil {
		log.Fatalf("%v.Delete_Group(_) = _, %v: ", client, err)
	}
	log.Println(response.Result)
}

func getThermostat(name string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	req := &proto.Thermostat{Name: string(name)}
	response, err := client.Get_Thermo(ctx, req)
	if err != nil {
		log.Fatalf("%v.Get_Thermo(_) = _, %v: ", client, err)
	}
	log.Println(response.ThermostatGet)
}
