package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aunz/api-mobile-dashboard-golang/internal/models"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

func NewPostgresDB() (*PostgresDB, error) {
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "password")
	dbname := getEnv("DB_NAME", "mobile_dashboard")
	sslmode := getEnv("DB_SSLMODE", "disable")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL database")

	postgresDB := &PostgresDB{DB: db}
	
	// Create tables if they don't exist
	if err := postgresDB.createTables(); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return postgresDB, nil
}

func (p *PostgresDB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS build_infos (
		id SERIAL PRIMARY KEY,
		start_time VARCHAR(255),
		end_time VARCHAR(255),
		duration VARCHAR(255),
		git_branch VARCHAR(255),
		git_author VARCHAR(255),
		scheme VARCHAR(255),
		machine_model VARCHAR(255),
		platform VARCHAR(255),
		cpu VARCHAR(255),
		memory_gb VARCHAR(255),
		disk_total VARCHAR(255),
		disk_available VARCHAR(255),
		file_change_count VARCHAR(255),
		compile_file_count VARCHAR(255),
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := p.DB.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create build_infos table: %w", err)
	}

	log.Println("Database tables created successfully")
	return nil
}

func (p *PostgresDB) GetAllBuildInfos() ([]models.BuildInfo, error) {
	query := `
		SELECT start_time, end_time, duration, git_branch, git_author, 
		       scheme, machine_model, platform, cpu, memory_gb, 
		       disk_total, disk_available, file_change_count, compile_file_count
		FROM build_infos 
		ORDER BY created_at DESC`

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to query build infos: %w", err)
	}
	defer rows.Close()

	var buildInfos []models.BuildInfo
	for rows.Next() {
		var buildInfo models.BuildInfo
		err := rows.Scan(
			&buildInfo.StartTime,
			&buildInfo.EndTime,
			&buildInfo.Duration,
			&buildInfo.GitBranch,
			&buildInfo.GitAuthor,
			&buildInfo.Scheme,
			&buildInfo.MachineModel,
			&buildInfo.Platform,
			&buildInfo.CPU,
			&buildInfo.MemoryGB,
			&buildInfo.DiskTotal,
			&buildInfo.DiskAvailable,
			&buildInfo.FileChangeCount,
			&buildInfo.CompileFileCount,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan build info: %w", err)
		}
		buildInfos = append(buildInfos, buildInfo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return buildInfos, nil
}

func (p *PostgresDB) CreateBuildInfo(buildInfo models.BuildInfo) error {
	query := `
		INSERT INTO build_infos (
			start_time, end_time, duration, git_branch, git_author,
			scheme, machine_model, platform, cpu, memory_gb,
			disk_total, disk_available, file_change_count, compile_file_count
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`

	_, err := p.DB.Exec(query,
		buildInfo.StartTime,
		buildInfo.EndTime,
		buildInfo.Duration,
		buildInfo.GitBranch,
		buildInfo.GitAuthor,
		buildInfo.Scheme,
		buildInfo.MachineModel,
		buildInfo.Platform,
		buildInfo.CPU,
		buildInfo.MemoryGB,
		buildInfo.DiskTotal,
		buildInfo.DiskAvailable,
		buildInfo.FileChangeCount,
		buildInfo.CompileFileCount,
	)

	if err != nil {
		return fmt.Errorf("failed to insert build info: %w", err)
	}

	return nil
}

func (p *PostgresDB) Close() error {
	return p.DB.Close()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}