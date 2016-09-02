package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	//"github.com/mrtomyum/nava-stock/model"
	//"strconv"
	"testing"
	"time"
	"github.com/mrtomyum/nava-stock/model"
)

const (
	MSSQL_SERVER   string = "olaf.nopadol.com"
	MSSQL_DATABASE string = "navatest"
	MSSQL_USER     string = "sa"
	MSSQL_PASS     string = "[ibdkifu"
)

var (
	myDB *sqlx.DB
	msDB *sqlx.DB
)

type Client struct {
	Host string
}

func init() {
	var err error
	myDSN := DB_USER + ":" + DB_PASS + "@" + DB_HOST + "/" + DB_NAME + "?parseTime=true"
	myDB, err = NewDB("mysql", myDSN)
	if err != nil {
		log.Panic("NewDB() Error:", err)
	}
	msDSN := "server=" + MSSQL_SERVER + ";database=" + MSSQL_DATABASE + ";user id=" + MSSQL_USER + ";password=" + MSSQL_PASS
	//msDSN := "server=olaf.nopadol.com; database= navatest ;user id= sa ;password= [ibdkifu"
	msDB, err = NewDB("mssql", msDSN)
	if err != nil {
		log.Panic("NewDB() Error:", err)
	}

}

type tb_sv_machine struct {
	ID             uint64      `json:"id"`
	MachineID      uint64      `json:"machine_id" db:"machineid"`
	Code           null.String `json:"code"`
	NameTH         null.String `json:"name_th" db:"nameth"`
	NameEN         null.String `json:"name_en" db:"nameen"`
	MachineTypeID  null.Int    `json:"machine_type_id" db:"machinetypeid"`
	TypeCode       null.String `json:"type_code" db:"typecode"`
	MachinePlaceID null.Int    `json:"machine_place_id" db:"machineplaceid"`
	PlaceCode      null.String `json:"place_code" db:"placecode"`
	MyDesc         null.String `json:"my_desc" db:"mydesc"`
	BranchID       uint64      `json:"branch_id" db:"branchid"`
	BranchCode     null.String `json:"branch_code" db:"branchcode"`
	ActiveStatus   null.Int    `json:"active_status" db:"activestatus"`
	CreatorCode    null.String `json:"creator_code" db:"creatorcode"`
	CreateDatetime *time.Time  `json:"create_datetime" db:"createdatetime"`
	EditorCode     null.String `json:"editor_code" db:"editorcode"`
	EditDatetime   *time.Time  `json:"edit_datetime" db:"editdatetime"`
	ColumnCount    null.Int    `json:"column_count" db:"columncount"`
	SaleCode       null.String `json:"sale_code" db:"salecode"`
	Target         []uint8     `json:"target" db:"target"`
}

//func TestGetTableMachine(t *testing.T) {
//	sql := `SELECT * FROM tb_sv_machine`
//	var machines []tb_sv_machine
//	err := msDB.Select(&machines, sql)
//	if err != nil {
//		t.Error(err)
//	}
//	log.Info(machines)
//}

//func TestImportTable_Machine(t *testing.T) {
//	sql := `SELECT * FROM tb_sv_machine WHERE typecode <> '' ORDER BY code`
//	var fromMachines []tb_sv_machine
//	err := msDB.Select(&fromMachines, sql)
//	if err != nil {
//		t.Error(err)
//	}
//	var count int
//	for _, m := range fromMachines {
//		code := m.Code.String
//		placecode := m.PlaceCode.String
//		var c model.Client
//		_ = myDB.Get(&c, `SELECT id FROM client WHERE code =?`, placecode)
//		clientID := c.ID
//		sql = `INSERT INTO machine(code, client_id) VALUES(?,?)`
//		_, err := myDB.Exec(sql, code, clientID)
//		if err != nil {
//			log.Error(err.Error())
//		}
//		count++
//		log.Info("Import machine row number = ", count)
//	}
//}

