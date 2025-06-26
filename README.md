# Namespaces Report
This is an experiment and Okteto does not officially support it.

This is an example on how you can use Okteto's public API to generate a CSV report of the Public Endpoints that Okteto manages.

* Create an Okteto Admin Token

* Export the token to a local variable:
```
export OKTETO_TOKEN=<<your-token>>
```
* Export the URL of your Okteto instance to a local variable:
```
export OKTETO_URL=<<your-okteto-url>>
```
* Clone the repository:
```
git clone https://github.com/okteto-community/get-public-endpoints
```
* Build the binary:
```
make build
```
* Make sure you have the correct okteto context and kubectl context in your terminal
```
okteto context use <your context>
okteto kubeconfig
```
* Run the command:
```
./public-endpoints
```
Once finished, the program will generate a CSV with all your Public Endpoints