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
	"net/http"
	"net/url"
	"strconv"

	"github.com/apache/incubator-devlake/config"
	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/feishu/apimodels"
	"github.com/apache/incubator-devlake/plugins/helper"
)

const RAW_OKR_USER_OKR_TABLE = "feishu_okr_user_okr"

var _ core.SubTaskEntryPoint = CollectChatMember

func CollectChatMember(taskCtx core.SubTaskContext) error {
	data := taskCtx.GetData().(*FeishuTaskData)
	pageSize := 100
	// NumOfDaysToCollectInt := int(data.Options.NumOfDaysToCollect)
	// iterator, err := helper.NewDateIterator(NumOfDaysToCollectInt)
	// if err != nil {
	// 	return err
	// }
	incremental := false

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: FeishuApiParams{
				ApiResName: "chat_member",
			},
			Table: RAW_OKR_USER_OKR_TABLE,
		},
		ApiClient:   data.ApiClient,
		Incremental: incremental,
		// Input:       iterator,
		// UrlTemplate: "/okr/v1/users/:user_id/okrs",
		UrlTemplate: "im/v1/chats/" + config.GetConfig().GetString("FEISHU_CHATID") + "/members",
		Query: func(reqData *helper.RequestData) (url.Values, error) {
			query := url.Values{}
			// input := reqData.Input.(*helper.DatePair)
			// query.Set("start_time", strconv.FormatInt(input.PairStartTime.Unix(), 10))
			// query.Set("end_time", strconv.FormatInt(input.PairEndTime.Unix(), 10))
			query.Set("page_size", strconv.Itoa(pageSize))
			// query.Set("order_by", "2")
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, error) {
			body := &apimodels.FeishuChatMemberResult{}
			err := helper.UnmarshalResponse(res, body)
			if err != nil {
				return nil, err
			}
			return body.Data.Items, nil
		},
	})
	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectChatMemberMeta = core.SubTaskMeta{
	Name:             "collectChatMember",
	EntryPoint:       CollectChatMember,
	EnabledByDefault: true,
	Description:      "Collect chat member data from Feishu api",
}
