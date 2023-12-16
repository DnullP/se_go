# What's this?

This is a backend for the SE work. You can check the API doc on [here](https://apifox.com/apidoc/shared-84283f24-2136-447e-80fb-706b21aeb387) which is written by [ApiFox](https://apifox.com/). It was developed by golang. The framework is gin and the database is sqlite3.

# About Auther
The backend is totally written by QiuXiangkai (Dnull), and you can check the commit history to see the details.

# How to run
1. You need to install the Golang env first, which version is 1.21.4.

2. Run the make to build the project to all platforms. 

    `make build-all`

3. Check and run the binary file in the build directory.

# Develop Convention

This project is layered into three layers: controller, service, and model.

- controller: In this layer, we will handle the request and response. The controller will call the service to do the business logic.

- service: In this layer, we will handle the business logic. The service will call the model to do the database operation.

- model: In this layer, we will handle the database operation. The model will call the database driver to do the database operation.

Thanks to the three-layer architecture, we can easily change the API or the schema and don't need to change the service.

If you want to add any service, you need to first define the interface in the `service`, and then implement it in the `service/impl` directory. 

`model/restAPI` is used for receiving the data and sending the response. You can NOT change it unless the API is changed.

`model/DTO` is used to transfer the data between different layers, you can define any if you need.

`schema` is used for defining the database schema. Changing it will cause the database to be recreated.

`mock` is used for mocking the air conditioners.

All the code should follow the [Golang Code Style from Google](https://google.github.io/styleguide/go/)