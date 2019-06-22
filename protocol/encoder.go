package protocol

type Encoder interface {
	PutBool(v bool)
	PutInt8(v int8)
	PutInt16(v int16)
	PutInt32(v int32)
	PutInt64(v int64)
	PutUint32(v uint32)
	// String utf8字符串并且长度一定是int16
	PutString(v string)
	// NullableString utf8字符串并且长度一定是int16,空值的编码长度为-1,并且没有后续字节
	PutNullableString(v *string)
	// Bytes 长度为int32给出.然后是N个字节
	PutBytes(v []byte)
	// NullableBytes 表示原始字节序列或null, 长度为int32给出,空值的编码长度为-1,并且没有后续字节
	PutNullableBytes(v []byte)
}

type ByteEncoder struct {
	offset int32
	b []byte
}

func (e *ByteEncoder) PutBool(v bool) {
	e.b[e.offset] = 0
	if v {
		e.b[e.offset] = byte(1)
	}
	e.offset ++
}

func (e *ByteEncoder) PutInt8(v int8) {
	e.b[e.offset] = byte(v)
	e.offset ++
}

func (e *ByteEncoder) PutInt16(v int16) {
	e.b[e.offset] = byte( v >> 8)
	e.offset ++
	e.b[e.offset] = byte(v)
	e.offset ++
}

func (e *ByteEncoder) PutInt32(v int32) {
	e.PutInt16(int16(v >> 16))
	e.PutInt16(int16(v))
}

func (e *ByteEncoder) PutInt64(v int64) {
	e.PutInt32(int32(v >> 32))
	e.PutInt32(int32(v))
}

func (e *ByteEncoder) PutUint32(v uint32) {
	e.PutInt32(int32(v))
}

func (e *ByteEncoder) PutString(v string) {
	l := int16(len(v))
	e.PutInt16(l)
	copy(e.b, v)
	e.offset += int32(l)
}

func (e *ByteEncoder) PutNullableString(v *string) {
	if v == nil {
		e.PutInt16(-1)
		return
	}
	e.PutString(*v)
}

func (e *ByteEncoder) PutBytes(v []byte) {
	l := int32(len(v))
	e.PutInt32(l)
	copy(e.b, v)
	e.offset += l
}

func (e *ByteEncoder) PutNullableBytes(v []byte) {
	if v == nil {
		e.PutInt32(-1)
		return
	}
	e.PutBytes(v)
}