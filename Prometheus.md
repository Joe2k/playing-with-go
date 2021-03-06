# Prometheus

## Metric Types

### Counters

Counters are a simple metric type that can only be incremented or be reset to zero on restart. It is often used to count primitive data like the total number of requests to a services or number of tasks completed. The absolute value of these counters is often irrelevant and does not give you much information about the applications state. The real information can be gathered by their evolution over time which can be obtained using the rate() function.

### Gauges

Gauges also represent a single numerical value but different to counters the value can go up as well as down. Therefore gauges are often used for measured values like temperature, humidity or current memory usage. Unlike with counters the current value of a gauge is meaningful and can be directly used in graphs and tests.

### Histograms

Histograms are used to measure the frequency of value observations that fall into specific predefined buckets. This means that they will provide information about the distribution of a metric like response time and signal outliers. By default Prometheus provides the following buckets: .005, .01, .025, .05, .075, .1, .25, .5, .75, 1, 2.5, 5, 7.5, 10. These buckets are not suitable for every measurement and can therefore easily be changed.

### Summaries

Summaries are very similar to Histograms because they both expose the distribution of a given data set. The one major difference is that a Histogram estimate quantiles on the Prometheus server while Summaries are calculated on the client side. Summaries are more accurate for some pre-defined quantiles but can be a lot more resource expensive because of the client-side calculations. That is why it is recommended to use Histograms for most use-cases.

## Resources

[Prometheus with Go and Grafana](https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang)
