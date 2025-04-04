package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"track5sprint/internal/personaldata"
	"track5sprint/internal/spentenergy"
)

// создайте структуру Training
type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	Personal     personaldata.Personal
}

// создайте метод Parse()
func (t *Training) Parse(data string) error {
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		return fmt.Errorf("incorrect data format: whant 3, have %d", len(parts))
	}

	// Преобразование количества шагов
	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("can't convert string steps to integer: %w", err)
	}
	t.Steps = steps

	// Проверка типа тренировки
	trainingTypes := []string{"Бег", "Ходьба"}
	found := false
	for _, tt := range trainingTypes {
		if parts[1] == tt {
			t.TrainingType = tt
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("unknown training type: %s", parts[1])
	}

	// Преобразование продолжительности
	dur, err := time.ParseDuration(parts[2])
	if err != nil {
		return fmt.Errorf("cant't parse duration to time.Duration: %w", err)
	}
	t.Duration = dur

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	// Вычисление дистанции
	distance := spentenergy.Distance(t.Steps)

	// Проверка продолжительности
	if t.Duration <= 0 {
		return "", fmt.Errorf("duration must be over 0")
	}

	// Вычисление средней скорости
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	// Расчет калорий
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}
	if err != nil {
		return "", err
	}

	// Формирование строки с информацией о тренировке
	info := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f ккал.",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		meanSpeed,
		calories,
	)

	return info, nil
}
func (t Training) Print() {
	fmt.Printf("Тип тренировки: %s\nДлительность: %s\n", t.TrainingType, t.Duration)
}
