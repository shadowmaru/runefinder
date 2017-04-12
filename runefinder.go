package main

import (
  "bufio"
  "fmt"
  "io"
  "log"
  "os"
  "strconv"
  "strings"
)

const linhas3Da43 = `
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
0043;LATIN CAPITAL LETTER C;Lu;0;L;;;;;N;;;;0063;
`

func main() {
  ucd, err := os.Open("UnicodeData.txt")
  if err != nil {
    log.Fatal(err.Error())
  }
  defer func() { ucd.Close() }()
  consulta := strings.Join(os.Args[1:], " ")
  Listar(ucd, strings.ToUpper(consulta))
}

// AnalisarLinha recebe uma linha de um Unicode e retorna o código e o nome
// deste caracter
func AnalisarLinha(ucdLine string) (rune, string) {
  campos := strings.Split(ucdLine, ";")
  código, _ := strconv.ParseInt(campos[0], 16, 32)
  return rune(código), campos[1]
}

// Listar exibe na saída padrão o código, a runa e o nome dos caracteres Unicode
// cujo nome contém o texto da consulta
func Listar(texto io.Reader, consulta string) {
  varredor := bufio.NewScanner(texto)
  for varredor.Scan() {
    linha := varredor.Text()
    if strings.TrimSpace(linha) == "" {
      continue
    }
    runa, nome := AnalisarLinha(linha)
    if strings.Contains(nome, consulta) {
      fmt.Printf("U+%04X\t%[1]c\t%s\n", runa, nome)
    }
  }
}
