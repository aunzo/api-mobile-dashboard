package models

// BuildInfo represents the build information data structure
type BuildInfo struct {
	StartTime        string `mapstructure:"start_time" json:"start_time" firestore:"start_time"`
	EndTime          string `mapstructure:"end_time" json:"end_time" firestore:"end_time"`
	Duration         string `mapstructure:"duration" json:"duration" firestore:"duration"`
	GitBranch        string `mapstructure:"git_branch" json:"git_branch" firestore:"git_branch"`
	GitAuthor        string `mapstructure:"git_author" json:"git_author" firestore:"git_author"`
	Scheme           string `mapstructure:"scheme" json:"scheme" firestore:"scheme"`
	MachineModel     string `mapstructure:"machine_model" json:"machine_model" firestore:"machine_model"`
	Platform         string `mapstructure:"platform" json:"platform" firestore:"platform"`
	CPU              string `mapstructure:"cpu" json:"cpu" firestore:"cpu"`
	MemoryGB         string `mapstructure:"memory_gb" json:"memory_gb" firestore:"memory_gb"`
	DiskTotal        string `mapstructure:"disk_total" json:"disk_total" firestore:"disk_total"`
	DiskAvailable    string `mapstructure:"disk_available" json:"disk_available" firestore:"disk_available"`
	FileChangeCount  string `mapstructure:"file_change_count" json:"file_change_count" firestore:"file_change_count"`
	CompileFileCount string `mapstructure:"compile_file_count" json:"compile_file_count" firestore:"compile_file_count"`
}
