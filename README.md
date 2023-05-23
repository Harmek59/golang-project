# golang-project

## Opis projektu

Celem projektu jest stworzenie prostej gry platformowej 2D przy użyciu języka Go w wersji 1.20 w połączeniu z OpenGL. Gra będzie zainspirowana grą [Google Chrome Dinosaur Game](https://chrome://dino/). W projekcie użyjemy wzorca architektnonicznego entity-component-system (ECS). Bazuje na podejściu zwanym kompozycją ponad dziedziczenie. ECS zapewnia izolację poszczególnych elementów, przejrzystość, modularność i łatwość rozszerzania kodu. W projekcie planujemy użyć następujących zewnętrznych bilbiotek:
- https://github.com/go-gl/gl
- https://github.com/go-gl/glfw
- https://github.com/go-gl/mathgl
 
 ## Zespól

- Michał Kacprzak
- Piotr Harmuszkiewicz

```shell
go mod tidy && go build && ./game2D
```
