# Workshop 2 - API REST in Go with Gin

In this workshop we will learn how to use the main functions of the Gin package in Go and how to create a REST API.

> :bulb: A quick [documentation](https://searchapparchitecture.techtarget.com/definition/RESTful-API) about REST API.

## Step 0: Initialization

All the required information to install the workshop's dependencies are given in the [SETUP.md](./SETUP.md)

## Step 0.5: The codebase

We designed a boilerplate for this project, it will be useful to start this project, but also to summarize all the required parts of an API
We will have :

- `controllers`: this is where you will design your routes endpoint. it will often be a wrapper that calls other API or the database itself
- `middlewares`: those are functions that can be called to print logs ,check access rights before entering a controller or define a default endpoint
- `routes`: the core of the router, setting up all the routes handler, with their middlewares and controllers

There are plenty of other important package you may need in a real API, but these are the main one you can be sure you'll need one time or another. Now let's code.

Add a route on the endpoint `/`:
- It uses the `GET` method
- It responds `"hello: world"` in a json format

> You can test it with `http://localhost:8080/` in a browser or with `postman`

## Step 1: Let's get started

You are now fine with the creation of endpoint.
Next step is to create different request method.
There is many [request method](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods), you already experimented the `GET` request.

For this step, we will use some others' method:
- a `GET` request `/item/template`: to get the template of the item as follows.
```json
{
  "name": "name",
  "description": "description",
  "price": 0
}
```

- a `POST` request `/item/display`: to display the item passed in the body as follows.
<img src="./assets/request-item-display.png">


also try to implement a default endpoint for the URL's that doesn't match any routes. (404 notfound error)



> This [tutorial](https://blog.logrocket.com/making-http-requests-in-go/) may help you

> You should create a structure item with the following fields:
> - name
> - description
> - price
> 
> and bind the request body to the structure to retrieve it (do it in a structure folder).
<details>
  <summary>See how to request with postman :satellite:</summary>

Enter your URL and the method you which to use in the title bar and click `Send`.

![Seek](../../.github/go-http/seek.png)

Then the result (if there is any) will be printed out at the bottom.

![Result](../../.github/go-http/result.png)

</details>

## Step 2: Database interaction

For the following we won't use a real database, but we will use a simple json file as the database.
(The configuration would have been too long to explain and to do here)
You can see the database.json in the database folder.


- Create a `GET` request `/item`: to retrieve all the items in the database.
- Create a `POST` request `/item`: to create a new item in the database.

> You must fill the `database/database.go` functions you need (`GetDatabase` & `AddItemInDB`) and call.

> Look at the [documentation](https://tutorialedge.net/golang/reading-writing-files-in-go/) to interact with the file

## Step 3: To go further (BONUS)

To go further you can add some others routes :
- a `DELETE` request `/item`: to delete an item in the database.
- a `PUT` request `/item`: to update an item in the database.
- a `POST` request `/items`: to create multiple items in the database.

## Go

If you still want to toy around with Go and all its usages, you can take a look at other go features like:
- `go routines`
- `interfaces`

You can also use a real database with this ORM packages:
- [Ent](https://github.com/facebookincubator/ent)
- [postgres/ent](https://entgo.io/docs/crud/)

## Authors

| [<img src="https://github.com/Doozers.png?size=85" width=85><br><sub>Ismaël FALL</sub>](https://github.com/Doozers) |
|:-------------------------------------------------------------------------------------------------------------------:|
<h2 align=center>
Organization
</h2>
<br/>
<p align='center'>
    <a href="https://www.linkedin.com/company/pocinnovation/mycompany/">
        <img src="https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white">
    </a>
    <a href="https://www.instagram.com/pocinnovation/">
        <img src="https://img.shields.io/badge/Instagram-E4405F?style=for-the-badge&logo=instagram&logoColor=white">
    </a>
    <a href="https://twitter.com/PoCInnovation">
        <img src="https://img.shields.io/badge/Twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white">
    </a>
    <a href="https://discord.com/invite/Yqq2ADGDS7">
        <img src="https://img.shields.io/badge/Discord-7289DA?style=for-the-badge&logo=discord&logoColor=white">
    </a>
</p>
<p align=center>
    <a href="https://www.poc-innovation.fr/">
        <img src="https://img.shields.io/badge/WebSite-1a2b6d?style=for-the-badge&logo=GitHub Sponsors&logoColor=white">
    </a>
</p>

> :rocket: Don't hesitate to follow us on our different networks, and put a star 🌟 on `PoC's` repositories.
