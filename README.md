# GoDeps 

`godeps` is tool to retrieve and visualize Go source code dependency trees and find not allowed import cycles

## Install

```sh
go install github.com/cgxarrie-go/godeps@latest
```

## Usage

```sh
godeps get <package-name>
```
 print all dependencies tree for the stated package.


```sh
godeps cycle <package-name>
```
 print all dependencies tree involved in an import cycle.



