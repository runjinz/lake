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

package archived

import (
	"encoding/json"

	"github.com/apache/incubator-devlake/models/common"
)

type FeishuOkrUserObjective struct {
	ID                                 string            `json:"id" gorm:"primaryKey;type:varchar(255)"`
	OkrID                              string            `json:"okr_id" gorm:"type:varchar(255)"`
	MemberID                           string            `json:"member_id" gorm:"type:varchar(255)"`
	Content                            string            `json:"content" gorm:"type:varchar(255)"`
	Deadline                           string            `json:"deadline" gorm:"type:varchar(255)"`
	Permission                         int               `json:"permission" `
	MentionedUsers                     []json.RawMessage `json:"mentioned_user_list"`
	ProgressRecords                    []json.RawMessage `json:"progress_record_list"`
	AlignedObjectives                  []json.RawMessage `json:"aligned_objective_list"`
	AligningObjectives                 []json.RawMessage `json:"aligning_objective_list"`
	ProgressRateStatus                 string            `json:"progress_rate_status" gorm:"type:varchar(255)"`
	ProgressRatePercent                int               `json:"progress_rate_percent"`
	ProgressReport                     string            `json:"progress_report" gorm:"type:varchar(255)"`
	ProgressRateStatusLastUpdatedTime  string            `json:"progress_rate_status_last_updated_time" gorm:"type:varchar(255)"`
	ProgressRatePercentLastUpdatedTime string            `json:"progress_rate_percent_last_updated_time" gorm:"type:varchar(255)"`
	ProgressRecordLastUpdatedTime      string            `json:"progress_record_last_updated_time" gorm:"type:varchar(255)"`
	ProgressReportLastUpdatedTime      string            `json:"progress_report_last_updated_time" gorm:"type:varchar(255)"`
	ScoreLastUpdatedTime               string            `json:"score_last_updated_time" gorm:"type:varchar(255)"`
	Score                              int               `json:"score"`
	Weight                             int               `json:"weight"`

	common.NoPKModel `json:"-"`
}

func (FeishuOkrUserObjective) TableName() string {
	return "_tool_feishu_okr_user_objectives"
}
