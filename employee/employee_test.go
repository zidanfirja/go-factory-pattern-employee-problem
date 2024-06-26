package employee_test

import (
	"go-factorypattern-company-case/employee"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Employee", func() {

	Context("Manager Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("manager")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Manager"))
			Expect(empl.GetSalary()).To(Equal(1000))
		})

		It("should return the correct bonus", func() {
			empl, err := employee.GetEmployeeFactory("manager")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetSalary()).To(Equal(1000))
			Expect(empl.GetBonus()).To(Equal(float64(200)))

			var PersenSalary float64 = float64(empl.GetBonus()/1000.0) * 100
			Expect(PersenSalary).To(Equal(float64(20)))

			// TODO Implement the test for the bonus
			// Salary is 1000
			// Bonus is 20% of the salary
			// Bonus is 200
		})

	})

	Context("Staff Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("staff")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Staff"))
			Expect(empl.GetSalary()).To(Equal(500))
		})

		It("should return the correct bonus", func() {
			empl, err := employee.GetEmployeeFactory("staff")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetSalary()).To(Equal(500))
			Expect(empl.GetBonus()).To(Equal(float64(50)))

			var PersenSalary float64 = float64(empl.GetBonus()/500.0) * 100
			Expect(PersenSalary).To(Equal(float64(10)))

			// TODO Implement the test for the bonus
			// Salary is 500
			// Bonus is 10% of the salary
			// Bonus is 50
		})

	})

	Context("Intern Object", func() {

		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("intern")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Intern"))
			Expect(empl.GetSalary()).To(Equal(100))
		})

		It("should return the correct bonus", func() {
			empl, err := employee.GetEmployeeFactory("intern")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetSalary()).To(Equal(100))
			Expect(empl.GetBonus()).To(Equal(0.0))
			var PersenSalary float64 = 0
			Expect(PersenSalary).To(Equal(0.0))
			// TODO Implement the test for the bonus
			// Salary is 100
			// Bonus is 0% of the salary
			// Bonus is 0
		})
	})

	// TODO: Implement the test for the Director object
	Context("Director Object", func() {
		It("should return the correct name and salary", func() {
			empl, err := employee.GetEmployeeFactory("director")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetName()).To(Equal("Director"))
			Expect(empl.GetSalary()).To(Equal(5000))
		})
		It("should return the correct bonus", func() {
			empl, err := employee.GetEmployeeFactory("director")
			Expect(err).NotTo(HaveOccurred())

			Expect(empl.GetBonus()).To(Equal(float64(1500)))

			var PersenSalary float64 = float64(empl.GetBonus()/5000.0) * 100
			Expect(PersenSalary).To(Equal(float64(30)))

			// TODO Implement the test for the bonus
			// Salary is 5000
			// Bonus is 30% of the salary
			// Bonus is 1500
		})
	})

	Context("Empty Employee", func() {

		It("should return an error", func() {
			_, err := employee.GetEmployeeFactory("non-existing")
			Expect(err).To(HaveOccurred())
		})

	})

})
