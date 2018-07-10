# mcp23017-golang v0.1
 A mcp23017 golang library.

## Info

This library has the intention to cover the [mcp23017](http://ww1.microchip.com/downloads/en/DeviceDoc/20001952C.pdf) features.

The examples are set to work with a raspberry pi. Tested in a [raspberry pi zero-w](https://www.raspberrypi.org/products/raspberry-pi-zero-w/).

Use `deploy_to_raspberrypi.sh` to deploy the selected example. You can configure wich example do you want to deploy in `deploy.json` file. You can configurate the raspberry ip address and deploy path in that file too.

## Api

* `NewMCP23017(i2cDevice *i2c.Device, name string, log *logging.Logger)` : Creates an Mcp23017 object.

* `(m MCP23017 GetName()` : Returns device name.

* `(m *MCP23017) SetAllPortsAsOutputs()` : Sets all
16 chips port in output mode.

* `(m *MCP23017) SetAllPortsAsInputs()` : Sets all 16 chips port in input mode.

* `(m *MCP23017) GetBankALogic()` : Returns a byte with the logic level state of the first 8 mcp23017 ports. This method does not read the real state from the chip, It's just a cache where status is stored.

* `(m *MCP23017) GetBankBLogic()` : Returns a byte with the logic level state of the last 8 mcp23017 ports. This method does not read the real state from the chip, It's just a cache where status is stored.

* `(m *MCP23017) SetBankALogic(logic byte)` : Sets logic level to the frist 8 mcp23017 ports.

* `(m *MCP23017) SetBankBLogic(logic byte)` : Sets logic level to the last 8 mcp23017 ports.

* `(m *MCP23017) GetPortLogic(port uint)` : Returns logic state of the desired port. This method does not read the real state from the chip, It's just a cache where status is stored.

* `(m *MCP23017) SetPortLogic(port uint, logic bool)` : Sets logic level to the desired port.

* `(m *MCP23017) SwichPortLogic(port uint)` : Switchs port logic level.

* `(m *MCP23017) SetLogic(logic uint16)` : Sets logic level to all 16 ports.

## Packages Needed
### Tools:
* [jq](https://stedolan.github.io/jq/)
