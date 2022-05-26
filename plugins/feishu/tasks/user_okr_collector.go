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
	"reflect"
	"strconv"

	"github.com/apache/incubator-devlake/plugins/core"
	"github.com/apache/incubator-devlake/plugins/feishu/apimodels"
	"github.com/apache/incubator-devlake/plugins/feishu/models"
	"github.com/apache/incubator-devlake/plugins/helper"
)

const RAW_OKR_USER_OKRS_TABLE = "feishu_okr_user_okrs"

var _ core.SubTaskEntryPoint = CollectUserOkrs

func CollectUserOkrs(taskCtx core.SubTaskContext) error {

	data := taskCtx.GetData().(*FeishuTaskData)
	db := taskCtx.GetDb()

	// filter out issue_ids that needed collection
	tx := db.Table("_tool_feishu_chat_members t").
		Select("t.member_id,t.name,t.member_id_type")

	// construct the input iterator
	cursor, err := tx.Rows()
	if err != nil {
		return err
	}
	// smaller struct can reduce memory footprint, we should try to avoid using big struct
	iterator, err := helper.NewCursorIterator(db, cursor, reflect.TypeOf(models.FeishuChatMember{}))
	if err != nil {
		return err
	}

	incremental := false

	collector, err := helper.NewApiCollector(helper.ApiCollectorArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: FeishuApiParams{
				ApiResName: "user_okrs",
			},
			Table: RAW_OKR_USER_OKRS_TABLE,
		},
		ApiClient:   data.ApiClient,
		Incremental: incremental,
		Input:       iterator,
		UrlTemplate: "/okr/v1/users/{{ .Input.MemberID }}/okrs",
		Query: func(reqData *helper.RequestData) (url.Values, error) {
			query := url.Values{}
			query.Set("offset", strconv.Itoa(0))
			query.Set("limit", strconv.Itoa(10))
			return query, nil
		},
		ResponseParser: func(res *http.Response) ([]json.RawMessage, error) {
			body := &apimodels.FeishuUserOkrResults{}
			err := helper.UnmarshalResponse(res, body)
			if err != nil {
				return nil, err
			}
			return body.Data.Okrs, nil
		},
	})
	if err != nil {
		return err
	}

	return collector.Execute()
}

var CollectUserOkrsMeta = core.SubTaskMeta{
	Name:             "collectUserOkrs",
	EntryPoint:       CollectUserOkrs,
	EnabledByDefault: true,
	Description:      "Collect user okrs data from Feishu api",
}
