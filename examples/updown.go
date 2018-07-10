package main

import (
	"os"
	"time"
	"github.com/op/go-logging"
	"github.com/sergiorb/mcp23017-golang/src/device"
	"golang.org/x/exp/io/i2c"
)

const (
	I2C_ADDR = "/dev/i2c-1"
	ADDR_01  = 0x20
)

func init() {

	stderrorLog := logging.NewLogBackend(os.Stderr, "", 0)

	stderrorLogLeveled := logging.AddModuleLevel(stderrorLog)
	stderrorLogLeveled.SetLevel(logging.DEBUG, "")

	logging.SetBackend(stderrorLogLeveled)
}

func main() {

	var mainLog = logging.MustGetLogger("PCA9685 Demo - Updown")

	i2cDevice, err := i2c.Open(&i2c.Devfs{Dev: I2C_ADDR}, ADDR_01)

	defer i2cDevice.Close()

	if err != nil {

		mainLog.Error(err)

	} else {

		var deviceLog = logging.MustGetLogger("PCA9685")

		mcp23017 := device.NewMCP23017(i2cDevice, "IO EXPANSOR", deviceLog)

		sleep := 500

		mcp23017.SetAllPortsAsOutputs()

		for i := 0; i < 16; i++ {

			mcp23017.SetPortLogic(uint(i), true)

			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}

		for i := 15; i >= 0; i-- {

			mcp23017.SetPortLogic(uint(i), false)

			time.Sleep(time.Duration(sleep) * time.Millisecond)
		}
	}
}
