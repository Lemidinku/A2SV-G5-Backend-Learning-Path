package main
import (
	"testing"
	"math"
)

func TestCalculateAverage(t *testing.T) {
	scores := map[string]float64{
		"Math":    80,
		"English": 90,
		"Science": 75,
	}

	expected := 81.67
	result := calculateAverage(scores)
	result = math.Round(result*100) / 100

	if result != expected {
		t.Errorf("Expected average %f, but got %f", expected, result)
	}
}

func TestCalculateAverage_EmptyScores(t *testing.T) {
	scores := map[string]float64{}

	expected := 0.0
	result := calculateAverage(scores)

	if result != expected {
		t.Errorf("Expected average %.2f, but got %.2f", expected, result)
	}
}

func TestCalculateAverage_SingleScore(t *testing.T) {
	scores := map[string]float64{
		"Math": 85,
	}

	expected := 85.0
	result := calculateAverage(scores)

	if result != expected {
		t.Errorf("Expected average %.2f, but got %.2f", expected, result)
	}
}