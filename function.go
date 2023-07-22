package function

import (
	"context"
	"fmt"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

func init() {
	functions.CloudEvent("CloudEventFunc", cloudEventFunc)
}

func cloudEventFunc(ctx context.Context, e cloudevents.Event) error {
	// Do something with event.Context and event.Data (via event.DataAs(foo)).
	fmt.Println("Hello world: ", string(e.Data()))
	return nil
}
