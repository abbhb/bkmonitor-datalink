// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package trace

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/offline-data-archive/eventbus"
)

// setDefaultConfig
func setDefaultConfig() {
	viper.SetDefault(KeysConfigPath, nil)
	viper.SetDefault(OtlpHostConfigPath, "127.0.0.1")
	viper.SetDefault(OtlpPortConfigPath, "4317")
	viper.SetDefault(OtlpTokenConfigPath, "")
	viper.SetDefault(OtlpTypeConfigPath, "grpc")

	viper.SetDefault(ServiceNameConfigPath, "offline-data-archive")
}

// InitConfig
func InitConfig() {

	for key, value := range configLabels {
		labels[key] = value
	}

	otlpHost = viper.GetString(OtlpHostConfigPath)
	otlpPort = viper.GetString(OtlpPortConfigPath)
	otlpToken = viper.GetString(OtlpTokenConfigPath)

	otlpType = viper.GetString(OtlpTypeConfigPath)

	ServiceName = viper.GetString(ServiceNameConfigPath)

}

// init
func init() {
	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPreParse, setDefaultConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for trace module for default config, maybe http module won't working.",
			eventbus.EventSignalConfigPreParse,
		)
	}

	if err := eventbus.EventBus.Subscribe(eventbus.EventSignalConfigPostParse, InitConfig); err != nil {
		fmt.Printf(
			"failed to subscribe event->[%s] for trace module for new config, maybe http module won't working.",
			eventbus.EventSignalConfigPostParse,
		)
	}
}
