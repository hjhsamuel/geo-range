# geo-range
A library with encoding/decoding geohashes and searching geohashes by coordinates and radius

## Usage

```shell
go get -u github.com/hjhsamuel/geo-range
```

Example Code:

```go
hashes := geo_range.RadiusSearch(30.2027354, 120.22543349, 5000, nil)
```

> Note:
> Geohash range searches are imprecise, sometimes you should implement `PrecisionDynamicFunc` and use your own function

## Digits and precision

| geohash length | Lat bits | lng bits | lat error   | lng error   | km error   |
|----------------|----------|----------|-------------|-------------|------------|
| 1              | 2        | 3        | ±23         | ±23         | ±2500      |
| 2              | 5        | 5        | ±2.8        | ±5.6        | ±630       |
| 3              | 7        | 8        | ±0.70       | ±0.7        | ±78        |
| 4              | 10       | 10       | ±0.087      | ±0.18       | ±20        |
| 5              | 12       | 13       | ±0.022      | ±0.022      | ±2.4       |
| 6              | 15       | 15       | ±0.0027     | ±0.0055     | ±0.61      |
| 7              | 17       | 18       | ±0.00068    | ±0.00068    | ±0.076     |
| 8              | 20       | 20       | ±0.000086   | ±0.000172   | ±0.01911   |
| 9              | 22       | 23       | ±0.000021   | ±0.000021   | ±0.00478   |
| 10             | 25       | 25       | ±0.00000268 | ±0.00000536 | ±0.0005971 |
| 11             | 27       | 28       | ±0.00000067 | ±0.00000067 | ±0.0001492 |
| 12             | 30       | 30       | ±0.00000008 | ±0.00000017 | ±0.0000186 |
