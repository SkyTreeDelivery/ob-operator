/*
Copyright (c) 2023 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package v1alpha1

type OperationContext struct {
	Name         string   `json:"name"`
	Tasks        []string `json:"tasks"`
	Task         string   `json:"task"`
	Idx          int      `json:"idx"`
	TaskStatus   string   `json:"taskStatus"`
	TaskId       string   `json:"taskId"`
	TargetStatus string   `json:"targetStatus"`
}
