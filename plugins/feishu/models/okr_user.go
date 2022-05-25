/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package models

import (
	"time"

	"github.com/apache/incubator-devlake/models/common"
)

type FeishuOkrUser struct {
	MemberID     string `json:"member_id" gorm:"type:varchar(255)"`
	Name         string `json:"name" gorm:"type:varchar(255)"`
	MemberIDType string `json:"member_id_type" gorm:"type:varchar(255)"`

	StartTime time.Time

	common.Model `json:"-"`
	common.RawDataOrigin
}

func (FeishuOkrUser) TableName() string {
	return "_tool_feishu_okr_user_okr"
}
