package repository_test

import (
	. "github.com/lovababu/golang-examples/cardprocessing-di/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("CardRepository", func() {

	var cRepository  = new(CardRepository)
	Context("Verifying CardRepository functions.", func() {
		It("Test PostTransaction().", func() {
			bool, err := cRepository.PostTransaction("dummy_card", 201.00, time.Now())
			Expect(bool).To(Equal(true))
			Expect(err).Should(BeNil())
		})
	})
})
