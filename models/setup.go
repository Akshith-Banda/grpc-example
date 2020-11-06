package models

import (
	"example/proto"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func dbase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database!")
	}
	return db
}

var db = dbase()

func Database() {
	db.AutoMigrate(&Thermostat{}, &Group{})
}

func Thermo(name string, ctemp int64, tempdsply bool) string {
	db.Create(&Thermostat{Name: name, CurrentTemp: ctemp, TempDisplay: tempdsply})
	return "Thermostat created successfully!"
}

func Grp(name string) string {
	db.Create(&Group{GroupName: name})
	return "Group created!"
}

func AddTG(tname, gname string) string {

	var thermostat Thermostat
	db.Where("name = ?", tname).Find(&thermostat)
	var group Group
	db.Where("group_name = ?", gname).Find(&group)
	db.Model(Thermostat{}).Where("id = ?", thermostat.ID).Update("group_id", group.ID)
	return "thermostat added to group!"
}

func DeleteGroup(name string) string {

	var group Group
	db.Where("group_name = ?", name).Find(&group)
	db.Model(&Thermostat{}).Where("group_id = ?", group.ID).Update("group_id", 0)
	db.Delete(&Group{}, group.ID)
	return "Group deleted!"
}

func GetThermo(name string) *proto.Thermostat {

	row := db.Model(Thermostat{}).Where("name = ?", name).Select("id,name,current_temp,temp_display,group_id").Row()
	if row.Err() != nil {
		fmt.Println("error occured at row!")
	}
	var tname string
	var tid, ctemp, gid int64
	var tempdsply bool
	row.Scan(&tid, &tname, &ctemp, &tempdsply, &gid)
	info := &proto.Thermostat{Id: tid, Name: tname, CurrentTemp: ctemp, TempDsply: tempdsply, GroupId: gid}
	return info
}
