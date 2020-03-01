package classfile

import "math"

type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = int32(bytes)
}

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	this.val = math.Float32frombits(bytes)
}

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	this.val = math.Float64frombits(bytes)
}












