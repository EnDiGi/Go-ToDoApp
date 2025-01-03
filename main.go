package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)


var functions []string = []string{"show", "add", "complete", "delete", "quit"} 

var tasks []string = []string{}
var completedTasks []string = []string{}

func getInstruction() (string, string) {

	reader := bufio.NewScanner(os.Stdin)

	var function string;
	var parameter string;
	var input string;

	fmt.Println("What would you like to do?")
	fmt.Println("Here are the available functions")	
	fmt.Println("1. show")
	fmt.Println("2. add newTaskName")
	fmt.Println("3. complete taskName")
	fmt.Println("4. delete taskName or --all or --completed")
	fmt.Println("5. quit")

	for {
		fmt.Print("> ")
		reader.Scan()
		input = reader.Text()

		parts := strings.Split(input, " ")
        if len(parts) == 0 {
            fmt.Println("Please enter a command.")
            continue
        }

		function = parts[0]
		if len(parts) > 1 {
			parameter = strings.Join(parts[1:], " ")
		}

		for i := 0; i < len(functions); i++ {
			if function == functions[i]{
				return function, parameter
			}
		}
		fmt.Printf("%s is not a valid function\n", function)
	}
}

func executeInstruction(instruction string, parameter string){
	switch instruction{
	case "show":
		show()
	case "add":
		add(parameter)
	case "complete":
		complete(parameter)
	case "delete":
		delete(parameter)
	case "quit":
		quit()
	}
	
	main()
	
}

func show(){
	fmt.Println("All tasks:")
	if len(tasks) == 0 {
		fmt.Println("No task")
	}
	for i := 0; i < len(tasks); i++{
		fmt.Println(tasks[i])
	}

	fmt.Println("Completed tasks:")
	if len(completedTasks) == 0{
		fmt.Println("No completed task")
	}
	for i := 0; i < len(completedTasks); i++{
		fmt.Println(completedTasks[i])
	}
}

func add(newTaskName string){
	tasks = append(tasks, newTaskName)
	fmt.Println("Task added!")
}

func complete(task string){
	var isTask bool;

	for i := 0; i < len(tasks); i++{
		if tasks[i] == task{
			isTask = true
			tasks = append(tasks[:i], tasks[i + 1:]...)
			completedTasks = append(completedTasks, task)
		}
	}
	if isTask{
		fmt.Println("Task marked as completed!")
	} else{
		fmt.Println("Task not in task list")
	}
}

func delete(task string){
	if task == "--all"{
		fmt.Println("Every task has been deleted!")
		tasks = tasks[:0]
		completedTasks = completedTasks[:0]
		return
	} else if task == "--completed"{
		fmt.Println("Every completed task has been deleted!")
		completedTasks = completedTasks[:0]
		return
	}
	var isTask bool;
	for i := 0; i < len(tasks); i++{
		if tasks[i] == task{
			isTask = true
			tasks = append(tasks[:i], tasks[i + 1:]...)
		}
	}

	if isTask{
		fmt.Println("Task deleted!")
	} else {
		fmt.Println("Task not in tasks list")
	}
}

func quit(){
	fmt.Println("Bye!")
	os.Exit(0)
}

func main(){
	fmt.Println("----------------------------------------")
	var instruction, parameter string = getInstruction()
	executeInstruction(instruction, parameter)
}