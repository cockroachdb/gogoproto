package omitempty

func (o *OmitEmpty_Inner) Empty() bool {
	return o.Foo == 0
}
