package di

import (
	"fmt"
	"io"
	"os"
)

// DI
// Our function doesn't need to care where or how the printing happens,
// so we should accept an interface rather than a concrete type.

// То бишь правило "Программируй на уровне интерфейсов"
// передавать зависимости (сервисы, которые нужны твоему коду для работы) извне,
// а не создавать их внутри функции/объекта
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
