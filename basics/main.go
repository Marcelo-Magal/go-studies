// O arquivo principal que inicia o programa e chama as funções definidas nos outros arquivos.

package main

import (
	"fmt"

	"github.com/Marcelo-Magal/go-studies/basics/greetings"
	"rsc.io/quote/v4"
)

func main() {
	greetings.SayHello()
	greetings.SayHi()
	fmt.Println(quote.Go())
}

// #### Explicação:

// - **package main**:
//   - Todo programa Go precisa de um pacote principal, que é o `package main`. O pacote `main` define o ponto de entrada do programa, onde a execução começa.

// - **import "go-studies/basics/greetings"**:
//   - A palavra-chave `import` é usada para incluir outros pacotes no seu programa. Aqui, estamos importando o pacote `greetings`, que está no diretório `basics/greetings`, para que possamos usar suas funções dentro do `main.go`.

// - **func main() { ... }**:
//   - Esta é a função principal (`main()`). É a função especial onde o programa começa a rodar. Todo código que for executado inicialmente será colocado dentro desta função.
