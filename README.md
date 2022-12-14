# gogen-cmd
Command line application for gogen project

---
## Installation
...

---
## Usage

### Project Initialization
Command options:

| Options | Description | Required |
|---------|-------------|----------|
| -m      | module name | true     |
| -n      | module name | true     |
| -f      | in folder   | false    |

example:
```
gogen init -m <module-name> -n <project-name> [-f]
```
or
```
gogen init --module-name <module-name> --project-name <project-name> [--in-folder]
```

## Command List
| Command    | Description        | Checklist |
|------------|--------------------|-----------|
| init       | initialize project | [x]       |
| endpoint   | endpoint module    | [x]       |
| usecase    | usecase module     | [x]       |
| repository | repository module  | [x]       |
| cmd        | cmd module         | [x]       |
| config     | config module      | [ ]       |
