package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNameValidate(t *testing.T) {

	g := NewGomegaWithT(t)
	n := Customer{

		Name:       "",
		Email:      "Emy@gmail.com",
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("Name cannot be blank"))

}
