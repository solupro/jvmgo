package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (this *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	length := reader.readUint16()
	this.lineNumberTable = make([]*LineNumberTableEntry, length)

	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}

type LineNumberTableEntry struct {
	startPc uint16
	lineNumber uint16
}


