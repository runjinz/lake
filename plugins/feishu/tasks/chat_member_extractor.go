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

var _ core.SubTaskEntryPoint = ExtractChatMember

func ExtractChatMember(taskCtx core.SubTaskContext) error {

	exetractor, err := helper.NewApiExtractor(helper.ApiExtractorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: FeishuApiParams{
				ApiResName: "chat_member",
			},
			Table: RAW_CHAT_MEMBERS_TABLE,
		},
		Extract: func(row *helper.RawData) ([]interface{}, error) {
			body := &models.FeishuChatMember{}
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
			results = append(results, &models.FeishuChatMember{
				MemberID:     body.MemberID,
				Name:         body.Name,
				MemberIDType: body.MemberIDType,
			})
			return results, nil
		},
	})
	if err != nil {
		return err
	}

	return exetractor.Execute()
}

var ExtractChatMemberMeta = core.SubTaskMeta{
	Name:             "extractChatMember",
	EntryPoint:       ExtractChatMember,
	EnabledByDefault: true,
	Description:      "Extrat raw chat member data into tool layer table",
}
