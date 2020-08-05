package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	/*
		http.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hola kiku")
		})

		log.Fatal(http.ListenAndServe(":8080", nil))
	*/

	//r := raspi.NewAdaptor()
	adaptor := raspi.NewAdaptor()
	led := gpio.NewLedDriver(adaptor, "12")
	sensor := gpio.NewPIRMotionDriver(adaptor, "22")

	work := func() {
		/*
			gobot.Every(1*time.Second, func() {
				led.Toggle()
			})
		*/

		sensor.On(gpio.MotionDetected, func(data interface{}) {
			fmt.Println(gpio.MotionDetected)
			led.On()
		})

		sensor.On(gpio.MotionStopped, func(data interface{}) {
			fmt.Println(gpio.MotionStopped)
			led.Off()
		})
	}

	robot := gobot.NewRobot("superBot",
		//[]gobot.Connection{r},
		[]gobot.Connection{adaptor},
		[]gobot.Device{sensor, led},
		work,
	)

	robot.Start()

}
