package utilityimplcs
import ("fmt"
"math"
"strconv")
type node struct {
	next *node
	key string
	}
	type queue struct {
	front *node
	back *node
	}
	func Day(month int32, day int32, year int32) int32 {

		mon := float64(month)
		yea := float64(year)
		da := float64(day)
		y := yea - math.Floor((14-mon)/12)
		x := y + math.Floor(y/4) - math.Floor(y/100) + math.Floor(y/400)
		m := mon + 12*math.Floor((14-mon)/12) - 2
		d := (int32)(da+x+math.Floor((31*m)/12)) % 7
		return d
		}
		//check leap year or not
		func LeapYear(year int32) bool {
		if (year%4 == 0) && (year%100 != 0) {
		return true
		}
		if year%400 == 0 {
		return true
		}
		return false
		}
	// using 2 two stacks print calender
	func GetCalenderStack() {
	var month, year int32
	fmt.Println("enter the month ")
	fmt.Scanf("%d", &month)
	fmt.Println("enter the year ")
	fmt.Scanf("%d", &year)
	months := []string{"", "January", "February", "March",
	"April", "May", "June", "July", "August", "September",
	"October", "November", "December"}
	days := []int32{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	
	if month == 2 && LeapYear(year) {
	days[month] = 29
	}
	fmt.Println(months[month], year)
	// insert into 1st stack
	st1 := new(stack)
	st1.Push("Su ")
	st1.Push("Mo ")
	st1.Push("Tu ")
	st1.Push("WE ")
	st1.Push("Th ")
	st1.Push("Fr ")
	st1.Push("Sa")
	d := Day(month, 1, year)
	var i int32
	
	for i = 0; i < d; i++ {
	st1.Push("   ")
	}
	for i = 1; i <= days[month]; i++ {
	typ := (int)(i)
	s := strconv.Itoa(typ)
	st1.Push(s)
	}
	// insert into 2nd stack
	st2 := new(stack)
	for i := 0; i < 7; i++ {
	st2.Push(st1.Pop())
	}
	fmt.Println()
	for i = 0; i < d; i++ {
	element := st1.Pop()
	st2.Push(element)
	}
	for i = 1; i <= days[month]; i++ {
	element := st1.Pop()
	st2.Push(element)
	
	}
	// using second stack display calender
	for i := 0; i < 7; i++ {
	fmt.Print(st2.Pop())
	}
	fmt.Println()
	for i = 0; i < d; i++ {
	element := st2.Pop()
	fmt.Print(element)
	}
	for i = 1; i <= days[month]; i++ {
	fmt.Printf("%2s ", st2.Pop())
	if ((i+d)%7 == 0) || (i == days[month]) {
	fmt.Println()
	}
	}
	
	}
	
	type nodestack struct {
	next *nodestack
	element string
	}
	type stack struct {
	front *nodestack
	}
	
	func (l *stack) Push(data string) bool {
	old := &nodestack{element: data}
	if l.front == nil {
	l.front = old
	} else {
	old.next = l.front
	l.front = old
	}
	return true
	}
	func (l *stack) Pop() string {
	if l.front == nil {
	return ""
	}
	value := l.front.element
	l.front = l.front.next
	return value
	}