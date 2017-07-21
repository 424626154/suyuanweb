package controllers

import (
	"github.com/Luxurioust/excelize"
	"github.com/astaxie/beego"
	"path"
	"strconv"
	"suyuanweb/models/bs"
	"time"
)

//excel 服务器
type XLSXController struct {
	beego.Controller
}

func (c *XLSXController) WriteXlsx() {
	xlsx := excelize.NewFile()
	sheet_name := "入库登记表"
	oldName := "Sheet1"
	xlsx.SetSheetName(oldName, sheet_name)
	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"size":20,"bold":true}}`)
	if err != nil {
		beego.Error(err)
	}
	//行高列高设置
	xlsx.SetColWidth(oldName, "A", "B", 12)
	xlsx.SetColWidth(oldName, "C", "C", 15)
	xlsx.SetColWidth(oldName, "D", "D", 20)
	xlsx.SetColWidth(oldName, "E", "J", 12)
	xlsx.SetColWidth(oldName, "K", "K", 15)
	xlsx.SetRowHeight(oldName, 0, 30)
	xlsx.SetRowHeight(oldName, 1, 18)
	//入库登记表标题
	xlsx.MergeCell(oldName, "A1", "K1")
	xlsx.SetCellStyle(oldName, "A1", "A1", style)
	xlsx.SetCellValue(oldName, "A1", "入库登记表")
	//入库登记表第二栏

	xlsx.MergeCell(oldName, "A2", "G2")
	xlsx.MergeCell(oldName, "H2", "I2")
	xlsx.MergeCell(oldName, "J2", "K2")
	xlsx.SetCellValue(oldName, "A2", "供货单位：")
	xlsx.SetCellValue(oldName, "H2", "入库类别:")
	xlsx.SetCellValue(oldName, "J2", "原库存:")

	categories := map[string]string{"A3": "序号", "B3": "入库时间", "C3": "订单号码", "D3": "产品名称", "E3": "单位", "F3": "规格", "G3": "单价格", "H3": "数量", "I3": "金额", "J3": "经办人", "K3": "备注"}
	// values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	for k, v := range categories {
		xlsx.SetCellValue(oldName, k, v)
	}
	t := time.Now().Format("2006-01-02_15-04-05")
	xlsx_name := "入库登记表" + t + ".xlsx"
	// xlsx_name = "入库登记表2006_01_02.xlsx"
	beego.Debug("xlsx_name:", xlsx_name)
	file_path := path.Join("bsfilehosting", xlsx_name)
	err = xlsx.SaveAs(file_path)
	if err != nil {
		beego.Error(err)
	}
	c.Ctx.WriteString("{\"path\":" + file_path + "}")
}

func (c *XLSXController) ReadXlsx() {
	xlsx, err := excelize.OpenFile("./bsfilehosting/入库登记表.xlsx")
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	// Get value from cell by given sheet index and axis.
	// cell := xlsx.GetCellValue("Sheet1", "B2")

	cell_str := ""
	beego.Debug("A1style", xlsx.GetCellStyle("Sheet1", "A1"))
	beego.Debug("A2style", xlsx.GetCellStyle("Sheet1", "A2"))
	beego.Debug("A3style", xlsx.GetCellStyle("Sheet1", "A3"))
	// for k, v := range xlsx.GetSheetMap() {
	// 	beego.Debug(k, v)
	// }

	// // Get sheet index.
	index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	for _, row := range rows {
		for _, colCell := range row {
			cell_str += colCell + "\t"
		}
	}

	c.Ctx.WriteString(cell_str)
}

//导出入库表
func (c *XLSXController) ExportInCommodity() {
	op := c.Input().Get("op")
	p := c.Input().Get("p")
	i_p := 1
	temp_p, err := strconv.Atoi(p)
	if err != nil {
		beego.Error(err)
	} else {
		i_p = temp_p
	}
	var incommoditys []bs.BsInCommodity
	switch op {
	case "epage":
		if len(p) == 0 {
			json_err := "参数错误"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		pageSize := 10
		incommoditys, err = bs.GetPageInCommoditys(pageSize, i_p)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	case "eall":
		incommoditys, err = bs.GetInCommoditys()
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	}
	title, path, err := WriteInCommodityXlsx(incommoditys)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	_, err = bs.AddBsExport(title, path)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	json_msg := title
	ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
	c.Ctx.WriteString(ajax_json)
}

func WriteInCommodityXlsx(incommoditys []bs.BsInCommodity) (string, string, error) {
	xlsx := excelize.NewFile()
	sheet_name := "入库登记表"
	oldName := "Sheet1"
	xlsx.SetSheetName(oldName, sheet_name)
	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"size":20,"bold":true}}`)
	if err != nil {
		beego.Error(err)
	}
	style1, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"}}`)
	if err != nil {
		beego.Error(err)
	}
	//行高列高设置
	xlsx.SetColWidth(oldName, "A", "B", 12)
	xlsx.SetColWidth(oldName, "C", "C", 15)
	xlsx.SetColWidth(oldName, "D", "D", 20)
	xlsx.SetColWidth(oldName, "E", "J", 12)
	xlsx.SetColWidth(oldName, "K", "K", 15)
	xlsx.SetRowHeight(oldName, 0, 30)
	xlsx.SetRowHeight(oldName, 1, 18)
	//入库登记表标题
	xlsx.MergeCell(oldName, "A1", "K1")
	xlsx.SetCellStyle(oldName, "A1", "A1", style)
	xlsx.SetCellValue(oldName, "A1", "入库登记表")

	categories := map[string]string{"A2": "序号", "B2": "入库时间", "C2": "订单号码", "D2": "产品名称", "E2": "单位", "F2": "规格", "G2": "单价格", "H2": "数量", "I2": "金额", "J2": "经办人", "K2": "备注"}
	for k, v := range categories {
		xlsx.SetCellValue(oldName, k, v)
	}

	index := 3
	xlsx.SetCellStyle(oldName, "A"+strconv.Itoa(index), "A"+strconv.Itoa(index+len(incommoditys)-1), style1)
	style_date, err := xlsx.NewStyle(`{"number_format": 14}`)
	if err != nil {
	    beego.Debug(err)
	}
	xlsx.SetCellStyle(oldName, "B"+strconv.Itoa(index), "B"+strconv.Itoa(index+len(incommoditys)-1), style_date)
	// beego.Debug("incommoditys:",incommoditys)
	for i := 0; i < len(incommoditys); i++ {
		incommodity := incommoditys[i]
		items := map[string]interface{}{}
		index_str := strconv.Itoa(index)
		items["A"+index_str] = i + 1
		items["B"+index_str] = incommodity.CDate
		items["C"+index_str] = incommodity.OrderNumber
		items["D"+index_str] = incommodity.Name
		items["E"+index_str] = incommodity.Unit
		items["F"+index_str] = incommodity.Spec
		items["G"+index_str] = incommodity.UnitPrice
		items["H"+index_str] = incommodity.Number
		items["I"+index_str] = incommodity.Total
		items["J"+index_str] = incommodity.Operator
		items["K"+index_str] = incommodity.Remarks

		for k, v := range items {
			xlsx.SetCellValue(oldName, k, v)
		}
		index++
	}

	t := time.Now()
	t1 := t.Format("2006-01-02_15-04-05")
	t2 := t.Format("2006-01-02 15:04:05")
	xlsx_name := "入库登记表" + t1 + ".xlsx"
	file_name := "入库登记表" + t2
	// xlsx_name = "入库登记表2006_01_02.xlsx"
	beego.Debug("xlsx_name:", xlsx_name)
	file_path := path.Join("bsfilehosting", xlsx_name)
	err = xlsx.SaveAs(file_path)
	if err != nil {
		beego.Error(err)
		return "", "", err
	}
	return file_name, file_path, nil
}

