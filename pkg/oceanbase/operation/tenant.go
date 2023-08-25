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

package operation

import (
	"fmt"
	"github.com/oceanbase/ob-operator/pkg/oceanbase/const/config"
	"github.com/oceanbase/ob-operator/pkg/oceanbase/const/sql"
	"github.com/oceanbase/ob-operator/pkg/oceanbase/model"
	"github.com/pkg/errors"
	"strings"
)

func (m *OceanbaseOperationManager) GetTenantByName(tenantName string) (*model.Tenant, error) {
	tenant := &model.Tenant{}
	err := m.QueryRow(tenant, sql.GetTenantByName, tenantName)
	if err != nil {
		return tenant, errors.Wrap(err, "Get tenantconst by tenantName")
	}
	return tenant, nil
}

func (m *OceanbaseOperationManager) GetPoolByName(poolName string) (*model.Pool, error) {
	pool := &model.Pool{}
	err := m.QueryRow(pool, sql.GetPoolByName, poolName)
	if err != nil {
		return pool, errors.Wrap(err, "Get pool by poolName")
	}
	return pool, nil
}

func (m *OceanbaseOperationManager) GetPoolList() ([]model.Pool, error) {
	var poolList []model.Pool
	err := m.QueryList(&poolList, sql.GetPoolList)
	if err != nil {
		return poolList, errors.Wrap(err, "Get get pool list")
	}
	return poolList, nil
}

func (m *OceanbaseOperationManager) GetResourceTotal(zoneName string) (*model.ResourceTotal, error) {
	resource := &model.ResourceTotal{}
	err := m.QueryRow(resource, sql.GetResourceTotal, zoneName)
	if err != nil {
		return resource, errors.Wrap(err, "Get resource by zoneName")
	}
	return resource, nil
}

func (m *OceanbaseOperationManager) GetUnitList() ([]model.Unit, error) {
	var unitList []model.Unit
	err := m.QueryList(&unitList, sql.GetUnitList)
	if err != nil {
		return unitList, errors.Wrap(err, "Get all unit list")
	}
	return unitList, nil
}

func (m *OceanbaseOperationManager) GetUnitConfigV4List() ([]model.UnitConfigV4, error) {
	var unitConfigV4List []model.UnitConfigV4
	err := m.QueryList(&unitConfigV4List, sql.GetUnitConfigV4List)
	if err != nil {
		return unitConfigV4List, errors.Wrap(err, "Get all unitConfigV4 list")
	}
	return unitConfigV4List, nil
}

func (m *OceanbaseOperationManager) GetUnitConfigV4ByName(name string) (*model.UnitConfigV4, error) {
	pool := &model.UnitConfigV4{}
	err := m.QueryRow(pool, sql.GetUnitConfigV4ByName, name)
	if err != nil {
		return pool, errors.Wrap(err, "Get unitConfigV4 By Name")
	}
	return pool, nil
}

func (m *OceanbaseOperationManager) GetCharset() (*model.Charset, error) {
	charset := &model.Charset{}
	err := m.QueryRow(charset, sql.GetCharset)
	if err != nil {
		return charset, errors.Wrap(err, "Get charset")
	}
	return charset, nil
}

func (m *OceanbaseOperationManager) GetVariable(name string) (*model.Variable, error) {
	variable := &model.Variable{}
	err := m.QueryRow(variable, sql.GetVariableLike, name)
	if err != nil {
		return variable, errors.Wrap(err, "Get variable")
	}
	return variable, nil
}

func (m *OceanbaseOperationManager) GetRsJob(reJobName string) (*model.RsJob, error) {
	rsJob := &model.RsJob{}
	err := m.QueryRow(rsJob, sql.GetRsJob, reJobName)
	if err != nil {
		return rsJob, errors.Wrap(err, "Get rsJob by reJobName")
	}
	return rsJob, nil
}

func (m *OceanbaseOperationManager) GetVersion() (*model.OBVersion, error) {
	version := &model.OBVersion{}
	err := m.QueryRow(version, sql.GetObVersion)
	if err != nil {
		return version, errors.Wrap(err, "Get version")
	}
	return version, nil
}

// ------------ delete ------------

func (m *OceanbaseOperationManager) DeleteTenant(tenantName string) error {
	preparedSQL, params := m.preparedSQLForDeleteTenant(tenantName)
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Delete tenantconst by tenantName")
	}
	return nil
}

