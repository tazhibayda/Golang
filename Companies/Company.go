package Companies

import "fmt"

type Person struct {
	Name       string
	Age        int
	Experience int
}

type Company struct {
	Name       string
	Experience int
	Employees  []Person
}

func (c *Company) Hire(p Person) string {
	if c.Experience <= p.Experience {
		c.Employees = append(c.Employees, p)
		return fmt.Sprintf("We successfully hired %s with experience %d", p.Name, p.Experience)
	}
	return fmt.Sprintf("Unfortunatelly we cant hire %s due to lack of his experience", p.Name)
}

func (p *Person) PersonInfo() string {
	return fmt.Sprintf("%s is %d old with %d year experience.", p.Name, p.Age, p.Experience)
}

func (c *Company) CompanyInfo() string {
	return fmt.Sprintf("%s has %d employees", c.Name, len(c.Employees))
}

func (c *Company) CompanyEmployees() []Person {
	return c.Employees
}
