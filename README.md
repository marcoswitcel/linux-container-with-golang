# Estudando um container simples implementado com a linguagem GO

Esse repositório foi criado para conter o código e a documentação da minha experiência testando o código provido pela usuário [https://github.com/lizrice](https://github.com/lizrice).

Segue os links do código fonte dela [https://github.com/lizrice/containers-from-scratch](https://github.com/lizrice/containers-from-scratch) e da apresentação feita na GOTO 2018 [https://www.youtube.com/watch?v=8fi7uSYlOdc](https://www.youtube.com/watch?v=8fi7uSYlOdc).

## Configurando o ambiente

Esse código foi feito para ser rodado dentro de um ambiente Linux, que é a base da tecnologia de containers atual.

Segue um tutorial da DigitalOcean de como instalar e rodar programas escritos em Go no Ubuntu [https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04](https://www.digitalocean.com/community/tutorials/how-to-install-go-on-ubuntu-20-04).

Você vai precisar de permissão root no seu Linux.

## Criando um módulo em GO

```bash
go mod init marcoswitcel/hello
```


## Rodando o código contido aqui e testando

Rodando um versão "containerized" do bash

```bash
go run main.go run /bin/bash
```

## Pendente

Abaixo alguns itens que precisam ser ajustados para que o repositório se torno uma documentação e implementação mais completa para estudo e compreensão dos conceitos apresentados no vídeo da Liz.

* Criar um filesystem e compartilhar junto com o repositório
* Ajustar os códigos para fazer uso do filesystem provido

## Montando um sistema de arquivo baseado no Alpine

* [https://alpinelinux.org/downloads/](https://alpinelinux.org/downloads/)
* [https://www.qemu.org/download/#windows](https://www.qemu.org/download/#windows)
* [https://wiki.alpinelinux.org/wiki/Qemu](https://wiki.alpinelinux.org/wiki/Qemu)
