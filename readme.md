goctl api go -api aula3.api -dir .

----------------- DEPLOY EC2 COMANDOS------------
login:ec2-user (amazon Linux) =================

sudo yum update -y

sudo yum install -y git

cd /usr/local

sudo wget https://go.dev/dl/go1.22.2.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz

echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc

source ~/.bashrc

go version

cd ..
cd ..

mkdir -p ~/projetos

cd ~/projetos

git clone ...



sudo yum install tmux -y

Inicie uma nova sessão tmux=======================

tmux

entrar na pasta do projeto que tem o arquivo .go principal e rodar o projeto======================

go run seu_arquivo.go

ctrl+b   seguido de  d
  
Para reanexar à sessão tmux:

tmux attach
