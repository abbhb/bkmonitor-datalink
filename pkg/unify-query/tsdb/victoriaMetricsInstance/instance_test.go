// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package victoriaMetricsInstance

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/curl"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
)

const (
	TestTime  = "2022-11-28 10:00:00"
	ParseTime = "2006-01-02 15:04:05"
)

func TestInstance_Query_Url(t *testing.T) {
	log.InitTestLogger()

	mockCurl := curl.NewMockCurl(map[string]string{
		`http://127.0.0.1/api/query?query=count%28container_cpu_system_seconds_total_value%29&step=60&time=1669600800`:                                                          `{"status":"success","isPartial":false,"data":{"resultType":"vector","result":[{"metric":{},"value":[1669600800,"31949"]}]}}`,
		`http://127.0.0.1/api/query?query=count+by+%28__bk_db__%2C+bk_biz_id%2C+bcs_cluster_id%29+%28container_cpu_system_seconds_total_value%7B%7D%29&step=60&time=1669600800`: `{"status":"success","isPartial":false,"data":{"resultType":"vector","result":[{"metric":{"__bk_db__":"mydb","bcs_cluster_id":"BCS-K8S-40949","bk_biz_id":"930"},"value":[1669600800,"31949"]}]}}`,
		`http://127.0.0.1/api/query?query=sum%28111gggggggggggggggg11&step=60&time=1669600800`:                                                                                  `{"status":"error","errorType":"422","error":"error when executing query=\"sum(111gggggggggggggggg11\" for (time=1669600800000, step=60000): argList: unexpected token \"gggggggggggggggg11\"; want \",\", \")\"; unparsed data: \"gggggggggggggggg11\""}`,
		`http://127.0.0.1/api/query?query=top%28sum%28kube_pod_container_resource_limits_value%29%29&step=60&time=1669600800`:                                                   `{"status":"error","errorType":"422","error":"unknown func \"top\""}`,
	}, log.OtLogger)

	ctx := context.Background()
	ins := &Instance{
		ctx:     ctx,
		address: "http://127.0.0.1/api",
		timeout: time.Minute,
		curl:    mockCurl,
	}

	endTime, _ := time.ParseInLocation(ParseTime, TestTime, time.Local)

	testCases := map[string]struct {
		promql   string
		expected string
		err      error
	}{
		"count": {
			promql:   `count(container_cpu_system_seconds_total_value)`,
			expected: `[{"metric":{},"values":[[1669600800,"31949"]]}]`,
		},
		"count rate metric": {
			promql:   `count by (__bk_db__, bk_biz_id, bcs_cluster_id) (container_cpu_system_seconds_total_value{})`,
			expected: `[{"metric":{"__bk_db__":"mydb","bcs_cluster_id":"BCS-K8S-40949","bk_biz_id":"930"},"values":[[1669600800,"31949"]]}]`,
		},
		"error metric 1": {
			promql: `sum(111gggggggggggggggg11`,
			err:    errors.New(`error when executing query="sum(111gggggggggggggggg11" for (time=1669600800000, step=60000): argList: unexpected token "gggggggggggggggg11"; want ",", ")"; unparsed data: "gggggggggggggggg11"`),
		},
		"error metric 2": {
			promql: `top(sum(kube_pod_container_resource_limits_value))`,
			err:    errors.New(`unknown func "top"`),
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			data, err := ins.Query(ctx, c.promql, endTime)
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.Nil(t, err)
				res, err1 := json.Marshal(data)
				assert.Nil(t, err1)
				assert.Equal(t, c.expected, string(res))
			}

		})
	}
}

