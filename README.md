# 🎂 FamBirthday Bot 🤖

![](https://github.com/Dedo-Finger2/fam-birthday-bot/blob/master/public/images/cover.png?raw=true)

<p align="center">
	<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" />
	<img src="https://img.shields.io/badge/Telegram-2CA5E0?style=for-the-badge&logo=telegram&logoColor=white" />
</p>

<p align="center">🚀 Receba <strong>notificações</strong> das <strong>datas de aniversário</strong> de pessoas queridas no <strong>Telegram</strong>! 🚀</p>

---

## 📔 Descrição

O Fam Birthday Bot é um cron job que roda um sistema de mensageria feito em Golang, integrando com a API do Telegram para criar um bot. 

Este bot é usado para enviar mensagens para uma lista de usuários com permissão de receber suas mensagens, e isso é definido em um arquivo .env. Com isso vem o cron job, que está configurado para rodar sempre as 5 horas da manhã de diariamente, nesse processo o sistema faz uma validação indo ao banco de dados. Que atualmente é um arquivo de configuração em YAML, e itera sobre todas as datas presentes lá. Cada data possui um array de pessoas, e cada pessoa possui nome, idade e complemento.

| O complemento serve para identificar pessoas com nomes iguais. Diferenciando eles.

Se a data atual bater com uma data no arquivo YAML, então todas as pessoas dentro do array de pessoas daquela data fazem aniversário no dia. Então é pego o complemento e nome das pessoas e enviado com uma mensagem pre-feita para os usuários com permissão de receber mensagens do bot. Depois disso o sistema entra em hiato por 24 horas, fazendo a validação apenas 24 horas depois.

## 🎯 Objetivo

O objetivo deste projeto é auxiliar o usuário que frequentemente esquece as datas de aniversário de pessoas conhecidas através do envio de mensagens no Telegram as 5 horas da manhã sempre que houver um aniversariante no dia. Dando assim tempo para o usuário dar feliz aniversário para esta pessoa.

## Como usar

Clone the repo
```bash
git clone https://github.com/Dedo-Finger2/fam-birthday-bot.git
```

Creates the birthdates config file
```bash
cp internal/config/birth_dates.example.yml internal/config/birth_dates.yml
```

Creates the .env file
`OBS: The TESTING_CHAT_ID can be your chat id`
```bash
cp .env.example .env
```

Execute the tests
```bash
make test
```

Runs the application
```bash
make dev
```

## ⚠️ Requisitos

### Funcionais

- [x] O sistema deve poder enviar mensagens para uma lista de usuários
- [x] O sistema deve usar um cron job para executar a validação de data em um determinado horário dia
- [x] O sistema deve poder integrar com a API do Telegram
- [x] O sistema deve poder enviar as mensagens através de um bot no Telegram
- [x] O sistema deve poder lidar com casos onde hajam mais de um aniversariante no dia, formatando a mensagem template para encaixar mais de um nome
- [x] O sistema só deve mandar mensagem para uma lista seleta de usuários com permissão para receber as mensagens
- [x] Deve ser usado um arquivo YAML para configuração
- [x] Deve ser usado um arquivo JSON para testes

### Não funcionais

- [ ] Deve haver um QR code para acessar o Bot no Telegram;
- [ ] O sistema deve constar com um subsistema de logs feitos a nível de linha de comando;
- [ ] Deve existir um Google Forms que seja capaz de coletar dados para serem usados no sistema;
- [x] Log de erros durante o envio de mensagens
- [x] Segurança dos dados do bot e do nome dos aniversariantes
- [x] Performance para lidar com vários envios sem sobrecarregar o servidor


### Regras de negócio

- [x] As mensagens só devem ser enviadas para IDs cadastrados no sistema;
- [x] A validação de data deve ser feita apenas uma vez por dia
- [x] As mensagens só devem ser enviadas caso haja um match com uma data de aniversário e a data atual da validação
- [x] Caso não hajam aniversariantes no dia o sistema deve aguardar 24 horas para validar novamente as datas

## ⚒️ Infraestrutura

### Fluxograma

![](https://github.com/Dedo-Finger2/fam-birthday-bot/blob/master/public/images/diagram.png?raw=true)

### 🖿 Estrutura de pastas

```markdown
- builds/
  - fam-birthdate-amd
  - fam-birthdate-arm
  - fam-birthdate.exe
- cmd/
  - main.go
- internal/
    - config/
      - birth_dates.yml
      - birth_dates.json
      - bot.go
    - types/
    - utils/
- public
  - images/
Makefile
README.md
LICENSE
go.mod
go.sum
```

### 🖥️ Tecnologias

| Tech             | Utilidade                                                    | Versão |
| :--------------- | :----------------------------------------------------------- | :----- |
| Golang           | Linguagem de programação usada                               | 1.22.6 |
| Neovim           | Editor de código via terminal                                | 10     |
| tgbotapi         | API do Telegram                                              | 5.5.1  |
| robfig/cron      | Biblioteca para criação de cron jobs em Golang               | 3.0.1  |
| viper            | Biblioteca para trabalho com variáveis de ambiente em Golang | 1.19.0 |
| gopkg.in/yaml.v3 | Biblioteca para trabalho com arquivos YAML em Golang         | 3.0.1  |

## 🌐 Implementações futuras

- Deploy;
- Cadastro de novas datas de aniversário através do bot;
- Envio de mensagem de parabéns para o aniversariante;
- Envio de mensagens seletas para usuários específicos;
	- Eu só quero ser notificado das datas de fulano, sicrano e beltrano.
- Tratamento de erros na hora de enviar mensagens, adiando o envio até que seja enviada;
- Separar em dois micro serviços dependentes;
- Separação da aplicação em 2 micro serviços, um de mensageiria e outro para tratar das datas de aniversário;

## ✏️ O que eu aprendi com este projeto

- Envio de mensagens mediante um bot no Telegram com a linguagem Go;
- Criação de um bot de Telegram usando Go;
- Agendamento de tarefas feitas em Go;
- Deploy de aplicações Go;
- Tratamento de erros em Go com SLog;
- Calcular quanto tempo falta para determinada data em Go;
- Criação de uma data customizada em Go;
- Obtenção do diretório raiz do projeto em Go;
- Formatação de horas e minutos oriundos da diferença de tempo entre duas datas em Go;
- Testes unitários em Go;
- Trabalho com variáveis de ambiente usando a biblioteca Viper em Go;

## Meus contatos

- Instagram: https://www.instagram.com/antonioalmeida2003/
- LinkedIn: https://www.linkedin.com/in/antonio-mauricio-4645832b3/
- Email: antonioimportant@gmail.com
