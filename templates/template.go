package templates

type Template struct {
	Framework        string `json:"framework"`
	BaseImage        string `json:"base_image"`
	WorkDir          string `json:"work_dir"`
	RequirementsFile string `json:"requirements_file"`
	RunCommand       string `json:"run_command"`
}

type Templates struct {
	Templates map[string]Template `json:"templates"`
}
