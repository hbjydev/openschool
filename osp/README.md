# OSP: OpenSchool Protocol

The OSP is the protocol driving most inter-service communications.

The specification is loosely based on the HTTP/1 specification, facilitating all
unary resource calls.

## Request Line

The request line consists of three components:

- `action`: The action to perform on the provided `osrn`.
- `osrn`: The [OSRN](../osrn/README.md) to request against.
- `version`: The version of OSP this request is being made with.

## Example Request

```http
LIST osrn::classes:* OSP/1.0
```
