// Licensed to ClickHouse, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. ClickHouse, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package column

import (
	"reflect"
	"strings"
	"time"

	"github.com/ClickHouse/ch-go/proto"
)

type AggregateFunction struct {
	base   Interface
	chType Type
	name   string
}

func (col *AggregateFunction) Reset() {
	col.base.Reset()
}

func (col *AggregateFunction) Name() string {
	return col.name
}

func (col *AggregateFunction) parse(t Type, tz *time.Location) (_ Interface, err error) {
	col.chType = t

	// throw away arguments
	base := strings.TrimSpace(strings.SplitN(t.params(), ",", 3)[1])
	if col.base, err = Type(base).Column(col.name, tz); err == nil {
		return col, nil
	}
	return nil, &UnsupportedColumnTypeError{
		t: t,
	}
}

func (col *AggregateFunction) Type() Type {
	return col.chType
}

func (col *AggregateFunction) ScanType() reflect.Type {
	return col.base.ScanType()
}

func (col *AggregateFunction) Rows() int {
	return col.base.Rows()
}

func (col *AggregateFunction) Row(i int, ptr bool) any {
	return col.base.Row(i, ptr)
}

func (col *AggregateFunction) ScanRow(dest any, rows int) error {
	return col.base.ScanRow(dest, rows)
}

func (col *AggregateFunction) Append(v any) ([]uint8, error) {
	return col.base.Append(v)
}

func (col *AggregateFunction) AppendRow(v any) error {
	return col.base.AppendRow(v)
}

func (col *AggregateFunction) Decode(reader *proto.Reader, rows int) error {
	return col.base.Decode(reader, rows)
}

func (col *AggregateFunction) Encode(buffer *proto.Buffer) {
	col.base.Encode(buffer)
}

var _ Interface = (*AggregateFunction)(nil)
