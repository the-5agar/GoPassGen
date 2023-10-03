# GoPassGen

## Overview

GoPassGen is a command-line password generator and manager tool for generating strong and secure passwords. It allows you to create new passwords with customizable strength and associate them with usernames, websites, and email addresses. Additionally, you can search for stored passwords based on websites.

## Installation

To use GoPassGen, follow these steps:

   ```bash
   git clone https://github.com/the-5agar/GoPassGen.git
   cd GoPassGen
   go build
   ```

## Usage

### Generating a New Password

To generate a new password, use the following command:

   ```bash
   ./GoPassGen new [flags]
   ```
#### Flags:
  ```bash
  -e, --email string: Email associated with the password.
  -m, --secret-pass string: Master password for encryption and decryption.
  -s, --strength int: Password strength (length) (default 8).
  -u, --username string: Username associated with the password.
  -w, --website string: Website associated with the password.
  ```

**Example:**

Generate a new password with custom length, username, and website:

  ```bash
  ./GoPassGen new -s 12 -u myusername -w example.com -m secretpass 
  ```

### Searching for Stored Passwords
To search for stored passwords, use the following command:
  
  ```bash
  ./GoPassGen search [flags]
  ```

#### Flags

  ```bash
  -h, --help: Show help for the search command.
  -m, --secret-pass string: Master password for encryption and decryption.
  -w, --website string: Search based on website. If no website is provided, it lists all websites.
 ```

**Examples**

Search for stored passwords for a specific website:
  ```bash
  ./GoPassGen search -w example.com -m secretpass
  ```
List all stored websites and their associated passwords:

 ```bash
  ./GoPassGen search -m secretpass
  ```

For more information on specific commands and their usage, use the -h or --help flag with the command.
