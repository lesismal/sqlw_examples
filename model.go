package examples

type ModelForTest struct {
	Id int64 `db:"id"`

	FieldBool bool `db:"field_bool"`

	FieldInt   int   `db:"field_int"`
	FieldInt8  int8  `db:"field_int8"`
	FieldInt16 int16 `db:"field_int16"`
	FieldInt32 int32 `db:"field_int32"`
	FieldInt64 int64 `db:"field_int64"`

	FieldUint   uint   `db:"field_uint"`
	FieldUint8  uint8  `db:"field_uint8"`
	FieldUint16 uint16 `db:"field_uint16"`
	FieldUint32 uint32 `db:"field_uint32"`
	FieldUint64 uint64 `db:"field_uint64"`

	FieldFloat32 float32 `db:"field_float32"`
	FieldFloat64 float64 `db:"field_float64"`
	FieldDecimal float64 `db:"field_decimal"`

	FieldDate      string `db:"field_date"`
	FieldTime      string `db:"field_time"`
	FieldTimestamp string `db:"field_timestamp"`

	FieldChar    string `db:"field_char"`
	FieldVarchar string `db:"field_varchar"`
	FieldText    string `db:"field_text"`
	FieldBinary  []byte `db:"field_binary"`
}