func TestInstance_QueryRange_Url(t *testing.T) {
	log.InitTestLogger()
	ctx := context.Background()

	mockCurl := curl.NewMockCurl(map[string]string{
		`http://127.0.0.1/api/query_range?end=1669600800&query=count%28kube_pod_container_resource_limits_value%29&start=1669600500&step=60`:                                                          `{"status":"success","isPartial":false,"data":{"resultType":"matrix","result":[{"metric":{},"values":[[1669600500,"61305"],[1669600560,"61305"],[1669600620,"61305"],[1669600680,"61311"],[1669600740,"61311"],[1669600800,"61314"]]}]}}`,
		`http://127.0.0.1/api/query_range?end=1669600800&query=count+by+%28__bk_db__%2C+bk_biz_id%2C+bcs_cluster_id%29+%28container_cpu_system_seconds_total_value%7B%7D%29&start=1669600500&step=60`: `{"status":"success","isPartial":false,"data":{"resultType":"matrix","result":[{"metric":{"__bk_db__":"mydb","bcs_cluster_id":"BCS-K8S-40949","bk_biz_id":"930"},"values":[[1669600500,"31949"],[1669600560,"31949"],[1669600620,"31949"],[1669600680,"31949"],[1669600740,"31949"],[1669600800,"31949"]]}]}}`,
		`http://127.0.0.1/api/query_range?end=1669600800&query=sum%28111gggggggggggggggg11&start=1669600500&step=60`:                                                                                  `{"status":"error","errorType":"422","error":"error when executing query=\"sum(111gggggggggggggggg11\" on the time range (start=1669600500000, end=1669600800000, step=60000): argList: unexpected token \"gggggggggggggggg11\"; want \",\", \")\"; unparsed data: \"gggggggggggggggg11\""}`,
		`http://127.0.0.1/api/query_range?end=1669600800&query=top%28sum%28kube_pod_container_resource_limits_value%29%29&start=1669600500&step=60`:                                                   `{"status":"error","errorType":"422","error":"unknown func \"top\""}`,
	}, log.OtLogger)

	ins := &Instance{
		ctx:     ctx,
		address: "http://127.0.0.1/api",
		timeout: time.Minute,
		curl:    mockCurl,
	}

	leftTime := time.Minute * -5

	endTime, _ := time.ParseInLocation(ParseTime, TestTime, time.Local)
	startTime := endTime.Add(leftTime)
	stepTime := time.Minute

	testCases := map[string]struct {
		promql   string
		expected string
		err      error
	}{
		"count": {
			promql:   `count(kube_pod_container_resource_limits_value)`,
			expected: `[{"metric":{},"values":[[1669600500,"61305"],[1669600560,"61305"],[1669600620,"61305"],[1669600680,"61311"],[1669600740,"61311"],[1669600800,"61314"]]}]`,
		},
		"count rate metric": {
			promql:   `count by (__bk_db__, bk_biz_id, bcs_cluster_id) (container_cpu_system_seconds_total_value{})`,
			expected: `[{"metric":{"__bk_db__":"mydb","bcs_cluster_id":"BCS-K8S-40949","bk_biz_id":"930"},"values":[[1669600500,"31949"],[1669600560,"31949"],[1669600620,"31949"],[1669600680,"31949"],[1669600740,"31949"],[1669600800,"31949"]]}]`,
		},
		"error metric 1": {
			promql: `sum(111gggggggggggggggg11`,
			err:    errors.New(`error when executing query="sum(111gggggggggggggggg11" on the time range (start=1669600500000, end=1669600800000, step=60000): argList: unexpected token "gggggggggggggggg11"; want ",", ")"; unparsed data: "gggggggggggggggg11"`),
		},
		"error metric 2": {
			promql: `top(sum(kube_pod_container_resource_limits_value))`,
			err:    errors.New(`unknown func "top"`),
		},
	}

	for name, c := range testCases {
		t.Run(name, func(t *testing.T) {
			data, err := ins.QueryRange(ctx, c.promql, startTime, endTime, stepTime)
			if c.err != nil {
				assert.Equal(t, c.err, err)
			} else {
				assert.Nil(t, err)
				res, err1 := json.Marshal(data)
				assert.Nil(t, err1)
				assert.Equal(t, c.expected, string(res))
			}

		})
	}

}
