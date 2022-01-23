# go-demo

## Steps

### Go Project setup

    1. Create your project folder.
    2. Use command prompt/terminal to navigate to the directory
    3. Run "go mod init github.com/examples/tfprovider" to initilise your go module which you are buiding in this case a Terraform provider

### Get Terraform modules

    $ go run main.go
    main.go:4:2: no required module provides package github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema; to add it:
            go get github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema
    main.go:5:2: no required module provides package github.com/hashicorp/terraform-plugin-sdk/v2/plugin; to add it:
            go get github.com/hashicorp/terraform-plugin-sdk/v2/plugin

    go get github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema
    go get github.com/hashicorp/terraform-plugin-sdk/v2/plugin

    line 112 you can see type ConfigureContextFunc func(context.Context, *ResourceData) (interface{}, diag.Diagnostics)

### Build the provider
    ./build.sh
    There has been difference in widnows, linux and MacOs plugin path which is very important to remember.

### Initialise the terraform
    terraform.exe init -plugin-dir=/c/Users/avinash.c.srivastava/.terraform.d/plugins/


### Trigger WebService hosting simple User API

curl -X POST -H "Content-Type: application/json" -d '{"ID": 0, "FirstName": "Avinash", "LastName": "Srivastava"}' http://localhost:3000/users



## Helper
    https://github.com/hashicorp/terraform-provider-hashicups