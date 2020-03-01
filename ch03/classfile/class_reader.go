package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (this *ClassReader) readUint8() uint8 {
	val := this.data[0]
	this.data = this.data[1:]

	return val
}

func (this *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(this.data)
	this.data = this.data[2:]

	return val
}

func (this *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(this.data)
	this.data = this.data[4:]

	return val
}

func (this *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(this.data)
	this.data = this.data[8:]

	return val
}

func (this *ClassReader) readUint16s() []uint16 {
	n := this.readUint16()
	table := make([]uint16, n)

	for i := range table {
		table[i] = this.readUint16()
	}

	return table;
}

func (this *ClassReader) readBytes(length uint32) []byte {
	bytes := this.data[:length]
	this.data = this.data[length:]

	return bytes
}

