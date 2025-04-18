package migrations

import "gofr.dev/pkg/gofr/migration"

const createResumesTable = `CREATE TABLE IF NOT EXISTS resumes (
    id VARCHAR(36) NOT NULL PRIMARY KEY,
    role VARCHAR(100) NOT NULL,
    company VARCHAR(100) NOT NULL,
    version VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    notes TEXT,
    file_name VARCHAR(255),
    file_content LONGBLOB,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_role (role),
    INDEX idx_company (company),
    INDEX idx_status (status)
);`

func createTableResumes() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createResumesTable)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
