package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	// "gorm.io/gorm"
)


func TestSuccessValidate(t *testing.T) {

	g := NewGomegaWithT(t)
	n := Customer{

		Name:       "Emy",
		Email:      "Emy@gmail.com",
		CustomerID: "M1234567",
	}

	ok, err := govalidator.ValidateStruct(n)
	g.Expect(ok).To(BeTrue())
	g.Expect(err).To(BeNil())
	// g.Expect(err.Error()).To(Equal("CustomerID: The following validator is invalid or can't be applied to the field: \"matches[LMH]^\\\\d{7}$\";Name cannot be blank"))

}
