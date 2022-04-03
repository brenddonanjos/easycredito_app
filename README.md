
<p align="center"><a href="https://easycredito.com.br" target="_blank"><img src="https://cdn2.easycredito.com.br/assets/main/logo-easycredito-6047341476fccf58a054d87a48cf1b8ab0f88b36b9af01dc0f54583ec18c93a7.png" width="180"></a></p>

<h2 align="center">Back-end Challenge 🏅 2021 - Space Flight News</h2>
<p align="center">Escrito em GoLang com  Echo Framework </p>
<p align="center">Módulo EasyCrédito para integração com a API Space Flight News</p>

<p align="center">
<a href="#install">Instalação</a> •
<a href="#docs">Documentação</a> • 

## Instalação
<div id="install">
<p>Para instalação é necessário que o docker e docker-compose estejam devidamente instalados em sua máquina</p>
<p>O sistema está configurado para iniciar no localhost (127.0.0.1) porta 8000, e o mysql iniciará na porta 3306, certifique-se que estas portas estejam livres antes de iniciar a instalação.</p>
<p>1. Entre na pasta do projeto: </p>

```
cd easycredito_app
```
<p>2. Iremos subir os containeres utilizando o docker-compose (Certifique-se de as portas 8000 e 3306 estejam livres): </p>   

```
docker-compose up -d
```

</div>

## Documentação
<div id="docs">
<p>Após iniciar os containeres corretamente e o sistema estiver funcionando, a página inicial do sistema (127.0.0.1:8000) disponibiliza uma descrição detalhada sobre as requisições e respostas</p>
<p>No diretório  /docs estará disponibilizado um diagrama EER do banco de dados</p>
<p>Também no diretório /docs foi disponibilizado uma collection e um enviroment para acesso via postman</p>
<p>Para visualização da collection é necessário que o postman esteja devidamente instalado na em máquina</p>
<p>1. Com o programa aberto importe a collection "go_easycredito_app.postman_collection" clicando no botão "import" no canto superior esquerdo.</p>
<p>2. Agora importe o envirnoment "go_easycredito_app.postman_environment.json" e selecione o environment importado no canto superior esquerdo</p>
</div>


As seguintes estão sendo foram usadas na construção do projeto:

- [GoLang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Mysql](https://www.mysql.com/)
