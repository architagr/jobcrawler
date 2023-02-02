package urlseeding

type HostName string
type JobType string
type JobModel string
type Role string
type JobTitle string
type ExperienceLevel string

const (
	HostName_Linkedin HostName = "www.linkedin.com"
	HostName_Indeed   HostName = "www.indeed.com"
)
const (
	ExperienceLevel_EntryLevel     = "Entry Level"
	ExperienceLevel_Internship     = "Internship"
	ExperienceLevel_Associate      = "Associate"
	ExperienceLevel_MidSeniorLevel = "Mid-Senior Level"
	ExperienceLevel_Director       = "Director"
)

const (
	JobType_FullTime   JobType = "Full Time"
	JobType_PartTime   JobType = "Part Time"
	JobType_Contract   JobType = "Contract"
	JobType_Temporary  JobType = "Temporary"
	JobType_Volunteer  JobType = "Volunteer"
	JobType_Internship JobType = "Internship"
	JobType_Other      JobType = "Other"
)

const (
	JobModel_OnSite JobModel = "On site"
	JobModel_Remote JobModel = "Remote"
	JobModel_Hybrid JobModel = "Hybrid"
)

const (
	Role_Manufacturing         Role = "Manufacturing"
	Role_Consulting            Role = "Consulting"
	Role_Finance               Role = "Finance"
	Role_Management            Role = "Management"
	Role_BusinessDevelopment   Role = "Business Development"
	Role_Marketing             Role = "Marketing"
	Role_ProjectManagement     Role = "Project Management"
	Role_HealthCareProvider    Role = "Health Care Provider"
	Role_Design                Role = "Design"
	Role_Sales                 Role = "Sales"
	Role_Engineering           Role = "Engineering"
	Role_InformationTechnology Role = "Information Technology"
	Role_Others                Role = "Others"
)

const (
	JobTitle_HumanResourcesIntern            JobTitle = "Human Resources Intern"
	JobTitle_ElectricalEngineer              JobTitle = "Electrical Engineer"
	JobTitle_EditorialStaff                  JobTitle = "Editorial Staff"
	JobTitle_InformationTechnologyArchitect  JobTitle = "Information Technology Architect"
	JobTitle_InformationTechnologyConsultant JobTitle = "Information Technology Consultant"
	JobTitle_Salesperson                     JobTitle = "Salesperson"
	JobTitle_SalesExecutive                  JobTitle = "Sales Executive"
	JobTitle_SoftwareEngineer                JobTitle = "Software Engineer"
	JobTitle_AdministrativeAssistant         JobTitle = "Administrative Assistant"
	JobTitle_CloudEngineer                   JobTitle = "Cloud Engineer"
	JobTitle_BusinessDevelopmentManager      JobTitle = "Business Development Manager"
	JobTitle_Reviewer                        JobTitle = "Reviewer"
	JobTitle_SalesManager                    JobTitle = "Sales Manager"
	JobTitle_MarketingManager                JobTitle = "Marketing Manager"
	JobTitle_JavaSoftwareEngineer            JobTitle = "Java Software Engineer"
	JobTitle_SeniorSoftwareEngineer          JobTitle = "Senior Software Engineer"
	JobTitle_FullStackEngineer               JobTitle = "Full Stack Engineer"
)
