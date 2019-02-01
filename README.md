# flagger
POSIX-like CLI Flag interpreter

## Example

See "test" folder for working examples

flagger allows much more freedom to the user when passing in flags.  It also allows flags to have multiple variations, such as a short and long form.  The following application has the available flags:
```sh
  -b, --bool, --boolean         Bool Flag
  -n, --newBool                 Another Bool Flag
  -i, --integer                 Integer Flag
  -s, --string                  String Flag
```

Flags can be used in short or long form.  Assignments for values works with either a space or an "="
```sh
$ ./goapp -b --integer 4 --string="hello"
```

Short Flags (single "-" and 1 letter) can be grouped together, any flags with assignments must come last in a group.

```sh
$ ./goapp -bi 4 -ns="hello"
```

## Usage
flagger follows the same methodology that the flags implementation in the Standard Library has.  To get started, you have to first create a "Flags" object.  It is best to use the function "New" to create these objects as this will also initialize the variables inside the object

```go
flags := flagger.New()
```

Now you are able to add new flags to it in multiple ways.

```go
// Creating a Flag will also return a pointer value
boolFlag := flags.Bool("Bool Flag", "-b", "--bool")
//It can be accessed by using *variable
fmt.Println(*boolFlag)

//You can also use the "Var" functions to manually assign a pointer
intFlag := 5
flags.IntVar(&intFlag, "Int Flag", "-i", "--integer")
```

Once all of your flags are in place, you can call the Parse() method to process the available flags.  Parse accepts a slice of strings that are the flags, and returns a slice of strings of any data not associated with a flag and an error if applicable.

```go
data, err := flags.Parse(os.Args[1:])
```

## Commands

flagger also has a sub-package named "Commands" that allows for variations of flags based on a root command given.  For instance:

```sh
$ ./goapp new -bn
  #Output of command "new"

$ ./goapp run -bni 9
  #Output of command "run"
```

Each command has its own set of flags.  To use command in an application, it is recommended that you create an object with the "New()" function

```go
cmd := commands.New()
```

To create a valid command, you must create a data type that satisfies the "Commander" interface. The most basic this could be is:

```go
type command struct {}
func (c *command) Prepare(flags *flagger.Flags) {}
func (c *command) Action(s []string, flags *flagger.Flags) error { return nil }
```

Once you have your data, you can use the function "Add" to place them into the Commands object

```go
c := command{}
cmd.Add("command", &c)
```

After all of the commands are in place, run the Parse Method to parse the flags and run the command specified.
```go
err := cmd.Parse(os.Args[1:])
```
