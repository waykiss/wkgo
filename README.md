GoApp is a Framework that implement Hexagonal Architecture organized by App (inspired by Django)

GoApp is like a Django for Go that implement Hexagonal Architecture

Goals

- Agnostic webserver(adapter) - use fiber, gin or create your own implementation for favority webserver
- Have a good common packages ready to use in real applications, helpers functions, validations, database

What's an App?
 App is like a module, just implement adpater interface and app interface, and you're ready to add your add to the 
framework
