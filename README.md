# Oficina de código: Desafio Go

### Objetivo:
A seguir, é colocado um desafio integrador que nos permitirá avaliar os tópicos vistos até agora.

## Enunciado
Uma CIA aérea pequena possui um sistema de reservas de bilhetes para diferentes países. Ele retorna um arquivo com as informações sobre os bilhetes comprados nas últimas 24 horas. A CIA precisa de um programa para extrair informações sobre as vendas do dia e, assim, analisar as tendências de compra.
O arquivo em questão é do tipo valores separados com vírgula (CSV), onde os campos são formados por: id, nome, email, país de destino, hora do voo e preço.
Boa sorte!

## Desafio
Realizar um programa que sirva como ferramenta para calcular diferentes dados estatísticos. Um repositório foi fornecido com o arquivo csv e os requisitos abaixo:

## Requisitos

- Uma função que calcule a quantidade de pessoas que viajam para um país determinado
  ````go
  func GetTotalTickets(destination string) (int,error){}
  
- Uma ou várias funções que calculam quantas pessoas viajam de madrugada (0 → 6),
  manhã (7 → 12), tarde (13 → 19), e noite (20 → 23).
  ````go
  func GetCountByPeriod(time string) (int,error){}
- Calcular a quantidade média de viagens.
  ````go
  func AverageDestination() (float64,error){}