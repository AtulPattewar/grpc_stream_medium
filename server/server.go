package main

import (
	"grpc_stream_medium/server/sensorpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) TempSensor(req *sensorpb.SensorRequest,
	stream sensorpb.Sensor_TempSensorServer) error {

	return nil
}

func (*server) HumiditySensor(req *sensorpb.SensorRequest,
	stream sensorpb.Sensor_HumiditySensorServer) error {

	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer()
	sensorpb.RegisterSensorServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}
}
