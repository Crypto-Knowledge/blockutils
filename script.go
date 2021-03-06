package blockutils

import (
	"errors"
)

// Bitcoin script type backed by a byte array
// The string function is particularly helpful for working
// with the stack and getting it into a string representation
type Script []byte

func (script Script) IsOpReturn() bool {
	return script[0] == 0x6a
}

func (script Script) IsP2PK() bool {
	if len(script) != 67 && len(script) != 35 {
		return false
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}
	pushLength := uint64(scriptReader.ReadByte())
	scriptReader.ReadBytes(pushLength)
	firstOp := scriptReader.ReadByte()

	if firstOp != 0xac {
		return false
	}

	return true
}

func (script Script) P2PKHash160() ([]byte, error) {
	if len(script) != 67 && len(script) != 35 {
		return nil, errors.New("Invalid script length for P2PK")
	}

	if !script.IsP2PK() {
		return nil, errors.New("Invalid script for P2PK address")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	pushLength := uint64(scriptReader.ReadByte())
	pubkey := scriptReader.ReadBytes(pushLength)
	return Hash160(Sha256(pubkey)), nil
}

func (script Script) IsP2PKH() bool {

	if len(script) != 25 {
		return false
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	firstOp := scriptReader.ReadByte()
	secondOp := scriptReader.ReadByte()

	if firstOp != 0x76 && secondOp != 0xa9 {
		return false
	}

	pushLength := uint64(scriptReader.ReadByte())
	scriptReader.ReadBytes(pushLength)

	thirdOp := scriptReader.ReadByte()
	fourthOp := scriptReader.ReadByte()

	if thirdOp != 0x88 && fourthOp != 0xac {
		return false
	}

	return true
}

func (script Script) P2PKHHash160() ([]byte, error) {
	if len(script) != 25 {
		return nil, errors.New("Invalid script length for P2PKH")
	}

	if !script.IsP2PKH() {
		return nil, errors.New("Invalid script for P2PKH address")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	scriptReader.ReadBytes(2)
	pushLength := uint64(scriptReader.ReadByte())
	return scriptReader.ReadBytes(pushLength), nil
}

func (script Script) IsP2SH() bool {
	if len(script) != 23 {
		return false
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	firstOp := scriptReader.ReadByte()
	if firstOp != 0xa9 {
		return false
	}

	pushLength := uint64(scriptReader.ReadByte())
	scriptReader.ReadBytes(pushLength)

	secondOp := scriptReader.ReadByte()

	if secondOp != 0x87 {
		return false
	}

	return true

}

func (script Script) P2SHHash160() ([]byte, error) {
	if len(script) != 23 {
		return nil, errors.New("Invalid script length for P2SH")
	}

	if !script.IsP2SH() {
		return nil, errors.New("Invalid script for P2SH address")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	scriptReader.ReadByte()
	pushLength := uint64(scriptReader.ReadByte())
	return scriptReader.ReadBytes(pushLength), nil
}

func (script Script) IsWitnessScript() bool {
	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	witnessVersion := scriptReader.ReadByte()
	if witnessVersion != 0x00 {
		return false
	}

	pushLengthByte := scriptReader.ReadByte()
	if pushLengthByte != 0x20 && pushLengthByte != 0x14 {
		return false
	}

	return true
}

func (script Script) WitnessVersion() (byte, error) {
	if !script.IsWitnessScript() {
		return 0xFF, errors.New("Invalid witness program")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	return scriptReader.ReadByte(), nil
}

func (script Script) WitnessProgram() ([]byte, error) {
	if !script.IsWitnessScript() {
		return nil, errors.New("Invalid witness program")
	}

	scriptReader := ByteReader{
		Bytes:  script,
		Cursor: 0,
	}

	scriptReader.ReadByte()
	pushLength := uint64(scriptReader.ReadByte())
	return scriptReader.ReadBytes(pushLength), nil
}
