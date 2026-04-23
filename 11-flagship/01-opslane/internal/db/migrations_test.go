package db

import (
	"strings"
	"testing"
)

func TestSchemaStatementsCoverCoreOpslaneTables(t *testing.T) {
	t.Parallel()

	wantSnippets := []string{
		"CREATE TABLE IF NOT EXISTS tenants",
		"CREATE TABLE IF NOT EXISTS users",
		"CREATE TABLE IF NOT EXISTS orders",
		"CREATE TABLE IF NOT EXISTS payments",
		"idx_orders_tenant_status",
		"idx_payments_tenant_order",
	}

	joined := ""
	for _, statement := range schemaStatements {
		joined += statement + "\n"
	}

	for _, snippet := range wantSnippets {
		if !strings.Contains(joined, snippet) {
			t.Fatalf("schema statements missing %q", snippet)
		}
	}
}
