package main

// define the interface
type Storage interface {
	Save(data string) error
	Get(id int) (string, error)
}

// create concrete implementation (MySQL)
type MySQL struct{}

func (m MySQL) Save(data string) error {
	println("Saving to MySQL:", data)
	return nil
}

func (m MySQL) Get(id int) (string, error) {
	return "Data from MySQL", nil
}

// Another implementation MongoDB
type MongoDB struct{}

func (m MongoDB) Save(data string) error {
	println("Saving to MongoDB:", data)
	return nil
}

func (m MongoDB) Get(id int) (string, error) {
	return "Data from MongoDB", nil
}

// Use the interface
func processData(s Storage) {
	s.Save("Hello World")
	data, _ := s.Get(1)
	println(data)
}

func main() {
	d1 := MySQL{}
	processData(d1)

	d2 := MongoDB{}
	processData(d2)
}
