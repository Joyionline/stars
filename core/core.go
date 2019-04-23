package core

import (
	"log"
	"stars/models"

	"github.com/go-xorm/xorm"
)

type StarsCore struct {
	engine *xorm.Engine
}

func NewStarsCore(engine *xorm.Engine) *StarsCore {
	return &StarsCore{
		engine: engine,
	}
}

func (d *StarsCore) GetById(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		return nil
		data.Id = 0
		return data
	}
	return nil
}

func (d *StarsCore) GetAll() []models.StarInfo {
	datalist := []models.StarInfo{}
	err := d.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
	return nil
}

func (d *StarsCore) Search(country string) []models.StarInfo {
	datalist := []models.StarInfo{}
	err := d.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	} else {
		return datalist
	}
	return nil
}

// 逻辑删除
func (d *StarsCore) DeleteById(id int) error {
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *StarsCore) UpdateByUserInfo(data *models.StarInfo, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *StarsCore) Create(data *models.StarInfo) error {
	_, err := d.engine.Insert(data)
	return err
}