func (c *XLSXController) ExportOutCommodity() {
	op := c.Input().Get("op")
	p := c.Input().Get("p")
	i_p := 1
	temp_p, err := strconv.Atoi(p)
	if err != nil {
		beego.Error(err)
	} else {
		i_p = temp_p
	}
	var incommoditys []bs.BsOutCommodity
	switch op {
	case "epage":
		if len(p) == 0 {
			json_err := "参数错误"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		pageSize := 10
		incommoditys, err = bs.GetPageOutCommoditys(pageSize, i_p)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	case "eall":
		incommoditys, err = bs.GetOutCommoditys()
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	}
	title, path, err := WriteOutCommodityXlsx(incommoditys)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	_, err = bs.AddBsExport(title, path)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	json_msg := title
	ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
	c.Ctx.WriteString(ajax_json)
}

func WriteOutCommodityXlsx(incommoditys []bs.BsOutCommodity) (string, string, error) {
	xlsx := excelize.NewFile()
	sheet_name := "入库登记表"
	oldName := "Sheet1"
	xlsx.SetSheetName(oldName, sheet_name)
	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"size":20,"bold":true}}`)
	if err != nil {
		beego.Error(err)
	}
	//行高列高设置
	xlsx.SetColWidth(oldName, "A", "B", 12)
	xlsx.SetColWidth(oldName, "C", "C", 15)
	xlsx.SetColWidth(oldName, "D", "D", 20)
	xlsx.SetColWidth(oldName, "E", "J", 12)
	xlsx.SetColWidth(oldName, "K", "K", 15)
	xlsx.SetRowHeight(oldName, 0, 30)
	xlsx.SetRowHeight(oldName, 1, 18)
	//出库登记表标题
	xlsx.MergeCell(oldName, "A1", "K1")
	xlsx.SetCellStyle(oldName, "A1", "A1", style)
	xlsx.SetCellValue(oldName, "A1", "出库登记表")
	//出库登记表第二栏

	xlsx.MergeCell(oldName, "A2", "I2")
	xlsx.MergeCell(oldName, "J2", "K2")
	xlsx.SetCellValue(oldName, "J2", "出库单号:")

	//出库登记表第三栏
	xlsx.MergeCell(oldName, "A3", "I3")
	xlsx.MergeCell(oldName, "J3", "K3")
	xlsx.SetCellValue(oldName, "A3", "领货单位:")
	xlsx.SetCellValue(oldName, "J3", "出库类别:")

	categories := map[string]string{"A4": "序号", "B4": "入库时间", "C4": "订单号码", "D4": "产品名称", "E4": "单位", "F4": "规格", "G4": "单价格", "H4": "数量", "I4": "金额", "J4": "经办人", "K4": "备注"}
	// values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	for k, v := range categories {
		xlsx.SetCellValue(oldName, k, v)
	}

	index := 5
	beego.Debug("incommoditys:", incommoditys)
	for i := 0; i < len(incommoditys); i++ {
		incommodity := incommoditys[i]
		items := map[string]interface{}{}
		index_str := strconv.Itoa(index)
		items["A"+index_str] = i + 1
		items["B"+index_str] = incommodity.CDate
		items["C"+index_str] = incommodity.OrderNumber
		items["D"+index_str] = incommodity.Name
		items["E"+index_str] = incommodity.Unit
		items["F"+index_str] = incommodity.Spec
		items["G"+index_str] = incommodity.UnitPrice
		items["H"+index_str] = incommodity.Number
		items["I"+index_str] = incommodity.Total
		items["J"+index_str] = incommodity.Operator
		items["K"+index_str] = incommodity.Remarks

		for k, v := range items {
			xlsx.SetCellValue(oldName, k, v)
		}
		index++
	}

	t := time.Now()
	t1 := t.Format("2006-01-02_15-04-05")
	t2 := t.Format("2006-01-02 15:04:05")
	xlsx_name := "出库登记表" + t1 + ".xlsx"
	file_name := "出库登记表" + t2
	// xlsx_name = "入库登记表2006_01_02.xlsx"
	beego.Debug("xlsx_name:", xlsx_name)
	file_path := path.Join("bsfilehosting", xlsx_name)
	err = xlsx.SaveAs(file_path)
	if err != nil {
		beego.Error(err)
		return "", "", err
	}
	return file_name, file_path, nil
}

func (c *XLSXController) ExportAllCommodity() {
	op := c.Input().Get("op")
	p := c.Input().Get("p")
	i_p := 1
	temp_p, err := strconv.Atoi(p)
	if err != nil {
		beego.Error(err)
	} else {
		i_p = temp_p
	}
	var allcommoditys []bs.BsAllCommodity
	switch op {
	case "epage":
		if len(p) == 0 {
			json_err := "参数错误"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		pageSize := 10
		allcommoditys, err = bs.GetPageAllCommoditys(pageSize, i_p)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	case "eall":
		allcommoditys, err = bs.GetAllCommoditys()
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
		}
		break
	}
	title, path, err := WriteAllCommodityXlsx(allcommoditys)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	_, err = bs.AddBsExport(title, path)
	if err != nil {
		beego.Error(err)
		json_err := err.Error()
		ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
		c.Ctx.WriteString(ajax_json)
	}
	json_msg := title
	ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
	c.Ctx.WriteString(ajax_json)
}

func WriteAllCommodityXlsx(allcommoditys []bs.BsAllCommodity) (string, string, error) {
	xlsx := excelize.NewFile()
	sheet_name := "进销存登记表"
	oldName := "Sheet1"
	xlsx.SetSheetName(oldName, sheet_name)
	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center","vertical":"center"},"font":{"size":20,"bold":true}}`)
	if err != nil {
		beego.Error(err)
	}
	//行高列高设置
	xlsx.SetColWidth(oldName, "A", "B", 12)
	xlsx.SetColWidth(oldName, "C", "C", 6)
	xlsx.SetColWidth(oldName, "D", "D", 12)
	xlsx.SetColWidth(oldName, "E", "N", 20)
	xlsx.SetRowHeight(oldName, 0, 30)
	xlsx.SetRowHeight(oldName, 1, 18)
	//入库登记表标题
	xlsx.MergeCell(oldName, "A1", "N1")
	xlsx.SetCellStyle(oldName, "A1", "N1", style)
	xlsx.SetCellValue(oldName, "A1", "进销存登记表")
	//入库登记表第二栏

	// xlsx.MergeCell(oldName, "A2", "G2")
	// xlsx.MergeCell(oldName, "H2", "I2")
	// xlsx.MergeCell(oldName, "J2", "K2")
	// xlsx.SetCellValue(oldName, "A2", "供货单位：")
	// xlsx.SetCellValue(oldName, "H2","入库类别:")
	// xlsx.SetCellValue(oldName, "J2","原库存:")

	categories := map[string]string{"A2": "商品名称", "B2": "规格", "C2": "单位", "D2": "日期", "E2": "期初结余数量", "F2": "期初结余金额", "G2": "本期购进数量", "H2": "本期购进单价", "I2": "本期购进金额", "J2": "本期发出数量", "K2": "本期发出单价", "L2": "本期发出金额", "M2": "本期结存数量", "N2": "本期结存金额"}
	for k, v := range categories {
		xlsx.SetCellValue(oldName, k, v)
	}

	index := 3
	// beego.Debug("incommoditys:",incommoditys)
	for i := 0; i < len(allcommoditys); i++ {
		allcommodity := allcommoditys[i]
		items := map[string]interface{}{}
		index_str := strconv.Itoa(index)
		items["A"+index_str] = allcommodity.Name
		items["B"+index_str] = allcommodity.Spec
		items["C"+index_str] = allcommodity.Unit
		items["D"+index_str] = allcommodity.CDate
		items["E"+index_str] = allcommodity.StartNumber
		items["F"+index_str] = allcommodity.StartTotal
		items["G"+index_str] = allcommodity.InNumber
		items["H"+index_str] = allcommodity.InUnitPrice
		items["I"+index_str] = allcommodity.InTotal
		items["J"+index_str] = allcommodity.OutNumber
		items["K"+index_str] = allcommodity.OutUnitPrice
		items["L"+index_str] = allcommodity.OutTotal
		items["M"+index_str] = allcommodity.EndNumber
		items["N"+index_str] = allcommodity.EndTotal

		for k, v := range items {
			xlsx.SetCellValue(oldName, k, v)
		}
		index++
	}

	t := time.Now()
	t1 := t.Format("2006-01-02_15-04-05")
	t2 := t.Format("2006-01-02 15:04:05")
	xlsx_name := "进销存登记表" + t1 + ".xlsx"
	file_name := "进销存记表" + t2
	// xlsx_name = "入库登记表2006_01_02.xlsx"
	beego.Debug("xlsx_name:", xlsx_name)
	file_path := path.Join("bsfilehosting", xlsx_name)
	err = xlsx.SaveAs(file_path)
	if err != nil {
		beego.Error(err)
		return "", "", err
	}
	return file_name, file_path, nil
}

