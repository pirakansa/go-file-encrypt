# File Encrypt

`fenc` is a Go-based tool for securely encrypting and decrypting files using AES-GCM encryption.

## Features

- **AES-GCM Encryption**: Ensures data integrity with authenticated encryption mode.
- **Simple Interface**: Easily perform encryption and decryption.

## Requirements

- Go 1.16 or later

## Installation

1. Clone this repository:

   ```bash
   $ git clone https://github.com/your-username/file-encrypt.git
   $ cd file-encrypt
   ```

2. Build the `fenc` binary:

   ```bash
   $ make build
   ```

## Usage

This tool provides a command-line interface for encrypting and decrypting files.

### Encrypt a File

To encrypt a file, use the following command:

```bash
$ fenc enc -i <input_file> -o <output_file>
```

- `-i`: Path to the plaintext file to encrypt.
- `-o`: Path to save the encrypted file.

### Decrypt a File

To decrypt a file, use the following command:

```bash
$ fenc dec -i <input_file> -o <output_file>
```

- `-i`: Path to the encrypted file to decrypt.
- `-o`: Path to save the decrypted file.

## Testing

Run unit tests with the following command:

```bash
$ make test
```
