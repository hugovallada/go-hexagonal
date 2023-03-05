package logger

type logValue[T, K, V any] struct {
	Journey      journey
	Where        string
	Method       string
	Step         string
	IsError      bool
	EntryData    T
	ResponseData K
	ErrorData    V
}

type journey struct {
	Flow    string
	SubFlow string
}

func NewJourney(flow, subFlow string) journey {
	return journey{
		Flow:    flow,
		SubFlow: subFlow,
	}
}

func NewLog[T, K, V any](journey journey, method, where, step string, isError bool, data T, dataResponse K, err V) logValue[T, K, V] {
	return logValue[T, K, V]{
		Journey:      journey,
		Method:       method,
		Where:        where,
		Step:         step,
		IsError:      isError,
		EntryData:    data,
		ResponseData: dataResponse,
		ErrorData:    err,
	}
}
