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

package core

import (
	cloudv1 "github.com/oceanbase/ob-operator/apis/cloud/v1"
	"github.com/pkg/errors"
)

func (ctrl *OBClusterCtrl) CreateUserForObproxy(statefulApp cloudv1.StatefulApp) error {
	sqlOperator, err := ctrl.GetSqlOperatorFromStatefulApp(statefulApp)
	if err != nil {
		return errors.Wrap(err, "get sql operator when create user for obproxy")
	}

	err = sqlOperator.CreateUser("proxyro", "")
	if err != nil {
		return err
	}

	err = sqlOperator.GrantPrivilege("select", "*.*", "proxyro")
	if err != nil {
		return err
	}
	return nil
}