func (m *OceanbaseOperationManager) DeletePool(poolName string) error {
	preparedSQL, params := m.preparedSQLForDeletePool(poolName)
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Delete pool by poolName")
	}
	return nil
}

func (m *OceanbaseOperationManager) DeleteUnitConfig(unitName string) error {
	preparedSQL, params := m.preparedSQLForDeleteUnitConfig(unitName)
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Delete unit by unitName")
	}
	return nil
}

// ------------ check function ------------

func (m *OceanbaseOperationManager) CheckTenantExistByName(tenantName string) (bool, error) {
	var count int
	err := m.QueryCount(&count, sql.GetTenantCountByName, tenantName)
	if err != nil {
		return false, errors.Wrap(err, "Get tenantconst by tenantName")
	}
	return count != 0, nil
}

func (m *OceanbaseOperationManager) CheckPoolExistByName(poolName string) (bool, error) {
	var count int
	err := m.QueryCount(&count, sql.GetPoolCountByName, poolName)
	if err != nil {
		return false, errors.Wrap(err, "Check whether pool exist by poolName")
	}
	return count != 0, nil
}

func (m *OceanbaseOperationManager) CheckUnitConfigExistByName(unitConfigName string) (bool, error) {
	var count int
	err := m.QueryCount(&count, sql.GetUnitConfigV4CountByName, unitConfigName)
	if err != nil {
		return false, errors.Wrap(err, "Check whether unitconfigV4 exist by poolName")
	}
	return count != 0, nil
}

func (m *OceanbaseOperationManager) CheckRsJobExistByTenantID(tenantName int) (bool, error) {
	var count int
	err := m.QueryCount(&count, sql.GetRsJobCount, tenantName)
	if err != nil {
		return false, errors.Wrap(err, "Check whether rsJob exist by poolName")
	}
	return count != 0, nil
}

// ------------ add ------------

func (m *OceanbaseOperationManager) AddTenant(tenantSQLParam model.TenantSQLParam) error {

	preparedSQL, params := preparedSQLForAddTenant(tenantSQLParam)
	err := m.ExecWithTimeout(config.TenantSqlTimeout, preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Add Tenant")
	}
	return nil
}

func (m *OceanbaseOperationManager) AddPool(pool model.PoolSQLParam) error {
	preparedSQL, params := preparedSQLForAddPool(pool)
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Add pool")
	}
	return nil
}

func (m *OceanbaseOperationManager) AddUnitConfigV4(unitConfigV4 model.UnitConfigV4SQLParam) error {
	preparedSQL, params := preparedSQLForAddUnitConfigV4(unitConfigV4)
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Add UnitConfigV4")
	}
	return nil
}

// ------------ modify ------------

func (m *OceanbaseOperationManager) SetTenantVariable(tenantName, variableList string) error {
	preparedSQL, params := m.preparedSQLForSetTenantVariable(tenantName, variableList)
	m.Logger.Info(fmt.Sprintf("sql: %s, parms: %v", preparedSQL, params))
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Set Tenant Variable")
	}
	return nil
}

func (m *OceanbaseOperationManager) SetUnitConfigV4(unitConfigV4 model.UnitConfigV4SQLParam) error {
	preparedSQL, params := preparedSQLForSetUnitConfigV4(unitConfigV4)
	m.Logger.Info(fmt.Sprintf("sql: %s, parms: %v", preparedSQL, params))
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Set UnitConfig")
	}
	return nil
}

func (m *OceanbaseOperationManager) SetTenantUnitNum(tenantName string,unitNum int) error {
	preparedSQL, params := m.preparedSQLForSetTenantUnitNum(tenantName, unitNum)
	m.Logger.Info(fmt.Sprintf("sql: %s, parms: %v", preparedSQL, params))
	err := m.ExecWithDefaultTimeout(preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Set pool UnitNum")
	}
	return nil
}

func (m *OceanbaseOperationManager) SetTenant(tenantSQLParam model.TenantSQLParam) error {
	preparedSQL, params := preparedSQLForSetTenant(tenantSQLParam)
	m.Logger.Info(fmt.Sprintf("sql: %s, parms: %v", preparedSQL, params))
	err := m.ExecWithTimeout(config.TenantSqlTimeout, preparedSQL, params...)
	if err != nil {
		return errors.Wrap(err, "Set tenant")
	}
	return nil
}

