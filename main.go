package main

import (
	"github.com/casbin/casbin"
	"github.com/casbin/xorm-adapter"
	_ "github.com/lib/pq"
	"fmt"
)

func main() {
	// Initialize a Xorm adapter and use it in a Casbin enforcer:
	// The adapter will use the Postgres database named "casbin".
	// If it doesn't exist, the adapter will create it automatically.
	a := xormadapter.NewAdapter("postgres", "user=postgres password=postgres host=postgresql.k2tf.marathon.mesos port=5432 sslmode=disable") // Your driver and data source.

	// Or you can use an existing DB "abc" like this:
	// The adapter will use the table named "casbin_rule".
	// If it doesn't exist, the adapter will create it automatically.
	// a := xormadapter.NewAdapter("postgres", "dbname=abc user=postgres_username password=postgres_password host=127.0.0.1 port=5432 sslmode=disable", true)

	e := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	e.Enforce("alice", "data1", "read")

	// Modify the policy.
	// e.AddPolicy("alice","data1","read","allow")
	// e.RemovePolicy(...)
	e.Enforce("alice", "data1", "read")
	rs := e.GetAllRoles()
	fmt.Sprint(rs)
	// Save the policy back to DB.
	e.SavePolicy()
}