package employee_skillset

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/carwin/employee-skills/bambooAPI"
)

type Directory struct {
	Employees []DirectoryEmployee `json:"employees"`
}

type DirectoryEmployee struct {
	ID          string
	FirstName   string
	DisplayName string
	LastName    string
}
type EmployeeSkillsTable struct {
	Rows []EmployeeSkillsRow `xml:"row"`
}

type EmployeeSkillsRow struct {
	EmployeeId int      `xml:"employeeId,attr"`
	Field      []string `xml:"field"`
}
type employeeSkillList struct {
	skills []employeeSkill
}

type employeeSkill struct {
	Skill         string
	SkillLevel    string
	Certified     string
	Certification string
}

func getIdFromDirectory(d Directory, e string) int {
	var i int
	for _, v := range d.Employees {
		if v.FirstName+" "+v.LastName == e {
			i, _ = strconv.Atoi(v.ID)
		}
	}
	return i
}

func createEmployeeSkillsList(t EmployeeSkillsTable) employeeSkillList {
	var employeeSkillList employeeSkillList

	for _, v := range t.Rows {
		var skill employeeSkill
		skill.Skill = v.Field[0]
		skill.SkillLevel = v.Field[1]
		skill.Certification = v.Field[2]
		skill.Certified = v.Field[3]
		employeeSkillList.skills = append(employeeSkillList.skills, skill)
	}

	return employeeSkillList

}

func GetEmployeeSkillset(employee string) {
	//var employee string = os.Args[2]
	path := "api/gateway.php/mobomo/v1/employees/directory"
	url := bambooAPI.GetBambooAPIURL() + path
	var data Directory
	err := json.Unmarshal([]byte(bambooAPI.GetAPIData(url, "application/json")), &data)

	if err != nil {
		fmt.Print("Errored unmarshalling data from directory \n", err)
		os.Exit(1)
	}

	var employeeID = getIdFromDirectory(data, employee)

	path2 := "api/gateway.php/mobomo/v1/employees/" + strconv.Itoa(employeeID) + "/tables/customSkills"
	url2 := bambooAPI.GetBambooAPIURL() + path2
	var data2 EmployeeSkillsTable
	err2 := xml.Unmarshal([]byte(bambooAPI.GetAPIData(url2, "application/xml")), &data2)

	if err2 != nil {
		fmt.Print("Errored unmarshalling data getting an Employee \n", err2)
		os.Exit(1)
	}

	var employeeSkillList = createEmployeeSkillsList(data2)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', tabwriter.AlignRight)
	defer w.Flush()
	fmt.Fprintf(w, "\n %s\t%s\t", employee+"'s Skills", "Skill Certification")
	fmt.Fprintf(w, "\n %s\t%s\t\n", "-------------", "------------------------")
	for _, v := range employeeSkillList.skills {
		fmt.Fprintf(w, "%s\t%s\t\n", v.Skill, strings.Replace(v.Certification, "\n", ", ", -1))
	}
}
