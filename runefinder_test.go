package main

import(
  "os"
  "testing"
  "strings"
)

const linhaLetraA = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`

func TestAnalisarLinha(t *testing.T) {
  runa, nome := AnalisarLinha(linhaLetraA)
  if runa != 'A' {
    t.Errorf("Esperava 'A', veio %q", runa)
  }
  const nomeA = "LATIN CAPITAL LETTER A"
  if nome != nomeA {
    t.Errorf("Esperava %q, veio %q", nomeA, nome)
  }
}

func ExampleListar() {
  texto := strings.NewReader(linhas3Da43)
  Listar(texto, "MARK")
  // Output: U+003F	?	QUESTION MARK
}

func ExampleListar_doisResultados() {
  texto := strings.NewReader(linhas3Da43)
  Listar(texto, "SIGN")
  // Output:
  // U+003D	=	EQUALS SIGN
  // U+003E	>	GREATER-THAN SIGN
}

func Example() {
  oldArgs := os.Args
  defer func() { os.Args = oldArgs }()
  os.Args = []string{"", "cruzeiro"}
  main()
  // Output:
  // U+20A2	₢	CRUZEIRO SIGN
}
