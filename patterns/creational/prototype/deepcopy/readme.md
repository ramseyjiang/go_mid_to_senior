The Clone(), in the address.go, is just used to clone a single address.

When you use the Clone in the address.go, if it has several address such as home address and work address, you should use Clone several times.

The Clone(), in the person.go, is used to clone serialization, it can clone one object into another.

Hence, even it has several address need to clone, it can use serialization clone many objects once.
