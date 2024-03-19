// `option` provides a flexible and expressive framework for configuring objects
// in Go using the functional options pattern. This pattern offers an intuitive
// way to customize objects with clear, readable code. It's particularly useful
// for APIs or libraries where the number of configuration parameters might
// change over time, offering a scalable and maintainable approach.
//
// For an in-depth understanding of this design pattern, consider exploring the
// following resources:
//   - https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html
//   - https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
//   - https://amirsoleimani.medium.com/functional-options-in-go-with-generic-863dbd68cc6f
//   - https://golang.design/research/generic-option/
//   - https://sagikazarmark.hu/blog/functional-options-on-steroids/
package option

// `Option` represents a configuration option for a type `T`. It's a function
// that modifies `T` in some way and returns an error if the modification fails.
// This allows options to perform validation or conditional application based on
// the state of `T` or other criteria.
type Option[T any] func(v *T) error

// `Apply` applies a series of options to a given instance of type `T`. It
// iterates through each option, applying them to `T`. If any option fails
// (indicated by returning an error), `Apply` halts and returns the encountered
// error.
func Apply[T any](v *T, opts ...Option[T]) error {
	for _, opt := range opts {
		if err := opt(v); err != nil {
			return err
		}
	}
	return nil
}

// `Group` combines multiple options into a single, composite option. This is
// particularly useful for organizing and reusing sets of related options. When
// the returned `Option[T]` is applied to an instance of `T`, it sequentially
// applies each grouped option, stopping and returning an error if any single
// option fails.
func Group[T any](opts ...Option[T]) Option[T] {
	return func(v *T) error {
		return Apply(v, opts...) // Leverage Exec to apply all grouped options.
	}
}
