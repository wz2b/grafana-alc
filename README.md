# Grafana ALC Data Source

A Grafana data source plugin for retrieving data from an Automated Logic
WebCTRL building automation system.

The plugin provides a Grafana backend data source that communicates with
WebCTRL using its SOAP web services. It allows WebCTRL data to be queried and
displayed using normal Grafana dashboards and visualizations.

## Features

Current functionality includes:

- WebCTRL SOAP authentication
- Historical trend data queries
- Present Value queries
- Browsing the WebCTRL geographic tree
- Grafana backend health checks
- Caching of frequently requested values
- Cache statistics for troubleshooting and diagnostics

The data source configuration requires the URL of the WebCTRL SOAP service and
a WebCTRL username and password.

## WebCTRL SOAP Limitations

There are some important limitations in the WebCTRL SOAP interface, particularly
when retrieving historical trend data.

WebCTRL trend histories contain records other than actual point samples,
including events such as:

- time synchronization records
- error records
- other trend/history metadata records

The WebCTRL SOAP interface does not reliably distinguish or filter these records
from normal trend samples when returning data.

As a result, these records may appear in Grafana as data points, frequently
with a value of zero. These apparent zero values are not necessarily real
measurements from the underlying point.

This can be confusing when viewing trend data and should be kept in mind when
interpreting graphs produced by this plugin.

This behavior originates in the data returned by the WebCTRL SOAP interface,
rather than Grafana itself.

## Future Development

The SOAP interface is increasingly limiting for this application.

Future versions of this plugin are expected to use **Cameron Vogt's WebCTRL REST
API** instead of the native WebCTRL SOAP interface. The REST API provides a
better foundation for retrieving and interpreting WebCTRL data and should allow
several of the limitations described above to be eliminated.

The current SOAP implementation will remain useful for existing installations
and as a reference implementation.

## Architecture

The plugin consists of:

- a React/TypeScript Grafana frontend
- a Go Grafana backend plugin
- the [`webctrl-soap-go`](https://github.com/wz2b/webctrl-soap-go) Go library for
  communication with WebCTRL

The backend performs WebCTRL requests and returns Grafana DataFrames to the
frontend.

## Status

This project is under active development.

It was originally developed for older versions of Grafana and has been updated
to work with current Grafana releases.

## License

See [LICENSE](LICENSE).