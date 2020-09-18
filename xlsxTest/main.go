package main

import (
	"archive/zip"
	"fmt"
	"github.com/tealeg/xlsx"
	models "golangDemo/xlsxTest/model"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	ComboExport()
	//GuestExport()
}
func ComboExport() {
	var orderList []Order
	col := models.Mongo.DB("mnOracle").
		C("SALE_TJDW_LIST")
	col.Find(bson.M{}).All(&orderList)
	xlsxFile := xlsx.NewFile()
	log.Printf(`%#v`, orderList[1].Combos)
	for _, combo := range orderList[1].Combos {

		var sheet *xlsx.Sheet
		var err error
		sheet, err = xlsxFile.AddSheet(combo.GROUPNAME)
		if err != nil {
			log.Panic(err)
		}
		title := sheet.AddRow()
		title.AddCell().Value = "项目代码"
		title.AddCell().Value = "体检项目"
		title.AddCell().Value = "体检项目指标"
		title.AddCell().Value = "临床意义"
		title.AddCell().Value = "男宾"
		title.AddCell().Value = "女未婚"
		title.AddCell().Value = "女已婚"
		title.AddCell().Value = "备注"
		for _, detail := range combo.Details {
			if len(detail.SUBITEM) == 0 {
				content := sheet.AddRow()
				ccc := content.AddCell()
				birthday, _ := time.Parse("2006-01-02 15:04:05", "1990-12-13 23:00:00")

				ccc.SetDate(birthday)

				content.AddCell().Value = detail.ITEM_NAME
				content.AddCell().Value = ""
				content.AddCell().Value = "临床意义"
				price := fmt.Sprintf(`%.f`, detail.PRICE)
				if detail.XB == "0" {
					content.AddCell().Value = "*"
					content.AddCell().Value = price
					content.AddCell().Value = price
				} else if detail.XB == "1" {
					content.AddCell().Value = price
					content.AddCell().Value = "*"
					content.AddCell().Value = "*"
				} else {
					content.AddCell().Value = price
					content.AddCell().Value = price
					content.AddCell().Value = price
				}
			} else {
				for _, subItem := range detail.SUBITEM {
					content := sheet.AddRow()
					content.AddCell().Value = detail.ITEM_ID
					content.AddCell().Value = detail.ITEM_NAME
					content.AddCell().Value = subItem.CHECK_ITEM
					content.AddCell().Value = "临床意义"
					price := fmt.Sprintf(`.2f`, detail.PRICE)
					if detail.XB == "0" {
						content.AddCell().Value = "*"
						content.AddCell().Value = price
						content.AddCell().Value = price
					} else if detail.XB == "1" {
						content.AddCell().Value = price
						content.AddCell().Value = "*"
						content.AddCell().Value = "*"
					} else {
						content.AddCell().Value = fmt.Sprintf(`.f`, detail.PRICE)
						content.AddCell().Value = fmt.Sprintf(`.f`, detail.PRICE)
						content.AddCell().Value = fmt.Sprintf(`.f`, detail.PRICE)
					}
				}
			}

		}
	}
	//for _, sheet := range xlsxFile.Sheets {
	//	for _, col := range sheet.Cols {
	//		col.SetType(0)
	//
	//	}
	//}
	xlsxFile.Save("./b.xlsx")
	//err := archiver.Archive([]string{"./a.xlsx"}, "test.zip")
	//if err != nil {
	//	log.Println(err)
	//}
}
func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
func GuestExport() {
	var orderList []Order
	var order Order
	col := models.Mongo.DB("mnOracle").
		C("SALE_TJDW_LIST")
	col.Find(bson.M{}).All(&orderList)
	log.Println(len(orderList))
	xlsxFile := xlsx.NewFile()
	for _, v := range orderList {
		if v.MsjBILLCODE == "MSJ8S1594346536" {
			order = v
		}
	}
	for _, combo := range order.Combos {
		var sheet *xlsx.Sheet
		var err error
		sheet, err = xlsxFile.AddSheet(combo.GROUPNAME)
		if err != nil {
			log.Panic(err)
		}
		title := sheet.AddRow()
		title.AddCell().Value = "姓名"
		title.AddCell().Value = "性别"
		title.AddCell().Value = "出生日期"
		title.AddCell().Value = "身份证号码"
		title.AddCell().Value = "部门2"
		title.AddCell().Value = "套餐"
		for _, guest := range combo.Guests {
			content := sheet.AddRow()
			content.AddCell().Value = guest.XM
			content.AddCell().Value = guest.XB
			content.AddCell().Value = guest.CSRQ
			content.AddCell().Value = guest.SFZHM
			content.AddCell().Value = guest.K3_ITEM
			content.AddCell().Value = combo.ID
		}
	}
	xlsxFile.Save("./b.xlsx")
}

