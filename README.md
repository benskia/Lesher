# Lesher

Manage Lenovo battery thresholds using power_supply class on Linux. List,
create, and delete charging profiles. Get the current health of each battery as
a percentage remaining of the battery's full-charge design specification.

* [Issues](#issues)
* [Install](#install)
* [Usage](#usage)

## Issues

* Setting profiles quickly results in hangs and failed or corrupted writes. I
suspect this has to do with those files being accessed after changes are
detected.

## Install

Note: Currently only supports Linux platforms that implement the power_supply class.

1. [Download and Install GO](https://go.dev/doc/install)
2. Add your user's go/bin (usually '$HOME/go/bin') to PATH.
3. [Clone this Repo](git@github.com:benskia/Lesher.git)
4. Install or build the project.

Install:

``` go
git clone git@github.com:benskia/Lesher.git
cd Lesher/cmd/Lesher
go install
```

Build:

``` go
git clone git@github.com:benskia/Lesher.git
cd Lesher
make build
./bin/Lesher
```

## Usage

`Lesher <command> [options]`

### Display Lesher documentation

* `help`

### Display available profiles, current battery thresholds, and charge status

* `list`

### Display battery full-charge stats

* `health`

### Create or update a profile with start and end thresholds

* `create <profile-name> <charge-start> <charge-end>`

### Delete a profile

* `delete <profile-name>`

### Active a profile by writing its start and end thresholds to file

* `set <profile-name>`
