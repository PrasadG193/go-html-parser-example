# go-html-parser-example
Go example to parse list of const from the given documentation page

This example scrapes package documentation page - https://pkg.go.dev/k8s.io/component-helpers/storage/volume
and parses list of constants declared in that package

## Running example

```
$ go run main.go 
List of consts declared in "k8s.io/component-helpers/storage/volume":
AnnBindCompleted 
AnnBoundByController 
AnnSelectedNode 
NotSupportedProvisioner 
AnnDynamicallyProvisioned 
AnnMigratedTo 
AnnStorageProvisioner     
AnnBetaStorageProvisioner 
PVDeletionProtectionFinalizer 
PVDeletionInTreeProtectionFinalizer 

```
