# tcgen: test case generator for coding problems

[Tiếng Việt](./README.vi.md)

Generate test cases for coding problems.

## Download

https://github.com/hoangvvo/tcgen/releases

## Instruction

https://user-images.githubusercontent.com/40987398/154596409-4c8c2dd9-1f42-4b72-851c-22b8812c56c3.mp4

Create 2 code files for the following purposes:

- **Case generation**: When run, this code should write the input file.
- **Case solver**: When run, this code should read the input file and write to the output file.

The input and output files should have the extensions listed in the config `testcase.extensions`.

Run the program, then type in the location of the above two files and the number of run `n`.

After `n` runs, the program will copy all input and output files into folder with the naming of `testcase.output`.

## Configuration file

An optional configuration file named `tcgen.yml` may placed in the same folder with `tcgen`. See the sample config file: [tcgen.yml](./gen/tcgen.yml).

### testcase.extensions

Files with such extensions will be copied to the result folder.

Example: `["out", "inp", "txt"]`

### testcase.output

Result location to store the generated the test cases, character `*` may be used to be later replaced with run number.

For example, using `./result/TEST*` will resulted in the following:

```
./result/TEST001/file.INP
./result/TEST001/file.OUT
./result/TEST002/file.INP
./result/TEST002/file.OUT

and more
```

If the character `*` is not included, you should have a strategy in case generation and solver code to make different file names (possibly by using the application arguments). Otherwise, these file may be overwritten.

After every run, the output folder will not be deleted. Make sure to do it yourself if you want a fresh run.

### testcase.{key}

Configure a language. See below.

## Configure a language

To configure a language, add a key with any name with the following properties:

- `name`: Name of the language
- `extensions`: Array containing file extensions to recognize this language
- `compile`: (only if the language needs compilation) an array of command line arguments to compile the code. Must include `$SOURCE` (to be replaced with the source path) and `$OUTPUT` (to be replaced with the output binary path). Languages like Python does not have this step.
- `run`: an array of command line arguments to run the binary. Must include `$OUTPUT` (to be replaced with the output binary).

See the sample config [tcgen.yml](./tcgen.yml) for some examples.

## Write case generation and solver code

See [example](./example)

### Write case generation code

The code file must write to an input file. It can write the file using any names (as long as its extensions listed in the config `testcase.extensions`).

The file will be called with two arguments: `run number` (starting with 1) and `total`, both can be utilized for different purposes.

An example making use of run number for case difficulty:

```cpp
// Generate the first 50% cases to be easy and the rest to be difficult
int main(int argc, char** argv)
{
  double runNumber = atoi(argv[1]);
  double total = atoi(argv[2]);
  if (runNumber/total <= 0.5) generateEasyTestCase();
  else generateDifficultTestCase();
}
```

### Write case solver code

The code file must read the generated input file and write to the output. Similarly, it can read from and write to files using any names (as long as its extensions listed in the config `testcase.extensions`).

The file will also be called with two arguments: `run number` (starting with 1) and `total`.

## LICENSE

[MIT](LICENSE)
