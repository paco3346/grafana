package conditions

import "github.com/grafana/grafana/pkg/tsdb"

type QueryReducer interface {
	Reduce(timeSeries *tsdb.TimeSeries) float64
}

type SimpleReducer struct {
	Type string
}

func (s *SimpleReducer) Reduce(series *tsdb.TimeSeries) float64 {
	var value float64 = 0

	switch s.Type {
	case "avg":
		for _, point := range series.Points {
			value += point[0]
		}
		value = value / float64(len(series.Points))
	}

	return value
}

func NewSimpleReducer(typ string) *SimpleReducer {
	return &SimpleReducer{Type: typ}
}