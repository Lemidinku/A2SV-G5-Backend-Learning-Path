package main

import (
	"fmt"
)

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Printf("Hello, %s!\n", name)

	var num_of_subjects int
    fmt.Print("How many subjects did you take? ")
    fmt.Scanln(&num_of_subjects)
	var subjects []string
	
	for i := 0; i < num_of_subjects; i++ {
		var subject_name string
		fmt.Printf("enter subject %d name: ", i+1)
		fmt.Scanln(&subject_name)
		subjects = append(subjects, subject_name)
	}

	var scores = make(map[string]float64);
	var grades = make(map[string]string);
	i :=0
	for i < num_of_subjects{
		var grade float64
		fmt.Printf("enter grade for %s: ", subjects[i])
		fmt.Scanln(&grade)
		if grade < 0 || grade > 100 {
			fmt.Println("Invalid grade. Grade should be between 0 and 100")
			continue
		} else {
			scores[subjects[i]] = grade
			grades[subjects[i]] = calculateGrade(grade)
			i++
		}
	}
	fmt.Printf("Your average grade is %.2f \n",calculateAverage(scores))
	for key, value := range grades {
		fmt.Printf("Subject: %s, Grade: %s\n", key, value)
	}
}

func calculateGrade(score float64) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func calculateAverage(scores map[string]float64) float64 {
	var sum float64 = 0
	for _, score := range scores {
		sum += score
	}
	return sum / float64(len(scores))
}


