package odbc

type OutputParam struct {
	buffer interface{}
}

func OUTPUT_PARAM(buffer interface{}) *OutputParam {
	/*
		if reflect.ValueOf(buffer).Kind() == reflect.Ptr {
		}
	*/
	return &OutputParam{
		buffer: buffer,
	}
}

func (op *OutputParam) GetValue() interface{} {
	return op.buffer
}
