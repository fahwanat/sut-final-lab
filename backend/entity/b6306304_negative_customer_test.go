package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestCustomerIDValidate(t *testing.T) {

	g := NewGomegaWithT(t)

	// t.Run("pass case"),
	// func TestCustomerValidate (t *testing.T){

	n := Customer{

		Name:       "Emy",
		Email:      "Emy@gmail.com",
		CustomerID: "M1234",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).ToNot(BeTrue())
	g.Expect(err).ToNot(BeNil())
	g.Expect(err.Error()).To(Equal("CustomerID: M1234 does not validate as matches(^[LMH]\\d{7}$)"))

}
