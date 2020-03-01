package classfile

type MarkerAttribute struct {

}

func (*MarkerAttribute) readInfo(reader *ClassReader) {
	// empty
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

type SyntheticAttribute struct {
	MarkerAttribute
}



