package charts

import (
	"strconv"

	"github.com/wcharczuk/go-chart"
)

func ParseContinuousSeries(r Request) (Graph, error) {
	result := make([]chart.Series, len(r.Series))
	for i, series := range r.Series {
		xValues := make([]float64, 0)
		yValues := make([]float64, 0)
		for _, x := range series.X {
			v, err := strconv.ParseFloat(x, 64)
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

		result[i] = chart.ContinuousSeries{
			Style: chart.Style{
				Show:        true,
				StrokeColor: color,
			},
			XValues: xValues,
			YValues: yValues,
		}
	}

	return &chart.Chart{
		YAxis: chart.YAxis{
			Style: chart.StyleShow(),
		},
		XAxis: chart.XAxis{
			Style: chart.StyleShow(),
		},
		Series: result,
	}, nil
}