type Order struct {
	IsDeleted     bool    `bson:"isDeleted" json:"isDeleted"`         // 软删除标识
	ComboLocked   bool    `bson:"comboLocked" json:"comboLocked"`     // 订单套餐锁定
	MsjBILLCODE   string  `bson:"MsjBILLCODE" json:"MsjBILLCODE"`     //自定义订单号
	BILLCODE      string  `bson:"BILLCODE" json:"BILLCODE"`           //订单号 （自动生成）
	DWDM          string  `bson:"DWDM" json:"DWDM"`                   //单位代码 （前端传参查询）*
	DWMC          string  `bson:"DWMC" json:"DWMC"`                   //订单单位 （通过DWDM查询）
	YWYXM         string  `bson:"YWYXM" json:"YWYXM"`                 // 业务员姓名
	YWYDM         string  `bson:"YWYDM" json:"YWYDM"`                 //业务员代码(制单时bas_ent_info_pre带入,显示查询BAS_OPERATOR_INFOR)
	DJSJ          string  `bson:"DJSJ" json:"DJSJ"`                   //登记时间 （制单时系统时间）
	FKFS          string  `bson:"FKFS" json:"FKFS"`                   //付款方式：1 套餐挂账+加选项自付费；2 全自付费；3 全挂账；4 套餐用销售卡+加选项自付费；5 全用销售卡 （前端列表带参）*
	QBGDFS        string  `bson:"QBGDFS" json:"QBGDFS"`               //查询报告的方式：0 全部；1 仅单位；2 仅个人；3 全禁止 （前端列表带参）*
	XSYQ          string  `bson:"XSYQ" json:"XSYQ"`                   //备注一（人员和套餐的关系） （前端列表带参）*
	JGXX          string  `bson:"JGXX" json:"JGXX"`                   //备注二 （前端列表带参）*
	YWY_ZG        string  `bson:"YWY_ZG" json:"YWY_ZG"`               //审核：审核后为录入人员，未审核为空
	YWY_ZG_SHSJ   string  `bson:"YWY_ZG_SHSJ" json:"YWY_ZG_SHSJ"`     //审核：审核时间，未审核为空
	AuditStatus   int     `bson:"auditStatus" json:"auditStatus"`     // 码上检审核状态：0 未审核；1 审核中；2 审核通过； 3 审核拒绝
	STATUS        string  `bson:"STATUS" json:"STATUS"`               //审核：0未审核；1内勤已审核；2内勤作废；3财务退回；4财务审核；5财务作废
	ZJ_ZRR        string  `bson:"ZJ_ZRR" json:"ZJ_ZRR"`               //门店代码
	XSBDJR        string  `bson:"XSBDJR" json:"XSBDJR"`               //财务审核前为NULL，财务审核后为操作人员代码
	ZLPG_SHSJ     string  `bson:"ZLPG_SHSJ" json:"ZLPG_SHSJ"`         //财务审核前为NULL，财务审核后为财务审核时间
	ZJ_QZSJ       string  `bson:"ZJ_QZSJ" json:"ZJ_QZSJ"`             //NULL
	STATUS5       string  `bson:"STATUS5" json:"STATUS5"`             //未知，默认值1
	STATUS4       string  `bson:"STATUS4" json:"STATUS4"`             //财务审核前为0，财务审核后为6
	LCDJR         string  `bson:"LCDJR" json:"LCDJR"`                 //NULL
	ZXWCSJ        string  `bson:"ZXWCSJ" json:"ZXWCSJ"`               //NULL
	ZRHS          string  `bson:"ZRHS" json:"ZRHS"`                   //录入人员代码，eg：8S0004
	ZRHSName      string  `bson:"ZRHSName" json:"ZRHSName"`           //录入人员代码，(制单人)
	READ_ORNOT    string  `bson:"READ_ORNOT" json:"READ_ORNOT"`       //体检类别 Y入职 N年度 X优先出报告 F妇检 W外检 Z职业病 U下午体检 S个检 （前端列表带参）*
	EMAIL         string  `bson:"EMAIL" json:"EMAIL"`                 //凭证卡订单  Y是 N否 （前端列表带参）*
	WCBGRS        int     `bson:"WCBGRS" json:"WCBGRS"`               //加项折扣率 （前端列表带参）*
	WCTJRS        int     `bson:"WCTJRS" json:"WCTJRS"`               //预计体检金额 元 （前端列表带参）*
	LXR           string  `bson:"LXR" json:"LXR"`                     //预计体检人数 （前端列表带参）*
	TDHF          string  `bson:"TDHF" json:"TDHF"`                   //套餐变更提醒：N 不做任何提醒 ；Y 女性提醒更改为已婚套餐；1 全客人提醒 （前端列表带参）*
	ZRYS          string  `bson:"ZRYS" json:"ZRYS"`                   //结算类别  -1一次性结算  30月度结算 90季度结算 （前端列表带参）*
	SFJZSJ        string  `bson:"SFJZSJ" json:"SFJZSJ"`               //订单付款日期 2019/12/18 10:31:14 （前端列表带参）*
	ZXFW          string  `bson:"ZXFW" json:"ZXFW"`                   //报告领取方式 1送达单位 2自取 3邮寄 4网络查询 5给业务员 6个人邮寄(到付) 7单位邮寄(到付) （前端列表带参）*
	HTSJ          string  `bson:"HTSJ" json:"HTSJ"`                   //开始体检时间 2019/12/18 10:31:14 （前端列表带参）*
	ZLPG_SHRY     string  `bson:"ZLPG_SHRY" json:"ZLPG_SHRY"`         //全国单是否：为空；Y 是 ；N 否 （前端列表带参）*
	YWB_ZG        string  `bson:"YWB_ZG" json:"YWB_ZG"`               //NULL
	FreeBreakfast string  `bson:"freeBreakfast" json:"freeBreakfast"` //免费早餐  1 ：有  2：无
	Combos        []Combo `bson:"combos" json:"combos"`               // 套餐
}

