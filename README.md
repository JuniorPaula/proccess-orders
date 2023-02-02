# Microservice Golang and RabbitMQ

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![MySQL](https://img.shields.io/badge/mysql-%2300f.svg?style=for-the-badge&logo=mysql&logoColor=white)
![Git](https://img.shields.io/badge/git-%23F05033.svg?style=for-the-badge&logo=git&logoColor=white)
![RabbitMQ](https://img.shields.io/badge/Rabbitmq-FF6600?style=for-the-badge&logo=rabbitmq&logoColor=white)
![Grafana](https://img.shields.io/badge/grafana-%23F46800.svg?style=for-the-badge&logo=grafana&logoColor=white)
![Prometheus](https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=Prometheus&logoColor=white)


A aplicação consiste em um **Microserviço** de processamentos de pedidos, fazendo um calculo simbólico do valor total de cada operação e salvando o resultado no banco de dados **mysql**. O **Producer ** irá publicar os pedidos no **RabbitMQ** via **amqp**, enquanto o **Consumer** irá consumir essas messages das filas do RabbitMQ e salvará o resultado no banco. Além desses serviços a aplicação ainda responde um servidor **http** devolvendo o numero de processos já realizados através do endpoit `/total`.  

Esse é um estudo de casos utilizando **Golang** e testando toda a sua capacidade e performace em lidar com situações que exigem uma alta demanda de processamento.

#### As tecnologias utilizadas foram:
- golang
- docker
- mysql
- rabbitmq
- grafana
- prometheus