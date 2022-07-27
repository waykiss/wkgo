GoApp is the project to create a complete backend for any application, inspired by Django.

The project goal is give a complete tool to start up quickly API with the common stack used by general applications.

The main architecture is inspired in Django, everything is an "App", use apps created by comunity and/or create your 
private own app 

GoApp brings things ready to use:

- Agnostic webserver - User http package, fiber, gin or create your own implementation for favority webserver
- Define input and output data model and have graphql, rest(with documentation) and grpc controllers ready to use  
- Utils package to deal with numbers, strings, slice, database 
- Simple interface to attach application and then create anything you want
