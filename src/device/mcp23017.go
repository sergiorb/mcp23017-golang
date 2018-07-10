package device

import (
	"fmt"
	"github.com/op/go-logging"
	"github.com/sergiorb/mcp23017-golang/src/utils"
	"golang.org/x/exp/io/i2c"
)

const (
	PORT_QUANTITY            uint = 16
	NON_BYTE                 byte = 0x00
	BYTE                     byte = 0xFF
	BANK_A_DIRECTION_ADDRESS byte = 0x00
	BANK_A_LOGIC_ADDRESS     byte = 0x14
	BANK_B_DIRECTION_ADDRESS byte = 0x01
	BANK_B_LOGIC_ADDRESS     byte = 0x15
)

var bankACurrentDirection byte = 0x00
var bankBCurrentDirection byte = 0x00
var bankACurrentLogic byte = 0x00
var bankBCurrentLogic byte = 0x00

type MCP23017 struct {
	i2cBus    *i2c.Device
	name      string
	initiated bool
	log       *logging.Logger
}

func NewMCP23017(i2cDevice *i2c.Device, name string, log *logging.Logger) *MCP23017 {

	log.Info(fmt.Sprintf("Creating a new MCP23017 device. Alias: %v", name))

	return &MCP23017{
		i2cBus:    i2cDevice,
		name:      name,
		log:       log,
		initiated: false,
	}
}

func (m *MCP23017) GetBankALogic() byte {
	return bankACurrentLogic
}

func (m *MCP23017) GetBankBLogic() byte {
	return bankBCurrentLogic
}

func (m *MCP23017) GetPortLogic(port uint) bool {

	if port < PORT_QUANTITY/2 {

		return !utils.HasBit(int(bankACurrentLogic), port)

	} else {

		return !utils.HasBit(int(bankBCurrentLogic), utils.GetBPort(port))
	}
}

func (m *MCP23017) SetPortLogic(port uint, logic bool) {

	if port < PORT_QUANTITY/2 {

		m.setPortLogicBankA(port, logic)

	} else {

		m.setPortLogicBankB(port, logic)
	}
}

func (m *MCP23017) SwichPortLogic(port uint) {

	if port < PORT_QUANTITY/2 {

		m.switchPortLogicBankA(port)

	} else {

		m.switchPortLogicBankB(port)
	}
}

func (m * MCP23017) SetLogic(logic uint16)  {

	m.SetLogicBankA(uint8(logic))
	m.SetLogicBankB(uint8(logic >> 8))
}

func (m * MCP23017) SetLogicBankA(logic byte) {

	bankACurrentLogic = logic

	m.write8(BANK_A_LOGIC_ADDRESS, int(bankACurrentLogic))
}

func (m * MCP23017) SetLogicBankB(logic byte) {

	bankBCurrentLogic = logic

	m.write8(BANK_B_LOGIC_ADDRESS, int(bankBCurrentLogic))
}

func (m *MCP23017) setPortLogicBankA(port uint, logic bool) {

	var addressLogic byte

	addressLogic = bankACurrentLogic

	if logic {

		addressLogic = byte(utils.SetBit(int(addressLogic), port))

	} else {

		addressLogic = byte(utils.ClearBit(int(addressLogic), port))
	}

	bankACurrentLogic = addressLogic

	m.write8(BANK_A_LOGIC_ADDRESS, int(addressLogic))
}

func (m *MCP23017) setPortLogicBankB(port uint, logic bool) {

	var addressLogic byte

	addressLogic = bankBCurrentLogic

	if logic {

		addressLogic = byte(utils.SetBit(int(addressLogic), utils.GetBPort(port)))

	} else {

		addressLogic = byte(utils.ClearBit(int(addressLogic), utils.GetBPort(port)))
	}

	bankBCurrentLogic = addressLogic

	m.write8(BANK_B_LOGIC_ADDRESS, int(addressLogic))
}

func (m *MCP23017) switchPortLogicBankA(port uint) {

	m.setPortLogicBankA(port, !utils.HasBit(int(bankACurrentLogic), port))
}

func (m *MCP23017) switchPortLogicBankB(port uint) {

	m.setPortLogicBankB(port, !utils.HasBit(int(bankBCurrentLogic), port))
}

func (m *MCP23017) SetAllPortsAsOutputs() {

	m.write8(BANK_A_DIRECTION_ADDRESS, 0x00)
	m.write8(BANK_B_DIRECTION_ADDRESS, 0x00)
}

func (m *MCP23017) SetAllPortsAsInput() {

	m.write8(BANK_A_DIRECTION_ADDRESS, 0xFF)
	m.write8(BANK_B_DIRECTION_ADDRESS, 0xFF)
}

func (m *MCP23017) GetName() string {

	return m.name
}

func (m *MCP23017) write8(reg byte, intVal int) {

	byteVal := byte(intVal) & BYTE

	m.writeByte(reg, byteVal)
}

func (m *MCP23017) read8(reg byte) int {

	byteVal := m.readByte(reg)

	return int(byteVal)
}

func (m *MCP23017) writeByte(reg byte, byteVal byte) {

	err := m.i2cBus.WriteReg(reg, []byte{byteVal})

	if err != nil {

		m.log.Error(fmt.Sprintf("Failed to read from register %#x.", reg))
		m.log.Error(err.Error())
	}

	m.log.Debug(fmt.Sprintf("Wrote %#x to register %#x.", byteVal, reg))
}

func (m *MCP23017) readByte(reg byte) byte {

	buf := make([]byte, 1)
	err := m.i2cBus.ReadReg(reg, buf)

	if err != nil {

		m.log.Error(fmt.Sprintf("Failed to read from register %#x.", reg))
		m.log.Error(err.Error())
	}

	m.log.Debug(fmt.Sprintf("Read %#x from register %#x.", buf[0], reg))

	return buf[0]
}
