// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Links is the golang structure of table links for DAO operations like Where/Data.
type Links struct {
	g.Meta    `orm:"table:links, do:true"`
	Id        interface{} //
	Name      interface{} //
	Url       interface{} //
	Sequence  interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
