drop database if exists sqlw_test;
create database sqlw_test;

drop table if exists sqlw_test.sqlw_test;
create table sqlw_test.sqlw_test (
	id bigint primary key auto_increment,

	field_bool BOOLEAN,

	field_int INTEGER,
	field_int8 TINYINT,
	field_int16 SMALLINT,
	field_int32 INTEGER,
	field_int64 BIGINT,

	field_uint INTEGER,
	field_uint8 TINYINT,
	field_uint16 SMALLINT,
	field_uint32 INTEGER,
	field_uint64 BIGINT,

	field_float32 FLOAT,
	field_float64 DOUBLE PRECISION,
	field_decimal DECIMAL(10,4),

	field_date DATE,
	field_time TIME,
	field_timestamp TIMESTAMP(3),

	field_char CHAR(32),
	field_varchar VARCHAR(32),
	field_text TEXT,
	field_binary BINARY(64)
);
