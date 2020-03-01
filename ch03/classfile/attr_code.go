package classfile

type CodeAttribute struct {
	cp ConstantPool
	maxStack uint16
	maxLocals uint16
	code []byte
	exceptionTable []*ExceptionTableEntry
	attributes []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}



func (this *CodeAttribute) readInfo(reader *ClassReader) {
	this.maxStack = reader.readUint16()
	this.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	this.code = reader.readBytes(codeLength)
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	length := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, length)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.readUint16(),
			endPc: reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}

	return exceptionTable
}
