# go-whosonfirst-riordan

Go package providing a tool to fetch Who's On First data from a GitHub organization and write it all to a GeoJSON FeatureCollection using the [whosonfirst/go-whosonfirst-iterwriter](https://github.com/whosonfirst/go-whosonfirst-iterwriter), [whosonfirst/go-whosonfirst-iterate-organization](https://github.com/whosonfirst/go-whosonfirst-iterate-organization/) and [whosonfirst/go-writer-featurecollection](https://github.com/whosonfirst/go-writer-featurecollection/) packages.

## Example

```
$> go build -mod vendor -ldflags="-s -w" -o bin/riordan cmd/riordan/main.go
$> ./bin/riordan \
	-writer-uri 'constant://?val=featurecollection://?writer=stdout://' \
	-iterator-uri org:// \
	'sfomuseum-data://?prefix=sfomuseum-data-architecture' \

	> architecture.geojson
```

## See also

* https://github.com/whosonfirst/go-whosonfirst-iterwriter/
* https://github.com/whosonfirst/go-whosonfirst-iterate-organization/
* https://github.com/whosonfirst/go-writer-featurecollection/