func (c *XLSXController) ImportInCommodity() {
	if c.Ctx.Input.IsPost() {
		file_path := "filepath"
		_, fh, err := c.GetFile(file_path)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
		if fh != nil {
			file_name := fh.Filename
			t := time.Now()
			t1 := t.Format("2006-01-02_15-04-05")
			save_file_name := t1 + "_" + file_name
			save_file_path := path.Join("bsfilehosting", save_file_name)
			err = c.SaveToFile(file_path, save_file_path)
			if err != nil {
				beego.Error(err)
				json_err := err.Error()
				ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
				c.Ctx.WriteString(ajax_json)
				return
			}
			ReadInCommodityXlsx(c, save_file_path)
			json_msg := "导入完成"
			ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
			c.Ctx.WriteString(ajax_json)
		} else {
			json_err := "获取文件失败"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
	}
}

func ReadInCommodityXlsx(c *XLSXController, file_path string) {
	xlsx, err := excelize.OpenFile(file_path)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	oldName := "Sheet1"
	index := xlsx.GetSheetIndex(oldName)
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	incommoditys := make([]bs.BsInCommodity, 0)
	for i, row := range rows {
		beego.Debug("row:", row)
		if i > 1 {
			var incommodity bs.BsInCommodity
			iserr := false
			for j, colCell := range row {
				if j == 1 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					incommodity.CDate = colCell
				}
				if j == 3 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					incommodity.Name = colCell
				}
				if j == 4 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					incommodity.Unit = colCell
				}
				if j == 5 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					incommodity.Spec = colCell
				}
				if j == 6 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					iunitprice, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					incommodity.UnitPrice = iunitprice
				}
				if j == 7 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					inumber, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					incommodity.Number = inumber
				}
				if j == 9 {
					incommodity.Operator = colCell
				}
				if j == 10 {
					incommodity.Remarks = colCell
				}
			}
			if iserr == false {
				incommoditys = append(incommoditys, incommodity)
			}
		}
	}
	if len(incommoditys) > 0{
		for i := 0; i < len(incommoditys); i++ {
				item := incommoditys[i]
				_,err := bs.AddInCommodityFObj(item)
				if err != nil{
					beego.Error(err)
				}
		}
	}
}

