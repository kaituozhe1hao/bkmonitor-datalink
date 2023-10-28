// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package operator

import (
	"strings"

	promv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/spf13/viper"
	"k8s.io/client-go/rest"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/common/define"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/operator/config"
	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/utils/logger"
)

const (
	configDryRunPath                = "operator.dry_run"
	confAPIServerHostPath           = "operator.apiserver_host"
	confKubeConfigPath              = "operator.kube_config"
	confTLSInsecurePath             = "operator.tls.tls_insecure"
	confTLSCertFilePath             = "operator.tls.tls_cert_file"
	confTLSKeyFilePath              = "operator.tls.tls_key_file"
	confTLSCAFilePath               = "operator.tls.tls_ca_file"
	confMonitorNamespacePath        = "operator.monitor_namespace"
	confDenyTargetNamespacesPath    = "operator.deny_target_namespaces"
	confTargetNamespacesPath        = "operator.target_namespaces"
	confTargetLabelSelectorPath     = "operator.target_label_selector"
	confEnableServiceMonitorPath    = "operator.enable_service_monitor"
	confEnablePodMonitorPath        = "operator.enable_pod_monitor"
	confEnableProbePath             = "operator.enable_probe"
	confEnableStatefulSetWorkerPath = "operator.enable_statefulset_worker"
	confEnableDaemonSetWorkerPath   = "operator.enable_daemonset_worker"
	confDisableMetricsPusherPath    = "operator.disable_metrics_pusher"
	confKubeletNamespacePath        = "operator.kubelet.namespace"
	confKubeletNamePath             = "operator.kubelet.name"
	confKubeletEnablePath           = "operator.kubelet.enable"
	confMaxNodeSecretRatioPath      = "operator.node_secret_ratio"
	confStatefulSetWorkerHpaPath    = "operator.statefulset_worker_hpa"
	confStatefulSetWorkerFactorPath = "operator.statefulset_worker_factor"
	confStatefulSetReplicasPath     = "operator.statefulset_replicas"
	confStatefulSetMaxReplicasPath  = "operator.statefulset_max_replicas"
	confStatefulSetMatchRulesPath   = "operator.statefulset_match_rules"
	confStatefulSetDispatchType     = "operator.statefulset_dispatch_type"
	confMonitorBlacklistMatchRules  = "operator.monitor_blacklist_match_rules"
)

const (
	dispatchTypeHash       = "hash"
	dispatchTypeRoundrobin = "roundrobin"
)

// StatefulSetMatchRule statefulset 匹配规则
// 提供一种机制可以通知 operator 将 monitor 资源调度到 statefulset worker 上
// 1) 如果 rule 中 name 为空表示命中所有的 resource
// 2) 如果 rule 中 name 不为空则要求精准匹配
type StatefulSetMatchRule struct {
	Kind      string `mapstructure:"kind"`
	Name      string `mapstructure:"name"`
	Namespace string `mapstructure:"namespace"`
}

// MonitorBlacklistMatchRule monitor 黑名单匹配规则
// 在 monitor namespace 黑名单机制外再提供一种 name 级别的屏蔽机制
// 要求 kind/name/namespace 三者同时不为空 且此配置项优先级最高
type MonitorBlacklistMatchRule struct {
	Kind      string `mapstructure:"kind" json:"kind"`
	Name      string `mapstructure:"name" json:"name"`
	Namespace string `mapstructure:"namespace" json:"namespace"`
}

func (r MonitorBlacklistMatchRule) Validate() bool {
	return r.Kind != "" && r.Namespace != "" && r.Name != ""
}

var (
	ConfDryRun                     bool
	ConfKubeConfig                 string // operator 连接 k8s 使用的 kubeconfig 文件路径
	ConfMonitorNamespace           string // operator 所处 namespace
	ConfDenyTargetNamespaces       []string
	ConfTargetNamespaces           []string
	ConfTargetLabelsSelector       string
	ConfAPIServerHost              string
	ConfTLSConfig                  *rest.TLSClientConfig
	ConfEnableServiceMonitor       bool
	ConfEnablePodMonitor           bool
	ConfEnableStatefulSetWorker    bool
	ConfEnableDaemonSetWorker      bool
	ConfDisableMetricsPusher       bool
	ConfKubeletNamespace           string
	ConfKubeletName                string
	ConfKubeletEnable              bool
	ConfMaxNodeSecretRatio         float64
	ConfPodName                    string
	ConfStatefulSetWorkerHpa       bool
	ConfStatefulSetWorkerFactor    int
	ConfStatefulSetReplicas        int
	ConfStatefulSetMaxReplicas     int
	ConfStatefulSetMatchRules      []StatefulSetMatchRule
	ConfStatefulSetDispatchType    string
	ConfMonitorBlacklistMatchRules []MonitorBlacklistMatchRule
)

// IfRejectServiceMonitor 判断是否拒绝 serviceMonitor
func IfRejectServiceMonitor(monitor *promv1.ServiceMonitor) bool {
	if monitor == nil {
		return false
	}
	for _, rule := range ConfMonitorBlacklistMatchRules {
		if !rule.Validate() {
			continue
		}
		if strings.ToUpper(rule.Kind) == strings.ToUpper(monitor.Kind) && rule.Namespace == monitor.Namespace && rule.Name == monitor.Name {
			return true
		}
	}
	return false
}

// IfRejectPodMonitor 判断是否拒绝 podMonitor
func IfRejectPodMonitor(monitor *promv1.PodMonitor) bool {
	if monitor == nil {
		return false
	}
	for _, rule := range ConfMonitorBlacklistMatchRules {
		if !rule.Validate() {
			continue
		}
		if strings.ToUpper(rule.Kind) == strings.ToUpper(monitor.Kind) && rule.Namespace == monitor.Namespace && rule.Name == monitor.Name {
			return true
		}
	}
	return false
}

