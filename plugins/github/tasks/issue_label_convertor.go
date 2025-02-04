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
	"reflect"

	"github.com/apache/incubator-devlake/models/domainlayer/didgen"
	"github.com/apache/incubator-devlake/models/domainlayer/ticket"
	"github.com/apache/incubator-devlake/plugins/core"
	githubModels "github.com/apache/incubator-devlake/plugins/github/models"
	"github.com/apache/incubator-devlake/plugins/helper"
)

var ConvertIssueLabelsMeta = core.SubTaskMeta{
	Name:             "convertIssueLabels",
	EntryPoint:       ConvertIssueLabels,
	EnabledByDefault: true,
	Description:      "Convert tool layer table github_issue_labels into  domain layer table issue_labels",
}

func ConvertIssueLabels(taskCtx core.SubTaskContext) error {
	db := taskCtx.GetDb()
	data := taskCtx.GetData().(*GithubTaskData)
	repoId := data.Repo.GithubId

	cursor, err := db.Model(&githubModels.GithubIssueLabel{}).
		Joins(`left join _tool_github_issues on _tool_github_issues.github_id = _tool_github_issue_labels.issue_id`).
		Where("_tool_github_issues.repo_id = ?", repoId).
		Order("issue_id ASC").
		Rows()
	if err != nil {
		return err
	}
	defer cursor.Close()
	issueIdGen := didgen.NewDomainIdGenerator(&githubModels.GithubIssue{})

	converter, err := helper.NewDataConverter(helper.DataConverterArgs{
		RawDataSubTaskArgs: helper.RawDataSubTaskArgs{
			Ctx: taskCtx,
			Params: GithubApiParams{
				Owner: data.Options.Owner,
				Repo:  data.Options.Repo,
			},
			Table: RAW_ISSUE_TABLE,
		},
		InputRowType: reflect.TypeOf(githubModels.GithubIssueLabel{}),
		Input:        cursor,
		Convert: func(inputRow interface{}) ([]interface{}, error) {
			issueLabel := inputRow.(*githubModels.GithubIssueLabel)
			domainIssueLabel := &ticket.IssueLabel{
				IssueId:   issueIdGen.Generate(issueLabel.IssueId),
				LabelName: issueLabel.LabelName,
			}
			return []interface{}{
				domainIssueLabel,
			}, nil
		},
	})
	if err != nil {
		return err
	}

	return converter.Execute()
}
