package charts

import (
	"context"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type Line interface {
	Y(x int) int
	Color() string
	LegendName() string
}

type TimeBuilder struct {
	log     *log.Logger
	lines   []Line
	periods []int
}

var NowTimeFunc = time.Now

func NewTimeBuilder(log *log.Logger, lines []Line, periods []int) *TimeBuilder {
	return &TimeBuilder{log: log, lines: lines, periods: periods}
}

func (b *TimeBuilder) Build(ctx context.Context) Request {
	series := make([]Series, 0, len(b.lines))
	for _, l := range b.lines {
		series = append(series, Series{
			Name:  l.LegendName(),
			Color: l.Color(),
			X:     make([]string, len(b.periods)),
			Y:     make([]string, len(b.periods)),
		})
	}

	chart := Request{
		Type:   TypeTime,
		Series: series,
	}

	for i, period := range b.periods {
		date := NowTimeFunc().Add(time.Duration(-period*24) * time.Hour).Format("2006-01-02")
		for j, l := range b.lines {
			chart.Series[j].X[i] = date
			chart.Series[j].Y[i] = strconv.Itoa(l.Y(period))
		}
	}

	return chart
}

type SimpleChartLine struct {
	data       map[int]int
	color      string
	legendName string
}

func NewSimpleChartLine(data map[int]int, color string, legendName string) *SimpleChartLine {
	return &SimpleChartLine{data: data, color: color, legendName: legendName}
}

func (l *SimpleChartLine) Y(x int) int {
	return l.data[x]
}

func (l *SimpleChartLine) Color() string {
	return l.color
}

func (l *SimpleChartLine) LegendName() string {
	return l.legendName
}