func initConfig() {
	viper.SetDefault(configDryRunPath, false)
	viper.SetDefault(confKubeConfigPath, "")
	viper.SetDefault(confAPIServerHostPath, "")
	viper.SetDefault(confTLSInsecurePath, false)
	viper.SetDefault(confTLSCertFilePath, "")
	viper.SetDefault(confTLSKeyFilePath, "")
	viper.SetDefault(confTLSCAFilePath, "")
	viper.SetDefault(confMonitorNamespacePath, "bkmonitor-operator")
	viper.SetDefault(confDenyTargetNamespacesPath, []string{})
	viper.SetDefault(confTargetNamespacesPath, []string{})
	viper.SetDefault(confTargetLabelSelectorPath, "")
	viper.SetDefault(confEnableServiceMonitorPath, true)
	viper.SetDefault(confEnablePodMonitorPath, true)
	viper.SetDefault(confEnableStatefulSetWorkerPath, true)
	viper.SetDefault(confEnableDaemonSetWorkerPath, true)
	viper.SetDefault(confEnableProbePath, false)
	viper.SetDefault(confDisableMetricsPusherPath, false)
	viper.SetDefault(confKubeletNamePath, "bkmonitor-operator-kubelet")
	viper.SetDefault(confKubeletNamespacePath, "bkmonitor-operator")
	viper.SetDefault(confKubeletEnablePath, true)
	viper.SetDefault(confMaxNodeSecretRatioPath, 2.0)
	viper.SetDefault(confStatefulSetWorkerHpaPath, true)
	viper.SetDefault(confStatefulSetWorkerFactorPath, defaultStatefulSetWorkerFactor)
	viper.SetDefault(confStatefulSetReplicasPath, 1)
	viper.SetDefault(confStatefulSetMaxReplicasPath, 10)
	viper.SetDefault(confStatefulSetDispatchType, dispatchTypeHash)
}

func updateConfig() {
	ConfDryRun = viper.GetBool(configDryRunPath)
	ConfKubeConfig = viper.GetString(confKubeConfigPath)
	ConfAPIServerHost = viper.GetString(confAPIServerHostPath)
	ConfMonitorNamespace = viper.GetString(confMonitorNamespacePath)
	ConfDenyTargetNamespaces = viper.GetStringSlice(confDenyTargetNamespacesPath)
	ConfTargetNamespaces = viper.GetStringSlice(confTargetNamespacesPath)
	ConfTargetLabelsSelector = viper.GetString(confTargetLabelSelectorPath)
	ConfEnableServiceMonitor = viper.GetBool(confEnableServiceMonitorPath)
	ConfEnablePodMonitor = viper.GetBool(confEnablePodMonitorPath)
	ConfEnableStatefulSetWorker = viper.GetBool(confEnableStatefulSetWorkerPath)
	ConfEnableDaemonSetWorker = viper.GetBool(confEnableDaemonSetWorkerPath)
	ConfKubeletNamespace = viper.GetString(confKubeletNamespacePath)
	ConfKubeletName = viper.GetString(confKubeletNamePath)
	ConfKubeletEnable = viper.GetBool(confKubeletEnablePath)
	ConfDisableMetricsPusher = viper.GetBool(confDisableMetricsPusherPath)
	ConfMaxNodeSecretRatio = viper.GetFloat64(confMaxNodeSecretRatioPath)
	ConfTLSConfig = &rest.TLSClientConfig{
		Insecure: viper.GetBool(confTLSInsecurePath),
		CertFile: viper.GetString(confTLSCertFilePath),
		KeyFile:  viper.GetString(confTLSKeyFilePath),
		CAFile:   viper.GetString(confTLSCAFilePath),
	}
	ConfPodName = viper.GetString(define.EnvPodName)
	ConfStatefulSetWorkerHpa = viper.GetBool(confStatefulSetWorkerHpaPath)
	ConfStatefulSetWorkerFactor = viper.GetInt(confStatefulSetWorkerFactorPath)
	ConfStatefulSetReplicas = viper.GetInt(confStatefulSetReplicasPath)
	ConfStatefulSetMaxReplicas = viper.GetInt(confStatefulSetMaxReplicasPath)
	ConfStatefulSetDispatchType = viper.GetString(confStatefulSetDispatchType)

	// reload 时状态需要置空
	if viper.IsSet(confStatefulSetMatchRulesPath) {
		if err := viper.UnmarshalKey(confStatefulSetMatchRulesPath, &ConfStatefulSetMatchRules); err != nil {
			logger.Errorf("failed to unmarshal ConfStatefulSetMatchRules, err: %v", err)
		}
	} else {
		ConfStatefulSetMatchRules = []StatefulSetMatchRule{}
	}

	if viper.IsSet(confMonitorBlacklistMatchRules) {
		if err := viper.UnmarshalKey(confMonitorBlacklistMatchRules, &ConfMonitorBlacklistMatchRules); err != nil {
			logger.Errorf("failed to unmarshal ConfMonitorBlacklistMatchRules, err: %v", err)
		}
	} else {
		ConfMonitorBlacklistMatchRules = []MonitorBlacklistMatchRule{}
	}
}

func init() {
	if err := config.EventBus.Subscribe(config.EventConfigPreParse, initConfig); err != nil {
		logger.Errorf("failed to subscribe event %s, err: %v", config.EventConfigPreParse, err)
	}

	if err := config.EventBus.Subscribe(config.EventConfigPostParse, updateConfig); err != nil {
		logger.Errorf("failed to subscribe event %s, err: %v", config.EventConfigPostParse, err)
	}
}