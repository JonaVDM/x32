package main

import (
	"time"

	"github.com/jonavdm/x32/pkg/x32"
)

func main() {
	// addr := "192.168.0.101:10023"
	addr := "10.0.0.21:10023"
	mixbus := 7

	c := x32.NewClient(addr)
	err := c.Connect()

	if err != nil {
		panic(err)
	}

	values := waves()

	for i := 0; i < 1000; i++ {
		for channel := 0; channel < 16; channel++ {
			index := (len(values) + i + channel*4) % len(values)
			c.SetMixbusFader(17+channel, mixbus, values[index])
		}
		time.Sleep(time.Millisecond * 75)
	}
}

func waves() []float32 {
	return []float32{0.000000,
		0.035191,
		0.070381,
		0.105572,
		0.140762,
		0.175953,
		0.211144,
		0.246334,
		0.281525,
		0.316716,
		0.351906,
		0.387097,
		0.422287,
		0.457478,
		0.492669,
		0.527859,
		0.563050,
		0.598240,
		0.633431,
		0.668622,
		0.703812,
		0.739003,
		0.774194,
		0.809384,
		0.844575,
		0.879765,
		0.914956,
		0.950147,
		0.985337,
		1.000000,
		0.964809,
		0.929619,
		0.894428,
		0.859238,
		0.824047,
		0.788856,
		0.753666,
		0.718475,
		0.683284,
		0.648094,
		0.612903,
		0.577713,
		0.542522,
		0.507331,
		0.472141,
		0.436950,
		0.401760,
		0.366569,
		0.331378,
		0.296188,
		0.260997,
		0.225806,
		0.190616,
		0.155425,
		0.120235,
		0.085044,
		0.049853,
		0.014663,
		0.000000,
	}
}
