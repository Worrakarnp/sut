package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestTAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	trader := Trader{
		Name:    "กกขขคค",
		Tel:     "",
		Address: "",
	}

	ok, err := govalidator.ValidateStruct(trader)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestNametMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	trader := Trader{
		Name:    "",
		Tel:     "",
		Address: "",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(trader)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
}
