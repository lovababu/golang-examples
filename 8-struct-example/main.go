package main

import "fmt"

func main()  {

	fmt.Println("============= Go struct example ==============")
	country := Country{
		name: "India",
		population: 12000000,
		areaInHectors: 2500,
	}

	fmt.Printf("Population of '%s' is - %d and density is: %.3f\n", country.name, country.population, country.calculateDensity())

	//initialize struct with new.

	var newCountry = new(Country)
	newCountry.name = "US"
	newCountry.population = 1000000
	newCountry.areaInHectors = 1500
	fmt.Printf("Population of %s is %d and density is : %.3f \n", newCountry.name, newCountry.getPopulation(), newCountry.calculateDensity())


	//
}

//struct declaration.
type Country struct {
	name string
	population int64
	areaInHectors float32
}

//method declaration for struct.
func (c *Country) getPopulation() int64  {
	return c.population;
}

func (c *Country) calculateDensity() float32  {
	return float32(c.population)/c.areaInHectors
}



