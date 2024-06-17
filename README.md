## core

Esta pasta é onde você coloca o código interno do aplicativo. Ele contém a lógica de negócios, como controladores, casos de uso, repositórios e entidades. A ideia por trás desta pasta é que o código aqui não seja acessível fora do pacote (módulo), o que ajuda a manter a encapsulação e a modularidade.

## api

Esta pasta pode conter tudo relacionado à interface da API, como definições de rotas, manipuladores de requisições HTTP e documentação da API. Ela fornece uma separação clara entre a lógica do aplicativo e a interface de comunicação com o mundo exterior.

## pkg

Esta pasta contém pacotes ou bibliotecas reutilizáveis que podem ser usados em diferentes partes do aplicativo ou até mesmo em outros projetos. Por exemplo, pode conter utilitários genéricos, funções de configuração, manipuladores de erros, entre outros.

---

#### app

Esta pasta é onde você coloca a lógica de aplicativo específica, como controladores, casos de uso, repositórios e entidades. Ela contém a maior parte da lógica de negócios do aplicativo.

#### controller

Aqui você coloca os controladores HTTP que recebem as requisições do cliente, chamam os casos de uso apropriados e retornam as respostas HTTP.

#### usecase

Esta pasta contém os casos de uso do aplicativo, que contêm a lógica de negócios específica do aplicativo. Os casos de uso orquestram as operações de negócios e interagem com os repositórios para acessar os dados.

#### repository

Os repositórios são responsáveis por interagir com a camada de armazenamento de dados (banco de dados, cache, etc.). Eles encapsulam a lógica de acesso aos dados e fornecem métodos para recuperar, criar, atualizar e excluir entidades.

#### entity

Esta pasta contém as definições de entidade do aplicativo, que representam os objetos de domínio principais com os quais o aplicativo trabalha. As entidades são frequentemente mapeadas diretamente para tabelas de banco de dados.
