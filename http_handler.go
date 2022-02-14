package charts

import (
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/wcharczuk/go-chart"
)

type ChartHandler struct {
	log *log.Logger
}

func NewChartHandler(log *log.Logger) *ChartHandler {
	return &ChartHandler{
		log: log,
	}
}

func (h *ChartHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	graph, err := h.parseRequest(req)
	if err != nil {
		h.log.WithError(err).Errorln("failed to parse chart request")
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	resp.Header().Set("Content-Type", "image/png")
	if err := graph.Render(chart.PNG, resp); err != nil {
		h.log.WithError(err).Errorln("failed to render chart")
		resp.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *ChartHandler) parseRequest(req *http.Request) (Graph, error) {
	var r Request
	err := json.Unmarshal([]byte(req.URL.Query().Get("q")), &r)
	if err != nil {
		//h.app.Logger.WithError(err).Error("failed to parse action JSON")
		return nil, err
	}

	h.log.WithField("request", r).Info("parsed chart request")

	switch r.Type {
	case TypeTime:
		return ParseTimeSeries(r)
	case TypeContinuous:
		return ParseContinuousSeries(r)
	}

	return nil, fmt.Errorf("unknown series type")
}
