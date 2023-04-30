$version: "2"
namespace example.hello

use aws.auth#sigv4
use aws.api#service
use aws.protocols#restJson1

@restJson1
@service(sdkId: "bmoffatt")
@sigv4(name: "lambda")
service Hello {
    version: "2023-04-28"
    operations: [Wave]
}

@http(method: "POST", uri: "/wave", code: 200)
operation Wave {
    input := {
        name: String
    }
    output := {
        @jsonName("Text")
        text: String
        @jsonName("Time")
        @timestampFormat("date-time")
        time: Timestamp
    }
}

