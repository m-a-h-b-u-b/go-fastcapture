# go-fastcapture

[![License:
Apache-2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)

High‑Speed Packet Capture in Go (10‑40 Gbps) --- A production‑grade
framework using **DPDK** or **PF_RING** for zero‑copy packet capture,
real‑time filtering, and low‑latency processing. Designed for network
analytics, intrusion detection, and performance monitoring.
([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))

------------------------------------------------------------------------

## Table of Contents

-   [Features](#features)\
-   [Use Cases](#use‑cases)\
-   [Prerequisites](#prerequisites)\
-   [Installation](#installation)\
-   [Usage](#usage)\
-   [Architecture](#architecture)\
-   [Examples](#examples)\
-   [Contributing](#contributing)\
-   [License](#license)

------------------------------------------------------------------------

## Features

-   Zero‑copy packet capture via **DPDK** or **PF_RING**
    ([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))\
-   Very high throughput (10‑40 Gbps) capture support
    ([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))\
-   Real‑time filtering to discard unneeded packets early
    ([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))\
-   Low latency processing for performance‑critical networking
    applications
    ([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))

------------------------------------------------------------------------

## Use Cases

-   Network analytics / monitoring\
-   Intrusion detection systems (IDS)\
-   Performance monitoring / network diagnostics\
-   Any application requiring high rate packet capture with minimal
    overhead

------------------------------------------------------------------------

## Prerequisites

To build and run go‑fastcapture, you'll need:

-   Go (version compatible with modules; check `go.mod`)\
-   Linux system with support for DPDK **or** PF_RING\
-   Proper kernel and network interface setup for zero‑copy mode\
-   Root or appropriate privileges to access network interface,
    hugepages (if using DPDK), etc.

------------------------------------------------------------------------

## Installation

1.  Clone this repository:

    ``` bash
    git clone https://github.com/m-a-h-b-u-b/go-fastcapture.git
    cd go-fastcapture
    ```

2.  Install dependencies:

    ``` bash
    go mod download
    ```

3.  Build the main binary:

    ``` bash
    cd cmd/hspc
    go build -o hspc
    ```

4.  (If needed) Compile or configure DPDK / PF_RING on your system,
    following their respective setup instructions.

------------------------------------------------------------------------

## Usage

Run the capture tool:

``` bash
sudo ./hspc [options]
```

Some possible command‑line options (example):

-   `--interface <name>`: network interface to capture from\
-   `--mode dpdk|pfring`: which backend to use\
-   `--filter <expression>`: packet filter (BPF‑like) to drop unwanted
    traffic early\
-   `--output <path>`: where to store or forward captured data

> **Note:** DPDK may require special kernel modules, hugepages setup,
> binding NICs to DPDK drivers, etc. PF_RING similarly has its own
> kernel or module configuration.

------------------------------------------------------------------------

## Architecture

-   `internal/capture`: core packet capture logic and integration with
    DPDK or PF_RING\
-   `cmd/hspc`: command‑line interface / main entry point for users\
-   `examples/throughput`: sample programs or benchmarks to test
    throughput and performance
    ([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))\
-   `scripts`: helper scripts (e.g. for setup, benchmarking)

------------------------------------------------------------------------

## Examples

Here is an example to test throughput with a simple filter:

``` bash
sudo ./hspc --mode dpdk --interface eth0 --filter "ip and tcp" --output /tmp/capture.pcap
```

Another example using PF_RING:

``` bash
sudo ./hspc --mode pfring --interface eth1 --filter "udp" --output /tmp/udp_dump.pcap
```

------------------------------------------------------------------------

## Contributing

Contributions are welcome! Here's how you can help:

1.  Fork the repo.\
2.  Create a feature branch: `git checkout ‑b feature/your‑feature`\
3.  Write tests for new functionality.\
4.  Ensure documentation / examples are updated.\
5.  Submit a pull request.

Please adhere to code style, keep performance in mind, and ensure
compatibility with both DPDK and PF_RING backends.

------------------------------------------------------------------------

## License

This project is licensed under the **Apache‑2.0 License**. See the
[LICENSE](LICENSE) file for details.
([github.com](https://github.com/m-a-h-b-u-b/go-fastcapture))

------------------------------------------------------------------------

## Contact / Author

Maintained by **m‑a‑h‑b‑u‑b**. Please open issues for bugs or feature
requests.

------------------------------------------------------------------------
