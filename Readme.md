# 🎯 Design Patterns in Go

This repository contains my notes, exercises, and implementations from a Udemy course on **Design Patterns in Go**.  
It serves as a structured reference for learning, experimenting, and applying **SOLID principles, creational, structural, and behavioral design patterns** using the Go programming language.  

![](https://images.squarespace-cdn.com/content/v1/5e10bdc20efb8f0d169f85f9/1641059684014-AL14YQXUONRLJW67G0RL/books.png)

## 📘 Table of Contents

### 🏛 SOLID Principles
- [SRP – Single Responsibility Principle](./1.SOLID/1.srp/main.go)
- [OCP – Open/Closed Principle](./1.SOLID/2.ocp/main.go)
- [LSP – Liskov Substitution Principle](./1.SOLID/3.lsp/main.go)
- [ISP – Interface Segregation Principle](./1.SOLID/4.isp/main.go)
- [DIP – Dependency Inversion Principle](./1.SOLID/5.dip/main.go)

---

### 🏗 Creational Patterns
- **[Builder](./2.Builder)**
  - [String Builder](./2.Builder/1.stringbuilder/main.go)
  - [Builder Parameter](./2.Builder/3.builderparameter/main.go)
  - [Builder Facets](./2.Builder/2.builderfacets/main.go)
  - [Functional Builder](./2.Builder/4.functionalbuilder/main.go)
- **[Factories](./3.Factories)**
  - [Factory Function](./3.Factories/1.factoryfunction/main.go)
  - [Interface Factory](./3.Factories/2.interfacefactory/main.go)
  - [Factory Generator](./3.Factories/3.factorygenerator/main.go)
  - [Prototype Factory](./3.Factories/4.prototypefactory/main.go)
- **[Prototype](./4.Prototype)**
  - [Deep Copying](./4.Prototype/1.deepcopying/main.go)
  - [Copy Method](./4.Prototype/2.copymethod/main.go)
  - [Copy Through Serialization](./4.Prototype/3.copytyhroughserialization/main.go)
  - [Prototype Factory](./4.Prototype/4.prototypefactory/main.go)
- **[Singleton](./5.Singleton)**
  - [Singleton](./5.Singleton/1.singleton/main.go)
  - [Problems with Singleton](./5.Singleton/2.problemswithsingleton/main.go)
  - [Singleton and DIP](./5.Singleton/3.singletonandDIP/main.go)

---

### 🏛 Structural Patterns
- **[Adapter](./6.Adapter)**
  - [Adapter](./6.Adapter/1.adapter/main.go)
  - [Adapter Caching](./6.Adapter/2.adaptercaching/main.go)
- **[Bridge](./7.Bridge)**
- **[Composite](./8.Composite)**
  - [Geometric Shapes](./8.Composite/1.geometricshapes/main.go)
  - [Neural Networks](./8.Composite/2.neuralnetworks/main.go)
- **[Decorator](./9.Decorator)**
  - [Multiple Aggregation](./9.Decorator/1.multipleaggregation/main.go)
  - [Decorator](./9.Decorator/2.decorator/main.go)
- **[Facade](./10.Facade)**
- **[Flyweight](./11.Flyweight)**
  - [Text Formatting](./11.Flyweight/1.textformatting/main.go)
  - [Usernames](./11.Flyweight/2.usernames/main.go)
- **[Proxy](./12.Proxy)**
  - [Protection Proxy](./12.Proxy/1.protectionproxy/main.go)
  - [Virtual Proxy](./12.Proxy/2.virtualproxy/main.go)

---

### 🤝 Behavioral Patterns
- **[Chain of Responsibility](./13.ChainOfResponsibility)**
  - [Method Chain](./13.ChainOfResponsibility/1.methodchain/main.go)
  - [Broker Chain](./13.ChainOfResponsibility/2.brokerchain/main.go)
- **[Command](./14.Command)**
  - [Command](./14.Command/1.command/main.go)
  - [Undo Operations](./14.Command/2.undooperations/main.go)
  - [Composite Command](./14.Command/3.compositecommand/main.go)
  - [Functional Component](./14.Command/4.functionalcomponent/main.go)
- **[Interpreter](./15.Interpreter)**
  - [Lexing](./15.Interpreter/1.lexing/main.go)
  - [Parsing](./15.Interpreter/2.parsing/main.go)
- **[Iterator](./16.Iterator)**
  - [Iteration](./16.Iterator/1.iteration/main.go)
  - [Tree Traversal](./16.Iterator/2.treetraversal/main.go)
- **[Mediator](./17.Mediator)**
  - [Chatroom](./17.Mediator/1.chatroom/main.go)
- **[Memento](./18.Memento)**
  - [Memento](./18.Memento/1.memento/main.go)
  - [Undo and Redo](./18.Memento/2.undoandredo/main.go)
- **[Observer](./19.Observer)**
  - [Observer](./19.Observer/1.observer/main.go)
  - [Property Observers](./19.Observer/2.propertyobservers/main.go)
  - [Property Dependencies](./19.Observer/3.propertdependencies/main.go)
- **[State](./20.State)**
  - [Classic Implementation](./20.State/1.classicimplementation/main.go)
  - [Handmade State Machine](./20.State/2.handmadestatemachine/main.go)
  - [Switch-based State Machine](./20.State/3.switchbasedstatemachine/main.go)
- **[Strategy](./21.Strategy)**
- **[Template Method](./22.TemplateMethod)**
  - [Template Method](./22.TemplateMethod/1.templatemethod/main.go)
  - [Functional Template Method](./22.TemplateMethod/2.functionaltemplatemethod/main.go)
- **[Visitor](./23.Visitor)**
  - [Intrusive Visitor](./23.Visitor/1.intrusivevisitor/main.go)
  - [Reflective Visitor](./23.Visitor/2.reflectivevisitor/main.go)
  - [Classic Visitor](./23.Visitor/3.classicvisitor/main.go)

---

## 📝 Extra Tasks
In addition to design patterns, I’ve included some Go practice tasks:
- [Intersection of 2 Slices](./tasks/1.Intersection_of_2_slices/main.go)
- [Random Number Generator](./tasks/2.random_number_generator/main.go)
- [Merge N Channels](./tasks/3.merge_n_channels/main.go)
- [Number Conveyor](./tasks/4.number_conveyer/main.go)
- [Worker Pool with Func](./tasks/5.worker_pool_with_func/main.go)
- [Custom Semaphore with WaitGroup](./tasks/6.custom_semaphore_wg/main.go)

---

## 🚀 How to Use
Clone the repo and navigate into pattern folders to explore the Go examples:

```bash
git clone git@github.com:mikebionic/go_course_notes.git
cd go_course_notes/2.Builder/1.stringbuilder
go run main.go
````

Each folder contains isolated examples that you can run independently.

---

## 📚 References

This repo is based on the **Design Patterns in Go** course on Udemy, restructured with additional notes and tasks for practice.
