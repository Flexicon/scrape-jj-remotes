package model

type Job struct {
	ID              string     `json:"id"`
	Title           string     `json:"title"`
	City            string     `json:"city"`
	CountryCode     string     `json:"country_code"`
	Remote          bool       `json:"remote"`
	CompanyName     string     `json:"company_name"`
	CompanyLogoUrl  string     `json:"company_logo_url"`
	ExperienceLevel string     `json:"experience_level"`
	SalaryFrom      int        `json:"salary_from"`
	SalaryTo        int        `json:"salary_to"`
	SalaryCurrency  string     `json:"salary_currency"`
	EmploymentType  string     `json:"employment_type"`
	PublishedAtISO  string     `json:"published_at"`
	Skills          []JobSkill `json:"skills"`
}

type JobSkill struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}
