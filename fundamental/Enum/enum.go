package enum

import "fmt"

// Statusと言う独自型（Stringの独自型）を定義
type Status string

// 上記の独自型で定義された定数を定義して、これをEnum型（列挙型）として利用する
// 以下の3つに限定したい場合に、この書き方を使えば、以降のコーディング時にタイポを防ぐことができる
const (
	StudentStatusHighSchool Status = Status("High School") // Status型にアサーション
	StudentStatusCollege    Status = Status("College")
	StudentStatusGraduated  Status = Status("Graduated")
)

type Student struct {
	ID     int
	Name   string
	Age    int
	Status Status
}

// 下記で言うレシーバーとは、実質オブジェクトのこと

// 独自型に紐づけてメソッドを定義することも出来る
// ゲッターなら値型でレシーバーを定義した方がいい？
func (s Status) GetStatus() string {
	return string(s)
}

// セッターならポインタ型でレシーバーを定義した方がいい？
func (s *Status) SetStatus(status Status) {
	*s = status
}

// ゲッターなら値型でレシーバーを定義した方がいい？
func (s Student) GetAge() int {
	return s.Age

}

// セッターならポインタ型でレシーバーを定義した方がいい？
func (s *Student) SetAge(age int) {
	s.Age = age
}

// ポインタ型のレシーバーなら、こんな感じでメソッドの中でswitch文を使って、状態を変更することもできる
func (s *Student) NextStep() {
	switch s.Status {
	case StudentStatusHighSchool:
		s.Status = StudentStatusCollege
	case StudentStatusCollege:
		s.Status = StudentStatusGraduated
	case StudentStatusGraduated:
		fmt.Println("You have already graduated")
	default:
		fmt.Println("Unknown status")
	}
}

func EnumPractice() {
	student := Student{
		ID:     1,
		Name:   "John",
		Age:    20,
		Status: StudentStatusHighSchool,
	}

	fmt.Println(student.GetAge()) // 20

	switch student.Status {
	case StudentStatusHighSchool:
		println(student.Name, "is", student.Status) // John is High School
	case StudentStatusCollege:
		println(student.Name, "is", student.Status)
	case StudentStatusGraduated:
		println(student.Name, "is", student.Status)
	default:
		println(student.Name, "is", "Unknown status")
	}

	// ポインタ型のレシーバーであっても、その呼び出し時は値型でも呼び出すことができる

	// セッターを使って年齢を変更
	student.SetAge(21)

	fmt.Println(student.GetAge()) // 21

	// 以下の2つは同じ意味
	fmt.Println(StudentStatusHighSchool.GetStatus()) // High School

	fmt.Println(student.Status.GetStatus()) // High School

	// ここでCollegeに変更
	student.NextStep()

	fmt.Println(student.Status.GetStatus()) // College

	// ここでGraduatedに変更
	student.Status.SetStatus(StudentStatusGraduated)

	// 以下の2つは同じ意味
	fmt.Println(student.Status.GetStatus()) // Graduated

	fmt.Println(student.Status) // Graduated

	student.NextStep() // You have already graduated

}
