---
label: Getting started
order: 200
tags: [ introduction, guide ]
---
# Getting started

The dependency mapping configuration in Parsley requires types, interfaces, and constructor methods - basically, the same things you need to wire dependencies manually. 


## Dependency mapping configuration

In Parsley, constructor methods (you may also know them as initializers or provider functions) are the centerpiece that defines the mappings between abstractions (interfaces) and implementation types.  Dependencies, i.e., the services required by the implementation type, are expressed as function parameters. The return type of a constructor method specifies the abstraction, whereby the method itself is responsible for creating the actual object instance, thus acting as the glue between implementation and interface types.

Let´s move on to the quick start tutorial, which explains the concept in detail.