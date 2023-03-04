package logger

type logValue[T, K any] struct {
	Journey      journey
	Where        string
	Method       string
	What         string
	IsError      bool
	EntryData    T
	ResponseData K
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

func NewLog[T, K any](journey journey, method, where, what string, isError bool, data T, dataResponse K) logValue[T, K] {
	return logValue[T, K]{
		Journey:      journey,
		Method:       method,
		Where:        where,
		What:         what,
		IsError:      isError,
		EntryData:    data,
		ResponseData: dataResponse,
	}
}
