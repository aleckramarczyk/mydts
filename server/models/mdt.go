package models

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/beego/beego/v2/client/orm"

	_ "github.com/go-sql-driver/mysql"
)

type Mdt struct {
	SerialNumber string `orm:"pk;size(128)"`
	UnitName     string `orm:"size(128)"`
	UnitId       string `orm:"size(128)"`
	VehicleId    string `orm:"size(128)"`
	SignedOn     bool
	InternalIp   string `orm:"size(128)"`
	RemoteIp     string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Mdt))

	// Get database connection string from environment variables
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		dbUser = "root"
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		dbHost = "127.0.0.1"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "mdts"
	}
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)

	orm.RegisterDataBase("default", "mysql", connString)

	orm.RunSyncdb("default", false, true)
}

// AddMdt insert a new Mdt into database and returns
// last inserted SerialNumber on success.
func AddMdt(m *Mdt) (SerialNumber int64, err error) {
	o := orm.NewOrm()
	SerialNumber, err = o.Insert(m)
	return
}

// GetMdtBySerialNumber retrieves Mdt by SerialNumber. Returns error if
// SerialNumber doesn't exist
func GetMdtBySerialNumber(SerialNumber string) (v *Mdt, err error) {
	o := orm.NewOrm()
	v = &Mdt{SerialNumber: SerialNumber}
	if err = o.QueryTable(new(Mdt)).Filter("SerialNumber", SerialNumber).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllMdt retrieves all Mdt matches certain condition. Returns empty list if
// no records exist
func GetAllMdt(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Mdt))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: InvalSerialNumber order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: InvalSerialNumber order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Mdt
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateMdt updates Mdt by SerialNumber and returns error if
// the record to be updated doesn't exist
func UpdateMdtBySerialNumber(m *Mdt) (err error) {
	o := orm.NewOrm()
	v := Mdt{SerialNumber: m.SerialNumber}
	// ascertain SerialNumber exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteMdt deletes Mdt by SerialNumber and returns error if
// the record to be deleted doesn't exist
func DeleteMdt(SerialNumber string) (err error) {
	o := orm.NewOrm()
	v := Mdt{SerialNumber: SerialNumber}
	// ascertain SerialNumber exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Mdt{SerialNumber: SerialNumber}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
