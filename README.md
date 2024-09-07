# 📘 Study Cycle Manager

![](https://github.com/Dedo-Finger2/study-cycle-manager/blob/study-cycle-manager-experiment/public/images/cover.png?raw=true)

<p align="center">
	<img src="https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white" />
</p>

<p align="center">🚀 Gerencie <strong>ciclos de estudos</strong> através do seu terminal! 🚀</p>

---

## 📔 Descrição

O projeto consiste em uma CLI projetada para auxiliar o usuário na gestão de seu ciclo de estudos. Com essa ferramenta, é possível adicionar e remover matérias, selecionar um ciclo de estudos específico para visualização, deletar ciclos e, principalmente, marcar checkboxes das matérias para acompanhar o progresso de cada uma.

## 🎯 Objetivo

Auxiliar o usuário a gerenciar um ciclo de estudos com uma aplicação via terminal.

## 📺 Demo

https://github.com/user-attachments/assets/f8f049e6-81ad-46c9-945f-ca1f23fceee7

## ⚙️ Comandos

- Criar um novo ciclo de estudos: `scm create --title="Nome do ciclo de estudos"`
- Listar todos os ciclos de estudos: `scm list`
- Listar ciclo de estudos selecionado: `COMMING SOON`
- Deletar ciclo de estudos: `COMMING SOON`
- Adicionar matéria ao ciclo de estudos: `scm add --name="Nome da matéria" --max-study-hours 10`
- Estudar uma matéria (add 1 hora): `scm study --id 1`
- Remover matéria do ciclo de estudos: `COMMING SOON`
- Visualizar progresso do ciclo de estudos: `scm view`

## ✨ Detalhes

Um ciclo de estudos é uma metodologia de organização do tempo de estudo que visa otimizar a aprendizagem e aumentar a produtividade. Em vez de seguir um cronograma rígido, o ciclo de estudos permite uma abordagem mais flexível e dinâmica. Onde você impõe um limite do quanto tem que estudar de cada matéria. Esse valor sendo representado pela flag `max-study-hours`. Se uma matéria chegar nesse valor você é obrigado a estudar outra matéria que não seja aquela até seu ciclo acabar. Garantindo assim que você estude todas as matérias necessárias.

> OBS: ESTE PROJETO APENAS GERENCIAR UM CICLO DE ESTUDOS JÁ DEFINIDO! SE VOCÊ NUNCA VEZ UM, FAÇA ANTES DE USAR O SOFTWARE.

## ⚠️ Requisitos

### Funcionais

- [x] Usuários devem poder registrar novos ciclos de estudos
- [x] O sistema deve listar todos os ciclos de estudos criados pelo usuário
- [x] O sistema deve registrar os ciclos de estudos com títulos únicos
- [x] O sistema deve permitir que o usuário selecione um dos ciclos de estudos criados
- [x] O sistema deve permitir a exclusão de ciclos de estudos mediante o seu ID
- [x] Usuários devem poder selcionar ciclos de estudos mediante seus IDs
- [x] Usuários devem poder adicionar novas matérias aos ciclos de estudos
- [x] O sistema deve exibir uma visão geral das matérias de um ciclo de estudos, com quadrados vázios para cada hora de estudo e preenchidos para cada hora estudada
- [x] Usuários devem poder remover matérias dos ciclos de estudos
- [x] O sistema deve registrar as horas estudadas de cada matéria do ciclo de estudos selecionado


### Não funcionais

- [x] Deve ser usado banco de dados SQLite
- [ ] O projeto deve ser testado
- [x] Deve haver um banco de dados dedicado para testes
- [x] Deve haver uma documentação detalhada dos comandods do projeto
- [x] O sistema deve possuir um comando de help interno
- [x] O projeto deve ser dockerizado
- [x] O projeto deve conter um Makefile com comandos de build e clean para melhor experiência de desenvolvimento


### Regras de negócio

- [x] O sistema deve impedir que os usuários atualizem as horas estudadas de matérias que atingiram o limite máximo de horas estudas até o fim do ciclo
- [x] O usuáio deve poder reiniciar o ciclo de estudos quando todas as matérias atingirem seu limite máixmo de horas estudadas, o reset deve zerar as horas estudadas de todas as matérias
- [x] Ao reiniciar um ciclo de estudos o sistema deverá somar mais um na contagem de ciclos concluídos
- [x] O sistema deve atualizar somente os dados do ciclo de estudos selecionado
- [x] O sistema deve impedir os usuários de selecionarem mais de um ciclo de estudos por vez
- [x] O sistema deve impedir os usuários atualizarem as horas estudadas de matérias fora do ciclo de estudos selecionado
- [x] O sistema deve exigir que os usuários selecionem um ciclo de estudos antes de gerenciar as horas estudadas das matérias
- [x] O sistema deve bloquear a inserção de matérias com nomes duplicados em um ciclo de estudos
- [x] O sistema deve formatar os nomes dos ciclos de estudos e das matérias, deixando o título dos ciclos de estudos com inicial maiúscula e separado por espaços enquanto o nome das matérias totalmente em minúsculo e separado por hífens
- [x] O sistema deve ser transparente, informando o usuário de erros que acontecerem durante a execução de tarefas


## ⚒️ Infraestrutura

### Banco de dados

```sql
CREATE TABLE IF NOT EXISTS study_cycles (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  title TEXT NOT NULL UNIQUE,
  completed_times INTEGER NOT NULL,
  selected BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS study_cycle_subjects (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  study_cycle_id INTEGER NOT NULL,
  name TEXT NOT NULL,
  max_study_hours INTEGER NOT NULL,
  user_studied_hours INTEGER NOT NULL,
  completed_times INTEGER NOT NULL,
  added_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP,

  FOREIGN KEY (study_cycle_id) REFERENCES study_cycles(id)
  UNIQUE (study_cycle_id, name)
);
```

### 🖿 Estrutura de pastas

```markdown
- bin/
- cmd/
  - scm/
    - main.go
- internal/
    - commands/
    - databasa/
    - store/
      - sqlite/
        - migrations/
        - repositorie/
	- test_database.db
	- database.db
- public
  - images/
  - videos/
Makefile
Dockerfile
README.md
LICENSE
go.mod
go.sum
```

### 🖥️ Tecnologias

| Tech             | Utilidade                                                    | Versão |
| :--------------- | :----------------------------------------------------------- | :----- |
| Golang           | Linguagem de programação usada                               | 1.23 |
| Neovim           | Editor de código via terminal                                | 10     |
| sqlite3         | Driver do banco de dados SQLite                                              | 1.14.23  |

## 🌐 Implementações futuras

- Criação de um ciclo de estudos do zero;
- Exportação do ciclo de estudos como imagem;
- Exportação do ciclo de estudos como PDF;
- Comando de limpeza do banco de dados;
- Comando de backup do banco de dados;
- Sugestão de comandos quando o usuário digitar um comando errado;

## ✏️ O que eu aprendi com este projeto

- Como trabalhar com SQLite e Golang;
- Como transformar dados vindos de um banco de dados relacional para uma variável ou slice de um tipo específico em Golang;
- Como imprimir dados de forma tabular em Golang;
- Como criar comandos para CLI em Golang;
- Como fazer uma CLI em Golang;
- Como formatar strings em Golang;
- Como ordenar um slice de structs em Golang;
- Como dockerizar uma aplicação Golang;
- Como obter o diretório padrão do projeto em Golang;
- Como organizar comandos e flags individuais de comandos em Golang;
- Como trabalhar com Makefiles em Golang;
- Como rodar uma aplicação Golang em ambiente Windows;

## Meus contatos

- Instagram: https://www.instagram.com/antonioalmeida2003/
- LinkedIn: https://www.linkedin.com/in/antonio-mauricio-4645832b3/
- Email: antonioimportant@gmail.com
