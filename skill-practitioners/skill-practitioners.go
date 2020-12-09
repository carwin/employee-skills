package skillpractitioners

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"github.com/carwin/employee-skills/bambooapi"
	employeeskillset "github.com/carwin/employee-skills/employee-skillset"
	skilldistribution "github.com/carwin/employee-skills/skill-distribution"
)

type employee struct {
	name string
}

type employeeList struct {
	names []employee
}

func createEmployeeList(t skill_distribution.SkillCountTable, s string) employeeList {
	var employeeList employeeList

	for _, v := range t.Row {
		var employee employee
		var employeeId = v.Id
		if employeeHasSkill(v, employeeId, s) {
			var employeeName = getEmployeeName(employeeId)
			employee.name = employeeName
			employeeList.names = append(employeeList.names, employee)
		}
	}

	return employeeList
}

func getEmployeeName(i int) string {

	// Replace {org} with your organization name.
	url := bambooapi.GetBambooAPIURL() + "api/gateway.php/{org}/v1/employees/" + strconv.Itoa(i) + "/?fields=firstName,lastName"
	//var data = getAPIData(url, "application/json")
	var data employeeskillset.DirectoryEmployee
	err := json.Unmarshal([]byte(bambooapi.GetAPIData(url, "application/json")), &data)

	if err != nil {
		fmt.Print("Errored unmarshalling data from directory \n", err)
		os.Exit(1)
	}
	return data.FirstName + " " + data.LastName
}

func employeeHasSkill(r skilldistribution.SkillCountRow, i int, s string) bool {
	if i == r.Id {
		if r.Field[0] == s {
			return true
		}
	}

	return false
}

// GetSkillPractitioners - Find which employees have the given skill.
func GetSkillPractitioners(skillName string) {
	// Replace {org} with your organization name.
	path := "api/gateway.php/{org}/v1/employees/all/tables/customSkills"
	url := bambooapi.GetBambooAPIURL() + path
	var data skilldistribution.SkillCountTable
	err := xml.Unmarshal([]byte(bamboopi.GetAPIData(url, "application/xml")), &data)
	if err != nil {
		fmt.Print("Errored unmarshaling skill %s\n", err)
		os.Exit(1)
	}

	var skillList = createEmployeeList(data, skillName)
	fmt.Println("People who know " + skillName)
	fmt.Println("-----------------------")
	for _, v := range skillList.names {
		fmt.Println(v.name)
	}
}
