package unit

import (
	"SE-B6527075/entity"
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestWeight(t *testing.T) {

	g := NewGomegaWithT(t)

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

		g.Expect(err.Error()).To(Equal("Username is required"))
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

		g.Expect(err.Error()).To(Equal("Username is required"))
	})

	


	t.Run(`weight is valid`, func(t *testing.T) {
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
func TestEmail(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run(`email is invalid`, func(t *testing.T) {
		member := entity.Member{
			Username: "Peet", 
			Password: "12345678", 
			Email:     "5555.go",  // ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is invalid"))

	})

	t.Run(`email is required`, func(t *testing.T) {
		member := entity.Member{ 
			Username: "Preem", 
			Password: "56781234", 
			Email:     "",  // ผิดตรงนี้
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Email is required"))
	})

	t.Run(`email is valid`, func(t *testing.T) {
		member := entity.Member{
			Username: "Patcharida", 
			Password: "23456789",
			Email:     "test@gmail.com",  // ถูกแล้ว
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())

	})	
}