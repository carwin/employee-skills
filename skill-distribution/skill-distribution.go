package skill_distribution

import (
	"encoding/xml"
	"fmt"
	"os"
	"text/tabwriter"

	bambooAPI "github.com/carwin/employee-skills/bambooAPI"
)

// Structs
// ----------------------------------------------------------------------------
type SkillCountTable struct {
	Row []SkillCountRow `xml:"row"`
}
type SkillCountRow struct {
	Id    int      `xml:"employeeId,attr"`
	Field []string `xml:"field"`
}

type skillList struct {
	skills []skill
}

type skill struct {
	name          string
	employeeCount int
}

// Functions
// ----------------------------------------------------------------------------
// Count the instances of a particular skill given a SkillCountTable
func countSkills(t SkillCountTable, s string) (i int) {
	for _, v := range t.Row {
		if v.Field[0] == s {
			i++
		}
	}
	return
}

// Find out if a particular string already exists in a given SkillList
func skillInSkillList(s skillList, n string) bool {
	var output bool
	for _, v := range s.skills {
		if v.name == n {
			return true
		}
	}
	return output
}

// Create a SkillList from a given SkillCountTable.
func createSkillList(t SkillCountTable) skillList {
	var skillList skillList

	for _, v := range t.Row {
		var skill skill
		skill.name = v.Field[0]
		skill.employeeCount = countSkills(t, v.Field[0])
		if skillInSkillList(skillList, v.Field[0]) == false {
			skillList.skills = append(skillList.skills, skill)
		}
	}
	return skillList
}

func GetSkillDistribution() {
	path := "api/gateway.php/mobomo/v1/employees/all/tables/customSkills"
	url := bambooAPI.GetBambooAPIURL() + path
	var data SkillCountTable
	err := xml.Unmarshal([]byte(bambooAPI.GetAPIData(url, "application/xml")), &data)
	if err != nil {
		fmt.Print("Errored unmarshaling skill %s\n", err)
		os.Exit(1)
	}
	var skillList = createSkillList(data)
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)
	defer w.Flush()
	fmt.Fprintf(w, "\n %s\t%s\t", "Skill Name", "Number of Employees with Skill")
	fmt.Fprintf(w, "\n %s\t%s\t", "-----------", "-----------")
	for _, v := range skillList.skills {
		fmt.Fprintf(w, "\n %s\t%d\t", v.name, v.employeeCount)
	}
}
