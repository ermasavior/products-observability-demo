package utils

type (
	traceIDType string
	dbTxType    string
)

const (
	TraceID  = traceIDType("trace_id")
	DBTxType = dbTxType("db_tx")
)
