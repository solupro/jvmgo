package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (this *ConstantValueAttribute) readInfo(reader *ClassReader) {
	this.constantValueIndex = reader.readUint16()
}

func (this *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return this.constantValueIndex
}


