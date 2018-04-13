# BitWallet

BitWallet is the result of the final project of 2ºDAW of Sergio Martinez Guerrero.

The project will be based on the idea of ​​building a web application, whose functionality will be established as a virtual cryptocurrency wallet.

 [here](#technologies)

## Technologies

The project is built with:

### FrontEnd
#### <img src="https://cdn.svgporn.com/logos/react.svg" width="18"></img> ReactJS
A javascript library open source property of Facebook, used to render views and control events, with the Virtual DOM makes the user experience smooth and enjoyable.
#### <img src="https://cdn.svgporn.com/logos/redux.svg" width="18"></img> Redux
Redux is an a javascript library to control the application state, this means that communication between client and server and testing actions are more easy, Redux librery are more effective on Redux projects, but it can also be used in AngularJS projects

Redux bases on three principles:

- The entire state of your application is contained in a single store
- The only way to modify the state is to issue an action that indicates that it changed
- To control how the store is modified by the actions, pure reducers are used
### BackEnd
#### <img src="https://cdn.svgporn.com/logos/gopher.svg" width="18"></img> Golang
GoLang is an open source programming lenguage developed by Google, it is a highly typed language and it is also a compiled language.

Advantages over other web programming languages in backend environment:
- It is easy to learn
- Supports thousands of simultaneous connections
- Speed and performance

On this project GoLang is used to do a backend and API RESTful

### Design
#### <img src="https://cdn.svgporn.com/logos/css-3.svg" width="18"></img> CSS3
Used to make website more beautiful and upgrade his look and feel, using transitions, transformations, and animations references CSS3
#### <img src="https://cdn.svgporn.com/logos/bem.svg" width="18"></img> BEM
Methodology used to oranize own code on components and code sharing on frontend.

BEM is based on Block, Element and Modifier
### Database
#### <img src="https://cdn.svgporn.com/logos/mysql.svg" width="18"></img> MySQL
Use MySQL to run own database and make queries, on MySQL they will be used triggers and procedures.

### Deploy and config
#### <img src="https://cdn.svgporn.com/logos/webpack.svg" width="18"></img> Webpack
A bundling system to prepare development on a web application, they have advantages in front Browserify, is based on modulable system, 
will be used to compile the SCSS styles, convert the code from ES6 to ES5 using babel, and gives the possibility to run the program in both development mode and production mode with her webpack.config.js
#### <img src="https://cdn.svgporn.com/logos/bash.svg" width="18"></img> Shell Scripts
Used to deploy own database with shell script, and make final deploy on VPS, all configuration, deploy and run it web application
#### <img src="https://cdn.svgporn.com/logos/babel.svg" width="18"></img> Babel Transpiler for ES6 
As the project is writing its frontend with ES6, we need a transpiler since it is not yet supported 100% by all browsers
https://kangax.github.io/compat-table/es6/

## Getting started
### Run development mode
FrontEnd
```
npm start
```
Backend
```
go run main.go
```
### Run on production mode
FrontEnd
```
npm run-script build
```
BackEnd
```
go build main.go
```
