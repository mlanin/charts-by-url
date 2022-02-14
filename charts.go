package charts

import (
	"encoding/json"
	"io"
	"net/url"

	"github.com/wcharczuk/go-chart"
)

const (
	BaseUrl        string = "https://bots.avito.ru/sheldon/chart?"
	TypeTime       string = "time"
	TypeContinuous string = "cont"
)

type Series struct {
	Name  string   `json:"l,omitempty"`
	Color string   `json:"c,omitempty"`
	X     []string `json:"x"`
	Y     []string `json:"y"`
}

type Request struct {
	Type   string   `json:"t"`
	Series []Series `json:"s"`
}

func (c Request) ToURL() string {
	chart, _ := json.Marshal(c)
	values := url.Values{"q": {string(chart)}}

	return BaseUrl + values.Encode()
}

type Graph interface {
	Render(rp chart.RendererProvider, w io.Writer) error
}
