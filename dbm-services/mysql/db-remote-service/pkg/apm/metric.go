package metric

import "db-apm/metrics"

var (
	ErrCnt     = "err_cnt"
	ExecuteCnt = "execute_cnt"
)

var CustomMetrics = []*metrics.Metric{
	{
		ID:          ErrCnt,
		Name:        ErrCnt,
		Description: "Counter test metric",
		Type:        "counter_vec",
		Labels:      []string{"url", "method", "code"},
	},
	{
		ID:          ExecuteCnt,
		Name:        ExecuteCnt,
		Description: "Counter test metric",
		Type:        "counter_vec",
		Labels:      []string{"url", "method", "code"},
	},
}
