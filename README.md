# Cleansys [![Go](https://img.shields.io/badge/go-1.17-blue)](https://golang.org/doc/go1.17) [![Release](https://img.shields.io/badge/release-1.0.0-success)](https://github.com/Lapp-coder/cleansys/releases) 
![image](images/cleansys.png)
***
## Installation
1. Clone the repository
  ```
  $ git clone https://github.com/Lapp-coder/cleansys
  ```
2. Go to the directory of the utility
  ```
  $ cd cleansys
  ```
3. Execute the make command depending on the architecture of your processor
* For M1-chip:
  ```
  $ chmod +x build/bin/cleansys-arm64 && make setup-arm64  
  ```
* For Intel-chip:
  ```
  $ chmod +x build/bin/cleansys-amd64 && make setup-amd64
  ```
## Usage
Just run this command
```
$ cleansys 
```

## Configuration
The configuration file is located at $HOME/.config/cleansys/config.json 
