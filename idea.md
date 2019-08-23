# IDEAS

## Operation

- create
  - register a file
    - input: absolute path of the file
    - output: none(?)
    - action: save object to database

- read
  - get file path
    - input: tags
    - output: list of file paths(mutch all)

- update
  - tag update

- delete
  - delete file from database
  - delete tag (equal to update?)

## Command line

```bash
$ tsearch get --tag test
/home/user/file1.txt
/opt/dir/movie.mp4

$ tsearch register /myfile.txt --tag test --tag text
ok
```
