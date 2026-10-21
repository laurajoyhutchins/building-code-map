// Bounded independent BCAST API interoperability probe.
// Source content is not printed and never becomes BCM adoption evidence.
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/laurajoyhutchins/building-code-map/bcastconsumer"
)

func main() {
	var endpoint, packageID string
	flag.StringVar(&endpoint, "base-url", "", "BCAST read-only API HTTP origin")
	flag.StringVar(&packageID, "package-id", "", "expected BCAST 0.2 package SHA-256 identity")
	flag.Parse()
	if endpoint == "" || packageID == "" || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "base-url and package-id are required; no positional args")
		os.Exit(2)
	}
	client, err := bcastconsumer.New(endpoint)
	if err != nil {
		fmt.Fprintln(os.Stderr, "BCAST consumer configuration rejected:", err)
		os.Exit(1)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	pkg, err := client.Fetch(ctx, packageID)
	if err == nil {
		err = client.VerifyNavigation(ctx, pkg)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "BCAST public contract interoperability failed:", err)
		os.Exit(1)
	}
	fmt.Printf("BCM_BCAST_HTTP_CONSUMER_VERIFIED package_id=%s objects=%d\n", pkg.ID, len(pkg.Objects))
}
