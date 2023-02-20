package main

import (
	"fmt"
)

func main() {
	//MAPS É PARECIDO COM O STRUCT, PORÉM MAIS RÍGIDO
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"ultimo":   "Pedro",
		},

		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Touro",
	}

	fmt.Println(usuario2)
	// -> acessar valores no maps fmt.Println(usuario2["nome"])
}
