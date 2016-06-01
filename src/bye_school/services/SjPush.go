package services

import (
	"bye_school/models"
	"bye_school/utiles"
	"errors"
	"fmt"
)

func SjPush(user_id string,msg string,title string) (err error){
				db := utiles.DB
        var devices []models.Avos
        query := db.Where("user_id = ?",user_id).Find(&devices)
        if len(devices) == 0 || query.Error != nil {
            return errors.New("devices not found")
        }

        for _,device := range devices{
            Installations := device.Installation
            fmt.Println(Installations)
            umeng := utiles.InitUmeng(Installations)
            _, err = umeng.ClientNotify(msg,title)
        }
        return err
}
	