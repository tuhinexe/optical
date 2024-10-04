<p align="center">
  <img src="https://p82.cooltext.com/Rendered/Cool%20Text%20-%20OPTICAL%20467187224148174.png" />
</p>

##

<h1 align="center">
  <a href="https://pkg.go.dev/github.com/TuhinBar/optical">
    <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat-square">
  </a>
<a href="https://github.com/TuhinBar/optical/releases"> 
  <img src="https://img.shields.io/github/v/release/TuhinBar/optical"/>
</a>
<a href="https://github.com/TuhinBar/optical/blob/main/LICENSE"> 
  <img src="https://img.shields.io/github/license/TuhinBar/optical"/>
</a>
<a href="https://github.com/TuhinBar/optical/actions"> 
  <img alt="GitHub Actions Workflow Status" src="https://img.shields.io/github/actions/workflow/status/TuhinBar/optical/.github%2Fworkflows%2Fbuild.yml">
</a>

  
  </h1>

<!--![Go Report Card](https://goreportcard.com/badge/github.com/TuhinBar/optical) -->
Optical is a CLI tool that generates a Go [Fiber](https://github.com/gofiber/fiber) project template. It is inspired by [express-generator](https://expressjs.com/en/starter/generator.html), a tool that produces a Node.js project template for the Express.js framework.

It is named Optical because it works with Fiber and combines 'fiber-optic'. Silly.

🔵 Although there are big-scale CLI tools exist like [create-go-app](https://github.com/create-go-app/cli) to help you create Full-Stack Apps using Go and JS/TS, optical is just a fiber-dependent CLI tool to generate API/Backend only.

---

## ⚙️ Installation

If you do not have Go installed on your device, you are required to install it to run this tool.


You can install Optical CLI using Go:

```bash
go install github.com/TuhinBar/optical@latest
```

Or you can use our installation script:

```bash
curl -sSL https://raw.githubusercontent.com/TuhinBar/optical/main/scripts/install.sh | bash
```

## ⚡Usage

To create a new Optical project:

```bash
optical -create
```

```bash

Optical CLI version 0.7.0
┃ What is the name of your project?(e.g 'my-fiber-api','./')                                     
┃ A new folder with this name will be created.                                                   
┃ > <your-project-name>                                                                                          
                                                                                                 
  Enter your GitHub username.                                                                    
  This is required to create the go.mod file                                                     
  > <your-gitub-username>                                                                                           
                                                                                                 
  Do you have air installed ?                                                                    
  Air is required for auto-reload                                                                
    Yes                                                                                          
  > No                                                                                           
                                                                                                 
enter next
```
or

To get help and see all the flags:
```bash
optical -h
```

After creating the project run this:
```bash
cd <your-project-name>
go mod tidy
air
```

It will run your Fiber App and the output should look like this:
<p align="center">
  <img src="https://github.com/user-attachments/assets/c7d06b42-bad8-46ec-9301-c4b0c1b637b9" />
</p>

Optical generates the starting files for your project. You can configure your project folder and files according to your requirements.

To set up your project further, refer to the [Fiber Official Docs](https://docs.gofiber.io/)

## 🔷 Example

If you run a command like:
```bash
optical -create
```
```bash
Optical CLI version 0.7.0
┃ What is the name of your project?(e.g 'my-fiber-api','./')                                     
┃ A new folder with this name will be created.                                                   
┃ > cyber-fiber                                                                                             
                                                                                                 
  Enter your GitHub username.                                                                    
  This is required to create the go.mod file                                                     
  > TuhinBar                                                                                             
                                                                                                 
  Do you have air installed ?                                                                    
  Air is required for auto-reload                                                                
    Yes                                                                                          
  > No                                                                                           
                                                                                                 
enter next
```

It will create a new Optical project in a directory named `cyber-fiber` with `go.mod` like this :
```bash
module github.com/TuhinBar/cyber-fiber

go 1.22

require github.com/gofiber/fiber/v2 v2.52.5
```

Please refer to [USAGE.md](docs/USAGE.md) for more detailed usage instructions.

## 👥 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](docs/CONTRIBUTING.md) for details on how to contribute to this project.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
