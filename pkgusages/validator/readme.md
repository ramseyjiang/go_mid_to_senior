This pkg is used to validate request. It can be used for struct and single value.

The most important thing is that you should not have any space after required keywords.

Not work: "validate.Var(myEmail, "required, email")"
Work: "validate.Var(myEmail, "required,email")"
