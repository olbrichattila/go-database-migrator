package migrate

import (
	"fmt"
	"strings"
)

type firebirdMigrationTableSQLProvider struct {
	tablePrefix string
}

func (p *firebirdMigrationTableSQLProvider) createMigrationSQL() string {
	upperCasePrefix := strings.ToUpper(p.tablePrefix)
	sql := `EXECUTE BLOCK AS BEGIN
		if (not exists(select 1 from rdb$relations where rdb$relation_name = '%s_MIGRATIONS')) then
		execute statement 'CREATE TABLE %s_MIGRATIONS (
			file_name VARCHAR(255),
			created_at VARCHAR(35),
			deleted_at TIMESTAMP)
			checksum CHAR(32);';
		END`

	return fmt.Sprintf(sql, upperCasePrefix, upperCasePrefix)
}

func (p *firebirdMigrationTableSQLProvider) createReportSQL() string {
	upperCasePrefix := strings.ToUpper(p.tablePrefix)
	sql := `EXECUTE BLOCK AS BEGIN
		if (not exists(select 1 from rdb$relations where rdb$relation_name = '%s_MIGRATION_REPORTS')) then
		execute statement 'CREATE TABLE %s_MIGRATION_REPORTS (
			file_name VARCHAR(255),
			result_status VARCHAR(12),
			created_at VARCHAR(35),
			message BLOB SUB_TYPE TEXT);';
		END`

	return fmt.Sprintf(sql, upperCasePrefix, upperCasePrefix)
}
