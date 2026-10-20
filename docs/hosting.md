# Website publishing

The public GitHub Pages site is deployed from this repository at https://laurajoyhutchins.github.io/building-code-map/.

Its four-file artifact contains the existing public HTML, JavaScript, CSS, and one reviewed fictional `DEMO-XX` project-code-basis JSON fixture. The static page displays the fixture and **does not perform real project verification**. Its data is not a new regulatory engine, or a live response from the Go verifier. The separately runnable Go reference app is the interactive canonical synthetic verifier.

Build and inspect the publication locally:

```sh
go test ./...
npm test
npm run validate
npm run build:pages
```

The publishing workflow runs public verification and uploads only `dist/pages/`. It never checks out the private `building-code-map-alpha` repository, and it must not publish the private maintained corpus, restricted legal evidence, geocoder databases, or source artifacts. The publication policy in `docs/publication-policy.md` governs further public additions.

When a production BCM API is available, connecting real lookup will require a separate explicit design and verification change. No real-world regulatory answer should be inferred from the fictional page.