type BCItem struct {
	Code  null.String `db:"code"`
	Name1 null.String
	Name2 null.String
	//DefStkUnitCode string
	//DefSaleUnitCode string
	//DefBuyUnitCode string
	//BrandCode string
	//TypeCode string
}

//func TestImportTable_Item(t *testing.T) {
//	log.Info(">>>1")
//	sql := `SELECT
//	 	code,
//	 	name1,
//	 	name2
//	 FROM BCItem ORDER BY code`
//	var items []BCItem
//	log.Info("var items")
//	err := msDB.Select(&items, sql)
//	if err != nil {
//		t.Error("Error on Select():", err)
//	}
//	log.Info("no err on msDB.Select()")
//	var count int
//	sql = `INSERT INTO item(
//			sku,
//			nameTH,
//			nameEN
//		) VALUES(?,?,?)`
//	stmt, err := myDB.Preparex(sql)
//	if err != nil {
//		t.Error("Error on DBPreparex()=", err)
//	}
//
//	for _, i := range items {
//		_, err := stmt.Exec(
//			i.Code,
//			i.Name1,
//			i.Name2,
//		)
//		if err != nil {
//			log.Error(err.Error())
//		}
//		count++
//		log.Info("Import Item row number = ", count)
//	}
//}

type MachineColumn struct {
	MachineCode string
	Code        string
	ItemCode    string
}

//func TestImportTable_MachineColumn(t *testing.T) {
//	sql := `SELECT
//	 	machinecode,
//	 	code,
//	 	itemcode
//	 FROM tb_sv_machineshelf ORDER BY machinecode ASC, code ASC`
//	var mcs []MachineColumn
//	log.Info("var mc")
//	err := msDB.Select(&mcs, sql)
//	if err != nil {
//		t.Error("Error on Select():", err)
//	}
//	log.Info("no err on msDB.Select()")
//	var count int
//	sql = `INSERT INTO machine_column(
//			machine_id,
//			column_no,
//			item_id
//		) VALUES(?,?,?)`
//	stmt, err := myDB.Preparex(sql)
//	if err != nil {
//		t.Error("Error on DBPreparex()=", err)
//	}
//	var (
//		m model.Machine
//		i model.Item
//	)
//	for _, mc := range mcs {
//
//		myDB.Get(&m, `SELECT id FROM machine WHERE code = ?`, mc.MachineCode)
//		MachineID := m.ID
//		myDB.Get(&i, `SELECT id FROM item WHERE sku = ?`, mc.ItemCode)
//		ItemID := i.ID
//		ColumnNo, err := strconv.ParseInt(mc.Code, 10, 64)
//		if err != nil {
//			log.Error(err.Error())
//		}
//		_, err = stmt.Exec(
//			MachineID,
//			ColumnNo,
//			ItemID,
//		)
//		if err != nil {
//			log.Error(err.Error())
//		}
//		count++
//		log.Info("Import MachineColumn row number = ", count)
//	}
//}

type MachinePlace struct {
	Code string
	NameTH null.String
	NameEN null.String
}

//func TestImportTable_Client(t *testing.T) {
//	sql := `SELECT
//	 	code,
//	 	nameth,
//	 	nameen
//	 FROM tb_sv_machineplace ORDER BY code ASC`
//	var ps []MachinePlace
//	err := msDB.Select(&ps, sql)
//	if err != nil {
//		t.Error("Error on Select():", err)
//	}
//	log.Info("no err on msDB.Select()")
//	var count int
//	sql = `INSERT INTO client(
//			code,
//			nameTH,
//			nameEN
//		) VALUES(?,?,?)`
//	stmt, err := myDB.Preparex(sql)
//	if err != nil {
//		t.Error("Error on DBPreparex()=", err)
//	}
//	for _, p := range ps {
//		_, err = stmt.Exec(
//			p.Code,
//			p.NameTH.String,
//			p.NameEN.String,
//		)
//		if err != nil {
//			log.Error(err.Error())
//		}
//		count++
//		log.Info("Import MachinePlace -> Client, row number = ", count)
//	}
//}
//
