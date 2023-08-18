<br />
<p>
    <a href="" target="_blank">
      <img
        src="https://github.com/YunusEmreAlps/Orion/blob/master/assets/orion.png?raw=true"
        alt="Orion"
        width="100%"
      />
    </a>
</p>
Orion is an open-source error detection system designed to monitor and detect errors in logs.With the release of version 1.0.0, Orion has gained new capabilities. This version is equipped to retrieve error information from a database and promptly alert the designated target team by sending error notifications via email.

## Orion Meaning for this project
**Mythological reference**: Orion is a prominent constellation in the night sky, named after the mythological hunter Orion. It is often depicted as a powerful figure with a bow and arrow. The name was chosen to symbolize the system's ability to hunt down and target errors in logs with accuracy and precision.

**Stellar association**: Orion is one of the most recognizable constellations visible from the Earth. It contains several prominent stars, such as Betelgeuse and Rigel. The name was chosen to signify the system's capability to identify and highlight critical errors or issues, just like the bright stars in the Orion constellation stand out in the night sky.

**Strength and reliability**: Orion is a strong and durable character in mythology, known for his physical prowess. The name was selected to convey the system's robustness and reliability in detecting errors. It suggests that the system is designed to endure and handle large volumes of logs while consistently identifying and reporting issues.

**Futuristic or technological connotations**: Orion is sometimes associated with the future and space exploration due to its connection with science fiction literature and movies. Choosing the name "Orion" might reflect the project's aspiration to be at the forefront of error detection technology or to evoke a sense of cutting-edge innovation.

## Prerequisites

- Go 1.15+
- PostgreSQL
- Docker & Docker Compose
- SMTP

## Quick start

We can run this **ORION** project with or without Docker. Here, I am providing both ways to run this project.

- Clone this project

```bash
# Move to your workspace
cd your-workspace

# Clone this project into your workspace
git clone ...

# Move to the project root directory
cd orion
```

### Run without Docker

Run the following command to execute the Go program:

- Make sure PostgreSQL is running and accessible with the credentials you provided in the .env file.
- Open a terminal or command prompt and navigate to the root directory of your project.
- Create a file `.env` similar to `.env.example` at the **/config directory** with your configuration.
- Install `go` if not installed on your machine.
- Install `PostgreSQL` if not installed on your machine.
- Important: Open the `.env` file and modify the values of `DB_HOST`, `DB_USER`, and `DB_PASSWORD` to match your PostgreSQL configuration. Update any other configuration variables if necessary.
- Run `go run main.go`.

### Run with Docker

- Create a file `.env` similar to `.env.example` at the **/config directory** with your configuration.
- Install Docker and Docker Compose.
- Run `docker-compose up -d`.

## Project structure

### `config`

Configuration. First, `config.yml` is read, then environment variables overwrite the yaml config if they match.
The config structure is in the `config.go`.
The `env-required: true` tag obliges you to specify a value (either in yaml, or in environment variables).

Reading the config from yaml contradicts the ideology of 12 factors, but in practice, it is more convenient than
reading the entire config from ENV.
It is assumed that default values are in yaml, and security-sensitive variables are defined in ENV.

### `main.go`

Core of this project (DB Connection, Excelize, Filtering and more...).

## Major Packages used in this project

- **postgeSQL go driver**: The Official Golang driver for PostgreSQL.
- **gorm**: The fantastic ORM library for Golang, aims to be developer friendly.
- **viper**: For loading configuration from the `.env` file. Go configuration with fangs. Find, load, and unmarshal a configuration file in JSON, TOML, YAML, HCL, INI, envfile, or Java properties formats.
- **bcrypt**: Package bcrypt implements Provos and Mazières's bcrypt adaptive hashing algorithm.
- **testify**: A toolkit with common assertions and mocks that plays nicely with the standard library.
- Check more packages in `go.mod`.

## Contributing

We welcome contributions to Orion! To contribute to the project, please follow these steps:

- Fork the repository.
- Create a new branch for your feature or bug fix.
- Make your changes and ensure that the tests pass.
- Commit your changes and push them to your fork.
- Submit a pull request to the main repository, describing your changes in detail.
- Please review the Contribution Guidelines for more information.
