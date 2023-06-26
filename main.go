package main

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
	clientv1 "github.com/influxdata/influxdb1-client/v2"
)

func main() {
	influxdb_v1_write()

	influxdb_v2_write()
}

// https://github.com/influxdata/influxdb1-client
func influxdb_v1_write() {
	c, err := clientv1.NewHTTPClient(clientv1.HTTPConfig{
		Addr: "http://localhost:4000/v1/influxdb",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	bp, _ := clientv1.NewBatchPoints(clientv1.BatchPointsConfig{
		Database:  "public",
		Precision: "s",
	})

	tags := map[string]string{"cpu": "cpu-total"}
	fields := map[string]interface{}{
		"idle":   10.1,
		"system": 53.3,
		"user":   46.6,
	}
	pt, err := clientv1.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	bp.AddPoint(pt)

	// Write the batch
	c.Write(bp)

}

// https://github.com/influxdata/influxdb-client-go
func influxdb_v2_write() {
	client := influxdb2.NewClient("http://localhost:4000/v1/influxdb", "")

	writeAPI := client.WriteAPIBlocking("", "public")

	line := fmt.Sprintf("stat,unit=123 avg=%f,max=%f %d", 23.5, 45.0, time.Now().UnixNano())

	err := writeAPI.WriteRecord(context.Background(), line)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	client.Close()
}
