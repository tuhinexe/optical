<p align="center">
  <img src="https://github-production-user-asset-6210df.s3.amazonaws.com/85868593/368620635-e9bcfdbe-1f4e-44cc-b3f0-9e5caffe6e17.png?X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAVCODYLSA53PQK4ZA%2F20240918%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20240918T145201Z&X-Amz-Expires=300&X-Amz-Signature=cbd77e7e782b9bf7d8fe8ed026b11ff8ded927dc46df9348aa51d0d17a5835cf&X-Amz-SignedHeaders=host&actor_id=85868593&key_id=0&repo_id=859126734" />
</p>

##

<h1 align="center">
<a href="https://github.com/TuhinBar/optical/releases"> 
  <img src="https://img.shields.io/github/v/release/TuhinBar/optical"/>
</a>
<a href="https://github.com/TuhinBar/optical/blob/main/LICENSE"> 
  <img src="https://img.shields.io/github/license/TuhinBar/optical"/>
</a>
  
  </h1>

<!--![Go Report Card](https://goreportcard.com/badge/github.com/TuhinBar/optical) -->
Optical is a CLI tool that generates a Go [Fiber](https://github.com/gofiber/fiber) project template. It is inspired by express-generator, a tool that produces a Node.js project template for the Express.js framework.

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
optical -gh <guthub-username> -init -name <your-project-name>
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
go run ./cmd
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
optical -gh TuhinBar -init -name cyber-fiber
```

It will create a new Optical project in a directory named `cyber-fiber` with `go.mod` like this :
```bash
module github.com/TuhinBar/cyber-fiber

go 1.22

require github.com/gofiber/fiber/v2 v2.52.5
```

For more detailed usage instructions, please refer to [USAGE.md](docs/USAGE.md).

## 👥 Contributing

We welcome contributions! Please see [CONTRIBUTING.md](docs/CONTRIBUTING.md) for details on how to contribute to this project.

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
