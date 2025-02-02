// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package limiter

import "time"

const (
	defaultMaxQueueSize = 1000
	defaultMaxWaitTime  = 15 * time.Second
)

type LimiterConfig struct {
	MaxQueueSize *int
	MaxWaitTime  *time.Duration
}

func setDefaultsForLimiterConfig(config LimiterConfig) LimiterConfig {
	maxQueueSize, maxWaitTime := defaultMaxQueueSize, defaultMaxWaitTime
	if config.MaxQueueSize == nil {
		config.MaxQueueSize = &maxQueueSize
	}

	if config.MaxWaitTime == nil {
		config.MaxWaitTime = &maxWaitTime
	}

	return config
}
