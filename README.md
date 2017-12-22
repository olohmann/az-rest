# az-rest [![Build Status](https://travis-ci.org/olohmann/az-rest.png?branch=master)](https://travis-ci.org/olohmann/az-rest)

`az-rest` is a simple extension to the Azure CLI 2.0 client that allows the execution of [Azure Resource Manager REST API](https://docs.microsoft.com/en-us/rest/api/resources/) calls.

The client re-uses the existing az CLI authentication and authorization scheme via `az account get-access-token`.

It supports JMESPath expressions to filter and project the results from REST calls.

## Installation
### Prerequisites

Install Azure CLI 2.0. Follow the [official installation guide](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli?view=azure-cli-latest).

### Installation on Linux, MacOS X, and Windows

`az-rest` is distributed as a single binary file. The [latest release](https://github.com/olohmann/az-rest/releases/latest) can be found in the releases section. 
Just download, unzip and put it in a directory that is registered in your system's `PATH` environment.

## Usage

```bash
Usage: az-rest [-v] COMMAND [arg...]

A simple Azure Resource Manager REST client.

Options:
  -v, --verbose   Verbose output mode

Commands:
  GET             Issue a GET request
  POST            Issue a POST request
  version         Print version information

Run 'az-rest COMMAND --help' for more information on a command.
```

## Example

```bash
# Login to Azure via az - az-rest will re-use the session.
az login
# ... follow the instructions to finish the log-in

# query the ARM REST API for specific information such as the primary key for AzureSearch
az-rest POST --api-version "2015-08-19" --query "primaryKey" /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/demo-resource-group/providers/Microsoft.Search/searchServices/my-search-service/listAdminKeys

# Response: "00000D607C080850082B30000110DF02"
```

## Acknowledgements

The build process from [terraform](https://github.com/hashicorp/terraform) was a great inspiration for `az-rest`'s build process.

