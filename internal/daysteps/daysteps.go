package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"track5sprint/internal/personaldata"
	"track5sprint/internal/spentenergy"
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	Steps      int
	Duration   time.Duration
	Personal   personaldata.Personal
	TotalSteps int
}

// создайте метод Parse()
func (ds *DaySteps) Parse(data string) error {
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return fmt.Errorf("incorrect data format: whant 2, have %d", len(parts))
	}

	// Преобразование количества шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("can't convert string steps to integer: %w", err)
	}
	ds.Steps = steps
	ds.TotalSteps += steps

	// Преобразование продолжительности
	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ant't parse duration to time.Duration: %w", err)
	}
	ds.Duration = dur

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	// Проверка продолжительности
	if ds.Duration <= 0 {
		return "", fmt.Errorf("duration must be over  0")
	}

	// Вычисление дистанции
	distance := spentenergy.Distance(ds.Steps)

	// Вычисление количества сожженных калорий
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	// Формирование строки с информацией
	info := fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.",
		ds.Steps,
		distance,
		calories,
	)

	return info, nil
}
func (ds DaySteps) Print() {
	fmt.Printf("Количество шагов: %d\nДлительность: %s\n", ds.Steps, ds.Duration)
}
