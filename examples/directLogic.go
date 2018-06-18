package main

import (
	"os"
	"time"
	"github.com/op/go-logging"
	"github.com/sergiorb/mcp23017-golang/src/device"
	"golang.org/x/exp/io/i2c"
  "fmt"
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

	var mainLog = logging.MustGetLogger("PCA9685 Demo - Direct Logic")

	i2cDevice, err := i2c.Open(&i2c.Devfs{Dev: I2C_ADDR}, ADDR_01)

	defer i2cDevice.Close()

	if err != nil {

		mainLog.Error(err)

	} else {

		var deviceLog = logging.MustGetLogger("PCA9685")

		mcp23017 := device.NewMCP23017(i2cDevice, "IO EXPANSOR", deviceLog)

		sleep := 1000

		mcp23017.SetAllPortsAsOutputs()

    // 00000010
    var logic uint = 2

    mainLog.Info(fmt.Sprintf("Logic = %08b\n", byte(logic)))

		mcp23017.SetLogicBankB(byte(logic))

    time.Sleep(time.Duration(sleep) * time.Millisecond)

    // 00000100
    logic = 4

    mainLog.Info(fmt.Sprintf("Logic = %08b\n", byte(logic)))

		mcp23017.SetLogicBankB(byte(logic))

    time.Sleep(time.Duration(sleep) * time.Millisecond)

    // 00000001
    logic = 1

    mainLog.Info(fmt.Sprintf("Logic = %08b\n", byte(logic)))

		mcp23017.SetLogicBankB(byte(logic))

    time.Sleep(time.Duration(sleep) * time.Millisecond)

    // 00000110
    logic = 6

    mainLog.Info(fmt.Sprintf("Logic = %08b\n", byte(logic)))

		mcp23017.SetLogicBankB(byte(logic))

    time.Sleep(time.Duration(sleep) * time.Millisecond)
	}
}
