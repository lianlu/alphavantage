// Copyright 2019 Miles Barr <milesbarr2@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package alphavantage

import "time"

// An Interval is the interval between time series data points.
type Interval string

// Intervals between time series data points.
const (
	Interval1Min   Interval = "1min"
	Interval5Min   Interval = "5min"
	Interval15Min  Interval = "15min"
	Interval30Min  Interval = "30min"
	Interval60Min  Interval = "60min"
	Interval1Day   Interval = "DAILY"
	Interval1Week  Interval = "WEEKLY"
	Interval1Month Interval = "MONTHLY"
)

// Duration returns the duration for the Interval.
func (i Interval) Duration() time.Duration {
	switch i {
	case Interval1Min:
		return 1 * time.Minute
	case Interval5Min:
		return 5 * time.Minute
	case Interval15Min:
		return 15 * time.Minute
	case Interval30Min:
		return 30 * time.Minute
	case Interval60Min:
		return 60 * time.Minute
	case Interval1Day:
		return 1 * time.Hour * 24
	case Interval1Week:
		return 7 * time.Hour * 24
	case Interval1Month:
		return 30 * time.Hour * 24
	}
	return 0
}