// 导入出库数据
func (c *XLSXController) ImportOutCommodity() {
	if c.Ctx.Input.IsPost() {
		file_path := "filepath"
		_, fh, err := c.GetFile(file_path)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
		if fh != nil {
			file_name := fh.Filename
			t := time.Now()
			t1 := t.Format("2006-01-02_15-04-05")
			save_file_name := t1 + "_" + file_name
			save_file_path := path.Join("bsfilehosting", save_file_name)
			err = c.SaveToFile(file_path, save_file_path)
			if err != nil {
				beego.Error(err)
				json_err := err.Error()
				ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
				c.Ctx.WriteString(ajax_json)
				return
			}
			ReadOutCommodityXlsx(c, save_file_path)
			json_msg := "导入完成"
			ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
			c.Ctx.WriteString(ajax_json)
		} else {
			json_err := "获取文件失败"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
	}
}

func ReadOutCommodityXlsx(c *XLSXController, file_path string) {
	xlsx, err := excelize.OpenFile(file_path)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	oldName := "Sheet1"
	index := xlsx.GetSheetIndex(oldName)
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	outcommoditys := make([]bs.BsOutCommodity, 0)
	for i, row := range rows {
		beego.Debug("row:", row)
		if i > 1 {
			var outcommodity bs.BsOutCommodity
			iserr := false
			for j, colCell := range row {
				if j == 1 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					outcommodity.CDate = colCell
				}
				if j == 3 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					outcommodity.Name = colCell
				}
				if j == 4 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					outcommodity.Unit = colCell
				}
				if j == 5 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					outcommodity.Spec = colCell
				}
				if j == 6 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					iunitprice, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					outcommodity.UnitPrice = iunitprice
				}
				if j == 7 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					inumber, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					outcommodity.Number = inumber
				}
				if j == 9 {
					outcommodity.CargoUnit = colCell
				}
				if j == 10 {
					outcommodity.Operator = colCell
				}
				if j == 11 {
					outcommodity.Remarks = colCell
				}
			}
			if iserr == false {
				outcommoditys = append(outcommoditys, outcommodity)
			}
		}
	}
	if len(outcommoditys) > 0{
		for i := 0; i < len(outcommoditys); i++ {
				item := outcommoditys[i]
				_,err := bs.AddOutCommodityFObj(item)
				if err != nil{
					beego.Error(err)
				}
		}
	}
}

