# Terminology

- **Overloading**: The ability to define multiple methods with the same name but different signatures. Particularly relevant when using composition to define a type that implements other types, where those other types define methods with the same name but different signatures.
- **Composition**: The ability to define a type that implements other types. This is done by embedding the other types in the new type, and then defining methods on the new type that call methods on the embedded types.