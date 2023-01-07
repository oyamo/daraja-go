<p align="center">
<img src="./assets/daraja-go.png" width="500" alt="# daraja-go"/><br><br>
<b>STILL IN BETA!</b><BR><BR>
<img src="https://img.shields.io/badge/UNIT%20TESTING-PASSING-green?style=flat" alt="Unit-testing">
<img src="https://img.shields.io/badge/GO%20REPORT-A+-green?style=flat" alt="GO REPORT" />
<img src="https://img.shields.io/badge/DOCS-GODOCS-green?style=flat" alt="DOCS" />
</p>

<strong>daraja-go</strong> is a library that simplifies integration with the Daraja API in Go
applications. The Daraja API allows developers to build applications that interact with M-PESA, a popular 
mobile money service in Kenya. With daraja-go, developers can easily build M-PESA-powered applications and
bring the power of the Daraja API to their users. Released under the MIT License, daraja-go is easy to use 
and integrate into any Go project.

### Highlighted features
* **Automatic re-authentication:** daraja-go automatically reauthenticates with the Daraja API when the expiry time is reached.

* **Automatic generation of cipher text** for secret credentials: daraja-go automatically generates a cipher text of secret credentials using RSA-OAEP.
* **Singleton instance of DarajaApi struct:** daraja-go creates a singleton instance of the DarajaApi struct, which represents an instance of the Daraja API and contains various methods for interacting with the API.

### Getting started
* To get started with and enjoy daraja-go, you need to have a Daraja account. If you don't have one, you can create one [here](https://developer.safaricom.co.ke).
* Create a new application and generate a consumer key and secret. You can do this by clicking on the `My Apps` tab and then clicking on the `Create App` button.
*  Install daraja-go using the command below:
    ```bash
    go get github.com/oyamo/daraja-go
    ```
*  Import daraja-go in your project:
    ```go
    import "github.com/oyamo/daraja-go"
    ```
### Usage
*  Create a new instance of the DarajaApi struct:
    ```go
    darajaApi := daraja.NewDarajaApi(consumerKey, consumerSecret, environment)
    ```
    *  `consumerKey` is the consumer key generated for your application.
    *  `consumerSecret` is the consumer secret generated for your application.
    *  `environment` is the environment you want to use. It can either be `DarajaApi.SANDBOX` or `DarajaApi.PRODUCTION`.


