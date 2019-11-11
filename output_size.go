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

import (
	"math"
	"time"
)

// An OutputSize is the output size of time series data.
type OutputSize string

// Output sizes of time series data.
const (
	OutputSizeCompact OutputSize = "compact"
	OutputSizeFull    OutputSize = "full"
)

// OutputSizeSince returns the smallest output size that covers a given date.
func OutputSizeSince(date time.Time) OutputSize {
	date = date.Truncate(24 * time.Hour)
	now := time.Now().Truncate(24 * time.Hour)
	if math.Ceil(now.Sub(date).Hours()/24) <= 100 {
		return OutputSizeCompact
	}
	return OutputSizeFull
}
