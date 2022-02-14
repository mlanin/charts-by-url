package charts

import (
	"strconv"
	"time"

	"github.com/wcharczuk/go-chart"
)

func ParseTimeSeries(r Request) (Graph, error) {
	result := make([]chart.Series, len(r.Series))
	for i, series := range r.Series {
		xValues := make([]time.Time, 0)
		yValues := make([]float64, 0)
		for _, x := range series.X {
			v, err := time.Parse("2006-01-02", x)
			if err != nil {
				return nil, err
			}
			xValues = append(xValues, v)
		}

		for _, y := range series.Y {
			v, err := strconv.ParseFloat(y, 64)
			if err != nil {
				return nil, err
			}
			yValues = append(yValues, v)
		}

		color := parseStrokeColor(i, series)

		result[i] = chart.TimeSeries{
			Name: series.Name,
			Style: chart.Style{
				Show:        true,
				StrokeColor: color,
			},
			XValues: xValues,
			YValues: yValues,
		}
	}

	graph := &chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top: 50,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		Series: result,
	}

	if graph.Series[0].GetName() != "" {
		graph.Elements = []chart.Renderable{
			chart.LegendThin(graph),
		}
	}

	return graph, nil
}
