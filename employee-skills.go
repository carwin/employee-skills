package main

import (
	"fmt"
	"os"

	employeeskillset "github.com/carwin/employee-skills/employee-skillset"
	skilldistribution "github.com/carwin/employee-skills/skill-distribution"
	skillpractitioners "github.com/carwin/employee-skills/skill-practitioners"
)

func printHelp() {
	fmt.Println("\n")
	fmt.Println("Usage: employee-skills <command> <options>")
	fmt.Println("\n")
	fmt.Println("Commands:")
	fmt.Println("  employee-skills help\t\tPrints this menu")
	fmt.Println("  employee-skills distribution\tPrints the distribution of skills across all employees")
	fmt.Println("  employee-skills skill\t\tPrints the employees that know the skill")
	fmt.Println("  \t\t\t\tOptions: ")
	fmt.Println("  \t\t\t\t  - A string representing a skill")
	fmt.Println("  employee-skills employee\tPrints the skills for a given employee")
	fmt.Println("  \t\t\t\tOptions: ")
	fmt.Println("  \t\t\t\t  - A string representing the first and last name of an employee")
	fmt.Println("\n")
	fmt.Println("Examples:")
	fmt.Println("  employee-skills distribution\t\t\tGet the company-wide skill distribution")
	fmt.Println("  employee-skills employee \"Carwin Young\"\tReturn all of Carwin Young's skills")
	fmt.Println("  employee-skills skill \"JavaScript\"\t\tReturn a list of everyone familiar with JavaScript")
}

// Conmmands
//
// distribution
//   prints the employee skills with a count of employees that know the skill
//
// employee
//   prints a specific employee's skills
//   @string employee name
//
// skill
//   prints a list of employees that know the given skill
//   @string skill name
func main() {

	if len(os.Args) < 2 {
		printHelp()
	} else {
		command := os.Args[1]
		switch command {
		case "distribution":
			skill_distribution.GetSkillDistribution()
		case "employee":
			var employee = os.Args[2]
			employee_skillset.GetEmployeeSkillset(employee)
			break
		case "skill":
			var skillName = os.Args[2]
			skill_practitioners.GetSkillPractitioners(skillName)
			break
		case "help":
			printHelp()
			break
		default:
			printHelp()
		}
	}

}
