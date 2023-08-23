package model

import "database/sql"

type Tenant struct {
	TenantID    int64  `json:"tenant_id" db:"tenant_id"`
	TenantName  string `json:"tenant_name" db:"tenant_name"`
	PrimaryZone string `json:"primary_zone" db:"primary_zone"`
	Locality    string `json:"locality" db:"locality"`
	Status      string `json:"status" db:"status"`
}

type Pool struct {
	ResourcePoolID int64  `json:"resource_pool_id" db:"resource_pool_id"`
	Name         string `json:"name" db:"name"`
	UnitNum      int64  `json:"unit_count" db:"unit_count"`
	UnitConfigID int64  `json:"unit_config_id" db:"unit_config_id"`
	ZoneList       string `json:"zone_list" db:"zone_list"`
	TenantID       int64  `json:"tenant_id" db:"tenant_id"`
}

type Unit struct {
	UnitID             int64  `json:"unit_id" db:"unit_id"`
	ResourcePoolID     int64  `json:"resource_pool_id" db:"resource_pool_id"`
	Zone               string `json:"zone" db:"zone"`
	SvrIP              string `json:"svr_ip" db:"svr_ip"`
	SvrPort            int64  `json:"svr_port" db:"svr_port"`
	MigrateFromSvrIP   sql.NullString `json:"migrate_from_svr_ip" db:"migrate_from_svr_ip"`
	MigrateFromSvrPort sql.NullInt64  `json:"migrate_from_svr_port" db:"migrate_from_svr_port"`
	Status             string `json:"status" db:"status"`
}

type UnitConfigV4 struct {
	UnitConfigID int64   `json:"unit_config_id" db:"unit_config_id"`
	Name         string  `json:"name" db:"name"`
	MaxCPU       float64 `json:"max_cpu" db:"max_cpu"`
	MinCPU       float64 `json:"min_cpu" db:"min_cpu"`
	MemorySize   int64   `json:"memory_size" db:"memory_size"`
	MaxIops      int64   `json:"max_iops" db:"max_iops"`
	MinIops      int64   `json:"min_iops" db:"min_iops"`
	LogDiskSize  int64   `json:"log_disk_size" db:"log_disk_size"`
	IopsWeight   int64   `json:"iops_weight" db:"iops_weight"`
	Options      string  `json:"options" db:"options"`
}

type ResourceTotal struct {
	CPUTotal  float64 `json:"cpu_total" db:"cpu_total"`
	MemTotal  int64   `json:"mem_total" db:"mem_total"`
	DiskTotal int64   `json:"disk_total" db:"disk_total"`
}

type Charset struct {
	Charset string `json:"charset" db:"charset"`
}

type Variable struct {
	VariableName string `json:"Variable_name" db:"Variable_name"`
	Value        string `json:"Value" db:"Value"`
}

type RsJob struct {
	JobID      int64  `json:"job_id" db:"job_id"`
	JobType    string `json:"job_type" db:"job_type"`
	JobStatus  string `json:"job_status" db:"job_status"`
	TenantID   int64  `json:"tenant_id" db:"tenant_id"`
	TenantName string `json:"tenant_name" db:"tenant_name"`
}

type OBVersion struct {
	Version string `json:"version" db:"version"`
}

type TenantSQLParam struct {
	TenantName string
	PrimaryZone string
	Locality string
	PoolList []string
	Charset string
	Collate string
	VariableList string
}

type PoolSQLParam struct {
	PoolName string
	UnitNum  int64
	ZoneList string
	UnitName string
}

type UnitConfigV4SQLParam struct {
	UnitConfigName string
	MaxCPU       float64
	MinCPU       float64
	MemorySize   int64
	MaxIops      int64
	MinIops      int64
	LogDiskSize  int64
	IopsWeight   int64
}