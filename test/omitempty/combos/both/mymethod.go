package omitempty

func (o *OmitEmpty_Inner) IsEmpty() bool {
	return o.Foo == 0
}
