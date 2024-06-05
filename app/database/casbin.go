package database

import (
	"fmt"
	"github.com/casbin/casbin/v2"

	gormadapter "github.com/casbin/gorm-adapter/v3"
)

func Casbin() *casbin.Enforcer {
	// Initialize  casbin adapter
	adapter, err := gormadapter.NewAdapterByDBUseTableName(DB, "user", "rules")
	if err != nil {
		panic(fmt.Sprintf("failed to initialize casbin adapter: %v", err))
	}

	// Load model configuration file and policy store adapter
	e, err := casbin.NewEnforcer("app/config/restful_rbac_model.conf", adapter)
	if err != nil {
		panic(fmt.Sprintf("failed to create casbin enforcer: %v", err))
	}

	// Add policy - One-time run
	if hasPolicy, _ := e.HasPolicy("R_SUPER", "/api/*", "(GET)|(POST)|(PUT)|(DELETE)"); !hasPolicy {
		_, err := e.AddPolicy("R_SUPER", "/api/*", "(GET)|(POST)|(PUT)|(DELETE)")
		if err != nil {
			return nil
		}
	}

	err = e.LoadPolicy()
	if err != nil {
		return nil
	}
	return e
}
