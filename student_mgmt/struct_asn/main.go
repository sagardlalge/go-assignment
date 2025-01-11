package main

import "fmt"

type stud_data struct {
	Name  string
	Id    string
	Marks float64
}

type marks struct {
	English int
	Hindi   int
	Math    int
	Science int
	CS      int
}

type student struct {
	student_details stud_data
	student_marks   marks
}

func main() {
	var s student

	fmt.Println("\nEnter the student details and marks for the following subjects:")

	fmt.Print("Name: ")
	fmt.Scan(&s.student_details.Name)

	fmt.Print("ID : ")
	fmt.Scan(&s.student_details.Id)

	fmt.Print("English: ")
	fmt.Scan(&s.student_marks.English)

	fmt.Print("Hindi: ")
	fmt.Scan(&s.student_marks.Hindi)

	fmt.Print("Math: ")
	fmt.Scan(&s.student_marks.Math)

	fmt.Print("Science: ")
	fmt.Scan(&s.student_marks.Science)

	fmt.Print("Computer Science (CS): ")
	fmt.Scan(&s.student_marks.CS)

	fmt.Println("\n Student Details: ", s.student_details, "\nUpdated Student Marks:", s.student_marks)
}
