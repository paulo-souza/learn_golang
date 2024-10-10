package lib

import "fmt"

type Fulano struct {
	nome  string
	idade int
}

// SetNome recebe um ponteiro para Fulano para que possa modificá-lo.
func (n *Fulano) SetNome(novoNome string) {
	n.nome = novoNome
}

// Nome recebe uma cópia de Fulano, pois não precisa modificá-lo.
func (n Fulano) GetNome() string {
	return n.nome
}

// SetIdade recebe um ponteiro para Fulano para que possa modificá-lo.
func (i *Fulano) SetIdade(novaIdade int) {
	i.idade = novaIdade
}

// Idade recebe uma cópia de Fulano, pois não precisa modificá-lo.
func (i Fulano) GetIdade() int {
	return i.idade
}

func ExecuteExemploGetESet() {
	// Observe o Fulano{}. O new(Fulano) era apenas um açúcar sintático para &Foo{}
	// e não precisamos de um ponteiro para o Foo, então eu o substituí.
	// no entanto, não é relevante para o problema.

	fulano := Fulano{}

	fulano.SetNome("Fulano de Tal")
	fulano.SetIdade(18)

	fmt.Printf("Nome: %s, Idade: %d\n", fulano.GetNome(), fulano.GetIdade())
	fmt.Println(fulano)
}
