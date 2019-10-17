![image](https://user-images.githubusercontent.com/12525476/66127607-615c6b00-e5ec-11e9-81e3-764b8bf5b8a5.png)
# Orsum Inflandi

[![Maintainability](https://api.codeclimate.com/v1/badges/213e7b22133dc6c11cc4/maintainability)](https://codeclimate.com/github/orsa-scholis/orsum-inflandi/maintainability)
![License](https://img.shields.io/github/license/orsa-scholis/orsum-inflandi.svg) ![Repo size](https://img.shields.io/github/repo-size/orsa-scholis/orsum-inflandi-II.svg)
[![DeepScan grade](https://deepscan.io/api/teams/5605/projects/7437/branches/75399/badge/grade.svg)](https://deepscan.io/dashboard#view=project&tid=5605&pid=7437&bid=75399)[![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)
[![BCH compliance](https://bettercodehub.com/edge/badge/orsa-scholis/orsum-inflandi?branch=develop)](https://bettercodehub.com/)
[![Backend CI Status](https://github.com/orsa-scholis/orsum-inflandi/workflows/Backend%20CI/badge.svg)](https://github.com/orsa-scholis/orsum-inflandi/actions)
[![Ruby CLI CI Status](https://github.com/orsa-scholis/orsum-inflandi/workflows/Ruby/badge.svg)](https://github.com/orsa-scholis/orsum-inflandi/actions)

Unnecessarily complicated and over-engineered 4 in a row game built with a Kotlin server and React/TypeScript client server

### CLI Usage

```bash
cli/app COMMAND [OPTIONS]
```

#### Available commands:

##### Start

```text
Usage:
  cli/app start

Options:
  [--backend], [--no-backend]              # Default: true
  [--frontend], [--no-frontend]            # Default: true
  [--force-build], [--no-force-build]      
  [--dev], [--no-dev]                      # Default: false
  [--dual-frontend], [--no-dual-frontend]  # Default: false

Start the app or a part of it
```

##### Help


```text
Usage:
  cli/app help [COMMAND]

Describe available commands or one specific command
```

##### Mock

```text
Usage:
  app mock

Options:
  [--port=N]  # Default: 4560

Start a dummy mock server to test frontend
```
