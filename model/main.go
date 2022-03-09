package model

import "xupt-scholarship/db"

// CheckDuplicatesRows 检查是否存在重复键值
func CheckDuplicatesRows(table interface{}, key string, value interface{}) (bool, error) {
	res := db.Mysql.First(&table, key+" = ?", value)
	if res.Error != nil {
		return false, res.Error
	} else {
		if res.RowsAffected > 0 {
			return true, nil
		} else {
			return true, nil
		}
	}
}
