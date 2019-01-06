package database

type Model struct {
	Db *DB
}

func (m *Model) DetailMaterial(name string) MaterialCredit {
	db := m.Db.GetConnection()
	var detailMaterial MaterialCredit
	db.Where("name = ?", name).First(&detailMaterial)
	return detailMaterial
}

func (m *Model) DetailUnit(name string) UnitNumber {
	db := m.Db.GetConnection()
	var detailUnit UnitNumber
	db.Where("name = ?", name).First(&detailUnit)
	return detailUnit
}

func (m *Model) CreateUnit(storeUnitNumber UnitNumber) UnitNumber {
	db := m.Db.GetConnection()
	db.Create(&storeUnitNumber)
	return storeUnitNumber
}

func (m *Model) UpdateMaterial(materialID uint, storeMaterial MaterialCredit) MaterialCredit {
	db := m.Db.GetConnection()
	db.Model(&storeMaterial).Where("id = ?", materialID).Updates(storeMaterial)
	return storeMaterial
}

func (m *Model) CreateMaterial(storeMaterialCredit MaterialCredit) MaterialCredit {
	db := m.Db.GetConnection()
	db.Create(&storeMaterialCredit)
	return storeMaterialCredit
}

func (m *Model) UpdateUnit(unitId uint, updateUnitNumber UnitNumber) UnitNumber {
	db := m.Db.GetConnection()
	db.Model(&updateUnitNumber).Where("id = ?", unitId).Updates(updateUnitNumber)
	return updateUnitNumber
}

func (m *Model) GetListUnitNumber(listUnit []string) []UnitNumber {
	db := m.Db.GetConnection()
	var listUnitNumber []UnitNumber
	db.Where("name in (?)", listUnit).Find(&listUnitNumber)
	return listUnitNumber
}

func NewModel() *Model {
	var model Model
	model.Db = NewDB()
	model.Db.SetPath("./database/roman.db")
	model.Db.CreateConnection()
	return &model
}
