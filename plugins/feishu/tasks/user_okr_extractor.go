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

package tasks

import (
	"encoding/json"

	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/feishu/models"
	"github.com/apache/incubator-devlake/plugins/helper"
)

var _ core.SubTaskEntryPoint = ExtractUserOkr

func ExtractUserOkr(taskCtx core.SubTaskContext) error {

	exetractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: FeishuApiParams{
				ApiResName: "user_okr",
			},
			Table: RAW_OKR_USER_OKRS_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, error) {
			body := &models.FeishuOkrUserOkr{}
			err := json.Unmarshal(row.Data, body)
			if err != nil {
				return nil, err
			}
			rawInput := &helper.DatePair{}
			rawErr := json.Unmarshal(row.Input, rawInput)
			if rawErr != nil {
				return nil, rawErr
			}
			results := make([]interface{}, 0)
			results = append(results, &models.FeishuOkrUserOkr{
				StartTime:     rawInput.PairStartTime.AddDate(0, 0, -1),
				MemberID:      body.MemberID,
				Name:          body.Name,
				ID:            body.ID,
				ConfirmStatus: body.ConfirmStatus,
				PeriodID:      body.PeriodID,
				Permission:    body.Permission,
			})
			return results, nil
		},
	})
	if err != nil {
		return err
	}

	return exetractor.Execute()
}

var ExtractUserOkrMeta = core.SubTaskMeta{
	Name:             "extractUserOkr",
	EntryPoint:       ExtractUserOkr,
	EnabledByDefault: true,
	Description:      "Extrat raw user okrs data into tool layer table",
}
