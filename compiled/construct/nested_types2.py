from construct import *
from construct.lib import *

nested_types2__subtype_a__subtype_c__subtype_d = Struct(
	'value_d' / Int8sb,
)

nested_types2__subtype_a__subtype_c = Struct(
	'value_c' / Int8sb,
	'typed_here' / LazyBound(lambda: nested_types2__subtype_a__subtype_c__subtype_d),
	'typed_parent' / LazyBound(lambda: nested_types2__subtype_a__subtype_cc),
	'typed_root' / LazyBound(lambda: nested_types2__subtype_b),
)

nested_types2__subtype_a__subtype_cc = Struct(
	'value_cc' / Int8sb,
)

nested_types2__subtype_a = Struct(
	'typed_at_root' / LazyBound(lambda: nested_types2__subtype_b),
	'typed_here1' / LazyBound(lambda: nested_types2__subtype_a__subtype_c),
	'typed_here2' / LazyBound(lambda: nested_types2__subtype_a__subtype_cc),
)

nested_types2__subtype_b = Struct(
	'value_b' / Int8sb,
)

nested_types2 = Struct(
	'one' / LazyBound(lambda: nested_types2__subtype_a),
	'two' / LazyBound(lambda: nested_types2__subtype_b),
)

_schema = nested_types2
