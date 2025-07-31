package models

// BuildInfo represents the build information data structure
type BuildInfo struct {
	StartTime        string `mapstructure:"start_time" json:"start_time" db:"start_time"`
	EndTime          string `mapstructure:"end_time" json:"end_time" db:"end_time"`
	Duration         string `mapstructure:"duration" json:"duration" db:"duration"`
	GitBranch        string `mapstructure:"git_branch" json:"git_branch" db:"git_branch"`
	GitAuthor        string `mapstructure:"git_author" json:"git_author" db:"git_author"`
	Scheme           string `mapstructure:"scheme" json:"scheme" db:"scheme"`
	MachineModel     string `mapstructure:"machine_model" json:"machine_model" db:"machine_model"`
	Platform         string `mapstructure:"platform" json:"platform" db:"platform"`
	CPU              string `mapstructure:"cpu" json:"cpu" db:"cpu"`
	MemoryGB         string `mapstructure:"memory_gb" json:"memory_gb" db:"memory_gb"`
	DiskTotal        string `mapstructure:"disk_total" json:"disk_total" db:"disk_total"`
	DiskAvailable    string `mapstructure:"disk_available" json:"disk_available" db:"disk_available"`
	FileChangeCount  string `mapstructure:"file_change_count" json:"file_change_count" db:"file_change_count"`
	CompileFileCount string `mapstructure:"compile_file_count" json:"compile_file_count" db:"compile_file_count"`
}
