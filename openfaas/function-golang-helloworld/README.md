
docker build -t coolbeevip/openfaas-function-golang-helloworld .

docker push coolbeevip/openfaas-function-golang-helloworld

faas-cli deploy stack.yml 

ab -n 50000  -c 50 -p 'data.json' -T 'application/json' http://127.0.0.1:31112/function/golang-helloworld
