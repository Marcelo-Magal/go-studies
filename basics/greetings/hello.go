// O arquivo `hello.go` contém a função `SayHello`, que exibe a mensagem `"Hello, World!"` quando chamada. Ele faz parte do pacote `greetings`.

package greetings

import "fmt"

func SayHello() {
	fmt.Println("Hello, World!")
}

// - **`package greetings`**: Este arquivo faz parte do pacote `greetings`. Todos os arquivos que compartilham o mesmo nome de pacote podem ter suas funções acessadas de forma consistente.

// - **`import "fmt"`**: Estamos importando o pacote `fmt` da biblioteca padrão de Go para exibir a mensagem no console. `fmt.Println()` é a função usada para imprimir texto no terminal.

// - **`func SayHello()`**: Define a função `SayHello`, que imprime `"Hello, World!"` quando chamada.
