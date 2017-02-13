// Package desc contains "rich descriptors" for protocol buffers. The built-in
// descriptor types are simple protobuf messages, each one representing a
// different kind of element in the AST of a .proto source file.
//
// Because of this inherent "tree" quality, a given message descriptor cannot
// refer to its enclosing file descriptor. Nor can a field descriptor refer to a
// message or enum descriptor that represents the field's type (for enum and
// nested message fields). This limitation makes them much harder to use for
// doing interesting things with reflection.
//
// Particularly tricky is resolving references to types -- a field's type, the
// message type an extension extends, the request and response types of an RPC
// method, etc. They are simply strings, and can even be relative type
// references, with their own lexical scoping and resolution rules.
//
// "Rich descriptors" avoid the need to deal with the complexities described
// above (like resolving type references). A rich descriptor has all type
// references resolved and provides methods to access other rich descriptors for
// all referenced elements. The rich descriptor does not try to mimic the full
// interface of the underlying descriptor proto. Instead, every rich descriptor
// provides access to that underlying proto.
//
// Rich descriptors can be accessed in much the same way that their "poor"
// cousins (descriptor protos) are accessed. Instead of using proto.FileDescriptor,
// use desc.LoadFileDescriptor.
//
// It is also possible create rich descriptors for proto messages that a given
// Go program doesn't even know about. For example, they could be loaded from a
// FileDescriptorSet file (which can be generated by protoc) or loaded from a
// server. This enables interesting things like dynamic clients: where a Go
// program can be an RPC client of a service it wasn't compiled to know about.
package desc