// ---------- replacer sql and collect params ----------

func preparedSQLForAddUnitConfigV4(unitConfigV4 model.UnitConfigV4SQLParam) (string,[]interface{}) {
	var optionSql string
	params := make([]interface{}, 0)
	params = append(params, unitConfigV4.MaxCPU, unitConfigV4.MemorySize)
	if unitConfigV4.MinCPU != 0 {
		optionSql = fmt.Sprint(optionSql, ", min_cpu ?")
		params = append(params, unitConfigV4.MinCPU)
	}
	if unitConfigV4.LogDiskSize != 0 {
		optionSql = fmt.Sprint(optionSql, ", log_disk_size ?")
		params = append(params, unitConfigV4.LogDiskSize)
	}
	if unitConfigV4.MaxIops != 0 {
		optionSql = fmt.Sprint(optionSql, ", max_iops ?")
		params = append(params, unitConfigV4.MaxIops)
	}
	if unitConfigV4.MinIops != 0 {
		optionSql = fmt.Sprint(optionSql, ", min_iops ?")
		params = append(params, unitConfigV4.MinIops)
	}
	if unitConfigV4.IopsWeight != 0 {
		optionSql = fmt.Sprint(optionSql, ", iops_weight ?")
		params = append(params, unitConfigV4.IopsWeight)
	}
	replacer := strings.NewReplacer("${UNIT_NAME}", unitConfigV4.UnitConfigName, "${OPTION}", optionSql)
	return replaceAll(sql.AddUnitConfigV4, replacer), params
}



func preparedSQLForAddPool(poolSQLParam model.PoolSQLParam) (string, []interface{}) {
	params := make([]interface{}, 0)
	params = append(params, poolSQLParam.UnitName, poolSQLParam.UnitNum, poolSQLParam.ZoneList)
	replacer := strings.NewReplacer("${POOL_NAME}", poolSQLParam.PoolName)
 	return replaceAll(sql.AddPool, replacer) , params
}

func preparedSQLForAddTenant(tenantSQLParam model.TenantSQLParam) (string, []interface{}) {
	var option string
	var variableList string
	params := make([]interface{}, 0)
	params = append(params, tenantSQLParam.Charset, tenantSQLParam.PrimaryZone)

	symbols := make([]string, 0)
	for i := 0; i < len(tenantSQLParam.PoolList); i++ {
		symbols = append(symbols, "?")
		params = append(params, tenantSQLParam.PoolList[i])
	}
	if tenantSQLParam.Locality!= "" {
		option = fmt.Sprint(option, ", LOCALITY= ?")
		params = append(params, tenantSQLParam.Locality)
	}
	if tenantSQLParam.Collate != "" {
		option = fmt.Sprint(option, ", COLLATE = ?")
		params = append(params, tenantSQLParam.Collate)
	}
	variableList = fmt.Sprintf("SET VARIABLES %s", tenantSQLParam.VariableList)
	replacer := strings.NewReplacer("${TENANT_NAME}", tenantSQLParam.TenantName, "${POOL_LIST}", strings.Join(symbols, ", "),
		"${OPTION}", option, "${VARIABLE_LIST}", variableList)
	return replaceAll(sql.AddTenant, replacer), params
}

func preparedSQLForSetTenant(tenantSQLParam model.TenantSQLParam) (string, []interface{}) {
	var alterItemStr string
	params := make([]interface{}, 0)
	alterItemList := make([]string, 0)
	if tenantSQLParam.PrimaryZone != "" {
		alterItemList = append(alterItemList, "PRIMARY_ZONE=?")
		params = append(params, tenantSQLParam.PrimaryZone)
	}
	if tenantSQLParam.Charset!= "" {
		alterItemList = append(alterItemList, "CHARSET=?")
		params = append(params, tenantSQLParam.Charset)
	}
	if len(tenantSQLParam.PoolList) != 0 {
		symbols := make([]string, 0)
		for i := 0; i < len(tenantSQLParam.PoolList); i++ {
			symbols = append(symbols, "?")
			params = append(params, tenantSQLParam.PoolList[i])
		}
		alterItemList = append(alterItemList, fmt.Sprintf("RESOURCE_POOL_LIST=(%s)", strings.Join(symbols, ", ")))
	}
	if tenantSQLParam.Locality!= "" {
		alterItemList = append(alterItemList, "LOCALITY=?")
		params = append(params, tenantSQLParam.Locality)
	}
	alterItemStr = strings.Join(alterItemList, ",")
	replacer := strings.NewReplacer("${TENANT_NAME}", tenantSQLParam.TenantName, "${ALTER_ITEM}", alterItemStr)
	return replaceAll(sql.SetTenant, replacer), params
}

