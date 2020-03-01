package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
	this.descriptorIndex = reader.readUint16()
}



