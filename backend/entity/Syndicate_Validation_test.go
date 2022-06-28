package entity

import (
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestSAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	syndicate := Syndicate{
		Name: "กขคกขค",
		Rank: "กข",
		Tel:  "",
	}

	ok, err := govalidator.ValidateStruct(syndicate)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestNameMore5(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	syndicate := Syndicate{
		Name: "กขคก",
		Rank: "กข",
		Tel:  "",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(syndicate)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Name must be more than 5"))
}

func TestRankMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	syndicate := Syndicate{
		Name: "กขคกขค",
		Rank: "",
		Tel:  "",
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(syndicate)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Rank cannot be blank"))
}
