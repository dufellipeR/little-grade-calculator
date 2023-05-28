package tests_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "little_grade_calculator/calculator"
)

var _ = Describe("CalculateGrade", func() {
	It("should test that the solution returns the correct value", func() {

		correct, wrong, total, grade := CalculateGrade("tftttttttttttttttttttttftffffffttttf")

		Expect(correct).To(Equal(int16(27)))
		Expect(wrong).To(Equal(int16(9)))
		Expect(total).To(Equal(int16(36)))
		Expect(grade).To(Equal(float32(75.00)))

	})
})