type Combo struct {
	MsjBILLCODE         string  `bson:"MsjBILLCODE" json:"MsjBILLCODE"`                 //自定义订单号
	BILLCODE            string  `bson:"BILLCODE" json:"BILLCODE"`                       //订单id
	ID                  string  `bson:"ID" json:"ID"`                                   //套餐基本id
	GROUPNAME           string  `bson:"GROUPNAME" json:"GROUPNAME"`                     //套餐名称
	GROUPID             string  `bson:"GROUPID" json:"GROUPID"`                         //后台拼接
	TRANS_STATUS        string  `bson:"TRANS_STATUS" json:"TRANS_STATUS"`               //自选项否 N否 Y是 1检查阳性加项 2检验阳性加项
	STATUS              string  `bson:"STATUS" json:"STATUS"`                           //0编辑 1审核 2作废
	OP_DATETIME         string  `bson:"OP_DATETIME" json:"OP_DATETIME"`                 //性别  0:全部  1:男  2:女
	PRICE               float64 `bson:"PRICE" json:"PRICE"`                             //成交价
	OLDPRICE            float64 `bson:"OLDPRICE" json:"OLDPRICE"`                       //原价
	CZSJ                string  `bson:"CZSJ" json:"CZSJ"`                               //未知,默认值null
	YYSJ                string  `bson:"YYSJ" json:"YYSJ"`                               //未知,默认值null
	CostCoefficient     float64 `bson:"costCoefficient" json:"costCoefficient"`         //成本系数(套餐)
	CostCoefficientRate string  `bson:"costCoefficientRate" json:"costCoefficientRate"` //成本系数率

	Details []Detail         `bson:"details" json:"details"` // 套餐明细
	Guests  []BAS_CUST_INFOR `bson:"guests" json:"guests"`   // 人员明细
}

