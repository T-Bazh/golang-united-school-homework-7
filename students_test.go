package coverage

import (
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestLenPeople(t *testing.T) {
	personA := Person{firstName: "Mark", lastName: "Nelson",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Mandy", lastName: "Smith",
		birthDay: time.Date(2005, time.July, 10, 10, 10, 0, 0, time.UTC),
	}
	expectedResult := 2
	listPeople := People{personA, personB}

	actualResult := listPeople.Len()
	if actualResult != expectedResult {
		t.Errorf("Expeсted Len: %d, Received Len: %d", expectedResult, actualResult)
	}
}

func TestLessPeopleFirstName(t *testing.T) {
	personA := Person{firstName: "Mars", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Marsy", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	peopleList := People{personA, personB}
	expectedResult := true
	actualResult := peopleList.Less(0, 1)
	if actualResult != expectedResult {
		t.Errorf("People Less Function: Expeсted Result: %t, Received Result: %t", expectedResult, actualResult)
	}
}

func TestLessPeopleLastName(t *testing.T) {
	personA := Person{firstName: "Mark", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Mark", lastName: "Smithfield",
		birthDay: time.Date(2002, time.March, 1, 0, 0, 0, 0, time.UTC),
	}
	peopleList := People{personA, personB}

	expectedResult := true
	actualResult := peopleList.Less(0, 1)
	if actualResult != expectedResult {
		t.Errorf("People Less Function: Expeсted Result: %t, Received Result: %t", expectedResult, actualResult)
	}
}

func TestLessPeopleBirthday(t *testing.T) {
	personA := Person{firstName: "Mark", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 10, 10, 0, 0, 0, time.UTC),
	}
	personB := Person{firstName: "Mark", lastName: "Smith",
		birthDay: time.Date(2002, time.March, 10, 0, 0, 0, 0, time.UTC),
	}
	peopleList := People{personA, personB}

	expectedResult := true
	actualResult := peopleList.Less(0, 1)
	if actualResult != expectedResult {
		t.Errorf("People Less Function: Expeсted Result: %t, Received Result: %t", expectedResult, actualResult)
	}
}

func TestLessSwap(t *testing.T) {
	birthdayA := time.Date(2002, time.May, 10, 0, 0, 0, 0, time.UTC)
	birthdayB := time.Date(2002, time.May, 5, 0, 0, 0, 0, time.UTC)
	personA := Person{firstName: "Mark", lastName: "Smith", birthDay: birthdayA}
	personB := Person{firstName: "Mandy", lastName: "Black", birthDay: birthdayB}
	peopleList := People{personA, personB}

	peopleList.Swap(0, 1)
	if peopleList[0].firstName != "Mandy" || peopleList[0].lastName != "Black" || peopleList[0].birthDay != birthdayB ||
		peopleList[1].firstName != "Mark" || peopleList[1].lastName != "Smith" || peopleList[1].birthDay != birthdayA {
		t.Errorf("People Swap Function: Expeсted Result is different from Received Result")
	}
}

func TestNewMatrix(t *testing.T) {
	matrixStr := "1 2\n3 4"
	m, err := New(matrixStr)
	if err != nil {
		t.Errorf("Failed to create a new Matrix")
		if m.cols != 2 {
			t.Errorf("Number of columns is not correct")
		}
		if m.rows != 2 {
			t.Errorf("Number of rows is not correct")
		}
	}
}

func TestNewMatrixOnEmptyString(t *testing.T) {
	_, err := New(" ")
	if err == nil {
		t.Errorf("String must not be empty")
	}
}

func TestNewMatrixInvalidRows(t *testing.T) {
	_, err := New("10 20 30\n40 50 60\n70 80")
	if err == nil {
		t.Errorf("Rows need to be the same length")
	}
}

func TestRowsMatrix(t *testing.T) {
	m, err := New("10 20\n30 40")
	if err != nil {
		t.Errorf("Failed to create a new Matrix")
	}
	expected := [][]int{{10, 20}, {30, 40}}
	result := m.Rows()
	for i, value := range result {
		for j, valCol := range value {
			if valCol != expected[i][j] {
				t.Errorf("Expected Result %d, Received Result %d", expected[i][j], valCol)
			}
		}
	}
}

func TestColsMatrix(t *testing.T) {
	m, err := New("10 20\n30 40")
	if err != nil {
		t.Errorf("Failed to create a new Matrix")
	}
	expected := [][]int{{10, 30}, {20, 40}}
	result := m.Cols()
	for i, value := range result {
		for j, valRow := range value {
			if valRow != expected[i][j] {
				t.Errorf("Expected Result %d, Received Result %d", expected[i][j], valRow)
			}
		}
	}
}

func TestSetMatrix(t *testing.T) {
	m, err := New("10 20 30\n40 50 60\n70 80 90")
	if err != nil {
		t.Errorf("Failed to create a new Matrix")
	}
	ok := m.Set(1, 1, 55)
	if !ok || m.data[4] != 55 {
		t.Errorf("Failed to set a new value")
	}
}

func TestSetMatrixInvalid(t *testing.T) {
	m, err := New("10 20 30\n40 50 60\n70 80 90")
	if err != nil {
		t.Errorf("Failed to create a new Matrix")
	}
	ok := m.Set(-1, 1, 55)
	if ok {
		t.Errorf("Failed to set an invalid value")
	}
}
