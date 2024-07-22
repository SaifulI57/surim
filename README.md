
# Surim

**Surim** is a tool for merging Suricata rule files into a single consolidated output file. This command-line utility is designed to simplify the management of Suricata rules by merging multiple files into one.

## Features

- **Version Display**: Shows the current version of the tool when no input or output flags are provided.
- **Automatic Directory Creation**: Creates the output directory if it does not already exist.
- **Efficient File Handling**: Reads and merges rule files from the specified input directory.
- **Detailed Error Messages**: Provides clear and informative messages for various errors.

## Installation

To use Surim, you need to have Go installed on your system. You can build Surim from source by cloning the repository and running the build command:

```bash
git clone https://github.com/yourusername/surim.git
cd surim
go build -o surim
```

## Usage

### Basic Command

To merge Suricata rule files from an input directory and save the result to an output file:

```bash
./surim --input /path/to/rules --output /path/to/output.rules
```

### Options

- `--input`, `-i`: Specifies the directory containing Suricata rule files to be merged.
- `--output`, `-o`: Specifies the file path where the merged rules will be saved.

### Example

```bash
./surim --input /etc/suricata/rules --output /etc/suricata/merged.rules
```

In this example, Surim will read all `.rules` files from `/etc/suricata/rules` and merge them into `/etc/suricata/merged.rules`.

### Version Information

To display the version of Surim:

```bash
./surim
```

If no input or output flags are provided, the tool will print the version information.

## Error Handling

- **Error reading file**: Indicates issues encountered while reading individual rule files.
- **No Suricata rules found**: No `.rules` files were found in the specified input directory.
- **Failed to create directory**: An error occurred while creating the output directory.
- **Failed to write output file**: An error occurred while writing the merged rules to the output file.

## License

Surim is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
---

