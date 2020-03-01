package classfile

type ConstantClassInfo struct {
	cp ConstantPool
	nameIndex uint16
}

func (this *ConstantClassInfo) readInfo(reader *ClassReader) {
	this.nameIndex = reader.readUint16()
}

func (this *ConstantClassInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}



