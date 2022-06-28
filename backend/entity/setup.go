package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("suta.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		&Officer{},
		&Syndicate{},
		&Member{},
		&Category{},
		&Trader{},
		&Gown{},
	)

	db = database

	//- Officer Data -
	password, err := bcrypt.GenerateFromPassword([]byte("123"), 14)

	db.Model(&Officer{}).Create(&Officer{
		Name:     "Worrakarn",
		Email:    "w@gmail.com",
		Password: string(password),
	})

	var Worrakarn Officer

	db.Raw("SELECT * FROM officers WHERE email = ?", "w@gmail.com").Scan(&Worrakarn)

	// - Branch Data -
	NoneB := Branch{
		Branch: "-",
	}
	db.Model(&Branch{}).Create(&NoneB)

	a1 := Branch{
		Branch: "สำนักวิชาวิศวกรรมศาสตร์",
	}
	db.Model(&Branch{}).Create(&a1)

	a2 := Branch{
		Branch: "วิศวกรรมการผลิต",
	}
	db.Model(&Branch{}).Create(&a2)

	a3 := Branch{
		Branch: "วิศวกรรมเกษตร",
	}
	db.Model(&Branch{}).Create(&a3)

	a4 := Branch{
		Branch: "วิศวกรรมขนส่ง",
	}
	db.Model(&Branch{}).Create(&a4)

	a5 := Branch{
		Branch: "วิศวกรรมคอมพิวเตอร์",
	}
	db.Model(&Branch{}).Create(&a5)

	a6 := Branch{
		Branch: "วิศวกรรมเคมี",
	}
	db.Model(&Branch{}).Create(&a6)

	a7 := Branch{
		Branch: "วิศวกรรมเครื่องกล",
	}
	db.Model(&Branch{}).Create(&a7)

	a8 := Branch{
		Branch: "วิศวกรรมเซรามิก",
	}
	db.Model(&Branch{}).Create(&a8)

	a9 := Branch{
		Branch: "วิศวกรรมโทรคมนาคม",
	}
	db.Model(&Branch{}).Create(&a9)

	a10 := Branch{
		Branch: "วิศวกรรมพอลิเมอร์",
	}
	db.Model(&Branch{}).Create(&a10)

	a11 := Branch{
		Branch: "วิศวกรรมไฟฟ้า",
	}
	db.Model(&Branch{}).Create(&a11)

	a12 := Branch{
		Branch: "วิศวกรรมโยธา",
	}
	db.Model(&Branch{}).Create(&a12)

	a13 := Branch{
		Branch: "วิศวกรรมโลหการ",
	}
	db.Model(&Branch{}).Create(&a13)

	a14 := Branch{
		Branch: "วิศวกรรมสิ่งแวดล้อม",
	}
	db.Model(&Branch{}).Create(&a14)

	a15 := Branch{
		Branch: "วิศวกรรมอิเล็กทรอนิกส์",
	}
	db.Model(&Branch{}).Create(&a15)

	a16 := Branch{
		Branch: "วิศวกรรมอุตสาหการ",
	}
	db.Model(&Branch{}).Create(&a16)

	a17 := Branch{
		Branch: "เทคโนโลยีการพิมพ์",
	}
	db.Model(&Branch{}).Create(&a17)

	a18 := Branch{
		Branch: "เทคโนโลยีธรณี",
	}
	db.Model(&Branch{}).Create(&a18)

	a19 := Branch{
		Branch: "วิศวกรรมยานยนต์",
	}
	db.Model(&Branch{}).Create(&a19)

	a20 := Branch{
		Branch: "แมคคาทรอนิกส์",
	}
	db.Model(&Branch{}).Create(&a20)

	a21 := Branch{
		Branch: "วิศวกรรมเมคคาทรอนิกส์",
	}
	db.Model(&Branch{}).Create(&a21)

	a22 := Branch{
		Branch: "วิศวกรรมอากาศยาน",
	}
	db.Model(&Branch{}).Create(&a22)

	a23 := Branch{
		Branch: "วิศวกรรมเกษตรและอาหาร",
	}
	db.Model(&Branch{}).Create(&a23)

	a24 := Branch{
		Branch: "วิศวกรรมขนส่งและโลจิสติกส์",
	}
	db.Model(&Branch{}).Create(&a24)

	a25 := Branch{
		Branch: "วิศวกรรมธรณี",
	}
	db.Model(&Branch{}).Create(&a25)

	a26 := Branch{
		Branch: "วิศวกรรมออกแบบผลิตภัณฑ์",
	}
	db.Model(&Branch{}).Create(&a26)

	a27 := Branch{
		Branch: "วิศวกรรมเครื่องมือ",
	}
	db.Model(&Branch{}).Create(&a27)

	a28 := Branch{
		Branch: "วิศวกรรมพรีซิชั่น",
	}
	db.Model(&Branch{}).Create(&a28)

	a29 := Branch{
		Branch: "วิศวกรรมปิโตรเลียมและเทคโนโลยีธรณี",
	}
	db.Model(&Branch{}).Create(&a29)

	a30 := Branch{
		Branch: "วิศวกรรมนวัตกรรมและการออกแบบวัสดุ",
	}
	db.Model(&Branch{}).Create(&a30)

	a31 := Branch{
		Branch: "วิศวกรรมปิโตรเคมีและพอลิเมอร์",
	}
	db.Model(&Branch{}).Create(&a31)

	a32 := Branch{
		Branch: "วิศวกรรมการผลิตอัตโนมัติและหุ่นยนต์",
	}
	db.Model(&Branch{}).Create(&a32)

	a33 := Branch{
		Branch: "วิศวกรรมโยธาและโครงสร้างพื้นฐาน",
	}
	db.Model(&Branch{}).Create(&a33)

	a34 := Branch{
		Branch: "วิศวกรรมไฟฟ้าอุตสาหกรรม",
	}
	db.Model(&Branch{}).Create(&a34)

	a35 := Branch{
		Branch: "วิศวกรรมปิโตรเคมีและพอลิเมอร์",
	}
	db.Model(&Branch{}).Create(&a35)

	a36 := Branch{
		Branch: "วิศวกรรมโยธาและการบริหารงานก่อสร้าง",
	}
	db.Model(&Branch{}).Create(&a36)

	a37 := Branch{
		Branch: "สำนักวิชาวิทยาศาสตร์",
	}
	db.Model(&Branch{}).Create(&a37)

	a38 := Branch{
		Branch: "คณิตศาสตร์ประยุกต์",
	}
	db.Model(&Branch{}).Create(&a38)

	a39 := Branch{
		Branch: "สำนักวิชาเทคโนโลยีสังคม",
	}
	db.Model(&Branch{}).Create(&a39)

	a40 := Branch{
		Branch: "นิเทศศาสตร์",
	}
	db.Model(&Branch{}).Create(&a40)

	a41 := Branch{
		Branch: "การจัดการการตลาด",
	}
	db.Model(&Branch{}).Create(&a41)

	a42 := Branch{
		Branch: "การจัดการโลจิสติกส์",
	}
	db.Model(&Branch{}).Create(&a42)

	a43 := Branch{
		Branch: "สำนักวิชาเทคโนโลยีการเกษตร",
	}
	db.Model(&Branch{}).Create(&a43)

	a44 := Branch{
		Branch: "เทคโนโลยีอาหาร",
	}
	db.Model(&Branch{}).Create(&a44)

	a45 := Branch{
		Branch: "เทคโนโลยีการผลิตพืช",
	}
	db.Model(&Branch{}).Create(&a45)

	a46 := Branch{
		Branch: "สำนักวิชาแพทยศาสตร์",
	}
	db.Model(&Branch{}).Create(&a46)

	a47 := Branch{
		Branch: "สำนักวิชาพยาบาลศาสตร์",
	}
	db.Model(&Branch{}).Create(&a47)

	a48 := Branch{
		Branch: "สำนักวิชาทันตแพทยศาสตร์",
	}
	db.Model(&Branch{}).Create(&a48)

	a49 := Branch{
		Branch: "สำนักวิชาสาธารณสุขศาสตร์",
	}
	db.Model(&Branch{}).Create(&a49)

	//Category data
	c1 := Category{
		Category: "สมทบ",
	}
	db.Model(&Category{}).Create(&c1)

	c2 := Category{
		Category: "สามัญ",
	}
	db.Model(&Category{}).Create(&c2)

	c3 := Category{
		Category: "กิตติมศักดิ์",
	}
	db.Model(&Category{}).Create(&c3)

	//SubDistrict data
	None := SubDistrict{
		SubDistrict: "-",
	}
	db.Model(&SubDistrict{}).Create(&None)

	s1 := SubDistrict{
		SubDistrict: "สุรนารี",
	}
	db.Model(&SubDistrict{}).Create(&s1)

	s2 := SubDistrict{
		SubDistrict: "ในเมือง",
	}
	db.Model(&SubDistrict{}).Create(&s2)

	s3 := SubDistrict{
		SubDistrict: "โพธิ์กลาง",
	}
	db.Model(&SubDistrict{}).Create(&s3)

	s4 := SubDistrict{
		SubDistrict: "หนองจะบก",
	}
	db.Model(&SubDistrict{}).Create(&s4)

	s5 := SubDistrict{
		SubDistrict: "โคกสูง",
	}
	db.Model(&SubDistrict{}).Create(&s5)

	s6 := SubDistrict{
		SubDistrict: "มะเริง",
	}
	db.Model(&SubDistrict{}).Create(&s6)

	s7 := SubDistrict{
		SubDistrict: "หนองระเวียง",
	}
	db.Model(&SubDistrict{}).Create(&s7)

	s8 := SubDistrict{
		SubDistrict: "ปรุใหญ่",
	}
	db.Model(&SubDistrict{}).Create(&s8)

	s9 := SubDistrict{
		SubDistrict: "หมื่นไวย",
	}
	db.Model(&SubDistrict{}).Create(&s9)

	s10 := SubDistrict{
		SubDistrict: "พลกรัง",
	}
	db.Model(&SubDistrict{}).Create(&s10)

	s11 := SubDistrict{
		SubDistrict: "หนองไผ่ล้อม",
	}
	db.Model(&SubDistrict{}).Create(&s11)

	s12 := SubDistrict{
		SubDistrict: "หัวทะเล",
	}
	db.Model(&SubDistrict{}).Create(&s12)

	s13 := SubDistrict{
		SubDistrict: "บ้านเกาะ",
	}
	db.Model(&SubDistrict{}).Create(&s13)

	s14 := SubDistrict{
		SubDistrict: "บ้านใหม่",
	}
	db.Model(&SubDistrict{}).Create(&s14)

	s15 := SubDistrict{
		SubDistrict: "พุดชา",
	}
	db.Model(&SubDistrict{}).Create(&s15)

	s16 := SubDistrict{
		SubDistrict: "บ้านโพธิ์",
	}
	db.Model(&SubDistrict{}).Create(&s16)

	s17 := SubDistrict{
		SubDistrict: "จอหอ",
	}
	db.Model(&SubDistrict{}).Create(&s17)

	s18 := SubDistrict{
		SubDistrict: "โคกกรวด",
	}
	db.Model(&SubDistrict{}).Create(&s18)

	s19 := SubDistrict{
		SubDistrict: "ไชยมงคล",
	}
	db.Model(&SubDistrict{}).Create(&s19)

	s20 := SubDistrict{
		SubDistrict: "หนองบัวศาลา",
	}
	db.Model(&SubDistrict{}).Create(&s20)

	s21 := SubDistrict{
		SubDistrict: "สีมุม",
	}
	db.Model(&SubDistrict{}).Create(&s21)

	s22 := SubDistrict{
		SubDistrict: "ตลาด",
	}
	db.Model(&SubDistrict{}).Create(&s22)

	s23 := SubDistrict{
		SubDistrict: "พะเนา",
	}
	db.Model(&SubDistrict{}).Create(&s23)

	s24 := SubDistrict{
		SubDistrict: "หนองกระทุ่ม",
	}
	db.Model(&SubDistrict{}).Create(&s24)

	s25 := SubDistrict{
		SubDistrict: "หนองไข่น้ำ",
	}
	db.Model(&SubDistrict{}).Create(&s25)

	//District data
	NoneD := District{
		District: "-",
	}
	db.Model(&District{}).Create(&NoneD)

	d1 := District{
		District: "เมือง",
	}
	db.Model(&District{}).Create(&d1)

	d2 := District{
		District: "ขามทะเลสอ",
	}
	db.Model(&District{}).Create(&d2)

	d3 := District{
		District: "โนนไทย",
	}
	db.Model(&District{}).Create(&d3)

	d4 := District{
		District: "โชคชัย",
	}
	db.Model(&District{}).Create(&d4)

	d5 := District{
		District: "เฉลิมพระเกียรติ",
	}
	db.Model(&District{}).Create(&d5)

	d6 := District{
		District: "โนนแดง",
	}
	db.Model(&District{}).Create(&d6)

	d7 := District{
		District: "ปักธงชัย",
	}
	db.Model(&District{}).Create(&d7)

	d8 := District{
		District: "สูงเนิน",
	}
	db.Model(&District{}).Create(&d8)

	d9 := District{
		District: "โนนสูง",
	}
	db.Model(&District{}).Create(&d9)

	d10 := District{
		District: "จักราช",
	}
	db.Model(&District{}).Create(&d10)

	d11 := District{
		District: "สีคิ้ว",
	}
	db.Model(&District{}).Create(&d11)

	d12 := District{
		District: "ขามสะแกแสง",
	}
	db.Model(&District{}).Create(&d12)

	d13 := District{
		District: "พระทองคำ",
	}
	db.Model(&District{}).Create(&d13)

	d14 := District{
		District: "หนองบุญมาก",
	}
	db.Model(&District{}).Create(&d14)

	d15 := District{
		District: "ครบุรี",
	}
	db.Model(&District{}).Create(&d15)

	d16 := District{
		District: "พิมาย",
	}
	db.Model(&District{}).Create(&d16)

	d17 := District{
		District: "ห้วยแถลง",
	}
	db.Model(&District{}).Create(&d17)

	d18 := District{
		District: "วังน้ำเขียว",
	}
	db.Model(&District{}).Create(&d18)

	d19 := District{
		District: "คง",
	}
	db.Model(&District{}).Create(&d19)

	d20 := District{
		District: "ด่านขุนทด",
	}
	db.Model(&District{}).Create(&d20)

	d21 := District{
		District: "ปากช่อง",
	}
	db.Model(&District{}).Create(&d21)

	d22 := District{
		District: "บ้านเหลื่อม",
	}
	db.Model(&District{}).Create(&d22)

	d23 := District{
		District: "เสิงสาง",
	}
	db.Model(&District{}).Create(&d23)

	d24 := District{
		District: "เทพารักษ์",
	}
	db.Model(&District{}).Create(&d24)

	d25 := District{
		District: "ประทาย",
	}
	db.Model(&District{}).Create(&d25)

	d26 := District{
		District: "ชุมพวง",
	}
	db.Model(&District{}).Create(&d26)

	d27 := District{
		District: "บัวใหญ่",
	}
	db.Model(&District{}).Create(&d27)

	d28 := District{
		District: "เมืองยาง",
	}
	db.Model(&District{}).Create(&d28)

	d29 := District{
		District: "แก้งสนามนาง",
	}
	db.Model(&District{}).Create(&d29)

	d30 := District{
		District: "บัวลาย",
	}
	db.Model(&District{}).Create(&d30)

	d31 := District{
		District: "สีดา",
	}
	db.Model(&District{}).Create(&d31)

	d32 := District{
		District: "ลำทะเมนชัย",
	}
	db.Model(&District{}).Create(&d32)

	//Province data
	NoneP := Province{
		Province: "-",
	}
	db.Model(&Province{}).Create(&NoneP)

	p1 := Province{
		Province: "กระบี่",
	}
	db.Model(&Province{}).Create(&p1)

	p2 := Province{
		Province: "กรุงเทพมหานคร",
	}
	db.Model(&Province{}).Create(&p2)

	p3 := Province{
		Province: "กาญจนบุรี",
	}
	db.Model(&Province{}).Create(&p3)

	p4 := Province{
		Province: "กาฬสินธุ์",
	}
	db.Model(&Province{}).Create(&p4)

	p5 := Province{
		Province: "กำแพงเพชร",
	}
	db.Model(&Province{}).Create(&p5)

	p6 := Province{
		Province: "ขอนแก่น",
	}
	db.Model(&Province{}).Create(&p6)

	p7 := Province{
		Province: "จันทบุรี",
	}
	db.Model(&Province{}).Create(&p7)

	p8 := Province{
		Province: "ฉะเชิงเทรา",
	}
	db.Model(&Province{}).Create(&p8)

	p9 := Province{
		Province: "ชลบุรี",
	}
	db.Model(&Province{}).Create(&p9)

	p10 := Province{
		Province: "ชัยนาท",
	}
	db.Model(&Province{}).Create(&p10)

	p11 := Province{
		Province: "ชัยภูมิ",
	}
	db.Model(&Province{}).Create(&p11)

	p12 := Province{
		Province: "ชุมพร",
	}
	db.Model(&Province{}).Create(&p12)

	p13 := Province{
		Province: "เชียงราย",
	}
	db.Model(&Province{}).Create(&p13)

	p14 := Province{
		Province: "เชียงใหม่",
	}
	db.Model(&Province{}).Create(&p14)

	p15 := Province{
		Province: "ตรัง",
	}
	db.Model(&Province{}).Create(&p15)

	p16 := Province{
		Province: "ตราด",
	}
	db.Model(&Province{}).Create(&p16)

	p17 := Province{
		Province: "ตาก",
	}
	db.Model(&Province{}).Create(&p17)

	p18 := Province{
		Province: "นครนายก",
	}
	db.Model(&Province{}).Create(&p18)

	p19 := Province{
		Province: "นครปฐม",
	}
	db.Model(&Province{}).Create(&p19)

	p20 := Province{
		Province: "นครพนม",
	}
	db.Model(&Province{}).Create(&p20)

	p21 := Province{
		Province: "นครราชสีมา",
	}
	db.Model(&Province{}).Create(&p21)

	p22 := Province{
		Province: "นครศรีธรรมราช",
	}
	db.Model(&Province{}).Create(&p22)

	p23 := Province{
		Province: "นครสวรรค์",
	}
	db.Model(&Province{}).Create(&p23)

	p24 := Province{
		Province: "นนทบุรี",
	}
	db.Model(&Province{}).Create(&p24)

	p25 := Province{
		Province: "นราธิวาส",
	}
	db.Model(&Province{}).Create(&p25)

	p26 := Province{
		Province: "น่าน",
	}
	db.Model(&Province{}).Create(&p26)

	p27 := Province{
		Province: "บึงกาฬ",
	}
	db.Model(&Province{}).Create(&p27)

	p28 := Province{
		Province: "บุรีรัมย์",
	}
	db.Model(&Province{}).Create(&p28)

	p29 := Province{
		Province: "ปทุมธานี",
	}
	db.Model(&Province{}).Create(&p29)

	p30 := Province{
		Province: "ประจวบคีรีขันธ์",
	}
	db.Model(&Province{}).Create(&p30)

	p31 := Province{
		Province: "ปราจีนบุรี",
	}
	db.Model(&Province{}).Create(&p31)

	p32 := Province{
		Province: "ปัตตานี",
	}
	db.Model(&Province{}).Create(&p32)

	p33 := Province{
		Province: "พระนครศรีอยุธยา",
	}
	db.Model(&Province{}).Create(&p33)

	p34 := Province{
		Province: "พะเยา",
	}
	db.Model(&Province{}).Create(&p34)

	p35 := Province{
		Province: "พังงา",
	}
	db.Model(&Province{}).Create(&p35)

	p36 := Province{
		Province: "พัทลุง",
	}
	db.Model(&Province{}).Create(&p36)

	p37 := Province{
		Province: "พิจิตร",
	}
	db.Model(&Province{}).Create(&p37)

	p38 := Province{
		Province: "พิษณุโลก",
	}
	db.Model(&Province{}).Create(&p38)

	p39 := Province{
		Province: "เพชรบุรี",
	}
	db.Model(&Province{}).Create(&p39)

	p40 := Province{
		Province: "เพชรบูรณ์",
	}
	db.Model(&Province{}).Create(&p40)

	p41 := Province{
		Province: "แพร่",
	}
	db.Model(&Province{}).Create(&p41)

	p42 := Province{
		Province: "ภูเก็ต",
	}
	db.Model(&Province{}).Create(&p42)

	p43 := Province{
		Province: "มหาสารคาม",
	}
	db.Model(&Province{}).Create(&p43)

	p44 := Province{
		Province: "มุกดาหาร",
	}
	db.Model(&Province{}).Create(&p44)

	p45 := Province{
		Province: "แม่ฮ่องสอน",
	}
	db.Model(&Province{}).Create(&p45)

	p46 := Province{
		Province: "ยโสธร",
	}
	db.Model(&Province{}).Create(&p46)

	p47 := Province{
		Province: "ยะลา",
	}
	db.Model(&Province{}).Create(&p47)

	p48 := Province{
		Province: "ร้อยเอ็ด",
	}
	db.Model(&Province{}).Create(&p48)

	p49 := Province{
		Province: "ระนอง",
	}
	db.Model(&Province{}).Create(&p49)

	p50 := Province{
		Province: "ระยอง",
	}
	db.Model(&Province{}).Create(&p50)

	p51 := Province{
		Province: "ราชบุรี",
	}
	db.Model(&Province{}).Create(&p51)

	p52 := Province{
		Province: "ลพบุรี",
	}
	db.Model(&Province{}).Create(&p52)

	p53 := Province{
		Province: "ลำปาง",
	}
	db.Model(&Province{}).Create(&p53)

	p54 := Province{
		Province: "ลำพูน",
	}
	db.Model(&Province{}).Create(&p54)

	p55 := Province{
		Province: "เลย",
	}
	db.Model(&Province{}).Create(&p55)

	p56 := Province{
		Province: "ศรีสะเกษ",
	}
	db.Model(&Province{}).Create(&p56)

	p57 := Province{
		Province: "สกลนคร",
	}
	db.Model(&Province{}).Create(&p57)

	p58 := Province{
		Province: "สงขลา",
	}
	db.Model(&Province{}).Create(&p58)

	p59 := Province{
		Province: "สตูล",
	}
	db.Model(&Province{}).Create(&p59)

	p60 := Province{
		Province: "สมุทรปราการ",
	}
	db.Model(&Province{}).Create(&p60)

	p61 := Province{
		Province: "สมุทรสงคราม",
	}
	db.Model(&Province{}).Create(&p61)

	p62 := Province{
		Province: "สมุทรสาคร",
	}
	db.Model(&Province{}).Create(&p62)

	p63 := Province{
		Province: "สระแก้ว",
	}
	db.Model(&Province{}).Create(&p63)

	p64 := Province{
		Province: "สระบุรี",
	}
	db.Model(&Province{}).Create(&p64)

	p65 := Province{
		Province: "สิงห์บุรี",
	}
	db.Model(&Province{}).Create(&p65)

	p66 := Province{
		Province: "สุโขทัย",
	}
	db.Model(&Province{}).Create(&p66)

	p67 := Province{
		Province: "สุพรรณบุรี",
	}
	db.Model(&Province{}).Create(&p67)

	p68 := Province{
		Province: "สุราษฎร์ธานี",
	}
	db.Model(&Province{}).Create(&p68)

	p69 := Province{
		Province: "สุรินทร์",
	}
	db.Model(&Province{}).Create(&p69)

	p70 := Province{
		Province: "หนองคาย",
	}
	db.Model(&Province{}).Create(&p70)

	p71 := Province{
		Province: "หนองบัวลำภู",
	}
	db.Model(&Province{}).Create(&p71)

	p72 := Province{
		Province: "อ่างทอง",
	}
	db.Model(&Province{}).Create(&p72)

	p73 := Province{
		Province: "อำนาจเจริญ",
	}
	db.Model(&Province{}).Create(&p73)

	p74 := Province{
		Province: "อุดรธานี",
	}
	db.Model(&Province{}).Create(&p74)

	p75 := Province{
		Province: "อุตรดิตถ์",
	}
	db.Model(&Province{}).Create(&p75)

	p76 := Province{
		Province: "อุทัยธานี",
	}
	db.Model(&Province{}).Create(&p76)

	p77 := Province{
		Province: "อุบลราชธานี",
	}
	db.Model(&Province{}).Create(&p77)

	//Size data
	si1 := Size{
		Size: "SS",
	}
	db.Model(&Size{}).Create(&si1)

	si2 := Size{
		Size: "S",
	}
	db.Model(&Size{}).Create(&si2)

	si3 := Size{
		Size: "M",
	}
	db.Model(&Size{}).Create(&si3)

	si4 := Size{
		Size: "L",
	}
	db.Model(&Size{}).Create(&si4)

	si5 := Size{
		Size: "XL",
	}
	db.Model(&Size{}).Create(&si5)

	si6 := Size{
		Size: "XXL",
	}
	db.Model(&Size{}).Create(&si6)

	si7 := Size{
		Size: "3XL",
	}
	db.Model(&Size{}).Create(&si7)

	si8 := Size{
		Size: "4XL",
	}
	db.Model(&Size{}).Create(&si8)

	si9 := Size{
		Size: "5XL",
	}
	db.Model(&Size{}).Create(&si9)

	si10 := Size{
		Size: "6XL",
	}
	db.Model(&Size{}).Create(&si10)

	//Blazer data
	Nonebl := Blazer{
		Blazer: "-",
	}
	db.Model(&Blazer{}).Create(&Nonebl)

	bl1 := Blazer{
		Blazer: "SS",
	}
	db.Model(&Blazer{}).Create(&bl1)

	bl2 := Blazer{
		Blazer: "S",
	}
	db.Model(&Blazer{}).Create(&bl2)

	bl3 := Blazer{
		Blazer: "M",
	}
	db.Model(&Blazer{}).Create(&bl3)

	bl4 := Blazer{
		Blazer: "L",
	}
	db.Model(&Blazer{}).Create(&bl4)

	bl5 := Blazer{
		Blazer: "XL",
	}
	db.Model(&Blazer{}).Create(&bl5)

	bl6 := Blazer{
		Blazer: "XXL",
	}
	db.Model(&Blazer{}).Create(&bl6)

	bl7 := Blazer{
		Blazer: "3XL",
	}
	db.Model(&Blazer{}).Create(&bl7)

	bl8 := Blazer{
		Blazer: "4XL",
	}
	db.Model(&Blazer{}).Create(&bl8)

	bl9 := Blazer{
		Blazer: "5XL",
	}
	db.Model(&Blazer{}).Create(&bl9)

	bl10 := Blazer{
		Blazer: "6XL",
	}
	db.Model(&Blazer{}).Create(&bl10)

	bl11 := Blazer{
		Blazer: "ตัดใหม่",
	}
	db.Model(&Blazer{}).Create(&bl11)

	// - Syndicate Data -
	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายรัฐมงคล เหมรัตนานนท์",
		Rank:   "นายกสมาคม",
		Branch: a13,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายศิริชัย นามวิเศษ",
		Rank:   "อุปนายก-กิจกรรม",
		Branch: a11,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายกฤดิภัค โกยดุลย์",
		Rank:   "กรรมการ-กิจกรรม",
		Branch: a8,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายกีรติคุณ บุญประสพ",
		Rank:   "กรรมการ-กิจกรรม",
		Branch: a11,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาวอรุณี อะโสต",
		Rank:   "กรรมการ-กิจกรรม",
		Branch: a7,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายทรรศสิน เดชะกูล",
		Rank:   "กรรมการ-กิจกรรม",
		Branch: a4,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายอมฤต ฉายะรถี",
		Rank:   "กรรมการ-กิจกรรม",
		Branch: a16,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "ดร.สมชาย สุขอินทร์",
		Rank:   "อุปนายก-วิชาการ",
		Branch: a38,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายวชิระ พุทธิแจ่ม",
		Rank:   "กรรมการ-วิชาการ",
		Branch: a7,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "ผศ.ดร.ปรัชญา เทพณรงค์",
		Rank:   "กรรมการ-วิชาการ",
		Branch: a18,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายเพิ่มพูน เสมอ",
		Rank:   "กรรมการ-วิชาการ",
		Branch: a11,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายกงจักร ลมวิชัย",
		Rank:   "กรรมการ-วิชาการ",
		Branch: a23,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาววิยภรณ์ กรองทอง",
		Rank:   "กรรมการ-วิชาการ",
		Branch: a8,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายวิศรุตน์ ไชยเพ็ชร",
		Rank:   "เลขาธิการ",
		Branch: a16,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายบรรพต สุทธิพันธ์",
		Rank:   "กรรมการ-เลขาธิการ",
		Branch: a18,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "ร.ท.รัฐพงษ์ อ่อนจันทร์",
		Rank:   "เหรัญญิก",
		Branch: a5,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายวีระพันธ์ เต็มดวง",
		Rank:   "กรรมการ-การเงิน",
		Branch: a16,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาวพันธกานต์ รื่นกลิ่น",
		Rank:   "ปฏิคม",
		Branch: a40,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายคมเดช ธูปกิ่ง",
		Rank:   "กรรมการ-ปฏิคม",
		Branch: a8,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายธนาทร อารยวริต",
		Rank:   "นายทะเบียน",
		Branch: a36,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาวฉัฐพร ทองวิจิตต์",
		Rank:   "กรรมการ-ทะเบียน",
		Branch: a14,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายอดิศักดิ์ ยั่งยืน",
		Rank:   "ประชาสัมพันธ์",
		Branch: a14,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายเอกชัย บุญรสศักดิ์",
		Rank:   "กรรมการ-ประชาสัมพันธ์",
		Branch: a5,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายอุดมศักดิ์ บัวสำราญ",
		Rank:   "กรรมการ-ประชาสัมพันธ์",
		Branch: a44,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นายสุเมธี ถีสูงเนิน",
		Rank:   "กรรมการ-ประชาสัมพันธ์",
		Branch: a21,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาวจิรนันท์ มหาดไทย",
		Rank:   "เจ้าหน้าที่บริหารทั่วไป",
		Branch: a45,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาวสุชาดา สงวนพันธุ์",
		Rank:   "เจ้าหน้าที่บริหารทั่วไป",
		Branch: a41,
		Tel:    "0801234567",
	})

	db.Model(&Syndicate{}).Create(&Syndicate{
		Name:   "นางสาววิจิตรา แสนโคตร",
		Rank:   "เจ้าหน้าที่บริหารทั่วไป",
		Branch: a42,
		Tel:    "0801234567",
	})

	// - Member Data -
	db.Model(&Member{}).Create(&Member{
		Name:        "ศ.ดร.วิจิตร ศรีสอ้าน",
		Num:         "0001",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c3,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "อ.ดร.วิศิษฎ์พร วัฒนวาทิน",
		Num:         "0002",
		Branch:      a39,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ผศ.ดร.เอมอร ทัศนศร",
		Num:         "0003",
		Branch:      a1,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ศ.ดร.หนึ่ง เตียอำรุง",
		Num:         "0004",
		Branch:      a43,
		Email:       "neung@sut.ac.th",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "รศ.ดร.ไทย ทิพย์สุวรรณกุล",
		Num:         "0005",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "ม.วลัยลักษณ์",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ผศ.ดร.อรชุน ไชยเสนะ",
		Num:         "0006",
		Branch:      a37,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "นางสาวจุฑามาศ สวัสดี",
		Num:         "0007",
		Branch:      NoneB,
		Email:       "chuthas@sut.ac.th",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "นายมนู อัศวเสนา",
		Num:         "0008",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ผศ.ดร.ตริตาภรณ์ ชูศรี",
		Num:         "0009",
		Branch:      a37,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "อ.ดร.ธวัชชัย ภิญโญรัฐกานต์",
		Num:         "0010",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "อ.ดร.วุฒิ ด่านกิตติกุล",
		Num:         "0011",
		Branch:      a1,
		Email:       "wut@sut.ac.th",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ผศ.ดร.อุทัย มีคำ",
		Num:         "0012",
		Branch:      a1,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "รศ.ร.อ.ดร.กนต์ธร ชำนิประศาสน์",
		Num:         "0013",
		Branch:      a1,
		Email:       "kontorn@sut.ac.th",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "นายราชัย อัศเวศน์",
		Num:         "0014",
		Branch:      NoneB,
		Email:       "rachai@sut.ac.th",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "อ.ดร.เบญจมาศ จิตรสมบูรณ์",
		Num:         "0015",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "รศ.พญ.อาภรณ์ ปราบริปูตลุง",
		Num:         "0016",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "Dr.Joewono Widjaja",
		Num:         "0017",
		Branch:      a37,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "รศ.ดร.เกษม ปราบริปูตลุง",
		Num:         "0018",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "อ.ดร.ณรงค์ อัครพัฒนากูล",
		Num:         "0019",
		Branch:      a1,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "111 30000",
		SubDistrict: s1,
		District:    d1,
		Province:    p21,
		Date:        time.Now(),
		Category:    c2,
	})

	db.Model(&Member{}).Create(&Member{
		Name:        "ผศ.ธารา เล็กอุทัย",
		Num:         "0020",
		Branch:      NoneB,
		Email:       "-",
		Tel:         "0801234567",
		Address:     "-",
		SubDistrict: None,
		District:    NoneD,
		Province:    NoneP,
		Date:        time.Now(),
		Category:    c2,
	})

	// - Trader Data -
	db.Model(&Trader{}).Create(&Trader{
		Name:    "วิจิตร ศรีสอ้าน",
		Tel:     "0801234567",
		Address: "111 ต.สุรนารี อ.เมือง จ.นครราชสีมา 30000",
	})

	//Degree data
	de1 := Degree{
		Degree: "ปริญญาตรี",
	}
	db.Model(&Degree{}).Create(&de1)

	de2 := Degree{
		Degree: "ปริญญาโท",
	}
	db.Model(&Degree{}).Create(&de2)

	de3 := Degree{
		Degree: "ปริญญาเอก",
	}
	db.Model(&Degree{}).Create(&de3)

	// - Gown Data -
	db.Model(&Gown{}).Create(&Gown{
		Name:    "วรกานต์ ปินชัย",
		Student: "B1111111",
		Status:  "ซื้อ",
		Size:    si3,
		Blazer:  Nonebl,
		Degree:  de1,
		Branch:  a5,
		Tel:     "0801234567",
		Date:    time.Now(),
	})
}
