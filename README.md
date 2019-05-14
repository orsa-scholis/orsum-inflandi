# Orsum-inflandi-II

[![Build Status](https://semaphoreci.com/api/v1/orsa-scholis/orsum-inflandi-ii/branches/develop/shields_badge.svg)](https://semaphoreci.com/orsa-scholis/orsum-inflandi-ii)
[![Maintainability](https://api.codeclimate.com/v1/badges/d853daa69ca35eb79268/maintainability)](https://codeclimate.com/github/orsa-scholis/orsum-inflandi-II/maintainability)
[![Known Vulnerabilities](https://snyk.io/test/github/orsa-scholis/orsum-inflandi-II/badge.svg?targetFile=frontend%2Fpackage.json)](https://snyk.io/test/github/orsa-scholis/orsum-inflandi-II?targetFile=frontend%2Fpackage.json)
![License](https://img.shields.io/github/license/orsa-scholis/orsum-inflandi-II.svg) ![Repo size](https://img.shields.io/github/repo-size/orsa-scholis/orsum-inflandi-II.svg)
[![DeepScan grade](https://deepscan.io/api/teams/3648/projects/5377/branches/41232/badge/grade.svg)](https://deepscan.io/dashboard#view=project&tid=3648&pid=5377&bid=41232)
[![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com)

Unnecessarily complicated and over-engineered 4 in a row game built with a Go server and React/TypeScript client server

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
