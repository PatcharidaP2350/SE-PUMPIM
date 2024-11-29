package config

import (
	"fmt"
	"time"

	"SE-B6527075/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func ConnectionDB() {
	database, err := gorm.Open(sqlite.Open("hospital.db?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connected database")
	db = database
}

func SetupDatabase() {
	// AutoMigrate ตารางทั้งหมด โดยไม่ต้องใส่ EmployeeDisease เพราะ GORM จะสร้างให้อัตโนมัติ
	db.AutoMigrate(
		&entity.BloodGroup{},
		// &entity.Building{},
		&entity.Department{},
		&entity.Disease{},
		// &entity.Drug{},
		&entity.Employee{},
		// &entity.Floor{},
		&entity.Gender{},
		// &entity.Inventory{},
		// &entity.MedicalEquipment{},
		// &entity.MedicalImage{},
		&entity.MedicalRecords{},
		&entity.Patient{},
		// &entity.PatientRoom{},
		&entity.Position{},
		&entity.Queue{},
		// &entity.Room{},
		// &entity.RoomLayout{},
		// &entity.RoomType{},
		// &entity.Schedule{},
		// &entity.Severity{},
		&entity.Specialist{},
		&entity.Status{},
		// &entity.Supplier{},
		&entity.TakeAHistory{},
		// &entity.Employee{},
		// &entity.Gender{},
		// &entity.Position{},
		// &entity.Department{},
		// &entity.Status{},
		// &entity.Specialist{},
		// &entity.Disease{},
		// &entity.BloodGroup{}, // เพิ่ม BloodGroup
		
	)

	// สร้างข้อมูลเริ่มต้นในแต่ละตาราง
	GenderMale := entity.Gender{GenderName: "Male"}
	GenderFemale := entity.Gender{GenderName: "Female"}
	db.FirstOrCreate(&GenderMale, &entity.Gender{GenderName: "Male"})
	db.FirstOrCreate(&GenderFemale, &entity.Gender{GenderName: "Female"})

	PositionDoctor := entity.Position{PositionName: "Doctor"}
PositionNurse := entity.Position{PositionName: "Nurse"}
PositionFinance := entity.Position{PositionName: "Finance Staff"}
PositionNurseCounter := entity.Position{PositionName: "Nurse counter"}
PositionAdmin := entity.Position{PositionName: "Admin"}
PositionPharmacy := entity.Position{PositionName: "Pharmacy"}

// ใช้ db.FirstOrCreate เพื่อป้องกันข้อมูลซ้ำ
db.FirstOrCreate(&PositionDoctor, &entity.Position{PositionName: "Doctor"})
db.FirstOrCreate(&PositionNurse, &entity.Position{PositionName: "Nurse"})
db.FirstOrCreate(&PositionFinance, &entity.Position{PositionName: "Finance Staff"})
db.FirstOrCreate(&PositionNurseCounter, &entity.Position{PositionName: "Nurse counter"})
db.FirstOrCreate(&PositionAdmin, &entity.Position{PositionName: "Admin"})
db.FirstOrCreate(&PositionPharmacy, &entity.Position{PositionName: "Pharmacy"})

	// สร้าง Department ตัวอย่าง
Department0 := entity.Department{DepartmentName: "None"}
Department1 := entity.Department{DepartmentName: "Emergency"}
Department2 := entity.Department{DepartmentName: "Pediatrics"}
Department3 := entity.Department{DepartmentName: "Cardiology"}
Department4 := entity.Department{DepartmentName: "Neurology"}
Department5 := entity.Department{DepartmentName: "Orthopedics"}
Department6 := entity.Department{DepartmentName: "Radiology"}
Department7 := entity.Department{DepartmentName: "Pharmacy"}

db.FirstOrCreate(&Department0, &entity.Department{DepartmentName: "None"})
db.FirstOrCreate(&Department1, &entity.Department{DepartmentName: "Emergency"})
db.FirstOrCreate(&Department2, &entity.Department{DepartmentName: "Pediatrics"})
db.FirstOrCreate(&Department3, &entity.Department{DepartmentName: "Cardiology"})
db.FirstOrCreate(&Department4, &entity.Department{DepartmentName: "Neurology"})
db.FirstOrCreate(&Department5, &entity.Department{DepartmentName: "Orthopedics"})
db.FirstOrCreate(&Department6, &entity.Department{DepartmentName: "Radiology"})
db.FirstOrCreate(&Department7, &entity.Department{DepartmentName: "Pharmacy"})


	// สร้าง Status ตัวอย่าง
StatusActive := entity.Status{StatusName: "Active"} // กำลังทำงานอยู่
StatusOnLeave := entity.Status{StatusName: "On Leave"} // กำลังลาพัก (เช่น ลาคลอด, ลาป่วย)
StatusSuspended := entity.Status{StatusName: "Suspended"} // ถูกพักงานชั่วคราว
StatusResigned := entity.Status{StatusName: "Resigned"} // ลาออก
StatusRetired := entity.Status{StatusName: "Retired"} // เกษียณอายุ
StatusTerminated := entity.Status{StatusName: "Terminated"} // ถูกยกเลิกสัญญาจ้าง

db.FirstOrCreate(&StatusActive, &entity.Status{StatusName: "Active"})
db.FirstOrCreate(&StatusOnLeave, &entity.Status{StatusName: "On Leave"})
db.FirstOrCreate(&StatusSuspended, &entity.Status{StatusName: "Suspended"})
db.FirstOrCreate(&StatusResigned, &entity.Status{StatusName: "Resigned"})
db.FirstOrCreate(&StatusRetired, &entity.Status{StatusName: "Retired"})
db.FirstOrCreate(&StatusTerminated, &entity.Status{StatusName: "Terminated"})

	// สร้าง Specialist ตัวอย่าง
SpecialistNone := entity.Specialist{SpecialistName: "None"} // หัวใจ
SpecialistCardiology := entity.Specialist{SpecialistName: "Cardiology"} // หัวใจ
SpecialistNeurology := entity.Specialist{SpecialistName: "Neurology"} // ประสาทวิทยา
SpecialistPediatrics := entity.Specialist{SpecialistName: "Pediatrics"} // กุมารเวชศาสตร์
SpecialistOrthopedics := entity.Specialist{SpecialistName: "Orthopedics"} // กระดูกและข้อ
SpecialistRadiology := entity.Specialist{SpecialistName: "Radiology"} // รังสีวิทยา
SpecialistDermatology := entity.Specialist{SpecialistName: "Dermatology"} // ผิวหนัง
SpecialistOncology := entity.Specialist{SpecialistName: "Oncology"} // มะเร็งวิทยา
SpecialistGynecology := entity.Specialist{SpecialistName: "Gynecology"} // นรีเวชศาสตร์
SpecialistOphthalmology := entity.Specialist{SpecialistName: "Ophthalmology"} // จักษุวิทยา
SpecialistPsychiatry := entity.Specialist{SpecialistName: "Psychiatry"} // จิตเวชศาสตร์

db.FirstOrCreate(&SpecialistNone, &entity.Specialist{SpecialistName: "None"})
db.FirstOrCreate(&SpecialistCardiology, &entity.Specialist{SpecialistName: "Cardiology"})
db.FirstOrCreate(&SpecialistNeurology, &entity.Specialist{SpecialistName: "Neurology"})
db.FirstOrCreate(&SpecialistPediatrics, &entity.Specialist{SpecialistName: "Pediatrics"})
db.FirstOrCreate(&SpecialistOrthopedics, &entity.Specialist{SpecialistName: "Orthopedics"})
db.FirstOrCreate(&SpecialistRadiology, &entity.Specialist{SpecialistName: "Radiology"})
db.FirstOrCreate(&SpecialistDermatology, &entity.Specialist{SpecialistName: "Dermatology"})
db.FirstOrCreate(&SpecialistOncology, &entity.Specialist{SpecialistName: "Oncology"})
db.FirstOrCreate(&SpecialistGynecology, &entity.Specialist{SpecialistName: "Gynecology"})
db.FirstOrCreate(&SpecialistOphthalmology, &entity.Specialist{SpecialistName: "Ophthalmology"})
db.FirstOrCreate(&SpecialistPsychiatry, &entity.Specialist{SpecialistName: "Psychiatry"})

	// สร้าง Disease ตัวอย่าง
Disease0 := entity.Disease{DiseaseName: "None"}
Disease1 := entity.Disease{DiseaseName: "Hypertension"}
Disease2 := entity.Disease{DiseaseName: "Diabetes"}
Disease3 := entity.Disease{DiseaseName: "Asthma"}
Disease4 := entity.Disease{DiseaseName: "Tuberculosis"}
Disease5 := entity.Disease{DiseaseName: "HIV/AIDS"}
Disease6 := entity.Disease{DiseaseName: "Cancer"}
Disease7 := entity.Disease{DiseaseName: "Chronic Kidney Disease"}
Disease8 := entity.Disease{DiseaseName: "Heart Disease"}
Disease9 := entity.Disease{DiseaseName: "Stroke"}
Disease10 := entity.Disease{DiseaseName: "Alzheimer's Disease"}
Disease11 := entity.Disease{DiseaseName: "Parkinson's Disease"}
Disease12 := entity.Disease{DiseaseName: "Pneumonia"}
Disease13 := entity.Disease{DiseaseName: "Dengue Fever"}
Disease14 := entity.Disease{DiseaseName: "Malaria"}
Disease15 := entity.Disease{DiseaseName: "COVID-19"}
Disease16 := entity.Disease{DiseaseName: "Hepatitis B"}
Disease17 := entity.Disease{DiseaseName: "Hepatitis C"}
Disease18 := entity.Disease{DiseaseName: "Arthritis"}
Disease19 := entity.Disease{DiseaseName: "Migraine"}
Disease20 := entity.Disease{DiseaseName: "Obesity"}

db.FirstOrCreate(&Disease0, &entity.Disease{DiseaseName: "None"})
db.FirstOrCreate(&Disease1, &entity.Disease{DiseaseName: "Hypertension"})
db.FirstOrCreate(&Disease2, &entity.Disease{DiseaseName: "Diabetes"})
db.FirstOrCreate(&Disease3, &entity.Disease{DiseaseName: "Asthma"})
db.FirstOrCreate(&Disease4, &entity.Disease{DiseaseName: "Tuberculosis"})
db.FirstOrCreate(&Disease5, &entity.Disease{DiseaseName: "HIV/AIDS"})
db.FirstOrCreate(&Disease6, &entity.Disease{DiseaseName: "Cancer"})
db.FirstOrCreate(&Disease7, &entity.Disease{DiseaseName: "Chronic Kidney Disease"})
db.FirstOrCreate(&Disease8, &entity.Disease{DiseaseName: "Heart Disease"})
db.FirstOrCreate(&Disease9, &entity.Disease{DiseaseName: "Stroke"})
db.FirstOrCreate(&Disease10, &entity.Disease{DiseaseName: "Alzheimer's Disease"})
db.FirstOrCreate(&Disease11, &entity.Disease{DiseaseName: "Parkinson's Disease"})
db.FirstOrCreate(&Disease12, &entity.Disease{DiseaseName: "Pneumonia"})
db.FirstOrCreate(&Disease13, &entity.Disease{DiseaseName: "Dengue Fever"})
db.FirstOrCreate(&Disease14, &entity.Disease{DiseaseName: "Malaria"})
db.FirstOrCreate(&Disease15, &entity.Disease{DiseaseName: "COVID-19"})
db.FirstOrCreate(&Disease16, &entity.Disease{DiseaseName: "Hepatitis B"})
db.FirstOrCreate(&Disease17, &entity.Disease{DiseaseName: "Hepatitis C"})
db.FirstOrCreate(&Disease18, &entity.Disease{DiseaseName: "Arthritis"})
db.FirstOrCreate(&Disease19, &entity.Disease{DiseaseName: "Migraine"})
db.FirstOrCreate(&Disease20, &entity.Disease{DiseaseName: "Obesity"})

	// สร้าง BloodGroup ตัวอย่าง
	BloodGroupA := entity.BloodGroup{BloodGroup: "A+"}
	BloodGroupA2 := entity.BloodGroup{BloodGroup: "A-"}
	BloodGroupB := entity.BloodGroup{BloodGroup: "B+"}
	BloodGroupB2 := entity.BloodGroup{BloodGroup: "B-"}
	BloodGroupO := entity.BloodGroup{BloodGroup: "O+"}
	BloodGroupO2 := entity.BloodGroup{BloodGroup: "O-"}
	BloodGroupAB := entity.BloodGroup{BloodGroup: "AB+"}
	BloodGroupAB2 := entity.BloodGroup{BloodGroup: "AB-"}
	
	db.FirstOrCreate(&BloodGroupA, &entity.BloodGroup{BloodGroup: "A+"})
	db.FirstOrCreate(&BloodGroupB, &entity.BloodGroup{BloodGroup: "B+"})
	db.FirstOrCreate(&BloodGroupAB, &entity.BloodGroup{BloodGroup: "AB+"})
	db.FirstOrCreate(&BloodGroupO, &entity.BloodGroup{BloodGroup: "O+"})
	db.FirstOrCreate(&BloodGroupA2, &entity.BloodGroup{BloodGroup: "A-"})
	db.FirstOrCreate(&BloodGroupB2, &entity.BloodGroup{BloodGroup: "B-"})
	db.FirstOrCreate(&BloodGroupAB2, &entity.BloodGroup{BloodGroup: "AB-"})
	db.FirstOrCreate(&BloodGroupO2, &entity.BloodGroup{BloodGroup: "O-"})
	// เข้ารหัสรหัสผ่าน
	hashedPassword, _ := HashPassword("123456")

	// สร้าง Employee พร้อมข้อมูลเริ่มต้น โดยเว้น profile ไว้เป็น null
EmployeeDoctor := entity.Employee{
	FirstName:           "John",
	LastName:            "Doe",
	Email:               "john.doe@hospital.com",
	DateOfBirth:         time.Date(1980, time.March, 15, 0, 0, 0, 0, time.UTC),
	Phone:               "111-111-1111",
	Address:             "123 Doctor Street",
	NationalID:          "1234567890123",
	Username:            "doctor",
	ProfessionalLicense: "DOC12345",
	Graduate:            "MD from XYZ University",
	Password:            hashedPassword,
	GenderID:            GenderMale.ID,
	PositionID:          PositionDoctor.ID,
	DepartmentID:        Department1.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistCardiology.ID,
	BloodGroupID:        BloodGroupA.ID, // ตั้งค่า BloodGroup เป็น A
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeeDoctor, &entity.Employee{Email: "john.doe@hospital.com"})

// เพิ่มความสัมพันธ์ Many-to-Many กับ Disease
DiseasesDoctor := []entity.Disease{Disease1, Disease2}
db.Model(&EmployeeDoctor).Association("Diseases").Append(DiseasesDoctor)

EmployeeNurse := entity.Employee{
	FirstName:           "Jane",
	LastName:            "Smith",
	Email:               "jane.smith@hospital.com",
	DateOfBirth:         time.Date(1990, time.July, 20, 0, 0, 0, 0, time.UTC),
	Phone:               "222-222-2222",
	Address:             "456 Nurse Lane",
	NationalID:          "2234567890123",
	Username:            "nurse",
	ProfessionalLicense: "NUR56789",
	Graduate:            "BSc Nursing from ABC University",
	Password:            hashedPassword,
	GenderID:            GenderFemale.ID,
	PositionID:          PositionNurse.ID,
	DepartmentID:        Department2.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistPediatrics.ID,
	BloodGroupID:        BloodGroupB.ID, // ตั้งค่า BloodGroup เป็น B
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeeNurse, &entity.Employee{Email: "jane.smith@hospital.com"})

DiseasesNurse := []entity.Disease{Disease3, Disease4}
db.Model(&EmployeeNurse).Association("Diseases").Append(DiseasesNurse)

EmployeeFinance := entity.Employee{
	FirstName:           "Alice",
	LastName:            "Johnson",
	Email:               "alice.johnson@hospital.com",
	DateOfBirth:         time.Date(1985, time.January, 10, 0, 0, 0, 0, time.UTC),
	Phone:               "333-333-3333",
	Address:             "789 Finance Blvd",
	NationalID:          "3234567890123",
	Username:            "finance",
	ProfessionalLicense: "",
	Graduate:            "MBA from DEF University",
	Password:            hashedPassword,
	GenderID:            GenderFemale.ID,
	PositionID:          PositionFinance.ID,
	DepartmentID:        Department3.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistRadiology.ID,
	BloodGroupID:        BloodGroupAB.ID, // ตั้งค่า BloodGroup เป็น AB
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeeFinance, &entity.Employee{Email: "alice.johnson@hospital.com"})

DiseasesFinance := []entity.Disease{Disease5}
db.Model(&EmployeeFinance).Association("Diseases").Append(DiseasesFinance)

EmployeeNurseCounter := entity.Employee{
	FirstName:           "Michael",
	LastName:            "Brown",
	Email:               "michael.brown@hospital.com",
	DateOfBirth:         time.Date(1992, time.April, 5, 0, 0, 0, 0, time.UTC),
	Phone:               "444-444-4444",
	Address:             "101 Nurse Counter Ave",
	NationalID:          "4234567890123",
	Username:            "counter",
	ProfessionalLicense: "",
	Graduate:            "Diploma in Healthcare",
	Password:            hashedPassword,
	GenderID:            GenderMale.ID,
	PositionID:          PositionNurseCounter.ID,
	DepartmentID:        Department4.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistNeurology.ID,
	BloodGroupID:        BloodGroupO.ID, // ตั้งค่า BloodGroup เป็น O
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeeNurseCounter, &entity.Employee{Email: "michael.brown@hospital.com"})

DiseasesNurseCounter := []entity.Disease{Disease6}
db.Model(&EmployeeNurseCounter).Association("Diseases").Append(DiseasesNurseCounter)

EmployeeAdmin := entity.Employee{
	FirstName:           "Emma",
	LastName:            "Davis",
	Email:               "emma.davis@hospital.com",
	DateOfBirth:         time.Date(1988, time.November, 12, 0, 0, 0, 0, time.UTC),
	Phone:               "555-555-5555",
	Address:             "202 Admin Plaza",
	NationalID:          "5234567890123",
	Username:            "admin",
	ProfessionalLicense: "",
	Graduate:            "BA in Administration",
	Password:            hashedPassword,
	GenderID:            GenderFemale.ID,
	PositionID:          PositionAdmin.ID,
	DepartmentID:        Department5.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistDermatology.ID,
	BloodGroupID:        BloodGroupA.ID, // ตั้งค่า BloodGroup เป็น A
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeeAdmin, &entity.Employee{Email: "emma.davis@hospital.com"})

DiseasesAdmin := []entity.Disease{Disease7}
db.Model(&EmployeeAdmin).Association("Diseases").Append(DiseasesAdmin)

EmployeePharmacy := entity.Employee{
	FirstName:           "Oliver",
	LastName:            "Wilson",
	Email:               "oliver.wilson@hospital.com",
	DateOfBirth:         time.Date(1987, time.December, 25, 0, 0, 0, 0, time.UTC),
	Phone:               "666-666-6666",
	Address:             "303 Pharmacy Court",
	NationalID:          "6234567890123",
	Username:            "pharmacy",
	ProfessionalLicense: "PHA67890",
	Graduate:            "Pharmacy Degree from GHI University",
	Password:            hashedPassword,
	GenderID:            GenderMale.ID,
	PositionID:          PositionPharmacy.ID,
	DepartmentID:        Department6.ID,
	StatusID:            StatusActive.ID,
	SpecialistID:        SpecialistOncology.ID,
	BloodGroupID:        BloodGroupB.ID, // ตั้งค่า BloodGroup เป็น B
	Profile:             "", // ตั้งค่า Profile ให้เป็นค่า null
}
db.FirstOrCreate(&EmployeePharmacy, &entity.Employee{Email: "oliver.wilson@hospital.com"})

DiseasesPharmacy := []entity.Disease{Disease8}
db.Model(&EmployeePharmacy).Association("Diseases").Append(DiseasesPharmacy)
}

// // tae ----------------------------------------------------------------
// // ตัวอย่างข้อมูลสำหรับ TypeRoom
// singleRoom := entity.RoomType{
// 	RoomName: "ห้อง ICU",
// 	PricePerNight: 3000.00,
// }
// doubleRoom := entity.RoomType{
// 	RoomName: "ห้องเดี่ยว",
// 	PricePerNight: 10000.00,
	
// }
// suiteRoom := entity.RoomType{
// 	RoomName: "ห้องรวม",
// 	PricePerNight: 15000.00,
// }
// db.FirstOrCreate(&singleRoom,entity.RoomType{RoomName: "ห้อง ICU"})
// db.FirstOrCreate(&doubleRoom,entity.RoomType{RoomName: "ห้องเดี่ยว"})
// db.FirstOrCreate(&suiteRoom,entity.RoomType{RoomName: "ห้องรวม"})

// // ตัวอย่างข้อมูลสำหรับ Building 
// buildingA := entity.Building{
// 	BuildingName: "อาคาร A",
// 	Location:     "มทส.",
// }
// buildingB := entity.Building{
// 	BuildingName: "อาคาร B",
// 	Location:     "มทส.",
// }
// db.FirstOrCreate(&buildingA,entity.Building{BuildingName: "อาคาร A"})
// db.FirstOrCreate(&buildingB,entity.Building{BuildingName: "อาคาร B"})

// // ตัวอย่างข้อมูลสำหรับ Floor
// floor1BA := entity.Floor{
// 	FloorNumber:     "1",
// 	BuildingID: buildingA.ID,
// }
// floor2BA := entity.Floor{
// 	FloorNumber:     "2",
// 	BuildingID: buildingA.ID,
// }
// floor1BB := entity.Floor{
// 	FloorNumber:     "1",
// 	BuildingID: buildingB.ID,
// }
// db.FirstOrCreate(&floor1BA,entity.Floor{FloorNumber: "1"})
// db.FirstOrCreate(&floor2BA,entity.Floor{FloorNumber: "2"})
// db.FirstOrCreate(&floor1BB,entity.Floor{FloorNumber: "1",BuildingID: buildingB.ID})

// // ตัวอย่างข้อมูลสำหรับ Room
// room101 := entity.Room{
// 	RoomNumber:  "101",
// 	RoomTypeID:  singleRoom.ID,
// 	//Status:      "Available",
// 	BedCapacity: 1,
// 	EmployeeID: EmployeeNurseCounter.ID,

// }
// room102 := entity.Room{
// 	RoomNumber:  "102",
// 	RoomTypeID:  doubleRoom.ID,
// 	//Status:      "Occupied",
// 	BedCapacity: 2,
// 	EmployeeID: EmployeeNurseCounter.ID,
// }
// room201 := entity.Room{
// 	RoomNumber:  "201",
// 	RoomTypeID:  suiteRoom.ID,
// 	//Status:      "Available",
// 	BedCapacity: 3,
// 	EmployeeID: EmployeeNurseCounter.ID,
// }
// db.FirstOrCreate(&room101,entity.Room{RoomNumber: "101"})
// db.FirstOrCreate(&room102,entity.Room{RoomNumber: "102"})
// db.FirstOrCreate(&room201,entity.Room{RoomNumber: "201"})


// // time zone Thailand
// loc, _ := time.LoadLocation("Asia/Bangkok")
// admissionDate := time.Now().In(loc)
// // outroom +7 day
// dischargeDate := admissionDate.AddDate(0, 0, 7)


// PatientRoom01:= entity.PatientRoom{
// 	PatientID :"สมหมาย",
// 	RoomID: room101.ID,
// 	AdmissionDate: admissionDate,
// 	DischargeDate: dischargeDate,
// 	Status: "Occupied",
// }

// PatientRoom02:= entity.PatientRoom{
// 	PatientID :"สมสัก",
// 	RoomID: room102.ID,
// 	AdmissionDate: admissionDate,
// 	DischargeDate: dischargeDate,
// 	Status: "Vacant",
// }
// db.FirstOrCreate(&PatientRoom01,entity.PatientRoom{PatientID :"สมหมาย",})
// db.FirstOrCreate(&PatientRoom02,entity.PatientRoom{PatientID :"สมสัก",})

// BookRoom01 := entity.RoomLayout{
// 	BuildingID: buildingA.ID,
// 	RoomID: PatientRoom01.ID,
// 	FloorID: floor1BA.ID,
// 	PositionX: 1,
// 	PositionY: 1,
// }
// BookRoom02 := entity.RoomLayout{
// 	BuildingID: buildingA.ID,
// 	RoomID: PatientRoom02.ID,
// 	FloorID: floor2BA.ID,
// 	PositionX: 1,
// 	PositionY: 2,
// }
// db.FirstOrCreate(&BookRoom01,entity.RoomLayout{RoomID: PatientRoom01.ID})
// db.FirstOrCreate(&BookRoom02,entity.RoomLayout{RoomID: PatientRoom02.ID})


// fmt.Println("Database setup completed.") 


// //ทัชเขียน
// schedule0 := entity.Schedule{ScheduleName: "นัด"}
// 	schedule1 := entity.Schedule{ScheduleName: "ไม่นัด"}

// 	db.FirstOrCreate(&schedule0, &entity.Schedule{ScheduleName: "นัด"})
// 	db.FirstOrCreate(&schedule1, &entity.Schedule{ScheduleName: "ไม่นัด"})

// 	severity1 := entity.Severity{
// 		SeverityLevel: 10,
// 		SeverityName:  "เล็กน้อยมาก",
// 	}
// 	severity2 := entity.Severity{
// 		SeverityLevel: 20,
// 		SeverityName:  "เล็กน้อย",
// 	}
// 	severity3 := entity.Severity{
// 		SeverityLevel: 30,
// 		SeverityName:  "เล็กน้อยถึงปานกลาง",
// 	}
// 	severity4 := entity.Severity{
// 		SeverityLevel: 40,
// 		SeverityName:  "ปานกลาง",
// 	}
// 	severity5 := entity.Severity{
// 		SeverityLevel: 50,
// 		SeverityName:  "ปานกลางถึงรุนแรง",
// 	}
// 	severity6 := entity.Severity{
// 		SeverityLevel: 60,
// 		SeverityName:  "รุนแรง",
// 	}
// 	severity7 := entity.Severity{
// 		SeverityLevel: 70,
// 		SeverityName:  "รุนแรงมาก",
// 	}
// 	severity8 := entity.Severity{
// 		SeverityLevel: 80,
// 		SeverityName:  "อันตราย",
// 	}
// 	severity9 := entity.Severity{
// 		SeverityLevel: 90,
// 		SeverityName:  "อันตรายถึงขั้นวิกฤต",
// 	}
// 	severity10 := entity.Severity{
// 		SeverityLevel: 100,
// 		SeverityName:  "ขั้นวิกฤตถึงเสียชีวิต",
// 	}

// 	db.FirstOrCreate(&severity1, &entity.Severity{SeverityName: "เล็กน้อยมาก"})
// 	db.FirstOrCreate(&severity2, &entity.Severity{SeverityName: "เล็กน้อย"})
// 	db.FirstOrCreate(&severity3, &entity.Severity{SeverityName: "เล็กน้อยถึงปานกลาง"})
// 	db.FirstOrCreate(&severity4, &entity.Severity{SeverityName: "ปานกลาง"})
// 	db.FirstOrCreate(&severity5, &entity.Severity{SeverityName: "ปานกลางถึงรุนแรง"})
// 	db.FirstOrCreate(&severity6, &entity.Severity{SeverityName: "รุนแรง"})
// 	db.FirstOrCreate(&severity7, &entity.Severity{SeverityName: "รุนแรงมาก"})
// 	db.FirstOrCreate(&severity8, &entity.Severity{SeverityName: "อันตราย"})
// 	db.FirstOrCreate(&severity9, &entity.Severity{SeverityName: "อันตรายถึงขั้นวิกฤต"})
// 	db.FirstOrCreate(&severity10, &entity.Severity{SeverityName: "ขั้นวิกฤตถึงเสียชีวิต"})

// 	patient1 := entity.Patient{
// 		NationID:"1234567891011",
// 		FirstName : "สมหญิง",
// 		LastName:"จริงใจ",
// 		DateOfBirth:time.Date(1999, time.March, 22, 0, 0, 0, 0, time.UTC),
// 		Address:"กรุงเทพ",
// 		PhoneNumber:"0601234567",
// 		GenderID:GenderFemale.ID,
// 		BloodGroupID:BloodGroupA.ID,
// 	}
// 	db.FirstOrCreate(&patient1, &entity.Patient{NationID: "1234567891011"})

// 	Diseasespatient1 := []entity.Disease{Disease8}
// 	db.Model(&patient1).Association("Diseases").Append(Diseasespatient1)

// 	patient2 := entity.Patient{
// 		NationID:"1234567891012",
// 		FirstName : "สมชาย",
// 		LastName:"สายสะอาด",
// 		DateOfBirth:time.Date(1990, time.April, 20, 0, 0, 0, 0, time.UTC),
// 		Address:"นครพนม",
// 		PhoneNumber:"0612345678",
// 		GenderID:GenderMale.ID,
// 		BloodGroupID:BloodGroupAB.ID,
// 	}
// 	db.FirstOrCreate(&patient2, &entity.Patient{NationID: "1234567891012"})

// 	Diseasespatient2 := []entity.Disease{Disease1,Disease10}
// 	db.Model(&patient2).Association("Diseases").Append(Diseasespatient2)

// 	patient3 := entity.Patient{
// 		NationID:"1234567891013",
// 		FirstName : "อับดุล",
// 		LastName:"สาโท",
// 		DateOfBirth:time.Date(1980, time.August, 1, 0, 0, 0, 0, time.UTC),
// 		Address:"นครพนม",
// 		PhoneNumber:"0699999999",
// 		GenderID:GenderMale.ID,
// 		BloodGroupID:BloodGroupA.ID,
// 	}
// 	db.FirstOrCreate(&patient3, &entity.Patient{NationID: "1234567891013"})

// 	Diseasespatient3 := []entity.Disease{Disease4,Disease6}
// 	db.Model(&patient3).Association("Diseases").Append(Diseasespatient3)

// 	MedicalRecords1 := entity.MedicalRecords{
// 		SymptomsDetails:"มีอาการไอแห้งอย่างต่อเนื่องเป็นเวลา 5 วันลักษณะการไอจะเกิดขึ้นถี่มากขึ้นในช่วงเวลากลางคืนทำให้นอนหลับได้ไม่เต็มที่และรู้สึกเพลียในตอนเช้า อาการไอมีลักษณะเสียงแห้ง ไม่มีเสมหะหรืออาการหายใจลำบากร่วมด้วย ผู้ป่วยแจ้งว่ารู้สึกเจ็บคอทุกครั้งหลังไอ โดยเฉพาะบริเวณลำคอส่วนหน้า ร่วมกับมีเสียงแหบซึ่งรบกวนการพูดคุยในชีวิตประจำวัน โดยเมื่อใช้เสียงเป็นเวลานานหรือพูดดัง อาการเสียงแหบจะรุนแรงขึ้น ผู้ป่วยยังระบุว่ารู้สึกระคายคออยู่ตลอดเวลา เหมือนมีสิ่งแปลกปลอมติดอยู่ในลำคอแต่ไม่มีอากากลืนลำบากหรือกลืนเจ็บไม่มีไข้ไม่มีอาการน้ำมูกไหลหรือแน่นหน้าอกผู้ป่วยเคยลองใช้ยาอมบรรเทาอาการเจ็บคอและจิบน้ำอุ่นบ่อย ๆ แต่อาการดีขึ้นเพียงเล็กน้อย",
// 		CheckResults:"อุณหภูมิร่างกาย 36.8°C ไม่มีอาการผิดปกติในระบบหัวใจและปอด ตรวจปอดทั้งสองข้างไม่มีเสียงผิดปกติ (No crackles or wheezing) ตรวจลำคอพบว่าบริเวณผนังคอด้านหลังมีลักษณะแดงเล็กน้อย แต่ไม่มีฝ้าขาว ไม่พบการบวมโตของต่อมทอนซิล ต่อมน้ำเหลืองบริเวณคอไม่โต หายใจปกติไม่มีอาการเหนื่อยหอบ",
// 		Diagnose:"อาจเกิดจากการติดเชื้อไวรัสในระบบทางเดินหายใจส่วนบน หรืออาจเกิดจากการใช้เสียงมากเกินไปในระยะหลัง โดยผู้ป่วยมีอาการไอแห้งติดต่อกันหลายวัน ไม่มีเสมหะ, เสียงแหบที่รุนแรงขึ้นหลังการใช้เสียงนานๆ, และการระคายเคืองคอซึ่งเป็นลักษณะเฉพาะของการอักเสบที่กล่องเสียง",
// 		OrderMedicine:"พาราเซตามอล 500 มิลลิกรัม 30 เม็ด ไอบูโพรเฟน 200 มิลลิกรัม 30 เม็ด",
// 		Instructions:"ให้ผู้ป่วยพักเสียง หลีกเลี่ยงการพูดดังหรือการพูดต่อเนื่องเป็นเวลานาน ควรดื่มน้ำอุ่นเพื่อบรรเทาอาการระคายคอ หลีกเลี่ยงอาหารที่เย็นจัดหรือรสจัด งดสูบบุหรี่หรืออยู่ในพื้นที่ที่มีควันหรือฝุ่นละอองจำนวนมาก หากอาการไม่ดีขึ้นใน 7 วัน หรือมีไข้หรืออาการอื่น เช่น เสียงแหบเรื้อรัง เจ็บคอมากขึ้น หรือเหนื่อยง่ายผิดปกติ ให้กลับมาพบแพทย์เพื่อการประเมินซ้ำ",
// 		Date:time.Now(),
// 		Price:500.00,
// 		SeverityID:severity2.ID,
// 		ScheduleID:schedule1.ID,
// 		EmployeeID:EmployeeDoctor.ID,
// 	}
// 	db.FirstOrCreate(&MedicalRecords1, &entity.MedicalRecords{OrderMedicine: "พาราเซตามอล 500 มิลลิกรัม 30 เม็ด ไอบูโพรเฟน 200 มิลลิกรัม 30 เม็ด"})

// 	DiseasesMedicalRecords1 := []entity.Disease{Disease6,Disease16}
// 	db.Model(&MedicalRecords1).Association("Diseases").Append(DiseasesMedicalRecords1)

// 	TakeAHistory1 := entity.TakeAHistory{
// 		Weight:45.00,
// 		Hight:150.00,
// 		PreliminarySymptoms: "รู้สึกเจ็บหน้าอกเวลาหายใจลึกๆเริ่มประมาณ3วันที่แล้วรู้สึกเจ็บตลอดเวลาแต่จะเจ็บมากขึ้นเวลาหายใจเข้าลึกๆหรือเวลาขยับตัวแรงๆ",
// 		SystolicBloodPressure:130,
// 		DiastolicBloodPressure:65,
// 		PulseRate:170,
// 		Smoking:"ใช่",
// 		DrinkAlcohol:"ใช่",
// 		LastMenstruationDate:time.Date(2024, time.November,10, 0, 0, 0, 0, time.UTC),
// 		Date:time.Now(),
// 		MedicalRecordsID:nil,
// 		PatientID:patient1.ID,
// 		EmployeeID:EmployeeNurse.ID,
// 	}
// 	db.FirstOrCreate(&TakeAHistory1, &entity.TakeAHistory{PreliminarySymptoms: "รู้สึกเจ็บหน้าอกเวลาหายใจลึกๆเริ่มประมาณ3วันที่แล้วรู้สึกเจ็บตลอดเวลาแต่จะเจ็บมากขึ้นเวลาหายใจเข้าลึกๆหรือเวลาขยับตัวแรงๆ"})

// 	TakeAHistory2 := entity.TakeAHistory{
// 		Weight:65.00,
// 		Hight:171.00,
// 		PreliminarySymptoms: "รู้สึกมีไข้ต่ำๆ มีอาการคัดจมูกหรือมีน้ำมูกใสไหลออกมาเป็นจำนวนมาก สึกไม่สบายตัวและหายใจลำบาก มักจะจามบ่อย ๆ โดยเฉพาะในช่วงเช้าหรือเมื่อเจออากาศเย็นๆ ",
// 		SystolicBloodPressure:120,
// 		DiastolicBloodPressure:70,
// 		PulseRate:172,
// 		Smoking:"ไม่",
// 		DrinkAlcohol:"ใช่",
// 		LastMenstruationDate:time.Time{},
// 		Date:time.Now(),
// 		MedicalRecordsID:&MedicalRecords1.ID,
// 		PatientID:patient2.ID,
// 		EmployeeID:EmployeeNurse.ID,
// 	}
// 	db.FirstOrCreate(&TakeAHistory2, &entity.TakeAHistory{PreliminarySymptoms: "รู้สึกมีไข้ต่ำๆ มีอาการคัดจมูกหรือมีน้ำมูกใสไหลออกมาเป็นจำนวนมาก สึกไม่สบายตัวและหายใจลำบาก มักจะจามบ่อย ๆ โดยเฉพาะในช่วงเช้าหรือเมื่อเจออากาศเย็นๆ "})

