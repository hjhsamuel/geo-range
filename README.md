# geo-range
A library with encoding/decoding geohashes and searching geohashes by coordinates and radius

## Usage

```shell
go get -u github.com/hjhsamuel/geo-range
```

### GeoHash

Example Code:

```go
hashes := geo_range.RadiusSearch(30.2027354, 120.22543349, 5000, 0.8, nil)
combined := geo_range.MergeHashes(hashes, 20)
```

> Note:
> 
> Geohash range searches are imprecise, sometimes you should implement `PrecisionDynamicFunc` and use your own function


### Coordinate

If you prefer not to use GeoHash-based tiling for retrieval, you can also use a simple bounding box (MBR) match followed by distance filtering. This approach also delivers excellent performance in MySQL 8.0+.

```sql
create table tb_marker (
    id biging unsigned primary key auto_increment,
    geohash char(12) not null,
    location point not null
) charset = utf8mb4;

create index idx_geohash on tb_marker (geohash);

create spatial index idx_location on tb_marker (location);
```

Example Code:

```go
point := &geo_hash.Location{Lng: 120.22543349, Lat: 30.2027354}
northPoint := geo_hash.GetPointAtDistance(point, 5000, 0)
eastPoint := geo_hash.GetPointAtDistance(point, 5000, 90)
southPoint := geo_hash.GetPointAtDistance(point, 5000, 180)
westPoint := geo_hash.GetPointAtDistance(point, 5000, 270)

rightTop := &Location{Lng: eastPoint.Lng, Lat: northPoint.Lat}
leftBottom := &Location{Lng: westPoint.Lng, Lat: southPoint.Lat}
```

then you could use MBR match in MySQL:

```sql
select * from tb_marker where MBRIntersects(ST_GEOMFROMTEXT('LINESTRING(lat1 lng1, lat2 lng2)', 4326), location)
```

of course, you could use `ST_DISTANCE_SPHERE` for further precise filtering.

## Specialized scenarios

1. When the search range crosses the antimeridian (180th meridian), you must split the bounding box into two separate rectangles.
2. When the search range encompasses a pole, the longitude span collapses to [-180, 180]. Since a traditional MBR becomes semantically invalid or excessively bloated in this case, it is recommended to constrain the search scope and perform the retrieval based solely on the latitude range.

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