type Detail struct {
	MsjBILLCODE     string         `bson:"MsjBILLCODE" json:"MsjBILLCODE"`         //自定义订单号
	BILLCODE        string         `bson:"BILLCODE" json:"BILLCODE"`               //订单id
	ID              string         `bson:"ID" json:"ID"`                           //套餐基本id
	ITEM_ID         string         `bson:"ITEM_ID" json:"ITEM_ID"`                 //项目代码
	ITEM_NAME       string         `bson:"ITEM_NAME" json:"ITEM_NAME"`             //项目名称
	XB              string         `bson:"XB" json:"XB"`                           //性别 0女 1男 3未知
	ITEM_MS         string         `bson:"ITEM_MS" json:"ITEM_MS"`                 //null
	PRICE           float64        `bson:"PRICE" json:"PRICE"`                     //价格
	STATUS          string         `bson:"STATUS" json:"STATUS"`                   //0编辑 1审核 2作废
	TRANS_STATUS    string         `bson:"TRANS_STATUS" json:"TRANS_STATUS"`       //null
	OP_DATETIME     string         `bson:"OP_DATETIME" json:"OP_DATETIME"`         //null
	SXH             int            `bson:"SXH" json:"SXH"`                         //总数
	CostCoefficient float64        `bson:"costCoefficient" json:"costCoefficient"` //成本系数
	SUBITEM         []ITEM_SUBITEM `bson:"SUBITEM" json:"SUBITEM"`                 //子项信息
	Tag             bool           `bson:"-" json:"tag,omitempty"`                 //必选项目标记字段 是:true 否:false 注：数据库不存储该字段
}

