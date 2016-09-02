package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/guregu/null"
	"github.com/jmoiron/sqlx"
	"testing"
	"time"
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

func TestImportTable_Machine(t *testing.T) {
	sql := `SELECT * FROM tb_sv_machine ORDER BY code`
	var fromMachines []tb_sv_machine
	err := msDB.Select(&fromMachines, sql)
	if err != nil {
		t.Error(err)
	}
	var count int
	for _, from := range fromMachines {
		sql = `INSERT INTO machine(code) VALUES(?)`
		code := from.Code.String
		_, err := myDB.Exec(sql, code)
		if err != nil {
			log.Error(err.Error())
		}
		count++
		log.Info("Import machine row number = ", count)
	}
}




