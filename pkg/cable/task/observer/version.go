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

package observer

import (
	"context"
	"strings"

	"github.com/oceanbase/ob-operator/pkg/config/constant"
	"github.com/oceanbase/ob-operator/pkg/util/shell"
	log "github.com/sirupsen/logrus"
)

func GetObVersion() (*shell.ExecuteResult, error) {
	res, err := shell.NewCommand(constant.OBSERVER_VERSION_COMMAND).WithContext(context.TODO()).WithUser(shell.AdminUser).Execute()
	if err != nil {
		log.WithError(err).Errorf("start observer command exec error %v", err)
	}
	lines := strings.Split(res.Output, "\n")
	log.Info("lines: ", len(lines), lines)
	if len(lines) > 2 {
		output := strings.Split(lines[1], " ")
		ver := output[len(output)-1]
		res.Output = ver[0 : len(ver)-1]
	}
	return res, err
}
