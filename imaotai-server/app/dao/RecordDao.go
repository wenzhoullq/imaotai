package dao

import (
	"github.com/jinzhu/gorm"
	"imaotai_helper/dto/record"
	"imaotai_helper/init/db"
)

type RecordDao struct {
	*gorm.DB
}

func NewRecordDao(ops ...func(*RecordDao)) *RecordDao {
	rd := &RecordDao{
		DB: db.DB,
	}
	for _, op := range ops {
		op(rd)
	}
	return rd
}

func (rd *RecordDao) AddRecords(records []*record.Record) error {
	record := &record.Record{}
	for _, r := range records {
		if err := rd.Table(record.TableName()).Create(r).Error; err != nil {
			return err
		}
	}
	return nil
}
