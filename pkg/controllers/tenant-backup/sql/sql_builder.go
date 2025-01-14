/*
Copyright (c) 2021 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package sql

import "strings"

func ReplaceAll(template string, replacers ...*strings.Replacer) string {
	s := template
	for _, replacer := range replacers {
		s = replacer.Replace(s)
	}
	return s
}

func SetParameterSQLReplacer(name, value string) *strings.Replacer {
	return strings.NewReplacer("${NAME}", name, "${VALUE}", value)
}

func SetBackupPasswordReplacer(pwd string) *strings.Replacer {
	return strings.NewReplacer("${pwd}", pwd)
}

func GetParameterSQLReplacer(name string) *strings.Replacer {
	return strings.NewReplacer("${NAME}", name)
}

func GetLeaderCountSQLReplacer(zoneName string) *strings.Replacer {
	return strings.NewReplacer("${ZONE_NAME}", zoneName)
}

func SetDeletePolicySQLReplacer(policyName, recoveryWindow string) *strings.Replacer {
	return strings.NewReplacer("${POLICY_NAME}", policyName, "${RECOVERY_WINDOW}", recoveryWindow)
}
