package charts

import (
	"context"
	"reflect"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"
)

func TestTimeBuilder_Build(t *testing.T) {
	type fields struct {
		lines   []Line
		periods []int
	}
	tests := []struct {
		name   string
		fields fields
		want   Request
	}{
		{
			name: "simple_chart",
			fields: fields{
				lines: []Line{
					NewSimpleChartLine(map[int]int{7: 3, 14: 4}, "red", "first line"),
					NewSimpleChartLine(map[int]int{7: 5, 14: 6}, "green", "second line"),
					NewSimpleChartLine(map[int]int{7: 1, 14: 12}, "blue", "third line"),
				},
				periods: []int{7, 14},
			},
			want: Request{
				Type: TypeTime,
				Series: []Series{
					{Name: "first line", Color: "red", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"3", "4"}},
					{Name: "second line", Color: "green", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"5", "6"}},
					{Name: "third line", Color: "blue", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"1", "12"}},
				},
			},
		},
		{
			name: "holed_chart",
			fields: fields{
				lines: []Line{
					NewSimpleChartLine(map[int]int{7: 3, 14: 4}, "red", "first line"),
					NewSimpleChartLine(map[int]int{14: 6}, "green", "second line"),
					NewSimpleChartLine(map[int]int{7: 1}, "blue", "third line"),
				},
				periods: []int{7, 14},
			},
			want: Request{
				Type: TypeTime,
				Series: []Series{
					{Name: "first line", Color: "red", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"3", "4"}},
					{Name: "second line", Color: "green", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"0", "6"}},
					{Name: "third line", Color: "blue", X: []string{"2021-04-08", "2021-04-01"}, Y: []string{"1", "0"}},
				},
			},
		},
	}

	NowTimeFunc = func() time.Time { return time.Date(2021, 4, 15, 12, 0, 0, 0, time.UTC) }
	for _, data := range tests {
		tt := data
		t.Run(tt.name, func(t *testing.T) {
			b := NewTimeBuilder(log.New(), tt.fields.lines, tt.fields.periods)
			if got := b.Build(context.Background()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
