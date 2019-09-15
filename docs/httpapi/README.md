# package httpapi

## GET /api/v1/ws

This is a Websocket endpoint that able to execute operation it receives.
The operation correspond with the HTTP API endpoint operation names.
For more you can check the swagger documentation

### operation IsFeatureEnabled

Get release flag status for a pilot

#### Example

request
```json
{
  "operation": "IsFeatureEnabled",
  "data": {
    "feature": "my-feature",
    "id":"public-pilot-uniq-id"
  }
}
```

response
```json
{"enrollment": true}
```

### operation IsFeatureGloballyEnabled

Get release flag status for global

#### Example

request
```json
{
  "operation": "IsFeatureEnabled",
  "data": {"feature": "my-feature"}
}
```

response
```json
{"enrollment": true}
```

## GET /api/v1/rollout/config.json

This endpoint able to answer multiple release flag state for a specific pilot.

### Example

request
```json
{
  "id": "public-pilot-uniq-id",
  "features": [
    "flag-name-a",
    "flag-name-b",
    "flag-name-c"
  ]
}
```

response
```json
{
  "states": {
    "flag-name-a": true,
    "flag-name-b": false,
    "flag-name-c": true
  }
}
```

## GET /api/v1/rollout/is-feature-enabled.json

Get release flag status for a pilot

### Example

request
```json
{
  "feature": "my-feature",
  "id":"public-pilot-uniq-id"
}
```

response
```json
{"enrollment": true}
```

## GET /api/v1/rollout/is-feature-globally-enabled.json

Get release flag status for global

### Example

request
```json
{"feature": "my-feature"}
```

response
```json
{"enrollment": true}
```
