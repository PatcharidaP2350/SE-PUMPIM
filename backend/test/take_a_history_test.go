package test

import (
	"SE-B6527075/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestWeight(t *testing.T) {

	g := NewGomegaWithT(t)

	// bloodgroup := entity.BloodGroup{
	// 	BloodGroup:"A+",
	// }
	// employee := entity.Employee{
	// 	FirstName:           "John",
	// 	LastName:            "Doe",
	// 	DateOfBirth:         time.Now().AddDate(-30, 0, 0),
	// 	Email:               "john.doe2@example.com",
	// 	Phone:               "0812345678",
	// 	Address:             "123 Main St",
	// 	Username:            "johndoe",
	// 	ProfessionalLicense: "12345",
	// 	Graduate:            "Bachelor's Degree",
	// 	NationalID:          "1234567890123",
	// 	InfoConfirm:         true,
	// 	FeedbackMessage:     "Great work!",
	// 	StatusExpiration:    time.Now().AddDate(0, 6, 0),
	// 	GenderID:            1,
	// 	PositionID:          1,
	// 	DepartmentID:        1,
	// 	StatusID:            1,
	// 	SpecialistID:        1,
	// 	BloodGroupID:        1,
	// 	Diseases:            []entity.Disease{}, // Empty diseases list
	// 	Profile:             "Experienced developer",
	// 	Password:            "securepassword",
	// 	ResetToken:          "123e4567-e89b-12d3-a456-426614174000",
	// 	ResetTokenExpiry:    time.Now().Add(24 * time.Hour),
	// 	Status: 			 status,
	// 	Gender:   			 gender,		
	// 	Position:   		 position,
	// 	Department: 		 department,
	// 	Specialist:  		 specialist,
	// 	BloodGroup: 		 bloodgroup,			
	// }
	// disease1 := entity.Disease{
	// 	DiseaseName: "Hypertension",
	// }
	// disease2 := entity.Disease{
	// 	DiseaseName: "Diabetes",
	// }

	t.Run(`weight is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   0,  // ผิดตรงนี้
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Weight is required"))
	})

	


	t.Run(`weight is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  // ถูกแล้ว
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}


func TestHeight(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`height is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    0,  // ผิดตรงนี้ // ส่วนสูงไม่ควรเป็น 0
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Height is required"))
	})

	


	t.Run(`height is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,  // ถูกแล้ว
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestPreliminarySymptoms(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`preliminary_symtomps is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "",    // ผิดตรงนี้
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PreliminarySymptoms is required"))
	})

	


	t.Run(`preliminary_symtomps is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  // ถูกแล้ว
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}


func TestSystolicBloodPressure(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`systolic_blood_pressure is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   0,  // ผิดตรงนี้
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SystolicBloodPressure is required"))
	})

	


	t.Run(`systolic_blood_pressure is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,  // ถูกแล้ว
			DiastolicBloodPressure:  80,
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestDiastolicBloodPressure(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`diastolic_blood_pressure is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   120,  
			DiastolicBloodPressure:  0,  // ผิดตรงนี้
			PulseRate:               70,
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DiastolicBloodPressure is required"))
	})

	


	t.Run(`diastolic_blood_pressure is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80, // ถูกแล้ว
			PulseRate:               70,  
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestPulseRate(t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`pulse_rate is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   50,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   120,  
			DiastolicBloodPressure:  80,  
			PulseRate:               0,  // ผิดตรงนี้
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("PulseRate is required"))
	})

	


	t.Run(`pulse_rate is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80, 
			PulseRate:               70,  // ถูกแล้ว
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestSmoking (t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`smoking is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   50,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   120,  
			DiastolicBloodPressure:  80,  
			PulseRate:               70,  
			// ผิดตรงนี้
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Smoking is required"))
	})

	


	t.Run(`smoking is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80, 
			PulseRate:               70,  
			Smoking:                 false,  // ถูกแล้ว
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestDrinkAlcohol (t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`drink_alcohol is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   50,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   120,  
			DiastolicBloodPressure:  80,  
			PulseRate:               70,  
			Smoking:                 false,
			// ผิดตรงนี้
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("DrinkAlcohol is required"))
	})

	


	t.Run(`drink_alcohol is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80, 
			PulseRate:               70,  
			Smoking:                 false,  
			DrinkAlcohol:            false, // ถูกแล้ว
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}

func TestLastMenstruationDate (t *testing.T) {

	g := NewGomegaWithT(t)

	t.Run(`last_menstruation_date is required`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   50,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน",   
			SystolicBloodPressure:   120,  
			DiastolicBloodPressure:  80,  
			PulseRate:               70,  
			Smoking:                 false,
			DrinkAlcohol:            false,
			LastMenstruationDate:    time.Now().AddDate(0, 1, 0), // วันที่ในอนาคต,  // ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("LastMenstruationDate is required"))
	})

	


	t.Run(`last_menstruation_date is valid`, func(t *testing.T) {
		take_a_history := entity.TakeAHistory{
			Weight:   55,  
			Height:    175.7,
			PreliminarySymptoms:  "ปวดหัว ตัวร้อน", 
			SystolicBloodPressure:   120,
			DiastolicBloodPressure:  80, 
			PulseRate:               70,  
			Smoking:                 false,  
			DrinkAlcohol:            false, 
			LastMenstruationDate:    time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC), // ถูกแล้ว
		}

		ok, err := govalidator.ValidateStruct(take_a_history)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})
}


