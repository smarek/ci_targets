from construct import *
from construct.lib import *

params_pass_struct__block = Struct(
	'foo' / Int8ub,
)

params_pass_struct__struct_type__baz = Struct(
	'qux' / Int8ub,
)

params_pass_struct__struct_type = Struct(
	'bar' / LazyBound(lambda: params_pass_struct__struct_type__baz),
)

params_pass_struct = Struct(
	'first' / LazyBound(lambda: params_pass_struct__block),
	'one' / LazyBound(lambda: params_pass_struct__struct_type),
)

_schema = params_pass_struct
