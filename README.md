# tcgen: test case generator for coding problems

[Tiếng Việt](./README.vi.md)

Generate test cases for coding problems.

## Download

https://github.com/hoangvvo/tcgen/releases

## Instruction

https://user-images.githubusercontent.com/40987398/154596409-4c8c2dd9-1f42-4b72-851c-22b8812c56c3.mp4

Craete 2 code files for the following purposes:

- **Case generation**: When run, this code should write the input file.
- **Case solver**: When run, this code should read the input file and write to the output file.

The input and output files should have the extensions listed in the config `testcase.extensions`.

Run the program, then type in the location of the above two files and the number of run `n`.

After `n` runs, the program will copy all input and output files into folder with the naming of `testcase.output`.

## Configuration file

An optional configuration file named `tcgen.yml` may placed in the same folder with `tcgen`. See the sample config file: [tcgen.yml](./gen/tcgen.yml).

- `testcase.extensions`: Files with such extensions will be copied to the result folder. (eg: `["out", "inp", "txt"]`)
- `testcase.output`: result folder to store test cases, character `*` should be included to be later replaced with formatted test number. (eg: `./result/TEST*`)
- `language.[key]`: Configure a language. See below.

## Configure a language

To configure a language, add a key with any name with the following properties:

- `name`: Name of the language
- `extensions`: Array containing file extensions to recognize this language
- `compile`: (only if the language needs compilation) an array of command line arguments to compile the code. Must include `$SOURCE` (to be replaced with the source path) and `$OUTPUT` (to be replaced with the output binary path). Languages like Python does not have this step.
- `run`: an array of command line arguments to run the binary. Must include `$OUTPUT` (to be replaced with the output binary).

See the sample config [tcgen.yml](./gen/tcgen.yml) to learn more.

## Write case generation and solver code

See [example](./example)

### Write case generation code

The code must write an input file.

This code will be called with the run number as the first argument (one right after caller name), starting with 1. This value can be used for different purposes.

An example making use of case number for case difficulty:

```cpp
// Generate the first 5 cases to be easy and the rest to be difficult
int main(int argc, char** argv)
{
  int caseNumber = atoi(argv[1]);
  if (caseNumber <= 5) generateEasyTestCase();
  else generateDifficultTestCase();
}
```

### Write case solver code

Case solver will read the generated input file and write to the output.

Similarly the code will be called with the run number as the first argument.

## LICENSE

[MIT](LICENSE)
