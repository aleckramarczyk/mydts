package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"server/models"
	"strings"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

// MdtController operations for Mdt
type MdtController struct {
	beego.Controller
}

// URLMapping ...
func (c *MdtController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Mdt
// @Param	body		body 	models.Mdt	true		"body for Mdt content"
// @Success 201 {int} models.Mdt
// @Failure 403 body is empty
// @router / [post]
func (c *MdtController) Post() {
	var v models.Mdt

	// Saves the remote IP address, first checks the X-Real-IP header for if the service is behind a reverse proxy, if not it will grab the remote address
	v.RemoteIp = c.Ctx.Input.Header("X-Real-IP")
	if v.RemoteIp == "" {
		v.RemoteIp = strings.Split(c.Ctx.Request.RemoteAddr, ":")[0]
	}

	// Set Updated to current time
	v.Updated = time.Now()

	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	// Determine the gateway manufacturer based on supplied MAC address
	v.GatewayManufacturer = determineGatewayManufacturer(v.GatewayMac)

	//If the MDT already exists, update the MDT
	if _, err := models.GetMdtBySerialNumber(v.SerialNumber); err == nil {
		if err := models.UpdateMdtBySerialNumber(&v); err == nil {
			c.Ctx.Output.SetStatus(200) // Status OK
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else { //If the MDT doesn't exist, add the MDT
		if _, err := models.AddMdt(&v); err == nil {
			c.Ctx.Output.SetStatus(201) // Status Created
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	}
	c.ServeJSON()
}

func determineGatewayManufacturer(mac string) string {
	switch oui := strings.Replace(strings.ToUpper(mac)[0:7], "-", ":", -1); {
	case oui == "00:14:3E":
		return "Sierra Wireless"
	case oui == "2A:30:44" || oui == "00:30:44":
		return "Cradlepoint"
	default:
		return getUnknownMac(mac)
	}
}

func getUnknownMac(mac string) string {
	resp, err := http.Get("https://api.maclookup.app/v2/macs/" + strings.Replace(mac, "-", ":", -1))
	if err != nil {
		log.Printf("Error looking up unknown MAC %s", err.Error())
		return "Unknown"
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	type macResponse struct {
		Success    bool   `json:"success"`
		Found      bool   `json:"found"`
		MacPrefix  string `json:"macPrefix"`
		Company    string `json:"company"`
		Address    string `json:"address"`
		Country    string `json:"country"`
		BlockStart string `json:"blockStart"`
		BlockEnd   string `json:"blockEnd"`
		BlockSize  int    `json:"blockSize"`
		BlockType  string `json:"blockType"`
		Updated    string `json:"updated"`
		IsRand     bool   `json:"isRand"`
		IsPrivate  bool   `json:"isPrivate"`
	}

	responseData := macResponse{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		log.Printf("Error unmarshaling MAC request data %s", err.Error())
		return "Unknown"
	}

	return responseData.Company
}

// GetOne ...
// @Title Get One
// @Description get Mdt by SerialNumber
// @Param	SerialNumber		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Mdt
// @Failure 403 :SerialNumber is empty
// @router /:SerialNumber [get]
func (c *MdtController) GetOne() {
	SerialNumber := c.Ctx.Input.Param(":SerialNumber")
	v, err := models.GetMdtBySerialNumber(SerialNumber)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Mdt
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Mdt
// @Failure 403
// @router / [get]
func (c *MdtController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalSerialNumber query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMdt(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Mdt
// @Param	SerialNumber		path 	string	true		"The SerialNumber you want to update"
// @Param	body		body 	models.Mdt	true		"body for Mdt content"
// @Success 200 {object} models.Mdt
// @Failure 403 :SerialNumber is not int
// @router /:SerialNumber [put]
func (c *MdtController) Put() {
	SerialNumber := c.Ctx.Input.Param(":SerialNumber")
	v := models.Mdt{SerialNumber: SerialNumber}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMdtBySerialNumber(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Mdt
// @Param	SerialNumber		path 	string	true		"The SerialNumber you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 SerialNumber is empty
// @router /:SerialNumber [delete]
func (c *MdtController) Delete() {
	SerialNumber := c.Ctx.Input.Param(":SerialNumber")
	if err := models.DeleteMdt(SerialNumber); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
