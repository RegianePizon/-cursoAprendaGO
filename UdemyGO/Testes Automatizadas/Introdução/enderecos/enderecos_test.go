//teste de unidade

package enderecos

import "testing"

type cenarioDeTest struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTest{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado!! Esperava %s e recebeu %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
	/*
		//TestNomedafunção
		enderecoParaTeste := "Rua Paulista"
		tipoDeEnderecoEsperado := "Avenida"
		tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

		if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado!! Esperava %s e recebeu %s",
				tipoDeEnderecoEsperado,
				tipoDeEnderecoRecebido,
			)
			//	t.Error("O tipo recebido é diferente do esperado!!")
		}
	*/
}
func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste Quebrou")
	}
}

//descreve quem esta sendo testado (  go test -v    )-
//Porcentagem do teste ( go test --cover)
//go tool cover --func=resultado.txt
