# meeko-poblano-directory #

Meeko agent that exports Poblano Directory API calls as native Meeko RPC methods.

## Agent Configuration ##

There are three Meeko variables that must be set:

* `API_BASE_URL` - the URL where the API can be found
* `API_TOKEN` - the access token to use for authentication
* `RPC_TOKEN` - the token that must be passed as an argument in all RPC requests

## Agent Interface ##

There are two methods being exported:

* `PoblanoDirectory@1.GetUser` - execute `GET API_BASE_URL/users?QUERY_STRING`
  where `QUERY_STRING` is the `query` argument.
* `PoblanoDirectory@1.GetProject` - the same as `PoblanoDirectory@1.GetUser`
  except the fact that `/projects` is used in the Poblano Directory API call.

### Arguments and Return Values ###

See the `methods` package, it's really small.

## Contributing ##

See `CONTRIBUTING.md`.

## License ##

MIT, see the `LICENSE` file.

## Original Authors ##

[tchap](https://github.com/tchap) on behalf of [Salsita](https://github.com/salsita)