type BAS_CUST_INFOR struct {
	MsjBILLCODE      string `bson:"MsjBILLCODE" json:"MsjBILLCODE"`           //自定义订单号
	MsjCID           string `bson:"MsjCID" json:"MsjCID"`                     //码上检健检号
	CID              string `bson:"CID" json:"CID"`                           //健检号
	XM               string `bson:"XM" json:"XM"`                             //姓名
	XB               string `bson:"XB" json:"XB"`                             //性别 0女 1男 3未知
	CSRQ             string `bson:"CSRQ" json:"CSRQ"`                         //出生日期
	SFZHM            string `bson:"SFZHM" json:"SFZHM"`                       //身份证号码
	MZ               string `bson:"MZ" json:"MZ"`                             //取 null
	WHCD             string `bson:"WHCD" json:"WHCD"`                         //未知 null
	GJ               string `bson:"GJ" json:"GJ"`                             //取 null
	HYZK             string `bson:"HYZK" json:"HYZK"`                         //婚姻状态 一半0001,一半null,取
	JTZZ             string `bson:"JTZZ" json:"JTZZ"`                         //家庭住址 取null
	YZBM             string `bson:"YZBM" json:"YZBM"`                         //邮政编码 取null
	LXDH             string `bson:"LXDH" json:"LXDH"`                         //联系电话 有则传,没有为null
	YDDH             string `bson:"YDDH" json:"YDDH"`                         //移动电话 有则传,没有为null
	DWDM             string `bson:"DWDM" json:"DWDM"`                         //单位代码 后台赋值
	DWMC             string `bson:"DWMC" json:"DWMC"`                         //单位名称 可为null 后面赋值
	EMAIL            string `bson:"EMAIL" json:"EMAIL"`                       //取 null
	BZ               string `bson:"BZ" json:"BZ"`                             //一半Y,一半null,可取null
	DJSJ             string `bson:"DJSJ" json:"DJSJ"`                         //到检时间,没填则为当前时间
	DJR              string `bson:"DJR" json:"DJR"`                           //操作人员代码? eg:8S0004
	PROVINCE         string `bson:"PROVINCE" json:"PROVINCE"`                 //省份 取 null
	CITY             string `bson:"CITY" json:"CITY"`                         //预约门店 eg:8S
	REGION           string `bson:"REGION" json:"REGION"`                     //地区 取 null
	CUST_TYPE        string `bson:"CUST_TYPE" json:"CUST_TYPE"`               //取 null
	ZWDM             string `bson:"ZWDM" json:"ZWDM"`                         //002/004/NQ2600011739 可取null?
	CW               string `bson:"CW" json:"CW"`                             //套餐分组 eg:0002 没有则为0000
	ZYCD             string `bson:"ZYCD" json:"ZYCD"`                         //取 null
	HYCD             string `bson:"HYCD" json:"HYCD"`                         //未知 一半9一半6
	KHLY             string `bson:"KHLY" json:"KHLY"`                         //未知 eg:01,4408436/10,419994/60,5426/null,2225653,可取 01
	ZHYE             string `bson:"ZHYE" json:"ZHYE"`                         //取 0
	MEMBER_TYPE      string `bson:"MEMBER_TYPE" json:"MEMBER_TYPE"`           //一半00/一半null,取null
	CARD_CODE        string `bson:"CARD_CODE" json:"CARD_CODE"`               //已回收/null
	MEMBER_CARD_CODE string `bson:"MEMBER_CARD_CODE" json:"MEMBER_CARD_CODE"` //取 null
	TRANS_STATUS     string `bson:"TRANS_STATUS" json:"TRANS_STATUS"`         //取 9
	OP_DATE          string `bson:"OP_DATE" json:"OP_DATE"`                   //系统时间
	YWYDM            string `bson:"YWYDM" json:"YWYDM"`                       //业务员代码,eg:1172
	GSGDDH           string `bson:"GSGDDH" json:"GSGDDH"`                     //dj/null
	KHTZ             string `bson:"KHTZ" json:"KHTZ"`                         //各种奇怪的东西都有,建议取null
	GSTZ             string `bson:"GSTZ" json:"GSTZ"`                         //公司**,建议取null
	TXTZ             string `bson:"TXTZ" json:"TXTZ"`                         //取 null
	YXJB             string `bson:"YXJB" json:"YXJB"`                         //未知 eg:0/null
	CUST_REGION      string `bson:"CUST_REGION" json:"CUST_REGION"`           //价格,eg:1/null
	K3_ITEM          string `bson:"K3_ITEM" json:"K3_ITEM"`                   //部门2(客户单位部门) 据说必填
	K3_ZG_ITEM       string `bson:"K3_ZG_ITEM" json:"K3_ZG_ITEM"`             //特殊标注(套餐名称) eg:POSTMAN01
	BGD_LB           string `bson:"BGD_LB" json:"BGD_LB"`                     //取 0
	GRGZ_YE          string `bson:"GRGZ_YE" json:"GRGZ_YE"`                   //取 0
	ZZYS             string `bson:"ZZYS" json:"ZZYS"`                         //取 null
	XGSJ             string `bson:"XGSJ" json:"XGSJ"`                         //没有则为当前时间
	USER_PASSWORD    string `bson:"USER_PASSWORD" json:"USER_PASSWORD"`       //cid密码 随机六位数 eg:394702
	PASSWORD_XGSJ    string `bson:"PASSWORD_XGSJ" json:"PASSWORD_XGSJ"`       //cid密码修改时间 建议:1900-01-01 00:00:00 null
	ZZHS             string `bson:"ZZHS" json:"ZZHS"`                         //门店编号 eg:8S
}

type ITEM_SUBITEM struct {
	CHECK_ITEM string `bson:"CHECK_ITEM" json:"CHECK_ITEM"` //检查项
	CHECK_ID   string `bson:"CHECK_ID" json:"CHECK_ID"`     //检查项代码
}
