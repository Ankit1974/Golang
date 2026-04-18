package slicepractice

func main() {

	var studentName = []string{"Ankit", "raj", "Gulshan"}

	studentName = append(studentName, "Kumar", "Ayush")

	// if we have to delete from slice at particular index

	index := 4

	studentName = append(studentName[:index], studentName[index+1:]...)
}
