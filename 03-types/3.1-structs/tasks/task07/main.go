// Задача 7: Интеграционная мини-модель
//
// Ожидаемый вывод:
//   Enrollment E-001: Alice enrolled in Go Basics [active]

package main

import "fmt"

// TODO: объяви структуру ContactInfo с полями Phone (string), Email (string)

// TODO: объяви структуру Course с полями ID (string), Title (string)

// TODO: объяви структуру Student с полями Name (string) и встроенным ContactInfo (embedding)

// TODO: объяви структуру CourseEnrollment с полями:
//       EnrollmentID (string), Status (string)
//       Student Student, Course Course

type ContactInfo struct {
	Phone string
	Email string
}

type Course struct {
	ID    string
	Title string
}

type Student struct {
	Name string
	ContactInfo
}

type CourseEnrollment struct {
	EnrollmentID string
	Status       string
	Student      Student
	Course       Course
}

func main() {
	// TODO: создай значение CourseEnrollment с осмысленными данными

	e := CourseEnrollment{
		EnrollmentID: "12300",
		Student:      Student{Name: "Boris"},
		Course:       Course{Title: "GO"},
		Status:       "student",
	}

	fmt.Printf("Enrollment %s: %s enrolled in %s [%s]\n",
		e.EnrollmentID, e.Student.Name, e.Course.Title, e.Status)
}
