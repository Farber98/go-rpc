package structs

import "fmt"

type Student struct {
	Id        int
	FirstName string
	LastName  string
}

func (s Student) FullName() string {
	return s.FirstName + " " + s.LastName
}

type College struct {
	studentsDb map[int]Student
}

func (c *College) Add(payload Student, reply *Student) error {
	if _, ok := c.studentsDb[payload.Id]; ok {
		return fmt.Errorf("Student with id %d already exists", payload.Id)
	}
	c.studentsDb[payload.Id] = payload
	*reply = payload
	return nil
}

func (c *College) Get(payload int, reply *Student) error {
	result, ok := c.studentsDb[payload]
	if !ok {
		return fmt.Errorf("student with id %d does not exist", payload)
	}
	*reply = result
	return nil
}

func NewCollege() *College {
	return &College{studentsDb: make(map[int]Student)}
}
