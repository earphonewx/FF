package middleware

import (
	"ff/g"
	"fmt"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a, _ := gormadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)",
		g.VP.GetString("mysql.user"), g.VP.GetString("mysql.password"), g.VP.GetString("mysql.host"), g.VP.GetInt("mysql.port")))
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy(...)
	// e.RemovePolicy(...)

	// Save the policy back to DB.
	e.SavePolicy()
}
