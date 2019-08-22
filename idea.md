# IDEAS

## Operation

- create
  - register a file
    - with symbolic link or file entity
    - file name is `<id>_<tags>.<ext>`

- read
  - get file path
    - input: tags
    - output: list of file paths(mutch all)
  - get file path
    - input: id
      output: file path

- update
  - tag update
  - change of file entity

- delete
  - delete with id
  - delete tag

## Command line

```bash
$ tsearch get --tag test
/path/to/dir/ba456ed_test_foo.txt
/path/to/dir/134fad6_bar_test.mp4

$ tsearch register /myfile.txt --tag test --tag text
# return registered path
/path/to/dir/cd5ad51_test_text.txt
# or return id
cd5ad51
```
