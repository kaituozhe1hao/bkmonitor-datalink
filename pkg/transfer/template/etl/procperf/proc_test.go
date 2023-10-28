// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package procperf_test

import (
	"context"
	_ "embed"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/transfer/define"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/transfer/template/etl/procperf"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/transfer/testsuite"
)

// ProcPerfTest
type ProcPerfTest struct {
	testsuite.ETLSuite
}

//go:embed fixture/proc_test_data.json
var procPerfData string

// TestUsage :
func (s *ProcPerfTest) TestUsage() {
	s.RunN(19,
		procPerfData,
		procperf.NewPerformanceProcProcessor(s.CTX, "test"),
		func(result map[string]interface{}) {
			var pid string
			if deviceValue, ok := result["dimensions"].(map[string]interface{})["pgid"].(string); ok {
				pid = deviceValue
			} else {
				panic(fmt.Errorf("convert fail no pgid found in %v", result))
			}
			s.EqualRecord(result, map[string]interface{}{
				"8942": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_adminserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8942",
						"pid":                "8942",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_adminserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0003,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          12.0,
						"io_read_bytes":   13230706688.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  3710976.0,
						"io_write_speed":  0.0,
						"mem_res":         14614528.0,
						"mem_usage_pct":   0.0009,
						"mem_virt":        331825152.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"5997": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "zk-java",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "org.apache.zookeeper.server.quorum.QuorumPeerMain",
						"pgid":               "5997",
						"pid":                "6162",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "java",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0027,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          38.0,
						"io_read_bytes":   2576191488.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  39948386304.0,
						"io_write_speed":  13448.640273,
						"mem_res":         158474240.0,
						"mem_usage_pct":   0.0095,
						"mem_virt":        7815032832.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"7066": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "redis-server",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        ":6379",
						"pgid":               "7066",
						"pid":                "7066",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "redis-server",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0053,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          149.0,
						"io_read_bytes":   2819203072.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  8192.0,
						"io_write_speed":  0.0,
						"mem_res":         19648512.0,
						"mem_usage_pct":   0.0012,
						"mem_virt":        190590976.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"7070": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "redis-sentinel",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "sentinel",
						"pgid":               "7070",
						"pid":                "7070",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "redis-server",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0018,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          83.0,
						"io_read_bytes":   2163552256.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  1077248.0,
						"io_write_speed":  0.0,
						"mem_res":         8548352.0,
						"mem_usage_pct":   0.0005,
						"mem_virt":        150745088.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"7458": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "mytest1_exporter",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "7458",
						"pid":                "7460",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "mytest1_exporter",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          8.0,
						"io_read_bytes":   2505338880.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  0.0,
						"io_write_speed":  0.0,
						"mem_res":         10113024.0,
						"mem_usage_pct":   0.0006,
						"mem_virt":        17113088.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"4771": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "mongod",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "4771",
						"pid":                "4772",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "mongod",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      10.0,
						"cpu_total_ticks": 12.0,
						"cpu_usage_pct":   0.0078,
						"cpu_user":        11.0,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          181.0,
						"io_read_bytes":   7360700416.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  14865457152.0,
						"io_write_speed":  5256.531578,
						"mem_res":         175132672.0,
						"mem_usage_pct":   0.0105,
						"mem_virt":        1626701824.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"6034": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "job-java",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "job.conf",
						"pgid":               "6034",
						"pid":                "6206",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "1",
						"proc_name":          "java",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0052,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          77.0,
						"io_read_bytes":   8387293184.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  12891226112.0,
						"io_write_speed":  1774.949073,
						"mem_res":         4166598656.0,
						"mem_usage_pct":   0.2501,
						"mem_virt":        8253329408.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"4867": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "consul-agent",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "4867",
						"pid":                "4867",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "4837",
						"proc_name":          "consul",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.014,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          22.0,
						"io_read_bytes":   32594833408.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  168698781696.0,
						"io_write_speed":  273067.424770,
						"mem_res":         81248256.0,
						"mem_usage_pct":   0.00490,
						"mem_virt":        184864768.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"11230": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_webserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "11230",
						"pid":                "11230",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_webserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0015,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          13.0,
						"io_read_bytes":   4060712960.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  3543040.0,
						"io_write_speed":  0.0,
						"mem_res":         43462656.0,
						"mem_usage_pct":   0.0026,
						"mem_virt":        438083584.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8931": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_toposerver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8931",
						"pid":                "8931",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_toposerver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0018,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          16.0,
						"io_read_bytes":   14256435200.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  4415488.0,
						"io_write_speed":  0.0,
						"mem_res":         44421120.0,
						"mem_usage_pct":   0.0027,
						"mem_virt":        211881984.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"10961": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_procserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "10961",
						"pid":                "10961",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_procserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0022,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          13.0,
						"io_read_bytes":   7959019520.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  1699840.0,
						"io_write_speed":  0.0,
						"mem_res":         31977472.0,
						"mem_usage_pct":   0.0019,
						"mem_virt":        405471232.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8952": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_proccontroller",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8952",
						"pid":                "8952",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_proccontroller",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0015,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          14.0,
						"io_read_bytes":   14398783488.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  12562432.0,
						"io_write_speed":  0.0,
						"mem_res":         38715392.0,
						"mem_usage_pct":   0.0023,
						"mem_virt":        406814720.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"11229": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_objectcontroller",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "11229",
						"pid":                "11229",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_objectcontroller",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0298,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          15.0,
						"io_read_bytes":   11135418368.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  4632576.0,
						"io_write_speed":  0.0,
						"mem_res":         41058304.0,
						"mem_usage_pct":   0.0025,
						"mem_virt":        333185024.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8930": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_hostserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8930",
						"pid":                "8930",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_hostserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0017,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          14.0,
						"io_read_bytes":   14267944960.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  155312128.0,
						"io_write_speed":  0.0,
						"mem_res":         37367808.0,
						"mem_usage_pct":   0.0022,
						"mem_virt":        476876800.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8929": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_hostcontroller",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8929",
						"pid":                "8929",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_hostcontroller",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0042,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          15.0,
						"io_read_bytes":   14568337408.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  3014656.0,
						"io_write_speed":  0.0,
						"mem_res":         38199296.0,
						"mem_usage_pct":   0.0023,
						"mem_virt":        410816512.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"13900": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_eventserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "13900",
						"pid":                "13900",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_eventserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.002,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          20.0,
						"io_read_bytes":   12224454656.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  23879680.0,
						"io_write_speed":  273.067960,
						"mem_res":         34099200.0,
						"mem_usage_pct":   0.002,
						"mem_virt":        332111872.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8941": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_datacollection",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8941",
						"pid":                "8941",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_datacollection",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0017,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          33.0,
						"io_read_bytes":   9053741056.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  1519407104.0,
						"io_write_speed":  682.671445,
						"mem_res":         37629952.0,
						"mem_usage_pct":   0.0023,
						"mem_virt":        339812352.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8935": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_auditcontroller",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8935",
						"pid":                "8935",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_auditcontroller",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0015,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          15.0,
						"io_read_bytes":   14459097088.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  4280320.0,
						"io_write_speed":  0.0,
						"mem_res":         36372480.0,
						"mem_usage_pct":   0.0022,
						"mem_virt":        272515072.0,
						"uptime":          nil,
					},
					"time": 1553480438,
				},
				"8936": map[string]interface{}{
					"dimensions": map[string]interface{}{
						"bk_supplier_id":     "0",
						"display_name":       "cmdb_apiserver",
						"hostname":           "VM_1_10_centos",
						"ip":                 "127.0.0.1",
						"bk_target_ip":       "127.0.0.1",
						"param_regex":        "",
						"pgid":               "8936",
						"pid":                "8936",
						"bk_cloud_id":        "0",
						"bk_target_cloud_id": "0",
						"bk_agent_id":        "010000525400c48bdc1670385834306k",
						"bk_biz_id":          "2",
						"bk_host_id":         "30145",
						"bk_target_host_id":  "30145",
						"ppid":               "8910",
						"proc_name":          "cmdb_apiserver",
						"state":              "sleeping",
						"username":           "root",
						"port":               nil,
					},
					"metrics": map[string]interface{}{
						"cpu_system":      nil,
						"cpu_total_ticks": nil,
						"cpu_usage_pct":   0.0023,
						"cpu_user":        nil,
						"fd_limit_hard":   102400.0,
						"fd_limit_soft":   102400.0,
						"fd_num":          16.0,
						"io_read_bytes":   14529892352.0,
						"io_read_speed":   0.0,
						"io_write_bytes":  948932608.0,
						"io_write_speed":  409.60367,
						"mem_res":         47943680.0,
						"mem_usage_pct":   0.0029,
						"mem_virt":        493924352.0,
						"uptime":          4341026.0,
					},
					"time": 1553480438,
				},
			}[pid].(map[string]interface{}))
		},
	)
}

func (s *ProcPerfTest) TestDisabledBizIDs() {
	processor := procperf.NewPerformanceProcProcessor(s.CTX, "test")
	processor.DisabledBizIDs = map[string]struct{}{"0": {}}
	s.RunN(0,
		procPerfData,
		processor,
		func(result map[string]interface{}) {},
	)
}

// TestProcPerfTest :
func TestProcPerfTest(t *testing.T) {
	suite.Run(t, new(ProcPerfTest))
}

// BenchmarkPerformanceProcProcessor_Process :
func BenchmarkPerformanceProcProcessor_Process(b *testing.B) {
	testsuite.ETLBenchmarkTest(b, func(ctx context.Context, name string) define.DataProcessor {
		return procperf.NewPerformanceProcProcessor(ctx, name)
	}, []byte(procPerfData))
}