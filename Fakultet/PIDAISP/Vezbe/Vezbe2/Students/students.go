package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Student struct {
	index     string
	firstName string
	lastName  string
	gpa       float64
}

func main() {
	for {
		exit := chooseOption()

		if exit {
			break
		}
	}
}

func chooseOption() bool {
	fmt.Println("Choose one of the following options:")
	fmt.Println("\t- 1. Create a new student record.")
	fmt.Println("\t- 2. Student overview (GPA descending).")
	fmt.Println("\t- 3. Student overview (GPA ascending).")
	fmt.Println("\t- 4. Overview of average GPA of all students.")
	fmt.Println("\t- 5. Exit.")

	var choice int
	fmt.Scanf("%d", &choice)

	for choice < 1 || choice > 4 {
		fmt.Println("You didn't enter a valid option. Try again:")
		fmt.Scanf("%d", &choice)
	}

	var err error
	if choice == 1 {
		err = createStudentRecord()
	} else if choice == 2 {
		listStudents(false)
	} else if choice == 3 {
		listStudents(true)
	} else {
		calculateAverageGPA()
	}

	fmt.Println()

	return choice == 5 || err != nil
}

func createStudentRecord() error {
	students, err := loadStudents()

	if err != nil {
		fmt.Println("Error while reading students.")
		return err
	}

	fmt.Println("Enter students index number:")
	var index string
	fmt.Scanf("%s", &index)

	if index == "" {
		fmt.Println("Index can't be empty.")
		return errors.New("Index can't be empty.")
	}

	for _, student := range students {
		if student.index == index {
			fmt.Printf("Student with index %s already exists.\n", index)
			return errors.New("Student with index already exists.")
		}
	}

	fmt.Println("Enter students first name:")
	var firstName string
	fmt.Scanf("%s", &firstName)

	if firstName == "" {
		fmt.Println("First name can't be empty.")
		return errors.New("First name can't be empty.")
	}

	fmt.Println("Enter students last name:")
	var lastName string
	fmt.Scanf("%s", &lastName)

	if firstName == "" {
		fmt.Println("Last name can't be empty.")
		return errors.New("Last name can't be empty.")
	}

	fmt.Println("Enter students GPA:")
	var gpa float64
	fmt.Scanf("%f", &gpa)

	if gpa < 5 || gpa > 10 {
		fmt.Println("GPA must be in range 5 to 10.")
		return errors.New("GPA must be in range 5 to 10.")
	}

	student := Student{index, firstName, lastName, gpa}

	err = writeStudentRecord(student)

	if err != nil {
		fmt.Println("Error while creating a student record.")
		return err
	}

	return nil
}

func listStudents(ascending bool) error {
	students, err := loadStudents()

	if err != nil {
		fmt.Println("Error while reading students.")
		return err
	}

	if ascending {
		sort.Slice(students, func(i, j int) bool {
			return students[i].gpa < students[j].gpa
		})
	} else {
		sort.Slice(students, func(i, j int) bool {
			return students[i].gpa > students[j].gpa
		})
	}

	fmt.Println("STUDENTS:")
	for _, student := range students {
		fmt.Printf("\t- INDEX: %s \t\t\tFIRST NAME: %s \t\t\tLAST NAME: %s \t\t\t GPA: %f\n", student.index, student.firstName, student.lastName, student.gpa)
	}

	return nil
}

func calculateAverageGPA() error {
	students, err := loadStudents()

	if err != nil {
		fmt.Println("Error while reading students.")
		return err
	}

	var gpa float64 = 0

	for _, student := range students {
		gpa += student.gpa
	}

	gpa /= float64(len(students))

	fmt.Printf("AVERAGE GPA: %f\n", gpa)

	return nil
}

func loadStudents() ([]Student, error) {
	file, err := os.OpenFile("./Students/students.txt", os.O_RDONLY, os.ModeExclusive)
	defer file.Close()

	if err != nil {
		fmt.Println("Error while opening file.")
		return nil, err
	}

	students := make([]Student, 0, 0)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		studentData := strings.Split(scanner.Text(), "|")
		gpa, err := strconv.ParseFloat(studentData[3], 64)
		if err != nil {
			fmt.Println("Error while reading student data.")
			return nil, err
		}
		students = append(students, Student{studentData[0], studentData[1], studentData[2], gpa})
	}

	return students, nil
}

func writeStudentRecord(student Student) error {
	file, err := os.OpenFile("./Students/students.txt", os.O_RDWR|os.O_APPEND, os.ModeAppend)
	defer file.Close()

	if err != nil {
		fmt.Println("Error while opening file.")
		return err
	}

	writer := bufio.NewWriter(file)

	studentData := student.index + "|" + student.firstName + "|" + student.lastName + "|" + strconv.FormatFloat(student.gpa, 'f', 2, 64) + "\n"

	_, err = writer.WriteString(studentData)

	if err != nil {
		fmt.Println("Error while writting into file.")
		return err
	}
	err = writer.Flush()

	return nil
}
