package main

import (
	"os"
	"time"
	"github.com/op/go-logging"
	"github.com/sergiorb/mcp23017-golang/src/device"
	"golang.org/x/exp/io/i2c"
  // "fmt"
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

    // 1011101111001100
    var logic uint16 = 0xBBCC

    mcp23017.SetLogic(logic)
    time.Sleep(time.Duration(sleep) * time.Millisecond)

    logic ^= 0xFFFF

    mcp23017.SetLogic(logic)
    time.Sleep(time.Duration(sleep) * time.Millisecond)

    logic ^= 0xFFFF

    mcp23017.SetLogic(logic)
    time.Sleep(time.Duration(sleep) * time.Millisecond)

    logic ^= 0xFFFF

    mcp23017.SetLogic(logic)
    time.Sleep(time.Duration(sleep) * time.Millisecond)

    mcp23017.SetLogic(0x0000)
	}
}
