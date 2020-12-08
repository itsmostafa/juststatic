# Just Static

A very simple Static Site Generator. 

- Create a `templates/{your-template-name}` directory
- Copy your go template files into `templates/{your-template-name`
- Build & Run
    - `go build && ./juststatic -t {your-template-name}`
- Your generated static files will live in the `public` directory