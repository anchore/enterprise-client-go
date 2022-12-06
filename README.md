# A Golang client for the Anchore Enterprise API

Modified from https://github.com/anchore/client-go. 
Based on generated code. No executables provided, only intended for use as a library in other projects.

To pull the swagger definition and re-generate all client go code, update `ENTERPRISE_REF` in the Makefile to point to
the desired enterprise commit. Then run the following to regenerate the client:
```bash
make 
```

Note: the version of the enterprise where the swagger spec is pulled from is pinned in the Makefile. 

## Release Strategy:

Currently releases are manual. This table shows the released versions of the client and it's associated enterprise version:

| client version | enterprise version |
|----------------|--------------------|
| v1.0.0         | v4.3.0             |

If breaking changes are introduced past this point, please follow this guide regarding how to release new
major versions of a Golang library:

[Module version numbering](https://go.dev/doc/modules/version-numbers)
[Module release and versioning](https://go.dev/doc/modules/release-workflow)
[Managing module source](https://go.dev/doc/modules/managing-source)

## Known Issues:
[OpenAPITool Bug](DEVELOPER.md)