func (c *XLSXController) ImportAllCommodity() {
	if c.Ctx.Input.IsPost() {
		file_path := "filepath"
		_, fh, err := c.GetFile(file_path)
		if err != nil {
			beego.Error(err)
			json_err := err.Error()
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
		if fh != nil {
			file_name := fh.Filename
			t := time.Now()
			t1 := t.Format("2006-01-02_15-04-05")
			save_file_name := t1 + "_" + file_name
			save_file_path := path.Join("bsfilehosting", save_file_name)
			err = c.SaveToFile(file_path, save_file_path)
			if err != nil {
				beego.Error(err)
				json_err := err.Error()
				ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
				c.Ctx.WriteString(ajax_json)
				return
			}
			ReadAllCommodityXlsx(c, save_file_path)
			json_msg := "导入完成"
			ajax_json := "{\"code\":0,\"msg\":\"" + json_msg + "\"}"
			c.Ctx.WriteString(ajax_json)
		} else {
			json_err := "获取文件失败"
			ajax_json := "{\"code\":1,\"err\":\"" + json_err + "\"}"
			c.Ctx.WriteString(ajax_json)
			return
		}
	}
}

func ReadAllCommodityXlsx(c *XLSXController, file_path string) {
	xlsx, err := excelize.OpenFile(file_path)
	if err != nil {
		c.Ctx.WriteString(err.Error())
		return
	}
	oldName := "Sheet1"
	index := xlsx.GetSheetIndex(oldName)
	rows := xlsx.GetRows("sheet" + strconv.Itoa(index))
	allcommoditys := make([]bs.BsAllCommodity, 0)
	for i, row := range rows {
		beego.Debug("row:", row)
		if i > 2 {
			var allcommodity bs.BsAllCommodity
			iserr := false
			for j, colCell := range row {
				if j == 0 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					allcommodity.Name = colCell
				}
				if j == 1 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					allcommodity.Spec = colCell
				}
				if j == 2 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					allcommodity.Unit = colCell
				}
				if j == 3 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					allcommodity.CDate = colCell
				}
				if j == 4 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					istartnumber, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					allcommodity.StartNumber = istartnumber
				}
				if j == 5 {
					if len(colCell) == 0 {
						iserr = true
						break
					}
					istarttotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iserr = true
						break
					}
					allcommodity.StartTotal = istarttotal
				}
				if j == 6 {
					iinnumber, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iinnumber = 0
					}
					allcommodity.InNumber = iinnumber
				}
				if j == 7 {
					iinunitprice, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iinunitprice = 0
					}
					allcommodity.InUnitPrice = iinunitprice
				}
				if j == 8 {
					iintotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iintotal = 0
					}
					allcommodity.InTotal = iintotal
				}
				if j == 9 {
					iouttotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iouttotal = 0
					}
					allcommodity.OutNumber = iouttotal
				}
				if j == 10 {
					iouttotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iouttotal = 0
					}
					allcommodity.OutUnitPrice = iouttotal
				}
				if j == 11 {
					iouttotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iouttotal = 0
					}
					allcommodity.OutTotal = iouttotal
				}
				if j == 12 {
					iendnumber, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iendnumber = 0
					}
					allcommodity.EndNumber = iendnumber
				}
				if j == 13 {
					iendtotal, err := strconv.Atoi(colCell)
					if err != nil {
						beego.Error(err)
						iendtotal = 0
					}
					allcommodity.EndTotal = iendtotal
				}
			}
			if iserr == false {
				allcommoditys = append(allcommoditys, allcommodity)
			}
		}
	}
	if len(allcommoditys) > 0{
		for i := 0; i < len(allcommoditys); i++ {
				item := allcommoditys[i]
				_,err := bs.AddAllCommodityFObj(item)
				if err != nil{
					beego.Error(err)
				}
		}
	}
}
