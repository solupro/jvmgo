package classfile

type MemberInfo struct {
	cp ConstantPool
	accessFlags uint16
	nameIndex uint16
	descriptorIndex uint16
	// attributesCount uint16
	attributes []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)

	for i := range members {
		members[i] = readMember(reader, cp)
	}

	return members
}



func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp: cp,
		accessFlags: reader.readUint16(),
		nameIndex: reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: readAttributes(reader, cp),
	}
}

func (this *MemberInfo) AccessFLags() uint16 {
	return this.accessFlags
}

func (this *MemberInfo) Name() string {
	return this.cp.getUtf8(this.nameIndex)
}

func (this *MemberInfo) Descriptor() string {
	return this.cp.getUtf8(this.descriptorIndex)
}

