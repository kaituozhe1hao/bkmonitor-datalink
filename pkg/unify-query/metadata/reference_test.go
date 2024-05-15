// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package metadata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/featureFlag"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
)

type checkExpected struct {
	ok           bool
	vmRtList     []string
	vmConditions map[string]string
}

func TestCheckVmQuery(t *testing.T) {
	ctx := context.Background()

	InitMetadata()
	log.InitTestLogger()

	err := featureFlag.MockFeatureFlag(
		ctx, `{
	"vm-query": {
		"variations": {
			"Default": false,
			"true": true,
			"false": false
		},
		"targeting": [{
			"query": "spaceUid in [\"vm-query\"]",
			"percentage": {
				"true": 100,
				"false": 0
			}
		}],
		"defaultRule": {
			"variation": "Default"
		}
	}
}`,
	)
	assert.Nil(t, err)

	refNameA := "a"
	refNameB := "b"

	tt := []struct {
		name     string
		ref      QueryReference
		spaceUid string
		expected checkExpected

		source string
	}{
		{
			name:     "测试单一查询符合 druid-query 双维度条件",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_net_raw",
							VmCondition:    `result_table_id="100147_ieod_system_net_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_cloud_id",
										"bk_obj_id",
										"bk_biz_id",
										"bk_inst_id",
										"bcs_cluster_id",
										"namespace",
										"pod",
										"container",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_net_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_net_cmdb",
				},
			},
		},
		{
			name:     "测试单一查询 conditions 符合 druid-query 双维度条件",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:                  "system",
							Measurement:         "cpu_detail",
							Field:               "usage",
							IsSingleMetric:      false,
							VmRt:                "100147_ieod_system_net_raw",
							Condition:           "(bk_inst_id='test' and bk_obj_id='demo') or bk_biz_id='test-1'",
							AggregateMethodList: []AggrMethod{},
							VmCondition:         `result_table_id="100147_ieod_system_net_raw", __name__="usage_value", bk_inst_id="test", bk_obj_id="demo" or result_table_id="100147_ieod_system_net_raw", __name__="usage_value", bk_biz_id="test-1"`,
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_net_cmdb", __name__="usage_value", bk_inst_id="test", bk_obj_id="demo" or result_table_id="100147_ieod_system_net_cmdb", __name__="usage_value", bk_biz_id="test-1"`,
				},
				vmRtList: []string{
					"100147_ieod_system_net_cmdb",
				},
			},
		},
		{
			name:     "测试单一查询开启 druid-query 特性开关，单维度",
			spaceUid: "test",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_net_raw",
							VmCondition:    `result_table_id="100147_ieod_system_net_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_cloud_id",
										"bk_biz_id",
										"bk_inst_id",
										"bcs_cluster_id",
										"namespace",
										"pod",
										"container",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_net_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_net_cmdb",
				},
			},
		},
		{
			name:     "测试多个符合 druid-query 的查询 - 2",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_summary_raw",
							Condition:      "bk_obj_id = '1' and bk_inst_id = '2'",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value",bk_obj_id="1", bk_inst_id="2"`,
							AggregateMethodList: []AggrMethod{
								{
									Name:       "sum",
									Dimensions: []string{},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_cmdb", __name__="usage_value" or result_table_id="100147_ieod_system_summary_cmdb", __name__="usage_value", bk_obj_id="1", bk_inst_id="2"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_cmdb",
					"100147_ieod_system_summary_cmdb",
				},
			},
		},
		{
			name:     "测试非单指标单表 vm 查询",
			spaceUid: "vm-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
									},
								},
							},
						},
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_summary_raw",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_cmdb", __name__="usage_value" or result_table_id="100147_ieod_system_summary_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_cmdb",
					"100147_ieod_system_summary_cmdb",
				},
			},
		},
		{
			name:     "测试单指标单表 vm 查询",
			spaceUid: "vm-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
									},
								},
							},
						},
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_summary_raw",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_cmdb", __name__="usage_value" or result_table_id="100147_ieod_system_summary_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_cmdb",
					"100147_ieod_system_summary_cmdb",
				},
			},
		},
		{
			name:     "测试多指标符合 druid-query 查询",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
				refNameB: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_summary_raw",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameB,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_cmdb", __name__="usage_value"`,
					refNameB: `result_table_id="100147_ieod_system_summary_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_cmdb",
					"100147_ieod_system_summary_cmdb",
				},
			},
		},
		{
			name:     "测试多指标多聚合符合 druid-query 查询",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
								{
									Name: "count",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
				refNameB: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "2_vm_system_mem",
							VmCondition:    `result_table_id="2_vm_system_mem", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
								{
									Name: "max",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameB,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_cmdb", __name__="usage_value"`,
					refNameB: `result_table_id="2_vm_system_mem_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_cmdb",
					"2_vm_system_mem_cmdb",
				},
			},
		},
		{
			name:     "测试不同环境下多指标多聚合符合 druid-query 查询",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "2_vm_system_cpu_detail",
							VmCondition:    `result_table_id="2_vm_system_cpu_detail", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
								{
									Name: "count",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
				refNameB: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "2_vm_system_mem",
							VmCondition:    `result_table_id="2_vm_system_mem", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
								{
									Name: "max",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameB,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="2_vm_system_cpu_detail_cmdb", __name__="usage_value"`,
					refNameB: `result_table_id="2_vm_system_mem_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"2_vm_system_cpu_detail_cmdb",
					"2_vm_system_mem_cmdb",
				},
			},
		},
		{
			name:     "测试多指标符合的 druid 和 vm 混合查询",
			spaceUid: "druid-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name:       "sum",
									Dimensions: []string{},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
				refNameB: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: false,
							VmRt:           "100147_ieod_system_summary_raw",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name: "sum",
									Dimensions: []string{
										"bk_obj_id",
										"bk_inst_id",
									},
								},
							},
						},
					},
					ReferenceName: refNameB,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
					refNameB: `result_table_id="100147_ieod_system_summary_cmdb", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_raw",
					"100147_ieod_system_summary_cmdb",
				},
			},
		},
		{
			name:     "测试多指标不符合的 vm 查询",
			spaceUid: "vm-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name:       "sum",
									Dimensions: []string{},
								},
							},
						},
					},
					ReferenceName: refNameA,
				},
				refNameB: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_summary",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_summary_raw",
							VmCondition:    `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
							AggregateMethodList: []AggrMethod{
								{
									Name:       "sum",
									Dimensions: []string{},
								},
							},
						},
					},
					ReferenceName: refNameB,
				},
			},
			expected: checkExpected{
				ok: true,
				vmConditions: map[string]string{
					refNameA: `result_table_id="100147_ieod_system_detail_raw", __name__="usage_value"`,
					refNameB: `result_table_id="100147_ieod_system_summary_raw", __name__="usage_value"`,
				},
				vmRtList: []string{
					"100147_ieod_system_detail_raw",
					"100147_ieod_system_summary_raw",
				},
			},
		},
		{
			name:     "测试 conditions 转义问题",
			spaceUid: "vm-query",
			ref: QueryReference{
				refNameA: &QueryMetric{
					QueryList: []*Query{
						{
							DB:             "system",
							Measurement:    "cpu_detail",
							Field:          "usage",
							IsSingleMetric: true,
							VmRt:           "100147_ieod_system_detail_raw",
							VmCondition:    `p1="{\"moduleType\":3}", result_table_id="table_id"`,
						},
					},
					ReferenceName: refNameA,
				},
			},
			expected: checkExpected{
				ok: true,
				vmRtList: []string{
					"100147_ieod_system_detail_raw",
				},
				vmConditions: map[string]string{
					refNameA: `p1="{\"moduleType\":3}", result_table_id="table_id"`,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx = InitHashID(ctx)

			SetUser(ctx, tc.source, tc.spaceUid, "")

			ok, vmExpand, err := tc.ref.CheckVmQuery(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tc.expected.ok, ok)
			if tc.expected.vmConditions != nil {
				assert.Equal(t, tc.expected.vmConditions, vmExpand.MetricFilterCondition)
			}
			if tc.expected.vmRtList != nil {
				assert.Equal(t, tc.expected.vmRtList, vmExpand.ResultTableList)
			}
		})
	}

}
