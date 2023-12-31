# ggt-simple-gin
[English](./README.md) || [中文](./README_zh.md)

ggt-simple-gin is a web framework project that imitates the hand-written implementation of gin. It aims to gain a deeper understanding of the underlying principles and design patterns of gin and ultimately create a simplified version of a web framework called gsg.

The name "ggt" is an abbreviation for "Gallifrey's GoTutoural," while "gsg" is the abbreviation for "ggt-simple-gin."

The project is primarily inspired by the blog posts of [GeekTutu](https://geektutu.com), specifically the tutorial [Building Web Framework Gee from Scratch in 7 Days](https://geektutu.com/post/gee.html) . For more details and considerations regarding the program design, please refer to the original blog.

## Development Plan

- [x] Construct the `Engine` structure to implement the `ServeHTTP` interface of the http package and add its constructor.
- [x] Encapsulate functions in `Engine` for handling GET and POST requests and implement basic networking functionality.
- [x] Extract the `router` to facilitate further feature development.
- [x] Design `Context` to encapsulate Request and Response and provide support for various response types like JSON and HTML.
- [x] Implement dynamic route parsing using a Trie tree.
- [x] Add support for two types of route matching: `:name` and `*filepath`.
- [x] Implement route grouping control.
- [x] Design and implement a middleware mechanism for the web framework.
- [x] Implement a universal `Logger` middleware that records the time taken from the request to the response.
- [x] Implement static resource serving.
- [x] Support HTML template rendering.
- [x] Implement error handling mechanism.
