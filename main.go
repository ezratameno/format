package main

import (
	"fmt"

	"github.com/ezratameno/format/pkg/format"
)

type metric struct {
	MetricName string  `label:"metric_name"`
	Value      float64 `label:"metric_value"`
	Cpu        string  `label:"label1"`
	Mem        string  `label:"label2"`
	G          int
	Sub        SubMetric
	Labels     map[string]string
}

type SubMetric struct {
	Sub string `label:"sub_metric"`
	H   Human
}
type Human struct {
	H string `label:"name"`
}

func main() {
	m := metric{
		MetricName: "metric_test",
		Value:      1.3,
		Cpu:        "CPU value",
		Mem:        "Mem value",
		Sub: SubMetric{
			Sub: "dfdf",
			H: Human{
				H: "fgdfg",
			},
		},
		Labels: map[string]string{
			"habana.ai/debug": "true",
			"habana.ai/sche":  "false",
		},
	}

	s, _ := format.FormatProm(&m)
	fmt.Println(s)

}
