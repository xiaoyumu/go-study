package main

type traceDataProxy struct{
	dispatcher Dispatcher
}

func NewTraceDataProxy(dispatcher Dispatcher) TraceDataHandler{
	return &traceDataProxy{ dispatcher: dispatcher }
}

func (p *traceDataProxy) ProcessTraceData(traceData *TraceMessage) error{
	return p.dispatcher.Dispatch(*traceData.Key, *traceData.Value)
}

func (p *traceDataProxy) Close(){
	if p.dispatcher == nil {
		return
	}
	p.dispatcher.Close()
}
