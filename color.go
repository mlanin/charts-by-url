package charts

import (
	"github.com/wcharczuk/go-chart"
	"github.com/wcharczuk/go-chart/drawing"
)

func parseStrokeColor(i int, series Series) drawing.Color {
	var color drawing.Color
	if series.Color != "" {
		color = drawing.ColorFromHex(series.Color)
	} else {
		color = chart.GetDefaultColor(i)
	}

	return color
}
