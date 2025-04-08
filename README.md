# AWS SSO ECR Docker Credential Helper

The AWS SSO ECR Docker Credential Helper is a credential helper for the Docker daemon that makes it easier to use [Amazon Elastic Container Registry](https://aws.amazon.com/ecr/) with AWS SSO credentials.

> [!NOTE]
> This project is not affiliated with Amazon Web Services, but it is heavily inspired by [Amazon ECR Docker Credential Helper](https://github.com/awslabs/amazon-ecr-credential-helper)

## Table of Contents
  * [Prerequisites](#prerequisites)
  * [Installing](#installing)
    + [macOS](#macos)
    + [Linux](#linux)
    + [Windows](#windows)
    + [From Source](#from-source)
  * [Configuration](#configuration)
    + [Docker](#docker)
    + [AWS credentials](#aws-credentials)
  * [Usage](#usage)
  * [Troubleshooting](#troubleshooting)
  * [License](#license)

## Prerequisites

You must have at least Docker 1.11 installed on your system.

You also must have AWS credentials available. See the [AWS credentials section](#aws-credentials) for details on how to use different AWS credentials.

## Installing

### macOS

A Homebrew formula is available in a custom tap.

```bash
brew tap kyosenergy/tap
brew install kyosenergy/tap/docker-credential-sso-ecr-login
```

Once you have installed the credential helper, see the [Configuration section](#configuration) for instructions on how to configure Docker to work with the helper.

### Linux

Linux executables are available via [GitHub releases](https://github.com/kyosenergy/docker-credential-sso-ecr-login/releases).

Download the latest release, extract the binary, place it in your `PATH` (usually `/usr/local/bin`), and make it executable:

```bash
chmod +x /usr/local/bin/docker-credential-sso-ecr-login
```

Once you have installed the credential helper, see the [Configuration section](#configuration) for instructions on how to configure Docker to work with the helper.

### Windows

Windows executables are available via [GitHub releases](https://github.com/kyosenergy/docker-credential-sso-ecr-login/releases).

Once you have installed the credential helper, see the [Configuration section](#configuration) for instructions on how to configure Docker to work with the helper.

### From Source

To build and install the AWS SSO ECR Docker Credential Helper, we suggest Go 1.19 or later, `git` and `make` installed on your system.

If you just installed Go, make sure you also have added it to your PATH or Environment Vars (Windows). For example:

```
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```

Or in Windows:

```
setx GOPATH %USERPROFILE%\go
<your existing PATH definitions>;%USERPROFILE%\go\bin
```

If you haven't defined the PATH, the command below will fail silently, and running `docker-credential-sso-ecr-login` will output: `command not found`

You can install this via the `go` command line tool.

To install, you can run:

```
go install github.com/kyosenergy/docker-credential-sso-ecr-login
```

## Configuration

### Docker

There is no need to use `docker login` or `docker logout`.

Place the `docker-credential-sso-ecr-login` binary on your `PATH`.
On Windows, depending on whether the executable is run in the User or System context, the corresponding `Path` user or system variable needs to be used.

Following that, the configuration for the Docker client needs to be updated in `~/.docker/config.json` to use the **sso-ecr-login** helper.
Depending on the operating system and context under which docker client will be executed, this configuration can be found in different places.
  
On macOS and Linux systems:
- `/home/<username>/.docker/config.json` for **user** context
- `/root/.docker/config.json` for **root** context
  
On Windows:
- `C:\Users\<username>\.docker\config.json` for **user** context
- `C:\Windows\System32\config\systemprofile\.docker\config.json` for the **SYSTEM** context

Create a `credHelpers` section with the URI of your ECR registry:

```json
{
	"credHelpers": {
		"<aws_account_id>.dkr.ecr.<region>.amazonaws.com": "sso-ecr-login"
	}
}
```

### AWS credentials

The AWS SSO ECR Docker Credential Helper allows you to use AWS credentials retrieved from AWS SSO.

You can configure your AWS SSO via the `aws configure sso` command, which will create a profile in your AWS config file.

For example:

```ini
[profile default]
sso_start_url = https://mycompany.awsapps.com/start/#
sso_region = us-west-1
sso_account_id = 123456789012
sso_role_name = MyAssumedRole
region = us-west-1
output = json
```

## Usage

`docker pull 123456789012.dkr.ecr.us-west-1.amazonaws.com/my-repo:my-tag`

If you have configured additional profiles for use with the AWS CLI, you can use those profiles by specifying the `AWS_PROFILE` environment variable when invoking `docker`.
For example:

`AWS_PROFILE=myprofile docker pull 123456789012.dkr.ecr.us-west-1.amazonaws.com/my-repo:my-tag`

There is no need to use `docker login` or `docker logout`.

## Troubleshooting

Logs from the AWS SSO ECR Docker Credential Helper are stored in `~/.sso-ecr-login/logs`.

## License

The AWS SSO ECR Docker Credential Helper is licensed under the MIT License.
