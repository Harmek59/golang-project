# golang-project

## Opis projektu

Celem projektu jest stworzenie prostej gry platformowej 2D przy użyciu języka Go w wersji 1.20 w połączeniu z OpenGL. Gra będzie zainspirowana grą [Google Chrome Dinosaur Game](https://chrome://dino/). W projekcie użyjemy wzorca architektnonicznego entity-component-system (ECS). Bazuje na podejściu zwanym kompozycją ponad dziedziczenie. ECS zapewnia izolację poszczególnych elementów, przejrzystość, modularność i łatwość rozszerzania kodu. W projekcie planujemy użyć następujących zewnętrznych bilbiotek:
- https://github.com/go-gl/gl
- https://github.com/go-gl/glfw
- https://github.com/go-gl/mathgl
 
## Zespól

- Michał Kacprzak
- Piotr Harmuszkiewicz

## Kompilacja
```shell
go mod tidy && go build && ./game2d
```
### Windows
 Biblioteki użyte w projekcie są stworzone w języku C i wymagają kompilator cgo (zwykle gcc). System Windows nie ma kompilatora gcc, więc, aby program poprawnie zadziałał, należy go zainstalować. Zalecamy użyć [tdm-gcc](https://jmeubank.github.io/tdm-gcc/): [link do pobrania](https://github.com/jmeubank/tdm-gcc/releases/download/v10.3.0-tdm64-2/tdm64-gcc-10.3.0-2.exe).
### Ubuntu/Debian
 Wymagany jest pakiet: **libgl1-mesa-dev**

## Sterowanie
### Esc - quit
### Space - jump
### Enter - restart after GameOver