func preparedSQLForSetUnitConfigV4(unitConfigV4 model.UnitConfigV4SQLParam) (string,[]interface{}) {
	var alterItemStr string
	params := make([]interface{}, 0)
	alterItemList := make([]string, 0)
	if unitConfigV4.MaxCPU!= 0 {
		alterItemList = append(alterItemList, "max_cpu=?")
		params = append(params, unitConfigV4.MaxCPU)
	}
	if unitConfigV4.MemorySize!= 0 {
		alterItemList = append(alterItemList, "memory_size=?")
		params = append(params, unitConfigV4.MemorySize)
	}
	if unitConfigV4.MinCPU != 0 {
		alterItemList = append(alterItemList, "min_cpu=?")
		params = append(params, unitConfigV4.MinCPU)
	}
	if unitConfigV4.MinCPU != 0 {
		alterItemList = append(alterItemList, "min_cpu=?")
		params = append(params, unitConfigV4.MinCPU)
	}
	if unitConfigV4.LogDiskSize != 0 {
		alterItemList = append(alterItemList, "log_disk_size=?")
		params = append(params, unitConfigV4.LogDiskSize)
	}
	if unitConfigV4.MaxIops != 0 {
		alterItemList = append(alterItemList, "max_iops=?")
		params = append(params, unitConfigV4.MaxIops)
	}
	if unitConfigV4.MinIops != 0 {
		alterItemList = append(alterItemList, "min_iops=?")
		params = append(params, unitConfigV4.MinIops)
	}
	if unitConfigV4.IopsWeight != 0 {
		alterItemList = append(alterItemList, "iops_weight=?")
		params = append(params, unitConfigV4.IopsWeight)
	}

	alterItemStr = strings.Join(alterItemList, ",")
	replacer := strings.NewReplacer("${UNIT_NAME}", unitConfigV4.UnitConfigName, "${ALTER_ITEM}", alterItemStr)
	return replaceAll(sql.SetUnitConfigV4, replacer), params
}

func (m *OceanbaseOperationManager) preparedSQLForSetTenantVariable(tenantName, variableList string) (string,[]interface{}) {
	params := make([]interface{}, 0)
	replacer := strings.NewReplacer("${TENANT_NAME}", tenantName, "${VARIABLE_LIST}", variableList)
	return replaceAll(sql.SetTenantVariable, replacer), params
}

func (m *OceanbaseOperationManager) preparedSQLForSetTenantUnitNum(tenantNum string, unitNum int) (string,[]interface{}) {
	params := make([]interface{}, 0)
	params = append(params, unitNum)
	replacer := strings.NewReplacer("${TENANT_NAME}", tenantNum)
	return replaceAll(sql.SetTenantUnitNum, replacer), params

}

func (m *OceanbaseOperationManager) preparedSQLForDeleteTenant(tenantName string) (string,[]interface{}) {
	params := make([]interface{}, 0)
	replacer := strings.NewReplacer("${TENANT_NAME}", tenantName)
	return replaceAll(sql.DeleteTenant, replacer), params
}

func (m *OceanbaseOperationManager) preparedSQLForDeletePool(poolName string) (string,[]interface{}) {
	params := make([]interface{}, 0)
	replacer := strings.NewReplacer("${POOL_NAME}", poolName)
	return replaceAll(sql.DeletePool, replacer), params
}

func (m *OceanbaseOperationManager) preparedSQLForDeleteUnitConfig(unitConfigName string) (string,[]interface{}) {
	params := make([]interface{}, 0)
	replacer := strings.NewReplacer("${UNIT_NAME}", unitConfigName)
	return replaceAll(sql.DeleteUnitConfig, replacer), params
}