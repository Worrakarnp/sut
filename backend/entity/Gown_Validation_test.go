package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
)

func TestGAllpass(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	gown := Gown{
		Name:    "กกขขคค",
		Student: "",
		Status:  "",
		Tel:     "",
		Date:    time.Now(),
	}

	ok, err := govalidator.ValidateStruct(gown)

	g.Expect(ok).To(gomega.BeTrue())
	g.Expect(err).To(gomega.BeNil())

}

func TestNamegMustNotBlank(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	gown := Gown{
		Name:    "",
		Student: "",
		Status:  "",
		Tel:     "",
		Date:    time.Now(),
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(gown)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Name cannot be blank"))
}

func TestGownDateNotBeFuture(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	gown := Gown{
		Name:    "กกขขคค",
		Student: "",
		Status:  "",
		Tel:     "",
		Date:    time.Now().Add(time.Hour * 24), // ผิด เป็นอนาคต
	}

	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(gown)

	// ok ต้องไม่เป็นค่า true แปลว่าต้องจับ error ได้
	g.Expect(ok).ToNot(gomega.BeTrue())

	// err ต้องไม่เป็นค่า nil แปลว่าต้องจับ error ได้
	g.Expect(err).ToNot(gomega.BeNil())

	// err.Error ต้องมี error message แสดงออกมา
	g.Expect(err.Error()).To(gomega.Equal("Date cannot be in the future"))
}
