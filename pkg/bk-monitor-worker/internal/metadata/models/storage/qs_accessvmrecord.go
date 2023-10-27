// Code generated by go-queryset. DO NOT EDIT.
package storage

import (
	"errors"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

// ===== BEGIN of all query sets

// ===== BEGIN of query set AccessVMRecordQuerySet

// AccessVMRecordQuerySet is an queryset type for AccessVMRecord
type AccessVMRecordQuerySet struct {
	db *gorm.DB
}

// NewAccessVMRecordQuerySet constructs new AccessVMRecordQuerySet
func NewAccessVMRecordQuerySet(db *gorm.DB) AccessVMRecordQuerySet {
	return AccessVMRecordQuerySet{
		db: db.Model(&AccessVMRecord{}),
	}
}

func (qs AccessVMRecordQuerySet) w(db *gorm.DB) AccessVMRecordQuerySet {
	return NewAccessVMRecordQuerySet(db)
}

func (qs AccessVMRecordQuerySet) Select(fields ...AccessVMRecordDBSchemaField) AccessVMRecordQuerySet {
	names := []string{}
	for _, f := range fields {
		names = append(names, f.String())
	}

	return qs.w(qs.db.Select(strings.Join(names, ",")))
}

// Create is an autogenerated method
// nolint: dupl
func (o *AccessVMRecord) Create(db *gorm.DB) error {
	return db.Create(o).Error
}

// Delete is an autogenerated method
// nolint: dupl
func (o *AccessVMRecord) Delete(db *gorm.DB) error {
	return db.Delete(o).Error
}

// All is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) All(ret *[]AccessVMRecord) error {
	return qs.db.Find(ret).Error
}

// BcsClusterIdEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdEq(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id = ?", bcsClusterId))
}

// BcsClusterIdGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdGt(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id > ?", bcsClusterId))
}

// BcsClusterIdGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdGte(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id >= ?", bcsClusterId))
}

// BcsClusterIdIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdIn(bcsClusterId ...string) AccessVMRecordQuerySet {
	if len(bcsClusterId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bcsClusterId in BcsClusterIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bcs_cluster_id IN (?)", bcsClusterId))
}

// BcsClusterIdLike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdLike(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id LIKE ?", bcsClusterId))
}

// BcsClusterIdLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdLt(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id < ?", bcsClusterId))
}

// BcsClusterIdLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdLte(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id <= ?", bcsClusterId))
}

// BcsClusterIdNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdNe(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id != ?", bcsClusterId))
}

// BcsClusterIdNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdNotIn(bcsClusterId ...string) AccessVMRecordQuerySet {
	if len(bcsClusterId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bcsClusterId in BcsClusterIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bcs_cluster_id NOT IN (?)", bcsClusterId))
}

// BcsClusterIdNotlike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BcsClusterIdNotlike(bcsClusterId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bcs_cluster_id NOT LIKE ?", bcsClusterId))
}

// BkBaseDataIdEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdEq(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id = ?", bkBaseDataId))
}

// BkBaseDataIdGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdGt(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id > ?", bkBaseDataId))
}

// BkBaseDataIdGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdGte(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id >= ?", bkBaseDataId))
}

// BkBaseDataIdIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdIn(bkBaseDataId ...uint) AccessVMRecordQuerySet {
	if len(bkBaseDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkBaseDataId in BkBaseDataIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_base_data_id IN (?)", bkBaseDataId))
}

// BkBaseDataIdLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdLt(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id < ?", bkBaseDataId))
}

// BkBaseDataIdLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdLte(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id <= ?", bkBaseDataId))
}

// BkBaseDataIdNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdNe(bkBaseDataId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("bk_base_data_id != ?", bkBaseDataId))
}

// BkBaseDataIdNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) BkBaseDataIdNotIn(bkBaseDataId ...uint) AccessVMRecordQuerySet {
	if len(bkBaseDataId) == 0 {
		qs.db.AddError(errors.New("must at least pass one bkBaseDataId in BkBaseDataIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("bk_base_data_id NOT IN (?)", bkBaseDataId))
}

// Count is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) Count() (int, error) {
	var count int
	err := qs.db.Count(&count).Error
	return count, err
}

// DataTypeEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeEq(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type = ?", dataType))
}

// DataTypeGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeGt(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type > ?", dataType))
}

// DataTypeGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeGte(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type >= ?", dataType))
}

// DataTypeIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeIn(dataType ...string) AccessVMRecordQuerySet {
	if len(dataType) == 0 {
		qs.db.AddError(errors.New("must at least pass one dataType in DataTypeIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data_type IN (?)", dataType))
}

// DataTypeLike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeLike(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type LIKE ?", dataType))
}

// DataTypeLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeLt(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type < ?", dataType))
}

// DataTypeLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeLte(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type <= ?", dataType))
}

// DataTypeNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeNe(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type != ?", dataType))
}

// DataTypeNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeNotIn(dataType ...string) AccessVMRecordQuerySet {
	if len(dataType) == 0 {
		qs.db.AddError(errors.New("must at least pass one dataType in DataTypeNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("data_type NOT IN (?)", dataType))
}

// DataTypeNotlike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DataTypeNotlike(dataType string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("data_type NOT LIKE ?", dataType))
}

// Delete is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) Delete() error {
	return qs.db.Delete(AccessVMRecord{}).Error
}

// DeleteNum is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DeleteNum() (int64, error) {
	db := qs.db.Delete(AccessVMRecord{})
	return db.RowsAffected, db.Error
}

// DeleteNumUnscoped is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) DeleteNumUnscoped() (int64, error) {
	db := qs.db.Unscoped().Delete(AccessVMRecord{})
	return db.RowsAffected, db.Error
}

// GetDB is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) GetDB() *gorm.DB {
	return qs.db
}

// GetUpdater is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) GetUpdater() AccessVMRecordUpdater {
	return NewAccessVMRecordUpdater(qs.db)
}

// Limit is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) Limit(limit int) AccessVMRecordQuerySet {
	return qs.w(qs.db.Limit(limit))
}

// Offset is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) Offset(offset int) AccessVMRecordQuerySet {
	return qs.w(qs.db.Offset(offset))
}

// One is used to retrieve one result. It returns gorm.ErrRecordNotFound
// if nothing was fetched
func (qs AccessVMRecordQuerySet) One(ret *AccessVMRecord) error {
	return qs.db.First(ret).Error
}

// OrderAscByBcsClusterId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByBcsClusterId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("bcs_cluster_id ASC"))
}

// OrderAscByBkBaseDataId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByBkBaseDataId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("bk_base_data_id ASC"))
}

// OrderAscByDataType is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByDataType() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("data_type ASC"))
}

// OrderAscByRemark is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByRemark() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("remark ASC"))
}

// OrderAscByResultTableId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByResultTableId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("result_table_id ASC"))
}

// OrderAscByStorageClusterID is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByStorageClusterID() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("storage_cluster_id ASC"))
}

// OrderAscByVmClusterId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByVmClusterId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("vm_cluster_id ASC"))
}

// OrderAscByVmResultTableId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderAscByVmResultTableId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("vm_result_table_id ASC"))
}

// OrderDescByBcsClusterId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByBcsClusterId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("bcs_cluster_id DESC"))
}

// OrderDescByBkBaseDataId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByBkBaseDataId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("bk_base_data_id DESC"))
}

// OrderDescByDataType is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByDataType() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("data_type DESC"))
}

// OrderDescByRemark is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByRemark() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("remark DESC"))
}

// OrderDescByResultTableId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByResultTableId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("result_table_id DESC"))
}

// OrderDescByStorageClusterID is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByStorageClusterID() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("storage_cluster_id DESC"))
}

// OrderDescByVmClusterId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByVmClusterId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("vm_cluster_id DESC"))
}

// OrderDescByVmResultTableId is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) OrderDescByVmResultTableId() AccessVMRecordQuerySet {
	return qs.w(qs.db.Order("vm_result_table_id DESC"))
}

// RemarkEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkEq(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark = ?", remark))
}

// RemarkGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkGt(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark > ?", remark))
}

// RemarkGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkGte(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark >= ?", remark))
}

// RemarkIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkIn(remark ...string) AccessVMRecordQuerySet {
	if len(remark) == 0 {
		qs.db.AddError(errors.New("must at least pass one remark in RemarkIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("remark IN (?)", remark))
}

// RemarkLike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkLike(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark LIKE ?", remark))
}

// RemarkLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkLt(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark < ?", remark))
}

// RemarkLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkLte(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark <= ?", remark))
}

// RemarkNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkNe(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark != ?", remark))
}

// RemarkNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkNotIn(remark ...string) AccessVMRecordQuerySet {
	if len(remark) == 0 {
		qs.db.AddError(errors.New("must at least pass one remark in RemarkNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("remark NOT IN (?)", remark))
}

// RemarkNotlike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) RemarkNotlike(remark string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("remark NOT LIKE ?", remark))
}

// ResultTableIdEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdEq(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id = ?", resultTableId))
}

// ResultTableIdGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdGt(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id > ?", resultTableId))
}

// ResultTableIdGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdGte(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id >= ?", resultTableId))
}

// ResultTableIdIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdIn(resultTableId ...string) AccessVMRecordQuerySet {
	if len(resultTableId) == 0 {
		qs.db.AddError(errors.New("must at least pass one resultTableId in ResultTableIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("result_table_id IN (?)", resultTableId))
}

// ResultTableIdLike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdLike(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id LIKE ?", resultTableId))
}

// ResultTableIdLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdLt(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id < ?", resultTableId))
}

// ResultTableIdLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdLte(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id <= ?", resultTableId))
}

// ResultTableIdNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdNe(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id != ?", resultTableId))
}

// ResultTableIdNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdNotIn(resultTableId ...string) AccessVMRecordQuerySet {
	if len(resultTableId) == 0 {
		qs.db.AddError(errors.New("must at least pass one resultTableId in ResultTableIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("result_table_id NOT IN (?)", resultTableId))
}

// ResultTableIdNotlike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) ResultTableIdNotlike(resultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("result_table_id NOT LIKE ?", resultTableId))
}

// StorageClusterIDEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDEq(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id = ?", storageClusterID))
}

// StorageClusterIDGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDGt(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id > ?", storageClusterID))
}

// StorageClusterIDGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDGte(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id >= ?", storageClusterID))
}

// StorageClusterIDIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDIn(storageClusterID ...uint) AccessVMRecordQuerySet {
	if len(storageClusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one storageClusterID in StorageClusterIDIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("storage_cluster_id IN (?)", storageClusterID))
}

// StorageClusterIDLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDLt(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id < ?", storageClusterID))
}

// StorageClusterIDLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDLte(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id <= ?", storageClusterID))
}

// StorageClusterIDNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDNe(storageClusterID uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("storage_cluster_id != ?", storageClusterID))
}

// StorageClusterIDNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) StorageClusterIDNotIn(storageClusterID ...uint) AccessVMRecordQuerySet {
	if len(storageClusterID) == 0 {
		qs.db.AddError(errors.New("must at least pass one storageClusterID in StorageClusterIDNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("storage_cluster_id NOT IN (?)", storageClusterID))
}

// VmClusterIdEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdEq(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id = ?", vmClusterId))
}

// VmClusterIdGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdGt(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id > ?", vmClusterId))
}

// VmClusterIdGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdGte(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id >= ?", vmClusterId))
}

// VmClusterIdIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdIn(vmClusterId ...uint) AccessVMRecordQuerySet {
	if len(vmClusterId) == 0 {
		qs.db.AddError(errors.New("must at least pass one vmClusterId in VmClusterIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("vm_cluster_id IN (?)", vmClusterId))
}

// VmClusterIdLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdLt(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id < ?", vmClusterId))
}

// VmClusterIdLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdLte(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id <= ?", vmClusterId))
}

// VmClusterIdNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdNe(vmClusterId uint) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_cluster_id != ?", vmClusterId))
}

// VmClusterIdNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmClusterIdNotIn(vmClusterId ...uint) AccessVMRecordQuerySet {
	if len(vmClusterId) == 0 {
		qs.db.AddError(errors.New("must at least pass one vmClusterId in VmClusterIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("vm_cluster_id NOT IN (?)", vmClusterId))
}

// VmResultTableIdEq is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdEq(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id = ?", vmResultTableId))
}

// VmResultTableIdGt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdGt(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id > ?", vmResultTableId))
}

// VmResultTableIdGte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdGte(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id >= ?", vmResultTableId))
}

// VmResultTableIdIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdIn(vmResultTableId ...string) AccessVMRecordQuerySet {
	if len(vmResultTableId) == 0 {
		qs.db.AddError(errors.New("must at least pass one vmResultTableId in VmResultTableIdIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("vm_result_table_id IN (?)", vmResultTableId))
}

// VmResultTableIdLike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdLike(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id LIKE ?", vmResultTableId))
}

// VmResultTableIdLt is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdLt(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id < ?", vmResultTableId))
}

// VmResultTableIdLte is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdLte(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id <= ?", vmResultTableId))
}

// VmResultTableIdNe is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdNe(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id != ?", vmResultTableId))
}

// VmResultTableIdNotIn is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdNotIn(vmResultTableId ...string) AccessVMRecordQuerySet {
	if len(vmResultTableId) == 0 {
		qs.db.AddError(errors.New("must at least pass one vmResultTableId in VmResultTableIdNotIn"))
		return qs.w(qs.db)
	}
	return qs.w(qs.db.Where("vm_result_table_id NOT IN (?)", vmResultTableId))
}

// VmResultTableIdNotlike is an autogenerated method
// nolint: dupl
func (qs AccessVMRecordQuerySet) VmResultTableIdNotlike(vmResultTableId string) AccessVMRecordQuerySet {
	return qs.w(qs.db.Where("vm_result_table_id NOT LIKE ?", vmResultTableId))
}

