# TraceGit

## Introduction
This is a TUI application that shows all your local git repositories in one place

## Demo
![](https://github.com/Cwjiee/tracegit/blob/main/tracegit_demo.gif)

## Installation

1. set system type

```bash
systype=$(uname -m)
```

2. install the package
### For Mac users
```bash
curl -L https://github.com/Cwjiee/tracegit/releases/latest/download/tracegit_Darwin_"$systype".tar.gz > tracegit.tar.gz
```
### For Linux/Wsl users
```bash
curl -L https://github.com/Cwjiee/tracegit/releases/latest/download/tracegit_Linux_"$systype".tar.gz > tracegit.tar.gz
```

3. extract the tar file
```bash
tar -xzf tracegit.tar.gz
```

4. move file to binaries
```bash
sudo mv tracegit /usr/local/bin
```
## Usage
1. execute the command
```bash
tracegit
```

2. when you first use the tool, it will prompt you to enter your code directory (path where you store your repos)

#### Example
```
# Linux/Wsl
/home/weijie/code

# Mac
/Users/weijie/code
```
3. input your path to the prompt

4. use `tracegit` anywhere in your terminal!

```bash
tracegit
```

