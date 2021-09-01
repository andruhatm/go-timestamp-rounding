package main

import (
	"fmt"
)

type MetricExternalResponse struct {
	Name string `json:"name"`
}

func main() {

	var metric = MetricExternalResponse{}
	metric.Name = "name"
	metric.SetName1()
	fmt.Println(metric.Name)
	metric.SetName2()
	fmt.Println(metric.Name)
}

func (m MetricExternalResponse) SetName1() {
	m.Name = "Name1"
}

func (m *MetricExternalResponse) SetName2() {
	m.Name = "Name2"
}