// SetBcsClusterId is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetBcsClusterId(bcsClusterId string) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.BcsClusterId)] = bcsClusterId
	return u
}

// SetBkBaseDataId is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetBkBaseDataId(bkBaseDataId uint) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.BkBaseDataId)] = bkBaseDataId
	return u
}

// SetDataType is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetDataType(dataType string) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.DataType)] = dataType
	return u
}

// SetRemark is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetRemark(remark string) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.Remark)] = remark
	return u
}

// SetResultTableId is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetResultTableId(resultTableId string) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.ResultTableId)] = resultTableId
	return u
}

// SetStorageClusterID is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetStorageClusterID(storageClusterID uint) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.StorageClusterID)] = storageClusterID
	return u
}

// SetVmClusterId is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetVmClusterId(vmClusterId uint) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.VmClusterId)] = vmClusterId
	return u
}

// SetVmResultTableId is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) SetVmResultTableId(vmResultTableId string) AccessVMRecordUpdater {
	u.fields[string(AccessVMRecordDBSchema.VmResultTableId)] = vmResultTableId
	return u
}

// Update is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) Update() error {
	return u.db.Updates(u.fields).Error
}

// UpdateNum is an autogenerated method
// nolint: dupl
func (u AccessVMRecordUpdater) UpdateNum() (int64, error) {
	db := u.db.Updates(u.fields)
	return db.RowsAffected, db.Error
}

// ===== END of query set AccessVMRecordQuerySet

// ===== BEGIN of AccessVMRecord modifiers

// AccessVMRecordDBSchemaField describes database schema field. It requires for method 'Update'
type AccessVMRecordDBSchemaField string

// String method returns string representation of field.
// nolint: dupl
func (f AccessVMRecordDBSchemaField) String() string {
	return string(f)
}

// AccessVMRecordDBSchema stores db field names of AccessVMRecord
var AccessVMRecordDBSchema = struct {
	DataType         AccessVMRecordDBSchemaField
	ResultTableId    AccessVMRecordDBSchemaField
	BcsClusterId     AccessVMRecordDBSchemaField
	StorageClusterID AccessVMRecordDBSchemaField
	VmClusterId      AccessVMRecordDBSchemaField
	BkBaseDataId     AccessVMRecordDBSchemaField
	VmResultTableId  AccessVMRecordDBSchemaField
	Remark           AccessVMRecordDBSchemaField
}{

	DataType:         AccessVMRecordDBSchemaField("data_type"),
	ResultTableId:    AccessVMRecordDBSchemaField("result_table_id"),
	BcsClusterId:     AccessVMRecordDBSchemaField("bcs_cluster_id"),
	StorageClusterID: AccessVMRecordDBSchemaField("storage_cluster_id"),
	VmClusterId:      AccessVMRecordDBSchemaField("vm_cluster_id"),
	BkBaseDataId:     AccessVMRecordDBSchemaField("bk_base_data_id"),
	VmResultTableId:  AccessVMRecordDBSchemaField("vm_result_table_id"),
	Remark:           AccessVMRecordDBSchemaField("remark"),
}

// Update updates AccessVMRecord fields by primary key
// nolint: dupl
func (o *AccessVMRecord) Update(db *gorm.DB, fields ...AccessVMRecordDBSchemaField) error {
	dbNameToFieldName := map[string]interface{}{
		"data_type":          o.DataType,
		"result_table_id":    o.ResultTableId,
		"bcs_cluster_id":     o.BcsClusterId,
		"storage_cluster_id": o.StorageClusterID,
		"vm_cluster_id":      o.VmClusterId,
		"bk_base_data_id":    o.BkBaseDataId,
		"vm_result_table_id": o.VmResultTableId,
		"remark":             o.Remark,
	}
	u := map[string]interface{}{}
	for _, f := range fields {
		fs := f.String()
		u[fs] = dbNameToFieldName[fs]
	}
	if err := db.Model(o).Updates(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}

		return fmt.Errorf("can't update AccessVMRecord %v fields %v: %s",
			o, fields, err)
	}

	return nil
}

// AccessVMRecordUpdater is an AccessVMRecord updates manager
type AccessVMRecordUpdater struct {
	fields map[string]interface{}
	db     *gorm.DB
}

// NewAccessVMRecordUpdater creates new AccessVMRecord updater
// nolint: dupl
func NewAccessVMRecordUpdater(db *gorm.DB) AccessVMRecordUpdater {
	return AccessVMRecordUpdater{
		fields: map[string]interface{}{},
		db:     db.Model(&AccessVMRecord{}),
	}
}

// ===== END of AccessVMRecord modifiers

// ===== END of all query sets
