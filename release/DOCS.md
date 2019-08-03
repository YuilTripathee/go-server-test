# Operational documentation

## Running servers

Firing up server in  development environment:

```c
./<server_binary> <arg1:DEV> >> <log_file>.log
// no LOG needed for debuggnig purpose but can be done over the need.
```

Firing up server in production environment:

```c
./<server_binary> <arg1:PROD> >> <log_file>.log
// LOG seems to be optional here also.
```

> **Note**:  
> Double debugging can be done by supplying custom port by setting up a command line flag with ease at any instant.

## About the log format

### LOG for DEVELOPMENT

Log output for debugging includes only the essential a backend developer may need at development phase of software development.

```cs
app_name-environment timestamp | status_code | protocol | client_ip | method path
```

Example:

```cs
[APP-DEBUG] 2019/04/25 - 00:06:12 | 300 | HTTP/1.1 |  127.0.0.1 | DELETE  /v1/app/
```

### LOG for PRODUCTION

Log output for production includes all the data essential for application analytics, tailored specifically to be processable is easy to use formats like CSV (and JSON as well). The logger for production is worth storing in a LOG for further appliments.

```cs
app_name-environment | host | time_custom | status | latency_human | remote_ip | bytes_in bytes_in | bytes_out bytes_out | method | uri
```

Example:

```cs
APP-PROD | 127.0.0.1:8000 | 2019/04/25 00:06:12 | 300 | 215.117µs | 127.0.0.1 | 0 bytes_in | 41 bytes_out | GET | /v1/app/
```

## Encoding log to analyzable CSV format

Algorithm to do so:

1. GET the CSV file with custom delimeter
2. Strip all the JSON based log printouts onto a separate file.
3. Relace all commas in a file (use URI decoding or Unicode decoding in advance).
4. Remove all the unwanted strings, texts, spaces from the file.
5. Replace all pipes '|' with comma, since the pipes are used to as a delimeter.

> **Note:**  
> Python is the best suited scripting language for this purpose. Although, GO or C++ can be applied on binary level.

## About saving binary log on linux machine

If you run binary on linux machine you could use shell script.

Overwrite into a file

```shell
./binaryapp > binaryapp.log
```

append into a file

```shell
./binaryapp >> binaryapp.log
```

overwrite stderr into a file

```shell
./binaryapp &> binaryapp.error.log
```

append stderr into a file

```shell
./binaryapp &>> binalyapp.error.log
```

It can be more dynamic using shell script file.  
Reference : [Solution from StackOverflow](https://stackoverflow.com/questions/19965795/go-golang-write-log-to-file)