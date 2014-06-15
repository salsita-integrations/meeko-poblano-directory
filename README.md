# meeko-poblano-directory #

Meeko agent that exports Poblano Directory calls as native Meeko RPC methods.

## Agent Configuration ##

There are two Meeko variables that must be set:

* `API_BASE_URL` - the URL where the API can be found
* `API_TOKEN` - the access token to use for authentication

## Agent Interface ##

There are two methods being exported:

* `Poblano.GetUser` - execute `GET API_BASE_URL/users?QUERY_STRING` where `QUERY_STRING`
  is the only argument that this method needs. The arguments object is thus `{"query":
  <string>}`. The return value is exactly what was returned by the API call.
* `Poblano.GetProject` - the same as `Poblano.GetUser` except the fact that
  `/projects` is used in the Poblano Directory API call.

## Contributing ##

See `CONTRIBUTING.md`.

## License ##

MIT, see the `LICENSE` file.

## Original Authors ##

[tchap](https://github.com/tchap) on behalf of [Salsita](https://github.com/salsita)
